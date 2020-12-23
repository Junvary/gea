package service

import (
	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
	"golang.org/x/crypto/ssh"
	"io"
	"log"
	"net"
	"net/url"
	"time"
)

// WebSSH 管理 Websocket 和 ssh 连接
type WebSSH struct {
	id string
	buffSize uint32
	term string
	sshConn net.Conn
	websocket *websocket.Conn
	connTimeout time.Duration
	logger *log.Logger
	currentUser string
}

// WebSSH 构造函数
func NewWebSSH() *WebSSH {
	return &WebSSH{
		buffSize: DefaultBuffSize,
		logger:   DefaultLogger,
		term: DefaultTerm,
		connTimeout: DefaultConnTimeout,
	}
}

func (ws *WebSSH) SetLogger(logger *log.Logger) {
	ws.logger = logger
}

// 设置 buff 大小
func (ws *WebSSH) SetBuffSize(buffSize uint32) {
	ws.buffSize = buffSize
}

// 设置日志输出
func (ws *WebSSH) SetLogOut(out io.Writer) {
	ws.logger.SetOutput(out)
}

// 设置终端 term 类型
func (ws *WebSSH) SetTerm(term string) {
	ws.term = term
}

// 设置连接 id
func (ws *WebSSH) SetId(id string) {
	ws.id = id
}

// 设置连接 id
func (ws *WebSSH) SetCurrentUser(currentUser string) {
	ws.currentUser = currentUser
}

// 设置连接超时时间
func (ws *WebSSH) SetConnTimeOut(connTimeout time.Duration) {
	ws.connTimeout = connTimeout
}

// 添加 websocket 连接
func (ws *WebSSH) AddWebsocket(conn *websocket.Conn) {
	ws.logger.Printf("(%s) (%s) ws连接", ws.id, ws.currentUser)
	ws.websocket = conn
	go func() {
		ws.logger.Printf("(%s) (%s) ws退出 %v", ws.id, ws.currentUser, ws.server())
	}()
}

// 添加 ssh 连接
func (ws *WebSSH) AddSSHConn(conn net.Conn) {
	ws.logger.Printf("(%s) (%s) ssh 连接", ws.id, ws.currentUser)
	ws.sshConn = conn
}

// 处理 websocket 连接发送过来的数据
func (ws *WebSSH) server() error {
	defer func(){
		_ = ws.websocket.Close()
	}()

	// 默认加密方式 aes128-ctr aes192-ctr aes256-ctr aes128-gcm@openssh.com arcfour256 arcfour128
	// 连 linux 通常没有问题，但是很多交换机其实默认只提供 aes128-cbc 3des-cbc aes192-cbc aes256-cbc 这些。
	// 因此我们还是加全一点比较好。
	sshConfig := ssh.Config{
		Ciphers: []string{"aes128-ctr", "aes192-ctr", "aes256-ctr", "aes128-gcm@openssh.com", "arcfour256", "arcfour128", "aes128-cbc", "3des-cbc", "aes192-cbc", "aes256-cbc"},
	}

	config := ssh.ClientConfig{
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         ws.connTimeout,
		Config: sshConfig,
	}

	var session *ssh.Session
	var stdin io.WriteCloser
	var hasAddr bool
	var hasLogin bool
	var hasAuth bool
	var hasTerm bool

	for {
		var msg message
		err := ws.websocket.ReadJSON(&msg)
		if err != nil {
			return errors.Wrap(err, "ws关闭或信息类型错误！")
		}
		switch msg.Type {
		case messageTypeCurrentUser:
			if hasAddr {
				continue
			}
			currentUser, _ := url.QueryUnescape(string(msg.Data))
			ws.SetCurrentUser(currentUser)

		case messageTypeAddr:
			if hasAddr {
				continue
			}
			addr, _ := url.QueryUnescape(string(msg.Data))
			ws.logger.Printf("(%s) (%s) 连接到 %s", ws.id, ws.currentUser, addr)
			conn, err := net.DialTimeout("tcp", addr, ws.connTimeout)
			if err != nil {
				_ = ws.websocket.WriteJSON(&message{Type: messageTypeStderr, Data: []byte("connect error\r\n")})
				return errors.Wrap(err, "连接到 " + addr + " 错误！")
			}
			ws.AddSSHConn(conn)
			defer func() {
				_ = ws.sshConn.Close()
			}()
			hasAddr = true
		case messageTypeTerm:
			if hasTerm {
				continue
			}
			term, _ := url.QueryUnescape(string(msg.Data))
			ws.logger.Printf("(%s) set term %s", ws.id, term)
			ws.SetTerm(term)
			hasTerm = true
		case messageTypeLogin:
			if hasLogin {
				continue
			}
			config.User, _ = url.QueryUnescape(string(msg.Data))
			ws.logger.Printf("(%s) (%s) 登陆用户名使用： %s", ws.id, ws.currentUser, config.User)
			hasLogin = true
		case messageTypePassword:
			if hasAuth {
				continue
			}
			if ws.sshConn == nil {
				ws.logger.Printf("must connect addr first")
				continue
			}
			if config.User == "" {
				ws.logger.Printf("must set user first")
				continue
			}
			password, _ := url.QueryUnescape(string(msg.Data))
			//ws.logger.Printf("(%s) auth with password %s", ws.id, password)
			ws.logger.Printf("(%s) (%s) 使用了密码 *不显示*", ws.id, ws.currentUser)
			config.Auth = append(config.Auth, ssh.Password(password))
			session, err = ws.newSSHXtermSession(ws.sshConn, &config, msg)
			if err != nil {
				_ = ws.websocket.WriteJSON(&message{Type: messageTypeStderr, Data: []byte("用户名密码错误！\r\n")})
				return errors.Wrap(err, "用户名密码错误！")
			}
			defer func() {
				_ = session.Close()
			}()

			stdin, err = session.StdinPipe()
			if err != nil {
				_ = ws.websocket.WriteJSON(&message{Type: messageTypeStderr, Data: []byte("get stdin channel error\r\n")})
				return errors.Wrap(err, "get stdin channel error")
			}
			defer func() {
				_ = stdin.Close()
			}()

			err = ws.transformOutput(session, ws.websocket)
			if err != nil {
				_ = ws.websocket.WriteJSON(&message{Type: messageTypeStderr, Data: []byte("get stdin & stderr channel error\r\n")})
				return errors.Wrap(err, "get stdin & stderr channel error")
			}

			err = session.Shell()
			if err != nil {
				_ = ws.websocket.WriteJSON(&message{Type: messageTypeStderr, Data: []byte("开启终端失败！\r\n")})
				return errors.Wrap(err, "开启终端失败！")
			}

			hasAuth = true
		case messageTypeStdin:
			if stdin == nil {
				ws.logger.Printf("stdin wait login")
				continue
			}
			_, err = stdin.Write(msg.Data)
			if err != nil {
				_ = ws.websocket.WriteJSON(&message{Type: messageTypeStderr, Data: []byte("write to stdin error\r\n")})
				return errors.Wrap(err, "write to stdin error")
			}

		case messageTypeResize:
			if session == nil {
				ws.logger.Printf("resize wait session")
				continue
			}
			err = session.WindowChange(msg.Rows, msg.Cols)
			if err != nil {
				_ = ws.websocket.WriteJSON(&message{Type: messageTypeStderr, Data: []byte("resize error\r\n")})
				return errors.Wrap(err, "resize error")
			}
		}
	}
}

// 创建 ssh 会话
func (ws *WebSSH) newSSHXtermSession(conn net.Conn, config *ssh.ClientConfig, msg message) (*ssh.Session, error) {
	c, chans, reqs, err := ssh.NewClientConn(conn, conn.RemoteAddr().String(), config)
	if err != nil {
		return nil, errors.Wrap(err, "打开终端失败！")
	}
	session, err := ssh.NewClient(c, chans, reqs).NewSession()
	if err != nil {
		return nil, errors.Wrap(err, "打开终端失败！")
	}
	modes := ssh.TerminalModes{
		ssh.ECHO: 1,
		ssh.TTY_OP_ISPEED: 8192,
		ssh.TTY_OP_OSPEED: 8192,
		ssh.IEXTEN: 0,
	}
	if msg.Cols <= 0 || msg.Cols > 500 {
		msg.Cols = 40
	}
	if msg.Rows <= 0 || msg.Rows > 1000 {
		msg.Rows = 80
	}
	err = session.RequestPty(ws.term, msg.Rows, msg.Cols, modes)
	if err != nil {
		return nil, errors.Wrap(err, "open pty error")
	}

	return session, nil
}

// 发送 ssh 会话的 stdout 和 stdin 数据到 websocket 连接
func (ws *WebSSH) transformOutput(session *ssh.Session, conn *websocket.Conn) error {
	stdout, err := session.StdoutPipe()
	if err != nil {
		return errors.Wrap(err, "get stdout channel error")
	}
	stderr, err := session.StderrPipe()
	if err != nil {
		return errors.Wrap(err, "get stderr channel error")
	}
	stdin, err := session.StdinPipe()
	if err != nil {
		return errors.Wrap(err, "get stdin channel error")
	}

	copyToMessage := func(t messageType, r io.Reader, w io.WriteCloser) {
		buff := make([]byte, ws.buffSize)
		for {
			n, err := r.Read(buff)
			if err != nil {
				return
			}
			err = conn.WriteJSON(&message{Type: t, Data: buff[:n]})
			if err != nil {
				ws.logger.Printf("%s write fail", t)
				return
			}
		}
	}
	go copyToMessage(messageTypeStdout, stdout, stdin)
	go copyToMessage(messageTypeStderr, stderr, stdin)
	return nil
}

package v1

import (
	"gin-element-admin/model/response"
	"gin-element-admin/service"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"os"
	"time"
)

func CreateWebSSH(c *gin.Context) {
	wsId := c.Request.Header.Get("Sec-Websocket-Key")
	webSSH := service.NewWebSSH()
	webSSH.SetTerm("xterm")
	webSSH.SetBuffSize(8192)
	webSSH.SetId(wsId)
	webSSH.SetConnTimeOut(5 * time.Second)
	webSSH.SetLogger(log.New(os.Stderr, "[gea-webSSH] ", log.Ltime|log.Ldate))
	upGrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
		Subprotocols: []string{c.Request.Header.Get("Sec-WebSocket-Protocol")},
		ReadBufferSize: 8192,
		WriteBufferSize: 8192,
	}
	wsConn, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err!= nil{
		response.FailWithMessage("创建websocket失败！", c)
	}
	webSSH.AddWebsocket(wsConn)
}

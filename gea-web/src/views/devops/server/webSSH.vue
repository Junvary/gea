<template>
    <div class="webSSH" :key="$route.id">
        <div id="terminal" ref="gea-console"></div>
        <div class="header">
            <span style="margin-bottom: 5px">主机名称：{{hostInfo.name}}</span>
            <span style="margin-bottom: 5px">IP：{{hostInfo.ip}}</span>
            <span style="margin-bottom: 5px">端口：{{hostInfo.port}}</span>
            <el-button type="danger" size="small" @click="closeWS">关闭连接</el-button>
        </div>
        <el-dialog title="登录" :visible.sync="dialogVisible" width="30%" :show-close="false" destroy-on-close
            :close-on-press-escape="false" :close-on-click-modal="false">
            <el-form :model="sshForm" :rules="sshFormRules" ref="sshForm" label-width="100px">
                <el-form-item label="主机名称：">
                    <span>{{hostInfo.name}}</span>
                </el-form-item>
                <el-form-item label="IP：">
                    <span>{{hostInfo.ip}}</span>
                </el-form-item>
                <el-form-item label="端口：">
                    <span>{{hostInfo.port}}</span>
                </el-form-item>
                <el-form-item label="用户名：" prop="username">
                    <el-input v-model="sshForm.username"></el-input>
                </el-form-item>
                <el-form-item label="密码：" prop="password">
                    <el-input v-model="sshForm.password" type="password"></el-input>
                </el-form-item>
            </el-form>
            <span slot="footer" class="dialog-footer">
                <el-button type="primary" @click="cancelWebSSH">取 消</el-button>
                <el-button type="primary" @click="submitWebSSH('sshForm')">登 录</el-button>
            </span>
        </el-dialog>
    </div>
</template>

<script>
import { Terminal } from 'xterm'
import { FitAddon } from 'xterm-addon-fit'
import 'xterm/css/xterm.css'
import 'xterm/lib/xterm.js'
export default {
    name: 'WebSSH',
    data() {
        return {
            websocket: null,
            term: null,
            lockReconnect: false,
            cols: Math.floor((window.innerWidth - 300) / 9),
            rows: Math.floor((window.innerHeight - 140) / 17),
            hostInfo: this.$route.params,
            userinfo: this.$store.getters['user/userinfo'],
            dialogVisible: true,
            fileDrawer: {
                visible: false,
                hostInfo: {},
            },
            fileButtonVisible: false,
            sshForm: {
                username: '',
                password: '',
            },
            sshFormRules: {
                username: [{ required: true, message: '请输入用户名', trigger: 'blur' }],
                password: [{ required: true, message: '请输入密码', trigger: 'blur' }],
            },
        }
    },
    beforeDestroy() {
        if (this.websocket) {
            this.closeWS()
        }
        if (this.term) {
            this.term.disposed()
        }
        this.dialogVisible = true
    },
    methods: {
        closeWS() {
            this.websocket.close()
        },
        cancelWebSSH() {
            this.dialogVisible = false
        },
        submitWebSSH(form) {
            this.$refs[form].validate((valid) => {
                if (!valid) {
                    return false
                } else {
                    this.initWS()
                    this.dialogVisible = false
                }
            })
        },
        initTerm() {
            this.term = new Terminal({
                rendererType: 'canvas',
                cursorBlink: true,
                cols: this.cols,
                rows: this.rows,
                theme: {
                    foreground: 'white',
                    background: 'black',
                },
            })
            this.term.open(this.$refs['gea-console'])
            var fitAddon = new FitAddon()
            this.term.loadAddon(fitAddon)
            fitAddon.fit()
            this.term.focus()
            this.term.writeln('你好，' + this.userinfo.nickName + '！')
        },
        initWS() {
            this.websocket = new WebSocket(window._CONFIG['wsURL'] + '/webSSH', /* Sec-WebSocket-Protocol */ ['gea-webSSH'])
            this.websocket.binaryType = 'arraybuffer'
            this.initTerm()
            this.websocket.onopen = this.wsOpen
            this.websocket.onerror = this.wsError
            this.websocket.onmessage = this.wsMessage
            this.websocket.onclose = this.wsClose
        },
        wsOpen: function () {
            var that = this
            that.fileButtonVisible = true
            that.websocket.send(JSON.stringify({ type: 'currentUser', data: this.utf8_to_b64(that.userinfo.nickName + '_' + that.userinfo.userName) }))
            that.websocket.send(JSON.stringify({ type: 'addr', data: that.utf8_to_b64(that.hostInfo.ip + ':' + that.hostInfo.port) }))
            that.websocket.send(JSON.stringify({ type: 'login', data: that.utf8_to_b64(that.sshForm.username) }))
            that.websocket.send(JSON.stringify({ type: 'password', data: that.utf8_to_b64(that.sshForm.password) }))
            that.websocket.send(JSON.stringify({ type: 'resize', cols: that.cols, rows: that.rows }))
            that.term.resize(this.cols, this.rows)
            that.term.onData(function (data) {
                that.websocket.send(JSON.stringify({ type: 'stdin', data: that.utf8_to_b64(data) }))
            })
        },
        wsError: function (e) {
            // eslint-disable-next-line
            console.log(e)
            this.term.writeln('貌似连接错误了？')
            this.fileDrawer.visible = false
            this.fileButtonVisible = false
            // this.wsReconnect()
        },
        wsReconnect() {
            var that = this
            if (that.lockReconnect) return
            that.lockReconnect = true
            setTimeout(function () {
                that.$message.error('ws连接出现问题，尝试重连中...')
                that.initWS()
                that.lockReconnect = false
            }, 5000)
        },
        wsMessage: function (recv) {
            try {
                let msg = JSON.parse(recv.data)
                switch (msg.type) {
                    case 'stdout':
                    case 'stderr':
                        this.term.write(this.b64_to_utf8(msg.data))
                        break
                    case 'console':
                        this.term.write(this.b64_to_utf8(msg.data))
                        break
                    case 'alert':
                        this.term.write(this.b64_to_utf8(msg.data))
                        break
                    default:
                        console.log('unsupport type msg', msg)
                }
            } catch (e) {
                console.log('unsupport data', recv.data)
            }
        },
        wsClose: function (e) {
            this.term.writeln('连接已断开，感谢使用...')
            this.fileDrawer.visible = false
            this.fileButtonVisible = false
            this.websocket.close()
        },
        utf8_to_b64(rawString) {
            return btoa(unescape(encodeURIComponent(rawString)))
        },
        b64_to_utf8(encodeString) {
            return decodeURIComponent(escape(atob(encodeString)))
        },
    },
}
</script>

<style scoped>
.webSSH {
    position: relative;
    height: 100%;
    width: 100%;
}

.header {
    color: #1976d2;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    position: absolute;
    right: 0;
    top: 0;
    margin: 10px;
    z-index: 999;
}
</style>
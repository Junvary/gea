<template>
    <div class="system">
        <el-button @click="update" type="danger">立即更新</el-button>
        <span style="margin-left: 20px; font-weight: bold; color: #C10015">在更改配置文件前，希望你清楚这将带来什么样的效果！</span>
        <el-form :model="config" ref="form">
            <el-row :gutter="24">
                <el-col :span="8">
                    <el-card>
                        <div slot="header" style="color: #C10015; font-weight: bold">
                            系统配置
                        </div>
                        <el-form-item label="环境值">
                            <el-input v-model="config.system.env"></el-input>
                        </el-form-item>
                        <el-form-item label="端口值">
                            <el-input v-model.number="config.system.addr"></el-input>
                        </el-form-item>
                        <el-form-item label="配置文件环境变量名">
                            <el-input v-model.number="config.system.configEnv"></el-input>
                        </el-form-item>
                        <el-form-item label="多点登录拦截">
                            <el-checkbox v-model="config.system.useMultipoint">开启</el-checkbox>
                        </el-form-item>
                    </el-card>
                    <el-card>
                        <div slot="header" style="color: #C10015; font-weight: bold">
                            Casbin
                        </div>
                        <el-form-item label="模型地址">
                            <el-input v-model="config.casbin.modelPath"></el-input>
                        </el-form-item>
                    </el-card>
                </el-col>
                <el-col :span="8">
                    <el-card>
                        <div slot="header" style="color: #C10015; font-weight: bold">
                            JWT
                        </div>
                        <el-form-item label="签名">
                            <el-input v-model="config.jwt.signingKey"></el-input>
                        </el-form-item>
                        <el-form-item label="过期事件">
                            <el-input v-model="config.jwt.expiresTime"></el-input>
                        </el-form-item>
                        <el-form-item label="缓冲期">
                            <el-input v-model="config.jwt.bufferTime"></el-input>
                        </el-form-item>
                    </el-card>
                    <el-card>
                        <div slot="header" style="color: #C10015; font-weight: bold">
                            验证码
                        </div>
                        <el-form-item label="字符数">
                            <el-input v-model.number="config.captcha.keyLong"></el-input>
                        </el-form-item>
                        <el-form-item label="图像宽度">
                            <el-input v-model.number="config.captcha.imgWidth"></el-input>
                        </el-form-item>
                        <el-form-item label="图像高度">
                            <el-input v-model.number="config.captcha.imgHeight"></el-input>
                        </el-form-item>
                    </el-card>

                </el-col>
                <el-col :span="8">
                    <el-card>
                        <div slot="header" style="color: #C10015; font-weight: bold">
                            上传
                        </div>
                        <el-form-item label="本地文件路径">
                            <el-input v-model="config.local.path"></el-input>
                        </el-form-item>
                    </el-card>
                    <el-card>
                        <div slot="header" style="color: #C10015; font-weight: bold">
                            Redis
                        </div>
                        <el-form-item label="数据库">
                            <el-input v-model="config.redis.db"></el-input>
                        </el-form-item>
                        <el-form-item label="IP端口">
                            <el-input v-model="config.redis.addr"></el-input>
                        </el-form-item>
                        <el-form-item label="密码">
                            <el-input v-model="config.redis.password"></el-input>
                        </el-form-item>
                    </el-card>
                </el-col>
            </el-row>

            <el-row :gutter="24">
                <el-col :span="8">
                    <el-card>
                        <div slot="header" style="color: #C10015; font-weight: bold">
                            Mysql
                        </div>
                        <el-form-item label="用户名">
                            <el-input v-model="config.mysql.username"></el-input>
                        </el-form-item>
                        <el-form-item label="密码">
                            <el-input v-model="config.mysql.password"></el-input>
                        </el-form-item>
                        <el-form-item label="IP端口">
                            <el-input v-model="config.mysql.path"></el-input>
                        </el-form-item>
                        <el-form-item label="数据库名">
                            <el-input v-model="config.mysql.dbname"></el-input>
                        </el-form-item>
                        <el-form-item label="闲置连接数">
                            <el-input v-model.number="config.mysql.maxIdleConns"></el-input>
                        </el-form-item>
                        <el-form-item label="最大连接数">
                            <el-input v-model.number="config.mysql.maxOpenConns"></el-input>
                        </el-form-item>
                        <el-form-item label="归档模式">
                            <el-checkbox v-model="config.mysql.logMode"></el-checkbox>
                        </el-form-item>
                    </el-card>
                </el-col>
                <el-col :span="8">
                    <el-card>
                        <div slot="header" style="color: #C10015; font-weight: bold">
                            Zap日志配置
                        </div>
                        <el-form-item label="级别">
                            <el-input v-model.number="config.zap.level"></el-input>
                        </el-form-item>
                        <el-form-item label="输出">
                            <el-input v-model="config.zap.format"></el-input>
                        </el-form-item>
                        <el-form-item label="日志前缀">
                            <el-input v-model="config.zap.prefix"></el-input>
                        </el-form-item>
                        <el-form-item label="日志文件夹">
                            <el-input v-model="config.zap.director"></el-input>
                        </el-form-item>
                        <el-form-item label="软链接名称">
                            <el-input v-model="config.zap.linkName"></el-input>
                        </el-form-item>
                        <el-form-item label="编码级">
                            <el-input v-model="config.zap.encodeLevel"></el-input>
                        </el-form-item>
                        <el-form-item label="栈名">
                            <el-input v-model="config.zap.stacktraceKey"></el-input>
                        </el-form-item>
                        <el-form-item label="显示行">
                            <el-checkbox v-model="config.zap.showLine"></el-checkbox>
                        </el-form-item>
                        <el-form-item label="输出控制台">
                            <el-checkbox v-model="config.zap.logInConsole"></el-checkbox>
                        </el-form-item>
                    </el-card>
                </el-col>
            </el-row>
        </el-form>
    </div>
</template>

<script>
import { getSystemConfig, setSystemConfig } from '@/api/system'
export default {
    name: 'Config',
    data() {
        return {
            config: {
                system: {},
                jwt: {},
                casbin: {},
                mysql: {},
                redis: {},
                captcha: {},
                zap: {},
                local: {},
            },
        }
    },
    async created() {
        await this.initForm()
    },
    methods: {
        async initForm() {
            const res = await getSystemConfig()
            if (res.code == 0) {
                this.config = res.data.config
            }
        },
        async update() {
            this.$confirm('不要轻易更改配置文件内容，希望你很清楚自己在做什么，是否继续?', '警告', {
                confirmButtonText: '确定更新',
                cancelButtonText: '我再想想',
                type: 'error',
                lockScroll: false,
                showClose: false,
                closeOnClickModal: false,
                closeOnPressEscape: false,
            })
                .then(async () => {
                    const res = await setSystemConfig({ config: this.config })
                    if (res.code === 0) {
                        this.$message({
                            type: 'success',
                            message: '配置文件设置成功',
                        })
                        await this.initForm()
                    }
                })
                .catch(() => {
                    this.$message({
                        type: 'info',
                        message: '你的选择是正确的！',
                    })
                })
        },
    },
}
</script>
<style lang="scss" scoped>
.el-card {
    margin-bottom: 30px;
}
</style>

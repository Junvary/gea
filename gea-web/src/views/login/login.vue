<template>
    <el-form style="width: 90%" :model="loginForm" :rules="rules" ref="loginForm" @keyup.enter.native="submitForm"
        v-loading="loading" element-loading-text="正在登录...">
        <el-form-item prop="username">
            <el-input placeholder="请输入用户名" v-model="loginForm.username">
                <i class="el-input__icon el-icon-user" slot="suffix"></i>
            </el-input>
        </el-form-item>
        <el-form-item prop="password">
            <el-input :type="lock === 'lock' ? 'password' : 'text'" placeholder="请输入密码" v-model="loginForm.password">
                <i :class="'el-input__icon el-icon-' + lock" @click="changeLock" slot="suffix"></i>
            </el-input>
        </el-form-item>
        <el-form-item prop="captcha">
            <el-input v-model="loginForm.captcha" name="logVerify" placeholder="请输入验证码" style="width:100%">
                <el-image slot="append" v-if="picPath" :src="picPath" @click="loginVefify()"
                    style="width: 120px; height: 35px" />
            </el-input>
        </el-form-item>
        <el-form-item>
            <el-button type="primary" @click="submitForm" style="width:100%" :loading="loading">登 录</el-button>
        </el-form-item>
    </el-form>
</template>

<script>
import { mapActions } from 'vuex'
import { captcha } from '@/api/user'

export default {
    name: 'login',
    data() {
        const checkUsername = (rule, value, callback) => {
            if (value.length < 5 || value.length > 12) {
                return callback(new Error('请输入正确的用户名'))
            } else {
                callback()
            }
        }
        const checkPassword = (rule, value, callback) => {
            if (value.length < 6 || value.length > 12) {
                return callback(new Error('请输入正确的密码'))
            } else {
                callback()
            }
        }
        return {
            loading: false,
            loginForm: {
                username: 'admin',
                password: '123456',
                captcha: '',
                captchaId: '',
            },
            rules: {
                username: [{ validator: checkUsername, trigger: 'blur' }],
                password: [{ validator: checkPassword, trigger: 'blur' }],
                captcha: [{ required: true, message: '请输入验证码', trigger: 'blur' }],
            },
            lock: 'lock',
            picPath: '',
        }
    },
    created() {
        this.loginVefify()
    },
    methods: {
        ...mapActions('user', ['Login']),
        async login() {
            return await this.Login(this.loginForm)
        },
        async submitForm() {
            this.$refs.loginForm.validate(async (v) => {
                if (v) {
                    this.loading = true
                    const flag = await this.login()
                    if (!flag) {
                        this.loginVefify()
                        this.loading = false
                    }
                } else {
                    this.$message({
                        type: 'error',
                        message: '请正确填写登录信息',
                        showClose: true,
                    })
                    this.loginVefify()
                    return false
                }
            })
        },
        changeLock() {
            this.lock === 'lock' ? (this.lock = 'unlock') : (this.lock = 'lock')
        },
        loginVefify() {
            captcha({}).then((ele) => {
                this.picPath = ele.data.picPath
                this.loginForm.captchaId = ele.data.captchaId
            })
        },
    },
}
</script>

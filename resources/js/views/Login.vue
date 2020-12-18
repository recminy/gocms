<template>
    <div class="login-container">
        <div class="login-title">
            <h2 class="text-center pb-2 pt-3 border-bottom">登录后台管理系统</h2>
        </div>
        <div class="login-form">
            <el-form ref="form" @keyup.enter.native="login()" :rules="rules" :model="form">
                <div class="login-input login-input-username">
                    <el-form-item label="" prop="username">
                        <div class="icon-container1"><i class="fa fa-user" @click="username = ''"></i></div>
                        <el-input placeholder="账号" id="username" v-model.trim="form.username" :autofocus="autoFocus" :maxlength="11"></el-input>
                    </el-form-item>
                </div>
                <div class="login-input login-input-password">
                    <el-form-item label="" prop="password">
                        <el-input type="password" id="password" placeholder="密码" v-model.trim="form.password" :autofocus="!autoFocus"></el-input>
                        <div class="icon-container2"><i class="fa fa-key" @click="password = ''"></i></div>
                    </el-form-item>
                </div>
                <div class="login-input">
                    <el-button type="primary" @click="login()" class="btn-block">登 录</el-button>
                </div>
            </el-form>
            <el-row>
                <el-col :span="6" :offset="18">
                    <div  class="forget-password pull-right" @click="forgetPassword">忘记密码?</div>
                </el-col>
            </el-row>
        </div>
    </div>
</template>
<style>
    .login-container {
        width: 450px;
        height: 350px;
        margin: 120px auto;
        background:#FFF;
        border-radius: 10px;
        border: 22px solid #eee
    }
    .login-title {
        text-align:center;
        border-bottom: 1px solid #eee;

    }
    .login-form {
        width: 88%;
        margin: 6%
    }
    .login-input {
        margin-bottom: 1.5rem;
    }
    .forget-password {
        float: right !important;
        text-decoration: none !important;
        color: #999;
        cursor: pointer;
    }
    .forget-password:hover {
        color: #666 !important;
    }
    .btn-block {
        width: 100%;
    }
</style>
<script>
    const axios = require("axios");
    export default {
        name: "Login",
        mounted() {
            if (!localStorage.username && localStorage.username != undefined) {
                this.username = localStorage.username;
            }
        },
        computed: {
            autoFocus() {
                return !this.username;
            }
        },
        data() {
            return {
                form: {
                    username: "",
                    password: "",
                },
                rules: {
                    username: [
                        {required: true, trigger: "blur", message: "用户名必填"},
                    ],
                    password: [
                        {required: true, trigger: "blur", message: "密码必填"}
                    ]
                }
            }
        },
        methods: {
            login() {
                this.$refs.form.validate(valid => {
                    if (valid) {
                        axios.patch("/admin/login", this.form).then(res => {

                        }).catch(err => {
                            //TODO
                        })
                    }
                })
            },
            forgetPassword() {
                alert("ForgetPassword")
            }
        }
    }
</script>

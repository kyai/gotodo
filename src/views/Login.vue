<template>
<div>
    <div class="panel">
        <div class="mb-5">GOTODO</div>
        <div class="signin">
            <v-text-field label="用户名" v-model="signinForm.username" required variant="outlined" />
            <v-text-field label="密码" type="password" v-model="signinForm.password" required variant="outlined" />
            <v-btn :loading="signinLoading" block color="success" size="large" @click="signin">登 录</v-btn>
        </div>
        <div class="signup"></div>
    </div>
</div>
</template>
<style lang="scss" scoped>
.panel{
    width: 360px;
    height: 100px;
    margin: 60px auto 0 auto;
}
</style>
<script>
import { useUserStore } from '@/stores/user'
export default {
    data(){
        return {
            signinForm: {},
            signinLoading: false,
        }
    },
    methods: {
        signin(){
            this.signinLoading = true
            useUserStore().login(this.signinForm).then(() => {
                this.$toast('登录成功')
                this.$router.push('/')
            }).catch(() => {
                this.signinLoading = false
            })
        }
    }
}
</script>
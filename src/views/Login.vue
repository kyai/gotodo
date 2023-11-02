<template>
<v-layout class="d-block bg-grey-lighten-3">
    <v-card class="mt-8 mx-auto pa-8" width="480" :loading="loading">
        <v-card-title class="text-center mb-8">Gotodo</v-card-title>
        <v-carousel
            :continuous="false"
            :show-arrows="false"
            hide-delimiters
            height="360"
            v-model="carousel"
        >
            <!-- 登录 -->
            <v-carousel-item>
                <sign-in :loading="loading" @submit="signin" @switch="carousel=1" />
            </v-carousel-item>
            <!-- 注册 -->
            <v-carousel-item>
                <sign-up :loading="loading" @submit="signup" @switch="carousel=0" />
            </v-carousel-item>
        </v-carousel>
    </v-card>
</v-layout>
</template>
<style lang="scss" scoped>
.carousel{
    width: 320px;
    margin: 64px auto 0 auto;
}
</style>
<script>
import SignIn from '@/components/SignIn.vue'
import SignUp from '@/components/SignUp.vue'
import { useUserStore } from '@/stores/user'
export default {
    data(){
        return {
            carousel: 0,
            loading: false,
        }
    },
    components: {
        SignIn,
        SignUp,
    },
    methods: {
        signin(form){
            this.loading = true
            useUserStore().login(form).then(() => {
                this.$toast('登录成功')
                this.$router.push('/')
            }).catch(() => {
                this.loading = false
            })
        },
        signup(form){
            this.loading = true
            this.$post('SignUp', form).then(() => {
                this.$toast('注册成功, 请登录')
                this.carousel = 0
            }).catch(() => {
                this.loading = false
            })
        },
    }
}
</script>
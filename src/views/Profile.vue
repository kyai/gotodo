<template>
<v-card width="360" class="mx-auto pa-5">
    <v-form validate-on="submit lazy" :disabled="loading" @submit.prevent="submit">
        <v-text-field label="账号" v-model="user.username" density="compact" disabled />
        <v-text-field label="昵称" v-model="user.nickname" density="compact" hint="其他用户可见" persistent-hint />
        <v-text-field label="密码" v-model="user.password" density="compact" hint="留空则不修改" persistent-hint class="mt-2" />
        <v-btn :loading="loading" type="submit" block elevation="0" color="primary" class="mt-4">保存修改</v-btn>
    </v-form>
</v-card>
</template>
<script>
import { useUserStore } from '@/stores/user'

export default {
    data(){
        return {
            user: {},
            loading: false,
        }
    },
    mounted(){
        this.getProfile()
    },
    methods: {
        getProfile(){
            this.loading = true
            useUserStore().profile().then(res => {
                this.user = res
                this.loading = false
            })
        },
        submit(){
            this.$post('PutProfile', this.user).then(() => {
                this.$toast('修改成功')
                useUserStore().profile()
            })
        },
    }
}
</script>
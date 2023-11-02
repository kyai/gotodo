<template>
<div class="app-navbar px-3 bg-grey-darken-1">
    <v-tabs class="float-left" density="compact" hide-slider>
        <v-tab v-if="showHome" to="/" class="tab-home">
            <v-icon>mdi-home-outline</v-icon>
        </v-tab>
        <slot name="left" />
    </v-tabs>
    <div class="app-navbar-title float-left">{{title}}</div>
    <v-tabs class="float-right" density="compact" hide-slider>
        <slot name="right" />
        <v-tab id="menu-user"><v-icon>mdi-account-outline</v-icon>{{uname}}</v-tab>
    </v-tabs>

    <v-menu activator="#menu-user">
        <v-list density="compact">
            <v-list-item to="/profile">
                <v-list-item-title>账号设置</v-list-item-title>
            </v-list-item>
            <v-list-item to="/system">
                <v-list-item-title>系统设置</v-list-item-title>
            </v-list-item>
            <v-list-item @click="logout">
                <v-list-item-title>退出登录</v-list-item-title>
            </v-list-item>
        </v-list>
    </v-menu>
</div>
</template>
<style lang="scss" scoped>
.app-navbar{
    width: 100%;
    height: $topHeight;
    position: fixed;
    z-index: 9999;
    &-title{
        line-height: $topHeight;
        cursor: default;
    }
}
.tab-home{
    min-width: 0;
}
</style>
<script>
import { useUserStore } from '@/stores/user'

export default {
    data(){
        return {}
    },
    props: {
        title: String,
        showHome: Boolean,
    },
    computed: {
        uname(){
            return useUserStore().name
        },
    },
    methods: {
        logout(){
            useUserStore().logout().then(() => {
                this.$router.push('/login')
            })
        },
    }
}
</script>
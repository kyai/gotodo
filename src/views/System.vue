<template>
<div class="app">
    <nav-bar show-home title="系统设置" />
    <div class="app-main">
        <v-tabs align-tabs="center" bg-color="teal-darken-3">
            <v-tab to="/system/account">账号管理</v-tab>
            <v-tab to="/system/project">项目管理</v-tab>
            <v-tab to="/system/setting">系统设置</v-tab>
        </v-tabs>
        <router-view />
    </div>
</div>
</template>
<script>
import NavBar from '@/components/NavBar.vue'
export default {
    data(){
        return {}
    },
    components: {
        NavBar,
    },
}
</script>
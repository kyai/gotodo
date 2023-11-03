<template>
<div class="app">
    <div class="app-head">
        <nav-bar show-home title="系统设置" />
        <v-tabs align-tabs="center" bg-color="grey-lighten-3">
            <v-tab to="/system/setting">系统管理</v-tab>
            <v-tab to="/system/project">项目管理</v-tab>
            <v-tab to="/system/account">账号管理</v-tab>
        </v-tabs>
    </div>
    <div class="app-main">
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
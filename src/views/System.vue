<template>
<div class="app">
    <div class="app-head">
        <nav-bar show-home title="设置" />
        <v-tabs align-tabs="center" bg-color="grey-lighten-3">
            <v-tab to="/system/profile">个人资料</v-tab>
            <v-tab to="/system/project">项目管理</v-tab>
            <v-tab to="/system/account">账号管理</v-tab>
            <v-tab to="/system/setting">系统设置</v-tab>
        </v-tabs>
    </div>
    <div class="app-main app-system">
        <router-view class="app-system-inner" />
    </div>
</div>
</template>
<style lang="scss">
.app-system{
    &-inner{
        width: 100%;
        max-width: 800px;
        min-height: 100%;
        margin: 0 auto;
        padding: 10px;
        box-shadow: 0 0 4px 1px rgba(0, 0, 0, 0.2);
    }
}
</style>
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
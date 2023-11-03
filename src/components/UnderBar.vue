<template>
<div class="app-underbar pb-1">
    <a :href="github" target="_blank" class="text-grey"><v-icon>mdi-github</v-icon></a>
</div>
</template>
<style lang="scss">
.app-underbar{
    width: 100%;
    position: absolute;
    bottom: 0;
    text-align: center;
}
</style>
<script>
export default {
    data(){
        return {
            github: import.meta.env.VITE_GITHUB,
        }
    }
}
</script>
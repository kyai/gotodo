<template>
<div class="topic elevation-2">
    <div class="title">{{topic.title}}</div>
    <div class="tasks">
        <li v-for="v in 100" :key="v">{{v}}</li>
    </div>
</div>
</template>
<style lang="scss" scoped>
.topic{
    width: 300px;
    height: 100%;
    display: inline-block;
    margin: 5px 10px 0 10px;
    border-radius: 20px 20px 0 0;
    .title{
        height: 30px;
    }
    .tasks{
        height: calc(100% - 35px);
        overflow-y: auto;
    }
}
</style>
<script>
export default {
    props: {
        topic: Object,
    },
}
</script>
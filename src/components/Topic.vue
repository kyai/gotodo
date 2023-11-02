<template>
<div class="topic elevation-2">
    <div class="title">{{topic.title}}</div>
    <div class="tasks">
        <li v-for="v,k in topic.tasks" :key="k">{{v.caption}}</li>
    </div>
</div>
</template>
<style lang="scss" scoped>
$marginTop: 10px;
$titleHeight: 48px;
.topic{
    width: 300px;
    height: 100%;
    display: inline-block;
    margin: $marginTop 10px 0 10px;
    border-radius: 20px 20px 0 0;
    overflow: hidden;
    .title{
        height: $titleHeight;
        line-height: $titleHeight;
        background: green;
    }
    .tasks{
        height: calc(100% - $marginTop - $titleHeight);
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
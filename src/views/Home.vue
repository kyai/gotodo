<template>
<div class="home">
    <div class="head px-3 bg-grey-darken-1">
        <div class="float-left">
            <pro-select ref="proSelect" @change="loadProject" />
        </div>
        <v-tabs class="float-right" density="compact" hide-slider>
            <v-tab @click="addProject">添加项目</v-tab>
            <v-tab @click="addTopic">添加主题</v-tab>
            <v-tab @click="addTask">添加任务</v-tab>
            <v-tab>username</v-tab>
        </v-tabs>
    </div>
    <div class="main">
        <topic v-for="topic,k in project.topics" :key="k" :topic="topic" />
    </div>
    <task-editor ref="taskEditor" @submit="loadProject" />
</div>
</template>
<style lang="scss">
.home{
    width: 100%;
    height: calc(100% - $topHeight);
    position: fixed;
}
.head{
    height: $topHeight;
    width: 100%;
}
.main{
    width: 100% !important;
    height: 100% !important;
    text-align: center !important;
    overflow-x: auto !important;
    overflow-y: hidden !important;
    white-space: nowrap !important;
}
</style>
<script>
import ProSelect from '@/components/ProSelect.vue'
import Topic from '@/components/Topic.vue'
import TaskEditor from '@/components/TaskEditor.vue'
export default {
    data(){
        return {
            project: {}, //当前项目
        }
    },
    components: {
        ProSelect,
        Topic,
        TaskEditor,
    },
    mounted(){
    },
    methods: {
        addProject(){
            var title = prompt('标题')
            this.$post('AddProject', {title}).then(id => {
                this.$toast('成功')
                this.$refs.proSelect.getList(id)
            })
        },
        loadProject(id){
            id = id || this.project.id || 0
            this.$post('GetProject', {id}).then(res => {
                this.project = res
            })
        },
        addTopic(){
            var title = prompt('标题')
            this.$post('AddTopic', {title, project_id: this.project.id}).then(() => {
                this.$toast('成功')
                this.loadProject()
            })
        },
        addTask(){
            this.$refs.taskEditor.open(this.project)
        },
    }
}
</script>
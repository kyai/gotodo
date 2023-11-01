<template>
<div class="app home">
    <nav-bar>
        <template #left>
            <pro-select ref="proSelect" @change="loadProject" />
        </template>
        <template #right>
            <v-tab @click="addProject">添加项目</v-tab>
            <v-tab @click="addTopic">添加主题</v-tab>
            <v-tab @click="addTask">添加任务</v-tab>
        </template>
    </nav-bar>
    <div class="app-main main">
        <topic v-for="topic,k in project.topics" :key="k" :topic="topic" />
    </div>
    <task-editor ref="taskEditor" @submit="loadProject" />
</div>
</template>
<style lang="scss" scoped>
.home{
    width: 100%;
    height: 100%;
    position: fixed;
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
import NavBar from '@/components/NavBar.vue'
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
        NavBar,
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
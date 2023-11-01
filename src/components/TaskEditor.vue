<template>
<v-dialog v-model="dialog" width="640">
    <v-card>
        <v-card-title>
            <span v-if="isNew">创建任务</span>
            <span v-else>编辑任务</span>
        </v-card-title>
        <v-card-text>
            <div>主题</div>
            <div>
                <v-radio-group inline v-model="task.topic_id">
                    <v-radio v-for="v,k in project.topics" :key="k" :label="v.title" :value="v.id" />
                </v-radio-group>
            </div>
            <div>标题</div>
            <div>
                <v-text-field v-model="task.caption" />
            </div>
            <div>内容</div>
            <div>
                <v-textarea v-model="task.content" />
            </div>
            <v-btn @click="submit">确认</v-btn>
        </v-card-text>
    </v-card>
</v-dialog>
</template>
<script>
export default {
    data(){
        return {
            dialog: false,
            isNew: false,
            project: {},
            task: {},
        }
    },
    methods: {
        open(project, task){
            this.project = project
            Object.assign(this.task, { project_id: project.id })
            this.isNew = !task
            this.dialog = true
        },
        close(){
            this.dialog = false
        },
        submit(){
            this.$post('AddTask', this.task).then(() => {
                this.$toast('成功')
                this.$emit('submit', this.project.id)
            })
        },
    }
}
</script>
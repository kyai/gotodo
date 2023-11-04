<template>
<div>
    <v-table>
        <thead>
            <tr>
                <th>名称</th>
                <th>最后更新时间</th>
                <th>操作</th>
            </tr>
        </thead>
        <tbody>
            <tr v-for="row,n in list" :key="n">
                <td>{{row.title}}</td>
                <td>{{$f.datetime(row.updated)}}</td>
                <td>
                    <v-btn variant="text" density="compact" color="success" prepend-icon="mdi-note-edit-outline" @click="edit(row)">编辑</v-btn>
                    <v-btn variant="text" density="compact" color="warning" prepend-icon="mdi-note-remove-outline">删除
                        <v-menu activator="parent" location="end">
                            <v-card class="pa-2">
                                <v-card-subtitle>确定删除项目</v-card-subtitle>
                                <v-card-text>{{row.title}}</v-card-text>
                                <v-card-actions><v-btn density="compact" color="error" @click="del(row)">删除</v-btn></v-card-actions>
                            </v-card>
                        </v-menu>
                    </v-btn>
                </td>
            </tr>
        </tbody>
    </v-table>
</div>
</template>
<script>
export default {
    data(){
        return {
            list: [],
        }
    },
    mounted(){
        this.getList()
    },
    methods: {
        getList(){
            this.$post('Projects').then(res => {
                this.list = res
            })
        },
        edit(row){
            row.title = prompt('项目标题', row.title)
            this.$post('PutProject', row).then(() => {
                this.$toast('成功')
                this.getList()
            })
        },
        del(row){
            this.$post('DelProject', row).then(() => {
                this.$toast('成功')
                this.getList()
            })
        },
    }
}
</script>
<template>
<div>
    <v-table>
        <thead>
            <tr>
                <th>名称</th>
                <th>最后更新时间</th>
            </tr>
        </thead>
        <tbody>
            <tr v-for="row,n in list" :key="n">
                <td>{{row.title}}</td>
                <td>{{$f.datetime(row.updated)}}</td>
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
    }
}
</script>
<template>
    <v-select
        class="pro-select"
        item-title="title"
        item-value="id"
        density="compact"
        bg-color="transparent"
        hide-details
        hide-no-data
        :loading="loading"
        :items="list"
        v-model="currentId"
        v-show="list.length"
    />
</template>
<style lang="scss">
.pro-select{
    height: $topHeight;
    .v-input__control{
        height: $topHeight;
    }
    .v-field__outline{
        --v-field-border-width: 0;
    }
    .v-field__field{
        height: $topHeight;
        .v-field__input{
            height: $topHeight;
            padding: 0 10px;
            .v-select__selection{
                height: $topHeight;
                margin: 0;
            }
        }
    }
}
</style>
<script>
export default {
    data(){
        return {
            list: [],
            loading: false,
            currentId: 0,
        }
    },
    watch: {
        currentId(v){
            this.$emit('change', v)
        }
    },
    mounted(){
        this.getList()
    },
    methods: {
        getList(id){
            this.loading = true
            this.$post('Projects').then(res => {
                this.list = res
                this.currentId = id || res[0].id
                this.loading = false
            })
        },
    }
}
</script>
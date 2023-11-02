<template>
<v-form :disabled="loading">
    <div class="text-subtitle-1 text-medium-emphasis">账号</div>
    <v-text-field
        density="compact"
        prepend-inner-icon="mdi-account-outline"
        variant="outlined"
        v-model="form.username"
    />

    <div class="text-subtitle-1 text-medium-emphasis d-flex align-center justify-space-between">密码
        <v-tooltip text="请联系管理员重置密码" location="top">
            <template v-slot:activator="{ props }">
                <a v-bind="props" class="text-caption text-decoration-none text-blue" href="#">忘记密码</a>
            </template>
        </v-tooltip>
    </div>
    <v-text-field
        :append-inner-icon="visible ? 'mdi-eye-off' : 'mdi-eye'"
        :type="visible ? 'text' : 'password'"
        density="compact"
        prepend-inner-icon="mdi-lock-outline"
        variant="outlined"
        @click:append-inner="visible = !visible"
        v-model="form.password"
    />

    <v-btn class="mt-8 mb-8" block color="blue" size="large" variant="tonal" :loading="loading" @click="submit">登录</v-btn>

    <div class="text-center">
        <a class="text-grey text-decoration-none" href="#" rel="noopener noreferrer" @click="switcc">
            <v-icon icon="mdi-empty" />注册<v-icon icon="mdi-chevron-right" />
        </a>
    </div>
</v-form>
</template>
<script>
export default {
    data(){
        return {
            form: {},
            visible: false,
        }
    },
    props: {
        loading: Boolean,
    },
    methods: {
        submit(){
            this.$emit('submit', this.form)
        },
        switcc(){
            this.$emit('switch')
        },
    }
}
</script>
const filters = {
    datetime(v){
        return new Date(v*1000).format('yyyy-MM-dd hh:mm:ss')
    },
}

export default function install(app){
    app.config.globalProperties.$f = filters
}
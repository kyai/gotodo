import { defineStore } from 'pinia'
import request from '@/plugins/request'
import cookies from 'vue-cookies'

const tokenName = 'token'

export const useUserStore = defineStore({
    id: 'user',
    state: () => ({
        name: '',
    }),
    actions: {
        logout(){
            return new Promise((resolve) => {
                this.$patch({
                    name: '',
                })
                cookies.remove(tokenName)
                resolve()
            })
        },
        login(data){
            return new Promise((resolve, reject) => {
                request({url: 'SignIn', data, method: 'POST'}).then(res => {
                    this.$patch({
                        name: res.user.nickname || res.user.username
                    })
                    cookies.set(tokenName, res.token, -1)
                    resolve(res)
                }).catch(err => {
                    reject(err)
                })
            })
        },
        profile(){
            return new Promise((resolve, reject) => {
                request({url: 'Profile', method: 'POST'}).then(user => {
                    this.$patch({
                        name: user.nickname || user.username
                    })
                    resolve(user)
                }).catch(err => {
                    reject(err)
                })
            })
        },
    },
})
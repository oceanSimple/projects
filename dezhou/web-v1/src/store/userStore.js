import {defineStore} from "pinia";
import requests from "../util/request.js";

export const useUserStore = defineStore("user", {
    state: () => ({
        staticInfo: {
            serverIp: "localhost:9000",
            wsUrl: "ws://localhost:9000/ws",
        },
        userInfo: {
            account: 'ocean'
        },
        roomInfo: {
            roomNumber: 0,
            roomPlayers: []
        },
        websocketConn: new WebSocket('ws://localhost:9000/ws'),
        handlerMap: new Map()
    }),
    actions: {
        async getBaidu() {
            requests.get("https://www.baidu.com").then(res => {
                console.log(res);
            }).catch(err => {
                console.log(err);
            })
        }
    }
})
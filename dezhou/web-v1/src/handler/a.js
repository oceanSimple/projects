import {useUserStore} from "../store/userStore.js";

export const newWebsocketConn = () => {
    let userStore = useUserStore()
    let wsConn = userStore.websocketConn
    wsConn.onopen = () => {
        console.log("wsConn open")
        userStore.websocketConn = wsConn
    }
    wsConn.onmessage = (msg) => {
        console.log("ws message", msg.data)
    }
    wsConn.onclose = () => {
        console.log("wsConn close")
    }
}

export const handleMessageEnter = (msg) => {
    console.log(msg)
    // 将msg转换成对象
    let msgObj = JSON.parse(msg)
    console.log("msg serverPath", msgObj.serverPath)
    console.log("msg data", msgObj.data)
}
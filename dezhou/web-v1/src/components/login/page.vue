<template>
  <h1>loginPage</h1>
  <span>账号</span>
  <el-input v-model="account" placeholder="请输入账号"></el-input>
  <span>密码</span>
  <el-input v-model="password" placeholder="请输入密码"></el-input>
  <br>
  <br>
  <br>
  <el-button type="primary" @click="login">goHome</el-button>
</template>

<script setup>
import {ref} from "vue";
import {useUserStore} from "../../store/userStore.js";
import requests from "../../util/request.js";
import {useRouter} from "vue-router";
import {newWebsocketConn} from "../../handler/a.js";

let userStore = useUserStore()
let account = ref('')
let password = ref('')

const router = useRouter()
const login = () => {
  requests.post("/login", {
    account: account.value,
    password: password.value
  }).then((res) => {
        // console.log(res)
        if (res.code === 200) {
          router.push("/home")

          // 建立websocket链接
          newWebsocketConn()
        }
      }
  )
}
</script>

<style scoped>

</style>
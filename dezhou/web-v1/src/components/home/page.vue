<template>
  <div v-for="(item, index) in roomList" :key="index">
    <div class="el-card card">
      <span>房间号: {{ item.Number }}</span>
      <span>房间人数: {{ item.PlayerCount }}/10</span>
      <span>房间状态: {{ item.Status }}</span>
      <el-button type="primary" @click="enterRoom(item.Number)">进入房间</el-button>
    </div>
  </div>
</template>

<script setup>
import {onMounted, ref} from "vue";
import requests from "../../util/request.js";
import {useUserStore} from "../../store/userStore.js";
import {useRouter} from "vue-router";

let roomList = ref([])
const router = useRouter();
let userStore = useUserStore();

onMounted(() => {
  requests.get("/getRoomList", {
    params: {page: 2}
  }).then((res) => {
    roomList.value = res.data
  })
})

const enterRoom = (roomNumber) => {
  // 获取account
  let account = userStore.userInfo.account

  // 发送请求
  requests.post("/enterRoom", {
    account: account,
    roomNumber: roomNumber
  }).then((res) => {
    if (res.code === 200) {
      userStore.roomInfo.roomNumber = roomNumber
      userStore.roomInfo.roomPlayers = res.data
      router.push("/room")
    } else {
      console.log("进入房间失败")
    }
  })
}
</script>

<style scoped>
.card > span {
  margin: 10px;
}
</style>
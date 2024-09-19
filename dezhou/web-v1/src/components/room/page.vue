<template>
  <div>
    <el-button type="primary" @click="returnHome">返回大厅</el-button>
  </div>
  <div>
    <el-card>
      <template #header>
        <span>房间信息</span>
      </template>
      <div>房间号: {{ roomNumber }}</div>
      <div>
        <span>玩家列表</span>
        <template v-for="(value, key) in roomPlayers">
          <el-tooltip
              :width="200"
              effect="light"
              placement="top-start"
              title="Title"
              trigger="click">
            <template #content>
              <div class="playerList-dialog-div">
                <div>转让房主权限</div>
                <div>加好友</div>
              </div>
            </template>

            <div>{{ key }}</div>

          </el-tooltip>
        </template>
      </div>
    </el-card>
  </div>
  <!--个人信息-->
  <div>
    <el-card>
      <template #header>
        <span>个人信息</span>
      </template>
      <div>
        <div>账号: {{ myInfo.account }}</div>
        <div>筹码: {{ myInfo.chip }}</div>
        <div>是否房主: {{ myInfo.isHost }}</div>
        <div>是否准备: {{ myInfo.isPrepared }}
        </div>
      </div>
    </el-card>
  </div>

  <!--准备按钮-->
  <div>
    <span v-if="!pageVar.isPrepared && !pageVar.isHost" class="button-prepare button-notPrepared"
          @click="changePrepareStatus">准备</span>
    <span v-else-if="pageVar.isPrepared && !pageVar.isHost" class="button-prepare button-yesPrepared"
          @click="changePrepareStatus">取消</span>
    <span v-else class="button-prepare button-host"
          @click="startGame">开始</span>
  </div>
</template>

<script setup>
import {onMounted, reactive, ref} from "vue";
import {useUserStore} from "../../store/userStore.js";
import {useRouter} from "vue-router";
import requests from "../../util/request.js";

let userStore = useUserStore();
let router = useRouter();

const roomNumber = userStore.roomInfo.roomNumber
const roomPlayers = ref(userStore.roomInfo.roomPlayers)
const myInfo = reactive({
  account: '',
  chip: 0,
  isHost: 0,
  isPrepared: 0,
})
const pageVar = reactive({
  isPrepared: false,
  isHost: false
})

onMounted(() => {
  // 从玩家列表中, 拿到自己的信息
  let myTempInfo = roomPlayers.value[userStore.userInfo.account]
  myInfo.account = myTempInfo.Account
  myInfo.chip = myTempInfo.Chip
  myInfo.isHost = myTempInfo.IsHost
  myInfo.isPrepared = myTempInfo.IsPrepared

  // 如果是房主, 则显示开始游戏按钮
  if (myInfo.isHost === 1) {
    pageVar.isHost = true
  }
})

const returnHome = () => {
  requests.post("/leaveRoom", {
    roomNumber: roomNumber,
    account: myInfo.account
  }).then((res) => {
    if (res.code === 200) {
      router.push("/home")
    } else {
      console.log("离开房间失败")
    }
  })
}

const changePrepareStatus = () => {
  if (pageVar.isPrepared) {
    myInfo.isPrepared = 0
  } else {
    myInfo.isPrepared = 1
  }
  userStore.websocketConn.send(JSON.stringify({
    path: '/prepare',
    data: {
      account: userStore.userInfo.account,
      roomNumber: roomNumber,
      flag: myInfo.isPrepared
    }
  }))
  pageVar.isPrepared = !pageVar.isPrepared
}

const startGame = () => {
  if (!pageVar.isHost) {
    return
  }
  console.log("开始游戏")
}
</script>

<style scoped>
.button-prepare {
  width: 100px; /* 按钮的宽度 */
  color: white; /* 按钮文本的颜色 */
  border: none; /* 移除按钮的默认边框 */
  padding: 10px 25px; /* 按钮的内边距，上下10px，左右40px */
  font-size: 16px; /* 按钮文本的字体大小 */
  border-radius: 50px; /* 圆角的半径，较大的值可使其呈现椭圆形 */
  cursor: pointer; /* 鼠标悬停时显示为指针 */
  transition: background-color 0.3s ease; /* 添加背景色的过渡效果 */
}

.button-notPrepared {
  background-color: #4CAF50; /* 按钮的背景颜色 */
}

.button-yesPrepared {
  background-color: #f44336; /* 按钮的背景颜色 */
}

.button-host {
  background-color: #5e635e; /* 按钮的背景颜色 */
}

.playerList-dialog-div > div {
  border-bottom: 1px solid #797474;
  padding: 10px;
}
</style>
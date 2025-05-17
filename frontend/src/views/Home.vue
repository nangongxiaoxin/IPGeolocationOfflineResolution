<template>
  <div class="container">
    <n-card title="IP 归属地信息检测">
      <n-space vertical :size="20">
        <!-- 输入框和按钮一行展示 -->
        <n-space align="center" justify="space-between" style="width: 100%; flex-wrap: nowrap;">
          <n-input
              v-model:value="ip"
              placeholder="请输入 IP 地址"
              style="width: 60vw"
          />
          <n-button
              :disabled="!ip"
              type="primary"
              :loading="loading"
              @click="getIPInfo"
          >
            获取 IP 地址信息
          </n-button>
        </n-space>

        <!-- 返回信息展示区域 -->
        <div v-if="ipInfo">
          <n-tag style="margin-bottom: 5px" type="info">您的 IP 地址信息为：</n-tag>
          <n-input v-model:value="ipInfo" readonly />
        </div>
      </n-space>

      <n-space style="background-color: snow;margin-top: 10vw">
        <div style="margin-right: 5vw">
          <n-tag style="margin-bottom: 5px" type="info">后台API监听[默认地址为http://localhost:45555/info?ip=1.2.3.4]：</n-tag>
        </div>
        <div>
          <n-switch v-model:value="apiEnabled" @update:value="onToggleApi" :loading="switchLoading">
            <template #checked>
              API接口运行中...
            </template>
            <template #unchecked>
              API接口已关闭.
            </template>
          </n-switch>
        </div>
      </n-space>
    </n-card>
  </div>

</template>

<script setup>
import {ref} from "vue"
import { NInput, NButton, NCard } from 'naive-ui'

const ip = ref(null);
const ipInfo = ref("请输入IP信息");
const loading = ref(false)
const switchLoading = ref(false)
const apiEnabled = ref(false)


const getIPInfo = async () =>{
  loading.value = true
  try {
    ipInfo.value = await window.go.main.App.GetIpInfoService(ip.value);
  } catch (err) {
    ipInfo.value = 'Error: ' + err.message
  } finally {
    loading.value = false
  }
}

const onToggleApi = async (value) =>{
  if (value){
    openApi()
    switchLoading.value = true;
    await sleep(500)
    switchLoading.value = false;
  }else {
    closeApi()
    switchLoading.value = true;
    await sleep(500)
    switchLoading.value = false;
  }
}

function openApi(){
  window.go.main.App.StartAPIService();
}

function closeApi(){
  window.go.main.App.StopAPIService();
}

const sleep = (ms) => new Promise(resolve => setTimeout(resolve, ms));
</script>

<style scoped>
.container {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 90vh;
  padding: 20px;
  box-sizing: border-box;
}

.n-card {
  width: 80vw;
  height: 60vw;
  max-width: 1000px;
  max-height: 800px;
  margin: 0 auto;
}
</style>

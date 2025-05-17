<template>
  <div class="about-container">
    <div class="title">感谢使用。</div>
    <div class="subtitle">数据来源为 ip2region</div>
    <div class="subtitle">version 1.0</div>
  </div>
</template>

<script setup>
// 无需逻辑，仅展示内容
</script>

<style scoped>
.about-container {
  display: flex;
  flex-direction: column;
  align-items: center;      /* 水平居中 */
  justify-content: center;  /* 垂直居中（可选）*/
  padding: 30px;
  text-align: center;
}

.title {
  font-size: 28px;
  font-weight: bold;
  margin-bottom: 16px;
}

.subtitle {
  font-size: 20px;
  font-weight: 600;
  color: #555;
}
</style>

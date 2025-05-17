<template>
  <div id="app">
    <nav>
      <index-top-menu></index-top-menu>
    </nav>



<!--    显示路由页面-->
    <router-view />
  </div>
</template>

<script>
import IndexTopMenu from "./components/IndexTopMenu.vue";
import {defineComponent} from "vue";

export default defineComponent({
  components: {IndexTopMenu}
})

</script>

<style scoped>
nav {
  margin-bottom: 20px;
}
</style>

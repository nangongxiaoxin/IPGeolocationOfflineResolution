<template>
  <n-card :embedded="true" :bordered="false" size="small">
    <n-menu
        :options="menuOptions"
        mode="horizontal"
        responsive
    />
  </n-card>

</template>

<script>
import {
  BookOutline as BookIcon,
  PersonOutline as PersonIcon,
  WineOutline as WineIcon,
  HomeOutline as HomeIcon,
  AlertCircleOutline as AboutIcon,
} from "@vicons/ionicons5";
import { NIcon } from "naive-ui";
import { defineComponent, h, ref } from "vue";
import { RouterLink } from "vue-router";

function renderIcon(icon) {
  return () => h(NIcon, null, { default: () => h(icon) });
}

const menuOptions = [
  {
    label: () => h(
        RouterLink,
        {
          to: {
            name: "Home",
          }
        },
        { default: () => "首页" }
    ),
    key: "homePage-router-link",
    icon: renderIcon(HomeIcon)
  },
  {
    label: () => h(
        RouterLink,
        {
          to: {
            name: "About",
          }
        },
        { default: () => "关于" }
    ),
    key: "aboutPage-router-link",
    icon: renderIcon(AboutIcon)
  },

];

export default defineComponent({
  setup() {
    return {
      menuOptions
    };
  }
});
</script>
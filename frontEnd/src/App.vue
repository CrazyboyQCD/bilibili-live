<template>
  <el-container>
    <el-aside width="230px">
      <SideMenu></SideMenu>
    </el-aside>
    <el-main>
      <router-view></router-view>
    </el-main>
  </el-container>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import SideMenu from "./components/SideMenu.vue";

export default defineComponent({
  name: "App",
  components: {
    SideMenu,
  },
  setup() {},
});
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>

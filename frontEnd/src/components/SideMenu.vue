<template>
  <div class="sidebar-logo">
    <a href="https://live.bilibili.com/" target="_blank">
      <img src="https://www.bilibili.com/favicon.ico" alt="logo" />
      <h1>BiliBili Monitor</h1>
    </a>
  </div>
  <el-menu :default-openeds="['1']">
    <el-submenu index="1">
      <template #title>
        <el-space wrap>
          <el-icon>
            <Monitor />
          </el-icon>
          <span>总览</span>
        </el-space>
      </template>
      <el-menu-item index="1-1">
        <router-link to="/">房间监控</router-link>
      </el-menu-item>
      <el-menu-item index="1-2">
        <router-link to="/areainfo">分区信息</router-link>
      </el-menu-item>
      <!-- <el-menu-item index="1-3">
        <router-link to="/livestatistics">直播统计</router-link>
      </el-menu-item> -->
      <!-- <el-menu-item index="1-4">
        <router-link to="/livewatch">观看直播</router-link>
      </el-menu-item> -->
    </el-submenu>
    <el-submenu index="2">
      <template #title>
        <el-space wwrap>
          <el-icon>
            <Setting />
          </el-icon>
          <span>操作</span>
        </el-space>
      </template>
      <el-menu-item index="2-1" @click="onAddRoom()">
        添加房间监控
      </el-menu-item>
      <el-menu-item index="2-2" @click="onAddArea()">
        添加分区
      </el-menu-item>
    </el-submenu>
  </el-menu>
  <AddRoomComponent></AddRoomComponent>
  <AddAreaComponent></AddAreaComponent>
</template>

<script lang="ts">
import mitter from '@/define/mitt'
import { defineComponent } from 'vue'
import AddRoomComponent from '@/components/AddRoom.vue'
import AddAreaComponent from '@/components/AddArea.vue'

export default defineComponent({
  name: 'SideMenu',
  components: {
    AddRoomComponent,
    AddAreaComponent
  },
  methods: {
    onAddRoom() {
      mitter.emit('changeAddRoomVisible', true);
    },
    onAddArea() {
      mitter.emit('showAddArea', true);
    }
  }
})
</script>


<style>
.sidebar-logo {
  position: relative;
  height: 64px;
  padding-left: 24px;
  overflow: hidden;
  line-height: 64px;
  background: #ffffff;
  transition: all 0.3s;
}

.sidebar-logo img {
  display: inline-block;
  height: 32px;
  width: 32px;
  vertical-align: middle;
}

.sidebar-logo h1 {
  display: inline-block;
  margin: 0 0 0 20px;
  color: rgb(62, 148, 230);
  font-weight: 600;
  font-size: 14px;
  font-family: Avenir, Helvetica Neue, Arial, Helvetica, sans-serif;
  vertical-align: middle;
}

a {
  text-decoration: none;
  color: #000000;
}
</style>

<template>
  <el-row :gutter="10">
    <el-col :span="12">
      <el-input v-model="searchTitle" placeholder="Search Title" clearable @change="searchRoomsWithTitle(searchTitle)">
        <template #append>
          <el-icon>
            <Search />
          </el-icon>
        </template>
      </el-input>
    </el-col>
    <el-col :span="12">
      <el-input v-model="searchAnchor" placeholder="Search Anchor" clearable
        @change="searchRoomsWithAnchor(searchAnchor)">
        <template #append>
          <el-icon>
            <Search />
          </el-icon>
        </template>
      </el-input>
    </el-col>
  </el-row>
  <el-row :gutter="10">
    <el-col :span="2" v-for="area in areas" :key="area">
      <el-tag @click="onClickArea(area)">
        {{ area }}
      </el-tag>
    </el-col>
  </el-row>
  <el-row :gutter="10" v-if="showSubArea">
    <el-col :span="2" v-for="area in subAreas" :key="area">
      <el-tag @click="onClickSubArea(area)">
        {{ area }}
      </el-tag>
    </el-col>
  </el-row>
  <div class="mainContainer" v-show="isShowLive">
    <el-row justify="center">
    <el-tag v-show="isShowTitle" round effect="light">{{ liveTitle }}</el-tag>
    </el-row>
    <!-- <el-input v-model="liveurl" clearable></el-input> -->
    <div class="video-container">
      <div>
        <video id="videoElement" class="centeredVideo" controls autoplay>
        </video>
      </div>
    </div>
    <div class="controls">
      <!-- <el-button @click="videoPlayer()">Load</el-button> -->
      <button @click="flvPlayer.play()">Start</button>
      <button @click="flvPlayer.pause()">Pause</button>
      <button @click="flv_destroy()">Destroy</button>
    </div>
  </div>
  <el-row :gutter="5">
    <el-col :xs="16" :sm="10" :md="8" :lg="5" :xl="4" v-for="item in showrooms" :key="item.RoomID">
      <el-card shadow="always" style="--el-card-padding: 10px">
        <el-space direction="vertical" alignment="start">
          <img :src="item.LiveCover" style="width: 256px; height: 144px" class="image" lazy @click="watchLive(item)" />
          <span>
            <i class="el-icon-user" style="font-size: small">{{
                item.Popularity
            }}</i>
            <span style="font-size: small; padding-left: 20px">
              {{ item.ParentName }}-{{ item.AreaName }}
            </span>
          </span>
          <div>{{ item.Title.slice(0, 15) }}</div>
          <div style="color: #909399">
            <span>
              <div>
                <Star style="
                    width: 2em;
                    height: 2em;
                    position: relative;
                    top: 1px;
                    padding-right: 4px;
                  " @click="favorite(item)" />
                <el-avatar size="small" :src="item.UserCover" @click="blockUser(item)"></el-avatar>
                <el-link class="anchor-name" :href="uID2URL(item)">{{
                    item.Uname
                }}</el-link>
              </div>
            </span>
          </div>
        </el-space>
      </el-card>
    </el-col>
  </el-row>
  <el-pagination layout="prev, pager, next" :page-count="pages" :current-page="pageIndex"
    @update:current-page="onCurrentChange" />
</template>

<script lang="ts">
import { MonitorRoom } from "@/define/interface";
import { defineComponent, ref } from "vue";
import { ip, port } from "@/define/const";
import { ElNotification } from "element-plus";
import { NotificationType } from "element-plus/lib/el-notification/src/notification.type";
import { Star } from "@element-plus/icons-vue";
import mitter from "@/define/mitt";
import flvjs from "flv.js";

export default defineComponent({
  setup() { },
  components: {
    Star,
  },
  data() {
    const searchTitle = ref("");
    const searchAnchor = ref("");
    let areas: string[] = [];
    let subAreas: string[] = [];
    let allrooms: MonitorRoom[] = [];
    let arearooms: MonitorRoom[] = [];
    let showrooms: MonitorRoom[] = [];
    let showSubArea: boolean = false;
    const perPageNum: number = 30;
    let pages: number = 0;
    let pageIndex: number = 1;
    let isShowLive: boolean = false;
    let liveurl: string = "";
    let flvPlayer = {} as flvjs.Player;
    let isShowTitle: boolean = false;
    let liveTitle: string = ''
    return {
      searchTitle,
      searchAnchor,
      areas,
      subAreas,
      allrooms,
      arearooms,
      showrooms,
      showSubArea,
      perPageNum,
      pages,
      pageIndex,
      isShowLive,
      liveurl,
      flvPlayer,
      isShowTitle,
      liveTitle
    };
  },
  mounted() {
    this.getAreaInfo(true);
    setInterval(this.getAreaInfo, 1000);
  },
  methods: {
    getAreaInfo(give: boolean = false) {
      fetch(`${ip}:${port}/areainfos`, {
        method: "get",
        headers: {
          "Access-Control-Allow-Origin": "*",
          "Access-Control-Allow-Methods":
            "POST, GET, OPTIONS, PUT, DELETE,UPDATE",
          "Access-Control-Allow-Headers":
            "Access-Control-Allow-Methods, Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers",
          "Access-Control-Allow-Credentials": "true",
          "Content-Type": "application/json; charset=utf-8",
        },
      })
        .then((res) => {
          if (!res.ok) {
            console.log("failed");
          }
          return res.json();
        })
        .then((json) => {
          let data: MonitorRoom[] = [];
          let tmpareas: string[] = [];
          for (let key in json) {
            data.push(json[key]);
            if (tmpareas.indexOf(json[key]["ParentName"]) === -1) {
              tmpareas.push(json[key]["ParentName"]);
            }
          }
          this.allrooms = data;
          this.areas = tmpareas;
          if (give) {
            this.arearooms = data;
            this.showrooms = this.arearooms.slice(0, this.perPageNum);
            this.pages = ((this.arearooms.length / this.perPageNum) | 0) + 1;
            this.pageIndex = 1;
          }
          // console.log(this.roomTable);
        });
    },
    onClickArea(area: string) {
      let tmpsubareas: string[] = [];
      let tmpdata: MonitorRoom[] = [];
      for (let key in this.allrooms) {
        if (this.allrooms[key]["ParentName"] === area) {
          tmpdata.push(this.allrooms[key]);
          if (tmpsubareas.indexOf(this.allrooms[key]["AreaName"]) === -1) {
            tmpsubareas.push(this.allrooms[key]["AreaName"]);
            this.showSubArea = true;
          }
        }
      }
      this.arearooms = tmpdata;
      this.subAreas = tmpsubareas;
      this.showrooms = this.arearooms.slice(0, this.perPageNum);
      this.pages = ((this.arearooms.length / this.perPageNum) | 0) + 1;
      this.pageIndex = 1;
    },
    onClickSubArea(subarea: string) {
      let tmpdata: MonitorRoom[] = [];
      for (let key in this.allrooms) {
        if (this.allrooms[key]["AreaName"] === subarea) {
          tmpdata.push(this.allrooms[key]);
        }
      }
      this.arearooms = tmpdata;
      this.showrooms = this.arearooms.slice(0, this.perPageNum);
      this.pages = ((this.arearooms.length / this.perPageNum) | 0) + 1;
      this.pageIndex = 1;
    },
    onCurrentChange(index: number) {
      this.showrooms = this.arearooms.slice(
        (index - 1) * this.perPageNum,
        index * this.perPageNum
      );
      this.pageIndex = index;
    },
    uID2URL(item: MonitorRoom): string {
      return `https://space.bilibili.com/${item.UID}`;
    },
    blockUser(item: MonitorRoom) {
      fetch(`${ip}:${port}/blockroom`, {
        method: "post",
        headers: {
          "Access-Control-Allow-Origin": "*",
          "Access-Control-Allow-Methods":
            "POST, GET, OPTIONS, PUT, DELETE,UPDATE",
          "Access-Control-Allow-Headers":
            "Access-Control-Allow-Methods, Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers",
          "Access-Control-Allow-Credentials": "true",
          "Content-Type": "application/json; charset=utf-8",
        },
        body: JSON.stringify({
          roomID: item.RoomID,
        }),
      })
        .then((res) => res.json())
        .then((json) => {
          console.log(json);
          let ttype: NotificationType = "";
          let title: string = "";
          let msg: string = "";
          if (json["msg"]) {
            ttype = "success";
            title = "Success";
            msg = "BLock user success";
          } else {
            ttype = "error";
            title = "Failed";
            msg = "BLock user failed";
          }
          ElNotification({
            title: title,
            message: msg,
            type: ttype,
          });
        })
        .catch((err) => {
          console.log(err);
        });
    },
    favorite(item: MonitorRoom) {
      mitter.emit("addMonitorRoom", {
        RoomID: item.RoomID,
        Platform: "bilibili",
      });
    },
    searchRoomsWithTitle(title: string) {
      let data: MonitorRoom[] = [];
      for (let key in this.allrooms) {
        if (this.allrooms[key].Title.indexOf(title) !== -1) {
          data.push(this.allrooms[key]);
        }
      }
      this.arearooms = data;
      this.showrooms = this.arearooms.slice(0, this.perPageNum);
      this.pages = ((this.arearooms.length / this.perPageNum) | 0) + 1;
      this.pageIndex = 1;
    },
    searchRoomsWithAnchor(anchor: string) {
      let data: MonitorRoom[] = [];
      for (let key in this.allrooms) {
        if (this.allrooms[key].Uname.indexOf(anchor) !== -1) {
          data.push(this.allrooms[key]);
        }
      }
      this.arearooms = data;
      this.showrooms = this.arearooms.slice(0, this.perPageNum);
      this.pages = ((this.arearooms.length / this.perPageNum) | 0) + 1;
      this.pageIndex = 1;
    },
    watchLive(item: MonitorRoom) {
      fetch(`${ip}:${port}/getliveurl`, {
        method: "post",
        headers: {
          "Access-Control-Allow-Origin": "*",
          "Access-Control-Allow-Methods":
            "POST, GET, OPTIONS, PUT, DELETE,UPDATE",
          "Access-Control-Allow-Headers":
            "Access-Control-Allow-Methods, Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers",
          "Access-Control-Allow-Credentials": "true",
          "Content-Type": "application/json; charset=utf-8",
        },
        body: JSON.stringify({
          roomID: item.RoomID,
          platform: "bilibili"
        }),
      })
        .then((res) => res.json())
        .then((json) => {
          console.log(json);
          if (json['succ']) {
            // router.push('/livewatch')
            // mitter.emit('updateliveurl', json['liveurl'])
            this.liveurl = json['liveurl']
            this.videoPlayer(json['liveurl'])
            this.isShowLive = true;
            this.isShowTitle = true;
            this.liveTitle = `${item.Uname}-${item.Title}`
          } else {
            ElNotification({
              title: '直播流解析失败',
              type: 'error',
            });
          }
        })
        .catch((err) => {
          console.log(err);
        });
    },
    videoPlayer(liveurl: string) {
      if (flvjs.isSupported()) {
        console.log(liveurl)
        var videoElement = document.getElementById("videoElement");
        if (this.isShowLive) {
          this.flvPlayer.unload();
          this.flvPlayer.detachMediaElement();
          this.flvPlayer.destroy();
          // this.flvPlayer = null;
        }
        console.log(liveurl)
        this.flvPlayer = flvjs.createPlayer(
          {
            type: "flv",
            url: liveurl,
            isLive: true,
            cors: true,
          },
          {
            autoCleanupSourceBuffer: true,//对SourceBuffer进行自动清理
            autoCleanupMaxBackwardDuration: 12,//    当向后缓冲区持续时间超过此值（以秒为单位）时，请对SourceBuffer进行自动清理
            autoCleanupMinBackwardDuration: 8,//指示进行自动清除时为反向缓冲区保留的持续时间（以秒为单位）。
            enableStashBuffer: false, //关闭IO隐藏缓冲区
            lazyLoad: false,
          });
        this.flvPlayer.attachMediaElement(videoElement as HTMLMediaElement);
        this.flvPlayer.load();
        this.flvPlayer.volume = 0.8;
        this.flvPlayer.on(flvjs.Events.ERROR, (errorType: any, errorDetail: any, errorInfo: any) => {
          //视频出错后销毁重新创建
          console.log(errorType, errorDetail, errorInfo);
          if (this.isShowLive) {
            // clearInterval(timer)
            this.flvPlayer.pause();
            this.flvPlayer.unload();
            this.flvPlayer.detachMediaElement();
            this.flvPlayer.destroy();
            // flvPlayer= null;
            //重新加载当前停止的视频流，根据个人的方法来配置
            this.videoPlayer(this.liveurl);
          }
        });
        this.flvPlayer.play();
      }
    },
    flv_destroy() {
      this.flvPlayer.pause();
      this.flvPlayer.unload();
      this.flvPlayer.detachMediaElement();
      this.flvPlayer.destroy();
      this.isShowLive = false;
      this.isShowTitle = false;
      // this.flvPlayer = null;
    },
  },
});
</script>

<style>
.el-row {
  margin-bottom: 10px;
}

.anchor-name {
  position: relative;
  bottom: 7px;
  left: 10px;
}

video {
  height: 50%;
  width: 50%;
}
</style>

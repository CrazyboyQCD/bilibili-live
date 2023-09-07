<template>
  <el-drawer
    v-model="isShowRoomInfo"
    title="房间信息"
    size="45%"
    @close="onCancel()"
  >
    <el-descriptions class="margin-top" :column="2" :data="roomInfo" border>
      <el-descriptions-item>
        <template #label>
          <i class="el-icon-user"></i>
          房间ID
        </template>
        <el-link
          :type="status2Type(roomInfo.LiveStatus)"
          :href="roomID2URL(roomInfo)"
        >
          {{ roomInfo.RoomID }}
        </el-link>
      </el-descriptions-item>
      <el-descriptions-item>
        <template #label>
          <i class="el-icon-mobile-phone"></i>
          主播
        </template>
        {{ roomInfo.Uname }}
      </el-descriptions-item>
      <el-descriptions-item>
        <template #label>
          <i class="el-icon-location-outline"></i>
          分区
        </template>
        {{ roomInfo.AreaName }}
      </el-descriptions-item>
      <el-descriptions-item>
        <template #label>
          <i class="el-icon-tickets"></i>
          直播标题
        </template>
        {{ roomInfo.Title }}
      </el-descriptions-item>
      <el-descriptions-item>
        <template #label>
          <i class="el-icon-office-building"></i>
          直播状态
        </template>
        <el-tag size="small" :type="status2Type(roomInfo.LiveStatus)"
          >{{ status2Name(roomInfo.LiveStatus) }}
        </el-tag>
      </el-descriptions-item>
      <el-descriptions-item>
        <template #label>
          <i class="el-icon-office-building"></i>
          开播时间
        </template>
        {{ roomInfo.LiveStartTime }}
      </el-descriptions-item>
      <el-descriptions-item>
        <template #label>
          <i class="el-icon-office-building"></i>
          开播时长
        </template>
        {{ roomInfo.LiveTime }}
      </el-descriptions-item>
      <el-descriptions-item>
        <template #label>
          <i class="el-icon-office-building"></i>
          录制时长
        </template>
        {{ roomInfo.RecordTime }}
      </el-descriptions-item>
      <el-descriptions-item>
        <template #label>
          <i class="el-icon-office-building"></i>
          转码时长
        </template>
        {{ roomInfo.DecodeTime }}
      </el-descriptions-item>
      <el-descriptions-item>
        <template #label>
          <i class="el-icon-office-building"></i>
          上传时长
        </template>
        {{ roomInfo.UploadTime }}
      </el-descriptions-item>
      <el-descriptions-item>
        <template #label>
          <i class="el-icon-office-building"></i>
          当前状态
        </template>
        <el-tag size="small" :type="state2Type(roomInfo.State)"
          >{{ state2Name(roomInfo.State) }}
        </el-tag>
      </el-descriptions-item>
      <el-descriptions-item>
        <template #label>
          <i class="el-icon-office-building"></i>
          占用空间
        </template>
        null
      </el-descriptions-item>
    </el-descriptions>
    <el-divider>
      <span>编辑房间信息</span>
    </el-divider>
    <el-form :model="roomInfo">
      <el-row>
        <el-col :span="12">
          <el-form-item>
            <el-switch
              v-model="roomInfo.RecordMode"
              active-color="#13ce66"
              inactive-color="#ff4949"
              active-text="24小时录制"
              inactive-text="仅在设定时间录制"
            />
          </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-form-item>
            <el-switch
              v-model="roomInfo.DivideByTitle"
              active-color="#13ce66"
              inactive-color="#ff4949"
              active-text="当天录像按照标题分P"
              inactive-text="当天录像不分P"
            />
          </el-form-item>
        </el-col>
      </el-row>
      <el-row>
        <el-col :span="12">
          <el-form-item label="录制开始时间" label-width="120px">
            <el-time-picker
              v-model="starttime"
              :disabled="roomInfo.RecordMode"
              format="HH:mm"
              @change="getStartTime()"
              size="small"
              style="width: 100px"
            ></el-time-picker>
          </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-form-item label="录制结束时间" label-width="120px">
            <el-time-picker
              v-model="endtime"
              :disabled="roomInfo.RecordMode"
              format="HH:mm"
              @change="getEndTime()"
              size="small"
              style="width: 100px"
            ></el-time-picker>
          </el-form-item>
        </el-col>
      </el-row>
      <el-row>
        <el-col :span="12">
          <el-form-item>
            <el-switch
              v-model="roomInfo.AutoRecord"
              active-color="#13ce66"
              inactive-color="#ff4949"
              active-text="自动录制"
              inactive-text="仅监控"
            />
          </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-form-item>
            <el-switch
              v-model="roomInfo.AutoUpload"
              disabled
              active-color="#13ce66"
              inactive-color="#ff4949"
              active-text="自动上传"
              inactive-text="禁止上传"
            />
          </el-form-item>
        </el-col>
      </el-row>
      <el-row>
        <el-col :span="12">
          <el-form-item>
            <el-switch
              v-model="roomInfo.NeedM4a"
              active-color="#13ce66"
              inactive-color="#ff4949"
              active-text="转码生成视频和音频"
              inactive-text="仅生成视频"
            />
          </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-form-item>
            <el-switch
              v-model="roomInfo.Mp4Compress"
              active-color="#13ce66"
              inactive-color="#ff4949"
              active-text="视频质量压缩"
              inactive-text="不压缩视频"
            />
          </el-form-item>
        </el-col>
      </el-row>
      <el-row>
        <el-col :span="12">
          <el-form-item>
            <el-switch
              v-model="roomInfo.CleanUpRegular"
              active-color="#13ce66"
              inactive-color="#ff4949"
              active-text="定期清理录制内容"
              inactive-text="不清理生成的文件"
            />
          </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-form-item label="文件保留时间" label-width="120px">
            <el-input
              v-model="roomInfo.SaveDuration"
              autocomplete="off"
              :disabled="!roomInfo.CleanUpRegular"
            ></el-input>
          </el-form-item>
        </el-col>
      </el-row>
      <el-row>
        <el-col :span="12">
          <el-form-item>
            <el-switch
              v-model="roomInfo.AreaLock"
              active-color="#13ce66"
              inactive-color="#ff4949"
              active-text="只录制指定分区下的直播"
              inactive-text="不限制直播分区"
            />
          </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-form-item label="指定的分区名" label-width="120px">
            <el-input
              v-model="roomInfo.AreaLimit"
              autocomplete="off"
              :disabled="!roomInfo.AreaLock"
            >
            </el-input>
          </el-form-item>
        </el-col>
      </el-row>
      <el-row>
        <el-col :span="12">
          <el-form-item>
            <el-switch
              v-model="roomInfo.SaveDanmu"
              active-color="#13ce66"
              inactive-color="#ff4949"
              active-text="录制弹幕"
              inactive-text="不录制弹幕"
            />
          </el-form-item>
        </el-col>
      </el-row>
    </el-form>
    <!-- <template #footer> -->
    <el-button-group>
      <el-button type="danger" @click="onDelete()">
        <el-icon><Delete /></el-icon>删除
      </el-button>
      <el-button @click="onCancel()">
        <el-icon><Close /></el-icon>取消
      </el-button>
      <el-button type="primary" @click="onSubmit()"
        ><el-icon><Check /></el-icon>确定</el-button
      >
      <el-button type="primary" @click="onDecode()"
        ><el-icon><Promotion /></el-icon>立即转码</el-button
      >
    </el-button-group>
    <!-- </template> -->
  </el-drawer>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { RoomInfo } from "@/define/interface";
import {
  status2Type,
  status2Name,
  state2Type,
  state2Name,
} from "@/define/methods";
import mitter from "@/define/mitt";
import { ip, port } from "@/define/const";
import { ElNotification } from "element-plus";
import { NotificationType } from "element-plus/lib/el-notification/src/notification.type";

export default defineComponent({
  data() {
    let roomInfo: RoomInfo = {
      RoomID: "",
      Platform: "",
      LiveStatus: 0,
      State: 0,
      RecordMode: true,
    };
    return {
      isShowRoomInfo: false,
      roomInfo,
      status2Name,
      status2Type,
      state2Type,
      state2Name,
      starttime: new Date(),
      endtime: new Date(),
    };
  },
  mounted() {
    const handleVisible = (e: any) => {
      console.log(e);
      this.isShowRoomInfo = e;
    };
    mitter.on("changeRoomInfoVisible", handleVisible);
    mitter.on("changeRoomInfo", (e: any) => {
      console.log(e);
      this.roomInfo = e;
      this.starttime = new Date(
        `2022T${this.roomInfo.StartTime?.substr(
          0,
          2
        )}:${this.roomInfo.StartTime?.substr(
          2,
          2
        )}:${this.roomInfo.StartTime?.substr(4, 2)}`
      );
      this.endtime = new Date(
        `2022T${this.roomInfo.EndTime?.substr(
          0,
          2
        )}:${this.roomInfo.EndTime?.substr(
          2,
          2
        )}:${this.roomInfo.EndTime?.substr(4, 2)}`
      );
    });
  },
  methods: {
    roomID2URL(item: RoomInfo) {
      if (item.Platform == "bilibili") {
        return `https://live.bilibili.com/${item.RoomID}`;
      } else if (item.Platform == "douyin") {
        return `https://live.douyin.com/${item.RoomID}`;
      } else if (item.Platform == "huya") {
        return `https://www.huya.com/${item.RoomID}`;
      }
      return "";
    },
    onDelete() {
      fetch(`${ip}:${port}/roomhandle`, {
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
          handle: "delete",
          data: this.getSubmitInfo(),
        }),
      })
        .then((res) => res.json())
        .then((json) => {
          console.log(json);
          let ttype: NotificationType = '';
          let title: string = '';
          let msg: string = ''
          if (json['msg']) {
            ttype = 'success';
            title = 'Success';
            msg = 'Delete success';
          } else {
            ttype = 'error';
            title = 'Failed';
            msg = 'Delete failed';
          }
          ElNotification({
            title: title,
            message: msg,
            type: ttype,
          });
          this.isShowRoomInfo = false;
        })
        .catch((err) => {
          console.log(err);
        });
    },
    onSubmit() {
      fetch(`${ip}:${this.roomInfo.port}/roomhandle`, {
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
          handle: "edit",
          data: this.getSubmitInfo(),
        }),
      })
        .then((res) => res.json())
        .then((json) => {
          console.log(json);
          let ttype: NotificationType = '';
          let title: string = '';
          let msg: string = ''
          if (json['msg']) {
            ttype = 'success';
            title = 'Success';
            msg = 'Edit success';
          } else {
            ttype = 'error';
            title = 'Failed';
            msg = 'Edit failed';
          }
          ElNotification({
            title: title,
            message: msg,
            type: ttype,
          });
          this.isShowRoomInfo = false;
        })
        .catch((err) => {
          console.log(err);
        });
    },
    onDecode() {
      fetch(`${ip}:${this.roomInfo.port}/decode`, {
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
          roomID: this.roomInfo.RoomID,
        }),
      })
        .then((res) => res.json())
        .then((json) => {
          console.log(json);
          let ttype: NotificationType = '';
          let title: string = '';
          let msg: string = ''
          if (json['msg']) {
            ttype = 'success';
            title = 'Success';
            msg = 'Decode success';
          } else {
            ttype = 'error';
            title = 'Failed';
            msg = 'Decode failed';
          }
          ElNotification({
            title: title,
            message: msg,
            type: ttype,
          });
          this.isShowRoomInfo = false;
        })
        .catch((err) => {
          console.log(err);
        });
    },
    onCancel() {
      this.isShowRoomInfo = false;
      console.log(this.roomInfo.StartTime);
      console.log(this.roomInfo.EndTime);
    },
    getStartTime() {
      let tmp = this.starttime.toTimeString().split(":");
      if (tmp?.length == 3) tmp[2] = "00";
      this.roomInfo.StartTime = tmp?.join("");
    },
    getEndTime() {
      let tmp = this.endtime.toTimeString().split(":");
      if (tmp?.length == 3) tmp[2] = "00";
      this.roomInfo.EndTime = tmp?.join("");
    },
    getSubmitInfo() {
      return {
        platform: this.roomInfo.Platform,
        roomID: this.roomInfo.RoomID,
        recordMode: this.roomInfo.RecordMode,
        startTime: this.roomInfo.StartTime,
        endTime: this.roomInfo.EndTime,
        autoRecord: this.roomInfo.AutoRecord,
        autoUpload: this.roomInfo.AutoUpload,
        needM4a: this.roomInfo.NeedM4a,
        mp4Compress: this.roomInfo.Mp4Compress,
        divideByTitle: this.roomInfo.DivideByTitle,
        cleanUpRegular: this.roomInfo.CleanUpRegular,
        saveDuration: this.roomInfo.SaveDuration,
        areaLock: this.roomInfo.AreaLock,
        areaLimit: this.roomInfo.AreaLimit,
        saveDanmu: this.roomInfo.SaveDanmu,
        name: this.roomInfo.Name,
        cookies: this.roomInfo.Cookies
      };
    },
  },
});
</script>

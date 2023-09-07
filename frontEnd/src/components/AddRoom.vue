<template>
  <el-drawer v-model="isShowAddRoom" title="添加房间" size="40%" direction="ltr">
    <el-form :model="roomInfo">
      <el-row>
        <el-col :span="6">
          <el-form-item label="房间ID" label-width="60px">
            <el-input v-model="roomInfo.roomID" autocomplete="off" style="width: 150px"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="2"></el-col>
        <el-col :span="6">
          <el-form-item label="主播名" label-width="60px">
            <el-input v-model="roomInfo.name" autocomplete="off" style="width: 150px"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="2"></el-col>
        <el-col :span="6">
          <el-form-item label="平台" label-width="60px">
            <el-input v-model="roomInfo.platform" style="width: 150px"></el-input>
          </el-form-item>
        </el-col>
      </el-row>

      <el-row>
        <el-col :span="12">
          <el-form-item>
            <el-switch v-model="roomInfo.recordMode" active-color="#13ce66" inactive-color="#ff4949"
              active-text="24小时录制" inactive-text="仅在设定时间录制" />
          </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-form-item>
            <el-switch v-model="roomInfo.divideByTitle" active-color="#13ce66" inactive-color="#ff4949"
              active-text="当天录像按照标题分P" inactive-text="当天录像不分P" />
          </el-form-item>
        </el-col>
      </el-row>
      <el-row>
        <el-col :span="6">
          <el-form-item label="录制开始时间" label-width="120px">
            <el-time-picker v-model="startTime" :disabled="roomInfo.recordMode" size="small" format="HH:mm"
              style="width: 100px" @change="getStartTime()"></el-time-picker>
          </el-form-item>
        </el-col>
        <el-col :span="4"></el-col>
        <el-col :span="8">
          <el-form-item label="录制结束时间" label-width="120px">
            <el-time-picker v-model="endTime" :disabled="roomInfo.recordMode" size="small" format="HH:mm"
              style="width: 100px" @change="getEndTime()"></el-time-picker>
          </el-form-item>
        </el-col>
      </el-row>
      <el-row>
        <el-col :span="12">
          <el-form-item>
            <el-switch v-model="roomInfo.autoRecord" active-color="#13ce66" inactive-color="#ff4949" active-text="自动录制"
              inactive-text="仅监控" />
          </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-form-item>
            <el-switch v-model="roomInfo.autoUpload" disabled active-color="#13ce66" inactive-color="#ff4949"
              active-text="自动上传" inactive-text="禁止上传" />
          </el-form-item>
        </el-col>
      </el-row>
      <el-row>
        <el-col :span="12">
          <el-form-item>
            <el-switch v-model="roomInfo.needM4a" active-color="#13ce66" inactive-color="#ff4949"
              active-text="转码生成视频和音频" inactive-text="仅生成视频" />
          </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-form-item>
            <el-switch v-model="roomInfo.mp4Compress" active-color="#13ce66" inactive-color="#ff4949"
              active-text="视频质量压缩" inactive-text="不压缩视频" />
          </el-form-item>
        </el-col>
      </el-row>
      <el-row>
        <el-col :span="12">
          <el-form-item>
            <el-switch v-model="roomInfo.cleanUpRegular" active-color="#13ce66" inactive-color="#ff4949"
              active-text="定期清理录制内容" inactive-text="不清理生成的文件" />
          </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-form-item label="文件保留时间" label-width="100px">
            <el-input v-model="roomInfo.saveDuration" :disabled="!roomInfo.cleanUpRegular" style="width: 100px">
            </el-input>
          </el-form-item>
        </el-col>
      </el-row>
      <el-row>
        <el-col :span="12">
          <el-form-item>
            <el-switch v-model="roomInfo.areaLock" active-color="#13ce66" inactive-color="#ff4949"
              active-text="只录制指定分区下的直播" inactive-text="不限制直播分区" />
          </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-form-item label="指定的分区名" label-width="100px">
            <el-input v-model="roomInfo.areaLimit" :disabled="!roomInfo.areaLock" style="width: 220px">
            </el-input>
          </el-form-item>
        </el-col>
      </el-row>
      <el-row>
        <el-col :span="12">
          <el-form-item>
            <el-switch v-model="roomInfo.saveDanmu" active-color="#13ce66" inactive-color="#ff4949" active-text="录制弹幕"
              inactive-text="不录制弹幕" />
          </el-form-item>
        </el-col>
      </el-row>
      <el-row>
        <el-col :span="14">
          <el-form-item label="Cookies" label-width="40px">
            <el-input v-model="roomInfo.cookies" autocomplete="off" style="width: 300px"></el-input>
          </el-form-item>
        </el-col>
      </el-row>
    </el-form>
    <div class="drawer-footer">
      <el-button-group>
        <el-button icon="el-icon-circle-close" @click="isShowAddRoom = false">取消</el-button>
        <el-button icon="el-icon-circle-check" type="primary" @click="onSubmit()">确定</el-button>
      </el-button-group>
    </div>
  </el-drawer>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { submitInfo } from "@/define/interface";
import mitter from "@/define/mitt";
import { ip, port } from "@/define/const";
import { ElNotification } from "element-plus";
import { NotificationType } from "element-plus/lib/el-notification/src/notification.type";

export default defineComponent({
  data() {
    let roomInfo: submitInfo = {
      platform: "",
      roomID: "",
      recordMode: true,
      startTime: "220000",
      endTime: "020000",
      autoRecord: true,
      autoUpload: false,
      needM4a: true,
      mp4Compress: true,
      divideByTitle: true,
      cleanUpRegular: true,
      saveDuration: "3d",
      areaLock: false,
      areaLimit: "放松电台,聊天电台,萌宅领域",
      saveDanmu: true,
      name: "",
      cookies: ""
    };
    return {
      roomInfo,
      isShowAddRoom: false,
      startTime: new Date(),
      endTime: new Date(),
    };
  },
  mounted() {
    mitter.on("changeAddRoomVisible", (e: any) => {
      this.isShowAddRoom = e;
    });
    mitter.on('addMonitorRoom', (e: any) => {
      this.roomInfo.roomID = e.RoomID;
      this.roomInfo.platform = e.Platform;
      this.isShowAddRoom = true;
    })
  },
  methods: {
    onSubmit() {
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
          handle: "add",
          data: this.roomInfo,
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
            msg = 'Add success';
          } else {
            ttype = 'error';
            title = 'Failed';
            msg = 'Add failed';
          }
          ElNotification({
            title: title,
            message: msg,
            type: ttype,
          });
          this.isShowAddRoom = false;
        })
        .catch((err) => {
          console.log(err);
        });
    },
    getStartTime() {
      let tmp = this.startTime.toTimeString().split(":");
      if (tmp?.length == 3) tmp[2] = "00";
      this.roomInfo.startTime = tmp?.join("");
    },
    getEndTime() {
      let tmp = this.endTime.toTimeString().split(":");
      if (tmp?.length == 3) tmp[2] = "00";
      this.roomInfo.endTime = tmp?.join("");
    },
  },
});
</script>
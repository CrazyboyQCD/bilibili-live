<template>
  <el-row :gutter="10">
    <el-col :span="5">
      <el-card class="box-card">
        <template #header>
          <div class="card-header">
            <span>总体信息</span>
          </div>
        </template>
        <div>下载量: {{ totalDownload }}</div>
        <div>已录制文件数: {{ fileNum }}</div>
      </el-card>
    </el-col>
    <el-col :xs="8" :sm="6" :md="4" :lg="3" :xl="1">
      <el-result :icon="result"></el-result>
    </el-col>
  </el-row>
  <el-table :default-sort="{ prop: 'State', order: 'descending' }" :data="roomTable">
    <el-table-column prop="RoomID" label="房间ID" sortable>
      <template #default="scope">
        <el-link :href="roomID2URL(scope.row)" type="info" :underline="false">{{ scope.row.RoomID }}</el-link>
      </template>
    </el-table-column>
    <el-table-column prop="Uname" label="主播">
      <template #default="scope">
        <router-link :to="'/livestatistics/' + scope.row.Uname">{{ scope.row.Uname }}</router-link>
      </template>
    </el-table-column>
    <el-table-column prop="Platform" label="平台" :filters="platformFilter" :filter-method="filterPlatform" />
    <el-table-column prop="AreaName" label="分区" :filters="areaFilter" :filter-method="filterArea" />
    <el-table-column prop="Title" label="直播标题" />
    <el-table-column prop="LiveStatus" label="直播状态" sortable>
      <template #default="scope">
        <el-tag :type="status2Type(scope.row.LiveStatus)">{{ status2Name(scope.row.LiveStatus) }}
        </el-tag>
      </template>
    </el-table-column>
    <el-table-column prop="LiveStartTime" label="开播时间" />
    <el-table-column prop="LiveTime" label="直播时长" />
    <el-table-column prop="RecordTime" label="录制时长" />
    <el-table-column prop="State" label="录制状态" sortable>
      <template #default="scope">
        <el-tag :type="state2Type(scope.row.State)">{{ state2Name(scope.row.State) }}
        </el-tag>
      </template>
    </el-table-column>
    <el-table-column label="操作" fixed="right">
      <template #default="scope">
        <el-button link type="primary" size="small" @click="onEdit(scope.row)">E</el-button>
        <el-button link type="primary" size="small" @click="refresh(scope.row)">R</el-button>
      </template>
    </el-table-column>
  </el-table>
  <RoomInfoComponent></RoomInfoComponent>
</template>

<script lang="ts">
import { RoomInfo } from "@/define/interface";
import { defineComponent, ref, Ref } from "vue";
import { ip, port } from "@/define/const";
import {
  status2Type,
  status2Name,
  state2Type,
  state2Name,
} from "@/define/methods";
import RoomInfoComponent from "@/components/RoomInfo.vue"
import mitter from '@/define/mitt'
import { ElNotification } from "element-plus";

interface filter {
  text: string;
  value: string;
}

export default defineComponent({
  name: "RoomMonitor",
  setup() {
    // return {roomTable}
  },
  components: {
    RoomInfoComponent
  },
  data() {
    let roomTable: RoomInfo[] = [];
    let result: string = "success";
    let totalDownload: string = "";
    let fileNum: number = 0;
    let platformFilter: filter[] = [
      {
        text: "抖音",
        value: "douyin",
      },
      {
        text: "虎牙",
        value: "huya",
      },
      {
        text: "哔哩哔哩",
        value: "bilibili",
      },
    ];
    let areaFilter: filter[] = [];
    let showRoomInfo: Ref<boolean> = ref(false);
    let room: Ref<RoomInfo> = ref(roomTable[0]);
    return {
      roomTable,
      result,
      totalDownload,
      fileNum,
      platformFilter,
      areaFilter,
      status2Name,
      status2Type,
      state2Type,
      state2Name,
      showRoomInfo,
      room,
    };
  },
  mounted() {
    this.getRooms();
    setInterval(this.getRooms, 1000);
    this.getInfos();
    setInterval(this.getInfos, 1000);
  },
  methods: {
    filterPlatform(value: string, row: RoomInfo) {
      return row.Platform === value;
    },
    filterArea(value: string, row: RoomInfo) {
      return row.AreaName === value;
    },
    getRooms() {
      fetch(`${ip}:${port}/livestatus`, {
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
        .then((res) => res.json())
        .then((json) => {
          let data: RoomInfo[] = [];
          let tmp: string[] = [];
          let tmpareas: filter[] = [];
          for (let key in json) {
            json[key]["LiveStartTime2"] = json[key]["LiveStartTime"];
            json[key]["LiveStartTime"] = this.date2string(
              new Date(parseInt(json[key]["LiveStartTime2"]) * 1000)
            );
            json[key].LiveTime = this.getTimeMiuns(
              json[key].LiveStartTime2,
              "0"
            );
            json[key].RecordTime = this.getTimeMiuns(
              json[key].RecordStartTime,
              json[key].RecordEndTime
            );
            json[key].DecodeTime = this.getTimeMiuns(
              json[key].DecodeStartTime,
              json[key].DecodeEndTime
            );
            json[key].UploadTime = this.getTimeMiuns(
              json[key].UploadStartTime,
              json[key].UploadEndTime
            );
            
            json[key].port = port;
            data.push(json[key]);
            if (tmp.indexOf(json[key]["AreaName"]) === -1) {
              tmp.push(json[key]["AreaName"]);
              let t: filter = {
                text: json[key]["AreaName"],
                value: json[key]["AreaName"],
              };
              tmpareas.push(t);
            }
          }
          this.roomTable = data;
          this.result = "success";
          this.areaFilter = tmpareas;
        })
        .catch((err) => {
          console.log(err);
          this.result = "error";
          console.log(this.result);
        });
    },
    fetchInfo(pport: string) {
      fetch(`${ip}:${pport}/livestatus`, {
        method: "get",
        headers: {
          "Access-Control-Allow-Origin": "*",
          "Access-Control-Allow-Methods":
            "POST, GET, OPTIONS, PUT, DELETE,UPDATE",
          "Access-Control-Allow-Headers":
            "Access-Control-Allow-Methods, Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers, Access-Control-Allow-Origin",
          "Access-Control-Allow-Credentials": "true",
          "Content-Type": "application/json; charset=utf-8",
        },
      })
        .then((res) => res.json())
        .then((json) => {
          for (let key in json) {
            json[key]["LiveStartTime2"] = json[key]["LiveStartTime"];
            json[key]["LiveStartTime"] = this.date2string(
              new Date(parseInt(json[key]["LiveStartTime2"]) * 1000)
            );
            json[key].LiveTime = this.getTimeMiuns(
              json[key].LiveStartTime2,
              "0"
            );
            json[key].RecordTime = this.getTimeMiuns(
              json[key].RecordStartTime,
              json[key].RecordEndTime
            );
            json[key].DecodeTime = this.getTimeMiuns(
              json[key].DecodeStartTime,
              json[key].DecodeEndTime
            );
            json[key].UploadTime = this.getTimeMiuns(
              json[key].UploadStartTime,
              json[key].UploadEndTime
            );
            json[key].port = pport
            this.roomTable.push(json[key]);
          }
        })
        .catch((err) => {
          console.log(err);
          this.result = "error";
          console.log(this.result);
        });
    },
    getInfos() {
      fetch(`${ip}:${port}/basestatus`, {
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
        .then((res) => res.json())
        .then((json) => {
          this.totalDownload = this.getReadableSizeString(
            json["totalDownload"]
          );
          this.fileNum = json["fileNum"];
          // console.log(this.roomTable);
        })
        .catch((err) => {
          console.log(err);
          this.result = "error";
          console.log(this.result);
        });
    },
    getReadableSizeString(fileSizeInBytes: number) {
      var i = -1;
      var byteUnits = [" K", " M", " G", " T", "P", "E", "Z", "Y"];
      do {
        fileSizeInBytes = fileSizeInBytes / 1024;
        i++;
      } while (fileSizeInBytes > 1024);

      return Math.max(fileSizeInBytes, 0.1).toFixed(1) + byteUnits[i];
    },
    stillTwo(num: number) {
      return ("0" + num).substr(-2);
    },
    date2string(date: Date) {
      return `${date.getFullYear()}-${this.stillTwo(
        date.getMonth() + 1
      )}-${this.stillTwo(date.getDate())} ${this.stillTwo(
        date.getHours()
      )}:${this.stillTwo(date.getMinutes())}:${this.stillTwo(
        date.getSeconds()
      )}`;
    },
    getTimeMiuns(st: string, et: string) {
      let nTime = 0;
      if (parseInt(st) < parseInt(et)) {
        nTime = parseInt(et) - parseInt(st);
      } else if (parseInt(st) > 0) {
        let dn = new Date();
        let start = new Date(parseInt(st) * 1000);
        nTime = dn.getTime() - start.getTime();
      }
      nTime = Math.floor(nTime / 1000);
      let day = Math.floor(nTime / 86400);
      let hour = Math.floor((nTime % 86400) / 3600);
      let minute = Math.floor(((nTime % 86400) % 3600) / 60);
      let second = nTime % 60;
      return `${day}天 ${hour}时 ${minute}分 ${second} 秒`;
    },
    onEdit(room: RoomInfo) {
      this.showRoomInfo = true;
      this.room = room;
      console.log(room)
      mitter.emit('changeRoomInfoVisible', this.showRoomInfo);
      mitter.emit('changeRoomInfo', this.room);
    },
    refresh(room: RoomInfo) {
      console.log(room.Uname)
      fetch(`${ip}:${room.port}/refresh/${room.RoomID}`, {
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
        .then((res) => res.json())
        .then((json) => {
           ElNotification({
            title: 'Refresh success',
            message: json['msg'],
            type: 'success',
          });
        })
        .catch((err) => {
          console.log(err);
        });
    },
    roomID2URL(item: RoomInfo) {
      if (item.Platform == 'bilibili') {
        return `https://live.bilibili.com/${item.RoomID}`
      } else if (item.Platform == 'douyin') {
        return `https://live.douyin.com/${item.RoomID}`
      } else if (item.Platform == 'huya') {
        return `https://www.huya.com/${item.RoomID}`
      }
      return ''
    }
  },
});
</script>

<template>
    <el-dialog v-model="isShowAddArea" title="添加分区">
        <el-form :model="areaInfo">
            <el-row :gutter="10">
                <el-col :span="12">
                    <el-form-item label="平台">
                        <el-input v-model="areaInfo.platform"></el-input>
                    </el-form-item>
                </el-col>
                <el-col :span="12">
                    <el-form-item label="分区名">
                        <el-input v-model="areaInfo.areaName"></el-input>
                    </el-form-item>
                </el-col>
            </el-row>
            <el-row :gutter="10">
                <el-col :span="12">
                    <el-form-item label="分区ID">
                        <el-input v-model="areaInfo.areaID"></el-input>
                    </el-form-item>
                </el-col>
                <el-col :span="12">
                    <el-form-item label="父ID">
                        <el-input v-model="areaInfo.parentID"></el-input>
                    </el-form-item>
                </el-col>
            </el-row>
        </el-form>
        <template #footer>
            <el-button @click="isShowAddArea = false">取消</el-button>
            <el-button type="primary" @click="onSubmit()">提交</el-button>
        </template>
    </el-dialog>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import mitter from "@/define/mitt";
import { ip, port } from "@/define/const";
import { ElNotification } from "element-plus";
import { NotificationType } from "element-plus/lib/el-notification/src/notification.type";

interface AreaInfo {
    areaName: string;
    areaID: string;
    parentID: string;
    platform: string;
}

export default defineComponent({
    data() {
        let isShowAddArea: boolean = false;
        let areaInfo: AreaInfo = {
            areaName: "",
            areaID: "",
            parentID: "",
            platform: "",
        };
        return {
            isShowAddArea,
            areaInfo,
        }
    },
    mounted() {
        mitter.on('showAddArea', (e: any) => {
            this.isShowAddArea = e;
        })
    },
    methods: {
        onSubmit() {
            fetch(`${ip}:${port}/areahandle`, {
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
                    data: this.areaInfo,
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
                        msg = 'Add area success';
                    } else {
                        ttype = 'error';
                        title = 'Failed';
                        msg = 'Add area failed';
                    }
                    ElNotification({
                        title: title,
                        message: msg,
                        type: ttype,
                    });
                    this.isShowAddArea = false;
                })
                .catch((err) => {
                    console.log(err);
                });
        }
    }
})
</script>

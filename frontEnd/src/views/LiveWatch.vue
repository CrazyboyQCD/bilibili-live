<template>
    <div class="mainContainer">
        <!-- <el-input v-model="liveurl" clearable></el-input> -->
        <div class="video-container">
            <div>
                <video id="videoElement" class="centeredVideo" controls autoplay>
                </video>
            </div>
        </div>
        <div class="controls">
            <el-button @click="videoPlayer()">Load</el-button>
            <button onclick="flv_start()">Start</button>
            <button onclick="flv_pause()">Pause</button>
            <button @click="flv_destroy()">Destroy</button>
        </div>
    </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import flvjs from "flv.js";
import mitter from "@/define/mitt";

export default defineComponent({
    name: "app",
    components: {},
    data() {
        let flvPlayer = {} as flvjs.Player;
        let liveurl: string = "";
        return {
            flvPlayer,
            liveurl
        }
    },
    updated() {
        mitter.on('updateliveurl', (e: any) => {
            this.liveurl = e;
            // console.log(this.liveurl)
        });
    },
    methods: {
        videoPlayer() {
            if (flvjs.isSupported()) {
                console.log(this.liveurl)
                var videoElement = document.getElementById("videoElement");
                if (typeof this.flvPlayer !== "undefined") {
                    if (this.flvPlayer != {} as flvjs.Player) {
                        this.flvPlayer.unload();
                        this.flvPlayer.detachMediaElement();
                        this.flvPlayer.destroy();
                        // this.flvPlayer = null;
                    }
                }
                console.log(this.liveurl)
                this.flvPlayer = flvjs.createPlayer(
                    {
                        type: "flv",
                        url: this.liveurl,
                        isLive: true,
                        cors: true,
                    },
                    {
                        headers: {

                            "Referer": "https://live.bilibili.com/"
                        }
                    });
                this.flvPlayer.attachMediaElement(videoElement as HTMLMediaElement);
                this.flvPlayer.load();
                this.flvPlayer.play();
            }
        },
        flv_destroy() {
            this.flvPlayer.pause();
            this.flvPlayer.unload();
            this.flvPlayer.detachMediaElement();
            this.flvPlayer.destroy();
            // this.flvPlayer = null;
        },
    },
});
</script>
<template>
    <el-dropdown>
        <span class="el-dropdown-link">
            Dropdown List<el-icon class="el-icon--right">
                <arrow-down />
            </el-icon>
        </span>
        <template #dropdown>
            <el-dropdown-menu>
                <el-dropdown-item v-for="brginfo in brgLst" :key="brginfo.timestamp" @click="fetchStatistics(brginfo)">
                    {{ date2string(new Date(brginfo.timestamp * 1000)) }} {{ brginfo.title }}
                </el-dropdown-item>

            </el-dropdown-menu>
        </template>
    </el-dropdown>
    <el-row justify="center">
        <el-tag v-show="showCharts" round effect="light">{{ selectedTitle }}</el-tag>
        <ArrowUpBold style="position:relative; width: 2em; height: 2em; top: 1px; left: 3px" @click="onClickPrev()">
        </ArrowUpBold>
        <ArrowDownBold style="position:relative; width: 2em; height: 2em; bottom: 2px; left: 6px"
            @click="onClickNext()"></ArrowDownBold>
    </el-row>
    <div style="width: 1200px; height: 500px">
        <v-chart v-show="showCharts" :autoresize="true" :option="medalChartOptions" @click="onClickMedal($event)"
            ref="medalChart">
        </v-chart>
    </div>
    <el-row>
        <div style="width: 800px; height: 500px">
            <v-chart v-show="showCharts" :autoresize="true" :option="revenueChartOptions" @click="onClickRich($event)">
            </v-chart>
        </div>
        <div style="width: 800px; height: 400px" v-show="showSengGiftsChart">
            <v-chart :autoresize="true" :option="sendGiftsChartOptions"></v-chart>
        </div>
    </el-row>
    <div style="width: 800px; height: 500px">
        <v-chart v-show="showWordCloud" :autoresize="true" :option="wordClodChartOptions"></v-chart>
    </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import { useRoute } from 'vue-router'
import { ip, port } from "@/define/const";
import VECharts from 'vue-echarts'
import { EChartsOption, ECBasicOption } from 'echarts/types/dist/shared';
import * as echarts from 'echarts/core';
import {
    TitleComponent,
    TooltipComponent,
    LegendComponent,
    GridComponent,
} from 'echarts/components';
import { PieChart, BarChart } from 'echarts/charts';
import { LabelLayout } from 'echarts/features';
import { CanvasRenderer } from 'echarts/renderers';
import { ECharts } from 'echarts/core';
import { ArrowUpBold, ArrowDownBold } from "@element-plus/icons-vue";
import 'echarts-wordcloud';
var randomColor = require('randomcolor');

echarts.use([
    TitleComponent,
    TooltipComponent,
    LegendComponent,
    PieChart,
    CanvasRenderer,
    LabelLayout,
    BarChart,
    GridComponent
]);

interface brgInfo {
    title: string,
    file_name: string,
    timestamp: number,
}

export default defineComponent({
    setup() {
        const route = useRoute();
        return {
            route,
        }
    },
    components: {
        'v-chart': VECharts,
        ArrowUpBold,
        ArrowDownBold,
    },
    data() {
        let brgLst: brgInfo[] = [];
        let showCharts: boolean = false;
        let showSengGiftsChart: boolean = false;
        let showWordCloud: boolean = false;
        let medalChartOptions: EChartsOption = {} as EChartsOption;
        let revenueChartOptions: EChartsOption = {} as EChartsOption;
        let sendGiftsChartOptions: EChartsOption = {} as EChartsOption;
        let wordClodChartOptions: any = {} as any;
        let initOptions = {
            renderer: 'canvas',
            width: 'auto',
            height: 'auto'
        };
        let selectedTitle: string = "";
        let receivedJson = {} as any;
        let menuIndex = 1;
        return {
            brgLst,
            showCharts,
            showSengGiftsChart,
            showWordCloud,
            medalChartOptions,
            revenueChartOptions,
            wordClodChartOptions,
            initOptions,
            sendGiftsChartOptions,
            selectedTitle,
            receivedJson,
            menuIndex
        }
    },
    mounted() {
        console.log(this.route.params.anchor)
        this.fetchBrgList()
    },
    methods: {
        fetchBrgList() {
            fetch(`${ip}:${port}/anchorlivebacklist/${this.route.params.anchor}`, {
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
                    let tmp: brgInfo[] = [];
                    for (let key in json) {
                        tmp.push(json[key])
                    }
                    this.brgLst = tmp;
                });
        },
        fetchStatistics(brginfo: brgInfo) {
            console.log(brginfo)
            for (let key in this.brgLst) {
                if (this.brgLst[key].file_name === brginfo.file_name) {
                    this.menuIndex = Number(key);
                }
            }
            this.selectedTitle = this.date2string(new Date(brginfo.timestamp * 1000)) + brginfo.title
            fetch(`${ip}:${port}/livebackstatistics`, {
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
                    anchor_name: this.route.params.anchor,
                    brg_file_name: brginfo.file_name,
                }),
            })
                .then((res) => res.json())
                .then((json) => {
                    console.log(json);
                    this.showCharts = true;
                    this.showSengGiftsChart = false;
                    this.setMedalChartOptions(json['medal_statistics']);
                    this.setRevenueChartOptions(json['revenue_statistics'])
                    this.receivedJson = json;
                })
                .catch((err) => {
                    console.log(err);
                });
            fetch(`${ip}:${port}/getwordcloud`, {
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
                    anchor_name: this.route.params.anchor,
                    brg_file_name: brginfo.file_name,
                }),
            })
                .then((res) => res.json())
                .then((json) => {
                    console.log(json);
                    if (json.length > 0) {
                        this.showWordCloud = true;
                        this.setWordCloud(json);
                    }
                })
                .catch((err) => {
                    console.log(err);
                });
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
        setMedalChartOptions(json: any) {
            this.medalChartOptions.title = {
                text: '直播间勋章统计',
                left: 'center',

            };
            this.medalChartOptions.tooltip = {
                trigger: 'item',
                formatter: '{b} : {c} ({d}%)'
            };
            let medals: string[] = []
            let data: any[] = [];
            for (let key in json) {
                medals.push(json[key]['medal_name'])
                data.push({ value: json[key]['num'], name: json[key]['medal_name'] })
            }
            this.medalChartOptions.legend = {
                orient: 'vertical',
                left: 'left',
                data: medals,
                type: 'scroll',
            }
            this.medalChartOptions.series = {
                type: 'pie',
                radius: '50%',
                center: ['40%', '50%'],
                selectedMode: 'single',
                // itemStyle: {
                //     borderColor: '#fff',
                //     borderWidth: 1
                // },
                data: data,
                emphasis: {
                    itemStyle: {
                        shadowBlur: 10,
                        shadowOffsetX: 0,
                        shadowColor: 'rgba(0, 0, 0, 0.5)'
                    }
                }
            }

        },
        setRevenueChartOptions(json: any) {
            this.revenueChartOptions.title = {
                text: `本次直播收到${json['total_revenue'] / 1000}元打赏`
            }
            this.revenueChartOptions.tooltip = {
                trigger: 'axis',
                axisPointer: {
                    type: 'shadow'
                }
            }
            this.revenueChartOptions.grid = {
                left: '3%',
                right: '4%',
                bottom: '3%',
                containLabel: true
            }
            this.revenueChartOptions.xAxis = {
                type: 'value',
                boundaryGap: [0, 0.01]
            }
            let ydata: string[] = [];
            let xdata: number[] = [];
            for (let key in json['rich_ranking']) {
                ydata.push(json['rich_ranking'][key]['uname'])
                xdata.push(json['rich_ranking'][key]['total_send_price'] / 1000)
            }
            this.revenueChartOptions.yAxis = {
                type: 'category',
                data: ydata.reverse(),
            };
            this.revenueChartOptions.series = {
                type: 'bar',
                data: xdata.reverse()
            }
        },
        setWordCloud(json: any) {
            this.wordClodChartOptions.tooltip = {
                trigger: 'item',
            };
            this.wordClodChartOptions.series = {
                type: 'wordCloud',
                shape: 'diamond',
                left: 'center',
                top: 'center',
                right: null,
                bottom: null,
                width: '70%',
                height: '80%',
                sizeRange: [18, 100],
                rotationRange: [0, 0],
                rotationStep: 45,
                gridSize: 10,
                drawOutOfBound: false,
                textStyle: {
                    // normal: {
                    fontFamily: 'PingFangSC-Regular, PingFang SC',
                    fontWeight: 'normal',
                    color: function () {
                        return randomColor({ hue: '#1785ff', count: 1 })[0];
                    }
                    // },
                },
                data: json,
                emphasis: {
                    focus: 'self',

                    textStyle: {
                        textShadowBlur: 10,
                        textShadowColor: '#333'
                    }
                },
            }
        },
        onClickMedal(param: any) {
            console.log(param.data.name)
            let data: any[] = [];
            for (let key in this.receivedJson['medal_statistics']) {
                if (this.receivedJson['medal_statistics'][key]['medal_name'] !== param.data.name) {
                    data.push({ name: this.receivedJson['medal_statistics'][key]['medal_name'], value: this.receivedJson['medal_statistics'][key]['num'] })
                } else {
                    data.push({
                        name: this.receivedJson['medal_statistics'][key]['medal_name'],
                        value: this.receivedJson['medal_statistics'][key]['num'],
                        label: {
                            alignTo: 'edge',
                            formatter: `{medal|${this.receivedJson['medal_statistics'][key]['medal_name']}}\n所属主播:${this.receivedJson['medal_statistics'][key]['anchor_name']}`,
                            // backgroundColor: '#eee',
                            // borderColor: '#777',
                            // borderWidth: 3,
                            // borderRadius: 4,
                            minMargin: 5,
                            // edgeDistance: 10,
                            lineHeight: 15,
                            rich: {
                                medal: {
                                    color: '#FFFFFF',
                                    backgroundColor: '#FA9DA6',
                                    padding: [2, 4],
                                    borderRadius: 2,
                                    align: 'center',
                                }
                            }
                        },
                        labelLine: {
                            length: 15,
                            length2: 0,
                            maxSurfaceAngle: 80
                        },
                    })
                }
            }
            let soptions: ECBasicOption = {
                series: {
                    type: 'pie',
                    radius: '50%',
                    center: ['50%', '50%'],
                    selectedMode: 'single',
                    itemStyle: {
                        borderColor: '#fff',
                        borderWidth: 1
                    },
                    data: data,
                    emphasis: {
                        itemStyle: {
                            shadowBlur: 10,
                            shadowOffsetX: 0,
                            shadowColor: 'rgba(0, 0, 0, 0.5)'
                        }
                    }
                }
            }
            let medalChart: ECharts = this.$refs.medalChart as ECharts;
            medalChart.setOption(soptions, {
                replaceMerge: ['series']
            })
        },
        onClickRich(param: any) {
            console.log(param.name);
            this.sendGiftsChartOptions.title = {
                text: `${param.name}送礼详情`,
                left: 'center',
            };
            this.sendGiftsChartOptions.tooltip = {
                trigger: 'item',
                formatter: '{b} : {c} ({d}%)'
            };
            let gifts: string[] = [];
            let data: any[] = [];
            for (let key of this.receivedJson['revenue_statistics']['rich_ranking']) {
                if (key['uname'] === param.name) {
                    for (let key2 of key['send_gifts']) {
                        if (gifts.indexOf(key2['gift_name']) === -1) {
                            gifts.push(key2['gift_name']);
                            data.push({ name: key2['gift_name'], value: key2['price'] / 1000 });
                        } else {
                            for (let key3 in data) {
                                if (data[key3].name === key2['gift_name']) {
                                    data[key3].value += key2['price'] / 1000;
                                }
                            }
                        }
                    }
                }
            }
            this.sendGiftsChartOptions.legend = {
                orient: 'vertical',
                left: 'left',
                data: gifts
            }
            this.sendGiftsChartOptions.series = {
                type: 'pie',
                radius: '50%',
                center: ['40%', '50%'],
                selectedMode: 'single',
                itemStyle: {
                    borderColor: '#fff',
                    borderWidth: 1
                },
                data: data,
                emphasis: {
                    itemStyle: {
                        shadowBlur: 10,
                        shadowOffsetX: 0,
                        shadowColor: 'rgba(0, 0, 0, 0.5)'
                    }
                }
            }
            this.showSengGiftsChart = true;
        },
        onClickPrev() {
            this.fetchStatistics(this.brgLst[this.menuIndex - 1]);
        },
        onClickNext() {
            this.fetchStatistics(this.brgLst[this.menuIndex + 1]);
        }
    },
})
</script>

<template>
  <div class="bg">
    <dv-loading v-if="loading">Loading...</dv-loading>
    <div style="margin-top: 2%">

      <el-row :gutter="20">
        <el-col :span="6"><div class="grid-content bg-purple">
          <dv-border-box-1>
            <dv-digital-flop :config="config_digital_flop" style="width:100px;height:50px; margin-left: 20%" />
            <dv-digital-flop :config="config_digital_flop" style="width:100px;height:50px; margin-left: 20%" />
            <dv-digital-flop :config="config_digital_flop" style="width:100px;height:50px; margin-left: 20%" />
            <dv-digital-flop :config="config_digital_flop" style="width:100px;height:50px; margin-left: 20%" />
            <dv-digital-flop :config="config_digital_flop" style="width:100px;height:50px; margin-left: 20%" />
            <dv-digital-flop :config="config_digital_flop" style="width:100px;height:50px; margin-left: 20%" />
          </dv-border-box-1>
        </div>
        </el-col>
        <el-col :span="6"><div class="grid-content bg-purple">
          <dv-border-box-8>
            <dv-water-level-pond :config="config_water_level" style="width:150px;height:200px" />
          </dv-border-box-8>
        </div></el-col>
        <el-col :span="6"><div class="grid-content bg-purple">
          <dv-border-box-9>
            <dv-active-ring-chart :config="config" style="width:300px;height:300px" />
          </dv-border-box-9>
        </div></el-col>
        <el-col :span="6"><div class="grid-content bg-purple">
          <dv-border-box-1>
            <dv-percent-pond :config="config" style="width:200px;height:100px;" />
          </dv-border-box-1>
        </div></el-col>
      </el-row>
      <el-row :gutter="20">
        <el-col :span="4"><div class="grid-content bg-purple">
          <dv-border-box-1>
            <dv-conical-column-chart :config="config" style="width:400px;height:200px;" />
            <dv-conical-column-chart :config="config" style="width:400px;height:200px;" />
            <dv-conical-column-chart :config="config" style="width:400px;height:200px;" />
          </dv-border-box-1>
        </div></el-col>
        <el-col :span="12" style="margin-left: 2%"><div class="grid-content bg-purple" style="margin-top: 1% ">
          <dv-flyline-chart :config="config_flyline" style="width: 100%;height:600px" />
        </div></el-col>
        <el-col :span="4"><div class="grid-content bg-purple">
          <dv-border-box-1>
            <dv-conical-column-chart :config="config" style="width:400px;height:200px;" />
            <dv-conical-column-chart :config="config" style="width:400px;height:200px;" />
            <dv-conical-column-chart :config="config" style="width:400px;height:200px;" />
          </dv-border-box-1>
        </div></el-col>
      </el-row>
    </div>

    <el-row :gutter="20">
      <el-col :span="6"><div class="grid-content bg-purple">
        <dv-border-box-1>
          <dv-active-ring-chart :config="config" style="width:300px;height:300px" />
        </dv-border-box-1>
      </div>
      </el-col>
      <el-col :span="6"><div class="grid-content bg-purple"><dv-border-box-8>   <dv-active-ring-chart :config="config" style="width:300px;height:300px" /></dv-border-box-8></div></el-col>
      <el-col :span="6"><div class="grid-content bg-purple"><dv-border-box-9>   <dv-active-ring-chart :config="config" style="width:300px;height:300px" /></dv-border-box-9></div></el-col>
      <el-col :span="6"><div class="grid-content bg-purple"><dv-border-box-11 title="dv-border-box-11">   <dv-active-ring-chart :config="config" style="width:300px;height:300px" /></dv-border-box-11></div></el-col>
    </el-row>
    <el-row :gutter="20">
      <el-col :span="6"><div class="grid-content bg-purple">
        <dv-border-box-1>
          <dv-capsule-chart :config="config_capsule" style="width:300px;height:200px" />
        </dv-border-box-1>
      </div>
      </el-col>
      <el-col :span="6"><div class="grid-content bg-purple"><dv-border-box-8>   <dv-active-ring-chart :config="config" style="width:300px;height:300px" /></dv-border-box-8></div></el-col>
      <el-col :span="6"><div class="grid-content bg-purple"><dv-border-box-9>   <dv-active-ring-chart :config="config" style="width:300px;height:300px" /></dv-border-box-9></div></el-col>
      <el-col :span="6"><div class="grid-content bg-purple">
        <dv-border-box-9>
          <dv-charts :option="option" style="width:300px;height:150px" />
          <dv-charts :option="option" style="width:300px;height:150px" />
        </dv-border-box-9>
      </div></el-col>

    </el-row>
    <el-row :gutter="20">
      <el-col :span="6"><div class="grid-content bg-purple">
        <dv-border-box-8>
          <dv-scroll-ranking-board :config="config_ranking_board" style="width:450px;height:300px ;margin-left: 2%" />
        </dv-border-box-8>
      </div>
      </el-col>
      <el-col :span="6"><div class="grid-content bg-purple">
        <dv-border-box-8>
          <dv-scroll-ranking-board :config="config_ranking_board" style="width:450px;height:300px ;margin-left: 2%" />
        </dv-border-box-8>
      </div></el-col>
      <el-col :span="6"><div class="grid-content bg-purple">
        <dv-border-box-8>
          <dv-scroll-ranking-board :config="config_ranking_board" style="width:450px;height:300px ;margin-left: 2%" />
        </dv-border-box-8>
      </div></el-col>
      <el-col :span="6"><div class="grid-content bg-purple">
        <dv-border-box-8>
          <dv-scroll-ranking-board :config="config_ranking_board" style="width:450px;height:300px ;margin-left: 2%" />
        </dv-border-box-8>
      </div></el-col>
    </el-row>

  </div>
</template>

<script>
import TabPane from './components/TabPane'
import PanelGroup from '../components/PanelGroup'
export default {
  name: 'Tab',
  components: { TabPane, PanelGroup },
  data() {
    return {
      loading: true,
      tabMapOptions: [
        { label: '主机(HOST)', key: 'Host' },
        { label: 'CPU', key: 'Cpu' },
        { label: '内存(MEM)', key: 'Mem' },
        { label: '硬盘(DISK)', key: 'Disk' },
        { label: '网络(NET)', key: 'Net' },
        { label: '容器(DOCKER)', key: 'Docker' },
        { label: '进程(PROCESS)', key: 'Process' },
        { label: '负载(LOAD)', key: 'Load' }
      ],
      activeName: 'Host',
      createdTimes: 0,
      config: {
        data: [
          {
            name: '周口',
            value: 55
          },
          {
            name: '南阳',
            value: 120
          },
          {
            name: '西峡',
            value: 78
          },
          {
            name: '驻马店',
            value: 66
          },
          {
            name: '新乡',
            value: 80
          }
        ]
      },
      config_digital_flop: {
        number: [10],
        toFixed: 2,
        content: '{nt}个'
      },
      option: {
        title: {
          text: '剩余油量表'
        },
        series: [
          {
            type: 'gauge',
            data: [
              { name: '油箱A', value: 59 },
              { name: '油箱B', value: 78, radius: '40%' },
              { name: '油箱C', value: 45, radius: '20%' }
            ],
            axisLabel: {
              formatter: '{value}%'
            },
            pointer: {
              valueIndex: 2,
              style: {
                scale: [0.6, 0.6]
              }
            },
            details: {
              show: true,
              formatter: '{name}剩余{value}%',
              position: 'start',
              offset: [10, 0],
              style: {
                fontSize: 13,
                textAlign: 'left'
              }
            }
          }
        ]
      },
      config_capsule: {
        data: [
          {
            name: '南阳',
            value: 167
          },
          {
            name: '周口',
            value: 123
          },
          {
            name: '漯河',
            value: 98
          },
          {
            name: '郑州',
            value: 75
          },
          {
            name: '西峡',
            value: 66
          }
        ],
        colors: ['#e062ae', '#fb7293', '#e690d1', '#32c5e9', '#96bfff'],
        unit: '单位',
        showValue: true
      },
      config_ranking_board: {
        data: [
          {
            name: '周口',
            value: 55123
          },
          {
            name: '南阳',
            value: 12022
          },
          {
            name: '西峡',
            value: 78932
          },
          {
            name: '驻马店',
            value: 63411
          },
          {
            name: '新乡',
            value: 44231
          }
        ],
        unit: '单位',
        valueFormatter({ value }) {
          console.warn(arguments)
          const reverseNumber = (value + '').split('').reverse()
          let valueStr = ''

          while (reverseNumber.length) {
            const seg = reverseNumber.splice(0, 3).join('')
            valueStr += seg
            if (seg.length === 3) valueStr += ','
          }

          return valueStr.split('').reverse().join('')
        }
      },
      config_water_level: {
        data: [66, 45],
        shape: 'roundRect'
      },
      config_flyline: {
        centerPoint: [0.48, 0.35],
        points: [
          {
            position: [0.52, 0.235],
            text: '新乡'
          },
          {
            position: [0.43, 0.29],
            text: '焦作'
          },
          {
            position: [0.59, 0.35],
            text: '开封'
          },
          {
            position: [0.53, 0.47],
            text: '许昌'
          },
          {
            position: [0.45, 0.54],
            text: '平顶山'
          },
          {
            position: [0.36, 0.38],
            text: '洛阳'
          },
          {
            position: [0.62, 0.55],
            text: '周口'
          },
          {
            position: [0.56, 0.56],
            text: '漯河'
          },
          {
            position: [0.37, 0.66],
            text: '南阳'
          },
          {
            position: [0.55, 0.81],
            text: '信阳'
          },
          {
            position: [0.55, 0.67],
            text: '驻马店'
          },
          {
            position: [0.37, 0.29],
            text: '济源'
          },
          {
            position: [0.20, 0.36],
            text: '三门峡'
          },
          {
            position: [0.76, 0.41],
            text: '商丘'
          },
          {
            position: [0.59, 0.18],
            text: '鹤壁'
          },
          {
            position: [0.68, 0.17],
            text: '濮阳'
          },
          {
            position: [0.59, 0.10],
            text: '安阳'
          }
        ],
        k: 0.5,
        bgImgUrl: 'http://datav.jiaminghi.com/img/flylineChart/map.jpg',
        centerPointImg: {
          url: 'http://datav.jiaminghi.com/img/flylineChart/mapCenterPoint.png'
        },
        pointsImg: {
          url: 'http://datav.jiaminghi.com/img/flylineChart/mapPoint.png'
        }
      }
    }
  },
  watch: {
    activeName(val) {
      this.$router.push(`${this.$route.path}?tab=${val}`)
    }
  },
  mounted() {
    this.cancelLoading()
  },
  created() {
    // init the default selected tab
    const tab = this.$route.query.tab
    if (tab) {
      this.activeName = tab
    }
  },
  methods: {
    showCreatedTimes() {
      this.createdTimes = this.createdTimes + 1
    },
    cancelLoading() {
      setTimeout(() => {
        this.loading = false
      }, 500)
    }
  }
}
</script>

<style lang="scss">

  @import "../../../assets/fullscreen/scss/index.scss";
</style>
<style>
  .bg{
    background-image: url("../../../assets/images/fullscreen/star-background.jpg");
  }

</style>

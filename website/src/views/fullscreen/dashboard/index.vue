<template>
  <div class="bg"  >
    <dv-loading v-if="loading">Loading...</dv-loading>
    <el-row :gutter="20">
      <el-col :span="6"><div class="grid-content bg-purple">
        <dv-border-box-1>
          <dv-active-ring-chart :config="config" style="width:300px;height:300px" />
        </dv-border-box-1>
      </div>
      </el-col>
      <el-col :span="6"><div class="grid-content bg-purple"><dv-border-box-8>dv-border-box-8</dv-border-box-8></div></el-col>
      <el-col :span="6"><div class="grid-content bg-purple"><dv-border-box-9>dv-border-box-9</dv-border-box-9></div></el-col>
      <el-col :span="6"><div class="grid-content bg-purple"><dv-border-box-11 title="dv-border-box-11">dv-border-box-11</dv-border-box-11></div></el-col>
    </el-row>


  </div>
</template>

<script>
import TabPane from './components/TabPane'

export default {
  name: 'Tab',
  components: { TabPane },
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
      config:{
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
      }
    }
  },
  watch: {
    activeName(val) {
      this.$router.push(`${this.$route.path}?tab=${val}`)
    }
  },
  mounted() {
    this.cancelLoading();
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
        this.loading = false;
      }, 500);
    }
  }
}
</script>

<style lang="scss">

  @import "../../../assets/fullscreen/scss/index.scss";
</style>
<style>
  .bg{
    background-image: url("../../../assets/images/fullscreen/pageBg.png");
  }

</style>

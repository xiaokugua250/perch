<template>
  <div class="tab-container">
<!--    <el-tag>mounted times ：{{ createdTimes }}</el-tag>-->
    <el-alert :closable="false" style="width:200px;display:inline-block;vertical-align: middle;margin-left:30px;" title="Tab with keep-alive" type="success" />
    <el-tabs v-model="activeName" style="margin-top:15px;" type="border-card" @tab-click="handleClick">
      <el-tab-pane label="CPU" name="cpu" >
      <CPUTablePane></CPUTablePane>

    </el-tab-pane>
      <el-tab-pane label="内存管理" name="memory" >
        <memory-table-pane></memory-table-pane>

      </el-tab-pane>

      <el-tab-pane label="HOST" name="host" >
        <HostTablePane></HostTablePane>
      </el-tab-pane>
      <el-tab-pane label="DISK" name="disk" >
        <DiskTablePane></DiskTablePane>
      </el-tab-pane>
      <el-tab-pane label="NETWORK" name="network" >
        <NetworkTablePane></NetworkTablePane>
      </el-tab-pane>
      <el-tab-pane label="DOCKER" name="docker" >
        <DockerTablePane></DockerTablePane>

      </el-tab-pane>
      <el-tab-pane label="LOAD" name="load" >
        <LoadTablePane></LoadTablePane>
      </el-tab-pane>
      <el-tab-pane label="PROECESS" name="process" >
        <ProcessTablePane></ProcessTablePane>

      </el-tab-pane>
  <!--    <el-tab-pane v-for="item in tabMapOptions" :key="item.key" :label="item.label" :name="item.key">
        <keep-alive>
          <tab-pane v-if="activeName==item.key" :type="item.key" @create="showCreatedTimes" />
        </keep-alive>
      </el-tab-pane>-->
    </el-tabs>
  </div>
</template>

<script>
import TabPane from './components/TabPane'
import CPUTablePane from './components/CPUTabPane'
import HostTablePane from './components/HostTabPane'
import MemoryTablePane from './components/MemoryTabPane'
import DiskTablePane from './components/DiskTabPane'
import NetworkTablePane from './components/NetworkTabPane'
import DockerTablePane from './components/DockerTabPane'
import ProcessTablePane from './components/ProcessTabPane'
import LoadTablePane from './components/LoadTabPane'

export default {
  name: 'Tab',
  components: { TabPane ,CPUTablePane,MemoryTablePane,HostTablePane,DiskTablePane,NetworkTablePane,DockerTablePane,ProcessTablePane,LoadTablePane},
  data() {
    return {
      roles:"",
      permissions:"",
      tabMapOptions: [
        { label: '用户角色', key: 'roles' },
        { label: '平台权限', key: 'permissions' },
      ],
      activeName: 'cpu',
      createdTimes: 0
    }
  },
  watch: {
    activeName(val) {
      this.$router.push(`${this.$route.path}?tab=${val}`)
    }
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
    handleClick(tab, event) {
      console.log(tab, event);
    }
  }
}
</script>

<style scoped>
  .tab-container {
    margin: 30px;
  }
</style>

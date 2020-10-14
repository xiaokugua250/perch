<template>
  <div class="tab-container">
    <el-tag>mounted times ：{{ createdTimes }}</el-tag>
    <el-alert :closable="false" style="width:200px;display:inline-block;vertical-align: middle;margin-left:30px;" title="Tab with keep-alive" type="success" />
    <el-tabs v-model="activeName" style="margin-top:15px;" type="border-card">

      <el-tab-pane label="Node" name="node">
        <NodeTabPane/>
      </el-tab-pane>
      <el-tab-pane label="Namespaces" name="namespaces">
      <NamespaceTabPane></NamespaceTabPane>

      </el-tab-pane>

      <el-tab-pane label="Service" name="service">

        <ServiceTabPane></ServiceTabPane>
      </el-tab-pane>
      <el-tab-pane label="ConfigMap" name="configmap">
        <ConfigMapTabPane></ConfigMapTabPane>
      </el-tab-pane>
      <el-tab-pane label="ServiceAccount" name="serviceaccount">
        <ServiceAccountTabPane></ServiceAccountTabPane>
      </el-tab-pane>
      <el-tab-pane label="Pod" name="pod">
        <PodTabPane></PodTabPane>
      </el-tab-pane>
      <el-tab-pane label="Job" name="job">
        <JobTabPane></JobTabPane>
      </el-tab-pane>
      <el-tab-pane label="BatchJob" name="batchjob">
        <BatchJobTabPane></BatchJobTabPane>
      </el-tab-pane>
      <el-tab-pane label="Deployment" name="deployment">
      <DeploymentTabPane></DeploymentTabPane>
    </el-tab-pane>
      <el-tab-pane label="Daemonset" name="daemonset">
      <DaemonsetTabPane></DaemonsetTabPane>
    </el-tab-pane>
      <el-tab-pane label="Replicaset" name="replicaset">
      <ReplicasetTabPane></ReplicasetTabPane>
    </el-tab-pane>
      <el-tab-pane label="Statefulset" name="statefulset">
      <StatefuleSetTabPane></StatefuleSetTabPane>
    </el-tab-pane>
      <el-tab-pane label="PV" name="pv">
     <PVTabPane/>
    </el-tab-pane>
      <el-tab-pane label="PVC" name="pvc">
     <PVCTabPane/>
    </el-tab-pane>







   <!--   <el-tab-pane v-for="item in tabMapOptions" :key="item.key" :label="item.label" :name="item.key">
        <keep-alive>
          <tab-pane v-if="activeName==item.key" :type="item.key" @create="showCreatedTimes" />
        </keep-alive>
      </el-tab-pane>-->
    </el-tabs>
  </div>
</template>

<script>
import TabPane from './components/TabPane'
import node from './components/NodeTabPane'
import NodeTabPane from "./components/NodeTabPane";
import NamespaceTabPane from "./components/NamespaceTabPane";
import ServiceTabPane from "./components/serviceTabPane";
import ConfigMapTabPane from "./components/configmapTabPane";
import ServiceAccountTabPane from "./components/serviceaccountTabPane";
import PodTabPane from "./components/podTabPane";
import JobTabPane from "./components/jobTabPane";
import BatchJobTabPane from "./components/batchjobTabPane";
import DeploymentTabPane from "./components/deploymentTabPane";
import DaemonsetTabPane from "./components/daemonsetTabPane";
import ReplicasetTabPane from "./components/replicasetTabPane";
import StatefuleSetTabPane from "./components/statefulsetTabPane";
import PVTabPane from "./components/pvTabPane";
import PVCTabPane from "./components/pvcTabPane";


export default {
  name: 'Tab',
  components: {NodeTabPane,NamespaceTabPane, ServiceTabPane,
    ConfigMapTabPane,ServiceAccountTabPane,PodTabPane,JobTabPane,BatchJobTabPane,DeploymentTabPane,
    DaemonsetTabPane,ReplicasetTabPane,StatefuleSetTabPane,PVTabPane,PVCTabPane,
    TabPane },
  data() {
    return {
      tabMapOptions: [
        { label: '节点(Node)', key: 'Node' },
        { label: 'NameSpaces', key: 'Namespaces' },
        { label: 'DeployMent', key: 'DeployMent' },
        { label: 'StatefulSet', key: 'StatefulSet' },
        { label: 'PODS', key: 'PODS' },
        { label: 'Service)', key: 'Service' },
        { label: 'Resources', key: 'Resources' },
        { label: 'Jobs', key: 'Jobs' },
        { label: 'CronJobs', key: 'CronJobs' },
        { label: 'CRDS', key: 'CRDS' }
      ],
      activeName: 'Node',
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
    }
  }
}
</script>

<style scoped>
  .tab-container {
    margin: 30px;
  }
</style>

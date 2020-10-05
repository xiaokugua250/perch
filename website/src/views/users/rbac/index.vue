<template>
  <div class="tab-container">
<!--    <el-tag>mounted times ：{{ createdTimes }}</el-tag>-->
    <el-alert :closable="false" style="width:200px;display:inline-block;vertical-align: middle;margin-left:30px;" title="Tab with keep-alive" type="success" />
    <el-tabs v-model="activeName" style="margin-top:15px;" type="border-card" @tab-click="handleClick">
      <el-tab-pane label="角色管理" name="roles" >
      <RoleTabPane/>

      </el-tab-pane>
      <el-tab-pane label="权限管理" name="permissions" >
        <PermissionTabPane/>

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
import RoleTabPane from  './components/RoleTabPane'
import PermissionTabPane from  './components/PermissionTabPane'

export default {
  name: 'Tab',
  components: { TabPane ,RoleTabPane,PermissionTabPane},
  data() {
    return {
      roles:"",
      permissions:"",
      tabMapOptions: [
        { label: '用户角色', key: 'roles' },
        { label: '平台权限', key: 'permissions' },
      ],
      activeName: 'roles',
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

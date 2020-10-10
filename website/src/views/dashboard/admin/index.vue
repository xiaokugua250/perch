<template>
  <div class="dashboard-editor-container">
    <github-corner class="github-corner" />

    <panel-group @handleSetLineChartData="handleSetLineChartData" />
    <el-row :gutter="20">
      <el-col v-for="(image, index) in 8" :key="index" :md="9" :lg="7" :xl="6">
        <!--   <el-col v-for="(image, index) in dchub_images" :key="index" :md="9" :lg="7" :xl="6">-->
        <el-card
          :body-style="{ padding: '0px' }"
          class="dchub_image_card"
          style="width: 350px"
        >
          <div style="cursor: pointer;display: inline-flex;" class="img" @click="DchubSpecImageGet(image.id)">
            <div style="width: 50%;">
              <img v-if="image.icon_uuid" width="100" height="100" style="display: block;margin: auto" :src="getImg(image.icon_uuid)">
              <img v-else width="100" height="100" style="display: block;margin: auto" src="../../../assets/images/dashboard/programming.png">
            </div>
            <div style="width: 50%;">
              <el-tag class="card_left">标题:&nbsp;&nbsp;&nbsp;<!--{{ 11 }}--></el-tag>
              <el-tag style="height: 30px;width: 180px;display: block;margin: auto;margin-bottom: 10px;margin-top: 2%;">领域:&nbsp;&nbsp;&nbsp;<span v-if="image.is_public ===true">公开镜像</span><span v-else>私有镜像</span></el-tag>
              <el-tag style="height: 30px;width: 180px;display: block;margin: auto;margin-bottom: 10px;margin-top: 2%;">日期:&nbsp;&nbsp;&nbsp;<span v-if="image.is_public ===true">公开镜像</span><span v-else>私有镜像</span></el-tag>
              <el-tag style="height: 30px;width: 180px;display: block;margin: auto;margin-bottom: 10px;margin-top: 2%;">详情链接:&nbsp;&nbsp;&nbsp;<!--{{ calculate(image.size) }}--></el-tag>
            </div>
          </div>
        </el-card>
      </el-col>
      <!--    <el-col :span="6">
      <div class="chart-wrapper">

        <div slot="header" class="clearfix">
          <span>平台服务(PLATADMIN)</span>
          <el-button style="float: right; padding: 3px 0" type="text">操作按钮</el-button>
        </div>
        <div v-for="o in 4" :key="o" class="text item">
          {{'列表内容 ' + o }}
        </div>

      </div>
    </el-col>
    <el-col :span="6">
      <div class="chart-wrapper">

        <div slot="header" class="clearfix">
          <span>数据库服务(MySQL)</span>
          <el-button style="float: right; padding: 3px 0" type="text">操作按钮</el-button>
        </div>
        <div v-for="o in 4" :key="o" class="text item">
          {{'列表内容 ' + o }}
        </div>

      </div>
    </el-col>
    <el-col :span="6">
      <div class="chart-wrapper">

        <div slot="header" class="clearfix">
          <span>系统监控(LINUX)</span>
          <el-button style="float: right; padding: 3px 0" type="text">操作按钮</el-button>
        </div>
        <div v-for="o in 4" :key="o" class="text item">
          {{'列表内容 ' + o }}
        </div>

      </div>
    </el-col>
    <el-col :span="6">       <div class="chart-wrapper">

      <div slot="header" class="clearfix">
        <span>爬虫服务(Colly)</span>
        <el-button style="float: right; padding: 3px 0" type="text">操作按钮</el-button>
      </div>
      <div v-for="o in 4" :key="o" class="text item">
        {{'列表内容 ' + o }}
      </div>

      &lt;!&ndash;   <el-card class="box-card">
           <div slot="header" class="clearfix">
             <span>卡片名称</span>
             <el-button style="float: right; padding: 3px 0" type="text">操作按钮</el-button>
           </div>
           <div v-for="o in 4" :key="o" class="text item">
             {{'列表内容 ' + o }}
           </div>
         </el-card>&ndash;&gt;
    </div></el-col>-->

    </el-row>
    <!--   <el-row :gutter="20">
      <el-col :span="6">
        <div class="chart-wrapper">

          <div slot="header" class="clearfix">
            <span>复杂网络(NETWORK)</span>
            <el-button style="float: right; padding: 3px 0" type="text">操作按钮</el-button>
          </div>
          <div v-for="o in 4" :key="o" class="text item">
            {{'列表内容 ' + o }}
          </div>

        </div>
      </el-col>
      <el-col :span="6">
        <div class="chart-wrapper">

          <div slot="header" class="clearfix">
            <span>容器与虚拟化(Vituralization)</span>
            <el-button style="float: right; padding: 3px 0" type="text">操作按钮</el-button>
          </div>
          <div v-for="o in 4" :key="o" class="text item">
            {{'列表内容 ' + o }}
          </div>

        </div>
      </el-col>
      <el-col :span="6">
        <div class="chart-wrapper">

          <div slot="header" class="clearfix">
            <span>调度系统(K8S)</span>
            <el-button style="float: right; padding: 3px 0" type="text">操作按钮</el-button>
          </div>
          <div v-for="o in 4" :key="o" class="text item">
            {{'列表内容 ' + o }}
          </div>

        </div>
      </el-col>
      <el-col :span="6">       <div class="chart-wrapper">

        <div slot="header" class="clearfix">
          <span>其他(OTHERS)</span>
          <el-button style="float: right; padding: 3px 0" type="text">操作按钮</el-button>
        </div>
        <div v-for="o in 4" :key="o" class="text item">
          {{'列表内容 ' + o }}
        </div>

        &lt;!&ndash;   <el-card class="box-card">
             <div slot="header" class="clearfix">
               <span>卡片名称</span>
               <el-button style="float: right; padding: 3px 0" type="text">操作按钮</el-button>
             </div>
             <div v-for="o in 4" :key="o" class="text item">
               {{'列表内容 ' + o }}
             </div>
           </el-card>&ndash;&gt;
      </div></el-col>

    </el-row>-->
    <el-row style="background:#fff;padding:16px 16px 0;margin-bottom:32px;margin-top: 1%">
      <line-chart :chart-data="lineChartData" />
    </el-row>

    <el-row :gutter="32">
      <el-col :xs="24" :sm="24" :lg="8">
        <div class="chart-wrapper">
          <raddar-chart />
        </div>
      </el-col>
      <el-col :xs="24" :sm="24" :lg="8">
        <div class="chart-wrapper">
          <pie-chart />
        </div>
      </el-col>
      <el-col :xs="24" :sm="24" :lg="8">
        <div class="chart-wrapper">
          <bar-chart />
        </div>
      </el-col>
    </el-row>

    <el-row :gutter="8">
      <el-col :xs="{span: 24}" :sm="{span: 24}" :md="{span: 24}" :lg="{span: 12}" :xl="{span: 12}" style="padding-right:8px;margin-bottom:30px;">
        <transaction-table />
      </el-col>
      <el-col :xs="{span: 24}" :sm="{span: 12}" :md="{span: 12}" :lg="{span: 6}" :xl="{span: 6}" style="margin-bottom:30px;">
        <todo-list />
      </el-col>
      <el-col :xs="{span: 24}" :sm="{span: 12}" :md="{span: 12}" :lg="{span: 6}" :xl="{span: 6}" style="margin-bottom:30px;">
        <box-card />
      </el-col>
    </el-row>
  </div>
</template>

<script>
import GithubCorner from '@/components/GithubCorner'
import PanelGroup from './components/PanelGroup'
import LineChart from './components/LineChart'
import RaddarChart from './components/RaddarChart'
import PieChart from './components/PieChart'
import BarChart from './components/BarChart'
import TransactionTable from './components/TransactionTable'
import TodoList from './components/TodoList'
import BoxCard from './components/BoxCard'

import { getResourceDocs } from '@/api/resources-docs'

const lineChartData = {
  newVisitis: {
    expectedData: [100, 120, 161, 134, 105, 160, 165],
    actualData: [120, 82, 91, 154, 162, 140, 145]
  },
  messages: {
    expectedData: [200, 192, 120, 144, 160, 130, 140],
    actualData: [180, 160, 151, 106, 145, 150, 130]
  },
  purchases: {
    expectedData: [80, 100, 121, 104, 105, 90, 100],
    actualData: [120, 90, 100, 138, 142, 130, 130]
  },
  shoppings: {
    expectedData: [130, 140, 141, 142, 145, 150, 160],
    actualData: [120, 82, 91, 154, 162, 140, 130]
  }
}

export default {
  name: 'DashboardAdmin',
  components: {
    GithubCorner,
    PanelGroup,
    LineChart,
    RaddarChart,
    PieChart,
    BarChart,
    TransactionTable,
    TodoList,
    BoxCard
  },
  data() {
    return {
      lineChartData: lineChartData.newVisitis,
      resourceArticles: []
    }
  },
  mounted() {
    // this.getList()
    this.resourceArticlesGet()
  },
  methods: {
    handleSetLineChartData(type) {
      this.lineChartData = lineChartData[type]
    },
    resourceArticlesGet() {
      this.listLoading = true
      getResourceDocs(this.listQuery).then(response => {
        //        this.list = response.data.items
        this.resourceArticles = response.spec

        this.total = response.total

        // Just to simulate the time of the request
        setTimeout(() => {
          this.listLoading = false
        }, 1.5 * 1000)
      })
    }
  }
}
</script>

<style lang="scss" scoped>
.dashboard-editor-container {
  padding: 32px;
  background-color: rgb(240, 242, 245);
  position: relative;

  .github-corner {
    position: absolute;
    top: 0px;
    border: 0;
    right: 0;
  }

  .chart-wrapper {
    background: #fff;
    padding: 16px 16px 0;
    margin-bottom: 32px;
  }
}
.card_left{
  height: 30px;width: 180px;display: block;margin: auto;margin-bottom: 10px;margin-top: 5%;
}

.dchub_image_card {
  display: -webkit-box;
  display: -ms-flexbox;
  display: flex;
  /* -webkit-box-align: center; */
  -ms-flex-align: center;
  align-items: center;
  height: 170px;
  margin-top: 10px;
  .img {
    transition: all 0.2s linear;
    &:hover {
      transform: scale(1.1, 1.1);
    }
  }
}

.dchub-image {
  max-width: 100px;
  max-height: 100px;
  margin-top: 30%;
  display: block;
  margin-left: 10px;
  margin-right: auto;

}

@media (max-width:1024px) {
  .chart-wrapper {
    padding: 8px;
  }

}
</style>

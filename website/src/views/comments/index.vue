<template>
  <div class="app-container">
    <aside>
      The guide page is useful for some people who entered the project for the first time. You can briefly introduce the
      features of the project. Demo is based on
      <a href="https://github.com/kamranahmedse/driver.js" target="_blank">driver.js.</a>
    </aside>
    <el-button icon="el-icon-question" type="primary" @click.prevent.stop="guide">
      Show Guide
    </el-button>
    <div style="margin-top: 1%">
      <el-row>
        <el-col :span="24">
          <div class="grid-content ">
            <div style="margin-top: 15px; ">
              <el-input v-model="search_model" class="input-with-select" placeholder="请输入搜索条件" @keyup.enter.native="Search()">
                <el-select slot="prepend" v-model="select_options" placeholder="请选择" style="width: 200px">
                  <el-option label="文本搜索" value="name" />
                  <el-option label="图片搜索" value="app_name" />
                </el-select>
                <el-button slot="append" icon="el-icon-search" @click="Search()" />
              </el-input>
            </div>

            <div style="margin-top: 1%" />
            <div>
              <el-row :gutter="20">
                <el-col v-for="(image, index) in images" :key="index" :md="9" :lg="7" :xl="6">
                  <el-card
                    :body-style="{ padding: '0px' }"
                    class="dchub_image_card"
                    style="width: 350px"
                  >
                    <div style="cursor: pointer;display: inline-flex;" class="img" @click="SpecImageGet(image.id)">
                      <div style="width: 50%;">
                        <img v-if="image.icon_uuid" width="100" height="100" style="display: block;margin: auto" :src="getImg(image.icon_uuid)">

                      </div>
                      <div style="width: 50%;">
                        <el-tag style="height: 30px;width: 180px;display: block;margin: auto;margin-bottom: 10px;margin-top: 5%;">名称:&nbsp;&nbsp;&nbsp;{{ getName(image.name) }}</el-tag>
                        <el-tag style="height: 30px;width: 180px;display: block;margin: auto;margin-bottom: 10px;margin-top: 5%;">是否公开:&nbsp;&nbsp;&nbsp;<span v-if="image.is_public ===true">公开镜像</span><span v-else>私有镜像</span></el-tag>
                        <el-tag style="height: 30px;width: 180px;display: block;margin: auto;margin-bottom: 10px;margin-top: 5%;">名称:&nbsp;&nbsp;&nbsp;{{ calculate(image.size) }}</el-tag>
                      </div>
                    </div>
                  </el-card>
                </el-col>

              </el-row>
              <br>
              <el-row :gutter="20" style="text-align:center">
                <el-paginationl
                  v-if="total >=12"
                  :current-page="listQuery.page"
                  :page-size="listQuery.limit"
                  :page-sizes="[12,20,30, 50]"
                  :total="total"
                  background
                  layout="total, sizes, prev, pager, next, jumper"
                  @current-change="handleCurrentChange"
                  @size-change="handleSizeChange"
                />
              </el-row>
            </div>
          </div>
        </el-col>
      </el-row>
      <!--      <el-row>

        <el-col :span="24" class="center">
          <el-select v-model="value" placeholder="请选择">
            <el-option
              v-for="item in options"
              :key="item.value"
              :label="item.label"
              :value="item.value">
            </el-option>
          </el-select>
          <el-input
            v-model="search"
            @focus="focus"
            @blur="blur"
            @keyup.enter.native="searchHandler"
            placeholder="搜索商家或地点"
          >
            <el-button slot="append" icon="el-icon-search" id="search" @click="searchHandler"></el-button>
          </el-input>
          &lt;!&ndash;-设置z-index优先级,防止被走马灯效果遮挡&ndash;&gt;
          <el-card
            @mouseenter="enterSearchBoxHanlder"
            v-if="isSearch"
            class="box-card"
            id="search-box"
            style="position:relative;z-index:15"
          >
            <dl v-if="isHistorySearch">
              <dt class="search-title" v-show="history">历史搜索</dt>
              <dt class="remove-history" v-show="history" @click="removeAllHistory">
                <i class="el-icon-delete"></i>清空历史记录
              </dt>
              <el-tag
                v-for="search in historySearchList"
                :key="search.id"
                closable
                :type="search.type"
                @close="closeHandler(search)"
                style="margin-right:5px; margin-bottom:5px;"
              >{{search.name}}</el-tag>
              <dt class="search-title">热门搜索</dt>
              <dd v-for="search in hotSearchList" :key="search.id">{{search}}</dd>
            </dl>
            <dl v-if="isSearchList">
              <dd v-for="search in searchList" :key="search.id">{{search}}</dd>
            </dl>
          </el-card>
        </el-col>
      </el-row>-->

    </div>

  </div>
</template>

<script>
import Driver from 'driver.js' // import driver.js
import 'driver.js/dist/driver.min.css' // import driver.js css
import steps from './steps'

export default {
  name: 'Guide',
  data() {
    return {
      driver: null,
      form: {
        name: null
      }
    }
  },
  computed: {

    isHistorySearch() {
      return this.isFocus && !this.search
    },
    isSearchList() {
      return this.isFocus && this.search
    },
    isSearch() {
      return this.isFocus
    }
  },
  mounted() {
    this.driver = new Driver()
  },
  methods: {
    guide() {
      this.driver.defineSteps(steps)
      this.driver.start()
    },
    focus() {
      this.isFocus = true
      this.historySearchList =
        Store.loadHistory() == null ? [] : Store.loadHistory()
      this.history = this.historySearchList.length != 0
    },
    blur() {
      var self = this
      this.searchBoxTimeout = setTimeout(function() {
        self.isFocus = false
      }, 300)
    },
    enterSearchBoxHanlder() {
      clearTimeout(this.searchBoxTimeout)
    },
    searchHandler() {
      // 随机生成搜索历史tag式样
      const n = RandomUtil.getRandomNumber(0, 5)
      const exist =
        this.historySearchList.filter(value => {
          return value.name == this.search
        }).length != 0
      if (!exist) {
        this.historySearchList.push({ name: this.search, type: this.types[n] })
        Store.saveHistory(this.historySearchList)
      }
      this.history = this.historySearchList.length != 0
    },
    closeHandler(search) {
      this.historySearchList.splice(this.historySearchList.indexOf(search), 1)
      Store.saveHistory(this.historySearchList)
      clearTimeout(this.searchBoxTimeout)
      if (this.historySearchList.length == 0) {
        this.history = false
      }
    },
    removeAllHistory() {
      Store.removeAllHistory()
    }
  }
}
</script>

<style>

</style>

<template>
  <div class="app-container">
    <div class="block">
      <aside>
        <a href="https://panjiachen.github.io/vue-element-admin-site/guide/advanced/icon.html" target="_blank">泵组可靠性试验-声学监控系统 系统设置页面
        </a>
      </aside>
      <el-tabs v-model="activeName" @tab-click="handleClick">
        <el-tab-pane label="系统设置" name="first">
          <el-row :gutter="20">

            <div slot="header">
              <a class="link-type link-title" target="_blank" href="https://panjiachen.github.io/vue-element-admin-site/guide/advanced/theme.html">
                Theme documentation
              </a>
            </div>
            <div class="box-item">
              <el-card class="box-card">
                <div slot="header" class="clearfix">
                  <span>监控设置</span>
                  <!--          <el-button style="float: right; padding: 3px 0" type="text">操作按钮</el-button>-->
                </div>
                <el-form ref="form" :model="monitorForm" :label-position="labelPosition" label-width="150px" size="mini">
                  <el-form-item label="配置文件路径">
                    <el-input v-model="monitorForm.name" type="file" />
                  </el-form-item>
                  <el-form-item label="设置">
                    <el-input-number v-model="monitorForm.sample_rate" style="width: 100%" />
                  </el-form-item>
                  <el-form-item label="点数">
                    <el-input-number v-model="monitorForm.sample_numbers" style="width: 100%" />
                  </el-form-item>
                  <el-form-item label="文件存储位置">
                    <el-input v-model="monitorForm.file_storage" />
                  </el-form-item>
                  <el-form-item size="large">
                    <el-button type="primary" @click="onSubmit">立即设置</el-button>
                    <el-button>取消</el-button>
                  </el-form-item>
                </el-form>
              </el-card>

            </div>

          </el-row></el-tab-pane>
        <!--      <el-tab-pane label="配置管理" name="second">配置管理</el-tab-pane>
        <el-tab-pane label="角色管理" name="third">角色管理</el-tab-pane>-->

      </el-tabs>

    </div>

    <!--
    <el-row :gutter="20">
      <el-col :span="6">
        <div class="chart-wrapper">
        <el-card class="box-card">
          <div slot="header">
            <a class="link-type link-title" target="_blank" href="https://panjiachen.github.io/vue-element-admin-site/guide/advanced/theme.html">
              Theme documentation
            </a>
          </div>
          <div class="box-item">
            <span class="field-label">Change Theme : </span>
            <el-switch v-model="theme" />
            <aside style="margin-top:15px;">
              Tips: It is different from the theme-pick on the navbar is two different skinning methods, each with different application scenarios. Refer to the documentation for details.
            </aside>
          </div>
        </el-card>
      </div>
      </el-col>
      <el-col :span="6">       <div class="chart-wrapper">
        <el-card class="box-card">
          <div slot="header">
            <a class="link-type link-title" target="_blank" href="https://panjiachen.github.io/vue-element-admin-site/guide/advanced/theme.html">
              Theme documentation
            </a>
          </div>
          <div class="box-item">
            <span class="field-label">Change Theme : </span>
            <el-switch v-model="theme" />
            <aside style="margin-top:15px;">
              Tips: It is different from the theme-pick on the navbar is two different skinning methods, each with different application scenarios. Refer to the documentation for details.
            </aside>
          </div>
        </el-card>
      </div></el-col>
      <el-col :span="6">       <div class="chart-wrapper">
        <el-card class="box-card">
          <div slot="header">
            <a class="link-type link-title" target="_blank" href="https://panjiachen.github.io/vue-element-admin-site/guide/advanced/theme.html">
              Theme documentation
            </a>
          </div>
          <div class="box-item">
            <span class="field-label">Change Theme : </span>
            <el-switch v-model="theme" />
            <aside style="margin-top:15px;">
              Tips: It is different from the theme-pick on the navbar is two different skinning methods, each with different application scenarios. Refer to the documentation for details.
            </aside>
          </div>
        </el-card>
      </div></el-col>
      <el-col :span="6">       <div class="chart-wrapper">
        <el-card class="box-card">
          <div slot="header">
            <a class="link-type link-title" target="_blank" href="https://panjiachen.github.io/vue-element-admin-site/guide/advanced/theme.html">
              Theme documentation
            </a>
          </div>
          <div class="box-item">
            <span class="field-label">Change Theme : </span>
            <el-switch v-model="theme" />
            <aside style="margin-top:15px;">
              Tips: It is different from the theme-pick on the navbar is two different skinning methods, each with different application scenarios. Refer to the documentation for details.
            </aside>
          </div>
        </el-card>
      </div></el-col>
    </el-row>
    <div class="block">
      <el-button type="primary">
        Primary
      </el-button>
      <el-button type="success">
        Success
      </el-button>
      <el-button type="info">
        Info
      </el-button>
      <el-button type="warning">
        Warning
      </el-button>
      <el-button type="danger">
        Danger
      </el-button>
    </div>

    <div class="block">
      <el-button type="primary" icon="el-icon-edit" />
      <el-button type="primary" icon="el-icon-share" />
      <el-button type="primary" icon="el-icon-delete" />
      <el-button type="primary" icon="el-icon-search">
        Search
      </el-button>
      <el-button type="primary">
        Upload
        <i class="el-icon-upload el-icon-right" />
      </el-button>
    </div>

    <div class="block">
      <el-tag v-for="tag in tags" :key="tag.type" :type="tag.type" class="tag-item">
        {{ tag.name }}
      </el-tag>
    </div>

    <div class="block">
      <el-radio-group v-model="radio">
        <el-radio :label="3">
          Option A
        </el-radio>
        <el-radio :label="6">
          Option B
        </el-radio>
        <el-radio :label="9">
          Option C
        </el-radio>
      </el-radio-group>
    </div>

    <div class="block">
      <el-slider v-model="slideValue" />
    </div>-->

  </div>
</template>

<script>
import { toggleClass } from '@/utils'
import '@/assets/custom-theme/index.css' // the theme changed version element-ui css

export default {
  name: 'Theme',
  data() {
    return {
      theme: false,
      activeName: 'first',
      labelPosition: 'right',
      tags: [
        { name: 'Tag One', type: '' },
        { name: 'Tag Two', type: 'info' },
        { name: 'Tag Three', type: 'success' },
        { name: 'Tag Four', type: 'warning' },
        { name: 'Tag Five', type: 'danger' }
      ],
      slideValue: 50,
      radio: 3,
      monitorForm: {
        name: '/data/',
        region: '',
        file_storage: '/data',
        sample_numbers: 100,
        sample_rate: 50,
        desc: ''
      }
    }
  },
  watch: {
    theme() {
      toggleClass(document.body, 'custom-theme')
    }
  }
}
</script>

<style scoped>
.field-label{
  vertical-align: middle;
}
.box-card {
  width: 4000px;
  max-width: 70%;
  margin: 20px auto;
  margin-left: 2%;
}

.block {
  padding: 30px 24px;
}

.tag-item {
  margin-right: 15px;
}
</style>

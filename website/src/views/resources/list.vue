<template>
  <div class="app-container">
    <el-table v-loading="listLoading" :data="list" border fit highlight-current-row style="width: 100%">
      <el-table-column align="center" label="ID" width="80">
        <template slot-scope="scope">
          <span>{{ scope.row.id }}</span>
        </template>
      </el-table-column>

    <el-table-column width="200px" align="center" label="Date">
        <template slot-scope="scope">
          <span>{{ scope.row.created_at  }}</span>
        </template>
      </el-table-column>
      <el-table-column width="120px" align="center" label="Name">
        <template slot-scope="scope">
          <span>{{ scope.row.doc_name }}</span>
        </template>
      </el-table-column>
          <el-table-column width="120px" align="center" label="Author">
            <template slot-scope="scope">
              <span>{{ scope.row.doc_author }}</span>
            </template>
          </el-table-column>

      <el-table-column width="120px" align="center" label="Category">
        <template slot-scope="scope">
          <span>{{ scope.row.doc_category }}</span>
        </template>
      </el-table-column>
      <el-table-column width="120px" align="center" label="Tags">
        <template slot-scope="scope">
          <span>{{ scope.row.doc_tags }}</span>
        </template>
      </el-table-column>
      <el-table-column width="120px" align="center" label="Links">
        <template slot-scope="scope">
          <span>{{ scope.row.doc_link }}</span>
        </template>
      </el-table-column>
      <el-table-column  label="Content">
        <template slot-scope="scope">
          <span>{{ scope.row.doc_content }}</span>
        </template>
      </el-table-column>
      <!--


              <el-table-column class-name="status-col" label="Status" width="110">
                <template slot-scope="{row}">
                  <el-tag :type="row.status | statusFilter">
                    {{ row.status }}
                  </el-tag>
                </template>
              </el-table-column>

              <el-table-column min-width="300px" label="Title">
                <template slot-scope="{row}">
                  <router-link :to="'/example/edit/'+row.id" class="link-type">
                    <span>{{ row.title }}</span>
                  </router-link>
                </template>
              </el-table-column>
-->
              <el-table-column align="right" label="Actions"  width="120px">
                <template slot-scope="scope">
                  <router-link :to="'/resource/edit/'+scope.row.id">
                    <el-button type="primary" size="small" icon="el-icon-edit">
                      Edit
                    </el-button>
                  </router-link>
                </template>
              </el-table-column>
    </el-table>

    <pagination v-show="total>0" :total="total" :page.sync="listQuery.page" :limit.sync="listQuery.limit" @pagination="getList" />
  </div>
</template>

<script>
import { getDocs } from '@/api/resources-docs'
import Pagination from '@/components/Pagination' // Secondary package based on el-pagination

export default {
  name: 'ArticleList',
  components: { Pagination },
  filters: {
    statusFilter(status) {
      const statusMap = {
        published: 'success',
        draft: 'info',
        deleted: 'danger'
      }
      return statusMap[status]
    }
  },
  data() {
    return {
      list: null,
      total: 0,
      listLoading: true,
      listQuery: {
        page: 1,
        limit: 20
      }
    }
  },
  created() {
    this.getList()
  },
  methods: {
    getList() {
      this.listLoading = true
      getDocs(this.listQuery).then(response => {

        this.list = response.spec
        this.total = response.total
        console.log("===?",this.list,this.total)
        this.listLoading = false
      })
    }
  }
}
</script>

<style scoped>
.edit-input {
  padding-right: 100px;
}
.cancel-btn {
  position: absolute;
  right: 15px;
  top: 10px;
}
</style>

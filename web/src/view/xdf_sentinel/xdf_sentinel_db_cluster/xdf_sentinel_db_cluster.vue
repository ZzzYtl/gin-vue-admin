<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">        
        <el-form-item label="哨兵集群">
          <el-cascader 
                :options="sentinelClusters"
                :props="{label:'name',value:'sentinel_cluster_id',emitPath:false}" 
                v-model="searchInfo.sentinel_cluster_id">
              </el-cascader>
        </el-form-item>    
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="openDialog" type="primary">新增监控</el-button>
        </el-form-item>
        <el-form-item>
          <el-popover placement="top" v-model="deleteVisible" width="160">
            <p>确定要删除吗？</p>
              <div style="text-align: right; margin: 0">
                <el-button @click="deleteVisible = false" size="mini" type="text">取消</el-button>
                <el-button @click="onDelete" size="mini" type="primary">确定</el-button>
              </div>
            <!-- 
              <el-button icon="el-icon-delete" size="mini" slot="reference" type="danger">批量删除</el-button>
            -->
          </el-popover>
        </el-form-item>
      </el-form>
    </div>
    <el-table
      :data="tableData"
      @selection-change="handleSelectionChange"
      border
      ref="multipleTable"
      stripe
      style="width: 100%"
      tooltip-effect="dark"
    >
    <el-table-column type="expand">
      <template slot-scope="scope">
        <el-form label-position="left" inline class="demo-table-expand">
          <el-row v-for="(domain,idx) in toJson(scope.row.dbs)" :key="idx">
          <el-col :span="6" :offset="2">
            <el-form-item label="ip:">
              {{domain.ip}}
            </el-form-item>
          </el-col>
          <el-col :span="6">
            <el-form-item label="角色:"> 
              {{filterDict(domain.type,"db_role")}}
            </el-form-item>
          </el-col>

          <el-col :span="6">
            <el-button v-if="domain.type === 4"
              @click="ChangeDBRole(scope.row, domain)"
              size="small"
              type="danger">变更为从库</el-button>
          </el-col>

        </el-row>
        </el-form>
      </template>
    </el-table-column>

    <!--
      <el-table-column type="selection" width="55"></el-table-column>
    -->
    <el-table-column label="集群" prop="cluster_id" width="120">
      <template slot-scope="scope">
        <div> 
          {{scope.row.cluster_id|clusterNameFilter}}
        </div>
      </template>
    </el-table-column> 
    
    <el-table-column label="实例" prop="logic_cluster_id" width="120">
      <template slot-scope="scope">
        <div> 
          {{scope.row.logic_cluster_id|logicClusterNameFilter}}
        </div>
      </template>
    </el-table-column> 
    
    <el-table-column label="LeaderEpoch" prop="leader_epoch" width="120"></el-table-column> 
    
    <el-table-column label="rlpc_user" prop="rlpc_user" width="120"></el-table-column> 
    
    <el-table-column label="rlpc_pwd" prop="rlpc_pwd" width="120"></el-table-column> 
    
    <el-table-column label="哨兵集群" prop="sentinel_cluster_id" width="120">
      <template slot-scope="scope">
        <div> 
          {{scope.row.sentinel_cluster_id|sentinelClusterFilter}}
        </div>
      </template>
    </el-table-column> 

      <el-table-column label="按钮组">
        <template slot-scope="scope">
          <!--
            <el-button class="table-button" @click="updateSentinelDBClusterInfo(scope.row)" size="small" type="primary" icon="el-icon-edit">变更</el-button>
          -->
          <el-popover placement="top" width="160" v-model="scope.row.visible">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin: 0">
              <el-button size="mini" type="text" @click="scope.row.visible = false">取消</el-button>
              <el-button type="primary" size="mini" @click="deleteSentinelDBClusterInfo(scope.row)">确定</el-button>
            </div>
            <el-button type="danger" icon="el-icon-delete" size="mini" slot="reference">删除</el-button>
          </el-popover>            
        </template>
      </el-table-column>
    </el-table>

    <el-pagination
      :current-page="page"
      :page-size="pageSize"
      :page-sizes="[10, 30, 50, 100]"
      :style="{float:'right',padding:'20px'}"
      :total="total"
      @current-change="handleCurrentChange"
      @size-change="handleSizeChange"
      layout="total, sizes, prev, pager, next, jumper"
    ></el-pagination>

    <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" title="新增监控">
      <el-form :model="formData" label-position="right" label-width="120px">
        <el-form-item label="集群名:">
          <el-cascader 
                :options="clusters"
                :props="{label:'name',value:'tag_id',emitPath:false}" 
                v-model="formData.cluster_id">
          </el-cascader>
        </el-form-item>
        <el-form-item label="实例名:">
          <el-cascader 
                :options="lClusters"
                :props="{label:'name',value:'logic_cluster_id',emitPath:false}" 
                v-model="formData.logic_cluster_id">
          </el-cascader>
        </el-form-item>

        <el-form-item label="rlpc_user:">
          <el-input v-model="formData.rlpc_user" clearable placeholder="请输入" ></el-input>
        </el-form-item>
       
        <el-form-item label="rlpc_pwd:">
          <el-input v-model="formData.rlpc_pwd" clearable placeholder="请输入" ></el-input>
        </el-form-item>
       
        <el-form-item label="哨兵集群:">
          <el-cascader 
              :options="sentinelClusters"
              :props="{label:'name',value:'sentinel_cluster_id',emitPath:false}" 
            v-model="formData.sentinel_cluster_id">
          </el-cascader>
        </el-form-item>
      </el-form>
      <div class="dialog-footer" slot="footer">
        <el-button @click="closeDialog">取 消</el-button>
        <el-button @click="enterDialog" type="primary">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import {
    createSentinelDBClusterInfo,
    deleteSentinelDBClusterInfo,
    changeOldMaster2Slave,
    deleteSentinelDBClusterInfoByIds,
    updateSentinelDBClusterInfo,
    findSentinelDBClusterInfo,
    getSentinelDBClusterInfoList
} from "@/api/xdf_sentinel_db_cluster";  //  此处请自行替换地址
import { formatTimeToStr } from "@/utils/date";
import { getAllSentinelClusters } from "@/api/xdf_sentinel_cluster";
import { getAllTags } from "@/api/xdf_tag";
import { getAllLogicClusters } from "@/api/xdf_logic_cluster";
import infoList from "@/mixins/infoList";

var SentinelClusters = [
  {
    sentinel_cluster_id: 6,
    name: "1111"
  }
];

var Clusters = [
  {
    tag_id: 6,
    name: "1111"
  }
];

var LClusters = [
  {
    logic_cluster_id: 6,
    logic_cluster_name: "1111"
  }
];

export default {
  name: "SentinelDBClusterInfo",
  mixins: [infoList],
  data() {
    return {
      listApi: getSentinelDBClusterInfoList,
      dialogFormVisible: false,
      visible: false,
      row_visible: false,
      type: "",
      sentinelClusters: SentinelClusters,
      clusters: Clusters,
      lClusters: LClusters,
      deleteVisible: false,
      multipleSelection: [],
      formData: {
            cluster_id:0,
            logic_cluster_id:0,
            leader_epoch:0,
            rlpc_user:"",
            rlpc_pwd:"",
            sentinel_cluster_id:0,
            
      }
    };
  },
  filters: {
    formatDate: function(time) {
      if (time != null && time != "") {
        var date = new Date(time);
        return formatTimeToStr(date, "yyyy-MM-dd hh:mm:ss");
      } else {
        return "";
      }
    },
    formatBoolean: function(bool) {
      if (bool != null) {
        return bool ? "是" :"否";
      } else {
        return "";
      }
    },
    sentinelClusterFilter(value) {
      const target = SentinelClusters.filter(item => item.sentinel_cluster_id === value)[0];
      return target && `${target.name}`;
    },
    clusterNameFilter(value) {
      const target = Clusters.filter(item=> item.tag_id === value)[0];
      return target && `${target.name}`;
    },
    logicClusterNameFilter(value) {
      const target = LClusters.filter(item=> item.logic_cluster_id === value)[0];
      return target && `${target.name}`;
    }
  },
  methods: {
      toJson:function(str){
        var _str =JSON.parse(str);
        return _str;
      },

      //条件搜索前端看此方法
      onSubmit() {
        this.page = 1
        this.pageSize = 10         
        this.getTableData()
      },
      handleSelectionChange(val) {
        this.multipleSelection = val
      },
      async onDelete() {
        const ids = []
        if(this.multipleSelection.length == 0){
          this.$message({
            type: 'warning',
            message: '请选择要删除的数据'
          })
          return
        }
        this.multipleSelection &&
          this.multipleSelection.map(item => {
            ids.push(item.cluster_id)
          })
        const res = await deleteSentinelDBClusterInfoByIds({ ids })
        if (res.code == 0) {
          this.$message({
            type: 'success',
            message: '删除成功'
          })
          if (this.tableData.length == ids.length) {
              this.page--;
          }
          this.deleteVisible = false
          this.getTableData()
        }
      },
    async updateSentinelDBClusterInfo(row) {
      const res = await findSentinelDBClusterInfo({ cluster_id: row.cluster_id });
      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.reSentinelDBCluster;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
          cluster_id:0,
          logic_cluster_id:0,
          leader_epoch:0,
          rlpc_user:"",
          rlpc_pwd:"",
          sentinel_cluster_id:0,
      };
    },
    async deleteSentinelDBClusterInfo(row) {
      this.visible = false;
      const res = await deleteSentinelDBClusterInfo({ id: row.id });
      if (res.code == 0) {
        this.$message({
          type: "success",
          message: "删除成功"
        });
        if (this.tableData.length == 1) {
            this.page--;
        }
        this.getTableData();
      }
    },
    async enterDialog() {
      let res;
      switch (this.type) {
        case "create":
          res = await createSentinelDBClusterInfo(this.formData);
          break;
        case "update":
          res = await updateSentinelDBClusterInfo(this.formData);
          break;
        default:
          res = await createSentinelDBClusterInfo(this.formData);
          break;
      }
      if (res.code == 0) {
        this.$message({
          type:"success",
          message:"创建/更改成功"
        })
        this.closeDialog();
        this.getTableData();
      }
    },
    openDialog() {
      this.type = "create";
      this.dialogFormVisible = true;
    },

    // 提升旧主为从库
    ChangeDBRole(row, domain) {
      this.$confirm("此操作将旧主库提升为从库, 是否继续?", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning"
      }).then(async () => {
          const res = await changeOldMaster2Slave({id:row.id, ip: domain.ip});
          if (res.code == 0) {
            this.$message({
              type: "success",
              message: "提升成功!"
            });
            if (this.tableData.length == 1) {
              this.page--;
            }
            this.getTableData();
          } else {
            this.$message({
              type: "warning",
              message: res.msg
            })
          }
        })
        .catch(() => {
          this.$message({
            type: "info",
            message: "已取消提升操作"
          });
        });
    }
  },

  async created() {
    var rst = await getAllSentinelClusters();
    if (rst.code == 0) {
      this.sentinelClusters = rst.data.list
      SentinelClusters = this.sentinelClusters
    }

    var cluster_rst = await getAllTags();
    if (cluster_rst.code == 0) {
      this.clusters = cluster_rst.data.list
      Clusters = this.clusters
    }

    var lcluster_rst = await getAllLogicClusters();
    if (lcluster_rst.code == 0) {
      this.lClusters = lcluster_rst.data.list
      LClusters = this.lClusters
    }

    await this.getTableData();
    await this.getDict("db_role");
  }
};
</script>

<style>
</style>
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
            <el-button icon="el-icon-delete" size="mini" slot="reference" type="danger">批量删除</el-button>
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
    <el-table-column type="selection" width="55"></el-table-column>
    
    <el-table-column label="数据库" prop="cluster_id" width="120">
      <template slot-scope="scope">
        <div> 
          {{scope.row.cluster_id|dbNameFilter}}
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
          <el-button class="table-button" @click="updateSentinelDBClusterInfo(scope.row)" size="small" type="primary" icon="el-icon-edit">变更</el-button>
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
         <el-form-item label="数据库:">
         <el-cascader 
                :options="dbs"
                :props="{label:'db_name',value:'cluster_id',emitPath:false}" 
                v-model="formData.cluster_id">
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
    deleteSentinelDBClusterInfoByIds,
    updateSentinelDBClusterInfo,
    findSentinelDBClusterInfo,
    getSentinelDBClusterInfoList
} from "@/api/xdf_sentinel_db_cluster";  //  此处请自行替换地址
import { formatTimeToStr } from "@/utils/date";
import { getAllSentinelClusters } from "@/api/xdf_sentinel_cluster";
import { getAllDBs } from "@/api/xdf_db_info";
import infoList from "@/mixins/infoList";

var SentinelClusters = [
  {
    sentinel_cluster_id: 6,
    name: "1111"
  }
];

var DBs = [
  {
    cluster_id: 6,
    db_name: "1111"
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
      type: "",
      sentinelClusters: SentinelClusters,
      dbs: DBs,
      deleteVisible: false,
      multipleSelection: [],formData: {
            cluster_id:0,
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
    dbNameFilter(value) {
      const target = DBs.filter(item=> item.cluster_id === value)[0];
      return target && `${target.db_name}`;
    }
  },
  methods: {
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
          leader_epoch:0,
          rlpc_user:"",
          rlpc_pwd:"",
          sentinel_cluster_id:0,
          
      };
    },
    async deleteSentinelDBClusterInfo(row) {
      this.visible = false;
      const res = await deleteSentinelDBClusterInfo({ cluster_id: row.cluster_id });
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
    }
  },
  async created() {
    var rst = await getAllSentinelClusters();
    if (rst.code == 0) {
      this.sentinelClusters = rst.data.list
      SentinelClusters = this.sentinelClusters
    }

    var db_rst = await getAllDBs();
    if (db_rst.code == 0) {
      this.dbs = db_rst.data.list
      DBs = this.dbs
    }

    await this.getTableData();
  
}
};
</script>

<style>
</style>
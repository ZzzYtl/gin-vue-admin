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
        <el-form-item label="IP">
          <el-input placeholder="搜索条件" v-model="searchInfo.ip"></el-input>
        </el-form-item>    
        <el-form-item label="port">
          <el-input placeholder="搜索条件" v-model="searchInfo.port"></el-input>
        </el-form-item>        
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="openDialog" type="primary">新增哨兵节点</el-button>
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
    
    <el-table-column label="ID" prop="sentinel_id" width="120"></el-table-column> 
    
    <el-table-column label="RunID" prop="run_id" width="120"></el-table-column> 
    
    <el-table-column label="哨兵集群" prop="sentinel_cluster_id" width="120">
        <template slot-scope="scope">
          <div> 
            {{scope.row.sentinel_cluster_id|sentinelClusterFilter}}
          </div>
      </template>
    </el-table-column> 
    
    <el-table-column label="IP" prop="ip" width="120"></el-table-column> 
    
    <el-table-column label="port" prop="port" width="120"></el-table-column> 
    
    <el-table-column label="CurrentEpoch" prop="current_epoch" width="120"></el-table-column> 
    
    <el-table-column label="ConfigEpoch" prop="config_epoch" width="120"></el-table-column> 
    
      <el-table-column label="按钮组">
        <template slot-scope="scope">
          <el-button class="table-button" @click="updateSentinelInfo(scope.row)" size="small" type="primary" icon="el-icon-edit">变更</el-button>
          <el-popover placement="top" width="160" v-model="scope.row.visible">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin: 0">
              <el-button size="mini" type="text" @click="scope.row.visible = false">取消</el-button>
              <el-button type="primary" size="mini" @click="deleteSentinelInfo(scope.row)">确定</el-button>
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

    <el-dialog :before-close="closeCreateDialog" :visible.sync="createDialogFormVisible" title="新增哨兵集群">
      <el-form :model="createData"  label-position="right" label-width="120px">
         <el-form-item label="哨兵集群名称:">
            <el-input v-model="createData.cluster_name" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
      <el-row>
          <el-col :span="12">
              <el-form-item label="sentinel_ip:">
                <el-input v-model="createData.sentinel1_ip" clearable placeholder="请输入" ></el-input>
          </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="sentinel_port:">
                <el-input v-model="createData.sentinel1_port" clearable placeholder="请输入" ></el-input>
            </el-form-item>
          </el-col>
      </el-row>
      <el-row>
          <el-col :span="12">
              <el-form-item label="sentinel_ip:">
                <el-input v-model="createData.sentinel2_ip" clearable placeholder="请输入" ></el-input>
          </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="sentinel_port:">
                <el-input v-model="createData.sentinel2_port" clearable placeholder="请输入" ></el-input>
            </el-form-item>
          </el-col>
      </el-row>

      <el-row>
          <el-col :span="12">
              <el-form-item label="sentinel_ip:">
                <el-input v-model="createData.sentinel3_ip" clearable placeholder="请输入" ></el-input>
          </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="sentinel_port:">
                <el-input v-model="createData.sentinel3_port" clearable placeholder="请输入" ></el-input>
            </el-form-item>
          </el-col>
      </el-row>
       </el-form>
      <div class="dialog-footer" slot="footer">
        <el-button @click="closeCreateDialog">取 消</el-button>
        <el-button @click="enterCreateDialog" type="primary">确 定</el-button>
      </div>
    </el-dialog>

    <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" title="弹窗操作">
      <el-form :model="formData" label-position="right" label-width="80px">              
         <el-form-item label="IP:">
            <el-input v-model="formData.ip" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="port:"><el-input v-model.number="formData.port" clearable placeholder="请输入"></el-input>
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
    createSentinelInfo,
    deleteSentinelInfo,
    deleteSentinelInfoByIds,
    updateSentinelInfo,
    findSentinelInfo,
    getSentinelInfoList
} from "@/api/xdf_sentinel_info";  //  此处请自行替换地址
import { formatTimeToStr } from "@/utils/date";
import infoList from "@/mixins/infoList";
import { getAllSentinelClusters } from "@/api/xdf_sentinel_cluster";

var SentinelClusters = [
  {
    sentinel_cluster_id: 6,
    name: "1111"
  }
];

export default {
  name: "SentinelInfo",
  mixins: [infoList],
  data() {
    return {
      listApi: getSentinelInfoList,
      dialogFormVisible: false,
      createDialogFormVisible: false,
      visible: false,
      sentinelClusters: SentinelClusters,
      type: "",
      deleteVisible: false,
      createData: {
          cluster_name: "",
          sentinel1_ip:"",
          sentinel1_port: 0,
          sentinel2_ip:"",
          sentinel2_port: 0,
          sentinel3_ip:"",
          sentinel3_port: 0,
      },
      multipleSelection: [],formData: {
            ip:"",
            port:0,    
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
            ids.push(item.sentinel_id)
          })
        const res = await deleteSentinelInfoByIds({ ids })
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
    async updateSentinelInfo(row) {
      const res = await findSentinelInfo({ sentinel_id: row.sentinel_id });
      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.resentinelinfo;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
          ip:"",
          port:0, 
      };
    },

    closeCreateDialog() {
      this.createDialogFormVisible = false;
      this.createData = {
          cluster_name: "",
          sentinel1_ip:"",
          sentinel1_port: 0,
          sentinel2_ip:"",
          sentinel2_port: 0,
          sentinel3_ip:"",
          sentinel3_port: 0,
      }
    },

    async deleteSentinelInfo(row) {
      this.visible = false;
      const res = await deleteSentinelInfo({ sentinel_id: row.sentinel_id });
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
        case "update":
          res = await updateSentinelInfo(this.formData);
          break;
        default:
          res = await createSentinelInfo(this.formData);
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

    async enterCreateDialog() {
        var sentinels = [
          {
            ip: "",
            port: 0
          }
        ]
        sentinels = []

        if (this.createData.sentinel1_ip.length != 0) {
          sentinels.push({
            ip: this.createData.sentinel1_ip,
            port: this.createData.sentinel1_port,
            })
        }

        if (this.createData.sentinel2_ip.length != 0) {
          sentinels.push({
            ip: this.createData.sentinel2_ip,
            port: this.createData.sentinel2_port,
            })
        }

        if (this.createData.sentinel3_ip.length != 0) {
          sentinels.push({
            ip: this.createData.sentinel3_ip,
            port: this.createData.sentinel3_port,
            })
        }
        let res = await createSentinelInfo({
          cluster_name:  this.createData.cluster_name,
          sentinels: sentinels,
          })

        if (res.code == 0) {
          this.$message({
            type:"success",
            message:"创建成功"
          })
          this.closeCreateDialog();
          this.getTableData();
        }
    },

    openDialog() {
      this.type = "create";
      this.createDialogFormVisible = true;
    }
  },
  async created() {
    var rst = await getAllSentinelClusters();
    if (rst.code == 0) {
      this.sentinelClusters = rst.data.list
      SentinelClusters = this.sentinelClusters
    }
    await this.getTableData();
  
}
};
</script>

<style>
</style>
<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">  
        <el-form-item label="集群名">
           <el-cascader 
            :options="tags"
            :props="{label:'name',value:'tag_id',emitPath:false}" 
            v-model="searchInfo.tag_id">
          </el-cascader>
        </el-form-item>    
        <el-form-item label="数据库类型">
          <el-select v-model="searchInfo.db_type" placeholder="请选择" clearable>
                 <el-option v-for="(item,key) in db_typeOptions" :key="key" :label="item.label" :value="item.value"></el-option>
             </el-select>
        </el-form-item>    
        <el-form-item label="业务线">
          <el-cascader 
            :options="depts" 
            :props="{label:'dept_name', value: 'dept_id', emitPath:false}"
            v-model="searchInfo.dept_id">
          </el-cascader>
        </el-form-item>    
        <el-form-item label="端口">
          <el-input placeholder="搜索条件" v-model="searchInfo.port"></el-input>
        </el-form-item>    
        <el-form-item label="数据库名">
          <el-input placeholder="搜索条件" v-model="searchInfo.db_name"></el-input>
        </el-form-item>              
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="openDialog" type="primary">新增数据库信息</el-button>
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

    <el-table-column type="expand">
        <template slot-scope="scope">
          <el-form label-position="left" inline class="demo-table-expand">
            <el-form-item label="用户名">
              {{scope.row.db_user}} <br>
            </el-form-item> 
    
            <el-form-item label="密码">
              {{scope.row.db_pwd}} <br>
            </el-form-item> 
            
            <el-form-item label="备份库">
              {{scope.row.backup_id | backupDBFilter}} <br>
            </el-form-item> 
            
            <el-form-item label="延迟备份库">
              {{scope.row.delay_backup_id | backupDBFilter}} <br>
            </el-form-item> 
          </el-form>
        </template>
      </el-table-column>
    
    <el-table-column label="业务线" prop="dept_id" width="120">
      <template slot-scope="scope">
        <div> 
          {{scope.row.dept_id|deptFilter}}
        </div>
      </template>
    </el-table-column> 

    <el-table-column label="集群名" prop="tag_id" width="120">
        <template slot-scope="scope">
          <div> 
            {{scope.row.tag_id|clusterFilter}}
          </div>
        </template>
    </el-table-column> 
    
    <el-table-column label="实例名" prop="cluster_name" width="120"></el-table-column> 
    
    <el-table-column label="数据库名" prop="db_name" width="120"></el-table-column> 

    <el-table-column label="端口" prop="port" width="120"></el-table-column> 

    <el-table-column label="数据库类型" prop="db_type" width="120">
      <template slot-scope="scope">
        {{filterDict(scope.row.db_type,"db_type")}}
      </template>
    </el-table-column>
  
    <el-table-column label="zk名" prop="zk_name" width="120"></el-table-column> 
    
      <el-table-column label="按钮组">
        <template slot-scope="scope">
          <el-button class="table-button" @click="updateDataBaseInfo(scope.row)" size="small" type="primary" icon="el-icon-edit">变更</el-button>
          <el-popover placement="top" width="160" v-model="scope.row.visible">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin: 0">
              <el-button size="mini" type="text" @click="scope.row.visible = false">取消</el-button>
              <el-button type="primary" size="mini" @click="deleteDataBaseInfo(scope.row)">确定</el-button>
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

    <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" title="修改">
      <el-form :model="formData" label-position="right" label-width="80px">
         <el-form-item label="集群名:">
          <el-cascader 
                :options="tags"
                :props="{label:'name',value:'tag_id',emitPath:false}" 
                v-model="formData.tag_id">
              </el-cascader>
      </el-form-item>
       
         <el-form-item label="实例名:">
            <el-input v-model="formData.cluster_name" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="数据库类型:">
             <el-select v-model="formData.db_type" placeholder="请选择" clearable>
                 <el-option v-for="(item,key) in db_typeOptions" :key="key" :label="item.label" :value="item.value"></el-option>
             </el-select>
      </el-form-item>
       
         <el-form-item label="业务线:">
          <el-cascader 
            :options="depts" 
            :props="{label:'dept_name', value: 'dept_id', emitPath:false}"
            v-model="formData.dept_id">
          </el-cascader>
      </el-form-item>
       
         <el-form-item label="端口:"><el-input v-model.number="formData.port" clearable placeholder="请输入"></el-input>
      </el-form-item>
       
         <el-form-item label="数据库名:">
            <el-input v-model="formData.db_name" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="用户名:">
            <el-input v-model="formData.db_user" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="密码:">
            <el-input v-model="formData.db_pwd" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="zk名:">
            <el-input v-model="formData.zk_name" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="备份库:">
          <el-cascader 
            :options="backupDBs" 
            :props="{label:'IP', value: 'backup_id', emitPath:false}"
            v-model="formData.backup_id">
          </el-cascader>
      </el-form-item>
       
         <el-form-item label="延迟备库:">
         <el-cascader 
            :options="backupDBs" 
            :props="{label:'IP', value: 'backup_id', emitPath:false}"
            v-model="formData.delay_backup_id">
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
    createDataBaseInfo,
    deleteDataBaseInfo,
    deleteDataBaseInfoByIds,
    updateDataBaseInfo,
    findDataBaseInfo,
    getDataBaseInfoList
} from "@/api/xdf_db_info";  //  此处请自行替换地址
import { formatTimeToStr } from "@/utils/date";
import infoList from "@/mixins/infoList";
import { getAllTags } from "@/api/xdf_tag";
import { getAllDepts } from "@/api/xdf_dept";
import { getAllBackUpDBs } from "@/api/xdf_backup_db";
 
var Tags = [
  {
    tag_id: 6,
    name: "1111"
  }
];

var Depts = [
  {
    dept_id: 1,
    dept_name: "111"
  }
];

var BackUpDBs = [
  {
    backup_id: 1,
    IP: ""
  }
];

export default {
  name: "DataBaseInfo",
  mixins: [infoList],
  data() {
    return {
      listApi: getDataBaseInfoList,
      dialogFormVisible: false,
      visible: false,
      type: "",
      tags : Tags,
      depts: Depts,
      backupDBs : BackUpDBs,
      deleteVisible: false,
      multipleSelection: [],
      db_typeOptions:[],
          formData: {
            cluster_id:0,
            tag_id:0,
            cluster_name:"",
            db_type:3,
            dept_id:0,
            port:0,
            db_name:"",
            db_user:"",
            db_pwd:"",
            zk_name:"",
            backup_id:0,
            delay_backup_id:0,
            
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
    clusterFilter(value) {
      const target = Tags.filter(item => item.tag_id === value)[0];
      return target && `${target.name}`;
    },
    deptFilter(value) {
      const target = Depts.filter(item => item.dept_id === value)[0];
      return target && `${target.dept_name}`;
    },
    backupDBFilter(value) {
      const target = BackUpDBs.filter(item => item.backup_id === value)[0];
      return target && `${target.IP}`;
    },
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
            ids.push(item.ID)
          })
        const res = await deleteDataBaseInfoByIds({ ids })
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
    async updateDataBaseInfo(row) {
      const res = await findDataBaseInfo({ cluster_id: row.cluster_id });
      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.reDBInfo;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
          cluster_id:0,
          tag_id:0,
          cluster_name:"",
          db_type:0,
          dept_id:0,
          port:0,
          db_name:"",
          db_user:"",
          db_pwd:"",
          zk_name:"",
          backup_id:0,
          delay_backup_id:0,
          
      };
    },
    async deleteDataBaseInfo(row) {
      this.visible = false;
      const res = await deleteDataBaseInfo({ cluster_id: row.cluster_id });
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
          res = await createDataBaseInfo(this.formData);
          break;
        case "update":
          res = await updateDataBaseInfo(this.formData);
          break;
        default:
          res = await createDataBaseInfo(this.formData);
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
    var rst = await getAllTags();
    if (rst.code == 0) {
      this.tags = rst.data.list
      Tags = this.tags
    }

    var dept_rst = await getAllDepts();
    if (dept_rst.code == 0) {
      this.depts = dept_rst.data.list
      Depts = this.depts
    }

    var backup_rst = await getAllBackUpDBs();
    if (backup_rst.code == 0) {
      this.backupDBs = backup_rst.data.list
      BackUpDBs = this.backupDBs
    }


    await this.getTableData();
    await this.getDict("db_type");
    
}
};
</script>

<style scoped>
  .demo-table-expand {
    font-size: 0;
  }
  .demo-table-expand label {
    width: 90px;
    color: #99a9bf;
  }
  .demo-table-expand .el-form-item {
    margin-right: 0;
    margin-bottom: 0;
    width: 50%;
  }
</style>
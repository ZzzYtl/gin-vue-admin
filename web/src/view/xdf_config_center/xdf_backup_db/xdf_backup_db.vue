<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">  
        <el-form-item label="IP">
          <el-input placeholder="搜索条件" v-model="searchInfo.IP"></el-input>
        </el-form-item>    
        <el-form-item label="备库/延时备库">
          <el-select v-model.trim="searchInfo.type">
                <el-option
                  v-for="item in backupDBTypes"
                  :key="item.value"
                  :label="item.label"
                  :value="item.value">
                </el-option>
              </el-select>
        </el-form-item>    
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="openDialog" type="primary">新增备库</el-button>
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
    <el-table-column label="ID" prop="backup_id" width="120"></el-table-column> 
    
    <el-table-column label="IP" prop="IP" width="120"></el-table-column> 
    
    <el-table-column label="备库/延时备库" prop="type" width="120">
      <template slot-scope="scope">
            <span style="color:black" v-if="scope.row.type === 1">备库</span>
            <span style="color:red" v-else-if="scope.row.type === 2">延时备库</span>
      </template>
    </el-table-column> 
    
    <el-table-column label="域" prop="domain" width="320"></el-table-column>

    <el-table-column label="操作">
        <template slot-scope="scope">
          <el-button class="table-button" @click="updateBackUpDB(scope.row)" size="small" type="primary" icon="el-icon-edit">变更</el-button>
          <el-popover placement="top" width="160" v-model="scope.row.visible">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin: 0">
              <el-button size="mini" type="text" @click="scope.row.visible = false">取消</el-button>
              <el-button type="primary" size="mini" @click="deleteBackUpDB(scope.row)">确定</el-button>
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
      :style="{float:'left',padding:'20px'}"
      :total="total"
      @current-change="handleCurrentChange"
      @size-change="handleSizeChange"
      layout="total, sizes, prev, pager, next, jumper"
    ></el-pagination>

    <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" title="新增备库">
      <el-form ref="apiForm" :model="formData"  :rules="formRules" label-position="right" label-width="120px">
      
      <el-form-item label="IP:" prop="IP">
            <el-input v-model="formData.IP" clearable placeholder="请输入" ></el-input>
      </el-form-item>

      <el-form-item label="domain:" prop="domain">
            <el-input v-model="formData.domain" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
      <el-form-item label="备库/延时备库:" prop="type">
        <el-select v-model.trim="formData.type">
            <el-option
              v-for="item in backupDBTypes"
              :key="item.value"
              :label="item.label"
              :value="item.value">
            </el-option>
          </el-select>
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
    createBackUpDB,
    deleteBackUpDB,
    deleteBackUpDBByIds,
    updateBackUpDB,
    findBackUpDB,
    getBackUpDBList
} from "@/api/xdf_backup_db";  //  此处请自行替换地址
import { formatTimeToStr } from "@/utils/date";
import infoList from "@/mixins/infoList";
export default {
  name: "BackUpDB",
  mixins: [infoList],
  data() {
    return {
      listApi: getBackUpDBList,
      dialogFormVisible: false,
      visible: false,
      type: "",
      deleteVisible: false,
      multipleSelection: [],formData: {
            backup_id:0,
            IP:"",
            domain:"",
            type : 1,           
      },
      formRules: {
        IP: [
          {required: true, message: '请输入IP地址', trigger: 'blur'}
        ],
        type: [
          {required: true, message: '请选择备库类型', trigger: 'blur'}
        ],
        domain: [
          {required: true, message: '请输入域名', trigger: 'blur'}
        ]
      },
      backupDBTypes: [
        {
          value: 1,
          label: '备库'
        },
        {
          value: 2,
          label: '延时从库'
        }
      ],
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
            ids.push(item.ID)
          })
        const res = await deleteBackUpDBByIds({ ids })
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
    async updateBackUpDB(row) {
      const res = await findBackUpDB({ backup_id: row.backup_id });
      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.reBackDB;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
          backup_id:0,
          IP:"",
          domain:"",
          type:1,
          
      };
    },
    async deleteBackUpDB(row) {
      this.visible = false;
      const res = await deleteBackUpDB({ backup_id: row.backup_id });
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
       this.$refs.apiForm.validate(async valid => {
        if (valid) {
          let res;
          switch (this.type) {
            case "create":
              res = await createBackUpDB(this.formData);
              break;
            case "update":
              res = await updateBackUpDB(this.formData);
              break;
            default:
              res = await createBackUpDB(this.formData);
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
        }
       })
    },
    openDialog() {
      this.type = "create";
      this.dialogFormVisible = true;
    }
  },
  async created() {
    await this.getTableData();
  
}
};
</script>

<style>
</style>
<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">  
        <el-form-item label="业务线名称">
          <el-input placeholder="搜索条件" v-model="searchInfo.dept_name"></el-input>
        </el-form-item>    
        <el-form-item label="负责人">
          <el-input placeholder="搜索条件" v-model="searchInfo.leader_name"></el-input>
        </el-form-item>    
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="openDialog" type="primary">新增业务线</el-button>
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
    <el-table-column label="ID" prop="dept_id" width="120"></el-table-column> 
    
    <el-table-column label="业务线名称" prop="dept_name" width="120"></el-table-column> 
    
    <el-table-column label="负责人" prop="leader_name" width="120"></el-table-column> 
    
      <el-table-column label="操作">
        <template slot-scope="scope">
          <el-button class="table-button" @click="updateDeptInfo(scope.row)" size="small" type="primary" icon="el-icon-edit">变更</el-button>
          <el-popover placement="top" width="160" v-model="scope.row.visible">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin: 0">
              <el-button size="mini" type="text" @click="scope.row.visible = false">取消</el-button>
              <el-button type="primary" size="mini" @click="deleteDeptInfo(scope.row)">确定</el-button>
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

    <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" title="新增业务线">
      <el-form :model="formData" :rules="formRules" label-position="right" label-width="120px">      
         <el-form-item label="业务线名称:" prop="dept_name">
            <el-input v-model="formData.dept_name" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="负责人:" prop="leader_name">
            <el-input v-model="formData.leader_name" clearable placeholder="请输入" ></el-input>
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
    createDeptInfo,
    deleteDeptInfo,
    deleteDeptInfoByIds,
    updateDeptInfo,
    findDeptInfo,
    getDeptInfoList
} from "@/api/xdf_dept";  //  此处请自行替换地址
import { formatTimeToStr } from "@/utils/date";
import infoList from "@/mixins/infoList";
export default {
  name: "DeptInfo",
  mixins: [infoList],
  data() {
    return {
      listApi: getDeptInfoList,
      dialogFormVisible: false,
      visible: false,
      type: "",
      deleteVisible: false,
      multipleSelection: [],formData: {
            dept_id:0,
            dept_name:"",
            leader_name:"",
            
      },
       formRules: {
        dept_name: [
          {required: true, message: '请输入业务线名称', trigger: 'blur'}
        ],
        leader_name: [
          {required: true, message: '请输入负责人', trigger: 'blur'}
        ]
      },
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
            ids.push(item.dept_id)
          })
        const res = await deleteDeptInfoByIds({ ids })
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
    async updateDeptInfo(row) {
      const res = await findDeptInfo({ dept_id: row.dept_id });
      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.redept;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
          dept_id:0,
          dept_name:"",
          leader_name:"",
          
      };
    },
    async deleteDeptInfo(row) {
      this.visible = false;
      const res = await deleteDeptInfo({ dept_id: row.dept_id });
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
              res = await createDeptInfo(this.formData);
              break;
            case "update":
              res = await updateDeptInfo(this.formData);
              break;
            default:
              res = await createDeptInfo(this.formData);
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
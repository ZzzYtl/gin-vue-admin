<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">  
        <el-form-item label="集群名称">
          <el-cascader 
            :options="tags"
            :props="{label:'name',value:'tag_id',emitPath:false}" 
            v-model="searchInfo.tag_id">
          </el-cascader>
        </el-form-item>    
        <el-form-item label="IP">
          <el-input placeholder="搜索条件" v-model="searchInfo.ip"></el-input>
        </el-form-item>    
        <el-form-item label="角色">
          <el-select v-model="searchInfo.role_id" placeholder="请选择" clearable>
            <el-option v-for="(item,key) in db_roleOptions" :key="key" :label="item.label" :value="item.value"></el-option>
          </el-select>
        </el-form-item>    
        <el-form-item label="机房">
          <el-cascader 
            :options="areas"
            :props="{label:'name',value:'area_id',emitPath:false}" 
            v-model="searchInfo.area_id">
          </el-cascader>
        </el-form-item>    
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="openDialog" type="primary">新增集群</el-button>
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
    <!--
    <el-table-column type="selection" width="55"></el-table-column>
    -->
    <el-table-column label="ID" prop="node_id" width="120"></el-table-column> 
    
    <el-table-column label="集群名称" prop="tag_id" width="120">
          <template slot-scope="scope">
          <div>
            
            {{scope.row.tag_id|clusterFilter}}
          </div>
        </template>
    </el-table-column> 
    
    <el-table-column label="IP" prop="ip" width="120">
    </el-table-column> 
    
    <el-table-column label="角色" prop="role_id" width="120">
      <template slot-scope="scope">
        <div>
          {{filterDict(scope.row.role_id,"db_role")}}
        </div>      
      </template>
    </el-table-column> 
    
    <el-table-column label="机房" prop="area_id" width="120">
      <template slot-scope="scope">
          <div>
            {{scope.row.area_id|areaFilter}}
          </div>
        </template>
    </el-table-column> 
    
      <el-table-column label="按钮组">
        <template slot-scope="scope">
          <el-button class="table-button" @click="updateNode(scope.row)" size="small" type="primary" icon="el-icon-edit">变更</el-button>
          <el-popover placement="top" width="160" v-model="scope.row.visible">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin: 0">
              <el-button size="mini" type="text" @click="scope.row.visible = false">取消</el-button>
              <el-button type="primary" size="mini" @click="deleteNode(scope.row)">确定</el-button>
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

    <el-dialog :before-close="closeCreateDialog" :visible.sync="createDialogFormVisible" title="新增集群">
      <el-form :model="createData" :rules="formRules"  label-position="right" label-width="90px">
        <el-form-item label="集群名称:" prop="cluster_name">
          <el-input v-model="createData.cluster_name" clearable placeholder="请输入" ></el-input>
        </el-form-item>
       
        <el-row>
          <el-col :span="12">
            <el-form-item label="master_ip:" prop="master_ip">
              <el-input v-model="createData.master_ip" clearable placeholder="请输入" ></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="所属机房:" prop="area_id">
              <el-cascader 
                :options="areas"
                :props="{label:'name',value:'area_id',emitPath:false}" 
                v-model="createData.master_area_id">
              </el-cascader>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="12">
              <el-form-item label="slave_ip:" prop="slave1_ip">
                <el-input v-model="createData.slave1_ip" clearable placeholder="请输入" ></el-input>
              </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="所属机房:" prop="area_id">
              <el-cascader 
                :options="areas"
                :props="{label:'name',value:'area_id',emitPath:false}" 
                v-model="createData.slave1_area_id">
              </el-cascader>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row>
          <el-col :span="12">
            <el-form-item label="slave_ip:" prop="slave2_ip">
              <el-input v-model="createData.slave2_ip" clearable placeholder="请输入" ></el-input>
          </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="所属机房:" prop="area_id">
              <el-cascader 
                :options="areas"
                :props="{label:'name',value:'area_id',emitPath:false}" 
                v-model="createData.slave2_area_id">
              </el-cascader>
            </el-form-item>
          </el-col>
      </el-row>
       </el-form>
      <div class="dialog-footer" slot="footer">
        <el-button @click="closeCreateDialog">取 消</el-button>
        <el-button @click="enterCreateDialog" type="primary">确 定</el-button>
      </div>
    </el-dialog>

    <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" title="更新">
      <el-form :model="formData" :rules="formRules"  label-position="right" label-width="90px">
        <el-form-item label="ip:" prop="ip">
          <el-input v-model="formData.ip" clearable placeholder="请输入IP" ></el-input>
        </el-form-item>
       

        <el-form-item label="角色:" prop="role_id">
          <el-select v-model="formData.role_id" placeholder="请选择" clearable>
            <el-option v-for="(item,key) in db_roleOptions" :key="key" :label="item.label" :value="item.value"></el-option>
          </el-select>
        </el-form-item>    
      <el-form-item label="所属机房:" prop="area_id">
        <el-cascader 
          :options="areas"
          :props="{label:'name',value:'area_id',emitPath:false}" 
          v-model="formData.area_id">
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
    createNode,
    deleteNode,
    deleteNodeByIds,
    updateNode,
    findNode,
    getNodeList,
} from "@/api/xdf_node";  //  此处请自行替换地址
import { formatTimeToStr } from "@/utils/date";
import { getAllTags } from "@/api/xdf_tag";
import { getAllAreas } from "@/api/xdf_area";
import infoList from "@/mixins/infoList";

var Tags = [
  {
    tag_id: 6,
    name: "1111"
  }
];

var Areas = [
  {
    area_id: 0,
    name: "1111"
  }
]

export default {
  name: "Node",
  mixins: [infoList],
  data() {
    return {
      listApi: getNodeList,
      dialogFormVisible: false,
      createDialogFormVisible: false,
      visible: false,
      type: "",
      tags : Tags,
      areas : Areas,
      deleteVisible: false,
      multipleSelection: [],
      db_roleOptions: [],
      createData: {
          cluster_name: "",
          master_ip:"",
          master_area_id: 0,
          slave1_ip:"",
          slave1_area_id: 0,
          slave2_ip:"",
          slave2_area_id: 0,
      },
      formData: {
            node_id:0,
            tag_id:0,
            ip:"",
            role_id:0,
            area_id:0,
            
      },

      formRules: {
        cluster_name: [
          {required: true, message: '请输入集群名称', trigger: 'blur'}
        ],
        master_ip: [
          {required: true, message: '请选主库地址', trigger: 'blur'}
        ],
        area_id: [
          {required: true, message: '请选择机房', trigger: 'blur'}
        ],
        slave1_ip: [
          {required: true, message: '请输入从库地址', trigger: 'blur'}
        ],
        slave2_ip: [
          {required: true, message: '请输入从库地址', trigger: 'blur'}
        ],
        ip: [
          {required: true, message: '请输入ip', trigger: 'blur'}
        ],
        role_id: [
          {required: true, message: '请选择角色', trigger: 'blur'}
        ]

      },
 
      roleTypes: [
        {
          value: 1,
          label: '主'
        },
        {
          value: 2,
          label: '从'
        }
      ],
    };
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
            ids.push(item.node_id)
          })
        const res = await deleteNodeByIds({ ids })
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
    async updateNode(row) {
      const res = await findNode({ node_id: row.node_id });
      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.renode;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
          node_id:0,
          tag_id:0,
          ip:"",
          role_id:0,
          area_id:0,
          
      };
    },

closeCreateDialog() {
      this.createDialogFormVisible = false;
      this.createData = {
          cluster_name: "",
          master_ip:"",
          master_area_id: 0,
          slave1_ip:"",
          slave1_area_id: 0,
          slave2_ip:"",
          slave2_area_id: 0,
      }
    },

    async deleteNode(row) {
      this.visible = false;
      const res = await deleteNode({ node_id: row.node_id });
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
    getTagByID(tag_id) {
       var index = this.tags.findNode( o => o.tag_id === tag_id)
       
       console.log(index)
    },

    async enterCreateDialog() {
        var nodes = [
          {
            ip: "",
            role_id: 0,
            area_id: 0,
          }
        ]
        nodes = []

        if (this.createData.master_ip.length != 0) {
          nodes.push({
            ip: this.createData.master_ip,
            role_id: 1,
            area_id: this.createData.master_area_id,
            })
        }

        if (this.createData.slave1_ip.length != 0) {
          nodes.push({
            ip: this.createData.slave1_ip,
            role_id: 2,
            area_id: this.createData.slave1_area_id,
            })
        }

        if (this.createData.slave2_ip.length != 0) {
          nodes.push({
            ip: this.createData.slave2_ip,
            role_id: 2,
            area_id: this.createData.slave2_area_id,
            })
        }
        let res = await createNode({
          cluster_name:  this.createData.cluster_name,
          nodes: nodes,
          })

        if (res.code == 0) {
          this.$message({
            type:"success",
            message:"创建成功"
          })
          this.closeDialog();
          this.getTableData();
        }
    },

    async enterDialog() {
      let res;
      switch (this.type) {
        case "update":
          res = await updateNode(this.formData);
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
      this.createDialogFormVisible = true;
    }
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
    areaFilter(value) {
      const target = Areas.filter(item => item.area_id === value)[0];
      return target && `${target.name}`;
    }
  },
  async created() {
    
    var rst = await getAllTags();
    if (rst.code == 0) {
      this.tags = rst.data.list
      Tags = this.tags
    }
    var rst2 = await getAllAreas();
    if (rst2.code == 0) {
      this.areas = rst2.data.list
      Areas = this.areas
    }
    await this.getDict("db_role");
    await this.getTableData();
}
};
</script>

<style>
</style>
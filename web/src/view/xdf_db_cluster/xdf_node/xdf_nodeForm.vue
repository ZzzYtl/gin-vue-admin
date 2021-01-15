<template>
<el_container>
  <task-sidebar></task-sidebar>
    <el-main>
    <el-form ref="form" :model="formData" label-position="right" label-width="80px">
             <el-form-item label="ID:"><el-input v-model.number="formData.node_id" clearable placeholder="请输入"></el-input>
          </el-form-item>
           
             <el-form-item label="集群名称:"><el-input v-model.number="formData.tag_id" clearable placeholder="请输入"></el-input>
          </el-form-item>
           
             <el-form-item label="IP:">
                <el-input v-model="formData.ip" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="角色:"><el-input v-model.number="formData.role_id" clearable placeholder="请输入"></el-input>
          </el-form-item>
           
             <el-form-item label="机房:"><el-input v-model.number="formData.area_id" clearable placeholder="请输入"></el-input>
          </el-form-item>
           <el-form-item>
           <el-button @click="save" type="primary">保存</el-button>
           <el-button @click="back" type="primary">返回</el-button>
           </el-form-item>
    </el-form>
    </el-main>
</el_container>
</template>

<script>
import {
    createNode,
    updateNode,
    findNode
} from "@/api/xdf_node";  //  此处请自行替换地址
import infoList from "@/mixins/infoList";
export default {
  name: "Node",
  mixins: [infoList],
  data() {
    return {
      type: "",formData: {
            node_id:0,
            tag_id:0,
            ip:"",
            role_id:0,
            area_id:0,
            
      }
    };
  },
  methods: {
    async save() {
      let res;
      switch (this.type) {
        case "create":
          res = await createNode(this.formData);
          break;
        case "update":
          res = await updateNode(this.formData);
          break;
        default:
          res = await createNode(this.formData);
          break;
      }
      if (res.code == 0) {
        this.$message({
          type:"success",
          message:"创建/更改成功"
        })
      }
    },
    back(){
        this.$router.go(-1)
    }
  },
  async created() {
   // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if(this.$route.query.id){
    const res = await findNode({ ID: this.$route.query.id })
    if(res.code == 0){
       this.formData = res.data.renode
       this.type == "update"
     }
    }else{
     this.type == "create"
   }
  
}
};
</script>

<style>
</style>
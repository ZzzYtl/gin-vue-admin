<template>
<div>
    <el-form :model="formData" label-position="right" label-width="80px">
             <el-form-item label="ID:"><el-input v-model.number="formData.cluster_id" clearable placeholder="请输入"></el-input>
          </el-form-item>
           
             <el-form-item label="集群名:"><el-input v-model.number="formData.tag_id" clearable placeholder="请输入"></el-input>
          </el-form-item>
           
             <el-form-item label="实例名:">
                <el-input v-model="formData.cluster_name" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="数据库类型:">
                 <el-select v-model="formData.db_type" placeholder="请选择" clearable>
                     <el-option v-for="(item,key) in db_typeOptions" :key="key" :label="item.label" :value="item.value"></el-option>
                 </el-select>
          </el-form-item>
           
             <el-form-item label="业务线:"><el-input v-model.number="formData.dept_id" clearable placeholder="请输入"></el-input>
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
           
             <el-form-item label="备份库:"><el-input v-model.number="formData.backup_id" clearable placeholder="请输入"></el-input>
          </el-form-item>
           
             <el-form-item label="延迟备份库:"><el-input v-model.number="formData.delay_backup_id" clearable placeholder="请输入"></el-input>
          </el-form-item>
           <el-form-item>
           <el-button @click="save" type="primary">保存</el-button>
           <el-button @click="back" type="primary">返回</el-button>
           </el-form-item>
    </el-form>
</div>
</template>

<script>
import {
    createDataBaseInfo,
    updateDataBaseInfo,
    findDataBaseInfo
} from "@/api/xdf_db_info";  //  此处请自行替换地址
import infoList from "@/mixins/infoList";
export default {
  name: "DataBaseInfo",
  mixins: [infoList],
  data() {
    return {
      type: "",
      db_typeOptions:[],
          formData: {
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
            
      }
    };
  },
  methods: {
    async save() {
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
      }
    },
    back(){
        this.$router.go(-1)
    }
  },
  async created() {
   // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if(this.$route.query.id){
    const res = await findDataBaseInfo({ ID: this.$route.query.id })
    if(res.code == 0){
       this.formData = res.data.reDBInfo
       this.type == "update"
     }
    }else{
     this.type == "create"
   }
  
    await this.getDict("db_type");
    
}
};
</script>

<style>
</style>
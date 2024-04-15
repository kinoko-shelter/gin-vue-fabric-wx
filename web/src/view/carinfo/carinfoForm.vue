<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="车辆id:" prop="car_id">
          <el-input v-model.number="formData.car_id" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="速度:" prop="speed">
          <el-input v-model="formData.speed" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="车辆日前预测里程:" prop="mileagePre">
          <el-input v-model="formData.mileagePre" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="车辆位置x:" prop="carLocationx">
          <el-input v-model="formData.carLocationx" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="车辆位置y:" prop="carLocationy">
          <el-input v-model="formData.carLocationy" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="电量:" prop="soc">
          <el-input v-model="formData.soc" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="累计行驶里程:" prop="totalMileage">
          <el-input v-model="formData.totalMileage" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="创建者:" prop="createdBy">
          <el-input v-model.number="formData.createdBy" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="更新者:" prop="updatedBy">
          <el-input v-model.number="formData.updatedBy" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="删除者:" prop="deletedBy">
          <el-input v-model.number="formData.deletedBy" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="车辆类型:" prop="type">
          <el-input v-model="formData.type" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="当前状态:" prop="currentState">
          <el-input v-model="formData.currentState" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createCarinfo,
  updateCarinfo,
  findCarinfo
} from '@/api/carinfo'

defineOptions({
    name: 'CarinfoForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            car_id: 0,
            speed: '',
            mileagePre: '',
            carLocationx: '',
            carLocationy: '',
            soc: '',
            totalMileage: '',
            createdBy: 0,
            updatedBy: 0,
            deletedBy: 0,
            type: '',
            currentState: '',
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findCarinfo({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.recarinfo
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createCarinfo(formData.value)
               break
             case 'update':
               res = await updateCarinfo(formData.value)
               break
             default:
               res = await createCarinfo(formData.value)
               break
           }
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>

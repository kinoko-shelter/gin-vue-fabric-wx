<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="车辆id:" prop="car_Id">
          <el-input v-model.number="formData.car_Id" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="速度:" prop="speed">
          <el-input v-model="formData.speed" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="车辆日前预测里程:" prop="mileage_pre">
          <el-input v-model="formData.mileage_pre" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="车辆位置x:" prop="car_locationx">
          <el-input v-model="formData.car_locationx" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="车辆位置y:" prop="car_locationy">
          <el-input v-model="formData.car_locationy" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="电量:" prop="soc">
          <el-input v-model="formData.soc" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="累计行驶里程:" prop="total_mileage">
          <el-input v-model="formData.total_mileage" :clearable="true" placeholder="请输入" />
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
  createVehiclesinfo,
  updateVehiclesinfo,
  findVehiclesinfo
} from '@/api/vehiclesinfo'

defineOptions({
    name: 'VehiclesinfoForm'
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
            car_Id: 0,
            speed: '',
            mileage_pre: '',
            car_locationx: '',
            car_locationy: '',
            soc: '',
            total_mileage: '',
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findVehiclesinfo({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reVehiclesinfoCar
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
               res = await createVehiclesinfo(formData.value)
               break
             case 'update':
               res = await updateVehiclesinfo(formData.value)
               break
             default:
               res = await createVehiclesinfo(formData.value)
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

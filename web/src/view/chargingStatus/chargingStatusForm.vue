<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="用户id:" prop="user_id">
          <el-input v-model="formData.user_id" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="车辆id:" prop="car_id">
          <el-input v-model="formData.car_id" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="充电站id:" prop="station_Id">
          <el-input v-model="formData.station_Id" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="到达状态:" prop="arrival_status">
          <el-input v-model.number="formData.arrival_status" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="充电状态:" prop="charging_status">
          <el-input v-model.number="formData.charging_status" :clearable="true" placeholder="请输入" />
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
  createChargingStatus,
  updateChargingStatus,
  findChargingStatus
} from '@/api/chargingStatus'

defineOptions({
    name: 'ChargingStatusForm'
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
            user_id: '',
            car_id: '',
            station_Id: '',
            arrival_status: 0,
            charging_status: 0,
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findChargingStatus({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reChargingStatusInfo
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
               res = await createChargingStatus(formData.value)
               break
             case 'update':
               res = await updateChargingStatus(formData.value)
               break
             default:
               res = await createChargingStatus(formData.value)
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

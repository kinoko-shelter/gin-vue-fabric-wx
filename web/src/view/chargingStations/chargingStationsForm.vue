<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="充电站id:" prop="station_Id">
          <el-input v-model.number="formData.station_Id" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="充电桩个数:" prop="pile_num">
          <el-input v-model.number="formData.pile_num" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="排队车辆数:" prop="wait_num">
          <el-input v-model.number="formData.wait_num" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="充电车辆数:" prop="charge_num">
          <el-input v-model="formData.charge_num" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="累计服务车次:" prop="total_service">
          <el-input v-model="formData.total_service" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="充电站位置x:" prop="station_locationx">
          <el-input v-model="formData.station_locationx" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="充电站位置y:" prop="station_locationy">
          <el-input v-model="formData.station_locationy" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="充电桩功率:" prop="station_name">
          <el-input v-model="formData.station_name" :clearable="true" placeholder="请输入" />
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
  createChargingStations,
  updateChargingStations,
  findChargingStations
} from '@/api/chargingStations'

defineOptions({
    name: 'ChargingStationsForm'
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
            station_Id: 0,
            pile_num: 0,
            wait_num: 0,
            charge_num: '',
            total_service: '',
            station_locationx: '',
            station_locationy: '',
            station_name: '',
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findChargingStations({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reChargingStationsInfo
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
               res = await createChargingStations(formData.value)
               break
             case 'update':
               res = await updateChargingStations(formData.value)
               break
             default:
               res = await createChargingStations(formData.value)
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

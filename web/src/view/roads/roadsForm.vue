<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="道路id:" prop="road_id">
          <el-input v-model.number="formData.road_id" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="统计时间:" prop="road_len">
          <el-input v-model="formData.road_len" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="道路限速:" prop="road_limit">
          <el-input v-model="formData.road_limit" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="路段起点x:" prop="r_startx">
          <el-input v-model="formData.r_startx" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="路段起点y:" prop="r_starty">
          <el-input v-model="formData.r_starty" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="路段终点x:" prop="r_endx">
          <el-input v-model="formData.r_endx" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="路段终点y:" prop="r_endy">
          <el-input v-model="formData.r_endy" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="车辆数:" prop="r_JamIndex">
          <el-input v-model="formData.r_JamIndex" :clearable="true" placeholder="请输入" />
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
  createRoads,
  updateRoads,
  findRoads
} from '@/api/roads'

defineOptions({
    name: 'RoadsForm'
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
            road_id: 0,
            road_len: '',
            road_limit: '',
            r_startx: '',
            r_starty: '',
            r_endx: '',
            r_endy: '',
            r_JamIndex: '',
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findRoads({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reRoadsInfo
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
               res = await createRoads(formData.value)
               break
             case 'update':
               res = await updateRoads(formData.value)
               break
             default:
               res = await createRoads(formData.value)
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

<template>
  <div class="gva-table-box">

    <section>
      <el-input
        class="query-ipt"
        placeholder="位置"
        clearable
      />
      <el-button
        type="primary"
        @click="startAnimation"
      >开始</el-button>
    <div id="map-location" class="map-location"></div>
  </section>
  <div class="gva-table-box">
    <div class="gva-search-box">
        <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" :rules="searchRule" @keyup.enter="onSubmit">
        <el-form-item label="创建日期" prop="createdAt">
        <template #label>
          <span>
            创建日期
            <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
              <el-icon><QuestionFilled /></el-icon>
            </el-tooltip>
          </span>
        </template>
        <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始日期" :disabled-date="time=> searchInfo.endCreatedAt ? time.getTime() > searchInfo.endCreatedAt.getTime() : false"></el-date-picker>
         —
        <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束日期" :disabled-date="time=> searchInfo.startCreatedAt ? time.getTime() < searchInfo.startCreatedAt.getTime() : false"></el-date-picker>
        </el-form-item>
        
          <el-form-item label="位置x" prop="location_x">
           <el-input v-model="searchInfo.location_x" placeholder="搜索条件" />
  
          </el-form-item>
          <el-form-item label="位置y" prop="location_y">
           <el-input v-model="searchInfo.location_y" placeholder="搜索条件" />
  
          </el-form-item>
          <el-form-item>
            <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
            <el-button icon="refresh" @click="onReset">重置</el-button>
          </el-form-item>
        </el-form>
      </div>
      <div class="gva-table-box">
          <div class="gva-btn-list">
              <el-button type="primary" icon="plus" @click="openDialog">新增</el-button>
              <el-popover v-model:visible="deleteVisible" :disabled="!multipleSelection.length" placement="top" width="160">
              <p>确定要删除吗？</p>
              <div style="text-align: right; margin-top: 8px;">
                  <el-button type="primary" link @click="deleteVisible = false">取消</el-button>
                  <el-button type="primary" @click="onDelete">确定</el-button>
              </div>
              <template #reference>
                  <el-button icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="deleteVisible = true">删除</el-button>
              </template>
              </el-popover>
          </div>
          <el-table
          ref="multipleTable"
          style="width: 100%"
          tooltip-effect="dark"
          :data="tableData"
          row-key="ID"
          @selection-change="handleSelectionChange"
          >
          <el-table-column type="selection" width="55" />
          
          <el-table-column align="left" label="日期" width="180">
              <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
          </el-table-column>
          
          <el-table-column align="left" label="车辆名称" prop="name" width="120" />
          <el-table-column align="left" label="纬度" prop="location_y" width="120" />
          <el-table-column align="left" label="经度" prop="location_x" width="120" />
          <el-table-column align="left" label="路径" prop="paths" width="500" />
          <el-table-column align="left" label="操作" fixed="right" min-width="240">
              <template #default="scope">
              <el-button type="primary" link class="table-button" @click="getDetails(scope.row)">
                  <el-icon style="margin-right: 5px"><InfoFilled /></el-icon>
                  查看详情
              </el-button>
              <el-button type="primary" link icon="edit" class="table-button" @click="updateCmapinfoFunc(scope.row)">变更</el-button>
              <el-button type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
              </template>
          </el-table-column>
          </el-table>
          <div class="gva-pagination">
              <el-pagination
              layout="total, sizes, prev, pager, next, jumper"
              :current-page="page"
              :page-size="pageSize"
              :page-sizes="[10, 30, 50, 100]"
              :total="total"
              @current-change="handleCurrentChange"
              @size-change="handleSizeChange"
              />
          </div>
      </div>
      <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" :title="type==='create'?'添加':'修改'" destroy-on-close>
        <el-scrollbar height="500px">
            <el-form :model="formData" label-position="right" ref="elFormRef" :rules="rule" label-width="80px">
              <el-form-item label="位置x:"  prop="location_x" >
                <el-input v-model="formData.location_x" :clearable="true"  placeholder="请输入位置x" />
              </el-form-item>
              <el-form-item label="位置y:"  prop="location_y" >
                <el-input v-model="formData.location_y" :clearable="true"  placeholder="请输入位置y" />
              </el-form-item>
            </el-form>
        </el-scrollbar>
        <template #footer>
          <div class="dialog-footer">
            <el-button @click="closeDialog">取 消</el-button>
            <el-button type="primary" @click="enterDialog">确 定</el-button>
          </div>
        </template>
      </el-dialog>
      <el-dialog v-model="detailShow" style="width: 800px" lock-scroll :before-close="closeDetailShow" title="查看详情" destroy-on-close>
        <el-scrollbar height="550px">
          <el-descriptions :column="1" border>
                  <el-descriptions-item label="车辆名称">
                          {{ formData.name }}
                  </el-descriptions-item>
                  <el-descriptions-item label="起始位置x">
                          {{ formData.location_x }}
                  </el-descriptions-item>
                  <el-descriptions-item label="起始位置y">
                          {{ formData.location_y }}
                  </el-descriptions-item>
                  <el-descriptions-item label="路径">
                          {{ formData.paths }}
                  </el-descriptions-item>
          </el-descriptions>
        </el-scrollbar>
      </el-dialog>
</div>

  </div>

</template>

<script setup>
import { onMounted, ref,reactive} from 'vue';
import AMapLoader from '@amap/amap-jsapi-loader';
import {
    createCmapinfo,
    deleteCmapinfo,
    deleteCmapinfoByIds,
    updateCmapinfo,
    findCmapinfo,
    getCmapinfoList
  } from '@/api/cmapinfo'
  defineOptions({
      name: 'Cmapinfo'
  })
    import { getDictFunc, formatDate, formatBoolean, filterDict, ReturnArrImg, onDownloadFile } from '@/utils/format'
  import { ElMessage, ElMessageBox } from 'element-plus'
let map = ref(null);
let marker = ref(null);
let polyline = ref(null);
let passedPolyline = ref(null);
const elFormRef = ref()
  const elSearchFormRef = ref()
  
  // =========== 表格控制部分 ===========
  const page = ref(1)
  const total = ref(0)
  const pageSize = ref(10)
  const tableData = ref([])
  const searchInfo = ref({})
  
  // 重置
  const onReset = () => {
    searchInfo.value = {}
    getTableData()
  }


const cars = ref([]);


async function loadMapAndCars() {
  const data = await getCmapinfoList({}); 
  console.log("data",data); // 假设这个函数返回所有需要的车辆数据
  if (data.code === 0) {
    const loadedCars = data.data.list.map(car => {
      const parsedPaths = JSON.parse(car.paths.replace(/\\/g, '')); // 确保去除额外的转义字符
      const carData = {
        position: [parseFloat(car.location_x), parseFloat(car.location_y)],
        path: parsedPaths.map(point => ([parseFloat(point.lng), parseFloat(point.lat)]))
      };
      console.log("Car Position", carData.position);
      console.log("Car Path", carData.path);
      return carData
      
    });
    cars.value = loadedCars;
    await renderMap();
  }
}

onMounted(async () => {
  await loadMapAndCars();
});
function getRandomColor(light) {
  // light参数决定颜色的浅重，light为true生成浅色调，否则生成重色调
  const opacity = light ? '0.5' : '1'; // 透明度，浅色更透明
  let color = 'rgba(';
  for (let i = 0; i < 3; i++) {
    color += Math.floor(Math.random() * 256) + ',';
  }
  color += opacity + ')';
  return color;
}


async function renderMap() {
  await AMapLoader.load({
    key: "8585680ecdc9060ba9b7f3e86eaf2a52",
    version: "2.0",
    plugins: ['AMap.MoveAnimation']
  }).then(AMap => {
    const map = new AMap.Map("map-location", {
      resizeEnable: true,
      center: [108.9878, 34.2325],  // 初始地图中心点
      zoom: 50
    });

    cars.value.forEach((car, index) => {
      let marker = new AMap.Marker({
        map: map,
        position: car.position,
        icon: "https://a.amap.com/jsapi_demos/static/demo-center-v2/car.png",
        offset: new AMap.Pixel(-26, -13),
      });

      let polyline = new AMap.Polyline({
        map: map,
        path: car.path,
        showDir: true,
        strokeColor:  getRandomColor(true),  // 循环使用三种颜色
        strokeWeight: 6,
      });

      let passedPolyline = new AMap.Polyline({
        map: map,
        strokeColor:  getRandomColor(false),
        strokeWeight: 10,
      });

      marker.on('moving', (e) => {
        passedPolyline.setPath(e.passedPath);
      });

      setTimeout(() => {
        marker.moveAlong(car.path, {
          duration: 3000,
          autoRotation: true,
          speed: 1
        });
      }, index * 1000);
    });

    map.setFitView();
  });
}


function pauseAnimation() {
  marker.value.pauseMove();
}

function resumeAnimation() {
  marker.value.resumeMove();
}

function stopAnimation() {
  marker.value.stopMove();
}


const onSubmit = () => {
    elSearchFormRef.value?.validate(async(valid) => {
      if (!valid) return
      page.value = 1
      pageSize.value = 10
      getTableData()
    })
  }
  
  // 分页
  const handleSizeChange = (val) => {
    pageSize.value = val
    getTableData()
  }
  
  // 修改页面容量
  const handleCurrentChange = (val) => {
    page.value = val
    getTableData()
  }
  
  // 查询
  const getTableData = async() => {
    const table = await getCmapinfoList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
    if (table.code === 0) {
      tableData.value = table.data.list
      total.value = table.data.total
      page.value = table.data.page
      pageSize.value = table.data.pageSize
    }
  }
  
  getTableData()
   // ============== 表格控制部分结束 ===============
  
  // 获取需要的字典 可能为空 按需保留
  const setOptions = async () =>{
  }
  
  // 获取需要的字典 可能为空 按需保留
  setOptions()
  
  
  // 多选数据
  const multipleSelection = ref([])
  // 多选
  const handleSelectionChange = (val) => {
      multipleSelection.value = val
  }
  
  // 删除行
  const deleteRow = (row) => {
      ElMessageBox.confirm('确定要删除吗?', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
      }).then(() => {
              deleteCmapinfoFunc(row)
          })
      }
  
  
  // 批量删除控制标记
  const deleteVisible = ref(false)
  
  // 多选删除
  const onDelete = async() => {
        const IDs = []
        if (multipleSelection.value.length === 0) {
          ElMessage({
            type: 'warning',
            message: '请选择要删除的数据'
          })
          return
        }
        multipleSelection.value &&
          multipleSelection.value.map(item => {
            IDs.push(item.ID)
          })
        const res = await deleteCmapinfoByIds({ IDs })
        if (res.code === 0) {
          ElMessage({
            type: 'success',
            message: '删除成功'
          })
          if (tableData.value.length === IDs.length && page.value > 1) {
            page.value--
          }
          deleteVisible.value = false
          getTableData()
        }
      }
  
  // 行为控制标记（弹窗内部需要增还是改）
  const type = ref('')
  
  // 更新行
  const updateCmapinfoFunc = async(row) => {
      const res = await findCmapinfo({ ID: row.ID })
      type.value = 'update'
      if (res.code === 0) {
          formData.value = res.data.recmapinfo
          dialogFormVisible.value = true
      }
  }
  
  
  // 删除行
  const deleteCmapinfoFunc = async (row) => {
      const res = await deleteCmapinfo({ ID: row.ID })
      if (res.code === 0) {
          ElMessage({
                  type: 'success',
                  message: '删除成功'
              })
              if (tableData.value.length === 1 && page.value > 1) {
              page.value--
          }
          getTableData()
      }
  }
  
  // 弹窗控制标记
  const dialogFormVisible = ref(false)
  
  
  // 查看详情控制标记
  const detailShow = ref(false)
  
  
  // 打开详情弹窗
  const openDetailShow = () => {
    detailShow.value = true
  }
  
  
  // 打开详情
  const getDetails = async (row) => {
    // 打开弹窗
    const res = await findCmapinfo({ ID: row.ID })
    if (res.code === 0) {
      formData.value = res.data.recmapinfo
      openDetailShow()
    }
  }
  
  
  // 关闭详情弹窗
  const closeDetailShow = () => {
    detailShow.value = false
    formData.value = {
            location_x: '',
            location_y: '',
            }
  }
  
  
  // 打开弹窗
  const openDialog = () => {
      type.value = 'create'
      dialogFormVisible.value = true
  }
  
  // 关闭弹窗
  const closeDialog = () => {
      dialogFormVisible.value = false
      formData.value = {
          location_x: '',
          location_y: '',
          }
  }
  // 弹窗确定
  const enterDialog = async () => {
       elFormRef.value?.validate( async (valid) => {
               if (!valid) return
                let res
                switch (type.value) {
                  case 'create':
                    res = await createCmapinfo(formData.value)
                    break
                  case 'update':
                    res = await updateCmapinfo(formData.value)
                    break
                  default:
                    res = await createCmapinfo(formData.value)
                    break
                }
                if (res.code === 0) {
                  ElMessage({
                    type: 'success',
                    message: '创建/更改成功'
                  })
                  closeDialog()
                  getTableData()
                }
        })
  }
</script>

<style scoped lang="scss">
.input-card .btn {
  margin-right: 1.2rem;
  width: 9rem;
}
.map-location {
  height: 60vh;
  border-radius: 10px;
  padding: 10px;
  margin-top: 30px;
}
.input-card .btn:last-child {
  margin-right: 0;
}
.secret{
  padding: 30px;
  margin-top: 20px;
  background: #F5F5F5;
    p {
    line-height: 30px;
  }
}
.query-ipt{
  width: 300px;
  margin-right: 30px;
}
.content{
  p {
    font-size: 16px;
    line-height: 20px;
  }
  padding: 10px;
  width: 100%;
  background: #F5F5F5;
  margin-top: 30px;
}
.tables{
  width: 100%;
  margin-top: 30px;
  .text{
    font-size: 18px;
    color: #6B7687;
    text-align: center;
  }
}
</style>

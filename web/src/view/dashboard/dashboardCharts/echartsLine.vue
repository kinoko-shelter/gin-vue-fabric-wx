<template>
  <div class="dashboard-line-box">
    <div class="dashboard-line-title">
      当日车流量情况
    </div>
    <div
      ref="echart"
      class="dashboard-line"
    />
  </div>
</template>
<script setup>
import * as echarts from 'echarts'
import { nextTick, onMounted, onUnmounted, ref } from 'vue'
import { useWindowResize } from '@/hooks/use-windows-resize'
var dataAxis = [];
for (var i = 0; i < 24; i++) {
  dataAxis.push(`${i}时`);
}

var data = [];
for (var i = 0; i < 24; i++) {
  if (i >= 7 && i <= 9 || i >= 17 && i <= 19) {
    data.push(Math.floor(Math.random() * (1001 - 800)) + 800);
  } else if (i >= 10 && i <= 16) {
    data.push(Math.floor(Math.random() * (701 - 500)) + 500);
  } else if (i >= 20 && i <= 23) {
    data.push(Math.floor(Math.random() * (401 - 600)) + 500);
  } else {
    data.push(Math.floor(Math.random() * (501 - 300)) + 300);
  }
}


var yMax = 1000;
var dataShadow = [];

for (let i = 0; i < data.length; i++) {
  dataShadow.push(yMax)
}

let chart = null
const echart = ref(null)

useWindowResize(() => {
  if (!chart) {
    return
  }
  chart.resize()
})

const initChart = () => {
  if (chart) {
    chart = null
  }
  chart = echarts.init(echart.value)
  setOptions()
}
const setOptions = () => {
  chart.setOption({
    grid: {
      left: '40',
      right: '20',
      top: '40',
      bottom: '20',
    },
    xAxis: {
      data: dataAxis,
      axisTick: {
        show: false,
      },
      axisLine: {
        show: false,
      },
      z: 10,
    },
    yAxis: {
      axisLine: {
        show: false,
      },
      axisTick: {
        show: false,
      },
      axisLabel: {
        textStyle: {
          color: '#999',
        },
      },
    },
    dataZoom: [
      {
        type: 'inside',
      },
    ],
    series: [
      {
        type: 'bar',
        barWidth: '40%',
        itemStyle: {
          borderRadius: [5, 5, 0, 0],
          color: '#188df0',
        },
        emphasis: {
          itemStyle: {
            color: '#188df0',
          },
        },
        data: data,
      },
    ],
  })
}

onMounted(async() => {
  await nextTick()
  initChart()
})

onUnmounted(() => {
  if (!chart) {
    return
  }
  chart.dispose()
  chart = null
})
</script>
<style lang="scss" scoped>
.dashboard-line-box {
  .dashboard-line {
    background-color: #fff;
    height: 360px;
    width: 100%;
  }
  .dashboard-line-title {
    font-weight: 1000;
    font-size: x-large;
    margin-bottom: 12px;
  }
}
</style>

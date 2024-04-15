<template>
  <div>
    <div class="data-center-box">
      <CenterCard title="充电情况">
        <template #action>
          <span class="gvaIcon-prompt" style="color: #999" />
        </template>
        <template #body>
          <ChainRatio />
        </template>
      </CenterCard>
      <CenterCard title="月度电量消耗趋势">
        <template #action>
          <span class="gvaIcon-prompt" style="color: #999" />
        </template>
        <template #body>
          <ReclaimMileage />
        </template>
      </CenterCard>
      <CenterCard title="月度充电量趋势">
        <template #action>
          <span class="gvaIcon-prompt" style="color: #999" />
        </template>
        <template #body>
          <Order />
        </template>
      </CenterCard>
      <CenterCard title="交易成功率">
        <template #action>
          <span class="gvaIcon-prompt" style="color: #999" />
        </template>
        <template #body>
          <RecoveryRate />
        </template>
      </CenterCard>



      <CenterCard title="年度充电量趋势" style="grid-column-start: span 3;">
        <template #action>
          <span class="gvaIcon-prompt" style="color: #999" />
        </template>
        <template #body>
          <lineCharts />
        </template>
      </CenterCard>
      <CenterCard title="月车型充电次数分布" style="grid-auto-columns: 1.5fr">
        <template #action>
          <span class="gvaIcon-prompt" style="color: #999" />
        </template>
        <template #body>
          <part />
        </template>
      </CenterCard>
    </div>
  </div>
</template>

<script setup>
import CenterCard from './dataCenterComponents/centerCard.vue'
import ChainRatio from './dataCenterComponents/chainRatio.vue'
import ReclaimMileage from './dataCenterComponents/ReclaimMileage.vue'
import RecoveryRate from './dataCenterComponents/RecoveryRate.vue'
import Order from './dataCenterComponents/order.vue'
import lineCharts from './dataCenterComponents/lineCharts.vue'
import part from './dataCenterComponents/part.vue'

</script>

<script>
export default {
  name: 'DataCenter',
}
</script>

<style lang="scss" scoped>

.data-center-box{
  width: 100%;
  display: grid;
  grid-template-columns: 1fr 1.5fr 2fr 1.5fr;
  column-gap: 10px;
}
</style>

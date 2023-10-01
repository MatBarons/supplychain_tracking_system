<template>
<main id="home">
  <Header pageName="Homepage" @search_data="handleSearch" @stock_data="handleReceivedData" @all_data="handleAllData"/>
  <div v-if="searchedStocks.length === 0">
    <HomepageData :stocks="homepageStocks" :key="key_2"/>
  </div>
  <div v-else>
    <HomepageSearchData :stocks="searchedStocks" :key="key_1"/>
  </div>
</main>
</template>

<script lang="ts" setup>
import Header from '@/components/Header.vue';
import HomepageData from '@/components/HomepageData.vue';
import HomepageSearchData from '@/components/HomepageSearchData.vue';
import type { Stock } from '@/types/stock';
import { onMounted, ref, type Ref } from 'vue';
import { useDataStore } from '@/stores/store';
import { getStocksSocket } from '@/api_calls/socket';

const store = useDataStore()
let homepageStocks : Ref<Stock[]> = ref([])
let searchedStocks : Ref<Stock[]> = ref([])
let key_1: number = 0
let key_2 : number = 0

function handleAllData(data: Stock[],search: Stock[]){
  handleSearch(search)
  handleReceivedData(data)

}

function handleReceivedData(data: Stock[]){
  homepageStocks.value = data
  key_2++;
}
function handleSearch(data : Stock[]){
  searchedStocks.value = data
  key_1++;
}

onMounted( () => {
  if(store.data.wallet !== ""){
    console.log("prv")
    getStocksSocket().then(response => {
      handleReceivedData(response)
    })
  }
})

</script>

<style lang="scss" scoped>
#home{
  min-height: 100%;
  width: 100%;
  display: flex;
  flex-direction: column;
  overflow-x: hidden;
  
}
</style>
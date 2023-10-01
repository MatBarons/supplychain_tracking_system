<template>
    <div class="filters">
        <button class="all" @click="changeShowedList(0)">
          <h3>All</h3>
          <h3>{{ n_total }}</h3>
        </button>
        <button class="stocks" @click="changeShowedList(1)">
          <h3>My Stocks</h3>
          <h3>{{ n_stocks }}</h3>
        </button>
        <button class="my-requests" @click="changeShowedList(2)">
          <h3>My Requests</h3>
          <h3>{{ n_myrequests }}</h3>
        </button>
        <button class="others-requests" @click="changeShowedList(3)">
          <h3>Others Requests</h3>
          <h3>{{ n_othersrequests }}</h3>
        </button>
      </div>
      <div class="sortButtons">
         <button class="sortID" @click="sortStocks('id')">
            <h2>ID</h2>
            <span class="material-icons">{{ icon[0] }}</span>
          </button>
          <button class="sortProducer" @click="sortStocks('producer')">
            <h2>Producer</h2>
            <span class="material-icons">{{ icon[1] }}</span>
          </button>
          <button class="sortOwner" @click="sortStocks('owner')">
            <h2>Owner</h2>
            <span class="material-icons">{{ icon[2] }}</span>
          </button>
          <button class="sortStatus" @click="sortStocks('status')">
            <h2>Status</h2>
            <span class="material-icons">{{ icon[3] }}</span>
          </button>
          <button class="sortRequester" @click="sortStocks('requester')">
            <h2>Requester</h2>
            <span class="material-icons">{{ icon[4] }}</span>
          </button>
          <form class="search"> 
            <input class="searchInput" v-model="searchBarInput" type="text" id="query" name="q" placeholder="Search..."  @input="searchMyStocks">
            <span class="material-icons">search</span>
          </form>
        </div>
      <div class="allStocks">
        <div class="stocks" v-for="stock in showedStocksList">
          <StockHomepage 
            :stock=stock :odd="isElementOdd(showedStocksList, stock)" :color="createStockStyle(stock)" :key="key_stock" 
            @updated_stock_list="updateStocksList"/>
        </div>
      </div>
</template>

<script lang="ts" setup>
import StockHomepage from '@/components/StockHomepage.vue';
import { Stock, Status } from "@/types/stock";
import { onMounted, ref } from "vue";
import { compareProducer, compareStatus, compareUuid, compareRequester, compareOwner } from "@/utils/compare"

const props = defineProps({
  stocks: {
    type: Array<Stock>,
    required: true,
  },
  key:{
    type: Number,
    required: true
  }
})
//data
let showedStocksList = ref<Stock[]>([])
let savedStocksList = ref<Stock[]>([])
let searchBarInput = ref("");

let icon = ref(["arrow_drop_up", "arrow_drop_up", "arrow_drop_up", "arrow_drop_up","arrow_drop_up"])
let whichSort = ref("")
let key_stock = ref(0)
let isSelected = ref([true, false, false, false,false])

let n_total = ref(savedStocksList.value.length)
let n_stocks = ref(savedStocksList.value.filter((stock) => stock.status === Status.owned).length)
let n_myrequests = ref(savedStocksList.value.filter((stock) => stock.status === Status.requested).length)
let n_othersrequests = ref(savedStocksList.value.filter((stock) => stock.status === Status.requested_by).length)

//functions
function selectFilterColor(n: number) {
    return isSelected.value[n] ? "#3b5998" : "grey"
}

function searchMyStocks() {
    showedStocksList.value = showedStocksList.value.filter((stock) => {
        if (searchBarInput.value === "") {
            return savedStocksList
        }
        return stock.id.toString().includes(searchBarInput.value) ||
            stock.producer.includes(searchBarInput.value) ||
            stock.status.includes(searchBarInput.value) ||
            stock.request?.requester?.includes(searchBarInput.value)
    })
    updateNumbersInFilters()
    key_stock.value++;
}

function updateNumbersInFilters() {
    n_total = ref(showedStocksList.value.length)
    n_stocks = ref(showedStocksList.value.filter((stock) => stock.status === Status.owned).length)
    n_myrequests = ref(showedStocksList.value.filter((stock) => stock.status === Status.requested).length)
    n_othersrequests = ref(showedStocksList.value.filter((stock) => stock.status === Status.requested_by).length)
}

function sortStocks(whichCliked: string) {
    if (whichCliked === whichSort.value) {
        if (whichCliked === "id") {
            changeArrowIcon(0)
            showedStocksList.value.reverse()
        } else if (whichCliked === "producer") {
            changeArrowIcon(1)
            showedStocksList.value.reverse()
        } else if (whichCliked === "status") {
            changeArrowIcon(2)
            showedStocksList.value.reverse()
        } else if (whichCliked === "owner") {
            changeArrowIcon(3)
            showedStocksList.value.reverse()
        }else{
          changeArrowIcon(4)
          let stocks1: Stock[];
          let stock2: Stock[];
          stocks1 = showedStocksList.value.filter((stock) => stock.request?.requester !== undefined);
          stocks1.reverse();
          stock2 = showedStocksList.value.filter((stock) => stock.request?.requester === undefined);
          showedStocksList.value = stocks1.concat(stock2);
        }
    } else {
        if (whichCliked === "id") {
            showedStocksList.value.sort(compareUuid);
            showedStocksList.value.reverse();
            changeArrowIcon(0)
        } else if (whichCliked === "producer") {
            showedStocksList.value.sort(compareProducer);
            showedStocksList.value.reverse();
            changeArrowIcon(1)
        } else if (whichCliked === "status") {
            showedStocksList.value.sort(compareStatus);
            showedStocksList.value.reverse();
            changeArrowIcon(2)
        } else if(whichCliked === "owner"){
            showedStocksList.value.sort(compareOwner);
            changeArrowIcon(3)
        } else {
            showedStocksList.value.sort(compareRequester)
            changeArrowIcon(4)
        }
        whichSort.value = whichCliked;
    }

    key_stock.value++;
}

function isElementOdd(stocks: Stock[], stock: Stock): boolean {
    return stocks.indexOf(stock) % 2 === 0;
}

function changeShowedList(whichClicked: number) {
    switch (whichClicked) {
        case 0: {
            showedStocksList.value = savedStocksList.value
            isSelected.value[0] = true
            isSelected.value[1] = false
            isSelected.value[2] = false
            isSelected.value[3] = false
        }
        break;
        case 1: {
            showedStocksList.value = savedStocksList.value.filter((stock) => stock.status === Status.owned)
            isSelected.value[0] = false
            isSelected.value[1] = true
            isSelected.value[2] = false
            isSelected.value[3] = false
        }
        break;
        case 2: {
            showedStocksList.value = savedStocksList.value.filter((stock) => stock.status === Status.requested)
            isSelected.value[0] = false
            isSelected.value[1] = false
            isSelected.value[2] = true
            isSelected.value[3] = false
        }
        break;
        case 3: {
            showedStocksList.value = savedStocksList.value.filter((stock) => stock.status === Status.requested_by)
            isSelected.value[0] = false
            isSelected.value[1] = false
            isSelected.value[2] = false
            isSelected.value[3] = true
        }
        break;
    }
    key_stock.value++;
}

function changeArrowIcon(whichIcon: number) {
    if (icon.value[whichIcon] === "arrow_drop_up") {
        icon.value[whichIcon] = "arrow_drop_down"
    } else {
        icon.value[whichIcon] = "arrow_drop_up"
    }
}

function updateStocksList(data: Stock[]){
  savedStocksList.value = data
  showedStocksList.value = savedStocksList.value
  key_stock.value += 1
  updateNumbersInFilters()
}

function createStockStyle(stock: Stock): string {
    if (stock.status === Status.owned) {
        return 'green'
    } else if (stock.status === Status.requested) {
        return 'red'
    } else {
        return '#409fff'
    }
}

onMounted(() => {
  savedStocksList.value = props.stocks
  showedStocksList.value = savedStocksList.value
  updateNumbersInFilters()
})

</script>

<style lang="scss" scoped>
.filters{
    display: flex;
    flex-direction: row;
    border-bottom: 2px solid #a6b1ad;
    margin-bottom: 2.5rem;
    margin-left: 2rem;
    width: 55%;
    
    h3{
      margin-left: 0.5rem;
      &+h3{
        background-color: #d9e4e0;
        padding: 5px;
        display: flex;
        align-items: center;
        justify-content: center;
        border-radius: 50%;
        width: 25px;
        height: 25px;
      } 
    }
    button{
      display: flex;
      flex-direction: row;
      align-items: center;
      margin-bottom: 1rem;
      margin-left: 5rem;
    }
    .all{
      margin-left: 1rem;
      color: v-bind(selectFilterColor(0));
     
    }
    .stocks{
      color: v-bind(selectFilterColor(1));
    }
    .my-requests{
      color: v-bind(selectFilterColor(2));
    }
    .others-requests{
      color: v-bind(selectFilterColor(3));
    }
  }
  span{
    color: #3b5998;
  }
  .sortButtons{
    margin-left: 2rem;
    display: flex;
    flex-direction: row;
    width: 100%;
    height: 5%;
    align-items: center;
    margin-bottom: 1rem;

    .sortID{
      display: flex;
      flex-direction: row;
      position: absolute;
      left: 19.5%;
    }
    .sortProducer{
      display: flex;
      flex-direction: row;
      position: absolute;
      left: 30%;
    }
    .sortOwner{
      display: flex;
      flex-direction: row;
      position: absolute;
      left: 45%;
    }
    .sortStatus{
      display: flex;
      flex-direction: row;
      position: absolute;
      left: 60%;
    }

    .sortRequester{
      display: flex;
      flex-direction: row;
      position: absolute;
      left: 75%;
    }
    .search{
      display: flex;
      flex-direction: row;
      position: absolute;
      left: 84.7%;
      align-items: center;
      background-color: #3b5998;
      padding: 0.5rem;
      border-radius: 0.25rem;
      .searchInput{
        border:none;
        background: none;
        outline:none;
        color: white;
        &::placeholder{
          color: #bebebe;
        }
      }
      span{
        color: white;
      }
    }
  }
  .allStocks{
    border: 1px solid #6e757e;
    margin-left: 2rem;
    width: 94%;
    display: flex;
    flex-direction: column;
    border-radius: 0.5rem;
    overflow: hidden;
    @media(max-width: 1450px) {
      margin-left: 1rem;
    }
  }
</style>
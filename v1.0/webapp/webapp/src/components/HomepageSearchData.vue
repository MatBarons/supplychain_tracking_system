<template>
    <main id="homepagesearchdata">
        <h1>Results </h1>
        <div v-if="stocks?.length !== 0">
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
                <form class="search"> 
                    <input class="searchInput" v-model="searchBarInput" type="text" id="query" name="q" placeholder="Search..."  @input="searchMyStocks">
                    <span class="material-icons">search</span>
                </form>
            </div>
            <div class="allStocks">
                <div class="stocks" v-for="stock in savedStocksList">
                   <StockSearch :stock=stock :odd="isElementOdd(savedStocksList, stock)" :key="key_stock"/>
                </div>
            </div>
        </div>
        <div v-else>
            <h1>La tua ricerca non ha prodotto risultati</h1>
        </div>
       
    </main>
</template>

<script lang="ts" setup>
import { Stock } from '@/types/stock';
import StockSearch from './StockSearch.vue';
import { onBeforeMount, ref } from 'vue';
import {compareUuid,compareProducer,compareOwner,compareStatus} from "@/utils/compare";

const props = defineProps({
    stocks: {
        type: Array<Stock>,
        required: false
    },
    key: {
        type: Number,
        required: true
    }
})

let showedStocksList = ref<Stock[]>([])
let savedStocksList = ref<Stock[]>([])
let key_stock = ref(0)
let searchBarInput = ref("")
let whichSort = ref("")
let icon = ref(["arrow_drop_up", "arrow_drop_up", "arrow_drop_up", "arrow_drop_up"])

function changeArrowIcon(whichIcon: number) {
    if (icon.value[whichIcon] === "arrow_drop_up") {
        icon.value[whichIcon] = "arrow_drop_down"
    } else {
        icon.value[whichIcon] = "arrow_drop_up"
    }
}

function sortStocks(whichCliked: string) {
    if (whichCliked === whichSort.value) {
        if (whichCliked === "id") {
            changeArrowIcon(0)
            showedStocksList.value.reverse()
        } else if (whichCliked === "owner") {
            changeArrowIcon(1)
            showedStocksList.value.reverse()
        } else if (whichCliked === "producer") {
            changeArrowIcon(2)
            showedStocksList.value.reverse()
        } else {
            changeArrowIcon(3)
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
        } else if (whichCliked === "owner") {
            showedStocksList.value.sort(compareProducer);
            showedStocksList.value.reverse();
            changeArrowIcon(1)
        } else if (whichCliked === "producer") {
            showedStocksList.value.sort(compareStatus);
            showedStocksList.value.reverse();
            changeArrowIcon(2)
        } else {
            showedStocksList.value.sort(compareOwner);
            changeArrowIcon(3)
        }
        whichSort.value = whichCliked;
    }

    key_stock.value++;
}

function isElementOdd(stocks: Stock[], stock: Stock): boolean {
    return stocks.indexOf(stock) % 2 === 0;
}

function searchMyStocks() {
    showedStocksList.value = showedStocksList.value.filter((stock) => {

        if (searchBarInput.value === "") {
            return savedStocksList
        }
        return stock.id.toString().includes(searchBarInput.value) ||
            stock.producer.includes(searchBarInput.value) ||
            stock.status.includes(searchBarInput.value) ||
            stock.request?.requester.includes(searchBarInput.value)
    })
    key_stock.value++;
}

onBeforeMount(() => {
    savedStocksList.value = props.stocks!
    showedStocksList.value = savedStocksList.value
})

</script>

<style lang="scss" scoped>
#homepagesearchdata{
    display: flex;
    flex-direction: column;
    h1{
        display: flex;
        color: black;
        justify-content: center;
        margin-bottom: 3rem;
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
        left: 35%;
        }
        .sortOwner{
        display: flex;
        flex-direction: row;
        position: absolute;
        left: 55%;
        }

        .sortStatus{
        display: flex;
        flex-direction: row;
        position: absolute;
        left: 70%;
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
        margin-bottom: 1rem;
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

}

</style>
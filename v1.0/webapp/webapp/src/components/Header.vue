<template>
    <header>
        <div class="searchBox">
            <input class="searchInput" v-model="searchInput" type="text"  name="" placeholder="Search a stock or a user"/>
            <button v-if="searchInput !== ''" class="removeText" @click="removeSearch">
                <i class="material-icons">cancel</i>
            </button>
            <button class="searchButton" @click="searchAll">
                <i class="material-icons">search</i>
            </button>
        </div>
        <div v-if="store.data.wallet === ''">
            <button class="connectWallet" @click="connectWallet">
                <h4>Connect wallet</h4>
            </button>
        </div>
        <div v-else>
            <button class="connectWallet" @click="disconnectWallet">
                <h4>Disconnect wallet</h4>
            </button>
        </div>
        <div class="createStockCollapsed" v-if="isExpanded === false">
            <button @click="toogle">
                <span class="material-icons">add</span> 
                <h4>Create stock</h4>
            </button>
        </div>
        <div class="createStockFull" v-else>
            <input class="createInput" v-model="createStockInput" id="query" name="q" placeholder="Input stock ID">
            <button @click="createStock">
                <span class="material-icons">{{ iconCreateStock }}</span>
            </button>
        </div>
    </header>
</template>

<script lang="ts" setup>
import { PeraWalletConnect } from "@perawallet/connect";
import { useDataStore } from "@/stores/store";
import { ref } from "vue";
import { addStock } from "@/api_calls/stocks";
import { createStockSocket, searchStocksSocket, walletConnectionSocket,getStocksSocket } from "@/api_calls/socket";
import { walletDisconnectionSocket } from "@/api_calls/socket";
import type { Stock } from "@/types/stock";


const emit = defineEmits<{
    (event: 'search_data', search: Stock[]): void
    (event: 'stock_data', data: Stock[]) : void
    (event: 'all_data',data: Stock[],search: Stock[]) : void
    
}>()

const peraWallet = new PeraWalletConnect({
    shouldShowSignTxnToast: false
});
const store = useDataStore();

let isExpanded = ref(false);
let createStockInput : string;
let iconCreateStock = ref("arrow_right_alt")
let searchInput = ref("")

function toogle(){
    isExpanded.value = !isExpanded.value
}

function removeSearch(){
    searchInput.value = ""
    let stocks : Stock[] = []
    getStocksSocket().then(response => {
        stocks = response
        return emit('all_data', stocks, [])
    })
    
}

async function createStock() {
    let stockList: Stock[] = []
    if (store.data.wallet !== "") {
        await addStock(createStockInput).then(response => {
            toogle()
            createStockSocket(createStockInput).then(response => {
                stockList = response
                return emit('stock_data', stockList)
            })
        }).catch(error => {
            console.log("Oh no: " + error)
        })
    }

}

function disconnectWallet() {
    peraWallet.disconnect().then(resolve => {
        store.changeWallet("")
        walletDisconnectionSocket()
    })
}

function connectWallet() {
    let stockList : Stock[] = []
    peraWallet.connect().then((accounts) => {
        peraWallet.connector?.on("disconnect", disconnectWallet);
        store.changeWallet(accounts[0])
        if(store.data.wallet !== undefined && store.data.wallet !== ""){
            walletConnectionSocket(store.data.wallet).then(response => {
                stockList = response
                return emit('stock_data', stockList)
            })
        }
    }).catch((e) => console.log(e));

    
}

function searchAll() {
    if(searchInput.value !== ""){
        searchStocksSocket(searchInput.value,store.data.wallet).then(response => {
            return emit("search_data",response)
        })
    }
    
}



/* 
onMounted(() => {
    peraWallet
        .reconnectSession()
        .then((accounts) => {
            peraWallet.connector?.on("disconnect", disconnectWallet);
            if (accounts.length) {
                store.changeWallet(accounts[0])
            }
        })
        .catch((error) => {
            if (error?.data?.type !== "CONNECT_MODAL_CLOSED") {
                console.log(error);
            }
        });
    console.log(store.data.wallet)
})*/
</script>

<style lang="scss" scoped>
header{
    display: flex;
    padding: 2rem;
    height: 20%;
    width: 100%;
    @media(max-width: 1450px){
        width: 60%;
    }
    .searchBox {
        display: flex;
        justify-content: center;
        position: absolute;
        left: 45%;
        background-color: #3b5998;
        width: 22%;
        border-radius: 0.35rem;
        padding: 0.5rem;
        .searchButton {
            color: white;
            float: right;
            width: fit-content;
            padding: 2%;
            border-radius: 50%;
            background-color: #3b5998;
            display: flex;
            justify-content: center;
            align-items: center;
            &:hover{
                background: white;
                color : #3b5998;
            }
            @media(max-width:1300px){
                
            }
        }
        .removeText{
            padding-right: 10px;
        }
        .searchInput {
            border:none;
            background: none;
            outline:none;
            float:left;
            padding: 0;
            color: white;
            font-size: 16px;
            line-height: 40px;
            width: 85%;
            padding-left: 1rem;
            &::placeholder{
                color: #bebebe;

            }
        }
    }
    .connectWallet{
        display: flex;
        font-weight: 600;
        padding: 1rem;
        position: absolute;
        right: 4.5%;
        height: 3rem;
        width: 10rem;
        border: 2px solid #3b5998;
        border-radius: 0.75rem;
        justify-content: center;
        align-items: center;
        h4{
            color: #274268;
            font-weight: bold;
        }
        &:hover{
            background-color: #3b5998;
        }
        &:hover > h4{
            color: white;
        }
    }

    .createStockCollapsed{
        display: flex;
        flex-direction: column;
        background-color: #3b5998;
        height: 25%;
        width: 7.5%;
        justify-content: center;
        border-radius: 0.5rem;
        transition: 0.5s;
        button{
            display: flex;
            flex-direction: column;
            align-items: center;
            
            span{
                color: white;
            }
            h4{
                color: white;
            }
        }
        
    }
    .createStockFull{
        display: flex;
        justify-content: center;
        height: 20%;
        border: 1px solid #3b5998;
        border-radius: 5px;
        overflow: hidden;
        .createInput{
            border: 0;
        }
        button{
            background-color: #3b5998;
            border: 0;
            border-radius: 0;
            &:hover{
                background-color: white;
            }
            &:hover > span{
                color: #3b5998;
            }
            span{
                color: white;
                width: 2.5rem;
            }
        }
    }
}
</style>

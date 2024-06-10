class userData{
    wallet: string;
    nomeAzienda: string;
    pIva: string;

    constructor(wallet: string, nomeAzienda: string, pIva: string){
        this.wallet = wallet;
        this.nomeAzienda = nomeAzienda;
        this.pIva = pIva;
    }
}

import { defineStore } from "pinia";

export const useDataStore = defineStore("data",{
    state: () => {
        return{
            data: new userData("","","")
        };
    },
    persist: {
        storage: sessionStorage
    },
    actions: {
        changeWallet(wallet: string){
            this.data.wallet = wallet
        },
        changeNomeAzienda(nomeAzienda: string){
            this.data.nomeAzienda = nomeAzienda
        },
        changePIVA(pIva: string){
            this.data.pIva = pIva
        },
        deleteAllData(){
            this.data.pIva = ""
            this.data.nomeAzienda = ""
            this.data.wallet = ""
        }
    },
})
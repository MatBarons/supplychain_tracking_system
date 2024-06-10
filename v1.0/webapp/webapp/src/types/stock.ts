import type { StockRequest } from "./request";

export enum Status{
    owned = "owned",
    requested = "requested",
    requested_by = "requested by",
    unavailable = "unavailable"
}

export class Stock {
    id: string;
    producer: string;
    status: Status;
    owner : string;
    request?: StockRequest;

    constructor(id: string, producer: string, status: Status, owner:string, request?: StockRequest){
        this.id = id;
        this.producer = producer;
        this.status = status;
        this.request = request;
        this.owner = owner
    }
}

export class StockHistory {
    id: string;
    historyNames: Array<string>
    historyWallets: Array<string>

    constructor(id: string, historyNames: Array<string>, historyWallets: Array<string>) {
        this.id = id;
        this.historyNames = historyNames;
        this.historyWallets = historyWallets;
    }
}

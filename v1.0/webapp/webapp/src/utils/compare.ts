import { Stock } from "@/types/stock";

function compareUuid(a: Stock, b: Stock): number{
    if (a.id < b.id) {
        return -1;
    }
    if (a.id > b.id) {
        return 1;
    }
    return 0;
}

function compareProducer(a: Stock, b: Stock): number{
    if (a.producer < b.producer) {
        return -1;
    }
    if (a.producer > b.producer) {
        return 1;
    }

    if (a.id < b.id) {
        return -1;
    }
    if (a.id > b.id) {
        return 1;
    }
    return 0;
}

function compareStatus(a: Stock, b: Stock): number{
    if (a.status < b.status) {
        return -1;
    }
    if (a.status > b.status) {
        return 1;
    }

    if (a.id < b.id) {
        return -1;
    }
    if (a.id > b.id) {
        return 1;
    }
    return 0;
}

function compareRequester(a: Stock, b: Stock) : number{
    if (a.request?.requester !== undefined && b.request?.requester === undefined) {
        return -1;
    }
    if (a.request?.requester === undefined && b.request?.requester !== undefined) {
        return 1;
    }
    if (a.request?.requester !== undefined && b.request?.requester !== undefined) {
        if (a.request?.requester < b.request?.requester) {
            return -1;
        }
        if (a.request?.requester > b.request?.requester) {
            return 1;
        }
    }
    return 0;
}

function compareOwner(a: Stock, b: Stock){
    if (a.owner !== undefined && b.owner === undefined) {
        return -1;
    }
    if (a.owner === undefined && b.owner !== undefined) {
        return 1;
    }
    if (a.owner !== undefined && b.owner !== undefined) {
        if (a.owner < b.owner) {
            return -1;
        }
        if (a.owner > b.owner) {
            return 1;
        }
    }
    return 0;
}

export { compareProducer, compareRequester, compareStatus, compareUuid, compareOwner}
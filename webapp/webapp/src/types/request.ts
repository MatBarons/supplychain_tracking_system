export class StockRequest{
    oldOwner: string;
    requester: string;
    isApproved: boolean;

    constructor(oldOwner: string, requester: string, isApproved: boolean){
        this.oldOwner = oldOwner
        this.requester = requester;
        this.isApproved = isApproved;
    }
}
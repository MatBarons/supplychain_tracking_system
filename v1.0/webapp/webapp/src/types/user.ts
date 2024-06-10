export class UserData{
    nomeAzienda : string;
    partitaIVA : string;
    walletAddress: string;

    constructor(nomeAzienda: string , partitaIVA: string , walletAddress: string){
        this.nomeAzienda = nomeAzienda;
        this.partitaIVA = partitaIVA;
        this.walletAddress = walletAddress;
    }
}
import * as sdk from "algosdk"
import * as pera from "@perawallet/connect";
import { useDataStore } from "../stores/store"
import appspec from '../dApp_schema/application.json'


let algodClient : any;
const appID = 264415535;
const contract = new sdk.ABIContract(appspec.contract)
const addStockMethodSelector = sdk.getMethodByName(contract.methods,"add_stock")
const changeOwnerMethodSelector = sdk.getMethodByName(contract.methods,"change_owner")
const deleteStockMethodSelector = sdk.getMethodByName(contract.methods,"delete_stock") //ADD THE METHOD

const store = useDataStore()
const connectPera = new pera.PeraWalletConnect();

function createAlgodClient() {
    return new sdk.Algodv2('', 'https://testnet-api.algonode.cloud', '')
}

async function signTxns(unsignedTxns: Array<sdk.Transaction>) {
    const signerTransactions = unsignedTxns.map(txn => {
        return {
            txn,
            signers: [sdk.encodeAddress(txn.from.publicKey)]
        }
    })
    return await connectPera.signTransaction([signerTransactions])
}

async function signer(txns: sdk.Transaction[]) {
    return await signTxns(txns)
}


export async function addStock(id: string) {
    if (algodClient === undefined) {
        algodClient = createAlgodClient();
    }
    try {
        let boxName = new Uint8Array(Buffer.from(id));
        const suggestedParams = await algodClient.getTransactionParams().do()
        const atc = new sdk.AtomicTransactionComposer();
        atc.addMethodCall({
            suggestedParams,
            sender: store.data.wallet,
            signer: signer,
            appID: appID,
            method: addStockMethodSelector,
            methodArgs: [id, store.data.wallet],
            boxes: [{ appIndex: appID, name: boxName }],
            note: boxName
        })
        const result = await atc.execute(algodClient, 3)
        console.log("confirmed round: " + result.confirmedRound)
        return result
    }catch(error){
        console.log("Something went wrong with addStock")
    }
}

export async function changeOwner(id: string){
    if (algodClient === undefined) {
        algodClient = createAlgodClient();
    }
    let boxName = new Uint8Array(Buffer.from(id));
    const suggestedParams = await algodClient.getTransactionParams().do()
    const atc = new sdk.AtomicTransactionComposer();
    atc.addMethodCall({
        suggestedParams,
        sender: store.data.wallet,
        signer: signer,
        appID: appID,
        method: changeOwnerMethodSelector,
        methodArgs: [id, store.data.wallet],
        boxes: [{ appIndex: appID, name: boxName }],
        note: boxName
    })
    const result = await atc.execute(algodClient, 3)
    console.log("confirmed round: " + result.confirmedRound)
}

export async function deleteStock(id: string){
    if (algodClient === undefined) {
        algodClient = createAlgodClient();
    }
    let boxName = new Uint8Array(Buffer.from(id));
    const suggestedParams = await algodClient.getTransactionParams().do()
    const atc = new sdk.AtomicTransactionComposer();
    atc.addMethodCall({
        suggestedParams,
        sender: store.data.wallet,
        signer: signer,
        appID: appID,
        method: deleteStockMethodSelector,
        methodArgs: [id],
        boxes: [{ appIndex: appID, name: boxName }],
        note: boxName

    })

    const result = await atc.execute(algodClient, 3)
    console.log("confirmed round: " + result.confirmedRound)
    return result;
}


export interface Stock{
    id: string,
    creationDate: string,
    remainingQuantity: number,
    totalQuantity: number,
    productType: string,
    expirationDate: string | null
}
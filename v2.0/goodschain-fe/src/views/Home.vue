<script setup lang="ts">
    import type { Stock } from '@/models/stock';

    let stocks : Array<Stock> = []

    let filterForm

    function deleteStock(id: string){

    }

    function search(){

    }
</script>

<template>
    <section>
        <button type="submit" @click="search()"></button>
    </section>
    <section>
        <table>
            <thead>
                <th>Id</th>
                <th>Creation date</th>
                <th>Expiration date</th>
                <th>Remaining quantity</th>
                <th>Total quantity</th>
                <th>Product type</th>
            </thead>
            <tbody>
                <tr v-for="stock in stocks">
                    <td>{{ stock.id }}</td>
                    <td>{{ stock.creationDate }}</td>
                    <td>{{ stock.expirationDate }}</td>
                    <td>{{ stock.remainingQuantity }}</td>
                    <td>{{ stock.totalQuantity }}</td>
                    <td>{{ stock.productType }}</td>
                    <td>
                        <button @click="deleteStock(stock.id)"></button>
                    </td>
                </tr>
            </tbody>
        </table>
    </section>
  
</template>

<style lang="scss"scoped>

</style>
<template>
    <main id="layout">
        <Sidebar/>
        <router-view/>
    </main>
</template>

<script lang="ts" setup>
import Sidebar from "@/components/Sidebar.vue";
import router from "@/router";
import { onMounted } from "vue";

onMounted(() => {
    router.push('/home')
})
</script>

<style lang="scss">
#layout {
    display: flex;
    height: 100%;
    min-width: 100%;
    @media(max-width: 1450px){
        min-width: fit-content;
    }
}
</style>
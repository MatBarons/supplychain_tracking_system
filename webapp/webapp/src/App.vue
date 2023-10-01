<script lang="ts" setup>
import { RouterView } from 'vue-router';

</script>

<template>
  <RouterView/>
</template>

<style lang="scss">
html{
  width: 100%;
  height: 100%;
  body{
    width: 100%;
    height: 100%;
    background-color: white;//rgba(0, 0, 0, 0.06);
    #app{
      width: 100%;
      height: 100%;
      padding: 0;
      margin: 0;
      @media(max-width: 768px){
       justify-items: flex-start;
      }
    }
  }
}


*,*::after,*::before{
  margin:0;
  padding: 0;

}
button{
  cursor: pointer;
  appearance: none;
  border: none;
  outline: none;
  background: none;
}

</style>

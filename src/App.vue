<template>
  <body>
    <router-view />
  </body>
</template>

<script>
import { provide,ref } from 'vue';
export default {
  setup(){
    const userID=ref("");
    provide('userID',userID);
    return{
      userID
    };
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 10px;
}
</style>

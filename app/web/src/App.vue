<template>
  <div id="app">
    <h1>Dice Online</h1>
    <Auth />
  </div>
</template>

<script>
import Auth from './components/Auth.vue'

export default {
  name: 'app',
  components: {
    Auth
  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin:0;
  background-image: linear-gradient(190deg, #0F2027 , #203A43, #2C5364);
  height:100vh;
  width:100vw;
}

h1{
  margin:0;
  color:white;
}

body{
  margin:0;
}
</style>

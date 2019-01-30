<template>
  <div id="app">
    <h1>Dice Online</h1>
    <component v-bind:is="authComponent" @clicked="authClicked"></component>
    <Game v-bind:userid="userid"/>
  </div>
</template>

<script>
import Auth from './components/Auth.vue'
import Game from './components/Game.vue'

export default {
  name: 'app',
  data(){
    return {
      userid: -1, 
      authComponent:"Auth"
    }
  },
  components: {
    Auth, 
    Game
  },
  methods:{
    authClicked: function(value){
      this.userid = value.data["ID"]
      this.authComponent = ""
    }
  }
}
</script>

<style>
#app {
  display:flex;
  flex-direction:column;
  align-content: center;
  align-items: center;
  justify-content: center;
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


.fade-enter-active, .fade-leave-active {
  transition: opacity .5s;
}
.fade-enter, .fade-leave-to /* .fade-leave-active below version 2.1.8 */ {
  opacity: 0;
}
</style>

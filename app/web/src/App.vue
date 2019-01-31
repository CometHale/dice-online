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
  flex-wrap:nowrap;
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

div {
  display:flex;
  flex-direction:column;
  flex-wrap:wrap;
  justify-content: center;
  align-items:center;

  min-width:100px;
  min-height:200px;
  border-radius:3px;
}

input{
  margin-bottom:5px;
  font-weight: bold;
  width:100%;
  font-size:12px;
  height:34px;

  border:none;
  border-radius:3px;
  
}


input::-webkit-input-placeholder {
  text-align: center;
  line-height:34px; /* must match height of input to vertically center*/
}

input:-moz-placeholder { /* Firefox 18- */
  text-align: center;
  line-height:34px; /* must match height of input to vertically center*/
}

input::-moz-placeholder {  /* Firefox 19+ */
  text-align: center;
  line-height:34px; /* must match height of input to vertically center*/
}

input:-ms-input-placeholder {  
  text-align: center;
  line-height:34px; /* must match height of input to vertically center*/
}

button{
  width:200px;
  min-height:50px;
  border:none;
  background-color:rgb(65, 124, 150);
  color:black;
  text-align: center;
  text-decoration: none;
  display: inline-block;
  margin-left:auto;
  margin-right:auto;
  margin-bottom:10px;
  font-weight: bold;
  font-size:15px;
}

button:hover{
  background-color:cadetblue;
}
.fade-enter-active, .fade-leave-active {
  transition: opacity .5s;
}
.fade-enter, .fade-leave-to /* .fade-leave-active below version 2.1.8 */ {
  opacity: 0;
}
</style>

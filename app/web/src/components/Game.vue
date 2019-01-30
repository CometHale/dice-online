<template>
  <div id="game-background">
    <div id="game">
      <input @change="makeGuess" type="number" name="goal" placeholder="Guess what the dice will show">
      <div id="dice">
        <button @click="rollDice" id="game-button">Roll The Dice</button>
        <Result />
      </div>
      <HighScore />
    </div>
  </div>
</template>

<script>
import Result from './Result.vue'
import HighScore from './HighScore.vue'
import axios from 'axios'
import qs from 'qs';

export default {
  name: 'Game',
  props:["userid"],
  components:{
    Result,
    HighScore
  },
  methods:{
    makeGuess: function(Event){

      var goal = Event.target.value
      const data = {
        'userid': this.userid,
        'goal': goal
      }

      axios({
        method:'post',
        url:'http://localhost:3000/roll-dice/',
        data: qs.stringify(data),
        config: {headers: {'Content-Type': 'application/x-www-form-urlencoded'}}
      }).then(function (response) {
        console.log(response)
      }).catch(function (response) {
          //handle error
          console.log(response);
          alert("An error occurred while creating your account. Please try again later.")
      });
    },
    rollDice: function(Event){

      var goal = Event.target.value
      const data = {
        'userid': this.userid,
        'goal': goal
      }

      axios({
        method:'get',
        url:'http://localhost:3000/roll-dice/',
        data: qs.stringify(data),
        config: {headers: {'Content-Type': 'application/x-www-form-urlencoded'}}
      }).then(function (response) {
        //handle success
        console.log(response);

      }).catch(function (response) {
          //handle error
          console.log(response);
          alert("An error occurred while creating your account. Please try again later.")
      });
    }
  }

}
</script>

<style scoped>

  div#game-background{
    margin-top: 60px;
    display:flex;
    flex-direction:column;
    width:45%;
    min-height:200px;
    background-color:white;
    border-radius:3px;
    justify-content: center;
    align-items:center;
  }
</style>

<template>
  <div id="game-background" v-if="userid != -1">
    <div id="tabs">
      <button @click="switchTab" id="game-tab" class="tab" v-bind:class="{ active: !showHighScores, inactive:showHighScores}">Play Game</button>
      <button @click="switchTab" id="high-score-tab" class="tab" v-bind:class="{ active: showHighScores, inactive:!showHighScores}">View High Scores</button>
    </div>
    <div id="pages">
      <div id="game" v-if="!showHighScores">
        <p id="instructions" v-if="resultComponent == ''">Guess what the dice will show, and then hit the button to roll it!</p>
        <div id="dice-game">
          <input id="goal-input" @change="makeGuess" type="number" name="goal" placeholder="Guess the dice roll">
          <button @click="rollDice" id="game-button">Roll the dice</button>
          <component v-bind:is="resultComponent" v-bind:userid="userid" v-bind:score="score" v-bind:result="result" v-bind:roll="roll" v-bind:goal="goal"></component>
        </div>
      </div>
      <HighScore v-if="showHighScores" v-bind:userid="userid"  v-bind:userhighscore="userhighscore" v-bind:allusers="allusers"/>
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
  data(){
    return {
      goal:-1,
      result:false,
      score:0,
      userhighscore:0,
      roll:-1,
      resultComponent:"",
      showHighScores:false,
      allusers:[]
    }
  },
  components:{
    Result,
    HighScore
  },
  methods:{
    makeGuess: function(Event){

      this.goal = Event.target.value
      
      const data = {
        'userid': this.userid,
        'goal': this.goal
      }

      axios({
        method:'post',
        url:'http://localhost:3000/start-game/',
        data: qs.stringify(data),
        config: {headers: {'Content-Type': 'application/x-www-form-urlencoded'}}
      }).then(function (response) {
        console.log(response);
      }).catch(function (response) {
          //handle error
          console.log(response);
          alert("An error occurred while setting up your game. Please try again later.")
      });
    },
    rollDice: function(Event){
  
      var $this = this;

      const data = {
        'userid': this.userid,
        'goal': this.goal
      }

      axios({
        method:'post',
        url:'http://localhost:3000/roll-dice/',
        data: qs.stringify(data),
        config: {headers: {'Content-Type': 'application/x-www-form-urlencoded'}}
      }).then(function (response) {
        //handle success
        $this.result = response.data["Result"]
        $this.score = response.data["Score"]
        $this.userhighscore = response.data["UserHighScore"]
        $this.roll = response.data["Roll"]
        $this.resultComponent = "Result"


        // hide the game instructions
        var instructions = document.getElementById("instructions");
        instructions.style.display = "none";

      }).catch(function (response) {
          //handle error
          console.log(response);
          alert("An error occurred while rolling the dice. Please try again later.")
      });
    },
    switchTab: function(Event){
        if(!Event.target.classList.contains("active")){
          this.showHighScores = !this.showHighScores;

          if(Event.target.id == "high-score-tab"){ // if the tab clicked was the high score tab

              var $this = this;
            // get the high scores of all users
            axios({
              method:'get',
              url:'http://localhost:3000/view-all/'
            }).then(function (response) {
              $this.allusers = response.data;
              console.log(response);
            }).catch(function (response) {
                //handle error
                console.log(response);
                alert("An error occurred while querying for user high scores. Please try again later.")
            });

          }
        }
    }
  }
}
</script>

<style scoped>

  #game-background{
    margin-top: 60px;
    display:flex;
    flex-direction:column-reverse;
    width:45vw;
    min-height:200px;
    background-image: linear-gradient(-190deg, whitesmoke , #d7d2cc);
    border-radius:3px;
    justify-content: flex-end;
    align-items:center;
  }

  #pages{
    display:flex;
    flex-direction:column;
    width:100%;
  }

  #game{
    width:100%;
    flex-direction:column;
    justify-content: space-evenly;
  }

  input{
    width:auto;    
  }

  p{
    text-align:center;
    width:50%;
    margin:0;
  }

  button{
    width:auto;
    min-height: 36px;
    font-size:12px;
  }


  #dice-game{
    min-height:auto;
    width:90%;
    flex-direction:row;
    justify-content: space-between;
  }

  #tabs{
    display:flex;
    flex-direction:row;
    position:fixed;
    min-height:40px;
    top:200px;
    left:197px;
    z-index:2;
  }

  .tab{
    background-color:whitesmoke;
    width: 100px;
    height:40px;
    color:black;
    border:none;
    font-weight:bold;
    font-size:12px;
  }

  #game-tab{
    border-top-left-radius: 3px;
  }

  #high-score-tab{
    border-top-right-radius: 3px;
    border-bottom-right-radius: -3px;
  }

  .active{
    text-decoration: underline;
  }

  .inactive{
    background-color:lightgray;
  }


</style>

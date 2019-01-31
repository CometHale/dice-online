<template>
  <div id="high-scores">
      <div id="user-high-score">
          <p>Your Highest Score is {{ userhighscore }}</p>
      </div>
      <hr>
      <div id="all-high-scores">
          <p>Site Score Board</p>
          <ul>
            <li v-for="(user, index) in allusers" :key="user.HighScore" v-bind:class="{ currentuser: user.ID == userid}">
              <p>#{{index + 1}}</p>
              <p v-if="user.ID == userid">
                  (This is You!)
              </p>
               <p v-if="user.ID != userid">
                  {{ user.Username }}
              </p>
              
              <p> with Score: {{ user.HighScore}}</p>
            </li>
          </ul>
      </div>
  </div>
</template>

<script>

export default {
  name: 'HighScore',
  props:["userhighscore", "userid", "allusers"]
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

    div{
        min-height:auto;
    }

    p{
        margin:0;
    }

    div#high-scores{
        display:flex;
        flex-direction:column;
        width:100%;
        min-height:170px;
        border-radius:3px;
        justify-content: center;
        align-items:center;
    }

    #all-high-scores{
        min-height: 50px;
        min-width:250px;
    }

    #all-high-scores p{
        text-decoration-line: underline;
    }

    .currentuser{
        color:red;
    }

    ul{
        overflow-y:scroll;
        overflow-x:hidden;
        height:75px;
        width:100%;
        margin:0;
        margin-top:10px;
        padding:0;
        background-color:whitesmoke;
    }
    li{
        text-decoration:none;
        list-style: none;
        display:flex;
        justify-content: space-evenly;
    }

   #all-high-scores ul li p{
        text-decoration-line: none;
    }

    hr{
        width:200px;
    }
</style>

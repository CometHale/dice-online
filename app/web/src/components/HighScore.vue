<template>
  <div id="high-scores">
      <div id="user-high-score">
          <p>Your Highest Score is {{ userhighscore }}</p>
      </div>
      <div id="all-high-scores">
          <ul>
            <li v-for="(user, index) in allusers" :key="user.HighScore" v-bind:class="{ currentuser: user.ID == userid}">
              {{index}} : {{ user }}
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

    div#high-scores{
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

    .currentuser{
        color:red;
    }
</style>

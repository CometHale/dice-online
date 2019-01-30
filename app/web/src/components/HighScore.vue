<template>
  <div id="high-scores">
      <div id="user-high-score">

      </div>
      <div id="all-high-scores">

      </div>
  </div>
</template>

<script>

export default {
  name: 'HighScore',

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
</style>

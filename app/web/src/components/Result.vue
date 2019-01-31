<template>
  <div id="result">
    <h2>Game Result</h2>
    <p>The dice landed on {{ roll }}.</p>
    <p>Your guess was {{ goal }}.</p>
    <p v-if="goal == roll">You Won!</p>
    <p v-if="goal != roll">You Lost</p>
  </div>
</template>

<script>

export default {
  name: 'Result',
  props:['userid', "score", "result", "roll", "goal"]
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>

<template>
  <div id="result">
    <h3>Game Result</h3>
    <div>
      <p>The dice landed on {{ roll }}.</p>
      <p>Your guess was {{ goal }}.</p>
      <p v-if="goal == roll">You Won!</p>
      <p v-if="goal != roll">You Lost.</p>
    </div>
  </div>
</template>

<script>

export default {
  name: 'Result',
  props:['userid', "score", "result", "roll", "goal"]
}
</script>

<style scoped>


#result{
  display:flex;
  margin:auto auto;
}

p{
  margin:0;
}

div{
  min-height:auto;
}
</style>

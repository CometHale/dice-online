<template>
  <div id="auth">
    <form id="sign-up" method="post" @submit.prevent="signup">
      <input type="text" name="email" placeholder="email">
      <input type="text" name="username" placeholder="username">
      <input type="password" name="password" placeholder="password">
      <button>Sign Up</button>
    </form>
    <form id="login" method="post" @submit.prevent="login">
      <input type="text" name="email" placeholder="email">
      <input type="password" name="password" placeholder="password">
      <button>Login</button>
    </form>
  </div>
</template>

<script>
import axios from 'axios';
import qs from 'qs';

export default {
  name: 'Auth',
  props:["isHidden"],
  methods: {
    signup: function (Event) {
      var email = Event.target.elements.email.value
      var username = Event.target.elements.username.value
      var password = Event.target.elements.password.value
      const data = {
        'email':email,
        'username':username,
        'password':password
      }
 
      axios({
        method:'post',
        url:'http://localhost:3000/user-create/',
        data: qs.stringify(data),
        config: {headers: {'Content-Type': 'application/x-www-form-urlencoded'}}
      }).then(function (response) {
        //handle success
        console.log(response);
        var auth = document.getElementById("auth")
        auth.style.display = "none";
      }).catch(function (response) {
          //handle error
          console.log(response);
          alert("An error occurred while creating your account. Please try again later.")
      });

    }, 
    login: function (Event){
      var email = Event.target.elements.email.value
      var password = Event.target.elements.password.value
      const data = {
        'email':email,
        'password':password
      }
 
      axios({
        method:'post',
        url:'http://localhost:3000/user-login/',
        data: qs.stringify(data),
        config: {headers: {'Content-Type': 'application/x-www-form-urlencoded'}}
      }).then(function (response) {
          //handle success
          console.log(response);
          var auth = document.getElementById("auth")
          auth.style.display = "none";
          
      }).catch(function (response) {
          //handle error
          console.log(response);
          alert("An error occurred while logging you in. Please try again later.")
      });
    }
  }
}


</script>
 
<style scoped>

#auth{

  display:flex;
  flex-direction:row;
  width:100%;
  justify-content: space-around;
  align-items:center;
  padding-top:60px;
}

form {
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

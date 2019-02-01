<template>
    <div id="auth">
      <div class="form-wrapper">
        <h2>Create an Account</h2>
        <form id="sign-up" method="post" @submit.prevent="signup">
          <input type="text" name="email" placeholder="email">
          <input type="text" name="username" placeholder="username">
          <input type="password" name="password" placeholder="password">
          <button>Sign Up</button>
        </form>
      </div>
      <div class="form-wrapper">
        <h2>Login</h2>
        <form id="login" method="post" @submit.prevent="login">
          <input type="text" name="email" placeholder="email">
          <input type="password" name="password" placeholder="password">
          <button>Login</button>
        </form>
      </div>
    </div>
</template>

<script>
import axios from 'axios';
import qs from 'qs';

export default {
  name: 'Auth',
  methods: {
    signup: function (Event) {
      
      var email = Event.target.elements.email.value
      var username = Event.target.elements.username.value
      var password = Event.target.elements.password.value
      var $this = this

      const data = {
        'email':email,
        'username':username,
        'password':password
      }

      axios({
        method:'post',
        url:' https://dice-online-api.herokuapp.com/user-create/',
        data: qs.stringify(data),
        config: {headers: {'Content-Type': 'application/x-www-form-urlencoded'}}
      }).then(function (response) {
        //handle success
        $this.$emit('clicked', response); // can't emit directly within axios

      }).catch(function (response) {
          //handle error
          console.log(response);
          alert("An error occurred while creating your account. Please try again later.")
      });

    }, 
    login: function (Event){
      var email = Event.target.elements.email.value;
      var password = Event.target.elements.password.value;

      var $this = this;
      const data = {
        'email':email,
        'password':password
      };
 
      axios({
        method:'post',
        url:' https://dice-online-api.herokuapp.com/user-login/',
        data: qs.stringify(data),
        config: {headers: {'Content-Type': 'application/x-www-form-urlencoded'}}
      }).then(function (response) {
          //handle success
          console.log(response);
          $this.$emit('clicked', response); // can't emit directly within axios
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

.form-wrapper{
  background-image: linear-gradient(-190deg, white , #d7d2cc);
}

form{
  display:flex;
  flex-direction:column;
  padding: 25px 75px;
}

</style>

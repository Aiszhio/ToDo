<template>
  <div id="Entrance">
    <form @submit.prevent="handleSubmit" class="EntForm">
      <p class="Entname">Авторизация</p>
      <label for="email" class="Enttext">Почта:</label>
      <input type="email" v-model="form.email" id="email" name="email" required><br><br>

      <label for="password" class="Enttext">Пароль:</label>
      <input type="password" v-model="form.password" id="password" name="password" required><br><br>

      <input type="submit" value="Войти" class="entbutton">
    </form>
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: 'Entrance',
  data() {
    return {
      form: {
        email: '',
        name: '',
        password: ''
      }
    }
  },
  methods: {
    async handleSubmit() {
      this.errorMessage = '';

      try {
        const response = await axios.post('http://localhost:5174/entrance', this.form);

        if (response.status === 200) {
          alert('Авторизация успешна!');
        }
      } catch (error) {
        if (error.response && error.response.status === 401) {
          this.errorMessage = 'Неверный логин или пароль';
        } else {
          this.errorMessage = 'Произошла ошибка, попробуйте позже';
        }
      }
    }
  }
};
</script>

<style scoped>
#Entrance {
  display: flex;
  flex-direction: row;
  justify-content: center;
  margin-top: 1%;
}

p {
  color: black;
}

.entbutton {
  padding: 15px;
  background-color: #FFDAB9;
  border-radius: 20px;
  width: 60%;
  font-family: "Roboto Light", Arial, sans-serif;
  font-size: large;
}

.entbutton:hover{
  opacity: .8;
  text-decoration: none;
  border-radius: 20px;
  background-color: #e3a95e;
  font-weight: bolder;
}

.Entname {
  margin-bottom: 5vh;
  font-family: "Roboto Light", Arial, sans-serif;
  font-size: x-large;
}

.EntForm{
  display: flex;
  flex-direction: column;
  align-items: center;
  width: 30%;
  border: 3px solid black;
  border-radius: 20px;
  padding: 35px;
}

.Enttext{
  font-family: "Roboto Light", Arial, sans-serif;
  font-size: large;
}

#email{
  margin-top: 2%;
  height: 5%;
  width: 45%;
  border: 2px solid black;
  border-radius: 4px;
  padding: 15px;
  font-size: 16px;
  box-sizing: border-box;
  outline: none;
  transition: border-color 0.3s;
}

#name{
  margin-top: 2%;
  height: 5%;
  width: 45%;
  border: 2px solid black;
  border-radius: 4px;
  padding: 15px;
  font-size: 16px;
  box-sizing: border-box;
  outline: none;
  transition: border-color 0.3s;
}

#password{
  margin-top: 2%;
  height: 5%;
  width: 45%;
  border: 2px solid black;
  border-radius: 4px;
  padding: 15px;
  font-size: 16px;
  box-sizing: border-box;
  outline: none;
  transition: border-color 0.3s;
}
</style>

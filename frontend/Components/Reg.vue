<template>
  <div id="Register">
      <form @submit.prevent="submitForm" class="RegForm">
        <p class="Regname">Регистрация</p>
        <label for="email" class="Regtext">Почта:</label>
        <input type="email" v-model="form.email" id="email" name="email" required><br><br>

        <label for="name" class="Regtext">Имя пользователя:</label>
        <input type="text" v-model="form.name" id="name" name="name" required><br><br>

        <label for="password" class="Regtext">Пароль:</label>
        <input type="password" v-model="form.password" id="password" name="password" required><br><br>

        <input type="submit" value="Зарегистрироваться" class="regbutton">
      </form>
  </div>
</template>

<script>
export default {
  name: 'Register',
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
    async submitForm() {
      try {
        const response = await fetch('http://localhost:5174/register', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json'
          },
          body: JSON.stringify(this.form)
        });

        if (!response.ok) {
          throw new Error('Сервер вернул ошибку');
        }

        alert('Регистрация прошла успешно!');
        this.$router.push('/')
        this.form.email = '';
        this.form.name = '';
        this.form.password = '';
      } catch (error) {
        console.error('Ошибка:', error);
        alert('Произошла ошибка при регистрации.');
      }
    }
  }
}
</script>

<style scoped>
#Register {
  display: flex;
  flex-direction: row;
  justify-content: center;
  margin-top: 1%;
}

p {
  color: black;
  font-family: Arial;
}

.regbutton {
  padding: 15px;
  border: 1px solid white;
  border-radius: 10px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
  width: 60%;
  font-size: large;
  cursor: pointer;
  transition: background-color 0.3s ease, transform 0.3s ease;
}

.regbutton:hover{
  opacity: .8;
  text-decoration: none;
  border-radius: 20px;
  background-color: #edbc91;
  font-weight: bolder;
  transform: scale(0.95);
}

.Regname {
  margin-bottom: 5vh;
  font-family: "Roboto Light", Arial, sans-serif;
  font-size: x-large;
}

.RegForm{
  display: flex;
  flex-direction: column;
  align-items: center;
  width: 30%;
  border: 1px solid #FFDAB9;
  border-radius: 10px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
  background-color: #FFDAB9;
  padding: 35px;
}

.Regtext{
  font-family: "Roboto Light", Arial, sans-serif;
  font-size: large;
}

#email{
  margin-top: 2%;
  height: 5%;
  width: 45%;
  border: 1px solid white;
  border-radius: 10px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
  padding: 15px;
  font-size: 16px;
  box-sizing: border-box;
  outline: none;
  transition: border-color 0.1s;
}

#name{
  margin-top: 2%;
  height: 5%;
  width: 45%;
  border: 1px solid white;
  border-radius: 10px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
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
  border: 1px solid white;
  border-radius: 10px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
  padding: 15px;
  font-size: 16px;
  box-sizing: border-box;
  outline: none;
  transition: border-color 0.3s;
}
</style>

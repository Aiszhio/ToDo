<template>
  <div id="Register">
    <div id="RegField">
      <form @submit.prevent="submitForm">
        <p class="Regname">Регистрация</p>
        <label for="email">Почта:</label>
        <input type="email" v-model="form.email" id="email" name="email" required><br><br>

        <label for="name">Имя пользователя:</label>
        <input type="text" v-model="form.name" id="name" name="name" required><br><br>

        <label for="password">Пароль:</label>
        <input type="password" v-model="form.password" id="password" name="password" required><br><br>

        <input type="submit" value="Зарегистрироваться" class="regbutton">
      </form>
    </div>
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
        const response = await fetch('/register', {
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
        // Очищаем форму после успешной регистрации
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

<style>
#Register {
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: center;
}

p {
  color: black;
}

#RegField {
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: center;
  height: 60vh;
  border: 2px solid black;
  border-radius: 25px;
  width: 30%;
}

.regbutton {
  padding: 15px;
  background-color: white;
  border-radius: 50px;
  width: 60%;
  margin-left: 5vh;
  margin-top: 2vh;
}

.Regname {
  margin-bottom: 5vh;
  margin-left: 8vh;
}
</style>

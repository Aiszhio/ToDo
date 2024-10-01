<script>
import axios from "axios";
import { jwtDecode } from 'jwt-decode';

export default {
  name: 'Entrance',
  data() {
    return {
      form: {
        email: '',
        password: ''
      },
      errorMessage: ''
    }
  },
  methods: {
    async handleSubmit() {
      this.errorMessage = '';
      try {
        // Отправляем запрос на вход
        const response = await axios.post('http://localhost:5174/entrance', this.form);

        if (response.status === 200) {
          alert("Authentification is good!");
          // Получаем токен после успешной аутентификации
          const tokenResponse = await axios.post('http://localhost:5174/receive-token');

          if (tokenResponse.status === 200 && tokenResponse.data.token) {
            const token = tokenResponse.data.token;
            localStorage.setItem('jwtToken', token);
            console.log(token);

            try {
              // Декодируем токен и сохраняем имя пользователя
              const decoded = jwtDecode(token);
              console.log('Decoded token:', decoded);
              localStorage.setItem('userName', decoded.user_name);

              // Проверка срока действия токена
              const expirationTime = decoded.exp * 1000; // преобразуем Unix time в миллисекунды
              const currentTime = Date.now();

              if (expirationTime <= currentTime) {
                this.logout(); // Если токен истек, выполняем логаут
              } else {
                // Устанавливаем таймер для автоматического логаута
                const remainingTime = expirationTime - currentTime;
                setTimeout(() => {
                  this.logout();
                }, remainingTime);
                // Перенаправляем на домашнюю страницу
                this.$router.push('/');
              }

            } catch (error) {
              console.error('Ошибка декодирования токена или перенаправления:', error);
            }
          } else {
            this.errorMessage = 'Не удалось получить токен.';
          }
        }
      } catch (error) {
        if (error.response && error.response.status === 401) {
          this.errorMessage = 'Неверный логин или пароль';
        } else {
          this.errorMessage = 'Произошла ошибка, попробуйте позже';
        }
      }
    },

    logout() {
      // Очистка токена и имени пользователя из localStorage
      localStorage.removeItem('jwtToken');
      localStorage.removeItem('userName');
      // Перенаправление на страницу логина
      this.$router.push('/login');
    }
  }
};
</script>

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

<style scoped>
#Entrance {
  display: flex;
  flex-direction: row;
  justify-content: center;
  margin-top: 1%;
}

p {
  color: black;
  font-family: Arial;
}

.entbutton {
  padding: 15px;
  border: 1px solid white;
  border-radius: 10px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
  width: 60%;
  font-family: "Roboto Light", Arial, sans-serif;
  font-size: large;
  cursor: pointer;
  transition: background-color 0.3s ease, transform 0.3s ease;
}

.entbutton:hover{
  opacity: .8;
  text-decoration: none;
  border-radius: 20px;
  background-color: #edbc91;
  font-weight: bolder;
  transform: scale(0.95);
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
  border: 1px solid #FFDAB9;
  border-radius: 10px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
  background-color: #FFDAB9;
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
  border: 1px solid white;
  border-radius: 10px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
  padding: 15px;
  font-size: 16px;
  box-sizing: border-box;
  outline: none;
  transition: border-color 0.1s;
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
  transition: border-color 0.1s;
}
</style>

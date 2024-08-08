<template>
    <div id="Register">
      <div id="RegField">
        <form action="/registerSend" method="post">
          <p class="Regname">Регистрация</p>
          <label for="email">Почта:</label>
          <input type="email" id="email" name="email" required><br><br>

          <label for="name">Имя пользователя:</label>
          <input type="text" id="username" name="username" required><br><br>

          <label for="password">Пароль:</label>
          <input type="password" id="password" name="password" required><br><br>

          <input type="submit" value="Зарегистрироваться" class="regbutton">
        </form>
      </div>
    </div>
</template>

<script>
export default {
  name: 'register'
}
</script>

<style>
#Register{
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: center;
}
p{
  color: black;
}
#RegField{
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: center;
  height: 60vh;
  border: 2px solid black;
  border-radius: 25px;
  width: 30%;
}
.regbutton{
  padding: 15px;
  background-color: white;
  border-radius: 50px;
  width: 60%;
  margin-left: 5vh;
  margin-top: 2vh;
}
.Regname{
  margin-bottom: 5vh;
  margin-left: 8vh;
}
</style>


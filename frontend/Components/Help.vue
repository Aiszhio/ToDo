<script>
export default{
  name:"help",
}
</script>

<template>
  <div id="help">
    <section class="help-section">
      <h2>Помощь</h2>
      <div class="help-content">
        <div class="help-box">
          <h3>Регистрация</h3>
          <p>Для создания учетной записи перейдите на страницу регистрации. Введите свои данные: имя, email и пароль. После успешной регистрации вы сможете войти в систему.</p>
        </div>
        <div class="help-box">
          <h3>Логин</h3>
          <p>На странице входа введите ваш email и пароль, которые вы использовали при регистрации. Если данные корректны, вы попадете в свой личный кабинет.</p>
        </div>
        <div class="help-box">
          <h3>Добавление заметки</h3>
          <p>Чтобы добавить заметку, перейдите на страницу "Заметки". Введите заголовок, текст заметки и выберите дату для планирования или сохранения события.</p>
        </div>
        <div class="help-box">
          <h3>Личный кабинет</h3>
          <p>В личном кабинете вы можете просматривать все свои заметки. Они будут отсортированы по дате, и вы сможете легко найти нужную информацию.</p>
        </div>
      </div>
    </section>
  </div>
</template>


<style scoped>
#help {
  display: flex;
  justify-content: center;
  margin-top: 1%;
}

.help-section {
  width: 80%;
  max-width: 800px;
  padding: 20px;
  background-color: #FFDAB9;
  border-radius: 8px;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
}

.help-section h2 {
  text-align: center;
  margin-bottom: 20px;
  font-size: 1.8em;
  color: #333;
}

.help-content {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.help-box {
  padding: 15px;
  background-color: white;
  border: 1px solid #ddd;
  border-radius: 6px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.05);
}

.help-box h3 {
  margin-bottom: 10px;
  font-size: 1.4em;
  color: #555;
}

.help-box p {
  font-size: 1em;
  color: #666;
  line-height: 1.5;
}
</style>

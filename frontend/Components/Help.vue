<script>
export default{
  name:"help",
}
</script>

<template>
  <div class="layout-center">
    <section class="card help-section">
      <h2 class="title-md text-center mb-2">Помощь</h2>

      <div class="help-content">
        <div class="help-box">
          <h3 class="title-sm">Регистрация</h3>
          <p>
            Для создания учетной записи перейдите на страницу регистрации. Введите свои данные:
            имя, email и пароль. После успешной регистрации вы сможете войти в систему.
          </p>
        </div>
        <div class="help-box">
          <h3 class="title-sm">Логин</h3>
          <p>
            На странице входа введите ваш email и пароль, которые вы использовали при регистрации.
            Если данные корректны, вы попадете в свой личный кабинет.
          </p>
        </div>
        <div class="help-box">
          <h3 class="title-sm">Добавление заметки</h3>
          <p>
            Чтобы добавить заметку, перейдите на страницу «Заметки». Введите заголовок, текст заметки
            и выберите дату для планирования или сохранения события.
          </p>
        </div>
        <div class="help-box">
          <h3 class="title-sm">Личный кабинет</h3>
          <p>
            В личном кабинете вы можете просматривать все свои заметки. Они будут отсортированы
            по дате, и вы сможете легко найти нужную информацию.
          </p>
        </div>
      </div>
    </section>
  </div>
</template>
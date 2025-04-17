<template>
  <div class="register-container">
    <form @submit.prevent="onSubmit" class="form">
      <h2 class="title">Регистрация</h2>

      <label class="label">
        Почта
        <input v-model="form.email" type="email" required />
      </label>

      <label class="label">
        Имя пользователя
        <input v-model="form.name" type="text" required />
      </label>

      <label class="label">
        Пароль
        <input v-model="form.password" type="password" required />
      </label>

      <button type="submit" class="btn">Зарегистрироваться</button>
    </form>
  </div>
</template>

<script setup>
import { reactive } from 'vue'
import { useRouter } from 'vue-router'
import axios from 'axios'

// реактивная форма
const form = reactive({
  email: '',
  name: '',
  password: ''
})

const router = useRouter()

async function onSubmit() {
  try {
    await axios.post('http://localhost:5174/register', form, {
      headers: { 'Content-Type': 'application/json' }
    })
    alert('Регистрация прошла успешно!')
    // сброс формы
    form.email = form.name = form.password = ''
    router.push('/')
  } catch (e) {
    console.error('Registration error:', e)
    alert('Ошибка при регистрации. Попробуйте ещё раз.')
  }
}
</script>

<style scoped>
.register-container {
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 2rem;
  min-height: 80vh;
  background: #fff;
}

.form {
  background: #FFDAB9;
  padding: 2rem;
  border-radius: 8px;
  box-shadow: 0 4px 8px rgba(0,0,0,0.1);
  width: 100%;
  max-width: 360px;
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

.title {
  text-align: center;
  font-size: 1.5rem;
  margin-bottom: 0.5rem;
}

.label {
  display: flex;
  flex-direction: column;
  font-size: 1rem;
  gap: 0.25rem;
}

.label input {
  padding: 0.5rem;
  border: 2px solid #ccc;
  border-radius: 6px;
  font-size: 1rem;
  outline: none;
  transition: border-color 0.2s;
}
.label input:focus {
  border-color: #edbc91;
}

.btn {
  padding: 0.75rem;
  background: #fff;
  border: 2px solid #ccc;
  border-radius: 6px;
  font-size: 1rem;
  cursor: pointer;
  transition: background 0.2s, transform 0.1s;
}
.btn:hover {
  background: #edbc91;
  transform: translateY(-1px);
}
</style>

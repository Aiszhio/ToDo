<template>
  <div class="auth-container">
    <form @submit.prevent="handleSubmit" class="auth-form">
      <h2 class="title">Авторизация</h2>
      <p v-if="errorMessage" class="error">{{ errorMessage }}</p>

      <label class="field">
        Почта
        <input v-model="form.email" type="email" required />
      </label>

      <label class="field">
        Пароль
        <input v-model="form.password" type="password" required />
      </label>

      <button type="submit" class="btn">Войти</button>
    </form>
  </div>
</template>

<script setup>
import { reactive, ref } from 'vue'
import { useRouter } from 'vue-router'
import axios from 'axios'
import { jwtDecode } from 'jwt-decode'

const router = useRouter()
const form = reactive({ email: '', password: '' })
const errorMessage = ref('')

// единый axios‑клиент
const API_URL = import.meta.env.VITE_API_URL || 'http://localhost:5174'
const api = axios.create({ baseURL: API_URL })

function logout() {
  localStorage.removeItem('jwtToken')
  localStorage.removeItem('userName')
  router.push('/login')
}

async function handleSubmit() {
  errorMessage.value = ''
  try {
    // запрос на аутентификацию
    await api.post('/entrance', form)

    // второй запрос за JWT
    const { data } = await api.post('/receive-token')
    const token = data.token
    if (!token) throw new Error('no token')

    localStorage.setItem('jwtToken', token)

    // декодируем токен
    const decoded = jwtDecode(token)
    const uname = decoded.user_name || decoded.userName
    localStorage.setItem('userName', uname)

    // рассчитываем время до истечения
    const expMs = decoded.exp * 1000 - Date.now()
    if (expMs <= 0) {
      logout()
    } else {
      setTimeout(logout, expMs)
      router.push('/')
    }
  } catch (err) {
    if (err.response?.status === 401) {
      errorMessage.value = 'Неверный логин или пароль'
    } else {
      errorMessage.value = 'Ошибка сервера, попробуйте позже'
    }
  }
}
</script>

<style scoped>
.auth-container {
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 2rem;
  min-height: 80vh;
  background: #fff;
}

.auth-form {
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

.error {
  color: #c00;
  text-align: center;
  margin: 0;
}

.field {
  display: flex;
  flex-direction: column;
  font-size: 1rem;
  gap: 0.25rem;
}

.field input {
  padding: 0.5rem;
  border: 2px solid #ccc;
  border-radius: 6px;
  font-size: 1rem;
  outline: none;
  transition: border-color 0.2s;
}

.field input:focus {
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

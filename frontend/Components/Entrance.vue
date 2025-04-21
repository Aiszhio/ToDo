<template>
  <div class="layout-center">
    <form @submit.prevent="handleSubmit" class="card" style="max-width:400px; width:100%;">
      <h2 class="title-lg text-center mb-2">Авторизация</h2>

      <p v-if="errorMessage" class="error text-center mb-2">{{ errorMessage }}</p>

      <label class="field">
        Почта
        <input
            class="input"
            v-model="form.email"
            type="email"
            placeholder="example@mail.com"
            required
        />
      </label>

      <label class="field">
        Пароль
        <input
            class="input"
            v-model="form.password"
            type="password"
            placeholder="••••••••"
            required
        />
      </label>

      <button type="submit" class="btn btn-primary mb-1">
        Войти
      </button>
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
    await api.post('/entrance', form)

    const { data } = await api.post('/receive-token')
    const token = data.token
    if (!token) throw new Error('no token')

    localStorage.setItem('jwtToken', token)

    const decoded = jwtDecode(token)
    const uname = decoded.user_name || decoded.userName
    localStorage.setItem('userName', uname)

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
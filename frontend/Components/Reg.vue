<template>
  <div class="layout-center">
    <form @submit.prevent="onSubmit" class="card">
      <h2 class="title-lg text-center">Регистрация</h2>

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
        Имя пользователя
        <input
            class="input"
            v-model="form.name"
            type="text"
            placeholder="Ваш ник"
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

      <button type="submit" class="btn btn-primary">
        Зарегистрироваться
      </button>
    </form>
  </div>
</template>

<script setup>
import { reactive } from 'vue'
import { useRouter } from 'vue-router'
import axios from 'axios'

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
    form.email = form.name = form.password = ''
    router.push('/')
  } catch (e) {
    console.error('Registration error:', e)
    alert('Ошибка при регистрации. Попробуйте ещё раз.')
  }
}
</script>
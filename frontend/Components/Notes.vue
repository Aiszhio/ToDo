<template>
  <div class="layout-center">
    <form @submit.prevent="sendNote" class="card note-form">
      <p class="note-title title-md text-center">Напишите свою заметку</p>

      <textarea
          class="input"
          v-model="form.note"
          placeholder="Введите заметку…"
          maxlength="100"
          rows="3"
          required
      />

      <input
          class="input"
          type="date"
          v-model="form.date"
          :min="minDate"
          required
      />

      <button type="submit" class="btn btn-primary">
        Добавить
      </button>
    </form>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import axios from 'axios'

const token = localStorage.getItem('jwtToken')
const api = axios.create({
  baseURL: import.meta.env.VITE_API_URL || 'http://localhost:5174',
  headers: { Authorization: `Bearer ${token}` }
})

const form = ref({ note: '', date: '' })

const minDate = computed(() => new Date().toISOString().slice(0, 10))

async function sendNote() {
  if (!token) return alert('Токен не найден – залогинься заново')
  try {
    await api.post('/notes', form.value)
    alert('Заметка добавлена 🎉')
    form.value.note = ''
    form.value.date = ''
  } catch (err) {
    console.error('sendNote error:', err)
    alert('Ошибка при добавлении – проверь авторизацию')
  }
}
</script>
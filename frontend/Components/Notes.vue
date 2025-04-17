<template>
  <div id="noteBlock">
    <form @submit.prevent="sendNote" class="note-form">
      <p class="note-title">Напишите свою заметку</p>
      <textarea
          v-model="form.note"
          placeholder="Введите заметку…"
          maxlength="100"
          rows="3"
          required
      />
      <input
          type="date"
          v-model="form.date"
          :min="minDate"
          required
      />
      <button type="submit" class="submit-btn">Добавить</button>
    </form>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import axios from 'axios'

// 1) раз и навсегда подтягиваем токен и базовый URL
const token = localStorage.getItem('jwtToken')
const api = axios.create({
  baseURL: import.meta.env.VITE_API_URL || 'http://localhost:5174',
  headers: { Authorization: `Bearer ${token}` }
})

// 2) реактивная форма
const form = ref({ note: '', date: '' })

// 3) минимальная дата для <input type="date">
const minDate = computed(() => new Date().toISOString().slice(0, 10))

// 4) отправка заметки
async function sendNote() {
  if (!token) return alert('Токен не найден – залогинься заново')
  try {
    await api.post('/notes', form.value)
    alert('Заметка добавлена 🎉')
    // сбрасываем форму
    form.value.note = ''
    form.value.date = ''
  } catch (err) {
    console.error('sendNote error:', err)
    alert('Ошибка при добавлении – проверь авторизацию')
  }
}
</script>

<style scoped>
#noteBlock {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 100vh;
  padding: 2rem;
  background: linear-gradient(135deg, #f8f9fa 0%, #e9ecef 100%);
}

.note-form {
  width: 100%;
  max-width: 500px;
  padding: 2rem;
  background: rgba(255, 255, 255, 0.98);
  border-radius: 12px;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.05);
  backdrop-filter: blur(8px);
  animation: slideIn 0.4s ease-out;
}

.note-title {
  font-family: 'Poppins', sans-serif;
  font-size: 1.8rem;
  color: #2d3436;
  margin-bottom: 1.5rem;
  text-align: center;
}

textarea {
  width: 90%;
  max-width: 90%;
  min-width: 70%;
  padding: 1rem;
  border: 2px solid #e0e0e0;
  border-radius: 8px;
  font-family: inherit;
  font-size: 1rem;
  transition: all 0.3s ease;
  background: #f8f9fa;
}

textarea:focus {
  outline: none;
  border-color: #74b9ff;
  box-shadow: 0 0 0 3px rgba(116, 185, 255, 0.2);
}

input[type="date"] {
  padding: 0.8rem;
  border: 2px solid #e0e0e0;
  border-radius: 8px;
  font-family: inherit;
  font-size: 1rem;
  background: #f8f9fa;
  transition: all 0.3s ease;
}

input[type="date"]:focus {
  outline: none;
  border-color: #74b9ff;
  box-shadow: 0 0 0 3px rgba(116, 185, 255, 0.2);
}

.submit-btn {
  padding: 1rem;
  background: linear-gradient(135deg, #74b9ff 0%, #0984e3 100%);
  color: white;
  border: none;
  border-radius: 8px;
  font-family: 'Poppins', sans-serif;
  font-size: 1.1rem;
  font-weight: 500;
  cursor: pointer;
  transition: transform 0.2s ease, box-shadow 0.2s ease;
}

.submit-btn:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(116, 185, 255, 0.3);
}

.submit-btn:active {
  transform: translateY(0);
}

@media (max-width: 480px) {
  #noteBlock {
    padding: 1rem;
  }

  .note-form {
    padding: 1.5rem;
  }

  .note-title {
    font-size: 1.5rem;
  }
}

@keyframes slideIn {
  from {
    opacity: 0;
    transform: translateY(15px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}
</style>
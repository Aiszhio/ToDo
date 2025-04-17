<template>
  <div id="accountInf">
    <section class="account-section">
      <h2>Личный кабинет</h2>
      <div class="user-info">
        <h3>Привет, {{ userName }}!</h3>
      </div>

      <div class="notes-section">
        <div class="notes-block">
          <h3>Предстоящие заметки</h3>
          <ul>
            <li v-for="note in upcomingNotes" :key="note.id" class="note-item">
              <strong>{{ note.title }}</strong> — {{ formatDate(note.createdTo) }}
              <p>{{ note.description }}</p>
            </li>
          </ul>
        </div>

        <div class="notes-block">
          <h3>Недавние заметки</h3>
          <ul>
            <li v-for="note in expiredNotes" :key="note.id" class="note-item">
              <strong>{{ note.title }}</strong> — {{ formatDate(note.createdTo) }}
              <p>{{ note.description }}</p>
            </li>
          </ul>
        </div>
      </div>
    </section>
  </div>
</template>

<script>
import axios from 'axios'
import { throttle } from 'lodash'

export default {
  name: 'Account',
  data: () => ({
    userName: localStorage.getItem('userName') || '',
    token: localStorage.getItem('jwtToken') || '',
    notes: [],
    offset: 0,
    limit: 20,
    loading: false,
  }),
  computed: {
    upcomingNotes() {
      const now = Date.now()
      return this.notes.filter(n => new Date(n.createdTo).getTime() >= now)
    },
    expiredNotes() {
      const now = Date.now()
      const twoDaysAgo = now - 2 * 24 * 60 * 60 * 1000
      return this.notes.filter(n => {
        const t = new Date(n.createdTo).getTime()
        return t < now && t >= twoDaysAgo
      })
    }
  },
  created() {
    // единый axios‑клиент
    this.api = axios.create({
      baseURL: 'http://localhost:5174',
      headers: { Authorization: `Bearer ${this.token}` }
    })
  },
  async mounted() {
    if (!this.token || !this.userName) {
      console.error('No token or userName')
      return
    }
    await this.loadNotes()
    window.addEventListener('scroll', this.onScroll)
  },
  beforeDestroy() {
    window.removeEventListener('scroll', this.onScroll)
  },
  methods: {
    formatDate(d) {
      return new Date(d).toLocaleDateString(undefined, {
        year: 'numeric', month: 'long', day: 'numeric'
      })
    },
    async loadNotes() {
      if (this.loading) return
      this.loading = true
      try {
        const res = await this.api.post('/user-notes',
            { userName: this.userName },
            { params: { offset: this.offset, limit: this.limit } }
        )
        const newNotes = res.data.notes || []
        this.notes.push(...newNotes)
        this.offset += this.limit
      } catch (e) {
        console.error('loadNotes error:', e)
      } finally {
        this.loading = false
      }
    },
    onScroll: throttle(function() {
      if (window.innerHeight + window.scrollY >= document.body.offsetHeight - 100) {
        this.loadNotes()
      }
    }, 200)
  }
}
</script>

<style scoped>
#accountInf {
  display: flex;
  justify-content: center;
  padding: 2rem 1rem;
  min-height: 100vh;
  background: linear-gradient(135deg, #f8f9fa 0%, #e9ecef 100%);
}

.account-section {
  width: 100%;
  max-width: 1200px;
  padding: 2rem;
  background: rgba(255, 255, 255, 0.95);
  border-radius: 16px;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.1);
  backdrop-filter: blur(10px);
  animation: slideIn 0.6s cubic-bezier(0.23, 1, 0.32, 1);
}

h2 {
  font-family: 'Poppins', sans-serif;
  font-size: 2.5rem;
  color: #2d3436;
  text-align: center;
  margin-bottom: 2rem;
  position: relative;
}

h2::after {
  content: '';
  display: block;
  width: 60px;
  height: 3px;
  background: #ff7675;
  margin: 0.5rem auto;
}

.user-info h3 {
  font-size: 1.8rem;
  color: #636e72;
  margin-bottom: 2rem;
  text-align: center;
}

.notes-section {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: 2rem;
  margin-top: 2rem;
}

.notes-block {
  background: white;
  border-radius: 12px;
  padding: 1.5rem;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.05);
  transition: transform 0.3s ease, box-shadow 0.3s ease;
}

.notes-block:hover {
  transform: translateY(-5px);
  box-shadow: 0 8px 15px rgba(0, 0, 0, 0.1);
}

.notes-block h3 {
  font-size: 1.4rem;
  color: #2d3436;
  margin-bottom: 1.5rem;
  padding-bottom: 0.5rem;
  border-bottom: 2px solid #74b9ff;
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.note-item {
  padding: 1rem;
  margin-bottom: 1rem;
  background: #f8f9fa;
  border-radius: 8px;
  transition: all 0.2s ease;
  cursor: pointer;
}

.note-item:hover {
  background: white;
  box-shadow: 0 2px 8px rgba(116, 185, 255, 0.15);
}

.note-item strong {
  font-size: 1.1rem;
  color: #2d3436;
  display: block;
  margin-bottom: 0.3rem;
}

.note-item p {
  font-size: 0.95rem;
  color: #636e72;
  line-height: 1.5;
  margin-top: 0.5rem;
}

/* Анимации */
@keyframes slideIn {
  from {
    opacity: 0;
    transform: translateY(20px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

/* Адаптивность */
@media (max-width: 768px) {
  .account-section {
    padding: 1.5rem;
    border-radius: 12px;
  }

  h2 {
    font-size: 2rem;
  }

  .notes-section {
    grid-template-columns: 1fr;
  }
}

/* Кастомный скролл */
.notes-block::-webkit-scrollbar {
  width: 6px;
}

.notes-block::-webkit-scrollbar-track {
  background: #f1f1f1;
  border-radius: 3px;
}

.notes-block::-webkit-scrollbar-thumb {
  background: #74b9ff;
  border-radius: 3px;
}

.notes-block::-webkit-scrollbar-thumb:hover {
  background: #0984e3;
}
</style>
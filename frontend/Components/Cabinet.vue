<template>
  <div class="layout-center">
    <section class="card w-100">
      <h2 class="title-lg text-center mb-2">Личный кабинет</h2>
      <p class="title-md text-center mb-2">Привет, {{ userName }}!</p>

      <div class="notes-grid mb-2">
        <div>
          <h3 class="title-sm mb-1">Предстоящие заметки</h3>
          <ul>
            <li
                v-for="note in upcomingNotes"
                :key="note.id"
                class="note-card"
            >
              <div>
                <strong>{{ note.title }}</strong>
                <p class="mb-1 text-sm">{{ formatDate(note.createdTo) }}</p>
                <p class="text-sm">{{ note.description }}</p>
              </div>
              <div class="note-actions">
                <button @click="onEdit(note)">✏️</button>
                <button @click="onDelete(note.id)">🗑️</button>
              </div>
            </li>
          </ul>
        </div>

        <div>
          <h3 class="title-sm mb-1">Недавние заметки</h3>
          <ul>
            <li
                v-for="note in expiredNotes"
                :key="note.id"
                class="note-card"
            >
              <div>
                <strong>{{ note.title }}</strong>
                <p class="mb-1 text-sm">{{ formatDate(note.createdTo) }}</p>
                <p class="text-sm">{{ note.description }}</p>
              </div>
              <div class="note-actions">
                <button @click="onEdit(note)">✏️</button>
                <button @click="onDelete(note.id)">🗑️</button>
              </div>
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
    this.api = axios.create({
      baseURL: 'http://localhost:5174',
      headers: { Authorization: `Bearer ${this.token}` }
    })
  },
  async mounted() {
    if (!this.token || !this.userName) return
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
        const res = await this.api.post(
            '/user-notes',
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
    }, 200),
    async onEdit(note) {
      const newTitle = prompt('Новая заметка:', note.title)
      if (newTitle == null) return
      const newDate = prompt('Новая дата (YYYY-MM-DD):', note.createdTo.slice(0,10))
      if (!newDate) return
      try {
        await this.api.put('/put-notes', {
          id: note.id,
          note: newTitle,
          date: newDate
        })
        note.title = newTitle
        note.createdTo = newDate
        alert('Заметка обновлена!')
      } catch (e) {
        console.error('edit error:', e)
        alert('Не получилось обновить.')
      }
    },
    async onDelete(id) {
      if (!confirm('Точно удалить?')) return
      try {
        await this.api.delete('/delete-notes', { data: { id } })
        this.notes = this.notes.filter(n => n.id !== id)
        alert('Заметка удалена!')
      } catch (e) {
        console.error('delete error:', e)
        alert('Не получилось удалить.')
      }
    }
  }
}
</script>
<script>
import axios from "axios";

export default {
  name: 'Account',
  data() {
    return {
      userName: '',        // Имя пользователя из localStorage
      notes: [],           // Все заметки пользователя
      upcomingNotes: [],   // Отфильтрованные предстоящие заметки
      expiredNotes: [],    // Отфильтрованные истекшие заметки
      offset: 0,           // Начальное смещение
      limit: 20,           // Количество заметок для загрузки
    };
  },
  async created() {
    try {
      this.userName = localStorage.getItem('userName');
      const token = localStorage.getItem('jwtToken');

      if (token && this.userName) {
        await this.loadNotes(); // Загружаем заметки при создании компонента
      } else {
        console.error("Token or username not found!");
      }
    } catch (error) {
      console.error("Ошибка при получении данных пользователя: ", error);
    }
  },
  mounted() {
    window.addEventListener('scroll', this.handleScroll);
  },
  beforeDestroy() {
    window.removeEventListener('scroll', this.handleScroll);
  },
  methods: {
    formatDate(dateString) {
      const options = { year: 'numeric', month: 'long', day: 'numeric' };
      return new Date(dateString).toLocaleDateString(undefined, options);
    },
    async loadNotes() {
      const token = localStorage.getItem('jwtToken');
      try {
        const response = await axios.post(
            'http://localhost:5174/user-notes',
            { userName: this.userName },
            {
              headers: {
                Authorization: `Bearer ${token}`
              },
              params: {
                limit: this.limit,
                offset: this.offset
              }
            }
        );

        const newNotes = response.data.notes;
        this.notes = [...this.notes, ...newNotes]; // Объединяем с уже загруженными заметками

        const currentDate = new Date();
        const twoDaysAgo = new Date(currentDate);
        twoDaysAgo.setDate(currentDate.getDate() - 2);

        // Фильтрация предстоящих и истекших заметок
        this.upcomingNotes = this.notes.filter(note => new Date(note.createdTo) >= currentDate);
        this.expiredNotes = this.notes.filter(note => {
          const noteDate = new Date(note.createdTo);
          return noteDate <= currentDate && noteDate >= twoDaysAgo;
        });

        this.offset += this.limit; // Увеличиваем смещение для следующего запроса
      } catch (error) {
        console.error("Ошибка при загрузке заметок: ", error);
      }
    },
    handleScroll() {
      const scrollTop = window.scrollY || document.documentElement.scrollTop;
      const windowHeight = window.innerHeight;
      const documentHeight = document.documentElement.offsetHeight;

      console.log("Scroll position:", scrollTop);
      console.log("Window height:", windowHeight);
      console.log("Document height:", documentHeight);

      if (scrollTop + windowHeight >= documentHeight - 100) {
        console.log("Loading more notes...");
        this.loadNotes(); // Загружаем больше заметок, когда достигнут низ страницы
      }
    },
  }
};
</script>

<template>
  <div id="accountInf">
    <section class="account-section">
      <h2>Личный кабинет</h2>
      <div class="user-info">
        <h3>Привет, {{ userName }}!</h3>
      </div>

      <div class="notes-section">
        <div class="upcoming-notes">
          <h3>Предстоящие заметки</h3>
          <ul>
            <li v-for="note in upcomingNotes" :key="note.id" class="note-item">
              <strong>{{ note.title }}</strong> — {{ formatDate(note.createdTo) }}
              <p>{{ note.description }}</p>
            </li>
          </ul>
        </div>

        <div class="expired-notes">
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

<style scoped>
#accountInf {
  display: flex;
  justify-content: center;
  margin-top: 2%;
  background-color: #fff;
  height: 85vh;
}

.account-section {
  width: 80%;
  max-width: 900px;
  height: 90%;
  padding: 20px;
  background-color: #FFDAB9;
  border-radius: 10px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
}

h2 {
  text-align: center;
  color: #333;
  margin-bottom: 20px;
  font-size: 2em;
}

.user-info {
  text-align: center;
  margin-bottom: 20px;
}

h3 {
  font-size: 1.5em;
  color: #555;
}

.notes-section {
  display: flex;
  justify-content: space-between;
  gap: 20px;
}

.upcoming-notes, .expired-notes {
  width: 48%;
  height: 50vh;
  background-color: #fff;
  border-radius: 8px;
  padding: 15px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.05);
  max-height: 80vh;
  overflow-y: auto;
}

.note-item {
  margin-bottom: 10px;
  padding: 10px;
  border-bottom: 1px solid #FFDAB9;
}

.note-item strong {
  color: #333;
  font-size: 1.2em;
}

.note-item p {
  color: #666;
  font-size: 0.9em;
}
</style>

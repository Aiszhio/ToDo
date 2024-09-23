<script>
import axios from "axios";

export default {
  name: "notes",
  data() {
    return {
      form: {
        note: '',
        date: '',
      }
    };
  },
  methods: {
    async sendNote() {
      try {
        const response = await axios.post('http://localhost:5174/notes', this.form);

        if (response.status === 200) {
          alert('Ваша заметка была добавлена в базу!');
        }
      } catch (error) {
        console.error('Ошибка при отправке заметки:', error);
        alert('Не удалось отправить заметку на сервер. Попробуйте позже.');
      }
    }
  },
  computed: {
    minDate() {
      const today = new Date();
      const year = today.getFullYear();
      const month = String(today.getMonth() + 1).padStart(2, '0');
      const day = String(today.getDate()).padStart(2, '0');
      return `${year}-${month}-${day}`;
    }
  }
};
</script>

<template>
  <div id="noteBlock">
    <form @submit.prevent="sendNote" class="NoteForm">
      <p class="NoteName">Напишите свою заметку</p>
      <textarea v-model="form.note" id="note" name="note" maxlength="100" rows="3" required placeholder="Введите заметку..."></textarea><br>
      <input type="date" v-model="form.date" id="notedate" name="notedate" required :min="minDate"><br>
      <input type="submit" value="Добавить" class="submitnote">
    </form>
  </div>
</template>

<style scoped>
#noteBlock {
  display: flex;
  flex-direction: row;
  justify-content: center;
  align-items: center;
  height: 80vh;
}
.NoteForm {
  display: flex;
  flex-direction: column;
  align-items: center;
  border: 3px solid black;
  border-radius: 20px;
  padding: 35px;
}
.NoteName {
  font-family: Arial;
  font-size: x-large;
  margin-bottom: 2vh;
}
#note {
  box-sizing: border-box;
  margin: 1vh;
  padding-left: 1vh;
  height: 15vh;
  width: 50vh;
  border: 3px solid #ccc;
  border-radius: 10px;
  font-family: Arial;
  font-size: larger;
  overflow-wrap: break-word;
  resize: horizontal;
  min-width: 50vh;
  max-width: 90vh;
}
#notedate {
  padding: 10px;
  border: 1px solid #ccc;
  border-radius: 5px;
  font-size: 16px;
  transition: border-color 0.3s;
}
.submitnote {
  font-family: Arial;
  font-size: larger;
  margin-top: 2vh;
  padding: 1vh;
  width: 80%;
  border: 2px solid #ccc;
  border-radius: 5px;
  background-color: #FFDAB9;
}
</style>

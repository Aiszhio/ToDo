<script>
export default {
  name: 'home',
  data() {
    return {
      userName: '',
    };
  },
  mounted() {
    const storedUserName = localStorage.getItem('userName');
    if (storedUserName) {
      this.userName = storedUserName;
    }
  },
  methods: {
    loadCabinet() {
      this.$router.push("cabinet");
    },
    loadNotes() {
      this.$router.push("notes");
    },
    loadHelp() {
      this.$router.push("help");
    },
  },
};
</script>

<template>
  <main class="layout-center layout-home">
    <aside class="aside">
      <nav class="nav-list">
        <p class="title-md navText">Меню</p>
        <hr class="underNav" />

        <button class="nav-link" @click="loadCabinet">
          Личный кабинет
        </button>
        <button class="nav-link" @click="loadNotes">
          Заметки
        </button>
        <button class="nav-link" @click="loadHelp">
          Помощь
        </button>
      </nav>
    </aside>

    <div>
      <p class="greet">
        Добро пожаловать, {{ userName || 'гость' }}!
      </p>
    </div>
  </main>
</template>

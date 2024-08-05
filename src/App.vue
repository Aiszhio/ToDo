<script>

export default {
  data() {
    return {
      message: ''
    }
  },
  methods: {
    async fetchMessage() {
      try {
        const response = await fetch('http://localhost:8080/registration');
        if (response.ok) {
          this.message = await response.text();
        } else {
          console.error('Failed to fetch message');
        }
      } catch (error) {
        console.error('Error:', error);
      }
    }
  }
}

</script>

<template>
  <main>
    <div id="asideInf">
      <nav>
        <p>Text 1</p>
      </nav>
    </div>
    <div id="buttons">
      <button class="but1" onclick="">
        Регистрация
      </button>
      <button class="but2">
        Вход
      </button>
    </div>
  </main>
</template>

<style>
main {
  display: flex;
  flex-direction: row;
  height: 100vh;
  margin: 0;
}

#asideInf {
  border: 2px solid rgb(110,110,110);
  border-radius: 10px;
  background-color: rgb(45,45,45);
  width: 15%;
  display: flex;
  box-sizing: border-box;
  margin: 0;
}

#buttons{
  margin-left: auto;
  margin-right: 1vh;
  margin-top: 1vh;
}

button {
  height: 5vh;
  width: 25vh;
  user-select: none;
  cursor: pointer;
}

.but1{
  margin-right: 1vh;
  background-color: rgb(120,120,120);
  color: white;
  border-radius: 5px;
}

.but1:hover{
  opacity: .8;
  text-decoration: none;
}

.but2{
  background-color: rgb(120,120,120);
  color: white;
  border-radius: 5px;
}

.but2:hover{
  opacity: .8;
  text-decoration: none;
}

h1, p {
  color: #F5F5F5;
}

html, body {
  margin: 0;
  height: 100%;
  background-color: #1C1C1C;
}

</style>

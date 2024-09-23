
<script>

export default {
  name: 'home',
  methods: {
    loadCabinet(){
      this.$router.push("cabinet")
    },
    loadNotes(){
      this.$router.push("notes")
    },
    loadHelp(){
      this.$router.push("help")
    }
  }
}

</script>

<template>
  <main>
    <div id="asideInf">
      <nav id="asideBut">
        <p class="navText">Меню</p>
        <hr class = "underNav">
        <button class="navBut" @click="loadCabinet"> Личный кабинет </button>
        <button class="navBut" @click="loadNotes"> Заметки </button>
        <button class="navBut" @click="loadHelp"> Помощь </button>
      </nav>
    </div>

    <p class="greet">Добро пожаловать!</p>

  </main>
</template>

<style scoped>
#asideInf {
  border: 2px solid white;
  border-radius: 10px;
  background-color: #FFDAB9;
  width: 15%;
  height: 55vh;
  margin-top: 10vh;
  margin-left: 1vh;
  padding-top: 1%;
}

#asideBut{
  height: 80%;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: space-evenly;
}

.navText{
  font-family: "Roboto Light", Arial, sans-serif;
  font-size: 3vh;
  margin: 0;
  color: black;
  font-weight: bold;
}

.underNav{
  width: 80%;
  margin: 0;
  background-color: black;
}

.navBut{
  height: 100%;
  width: 98%;
  color: black;
  font-size: 18px;
  background-color: #FFDAB9;
  border: white;
  border-radius: 5vh;
  margin-top: 4vh;
}

  .navBut:hover{
    opacity: .8;
    text-decoration: none;
    background-color: #e3a95e;
    font-weight: bolder;
  }

.greet{
  color: black;
  align-content: center;
  margin-left: 25%;
  font-family: "Roboto Light", Arial, sans-serif;
  font-size: xxx-large;
}
</style>
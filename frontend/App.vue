
<script>

import home from './Components/HomePage.vue';

export default {
  components: {
    home,
  },
  methods:{
    loadReg(){
      this.$router.push('/register')
    },
    loadHome(){
      this.$router.push('/')
    },
    loadEnt(){
      this.$router.push('/entrance')
    }
  }
};

</script>

<template>
  <div id="app">
    <button class="but1" @click="loadReg()">
      Регистрация
    </button>
    <button class="but2" @click="loadEnt()">
      Вход
    </button>
    <button class="but3" @click="loadHome()">
      Главная
    </button>

  </div>
  <router-view v-slot="{ Component }">
    <transition name="fade" mode="out-in">
      <component :is="Component" />
    </transition>
  </router-view>
</template>


<style scoped>
#app{
  display: flex;
  justify-content: right;
  background-color: #FFDAB9;
}
</style>

<style>
fade-enter-active,
.fade-leave-active {
  transition: opacity 0.4s;
}
.fade-enter, .fade-leave-to {
  opacity: 0;
}
</style>
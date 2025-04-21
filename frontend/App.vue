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
    <header class="header">
      <button class="btn btn-primary" @click="loadReg">Регистрация</button>
      <button class="btn btn-primary" @click="loadEnt">Вход</button>
      <button class="btn btn-primary" @click="loadHome">Главная</button>
    </header>

    <transition name="fade" mode="out-in">
      <router-view class="layout-center" />
    </transition>
  </div>
</template>


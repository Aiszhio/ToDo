
<script>

import home from './Components/HomePage.vue';
import register from './Components/Reg.vue';

export default {
  components: {
    home,
    register,
  },
  methods:{
    loadReg(){
      this.$router.push('/register')
    },
    loadHome(){
      this.$router.push('/')
    }
  }
};

</script>

<template>
  <div id="app">
    <button class="but1" @click="loadReg()">
      Регистрация
    </button>
    <button class="but2">
      Вход
    </button>
    <button class="but3" @click="loadHome()">
      Главная
    </button>
    <router-view></router-view>
  </div>
</template>

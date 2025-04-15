<script setup lang="ts">
import { RouterView } from 'vue-router';
import LoginForm from './components/form/LoginForm.vue';
</script>

<template>
  <header>
    <div class="flex justify-between items-center p-4">
      <h1 class="text-2xl font-bold">Home</h1>


          <LoginForm />

    </div>
  </header>
  <RouterView />
</template>

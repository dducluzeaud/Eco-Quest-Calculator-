<script setup lang="ts">
import LoginForm from '@/components/form/LoginForm.vue'
import EcoIcon from '@/components/icons/EcoIcon.vue'
import { Toaster } from '@/components/ui/sonner'
import { RouterView } from 'vue-router'
</script>

<template>
  <header>
    <div class="flex justify-between items-center p-4">
      <h1 class="text-2xl font-bold">
        <EcoIcon class="w-40 h-10" />
      </h1>
      <LoginForm />
    </div>
  </header>
  <Toaster />
  <RouterView />
</template>

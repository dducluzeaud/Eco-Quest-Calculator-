<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
import { Avatar, AvatarFallback, AvatarImage } from '@/components/ui/avatar'
</script>

<template>
  <header>
    <div class="flex justify-between items-center p-4">
      <h1 class="text-2xl font-bold">Home</h1>
      <Avatar>
        <AvatarImage src="https://github.com/unovue.png" alt="@unovue" />
        <AvatarFallback>CN</AvatarFallback>
      </Avatar>
    </div>
  </header>
  <RouterView />
</template>

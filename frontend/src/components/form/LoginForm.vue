<script lang="ts">
export const description =
  "A login form with email and password. There's an option to login with Google and a link to sign up if you don't have an account."
export const iframeHeight = '600px'
export const containerClass = 'w-full h-screen flex items-center justify-center px-4'
</script>

<script setup lang="ts">
import { Avatar, AvatarFallback } from '@/components/ui/avatar'
import { Button } from '@/components/ui/button'
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogTitle,
  DialogTrigger,
} from '@/components/ui/dialog'
import { FormField, FormItem, FormLabel, FormMessage } from '@/components/ui/form'
import { Input } from '@/components/ui/input'
import { useAuthStore } from '@/stores/auth'
import { toTypedSchema } from '@vee-validate/zod'
import { Eye, EyeOff, User as UserIcon } from 'lucide-vue-next'
import { useForm } from 'vee-validate'
import { ref } from 'vue'
import { toast } from 'vue-sonner'
import * as z from 'zod'

const authStore = useAuthStore()

const formSchema = toTypedSchema(
  z.object({
    email: z.string().email(),
    password: z.string().min(6).max(50),
  }),
)

const form = useForm({
  validationSchema: formSchema,
  initialValues: {
    email: '',
    password: '',
  },

})

const open = ref(false)
const showPassword = ref(false)

function togglePasswordVisibility()  {
  showPassword.value = !showPassword.value
}

function closeAndNavigateToSignup() {
  open.value = false
}



const onSubmit = form.handleSubmit(async (values) => {


  try {
    await authStore.login(values)


    toast.success('Login successful')
  } catch (error) {
    console.error('Login error', error)
    toast.error('Login failed')
    return
  }


    open.value = false

})
</script>

<template>
  <Dialog v-model:open="open">
    <DialogTrigger as-child>
      <Avatar>
        <AvatarFallback>
          <div class="uppercase" v-if="authStore.isAuthenticated">{{ authStore.getUsername.slice(0, 2) }}</div>
          <div v-else>
            <UserIcon />
          </div>
        </AvatarFallback>
      </Avatar>
    </DialogTrigger>

    <DialogContent>
      <div v-if="authStore.isAuthenticated" class="text-sm text-center">
        {{ authStore.getUsername}}
        <Button variant="link" class="ml-1" @click="authStore.logout">Logout</Button>
      </div>
      <div v-else class="text-sm text-center">
        <DialogTitle>Login</DialogTitle>
        <DialogDescription>Enter your email below to login to your account </DialogDescription>
        <form @submit="onSubmit" class="grid gap-4">
          <FormField class="grid gap-2" v-slot="{ componentField }" name="email">
            <FormItem>
              <FormLabel>Email</FormLabel>
              <Input
                id="email"
                type="email"
                autocomplete="username"
                placeholder="m@example.com"
                required
                v-bind="componentField"
                autofocus
              />
              <FormMessage />
            </FormItem>
          </FormField>
          <FormField class="grid gap-2" v-slot="{ componentField }" name="password">
            <FormItem>
              <div class="flex items-center">
                <FormLabel>Password</FormLabel>
                <a href="#" class="ml-auto inline-block text-sm underline">
                  Forgot your password?
                </a>
              </div>
              <div class="relative">
                <Input
                id="password"
                autocomplete="current-password"
                required
                v-bind="componentField"
                :type="showPassword ? 'text' : 'password'"
                />
                <button
                type="button"
                class="absolute right-2 top-1/2 transform -translate-y-1/2"
                @click="togglePasswordVisibility"
                >
                <EyeOff v-if="showPassword" class="h-4 w-4" />
                <Eye v-else class="h-4 w-4" />
              </button>
            </div>

              <FormMessage />
            </FormItem>
          </FormField>
          <Button type="submit" class="w-full">Login</Button>
        </form>
        <div class="mt-4 text-center text-sm">
          Don't have an account?
          <Button variant="link" class="ml-1" @click="closeAndNavigateToSignup">Sign up</Button>
        </div>
      </div>
    </DialogContent>
  </Dialog>
</template>

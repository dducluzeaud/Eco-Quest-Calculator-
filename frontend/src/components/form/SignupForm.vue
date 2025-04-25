<script lang="ts">
export const description = 'A signup form with email, username and password validation.'
export const containerClass = 'w-full h-screen flex items-center justify-center px-4'
</script>

<script setup lang="ts">
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
import { Eye, EyeOff } from 'lucide-vue-next'
import { useForm } from 'vee-validate'
import { ref } from 'vue'
import { toast } from 'vue-sonner'
import * as z from 'zod'

const authStore = useAuthStore()

const formSchema = toTypedSchema(
  z
    .object({
      username: z.string().min(3).max(50),
      email: z.string().email(),
      password: z.string().min(6).max(50),
      confirmPassword: z.string(),
    })
    .refine((data) => data.password === data.confirmPassword, {
      message: "Passwords don't match",
      path: ['confirmPassword'],
    }),
)

const form = useForm({
  validationSchema: formSchema,
  initialValues: {
    username: '',
    email: '',
    password: '',
    confirmPassword: '',
  },
})

const open = ref(false)
const showPassword = ref(false)
const showConfirmPassword = ref(false)

function togglePasswordVisibility() {
  showPassword.value = !showPassword.value
}

function toggleConfirmPasswordVisibility() {
  showConfirmPassword.value = !showConfirmPassword.value
}

function closeAndNavigateToLogin() {
  open.value = false
}

const onSubmit = form.handleSubmit(async (values) => {
  try {
    await authStore.register({
      username: values.username,
      email: values.email,
      password: values.password,
    })

    toast.success('Registration successful')
    open.value = false
  } catch (error) {
    console.error('Registration error:', error)
    toast.error('Registration failed')
  }
})
</script>

<template>
  <Dialog v-model:open="open">
    <DialogTrigger asChild>
      <Button variant="outline">Sign up</Button>
    </DialogTrigger>

    <DialogContent>
      <div class="text-sm text-center">
        <DialogTitle>Create an account</DialogTitle>
        <DialogDescription> Enter your details below to create your account </DialogDescription>
        <form @submit="onSubmit" class="grid gap-4">
          <FormField class="grid gap-2" v-slot="{ componentField }" name="username">
            <FormItem>
              <FormLabel>Username</FormLabel>
              <Input
                id="username"
                type="text"
                autocomplete="username"
                required
                v-bind="componentField"
                autofocus
              />
              <FormMessage />
            </FormItem>
          </FormField>

          <FormField class="grid gap-2" v-slot="{ componentField }" name="email">
            <FormItem>
              <FormLabel>Email</FormLabel>
              <Input
                id="email"
                type="email"
                autocomplete="email"
                required
                v-bind="componentField"
              />
              <FormMessage />
            </FormItem>
          </FormField>

          <FormField class="grid gap-2" v-slot="{ componentField }" name="password">
            <FormItem>
              <FormLabel>Password</FormLabel>
              <div class="relative">
                <Input
                  id="password"
                  autocomplete="new-password"
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

          <FormField class="grid gap-2" v-slot="{ componentField }" name="confirmPassword">
            <FormItem>
              <FormLabel>Confirm Password</FormLabel>
              <div class="relative">
                <Input
                  id="confirmPassword"
                  autocomplete="new-password"
                  required
                  v-bind="componentField"
                  :type="showConfirmPassword ? 'text' : 'password'"
                />
                <button
                  type="button"
                  class="absolute right-2 top-1/2 transform -translate-y-1/2"
                  @click="toggleConfirmPasswordVisibility"
                >
                  <EyeOff v-if="showConfirmPassword" class="h-4 w-4" />
                  <Eye v-else class="h-4 w-4" />
                </button>
              </div>
              <FormMessage />
            </FormItem>
          </FormField>

          <Button type="submit" class="w-full">Create account</Button>
        </form>
        <div class="mt-4 text-center text-sm">
          Already have an account?
          <Button variant="link" class="ml-1" @click="closeAndNavigateToLogin"> Login </Button>
        </div>
      </div>
    </DialogContent>
  </Dialog>
</template>

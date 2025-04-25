// auth.js

import auth, { type LoginCredentials } from '@/api/auth'
import { acceptHMRUpdate, defineStore } from 'pinia'
import { z } from 'zod'

const UserSchema = z.object({
  id: z.number(),
  username: z.string(),
  email: z.string().email(),
})

const AuthStoreSchema = z.object({
  user: UserSchema.nullable(),
  accessToken: z.string().nullable(),
  refreshToken: z.string().nullable(),
})

const AuthStoreValidation = z.object({
  user: UserSchema,
  accessToken: z.string().min(1),
  refreshToken: z.string().min(1),
})

export const useAuthStore = defineStore('auth', {
  state: (): z.infer<typeof AuthStoreSchema> => ({
    user: null,
    accessToken: null,
    refreshToken: null,
  }),
  getters: {
    isAuthenticated(): boolean {
      return !!this.user
    },
    getUsername(): string {
      if (!this.user?.username) {
        throw new Error('User name is not available')
      }
      return this.user.username
    },
    getEmail(): string {
      if (!this.user?.email) {
        throw new Error('User name is not available')
      }
      return this.user.email
    },
  },
  actions: {
    async login(credentials: LoginCredentials) {
      try {
        const response = await auth.login(credentials)

        console.log('Login response:', response)
        const validated = AuthStoreValidation.parse(response.data)
        console.log('Validated response:', validated)

        this.user = validated.user
        this.accessToken = validated.accessToken
        this.refreshToken = validated.refreshToken

        return true
      } catch (error) {
        if (error instanceof z.ZodError) {
          console.error('Validation error:', error.errors)
        }
        console.error('Login error:', error)
        throw error
      }
    },
    logout() {
      this.user = null
      this.accessToken = null
      this.refreshToken = null
    },
  },
  persist: true,
})

// make sure to pass the right store definition, `useAuth` in this case.
if (import.meta.hot) {
  import.meta.hot.accept(acceptHMRUpdate(useAuthStore, import.meta.hot))
}

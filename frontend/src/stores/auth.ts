// auth.js
import { acceptHMRUpdate, defineStore } from 'pinia'

type AuthStore =
  | {
      user: {
        id: number
        name: string
        email: string
        role: string
      }
      token: string
      refreshToken: string
    }
  | {
      user: null
      token: null
      refreshToken: null
    }

export const useAuthStore = defineStore('auth', {
  state: (): AuthStore => ({
    user: null,
    token: null,
    refreshToken: null,
  }),
  getters: {
    isAuthenticated(state) {
      return !!state.user
    },
    getUser(state) {
      return state.user
    },
    getToken(state) {
      return state.token
    },
  },
  actions: {
    setAuh(auth: AuthStore) {
      this.user = auth.user
      this.token = auth.token
    },
    clearAuth() {
      this.user = null
      this.token = null
    },
  },
  persist: true,
})

// make sure to pass the right store definition, `useAuth` in this case.
if (import.meta.hot) {
  import.meta.hot.accept(acceptHMRUpdate(useAuthStore, import.meta.hot))
}

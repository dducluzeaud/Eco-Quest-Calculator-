// auth.js
import { acceptHMRUpdate, defineStore } from 'pinia'

type AuthStore = {
  user: null | {
    id: number
    name: string
    email: string
    role: string
  }
  token: null | string
  refreshToken: null | string
}

export const useAuth = defineStore('auth', {
  state: (): AuthStore => ({
    user: null,
    token: null,
    refreshToken: null,
  }),
  getters: {
    isAuthenticated(state) {
      return !!state.token
    },
  },
  actions: {
    setAuh(user: null) {
      this.user = user
    },
    setToken(token: null) {
      this.token = token
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
  import.meta.hot.accept(acceptHMRUpdate(useAuth, import.meta.hot))
}

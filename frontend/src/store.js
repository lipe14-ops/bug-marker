
import { defineStore } from 'pinia'

export const useTokenStore = defineStore('token', {
  state: () => {
    return { token: '' }
  },
  actions: {
    setToken(newToken) {
      this.token = newToken
    },
  },
})

import ky from 'ky'
import auth from './auth'

let refreshTokenPromise: Promise<unknown> | null = null

const api = ky.create({
  prefixUrl: import.meta.env.VITE_API_URL,
  headers: {
    'Content-Type': 'application/json',
  },
  hooks: {
    beforeRequest: [
      (request) => {
        const token = localStorage.getItem('token')
        if (token) {
          request.headers.set('Authorization', `Bearer ${token}`)
        }
      },
    ],
    afterResponse: [
      async (request, _options, response) => {
        if (response.status === 401) {
          refreshTokenPromise ??= auth.refreshToken().finally(() => {
            refreshTokenPromise = null
          })
          await refreshTokenPromise
          const newToken = localStorage.getItem('accessToken')
          request.headers.set('Authorization', `Bearer ${newToken}`)
          return ky(request)
        }
        return response
      },
    ],
  },
})

export default api

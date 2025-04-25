import ky from 'ky'

type LoginResponse = {
  data: {
    accessToken: string
    refreshToken: string
    user: {
      id: string
      username: string
      email: string
    }
  }
}

const unconnectedApi = ky.create({
  prefixUrl: import.meta.env.VITE_API_URL,
  headers: {
    'Content-Type': 'application/json',
  },
})

export interface LoginCredentials {
  email: string
  password: string
}

function login(credentials: LoginCredentials) {
  return unconnectedApi
    .post<LoginResponse>('api/auth/login', {
      json: credentials,
    })
    .json()
}

function refreshToken() {
  return ky
    .post('/api/refresh-token', {
      headers: {
        Authorization: `Bearer ${localStorage.getItem('refreshToken')}`,
      },
    })
    .json()
}

export default { login, refreshToken }

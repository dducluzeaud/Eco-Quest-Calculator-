import type { LoginCredentials } from '@/api/auth'
import auth from '@/api/auth'
import { defineMutation } from '@pinia/colada'

export const useLoginMutation = defineMutation(() => ({
  mutation: (credentials: LoginCredentials) => auth.login(credentials),
}))

import type { ApiError, ApiResponse, User, Video } from '@/types'

const API_BASE = import.meta.env.VITE_API_BASE ?? '/api/v1'

async function request<T>(
  path: string,
  options: RequestInit = {},
): Promise<ApiResponse<T>> {
  const res = await fetch(`${API_BASE}${path}`, {
    credentials: 'include',
    headers: {
      ...(options.body instanceof FormData
        ? {}
        : { 'Content-Type': 'application/json' }),
      ...options.headers,
    },
    ...options,
  })

  const data = (await res.json().catch(() => ({}))) as ApiResponse<T> & ApiError

  if (!res.ok) {
    throw new Error(data.error ?? data.message ?? 'Something went wrong')
  }

  return data
}

export const authApi = {
  register: (body: {
    username: string
    name: string
    email: string
    password: string
  }) =>
    request<{ user: User }>('/auth/register', {
      method: 'POST',
      body: JSON.stringify(body),
    }),

  login: (body: { email: string; password: string }) =>
    request('/auth/login', {
      method: 'POST',
      body: JSON.stringify(body),
    }),

  logout: () =>
    request('/auth/logout', { method: 'POST' }),

  verifyToken: () =>
    request('/auth/verifytoken', { method: 'POST' }),
}

export const userApi = {
  getCurrentUser: () =>
    request<{user: User }>('/user/current-user'),

  changePassword: (new_password: string) =>
    request('/user/change-password', {
      method: 'POST',
      body: JSON.stringify({ new_password }),
    }),

  updateAccount: (form: FormData) =>
    request<{user: User }>('/user/update-account', {
      method: 'POST',
      body: form,
    }),
}

export const videoApi = {
  getAll: () =>
    request<{videos: Video[] }>('/video/all'),
}

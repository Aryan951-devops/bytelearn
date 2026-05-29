import type {
  ApiError,
  ApiResponse,
  Course,
  CoursePlaylist,
  CourseWithPlaylists,
  Playlist,
  PlaylistWithVideos,
  User,
  Video,
} from '@/types'

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

  logout: () => request('/auth/logout', { method: 'POST' }),

  verifyToken: () => request('/auth/verifytoken', { method: 'POST' }),
}

export const userApi = {
  getCurrentUser: () => request<{ user: User }>('/user/current-user'),

  changePassword: (new_password: string) =>
    request('/user/change-password', {
      method: 'POST',
      body: JSON.stringify({ new_password }),
    }),

  updateAccount: (form: FormData) =>
    request<{ user: User }>('/user/update-account', {
      method: 'PATCH',
      body: form,
    }),
}

export const courseApi = {
  getAll: () => request<{ courses: Course[] }>('/course/all'),

  getById: (courseId: string) =>
    request<{ course: CourseWithPlaylists }>(`/course/${courseId}`),

  create: (body: {
    title: string
    description?: string
    category?: string
  }) =>
    request<{ course: Course }>('/course/', {
      method: 'POST',
      body: JSON.stringify(body),
    }),

  update: (
    courseId: string,
    body: {
      title?: string
      description?: string
      category?: string
    },
  ) =>
    request<{ course: Course }>(`/course/${courseId}`, {
      method: 'PATCH',
      body: JSON.stringify(body),
    }),
}

export const playlistApi = {
  create: (body: {
    type: string
    title: string
    description?: string
    course_id?: string
  }) =>
    request<{ playlist: Playlist }>('/playlist/', {
      method: 'POST',
      body: JSON.stringify(body),
    }),

  getUserPlaylists: () =>
    request<{ playlist: Playlist[] }>('/playlist/user-playlists'),

  getUserPlaylist: (playlistId: string) =>
    request<{ playlist: PlaylistWithVideos }>(`/playlist/user/${playlistId}`),

  getCoursePlaylist: (playlistId: string) =>
    request<{ playlist: CoursePlaylist }>(`/playlist/course/${playlistId}`),

  update: (
    playlistId: string,
    body: { title?: string; description?: string },
  ) =>
    request<Playlist>(`/playlist/${playlistId}`, {
      method: 'PATCH',
      body: JSON.stringify(body),
    }),

  addVideo: (videoId: string, playlistId: string) =>
    request(`/playlist/add/${videoId}/${playlistId}`, { method: 'PATCH' }),

  removeVideo: (videoId: string, playlistId: string) =>
    request(`/playlist/remove/${videoId}/${playlistId}`, { method: 'PATCH' }),
}

export const videoApi = {
  getAll: () => request<{ videos: Video[] }>('/video/all'),

  getMyVideos: () => request<{videos: Video[]}>('/video/'),

  getById: (videoId: string) =>
    request<{ video: Video }>(`/video/${videoId}`),

  getUploadSignature: () =>
    request<{
      timestamp: number
      signature: string
      api_key: string
      cloud_name: string
      folder: string
    }>('/video/upload/signature'),

  upload: (body: {
    title: string
    description?: string
    videofile_url: string
    videofile_public_id: string
    duration_seconds?: number
  }) =>
    request<{ video: Video }>('/video/upload', {
      method: 'POST',
      body: JSON.stringify(body),
    }),

  update: (videoId: string, formData: FormData) =>
    request<{ user: Video }>(`/video/${videoId}`, {
      method: 'PATCH',
      body: formData,
    }),

  delete: (videoId: string) =>
    request<{ video: Video }>(`/video/${videoId}`, {
      method: 'DELETE',
    }),
}

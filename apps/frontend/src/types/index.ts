export interface User {
  user_id: string
  username: string
  name: string
  email: string
  phone_no: string | null
  profile_pic: string | null
  city: string | null
  state: string | null
  pincode: string | null
  created_at: string
  updated_at: string
}

export interface Video {
  video_id: string
  title: string
  description: string | null
  thumbnail_url: string | null
  videofile_url: string
  duration_seconds: number | string
  views: number
  uploaded_by: string
  created_at: string
  updated_at: string
}

export interface ApiError {
  error?: string
  message?: string
}


export interface ApiResponse<T = undefined> {
  message: string
  data?: T // Optional (?) because omitempty hides it when nil on the backend
}
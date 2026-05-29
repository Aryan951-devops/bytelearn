export interface User {
  user_id: string
  username: string
  name: string
  email: string
  phone_no: string | null
  profile_pic_url: string | null
  city: string | null
  state: string | null
  pincode: string | null
  role: string
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

export interface VideoMetadata {
  video_id: string
  title: string
  thumbnail_url: string | null
  views: number
}

export interface Course {
  course_id: string
  title: string
  description: string | null
  category: string | null
  created_at: string
  updated_at: string
}

export interface EducatorDetail {
  user_id: string
  username: string
  name: string
  profile_pic_url: string
}

export interface PlaylistPreview {
  playlist_id: string
  title: string
  description: string | null
}

export interface CourseWithPlaylists {
  course_id: string
  title: string
  description: string | null
  category: string | null
  created_at: string
  updated_at: string
  playlists: PlaylistPreview[]
  educators: EducatorDetail[]
}

export interface Playlist {
  playlist_id: string
  type: string
  title: string
  description: string | null
  user_id: string
  course_id: string | null
  created_at: string
  updated_at: string
}

export interface PlaylistWithVideos {
  playlist_id: string
  type: string
  title: string
  description: string
  user_id: string
  course_id: string | null
  videos: VideoMetadata[]
  created_at: string
  updated_at: string
}

export interface CoursePlaylist extends PlaylistWithVideos {
  educator_user_id: string
  educator_username: string
  educator_name: string
  educator_profile_pic: string | null
}

export interface ApiError {
  error?: string
  message?: string
}

export interface Comment {
  comment_id: string
  video_id: string
  user_id: string
  user_name: string
  content: string
  created_at: string
  likes_count?: number
}

export interface ApiResponse<T = undefined> {
  message: string
  data?: T
}

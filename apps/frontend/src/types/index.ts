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
  content: string
  created_at: string
  updated_at: string
  commented_by?: string
}

export interface CodingPracticeSummary {
  contest_id: string
  title: string
  description: string | null
  playlist_id: string
  created_at: string
  updated_at: string
}

export interface CodingPractice extends CodingPracticeSummary {
  questions: CodingQuestionMeta[]
}

export interface CodingQuestionMeta {
  question_id: string
  title: string
  difficulty: string
}

export interface CodingQuestion extends CodingQuestionMeta {
  contest_id: string
  statement: string
  constraints: string | null
  input_format: string | null
  output_format: string | null
  time_limit_ms: number
  memory_limit_mb: number
  created_at?: string
  updated_at?: string
}

export interface TestCase {
  testcase_id: string
  question_id: string
  input: string
  expected_output: string
  is_hidden: boolean
  created_at: string
}

export interface Submission {
  submission_id: string
  question_id: string
  user_id: string
  code: string
  language: string
  status: string
  passed_cases: number
  total_cases: number
  started_at: string | null
  finished_at: string | null
  submitted_at?: string
  updated_at?: string
}

export interface SubmissionStatus {
  submission_id: string
  status: string
  passed_cases: number
  total_cases: number
  started_at: string | null
  finished_at: string | null
}

export interface SubmissionResult {
  submission_id: string
  testcase_id: string
  input: string
  expected_output: string
  actual_output: string | null
  error_output: string | null
  is_passed: boolean
  verdict: string
  runtime_ms: number | null
  memory_kb: number | null
}

export interface ApiResponse<T = undefined> {
  message: string
  data?: T
}

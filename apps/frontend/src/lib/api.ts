import type {
  ApiError,
  ApiResponse,
  CodingPractice,
  CodingPracticeSummary,
  CodingQuestion,
  Comment,
  Course,
  CoursePlaylist,
  CourseWithPlaylists,
  Playlist,
  PlaylistWithVideos,
  Submission,
  SubmissionResult,
  SubmissionStatus,
  TestCase,
  User,
  Video,
  VideoMetadata,
  WatchHistoryResponse,
  QuizSummary,
  CreateQuizQuestionInput,
  StartQuizResponse,
  UserSubmittedAnswer,
  SubmitQuizResponse,
  QuizAttemptResult,
  QuizAttemptSummary,
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

  getWatchHistory: () => request<WatchHistoryResponse>('/user/watch-history'),
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

  search: (query: string) =>
    request<{ videos: VideoMetadata[] }>(
      `/video/search/${encodeURIComponent(query)}`,
    ),
}

export const commentApi = {
  getForVideo: (videoId: string) =>
    request<{ comments: Comment[] }>(`/comment/v/${videoId}`),

  create: (videoId: string, body: { content: string }) =>
    request<{ comment: Comment }>(`/comment/v/${videoId}`, {
      method: 'POST',
      body: JSON.stringify(body),
    }),

  update: (commentId: string, body: { content: string }) =>
    request<{ comment: Comment }>(`/comment/c/${commentId}`, {
      method: 'PATCH',
      body: JSON.stringify(body),
    }),

  delete: (commentId: string) =>
    request(`/comment/c/${commentId}`, {
      method: 'DELETE',
    }),
}

export const likeApi = {
  toggleVideo: (videoId: string) =>
    request<{ liked: boolean }>(`/like/toggle/v/${videoId}`, { method: 'POST' }),

  toggleComment: (commentId: string) =>
    request<{ liked: boolean }>(`/like/toggle/c/${commentId}`, { method: 'POST' }),

  checkVideo: (videoId: string) =>
    request<{ liked: boolean }>(`/like/check/v/${videoId}`),

  checkComment: (commentId: string) =>
    request<{ liked: boolean }>(`/like/check/c/${commentId}`),

  getTotalVideoLikes: (videoId: string) =>
    request<{ total_likes: number }>(`/like/total/v/${videoId}`),

  getTotalCommentLikes: (commentId: string) =>
    request<{ total_likes: number }>(`/like/total/c/${commentId}`),
}

export const codeApi = {
  createPractice: (body: {
    title: string
    description?: string
    playlist_id: string
  }) =>
    request<{ practice: CodingPractice }>('/code/cp/', {
      method: 'POST',
      body: JSON.stringify(body),
    }),

  getPractice: (contestId: string) =>
    request<{ practice: CodingPractice }>(`/code/cp/${contestId}`),

  getPracticesByPlaylist: (playlistId: string) =>
    request<{ practices: CodingPracticeSummary[] }>(
      `/code/cp/playlist/${playlistId}`,
    ),

  getQuestion: (questionId: string) =>
    request<{ question: CodingQuestion }>(`/code/cp/question/${questionId}`),

  createQuestion: (body: {
    contest_id: string
    title: string
    difficulty: string
    statement: string
    constraints?: string
    input_format?: string
    output_format?: string
    time_limit_ms: number
    memory_limit_mb: number
  }) =>
    request<{ question: CodingQuestion }>('/code/cp/question', {
      method: 'POST',
      body: JSON.stringify(body),
    }),

  createTestCase: (body: {
    question_id: string
    input: string
    expected_output: string
    is_hidden: boolean
  }) =>
    request<{ testcase: TestCase }>('/code/cp/testcase', {
      method: 'POST',
      body: JSON.stringify(body),
    }),

  getSampleTestCases: (questionId: string) =>
    request<{ testcases: TestCase[] }>(`/code/cp/testcases-sample/${questionId}`),

  submitSample: (body: {
    question_id: string
    code: string
    language: string
  }) =>
    request<{ submission: Submission }>('/code/cp/submit-sample', {
      method: 'POST',
      body: JSON.stringify(body),
    }),

  submitFinal: (body: {
    question_id: string
    code: string
    language: string
  }) =>
    request<{ submission: Submission }>('/code/cp/submit', {
      method: 'POST',
      body: JSON.stringify(body),
    }),

  pollSubmission: (submissionId: string) =>
    request<{ submission: SubmissionStatus }>(`/code/cp/poll/${submissionId}`),

  getSubmissionResult: (submissionId: string) =>
    request<{ results: SubmissionResult[] }>(
      `/code/cp/submission-result/${submissionId}`,
    ),
}

export const quizApi = {
  createQuiz: (body: {
    title: string
    playlist_id: string
    duration_minutes: number
    questions: CreateQuizQuestionInput[]
  }) =>
    request<{ quiz: QuizSummary }>('/quiz/', {
      method: 'POST',
      body: JSON.stringify(body),
    }),

  getQuizzesByPlaylist: (playlistId: string) =>
    request<QuizSummary[]>(`/quiz/quizzes/${playlistId}`),

  startQuiz: (quizId: string) =>
    request<StartQuizResponse>(`/quiz/start/${quizId}`),

  submitQuiz: (attemptId: string, body: { answers: UserSubmittedAnswer[] }) =>
    request<SubmitQuizResponse>(`/quiz/submit/${attemptId}`, {
      method: 'POST',
      body: JSON.stringify(body),
    }),

  getAttemptResult: (attemptId: string) =>
    request<QuizAttemptResult>(`/quiz/result/${attemptId}`),

  getAttempts: (quizId: string) =>
    request<QuizAttemptSummary[]>(`/quiz/attempts/${quizId}`),
}
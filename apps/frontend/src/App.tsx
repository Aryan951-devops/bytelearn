import { BrowserRouter, Navigate, Route, Routes } from 'react-router-dom'
import { AuthProvider } from '@/context/AuthContext'
import { ThemeProvider } from '@/context/ThemeContext'
import { Layout } from '@/components/layout/Layout'
import { ProtectedRoute } from '@/components/ProtectedRoute'
import { AdminRoute } from '@/components/AdminRoute'
import { HomePage } from '@/pages/HomePage'
import { LoginPage } from '@/pages/LoginPage'
import { RegisterPage } from '@/pages/RegisterPage'
import { VideosPage } from '@/pages/VideosPage'
import { ProfilePage } from '@/pages/ProfilePage'
import { SettingsPage } from '@/pages/SettingsPage'
import { WatchHistoryPage } from '@/pages/WatchHistoryPage'
import { EducatorDashboard } from '@/pages/EducatorDashboard'
import { CourseDetailPage } from '@/pages/CourseDetailPage'
import { CoursePlaylistPage } from '@/pages/CoursePlaylistPage'
import { VideoWatchPage } from '@/pages/VideoWatchPage'
import { MyPlaylistsPage } from '@/pages/MyPlaylistsPage'
import { UserPlaylistPage } from '@/pages/UserPlaylistPage'
import { AdminCoursesPage } from '@/pages/AdminCoursesPage'
import { CodingPracticePage } from '@/pages/CodingPracticePage'
import { CodingQuestionManagePage } from '@/pages/CodingQuestionManagePage'
import { CodingSolvePage } from '@/pages/CodingSolvePage'
import { PlaylistPracticesPage } from '@/pages/PlaylistPracticesPage'
import { PlaylistQuizzesPage } from '@/pages/PlaylistQuizzesPage'
import { QuizTakePage } from '@/pages/QuizTakePage'
import { QuizResultPage } from '@/pages/QuizResultPage'
import { QuizAttemptsPage } from '@/pages/QuizAttemptsPage'

export default function App() {
  return (
    <ThemeProvider>
      <AuthProvider>
        <BrowserRouter>
          <Routes>
            <Route element={<Layout />}>
              <Route index element={<HomePage />} />
              <Route path="courses/:courseId" element={<CourseDetailPage />} />
              <Route
                path="courses/:courseId/playlists/:playlistId"
                element={<CoursePlaylistPage />}
              />
              <Route
                path="courses/:courseId/playlists/:playlistId/practices"
                element={<PlaylistPracticesPage />}
              />
              <Route
                path="courses/:courseId/playlists/:playlistId/quizzes"
                element={<PlaylistQuizzesPage />}
              />
              <Route
                path="courses/:courseId/playlists/:playlistId/quizzes/:quizId/attempts"
                element={
                  <ProtectedRoute>
                    <QuizAttemptsPage />
                  </ProtectedRoute>
                }
              />
              <Route
                path="courses/:courseId/playlists/:playlistId/quizzes/:quizId/attempt/:attemptId"
                element={
                  <ProtectedRoute>
                    <QuizTakePage />
                  </ProtectedRoute>
                }
              />
              <Route
                path="courses/:courseId/playlists/:playlistId/quizzes/:quizId/result/:attemptId"
                element={
                  <ProtectedRoute>
                    <QuizResultPage />
                  </ProtectedRoute>
                }
              />
              <Route
                path="courses/:courseId/playlists/:playlistId/coding/:contestId"
                element={<CodingPracticePage />}
              />
              <Route
                path="courses/:courseId/playlists/:playlistId/coding/:contestId/questions/:questionId"
                element={<CodingSolvePage />}
              />
              <Route
                path="courses/:courseId/playlists/:playlistId/coding/:contestId/questions/:questionId/manage"
                element={
                  <ProtectedRoute>
                    <CodingQuestionManagePage />
                  </ProtectedRoute>
                }
              />
              <Route path="watch/:videoId" element={<VideoWatchPage />} />
              <Route path="login" element={<LoginPage />} />
              <Route path="register" element={<RegisterPage />} />
              <Route path="videos" element={<VideosPage />} />

              <Route
                path="my-playlists"
                element={
                  <ProtectedRoute>
                    <MyPlaylistsPage />
                  </ProtectedRoute>
                }
              />
              <Route
                path="watch-history"
                element={
                  <ProtectedRoute>
                    <WatchHistoryPage />
                  </ProtectedRoute>
                }
              />
              <Route
                path="my-playlists/:playlistId"
                element={
                  <ProtectedRoute>
                    <UserPlaylistPage />
                  </ProtectedRoute>
                }
              />

              <Route
                path="admin/courses"
                element={
                  <AdminRoute>
                    <AdminCoursesPage />
                  </AdminRoute>
                }
              />

              <Route
                path="educator/dashboard"
                element={
                  <ProtectedRoute>
                    <EducatorDashboard />
                  </ProtectedRoute>
                }
              />
              <Route
                path="profile"
                element={
                  <ProtectedRoute>
                    <ProfilePage />
                  </ProtectedRoute>
                }
              />
              <Route
                path="settings"
                element={
                  <ProtectedRoute>
                    <SettingsPage />
                  </ProtectedRoute>
                }
              />
              <Route path="*" element={<Navigate to="/" replace />} />
            </Route>
          </Routes>
        </BrowserRouter>
      </AuthProvider>
    </ThemeProvider>
  )
}

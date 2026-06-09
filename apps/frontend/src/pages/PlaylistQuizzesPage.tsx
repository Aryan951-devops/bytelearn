import { useCallback, useEffect, useState } from 'react'
import { Link, useNavigate, useParams } from 'react-router-dom'
import { ClipboardList, Clock, History, Play } from 'lucide-react'
import { courseApi, playlistApi, quizApi } from '@/lib/api'
import type { QuizSummary } from '@/types'
import { CreateQuizCard } from '@/components/quiz/CreateQuizCard'
import { Breadcrumbs } from '@/components/ui/Breadcrumbs'
import { Button } from '@/components/ui/Button'
import { Card } from '@/components/ui/Card'
import { LoadingSpinner } from '@/components/ui/LoadingSpinner'
import { useAuth } from '@/context/AuthContext'

export function PlaylistQuizzesPage() {
  const { courseId, playlistId } = useParams<{
    courseId: string
    playlistId: string
  }>()
  const navigate = useNavigate()
  const { user, isAuthenticated } = useAuth()
  const [quizzes, setQuizzes] = useState<QuizSummary[]>([])
  const [playlistTitle, setPlaylistTitle] = useState('Playlist')
  const [courseTitle, setCourseTitle] = useState('Course')
  const [educatorUserId, setEducatorUserId] = useState('')
  const [loading, setLoading] = useState(true)
  const [startingId, setStartingId] = useState<string | null>(null)
  const [error, setError] = useState('')

  const load = useCallback(async () => {
    if (!playlistId) return
    setLoading(true)
    setError('')
    try {
      const [quizzesRes, playlistRes, courseRes] = await Promise.all([
        quizApi.getQuizzesByPlaylist(playlistId),
        playlistApi.getCoursePlaylist(playlistId),
        courseId ? courseApi.getById(courseId).catch(() => null) : Promise.resolve(null),
      ])
      const list = Array.isArray(quizzesRes.data) ? quizzesRes.data : []
      setQuizzes(list)
      const playlist = playlistRes.data?.playlist
      if (playlist) {
        setPlaylistTitle(playlist.title)
        setEducatorUserId(playlist.educator_user_id || playlist.user_id)
      }
      if (courseRes?.data?.course?.title) {
        setCourseTitle(courseRes.data.course.title)
      }
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Failed to load quizzes')
    } finally {
      setLoading(false)
    }
  }, [playlistId, courseId])

  useEffect(() => {
    load()
  }, [load])

  const isOwner = user?.role === 'educator' && user.user_id === educatorUserId
  const playlistPath = `/courses/${courseId}/playlists/${playlistId}`
  const basePath = `/courses/${courseId}/playlists/${playlistId}/quizzes`

  const openAttempts = (quizId: string) => {
    if (!isAuthenticated) {
      navigate('/login', { state: { from: `${basePath}/${quizId}/attempts` } })
      return
    }
    navigate(`${basePath}/${quizId}/attempts`)
  }

  const startQuiz = async (quizId: string) => {
    if (!isAuthenticated) {
      navigate('/login', { state: { from: `${basePath}` } })
      return
    }
    setStartingId(quizId)
    setError('')
    try {
      const res = await quizApi.startQuiz(quizId)
      const session = res.data
      if (!session?.attempt_id) throw new Error('Could not start quiz')
      navigate(`${basePath}/${quizId}/attempt/${session.attempt_id}`, {
        state: { session },
      })
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Failed to start quiz')
    } finally {
      setStartingId(null)
    }
  }

  if (loading) return <LoadingSpinner className="min-h-[40vh]" />

  return (
    <div className="animate-fade-in space-y-8">
      <Breadcrumbs
        items={[
          { label: 'Home', to: '/' },
          { label: courseTitle, to: `/courses/${courseId}` },
          { label: playlistTitle, to: playlistPath },
          { label: 'Quizzes' },
        ]}
      />

      <header>
        <h1 className="flex items-center gap-2 font-display text-3xl font-semibold text-ink-900 dark:text-ink-50">
          <ClipboardList className="size-8 text-mist-500" />
          Quizzes
        </h1>
        <p className="mt-2 text-ink-500 dark:text-ink-400">
          Test your knowledge with timed quizzes for this playlist.
        </p>
      </header>

      {error ? <p className="text-sm text-red-500">{error}</p> : null}

      {quizzes.length === 0 ? (
        isOwner ? (
          <CreateQuizCard
            playlistId={playlistId!}
            onCreated={() => load()}
          />
        ) : (
          <Card className="py-12 text-center text-ink-500">
            No quizzes available yet.
          </Card>
        )
      ) : (
        <div className="grid gap-4 sm:grid-cols-2">
          {quizzes.map((quiz) => (
              <Card key={quiz.quiz_id} hover className="flex flex-col gap-3">
                <div>
                  <h2 className="font-display text-lg font-semibold text-ink-900 dark:text-ink-50">
                    {quiz.title}
                  </h2>
                  <p className="mt-1 flex items-center gap-1.5 text-sm text-ink-500">
                    <Clock className="size-3.5" />
                    {quiz.duration_minutes} minutes
                  </p>
                </div>
                <div className="flex flex-wrap gap-2">
                  <Button
                    size="sm"
                    isLoading={startingId === quiz.quiz_id}
                    onClick={() => startQuiz(quiz.quiz_id)}
                  >
                    <Play className="size-4" />
                    Start quiz
                  </Button>
                  <Button
                    variant="secondary"
                    size="sm"
                    onClick={() => openAttempts(quiz.quiz_id)}
                  >
                    <History className="size-4" />
                    See previous attempts
                  </Button>
                </div>
              </Card>
          ))}
        </div>
      )}

      {isOwner && quizzes.length > 0 ? (
        <CreateQuizCard playlistId={playlistId!} onCreated={() => load()} />
      ) : null}

      <Link to={playlistPath}>
        <Button variant="ghost">Back to playlist</Button>
      </Link>
    </div>
  )
}

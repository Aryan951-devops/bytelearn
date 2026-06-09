import { useCallback, useEffect, useState } from 'react'
import { Link, useNavigate, useParams } from 'react-router-dom'
import { ChevronRight, ClipboardList } from 'lucide-react'
import { courseApi, playlistApi, quizApi } from '@/lib/api'
import type { QuizAttemptSummary, QuizSummary } from '@/types'
import { Breadcrumbs } from '@/components/ui/Breadcrumbs'
import { Button } from '@/components/ui/Button'
import { Card } from '@/components/ui/Card'
import { LoadingSpinner } from '@/components/ui/LoadingSpinner'
import { useAuth } from '@/context/AuthContext'

function formatAttemptDate(value: string) {
  if (!value) return '—'
  return new Date(value).toLocaleString(undefined, {
    dateStyle: 'medium',
    timeStyle: 'short',
  })
}

export function QuizAttemptsPage() {
  const { courseId, playlistId, quizId } = useParams<{
    courseId: string
    playlistId: string
    quizId: string
  }>()
  const navigate = useNavigate()
  const { isAuthenticated } = useAuth()

  const [quiz, setQuiz] = useState<QuizSummary | null>(null)
  const [attempts, setAttempts] = useState<QuizAttemptSummary[]>([])
  const [playlistTitle, setPlaylistTitle] = useState('Playlist')
  const [courseTitle, setCourseTitle] = useState('Course')
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState('')

  const basePath = `/courses/${courseId}/playlists/${playlistId}/quizzes`
  const resultBasePath = `${basePath}/${quizId}/result`

  const load = useCallback(async () => {
    if (!playlistId || !quizId) return
    if (!isAuthenticated) {
      navigate('/login', { state: { from: `${basePath}/${quizId}/attempts` } })
      return
    }

    setLoading(true)
    setError('')
    try {
      const [quizzesRes, attemptsRes, playlistRes, courseRes] = await Promise.all([
        quizApi.getQuizzesByPlaylist(playlistId),
        quizApi.getAttempts(quizId),
        playlistApi.getCoursePlaylist(playlistId),
        courseId ? courseApi.getById(courseId).catch(() => null) : Promise.resolve(null),
      ])

      const quizzes = Array.isArray(quizzesRes.data) ? quizzesRes.data : []
      setQuiz(quizzes.find((q) => q.quiz_id === quizId) ?? null)

      const list = Array.isArray(attemptsRes.data) ? attemptsRes.data : []
      setAttempts(
        [...list].sort(
          (a, b) =>
            new Date(b.submitted_at).getTime() - new Date(a.submitted_at).getTime(),
        ),
      )

      const playlist = playlistRes.data?.playlist
      if (playlist) setPlaylistTitle(playlist.title)
      if (courseRes?.data?.course?.title) setCourseTitle(courseRes.data.course.title)
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Failed to load attempts')
    } finally {
      setLoading(false)
    }
  }, [playlistId, quizId, courseId, isAuthenticated, navigate, basePath])

  useEffect(() => {
    load()
  }, [load])

  if (loading) return <LoadingSpinner className="min-h-[40vh]" />

  return (
    <div className="animate-fade-in space-y-8">
      <Breadcrumbs
        items={[
          { label: 'Home', to: '/' },
          { label: courseTitle, to: `/courses/${courseId}` },
          { label: playlistTitle, to: `/courses/${courseId}/playlists/${playlistId}` },
          { label: 'Quizzes', to: basePath },
          { label: quiz?.title ?? 'Previous attempts' },
        ]}
      />

      <header>
        <h1 className="flex items-center gap-2 font-display text-3xl font-semibold text-ink-900 dark:text-ink-50">
          <ClipboardList className="size-8 text-mist-500" />
          Previous attempts
        </h1>
        {quiz ? (
          <p className="mt-2 text-ink-500 dark:text-ink-400">{quiz.title}</p>
        ) : null}
      </header>

      {error ? <p className="text-sm text-red-500">{error}</p> : null}

      {attempts.length === 0 ? (
        <Card className="py-12 text-center text-ink-500">
          No submitted attempts yet. Start the quiz to record your first attempt.
        </Card>
      ) : (
        <div className="space-y-3">
          {attempts.map((attempt, index) => (
            <Link
              key={attempt.attempt_id}
              to={`${resultBasePath}/${attempt.attempt_id}`}
              className="block"
            >
              <Card
                hover
                className="flex items-center justify-between gap-4 transition-colors"
              >
                <div className="min-w-0 space-y-1">
                  <p className="font-medium text-ink-900 dark:text-ink-50">
                    Attempt {attempts.length - index}
                  </p>
                  <p className="text-sm text-ink-500">
                    Submitted {formatAttemptDate(attempt.submitted_at)}
                  </p>
                  {attempt.started_at ? (
                    <p className="text-xs text-ink-400">
                      Started {formatAttemptDate(attempt.started_at)}
                    </p>
                  ) : null}
                </div>
                <div className="flex shrink-0 items-center gap-3">
                  <div className="text-right">
                    <p className="text-lg font-semibold text-sage-600 dark:text-sage-400">
                      {attempt.score} / {attempt.total_marks}
                    </p>
                    {attempt.status ? (
                      <p className="text-xs uppercase tracking-wide text-ink-400">
                        {attempt.status}
                      </p>
                    ) : null}
                  </div>
                  <ChevronRight className="size-5 text-ink-400" />
                </div>
              </Card>
            </Link>
          ))}
        </div>
      )}

      <div className="flex flex-wrap gap-3">
        <Link to={basePath}>
          <Button variant="ghost">Back to quizzes</Button>
        </Link>
      </div>
    </div>
  )
}

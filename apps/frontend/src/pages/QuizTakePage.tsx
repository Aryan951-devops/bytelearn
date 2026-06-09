import { useCallback, useEffect, useMemo, useRef, useState } from 'react'
import { Link, useLocation, useNavigate, useParams } from 'react-router-dom'
import { Clock } from 'lucide-react'
import { quizApi } from '@/lib/api'
import type { QuizQuestionMetadata, StartQuizResponse, UserSubmittedAnswer } from '@/types'
import { Breadcrumbs } from '@/components/ui/Breadcrumbs'
import { Button } from '@/components/ui/Button'
import { Card } from '@/components/ui/Card'
import { LoadingSpinner } from '@/components/ui/LoadingSpinner'

type AnswerState = Record<
  string,
  { selected_options?: number[]; text_answer?: string }
>

function formatTimeLeft(ms: number) {
  const totalSec = Math.max(0, Math.floor(ms / 1000))
  const min = Math.floor(totalSec / 60)
  const sec = totalSec % 60
  return `${min}:${sec.toString().padStart(2, '0')}`
}

export function QuizTakePage() {
  const { courseId, playlistId, quizId, attemptId } = useParams<{
    courseId: string
    playlistId: string
    quizId: string
    attemptId: string
  }>()
  const location = useLocation()
  const navigate = useNavigate()

  const session = location.state?.session as StartQuizResponse | undefined

  const [questions, setQuestions] = useState<QuizQuestionMetadata[]>(session?.questions ?? [])
  const [startedAt, setStartedAt] = useState(session?.started_at ?? '')
  const [durationMinutes, setDurationMinutes] = useState(session?.duration_minutes ?? 0)
  const [answers, setAnswers] = useState<AnswerState>({})
  const [now, setNow] = useState(Date.now())
  const [submitting, setSubmitting] = useState(false)
  const [error, setError] = useState('')
  const [loading, setLoading] = useState(!session)
  const autoSubmitted = useRef(false)

  const basePath = `/courses/${courseId}/playlists/${playlistId}/quizzes`
  const quizzesPath = basePath

  const endTime = useMemo(() => {
    if (!startedAt || !durationMinutes) return 0
    return new Date(startedAt).getTime() + durationMinutes * 60 * 1000
  }, [startedAt, durationMinutes])

  const timeLeftMs = endTime - now
  const expired = endTime > 0 && timeLeftMs <= 0

  useEffect(() => {
    if (session) return
    if (!quizId) return
    setLoading(true)
    quizApi
      .startQuiz(quizId)
      .then((res) => {
        const data = res.data
        if (!data) throw new Error('Could not load quiz session')
        setQuestions(data.questions)
        setStartedAt(data.started_at)
        setDurationMinutes(data.duration_minutes)
        if (data.attempt_id !== attemptId) {
          navigate(
            `/courses/${courseId}/playlists/${playlistId}/quizzes/${quizId}/attempt/${data.attempt_id}`,
            { replace: true, state: { session: data } },
          )
        }
      })
      .catch((err) => {
        setError(err instanceof Error ? err.message : 'Failed to start quiz')
      })
      .finally(() => setLoading(false))
  }, [session, quizId, attemptId, courseId, playlistId, navigate])

  useEffect(() => {
    const id = window.setInterval(() => setNow(Date.now()), 1000)
    return () => window.clearInterval(id)
  }, [])

  const buildAnswers = useCallback((): UserSubmittedAnswer[] => {
    return questions.map((q) => {
      const a = answers[q.question_id]
      if (q.type === 'one_word') {
        return { question_id: q.question_id, text_answer: a?.text_answer ?? '' }
      }
      return { question_id: q.question_id, selected_options: a?.selected_options ?? [] }
    })
  }, [questions, answers])

  const submit = useCallback(async () => {
    if (!attemptId || submitting) return
    setSubmitting(true)
    setError('')
    try {
      const res = await quizApi.submitQuiz(attemptId, { answers: buildAnswers() })
      const result = res.data
      navigate(`${basePath}/${quizId}/result/${attemptId}`, {
        state: { submitResult: result },
      })
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Failed to submit quiz')
    } finally {
      setSubmitting(false)
    }
  }, [attemptId, submitting, buildAnswers, navigate, basePath, quizId])

  useEffect(() => {
    if (expired && !submitting && !autoSubmitted.current) {
      autoSubmitted.current = true
      submit()
    }
  }, [expired, submitting, submit])

  const setSingleOption = (questionId: string, optIndex: number) => {
    setAnswers((prev) => ({
      ...prev,
      [questionId]: { selected_options: [optIndex] },
    }))
  }

  const toggleMultipleOption = (questionId: string, optIndex: number) => {
    setAnswers((prev) => {
      const current = new Set(prev[questionId]?.selected_options ?? [])
      if (current.has(optIndex)) current.delete(optIndex)
      else current.add(optIndex)
      return {
        ...prev,
        [questionId]: { selected_options: [...current].sort((a, b) => a - b) },
      }
    })
  }

  const setTextAnswer = (questionId: string, text: string) => {
    setAnswers((prev) => ({
      ...prev,
      [questionId]: { text_answer: text },
    }))
  }

  if (loading) return <LoadingSpinner className="min-h-[40vh]" />

  if (!questions.length) {
    return (
      <Card className="space-y-4 py-8 text-center">
        <p className="text-ink-500">{error || 'Quiz session not found.'}</p>
        <Link to={quizzesPath}>
          <Button variant="ghost">Back to quizzes</Button>
        </Link>
      </Card>
    )
  }

  return (
    <div className="animate-fade-in space-y-6">
      <Breadcrumbs
        items={[
          { label: 'Quizzes', to: quizzesPath },
          { label: 'Attempt' },
        ]}
      />

      <div className="flex flex-wrap items-center justify-between gap-3">
        <h1 className="font-display text-2xl font-semibold text-ink-900 dark:text-ink-50">
          Quiz in progress
        </h1>
        <div
          className={`flex items-center gap-2 rounded-full px-4 py-1.5 text-sm font-medium ${
            timeLeftMs < 60_000
              ? 'bg-red-100 text-red-700 dark:bg-red-900/40 dark:text-red-300'
              : 'bg-mist-100 text-mist-700 dark:bg-mist-900/40 dark:text-mist-300'
          }`}
        >
          <Clock className="size-4" />
          {expired ? 'Time up' : formatTimeLeft(timeLeftMs)}
        </div>
      </div>

      {error ? <p className="text-sm text-red-500">{error}</p> : null}

      <div className="space-y-4">
        {questions.map((q, index) => (
          <Card key={q.question_id} className="space-y-3">
            <div className="flex flex-wrap items-start justify-between gap-2">
              <p className="font-medium text-ink-900 dark:text-ink-50">
                {index + 1}. {q.question}
              </p>
              <span className="text-xs text-ink-400">
                {q.marks} mark{q.marks !== 1 ? 's' : ''}
                {q.negative_marks > 0 ? ` · −${q.negative_marks} wrong` : ''}
              </span>
            </div>

            {q.type === 'one_word' ? (
              <input
                value={answers[q.question_id]?.text_answer ?? ''}
                onChange={(e) => setTextAnswer(q.question_id, e.target.value)}
                placeholder="Your answer"
                className="w-full rounded-xl border border-ink-200/80 bg-white/80 px-4 py-2.5 text-sm dark:border-ink-700 dark:bg-ink-900/50"
              />
            ) : (
              <div className="space-y-2">
                {(q.options ?? []).map((opt, optIndex) => {
                  const selected = answers[q.question_id]?.selected_options ?? []
                  const isChecked = selected.includes(optIndex)
                  return (
                    <label
                      key={optIndex}
                      className={`flex cursor-pointer items-center gap-3 rounded-xl border px-4 py-3 text-sm transition-colors ${
                        isChecked
                          ? 'border-sage-400 bg-sage-50 dark:border-sage-600 dark:bg-sage-900/30'
                          : 'border-ink-200/60 hover:bg-cream-50 dark:border-ink-700 dark:hover:bg-ink-800/50'
                      }`}
                    >
                      <input
                        type={q.type === 'multiple' ? 'checkbox' : 'radio'}
                        name={`question-${q.question_id}`}
                        checked={isChecked}
                        onChange={() =>
                          q.type === 'multiple'
                            ? toggleMultipleOption(q.question_id, optIndex)
                            : setSingleOption(q.question_id, optIndex)
                        }
                        className="size-4 accent-sage-600"
                      />
                      <span>{opt}</span>
                    </label>
                  )
                })}
              </div>
            )}
          </Card>
        ))}
      </div>

      <div className="flex flex-wrap gap-3">
        <Button onClick={submit} isLoading={submitting} disabled={expired && submitting}>
          Submit quiz
        </Button>
        <Link to={quizzesPath}>
          <Button variant="ghost">Cancel</Button>
        </Link>
      </div>
    </div>
  )
}

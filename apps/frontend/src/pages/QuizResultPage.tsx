import { useCallback, useEffect, useState } from 'react'
import { Link, useLocation, useParams } from 'react-router-dom'
import { CheckCircle2, XCircle } from 'lucide-react'
import { quizApi } from '@/lib/api'
import type { QuizAnswerResult, QuizAttemptResult, SubmitQuizResponse } from '@/types'
import { Breadcrumbs } from '@/components/ui/Breadcrumbs'
import { Button } from '@/components/ui/Button'
import { Card } from '@/components/ui/Card'
import { LoadingSpinner } from '@/components/ui/LoadingSpinner'

function isAnswerCorrect(answer: QuizAnswerResult): boolean {
  if (answer.type === 'one_word') {
    return (
      answer.text_answer?.trim().toLowerCase() ===
      answer.correct_answer?.trim().toLowerCase()
    )
  }
  const selected = [...(answer.selected_options ?? [])].sort((a, b) => a - b)
  const correct = [...(answer.correct_options ?? [])].sort((a, b) => a - b)
  if (selected.length !== correct.length) return false
  return selected.every((v, i) => v === correct[i])
}

function formatOptions(options: string[] | undefined, indices: number[] | undefined) {
  if (!options?.length || !indices?.length) return '—'
  return indices.map((i) => options[i] ?? `Option ${i + 1}`).join(', ')
}

export function QuizResultPage() {
  const { courseId, playlistId, quizId, attemptId } = useParams<{
    courseId: string
    playlistId: string
    quizId: string
    attemptId: string
  }>()
  const location = useLocation()
  const submitResult = location.state?.submitResult as SubmitQuizResponse | undefined

  const [result, setResult] = useState<QuizAttemptResult | null>(null)
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState('')

  const basePath = `/courses/${courseId}/playlists/${playlistId}/quizzes`

  const load = useCallback(async () => {
    if (!attemptId) return
    setLoading(true)
    setError('')
    try {
      const res = await quizApi.getAttemptResult(attemptId)
      setResult(res.data ?? null)
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Failed to load result')
    } finally {
      setLoading(false)
    }
  }, [attemptId])

  useEffect(() => {
    load()
  }, [load])

  if (loading) return <LoadingSpinner className="min-h-[40vh]" />

  const score = result?.score ?? submitResult?.score ?? 0
  const total = result?.total_marks ?? submitResult?.total_marks ?? 0
  const answers = result?.submitted_answers ?? []

  return (
    <div className="animate-fade-in space-y-8">
      <Breadcrumbs
        items={[
          { label: 'Quizzes', to: basePath },
          ...(quizId
            ? [{ label: 'Previous attempts', to: `${basePath}/${quizId}/attempts` }]
            : []),
          { label: 'Result' },
        ]}
      />

      <header className="space-y-2">
        <h1 className="font-display text-3xl font-semibold text-ink-900 dark:text-ink-50">
          Quiz result
        </h1>
        <p className="text-lg text-ink-600 dark:text-ink-300">
          You scored{' '}
          <span className="font-semibold text-sage-600 dark:text-sage-400">
            {score} / {total}
          </span>
        </p>
      </header>

      {error ? <p className="text-sm text-red-500">{error}</p> : null}

      {answers.length > 0 ? (
        <div className="space-y-4">
          {answers.map((answer, index) => {
            const correct = isAnswerCorrect(answer)
            return (
              <Card key={answer.question_id} className="space-y-3">
                <div className="flex items-start gap-3">
                  {correct ? (
                    <CheckCircle2 className="mt-0.5 size-5 shrink-0 text-sage-500" />
                  ) : (
                    <XCircle className="mt-0.5 size-5 shrink-0 text-red-500" />
                  )}
                  <div className="min-w-0 flex-1 space-y-2">
                    <p className="font-medium text-ink-900 dark:text-ink-50">
                      {index + 1}. {answer.question}
                    </p>
                    <div className="grid gap-2 text-sm sm:grid-cols-2">
                      <div>
                        <p className="text-xs font-medium uppercase text-ink-400">Your answer</p>
                        <p className="text-ink-700 dark:text-ink-200">
                          {answer.type === 'one_word'
                            ? answer.text_answer || '—'
                            : formatOptions(answer.options, answer.selected_options)}
                        </p>
                      </div>
                      <div>
                        <p className="text-xs font-medium uppercase text-ink-400">Correct answer</p>
                        <p className="text-ink-700 dark:text-ink-200">
                          {answer.type === 'one_word'
                            ? answer.correct_answer || '—'
                            : formatOptions(answer.options, answer.correct_options)}
                        </p>
                      </div>
                    </div>
                    {answer.explanation ? (
                      <p className="rounded-lg bg-cream-100/80 px-3 py-2 text-sm text-ink-600 dark:bg-ink-800/60 dark:text-ink-300">
                        {answer.explanation}
                      </p>
                    ) : null}
                  </div>
                </div>
              </Card>
            )
          })}
        </div>
      ) : (
        <Card className="py-8 text-center text-ink-500">
          Detailed answers are not available for this attempt.
        </Card>
      )}

      <div className="flex flex-wrap gap-3">
        {quizId ? (
          <Link to={`${basePath}/${quizId}/attempts`}>
            <Button variant="secondary">All attempts</Button>
          </Link>
        ) : null}
        <Link to={basePath}>
          <Button>Back to quizzes</Button>
        </Link>
        <Link to={`/courses/${courseId}/playlists/${playlistId}`}>
          <Button variant="ghost">Back to playlist</Button>
        </Link>
      </div>
    </div>
  )
}

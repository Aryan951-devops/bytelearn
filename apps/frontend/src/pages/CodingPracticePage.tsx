import { type FormEvent, useEffect, useState } from 'react'
import { Link, useNavigate, useParams } from 'react-router-dom'
import { Code2, Plus, Settings } from 'lucide-react'
import { codeApi } from '@/lib/api'
import type { CodingPractice } from '@/types'
import { Breadcrumbs } from '@/components/ui/Breadcrumbs'
import { Button } from '@/components/ui/Button'
import { Card } from '@/components/ui/Card'
import { Input } from '@/components/ui/Input'
import { LoadingSpinner } from '@/components/ui/LoadingSpinner'
import { useAuth } from '@/context/AuthContext'

const DIFFICULTIES = ['easy', 'medium', 'hard']

export function CodingPracticePage() {
  const { courseId, playlistId, contestId } = useParams<{
    courseId: string
    playlistId: string
    contestId: string
  }>()
  const navigate = useNavigate()
  const { user } = useAuth()
  const [practice, setPractice] = useState<CodingPractice | null>(null)
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState('')
  const [showQuestionForm, setShowQuestionForm] = useState(false)
  const [qTitle, setQTitle] = useState('')
  const [qDifficulty, setQDifficulty] = useState('easy')
  const [qStatement, setQStatement] = useState('')
  const [qConstraints, setQConstraints] = useState('')
  const [qInputFormat, setQInputFormat] = useState('')
  const [qOutputFormat, setQOutputFormat] = useState('')
  const [qTimeLimit, setQTimeLimit] = useState('2000')
  const [qMemoryLimit, setQMemoryLimit] = useState('256')
  const [creatingQ, setCreatingQ] = useState(false)

  const isEducator = user?.role === 'educator'
  const basePath = `/courses/${courseId}/playlists/${playlistId}/coding/${contestId}`

  const load = () => {
    if (!contestId) return
    setLoading(true)
    codeApi
      .getPractice(contestId)
      .then((res) => setPractice(res.data?.practice ?? null))
      .catch((err) =>
        setError(err instanceof Error ? err.message : 'Failed to load practice'),
      )
      .finally(() => setLoading(false))
  }

  useEffect(() => {
    load()
  }, [contestId])

  const handleCreateQuestion = async (e: FormEvent) => {
    e.preventDefault()
    if (!contestId) return
    setCreatingQ(true)
    setError('')
    try {
      const res = await codeApi.createQuestion({
        contest_id: contestId,
        title: qTitle.trim(),
        difficulty: qDifficulty,
        statement: qStatement.trim(),
        constraints: qConstraints.trim() || undefined,
        input_format: qInputFormat.trim() || undefined,
        output_format: qOutputFormat.trim() || undefined,
        time_limit_ms: Number(qTimeLimit),
        memory_limit_mb: Number(qMemoryLimit),
      })
      const question = res.data?.question
      if (question) {
        setShowQuestionForm(false)
        navigate(`${basePath}/questions/${question.question_id}/manage`)
      }
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Failed to create question')
    } finally {
      setCreatingQ(false)
    }
  }

  if (loading) return <LoadingSpinner className="min-h-[40vh]" />

  if (error && !practice) {
    return (
      <Card>
        <p className="text-center text-red-500">{error}</p>
      </Card>
    )
  }

  if (!practice) return null

  return (
    <div className="animate-fade-in space-y-8">
      <Breadcrumbs
        items={[
          { label: 'Home', to: '/' },
          {
            label: 'Playlist',
            to: `/courses/${courseId}/playlists/${playlistId}`,
          },
          { label: practice.title },
        ]}
      />

      <header className="flex flex-wrap items-start justify-between gap-4">
        <div>
          <h1 className="flex items-center gap-2 font-display text-3xl font-semibold text-ink-900 dark:text-ink-50">
            <Code2 className="size-8 text-sage-500" />
            {practice.title}
          </h1>
          {practice.description ? (
            <p className="mt-2 max-w-2xl text-ink-600 dark:text-ink-300">
              {practice.description}
            </p>
          ) : null}
        </div>
        {isEducator ? (
          <Button variant="secondary" size="sm" onClick={() => setShowQuestionForm((v) => !v)}>
            <Plus className="size-4" />
            Add question
          </Button>
        ) : null}
      </header>

      {error ? <p className="text-sm text-red-500">{error}</p> : null}

      {showQuestionForm && isEducator ? (
        <Card className="space-y-4">
          <h2 className="font-display text-lg font-semibold">New coding question</h2>
          <form onSubmit={handleCreateQuestion} className="space-y-3">
            <Input label="Title" required value={qTitle} onChange={(e) => setQTitle(e.target.value)} />
            <div className="space-y-1.5">
              <label className="text-sm font-medium">Difficulty</label>
              <select
                value={qDifficulty}
                onChange={(e) => setQDifficulty(e.target.value)}
                className="w-full rounded-xl border border-ink-200/80 bg-white/80 px-3 py-2 text-sm dark:border-ink-700 dark:bg-ink-900/50"
              >
                {DIFFICULTIES.map((d) => (
                  <option key={d} value={d}>
                    {d}
                  </option>
                ))}
              </select>
            </div>
            <div className="space-y-1.5">
              <label className="text-sm font-medium">Problem statement *</label>
              <textarea
                required
                rows={5}
                value={qStatement}
                onChange={(e) => setQStatement(e.target.value)}
                className="w-full rounded-xl border border-ink-200/80 px-3 py-2 text-sm dark:border-ink-700 dark:bg-ink-900/50"
              />
            </div>
            <Input label="Constraints" value={qConstraints} onChange={(e) => setQConstraints(e.target.value)} />
            <Input label="Input format" value={qInputFormat} onChange={(e) => setQInputFormat(e.target.value)} />
            <Input label="Output format" value={qOutputFormat} onChange={(e) => setQOutputFormat(e.target.value)} />
            <div className="grid gap-3 sm:grid-cols-2">
              <Input label="Time limit (ms)" type="number" value={qTimeLimit} onChange={(e) => setQTimeLimit(e.target.value)} />
              <Input label="Memory limit (MB)" type="number" value={qMemoryLimit} onChange={(e) => setQMemoryLimit(e.target.value)} />
            </div>
            <Button type="submit" isLoading={creatingQ}>
              Create question
            </Button>
          </form>
        </Card>
      ) : null}

      <section className="space-y-4">
        <h2 className="font-display text-xl font-semibold">Problems</h2>
        {practice.questions.length === 0 ? (
          <Card className="py-12 text-center text-ink-500">
            {isEducator
              ? 'No questions yet. Add your first problem above.'
              : 'No coding problems available yet.'}
          </Card>
        ) : (
          <div className="space-y-3">
            {practice.questions.map((q) => (
              <Card key={q.question_id} hover className="flex flex-wrap items-center justify-between gap-3">
                <div>
                  <h3 className="font-medium text-ink-900 dark:text-ink-50">{q.title}</h3>
                  <span className="mt-1 inline-block rounded-full bg-mist-100 px-2 py-0.5 text-xs capitalize text-mist-700 dark:bg-mist-900/40 dark:text-mist-300">
                    {q.difficulty}
                  </span>
                </div>
                <div className="flex gap-2">
                  <Link to={`${basePath}/questions/${q.question_id}`}>
                    <Button size="sm">Solve</Button>
                  </Link>
                  {isEducator ? (
                    <Link to={`${basePath}/questions/${q.question_id}/manage`}>
                      <Button variant="ghost" size="sm">
                        <Settings className="size-4" />
                        Test cases
                      </Button>
                    </Link>
                  ) : null}
                </div>
              </Card>
            ))}
          </div>
        )}
      </section>
    </div>
  )
}

export function CreateCodingPracticeCard({
  playlistId,
  onCreated,
}: {
  playlistId: string
  onCreated: (contestId: string) => void
}) {
  const [title, setTitle] = useState('Coding Practice')
  const [description, setDescription] = useState('')
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState('')

  const handleSubmit = async (e: FormEvent) => {
    e.preventDefault()
    setLoading(true)
    setError('')
    try {
      const res = await codeApi.createPractice({
        title: title.trim(),
        description: description.trim() || undefined,
        playlist_id: playlistId,
      })
      const contestId = res.data?.practice?.contest_id
      if (contestId) onCreated(contestId)
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Failed to create coding practice')
    } finally {
      setLoading(false)
    }
  }

  return (
    <Card className="space-y-3">
      <h3 className="flex items-center gap-2 font-display text-lg font-semibold">
        <Code2 className="size-5 text-sage-500" />
        Create coding practice
      </h3>
      <form onSubmit={handleSubmit} className="space-y-3">
        <Input label="Title" required value={title} onChange={(e) => setTitle(e.target.value)} />
        <Input label="Description" value={description} onChange={(e) => setDescription(e.target.value)} />
        {error ? <p className="text-sm text-red-500">{error}</p> : null}
        <Button type="submit" isLoading={loading}>
          Create section
        </Button>
      </form>
    </Card>
  )
}

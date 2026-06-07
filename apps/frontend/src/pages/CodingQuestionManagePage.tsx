import { type FormEvent, useEffect, useState } from 'react'
import { Link, useParams } from 'react-router-dom'
import { FlaskConical } from 'lucide-react'
import { codeApi } from '@/lib/api'
import type { TestCase } from '@/types'
import { Breadcrumbs } from '@/components/ui/Breadcrumbs'
import { Button } from '@/components/ui/Button'
import { Card } from '@/components/ui/Card'
import { LoadingSpinner } from '@/components/ui/LoadingSpinner'
import { useAuth } from '@/context/AuthContext'

export function CodingQuestionManagePage() {
  const { courseId, playlistId, contestId, questionId } = useParams<{
    courseId: string
    playlistId: string
    contestId: string
    questionId: string
  }>()
  const { user } = useAuth()
  const [questionTitle, setQuestionTitle] = useState('')
  const [sampleCases, setSampleCases] = useState<TestCase[]>([])
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState('')
  const [message, setMessage] = useState('')

  const [input, setInput] = useState('')
  const [expectedOutput, setExpectedOutput] = useState('')
  const [isHidden, setIsHidden] = useState(false)
  const [creating, setCreating] = useState(false)

  const basePath = `/courses/${courseId}/playlists/${playlistId}/coding/${contestId}`

  useEffect(() => {
    if (!questionId) return

    Promise.all([
      codeApi.getQuestion(questionId),
      codeApi.getSampleTestCases(questionId),
    ])
      .then(([questionRes, tcRes]) => {
        if (questionRes.data?.question) {
          setQuestionTitle(questionRes.data.question.title)
        }
        setSampleCases(tcRes.data?.testcases ?? [])
      })
      .catch((err) => {
        setError(err instanceof Error ? err.message : 'Failed to load question')
      })
      .finally(() => setLoading(false))
  }, [questionId])

  if (user?.role !== 'educator') {
    return (
      <Card>
        <p className="text-center text-ink-500">Only educators can manage test cases.</p>
      </Card>
    )
  }

  const handleCreate = async (e: FormEvent) => {
    e.preventDefault()
    if (!questionId) return
    setCreating(true)
    setError('')
    try {
      await codeApi.createTestCase({
        question_id: questionId,
        input: input.trim(),
        expected_output: expectedOutput.trim(),
        is_hidden: isHidden,
      })
      setMessage(isHidden ? 'Hidden test case added.' : 'Sample test case added.')
      setInput('')
      setExpectedOutput('')
      setIsHidden(false)
      const res = await codeApi.getSampleTestCases(questionId)
      setSampleCases(res.data?.testcases ?? [])
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Failed to add test case')
    } finally {
      setCreating(false)
    }
  }

  if (loading) return <LoadingSpinner className="min-h-[40vh]" />

  return (
    <div className="animate-fade-in mx-auto max-w-3xl space-y-8">
      <Breadcrumbs
        items={[
          { label: 'Coding practice', to: basePath },
          { label: questionTitle || 'Question' },
          { label: 'Test cases' },
        ]}
      />

      <header>
        <h1 className="flex items-center gap-2 font-display text-2xl font-semibold">
          <FlaskConical className="size-6 text-mist-500" />
          Test cases — {questionTitle}
        </h1>
        <p className="mt-2 text-sm text-ink-500">
          Sample cases (visible to learners) use <code className="text-xs">is_hidden: false</code>.
          Hidden cases are used for final submission only.
        </p>
      </header>

      <Card className="space-y-4">
        <h2 className="font-medium">Add test case</h2>
        <form onSubmit={handleCreate} className="space-y-3">
          <div className="space-y-1.5">
            <label className="text-sm font-medium">Input</label>
            <textarea
              required
              rows={3}
              value={input}
              onChange={(e) => setInput(e.target.value)}
              className="w-full rounded-xl border border-ink-200/80 px-3 py-2 font-mono text-sm dark:border-ink-700 dark:bg-ink-900/50"
            />
          </div>
          <div className="space-y-1.5">
            <label className="text-sm font-medium">Expected output</label>
            <textarea
              required
              rows={3}
              value={expectedOutput}
              onChange={(e) => setExpectedOutput(e.target.value)}
              className="w-full rounded-xl border border-ink-200/80 px-3 py-2 font-mono text-sm dark:border-ink-700 dark:bg-ink-900/50"
            />
          </div>
          <label className="flex items-center gap-2 text-sm">
            <input
              type="checkbox"
              checked={isHidden}
              onChange={(e) => setIsHidden(e.target.checked)}
            />
            Hidden test case (final submission only)
          </label>
          {error ? <p className="text-sm text-red-500">{error}</p> : null}
          {message ? <p className="text-sm text-sage-600">{message}</p> : null}
          <Button type="submit" isLoading={creating}>
            Add test case
          </Button>
        </form>
      </Card>

      <section className="space-y-3">
        <h2 className="font-medium">Sample test cases (visible)</h2>
        {sampleCases.length === 0 ? (
          <Card className="py-8 text-center text-sm text-ink-500">
            No sample test cases yet.
          </Card>
        ) : (
          sampleCases.map((tc) => (
            <Card key={tc.testcase_id} className="space-y-2 font-mono text-xs">
              <p className="font-sans text-sm font-medium text-ink-700 dark:text-ink-200">
                Test case
              </p>
              <div>
                <span className="text-ink-400">Input:</span>
                <pre className="mt-1 whitespace-pre-wrap">{tc.input}</pre>
              </div>
              <div>
                <span className="text-ink-400">Expected:</span>
                <pre className="mt-1 whitespace-pre-wrap">{tc.expected_output}</pre>
              </div>
            </Card>
          ))
        )}
      </section>

      <Link to={basePath}>
        <Button variant="ghost">Back to practice</Button>
      </Link>
    </div>
  )
}

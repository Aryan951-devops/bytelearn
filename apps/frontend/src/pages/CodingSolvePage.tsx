import { useCallback, useEffect, useState } from 'react'
import { Link, useNavigate, useParams } from 'react-router-dom'
import { Play, Send } from 'lucide-react'
import { codeApi } from '@/lib/api'
import { CODING_LANGUAGES, getLanguage } from '@/lib/codingLanguages'
import { pollSubmission } from '@/lib/pollSubmission'
import type { CodingQuestion, SubmissionResult, SubmissionStatus, TestCase } from '@/types'
import { Breadcrumbs } from '@/components/ui/Breadcrumbs'
import { Button } from '@/components/ui/Button'
import { Card } from '@/components/ui/Card'
import { LoadingSpinner } from '@/components/ui/LoadingSpinner'
import { CodeEditor } from '@/components/code/CodeEditor'
import { SubmissionResultsPanel } from '@/components/code/SubmissionResultsPanel'
import { useAuth } from '@/context/AuthContext'

async function mockUnsupportedSubmit(): Promise<{
  status: SubmissionStatus
  results: SubmissionResult[]
}> {
  await new Promise((r) => setTimeout(r, 1200))
  return {
    status: {
      submission_id: 'mock',
      status: 'COMPLETED',
      passed_cases: 0,
      total_cases: 0,
      started_at: null,
      finished_at: null,
    },
    results: [],
  }
}

export function CodingSolvePage() {
  const { courseId, playlistId, contestId, questionId } = useParams<{
    courseId: string
    playlistId: string
    contestId: string
    questionId: string
  }>()
  const navigate = useNavigate()
  const { isAuthenticated } = useAuth()

  const [question, setQuestion] = useState<CodingQuestion | null>(null)
  const [sampleCases, setSampleCases] = useState<TestCase[]>([])
  const [loading, setLoading] = useState(true)

  const [language, setLanguage] = useState('python')
  const [code, setCode] = useState(getLanguage('python').defaultCode)
  const [running, setRunning] = useState(false)
  const [submitting, setSubmitting] = useState(false)
  const [error, setError] = useState('')
  const [submissionStatus, setSubmissionStatus] = useState<SubmissionStatus | null>(null)
  const [results, setResults] = useState<SubmissionResult[]>([])

  const basePath = `/courses/${courseId}/playlists/${playlistId}/coding/${contestId}`

  const loadQuestion = useCallback(async () => {
    if (!questionId) return
    setLoading(true)
    setError('')

    try {
      const [questionRes, tcRes] = await Promise.all([
        codeApi.getQuestion(questionId),
        codeApi.getSampleTestCases(questionId).catch(() => null),
      ])
      setQuestion(questionRes.data?.question ?? null)
      setSampleCases(tcRes?.data?.testcases ?? [])
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Failed to load question')
      setQuestion(null)
    } finally {
      setLoading(false)
    }
  }, [questionId])

  useEffect(() => {
    loadQuestion()
  }, [loadQuestion])

  const onLanguageChange = (langId: string) => {
    setLanguage(langId)
    setCode(getLanguage(langId).defaultCode)
  }

  const runSubmission = async (mode: 'sample' | 'final') => {
    if (!questionId || !code.trim()) return
    if (!isAuthenticated) {
      navigate('/login')
      return
    }

    const lang = getLanguage(language)
    setError('')
    setSubmissionStatus(null)
    setResults([])

    if (!lang.supported) {
      setRunning(true)
      try {
        const mock = await mockUnsupportedSubmit()
        setSubmissionStatus(mock.status)
        setResults(mock.results)
        setError(`${lang.label} execution is not available yet. Use Python for now.`)
      } finally {
        setRunning(false)
        setSubmitting(false)
      }
      return
    }

    const setBusy = mode === 'sample' ? setRunning : setSubmitting
    setBusy(true)

    try {
      const submitRes =
        mode === 'sample'
          ? await codeApi.submitSample({
              question_id: questionId,
              code,
              language,
            })
          : await codeApi.submitFinal({
              question_id: questionId,
              code,
              language,
            })

      const submissionId = submitRes.data?.submission?.submission_id
      if (!submissionId) throw new Error('No submission id returned')

      const status = await pollSubmission(submissionId)
      setSubmissionStatus(status)

      const resultRes = await codeApi.getSubmissionResult(submissionId)
      setResults(resultRes.data?.results ?? [])
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Submission failed')
    } finally {
      setRunning(false)
      setSubmitting(false)
    }
  }

  if (loading) return <LoadingSpinner className="min-h-[40vh]" />

  if (!question) {
    return (
      <Card>
        <p className="text-center text-ink-500">Question not found.</p>
        <Link to={basePath} className="mt-4 block text-center text-sage-600">
          Back
        </Link>
      </Card>
    )
  }

  return (
    <div className="animate-fade-in space-y-6">
      <Breadcrumbs
        items={[
          { label: 'Coding practice', to: basePath },
          { label: question.title },
        ]}
      />

      <div className="grid gap-6 lg:grid-cols-2">
        <div className="space-y-4">
          <Card className="space-y-4">
            <div className="flex flex-wrap items-center gap-2">
              <h1 className="font-display text-xl font-semibold">{question.title}</h1>
              <span className="rounded-full bg-mist-100 px-2 py-0.5 text-xs capitalize dark:bg-mist-900/40">
                {question.difficulty}
              </span>
            </div>
            <div className="prose prose-sm max-w-none dark:prose-invert">
              <p className="whitespace-pre-wrap text-sm leading-relaxed text-ink-700 dark:text-ink-200">
                {question.statement}
              </p>
            </div>
            {question.constraints ? (
              <div>
                <h3 className="text-sm font-medium">Constraints</h3>
                <p className="text-sm text-ink-500">{question.constraints}</p>
              </div>
            ) : null}
            {question.input_format ? (
              <div>
                <h3 className="text-sm font-medium">Input format</h3>
                <p className="whitespace-pre-wrap text-sm text-ink-500">{question.input_format}</p>
              </div>
            ) : null}
            {question.output_format ? (
              <div>
                <h3 className="text-sm font-medium">Output format</h3>
                <p className="whitespace-pre-wrap text-sm text-ink-500">{question.output_format}</p>
              </div>
            ) : null}
            <p className="text-xs text-ink-400">
              Time limit: {question.time_limit_ms} ms · Memory: {question.memory_limit_mb} MB
            </p>
          </Card>

          {sampleCases.length > 0 ? (
            <Card className="space-y-3">
              <h2 className="text-sm font-semibold">Sample test cases</h2>
              {sampleCases.map((tc) => (
                <div key={tc.testcase_id} className="rounded-lg bg-cream-100/80 p-3 font-mono text-xs dark:bg-ink-800/50">
                  <p>
                    <span className="text-ink-400">Input:</span> {tc.input}
                  </p>
                  <p className="mt-1">
                    <span className="text-ink-400">Output:</span> {tc.expected_output}
                  </p>
                </div>
              ))}
            </Card>
          ) : null}
        </div>

        <div className="space-y-4">
          <div className="flex flex-wrap items-center gap-3">
            <select
              value={language}
              onChange={(e) => onLanguageChange(e.target.value)}
              className="rounded-xl border border-ink-200/80 bg-white px-3 py-2 text-sm dark:border-ink-700 dark:bg-ink-900/50"
            >
              {CODING_LANGUAGES.map((l) => (
                <option key={l.id} value={l.id}>
                  {l.label}
                  {!l.supported ? ' (soon)' : ''}
                </option>
              ))}
            </select>
            <Button
              variant="secondary"
              size="sm"
              isLoading={running}
              disabled={submitting}
              onClick={() => runSubmission('sample')}
            >
              <Play className="size-4" />
              Run sample
            </Button>
            <Button
              size="sm"
              isLoading={submitting}
              disabled={running}
              onClick={() => runSubmission('final')}
            >
              <Send className="size-4" />
              Submit
            </Button>
          </div>

          {!getLanguage(language).supported ? (
            <p className="text-xs text-amber-600 dark:text-amber-400">
              {getLanguage(language).label} uses a placeholder runner until the execution service
              supports it. Python is fully wired to the queue.
            </p>
          ) : null}

          <CodeEditor language={language} value={code} onChange={setCode} />

          {error ? <p className="text-sm text-red-500">{error}</p> : null}

          <SubmissionResultsPanel
            status={submissionStatus}
            results={results}
            loading={running || submitting}
          />
        </div>
      </div>
    </div>
  )
}

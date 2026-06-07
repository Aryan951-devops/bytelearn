import { useState } from 'react'
import type { SubmissionResult, SubmissionStatus } from '@/types'
import { Card } from '@/components/ui/Card'

interface SubmissionResultsPanelProps {
  status: SubmissionStatus | null
  results: SubmissionResult[]
  loading?: boolean
}

function verdictStyle(verdict: string, isPassed: boolean) {
  const v = verdict.toUpperCase()
  if (isPassed || v === 'AC' || v === 'ACCEPTED') {
    return 'bg-sage-100 text-sage-800 dark:bg-sage-900/50 dark:text-sage-200'
  }
  if (v === 'WA' || v === 'WRONG ANSWER') {
    return 'bg-amber-100 text-amber-800 dark:bg-amber-900/40 dark:text-amber-200'
  }
  if (v === 'TLE' || v === 'TIME LIMIT EXCEEDED') {
    return 'bg-orange-100 text-orange-800 dark:bg-orange-900/40 dark:text-orange-200'
  }
  if (v === 'MLE' || v === 'MEMORY LIMIT EXCEEDED') {
    return 'bg-purple-100 text-purple-800 dark:bg-purple-900/40 dark:text-purple-200'
  }
  if (v === 'RE' || v === 'RUNTIME ERROR' || v.includes('ERROR')) {
    return 'bg-red-100 text-red-800 dark:bg-red-900/40 dark:text-red-200'
  }
  return 'bg-ink-100 text-ink-700 dark:bg-ink-800 dark:text-ink-200'
}

function ResultField({
  label,
  value,
  variant = 'default',
}: {
  label: string
  value: string | null | undefined
  variant?: 'default' | 'error'
}) {
  if (value == null || value === '') return null
  return (
    <div className="space-y-1">
      <p className="text-xs font-medium uppercase tracking-wide text-ink-400">{label}</p>
      <pre
        className={`max-h-32 overflow-auto rounded-lg px-3 py-2 font-mono text-xs whitespace-pre-wrap ${
          variant === 'error'
            ? 'bg-red-50 text-red-700 dark:bg-red-950/30 dark:text-red-300'
            : 'bg-cream-100/80 text-ink-700 dark:bg-ink-800/60 dark:text-ink-200'
        }`}
      >
        {value}
      </pre>
    </div>
  )
}

export function SubmissionResultsPanel({
  status,
  results,
  loading,
}: SubmissionResultsPanelProps) {
  const [activeIndex, setActiveIndex] = useState(0)

  if (loading) {
    return (
      <Card className="py-8 text-center text-sm text-ink-500">
        <span className="inline-block size-5 animate-spin rounded-full border-2 border-sage-500 border-t-transparent" />
        <p className="mt-2">Running your code…</p>
      </Card>
    )
  }

  if (!status) return null

  const passed = status.passed_cases
  const total = status.total_cases
  const allPassed = status.status === 'COMPLETED' && passed === total && total > 0
  const active = results[activeIndex]

  return (
    <Card className="space-y-4">
      <div className="flex flex-wrap items-center gap-3">
        <span
          className={`rounded-full px-3 py-1 text-xs font-semibold uppercase tracking-wide ${
            status.status === 'COMPLETED'
              ? allPassed
                ? 'bg-sage-100 text-sage-700 dark:bg-sage-900/50 dark:text-sage-300'
                : 'bg-amber-100 text-amber-700 dark:bg-amber-900/40 dark:text-amber-300'
              : status.status === 'FAILED'
                ? 'bg-red-100 text-red-700 dark:bg-red-900/40 dark:text-red-300'
                : 'bg-mist-100 text-mist-700 dark:bg-mist-900/40 dark:text-mist-300'
          }`}
        >
          {status.status}
        </span>
        {total > 0 ? (
          <span className="text-sm text-ink-600 dark:text-ink-300">
            {passed} / {total} passed
          </span>
        ) : null}
      </div>

      {results.length > 0 ? (
        <div className="space-y-3">
          <div className="flex flex-wrap gap-1.5">
            {results.map((r, i) => (
              <button
                key={`${r.testcase_id}-${i}`}
                type="button"
                onClick={() => setActiveIndex(i)}
                className={`rounded-lg px-3 py-1.5 text-xs font-medium transition-colors ${
                  i === activeIndex
                    ? 'bg-sage-600 text-white shadow-sm dark:bg-sage-500'
                    : r.is_passed
                      ? 'bg-sage-100 text-sage-700 hover:bg-sage-200 dark:bg-sage-900/40 dark:text-sage-200'
                      : 'bg-ink-100 text-ink-600 hover:bg-ink-200 dark:bg-ink-800 dark:text-ink-300'
                }`}
              >
                Test {i + 1}
              </button>
            ))}
          </div>

          {active ? (
            <div className="space-y-3 rounded-xl border border-ink-200/60 p-4 dark:border-ink-700">
              <div className="flex flex-wrap items-center justify-between gap-2">
                <span
                  className={`rounded-full px-2.5 py-0.5 text-xs font-semibold uppercase ${verdictStyle(active.verdict, active.is_passed)}`}
                >
                  {active.verdict}
                </span>
                <div className="flex gap-3 text-xs text-ink-400">
                  {active.runtime_ms != null ? <span>{active.runtime_ms} ms</span> : null}
                  {active.memory_kb != null ? <span>{active.memory_kb} KB</span> : null}
                </div>
              </div>

              <ResultField label="Input" value={active.input} />
              <ResultField label="Expected output" value={active.expected_output} />
              <ResultField label="Actual output" value={active.actual_output} />
              <ResultField label="Error" value={active.error_output} variant="error" />
            </div>
          ) : null}
        </div>
      ) : status.status === 'COMPLETED' ? (
        <p className="text-sm text-ink-500">No per-testcase results returned.</p>
      ) : null}
    </Card>
  )
}

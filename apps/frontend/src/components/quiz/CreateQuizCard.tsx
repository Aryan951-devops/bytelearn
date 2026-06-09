import { type FormEvent, useState } from 'react'
import { ClipboardList, Plus, Trash2 } from 'lucide-react'
import { quizApi } from '@/lib/api'
import type { CreateQuizQuestionInput, QuizQuestionType } from '@/types'
import { Button } from '@/components/ui/Button'
import { Card } from '@/components/ui/Card'
import { Input } from '@/components/ui/Input'

const QUESTION_TYPES: { value: QuizQuestionType; label: string }[] = [
  { value: 'mcq', label: 'Single choice (MCQ)' },
  { value: 'multiple', label: 'Multiple choice' },
  { value: 'true_false', label: 'True / False' },
  { value: 'one_word', label: 'One word answer' },
]

function defaultQuestion(type: QuizQuestionType = 'mcq'): CreateQuizQuestionInput {
  if (type === 'true_false') {
    return {
      type,
      question: '',
      options: ['True', 'False'],
      correct_options: [0],
      correct_answer: '',
      marks: 1,
      negative_marks: 0,
    }
  }
  if (type === 'one_word') {
    return {
      type,
      question: '',
      options: [],
      correct_options: [],
      correct_answer: '',
      marks: 1,
      negative_marks: 0,
    }
  }
  return {
    type,
    question: '',
    options: ['', ''],
    correct_options: [0],
    correct_answer: '',
    marks: 1,
    negative_marks: 0,
  }
}

export function CreateQuizCard({
  playlistId,
  onCreated,
}: {
  playlistId: string
  onCreated: (quizId: string) => void
}) {
  const [title, setTitle] = useState('Quiz')
  const [durationMinutes, setDurationMinutes] = useState(30)
  const [questions, setQuestions] = useState<CreateQuizQuestionInput[]>([
    defaultQuestion(),
  ])
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState('')

  const updateQuestion = (index: number, patch: Partial<CreateQuizQuestionInput>) => {
    setQuestions((prev) =>
      prev.map((q, i) => {
        if (i !== index) return q
        const next = { ...q, ...patch }
        if (patch.type) {
          const base = defaultQuestion(patch.type)
          return { ...base, question: q.question, marks: q.marks, negative_marks: q.negative_marks }
        }
        return next
      }),
    )
  }

  const addQuestion = () => {
    setQuestions((prev) => [...prev, defaultQuestion()])
  }

  const removeQuestion = (index: number) => {
    setQuestions((prev) => (prev.length <= 1 ? prev : prev.filter((_, i) => i !== index)))
  }

  const updateOption = (qIndex: number, optIndex: number, value: string) => {
    setQuestions((prev) =>
      prev.map((q, i) => {
        if (i !== qIndex) return q
        const options = [...q.options]
        options[optIndex] = value
        return { ...q, options }
      }),
    )
  }

  const addOption = (qIndex: number) => {
    setQuestions((prev) =>
      prev.map((q, i) => (i === qIndex ? { ...q, options: [...q.options, ''] } : q)),
    )
  }

  const removeOption = (qIndex: number, optIndex: number) => {
    setQuestions((prev) =>
      prev.map((q, i) => {
        if (i !== qIndex || q.options.length <= 2) return q
        const options = q.options.filter((_, oi) => oi !== optIndex)
        const correct_options = q.correct_options
          .filter((ci) => ci !== optIndex)
          .map((ci) => (ci > optIndex ? ci - 1 : ci))
        return { ...q, options, correct_options }
      }),
    )
  }

  const toggleCorrectOption = (qIndex: number, optIndex: number, multiple: boolean) => {
    setQuestions((prev) =>
      prev.map((q, i) => {
        if (i !== qIndex) return q
        if (multiple) {
          const set = new Set(q.correct_options)
          if (set.has(optIndex)) set.delete(optIndex)
          else set.add(optIndex)
          return { ...q, correct_options: [...set].sort((a, b) => a - b) }
        }
        return { ...q, correct_options: [optIndex] }
      }),
    )
  }

  const handleSubmit = async (e: FormEvent) => {
    e.preventDefault()
    setLoading(true)
    setError('')

    const payload = questions.map((q) => ({
      ...q,
      question: q.question.trim(),
      options: q.type === 'one_word' ? [] : q.options.map((o) => o.trim()).filter(Boolean),
      correct_answer: q.correct_answer.trim(),
      explanation: q.explanation?.trim() || undefined,
    }))

    for (const q of payload) {
      if (!q.question) {
        setError('Every question needs text.')
        setLoading(false)
        return
      }
      if (q.type !== 'one_word' && q.options.length < 2) {
        setError('Choice questions need at least two options.')
        setLoading(false)
        return
      }
      if (q.type === 'one_word' && !q.correct_answer) {
        setError('One-word questions need a correct answer.')
        setLoading(false)
        return
      }
      if (q.type !== 'one_word' && q.correct_options.length === 0) {
        setError('Mark at least one correct option per question.')
        setLoading(false)
        return
      }
    }

    try {
      const res = await quizApi.createQuiz({
        title: title.trim(),
        playlist_id: playlistId,
        duration_minutes: durationMinutes,
        questions: payload,
      })
      const quizId = res.data?.quiz?.quiz_id
      if (quizId) onCreated(quizId)
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Failed to create quiz')
    } finally {
      setLoading(false)
    }
  }

  return (
    <Card className="space-y-4">
      <h3 className="flex items-center gap-2 font-display text-lg font-semibold">
        <ClipboardList className="size-5 text-mist-500" />
        Create quiz
      </h3>
      <form onSubmit={handleSubmit} className="space-y-5">
        <div className="grid gap-3 sm:grid-cols-2">
          <Input label="Title" required value={title} onChange={(e) => setTitle(e.target.value)} />
          <Input
            label="Duration (minutes)"
            type="number"
            min={1}
            required
            value={durationMinutes}
            onChange={(e) => setDurationMinutes(Number(e.target.value))}
          />
        </div>

        <div className="space-y-4">
          <div className="flex items-center justify-between">
            <p className="text-sm font-medium text-ink-700 dark:text-ink-200">Questions</p>
            <Button type="button" variant="ghost" size="sm" onClick={addQuestion}>
              <Plus className="size-4" />
              Add question
            </Button>
          </div>

          {questions.map((q, qIndex) => (
            <Card key={qIndex} className="space-y-3 border border-ink-200/60 dark:border-ink-700">
              <div className="flex items-start justify-between gap-2">
                <p className="text-sm font-semibold text-ink-600 dark:text-ink-300">
                  Question {qIndex + 1}
                </p>
                {questions.length > 1 ? (
                  <button
                    type="button"
                    onClick={() => removeQuestion(qIndex)}
                    className="text-ink-400 hover:text-red-500"
                    aria-label="Remove question"
                  >
                    <Trash2 className="size-4" />
                  </button>
                ) : null}
              </div>

              <div className="grid gap-3 sm:grid-cols-3">
                <div className="space-y-1.5 sm:col-span-1">
                  <label className="block text-sm font-medium text-ink-700 dark:text-ink-200">
                    Type
                  </label>
                  <select
                    value={q.type}
                    onChange={(e) =>
                      updateQuestion(qIndex, { type: e.target.value as QuizQuestionType })
                    }
                    className="w-full rounded-xl border border-ink-200/80 bg-white/80 px-3 py-2.5 text-sm dark:border-ink-700 dark:bg-ink-900/50"
                  >
                    {QUESTION_TYPES.map((t) => (
                      <option key={t.value} value={t.value}>
                        {t.label}
                      </option>
                    ))}
                  </select>
                </div>
                <div className="sm:col-span-2">
                  <Input
                    label="Question text"
                    required
                    value={q.question}
                    onChange={(e) => updateQuestion(qIndex, { question: e.target.value })}
                  />
                </div>
              </div>

              <div className="grid gap-3 sm:grid-cols-2">
                <Input
                  label="Marks"
                  type="number"
                  min={0}
                  value={q.marks}
                  onChange={(e) => updateQuestion(qIndex, { marks: Number(e.target.value) })}
                />
                <Input
                  label="Negative marks"
                  type="number"
                  min={0}
                  value={q.negative_marks}
                  onChange={(e) =>
                    updateQuestion(qIndex, { negative_marks: Number(e.target.value) })
                  }
                />
              </div>

              {q.type === 'one_word' ? (
                <Input
                  label="Correct answer"
                  required
                  value={q.correct_answer}
                  onChange={(e) => updateQuestion(qIndex, { correct_answer: e.target.value })}
                />
              ) : (
                <div className="space-y-2">
                  <div className="flex items-center justify-between">
                    <p className="text-sm font-medium text-ink-700 dark:text-ink-200">Options</p>
                    {q.type !== 'true_false' ? (
                      <Button type="button" variant="ghost" size="sm" onClick={() => addOption(qIndex)}>
                        <Plus className="size-3.5" />
                        Option
                      </Button>
                    ) : null}
                  </div>
                  {q.options.map((opt, optIndex) => (
                    <div key={optIndex} className="flex items-center gap-2">
                      <input
                        type={q.type === 'multiple' ? 'checkbox' : 'radio'}
                        name={`q-${qIndex}-correct`}
                        checked={q.correct_options.includes(optIndex)}
                        onChange={() =>
                          toggleCorrectOption(qIndex, optIndex, q.type === 'multiple')
                        }
                        className="size-4 accent-sage-600"
                        title="Mark as correct"
                      />
                      <input
                        value={opt}
                        disabled={q.type === 'true_false'}
                        onChange={(e) => updateOption(qIndex, optIndex, e.target.value)}
                        placeholder={`Option ${optIndex + 1}`}
                        className="flex-1 rounded-lg border border-ink-200/80 bg-white/80 px-3 py-2 text-sm dark:border-ink-700 dark:bg-ink-900/50"
                      />
                      {q.type !== 'true_false' && q.options.length > 2 ? (
                        <button
                          type="button"
                          onClick={() => removeOption(qIndex, optIndex)}
                          className="text-ink-400 hover:text-red-500"
                        >
                          <Trash2 className="size-4" />
                        </button>
                      ) : null}
                    </div>
                  ))}
                  <p className="text-xs text-ink-400">
                    Select the correct {q.type === 'multiple' ? 'options' : 'option'}.
                  </p>
                </div>
              )}

              <Input
                label="Explanation (optional)"
                value={q.explanation ?? ''}
                onChange={(e) => updateQuestion(qIndex, { explanation: e.target.value })}
              />
            </Card>
          ))}
        </div>

        {error ? <p className="text-sm text-red-500">{error}</p> : null}
        <Button type="submit" isLoading={loading}>
          Create quiz
        </Button>
      </form>
    </Card>
  )
}

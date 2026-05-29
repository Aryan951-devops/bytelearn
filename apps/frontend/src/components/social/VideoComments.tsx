import { type FormEvent, useEffect, useState } from 'react'
import { MessageCircle, Send } from 'lucide-react'
import { commentApi } from '@/lib/socialApi'
import type { Comment } from '@/types'
import { useAuth } from '@/context/AuthContext'
import { Button } from '@/components/ui/Button'
import { Card } from '@/components/ui/Card'

interface VideoCommentsProps {
  videoId: string
}

export function VideoComments({ videoId }: VideoCommentsProps) {
  const { isAuthenticated, user } = useAuth()
  const [comments, setComments] = useState<Comment[]>([])
  const [text, setText] = useState('')
  const [loading, setLoading] = useState(true)
  const [submitting, setSubmitting] = useState(false)

  const load = () => {
    setLoading(true)
    commentApi
      .getForVideo(videoId)
      .then(setComments)
      .finally(() => setLoading(false))
  }

  useEffect(() => {
    load()
  }, [videoId])

  const handleSubmit = async (e: FormEvent) => {
    e.preventDefault()
    if (!text.trim() || !user) return
    setSubmitting(true)
    try {
      const comment = await commentApi.create(videoId, text.trim(), {
        user_id: user.user_id,
        name: user.name,
      })
      setComments((prev) => [comment, ...prev])
      setText('')
    } finally {
      setSubmitting(false)
    }
  }

  return (
    <Card className="space-y-4">
      <h2 className="flex items-center gap-2 font-display text-lg font-semibold text-ink-900 dark:text-ink-50">
        <MessageCircle className="size-5 text-sage-500" />
        Comments
        <span className="text-sm font-normal text-ink-400">({comments.length})</span>
      </h2>

      {isAuthenticated ? (
        <form onSubmit={handleSubmit} className="flex gap-2">
          <input
            type="text"
            placeholder="Add a comment…"
            value={text}
            onChange={(e) => setText(e.target.value)}
            className="min-w-0 flex-1 rounded-xl border border-ink-200/80 bg-white/80 px-4 py-2.5 text-sm focus:border-sage-400 focus:outline-none focus:ring-2 focus:ring-sage-300/40 dark:border-ink-700 dark:bg-ink-900/50 dark:text-ink-50"
          />
          <Button type="submit" size="sm" isLoading={submitting} disabled={!text.trim()}>
            <Send className="size-4" />
          </Button>
        </form>
      ) : (
        <p className="text-sm text-ink-500">Sign in to join the discussion.</p>
      )}

      {loading ? (
        <p className="text-sm text-ink-400">Loading comments…</p>
      ) : comments.length === 0 ? (
        <p className="text-sm text-ink-400">No comments yet. Be the first!</p>
      ) : (
        <ul className="max-h-80 space-y-3 overflow-y-auto">
          {comments.map((c) => (
            <li
              key={c.comment_id}
              className="rounded-xl bg-cream-100/80 px-4 py-3 dark:bg-ink-800/50"
            >
              <div className="flex items-baseline justify-between gap-2">
                <span className="text-sm font-medium text-ink-800 dark:text-ink-100">
                  {c.user_name}
                </span>
                <time className="shrink-0 text-xs text-ink-400">
                  {new Date(c.created_at).toLocaleDateString()}
                </time>
              </div>
              <p className="mt-1 text-sm text-ink-600 dark:text-ink-300">{c.content}</p>
            </li>
          ))}
        </ul>
      )}

      <p className="text-xs text-ink-400">
        Comments use placeholder API paths until backend routes are enabled.
      </p>
    </Card>
  )
}

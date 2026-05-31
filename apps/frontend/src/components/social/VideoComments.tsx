import { type FormEvent, useEffect, useState } from 'react'
import { MessageCircle, Pencil, Send, Trash2, X } from 'lucide-react'
import { commentApi } from '@/lib/api'
import type { Comment } from '@/types'
import { useAuth } from '@/context/AuthContext'
import { Button } from '@/components/ui/Button'
import { Card } from '@/components/ui/Card'
import { CommentLikeButton } from '@/components/social/CommentLikeButton'

interface VideoCommentsProps {
  videoId: string
}

export function VideoComments({ videoId }: VideoCommentsProps) {
  const { isAuthenticated, user } = useAuth()
  const [comments, setComments] = useState<Comment[]>([])
  const [text, setText] = useState('')
  const [loading, setLoading] = useState(true)
  const [submitting, setSubmitting] = useState(false)
  const [error, setError] = useState('')
  const [editingId, setEditingId] = useState<string | null>(null)
  const [editText, setEditText] = useState('')
  const [savingId, setSavingId] = useState<string | null>(null)
  const [deletingId, setDeletingId] = useState<string | null>(null)

  const load = () => {
    setLoading(true)
    setError('')
    commentApi
      .getForVideo(videoId)
      .then((res) => setComments(res.data?.comments ?? []))
      .catch((err) => {
        setComments([])
        setError(err instanceof Error ? err.message : 'Failed to load comments')
      })
      .finally(() => setLoading(false))
  }

  useEffect(() => {
    load()
  }, [videoId])

  const handleSubmit = async (e: FormEvent) => {
    e.preventDefault()
    if (!text.trim() || !user) return
    setSubmitting(true)
    setError('')
    try {
      const res = await commentApi.create(videoId, { content: text.trim() })
      const created = res.data?.comment
      if (created) {
        setComments((prev) => [
          { ...created, commented_by: created.commented_by ?? user.username },
          ...prev,
        ])
      }
      setText('')
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Failed to post comment')
    } finally {
      setSubmitting(false)
    }
  }

  const startEdit = (comment: Comment) => {
    setEditingId(comment.comment_id)
    setEditText(comment.content)
    setError('')
  }

  const cancelEdit = () => {
    setEditingId(null)
    setEditText('')
  }

  const saveEdit = async (commentId: string) => {
    if (!editText.trim()) return
    setSavingId(commentId)
    setError('')
    try {
      const res = await commentApi.update(commentId, { content: editText.trim() })
      const updated = res.data?.comment
      if (updated) {
        setComments((prev) =>
          prev.map((c) =>
            c.comment_id === commentId
              ? {
                  ...updated,
                  commented_by: c.commented_by ?? updated.commented_by,
                }
              : c,
          ),
        )
      }
      cancelEdit()
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Failed to update comment')
    } finally {
      setSavingId(null)
    }
  }

  const deleteComment = async (commentId: string) => {
    if (!confirm('Delete this comment?')) return
    setDeletingId(commentId)
    setError('')
    try {
      await commentApi.delete(commentId)
      setComments((prev) => prev.filter((c) => c.comment_id !== commentId))
      if (editingId === commentId) cancelEdit()
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Failed to delete comment')
    } finally {
      setDeletingId(null)
    }
  }

  const isOwner = (comment: Comment) =>
    !!user && user.user_id === comment.user_id

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

      {error ? <p className="text-sm text-red-500">{error}</p> : null}

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
              <div className="flex items-start justify-between gap-2">
                <span className="text-sm font-medium text-ink-800 dark:text-ink-100">
                  {c.commented_by ?? 'User'}
                </span>
                <div className="flex shrink-0 items-center gap-1">
                  <time className="text-xs text-ink-400">
                    {new Date(c.created_at).toLocaleDateString()}
                  </time>
                  {isOwner(c) && editingId !== c.comment_id ? (
                    <>
                      <button
                        type="button"
                        onClick={() => startEdit(c)}
                        className="rounded-lg p-1.5 text-ink-400 transition-colors hover:bg-ink-200/50 hover:text-sage-600 dark:hover:bg-ink-700 dark:hover:text-sage-300"
                        aria-label="Edit comment"
                      >
                        <Pencil className="size-3.5" />
                      </button>
                      <button
                        type="button"
                        onClick={() => deleteComment(c.comment_id)}
                        disabled={deletingId === c.comment_id}
                        className="rounded-lg p-1.5 text-ink-400 transition-colors hover:bg-red-50 hover:text-red-500 disabled:opacity-50 dark:hover:bg-red-950/30"
                        aria-label="Delete comment"
                      >
                        <Trash2 className="size-3.5" />
                      </button>
                    </>
                  ) : null}
                </div>
              </div>

              {editingId === c.comment_id ? (
                <div className="mt-2 space-y-2">
                  <textarea
                    value={editText}
                    onChange={(e) => setEditText(e.target.value)}
                    rows={2}
                    className="w-full rounded-xl border border-ink-200/80 bg-white px-3 py-2 text-sm text-ink-900 focus:border-sage-400 focus:outline-none focus:ring-2 focus:ring-sage-300/40 dark:border-ink-600 dark:bg-ink-900 dark:text-ink-50"
                  />
                  <div className="flex gap-2">
                    <Button
                      size="sm"
                      isLoading={savingId === c.comment_id}
                      disabled={!editText.trim()}
                      onClick={() => saveEdit(c.comment_id)}
                    >
                      Save
                    </Button>
                    <Button size="sm" variant="ghost" onClick={cancelEdit}>
                      <X className="size-4" />
                      Cancel
                    </Button>
                  </div>
                </div>
              ) : (
                <p className="mt-1 text-sm text-ink-600 dark:text-ink-300">{c.content}</p>
              )}

              {editingId !== c.comment_id ? (
                <div className="mt-2">
                  <CommentLikeButton commentId={c.comment_id} />
                </div>
              ) : null}
            </li>
          ))}
        </ul>
      )}
    </Card>
  )
}

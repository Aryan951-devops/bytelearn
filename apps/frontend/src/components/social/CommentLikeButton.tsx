import { useCallback, useEffect, useState } from 'react'
import { Heart } from 'lucide-react'
import { likeApi } from '@/lib/api'
import { useAuth } from '@/context/AuthContext'

interface CommentLikeButtonProps {
  commentId: string
}

export function CommentLikeButton({ commentId }: CommentLikeButtonProps) {
  const { isAuthenticated } = useAuth()
  const [count, setCount] = useState(0)
  const [liked, setLiked] = useState(false)
  const [loading, setLoading] = useState(false)

  const load = useCallback(async () => {
    const totalRes = await likeApi.getTotalCommentLikes(commentId)
    setCount(totalRes.data?.total_likes ?? 0)

    if (isAuthenticated) {
      try {
        const checkRes = await likeApi.checkComment(commentId)
        setLiked(checkRes.data?.liked ?? false)
      } catch {
        setLiked(false)
      }
    } else {
      setLiked(false)
    }
  }, [commentId, isAuthenticated])

  useEffect(() => {
    load()
  }, [load])

  const toggle = async () => {
    if (!isAuthenticated) return
    setLoading(true)
    try {
      const res = await likeApi.toggleComment(commentId)
      setLiked(res.data?.liked ?? false)
      const totalRes = await likeApi.getTotalCommentLikes(commentId)
      setCount(totalRes.data?.total_likes ?? 0)
    } finally {
      setLoading(false)
    }
  }

  return (
    <button
      type="button"
      disabled={!isAuthenticated || loading}
      onClick={toggle}
      title={isAuthenticated ? undefined : 'Sign in to like'}
      className={`inline-flex items-center gap-1 rounded-lg px-2 py-1 text-xs transition-colors ${
        liked
          ? 'text-rose-600 bg-rose-50 dark:bg-rose-950/30 dark:text-rose-400'
          : 'text-ink-400 hover:bg-ink-100 dark:hover:bg-ink-800'
      } disabled:cursor-not-allowed disabled:opacity-50`}
    >
      <Heart className={`size-3.5 ${liked ? 'fill-current' : ''}`} />
      {count}
    </button>
  )
}

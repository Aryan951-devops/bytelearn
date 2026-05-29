import { useEffect, useState } from 'react'
import { Heart } from 'lucide-react'
import { likeApi } from '@/lib/socialApi'
import { useAuth } from '@/context/AuthContext'
import { Button } from '@/components/ui/Button'

interface VideoLikeButtonProps {
  videoId: string
}

export function VideoLikeButton({ videoId }: VideoLikeButtonProps) {
  const { isAuthenticated, user } = useAuth()
  const [count, setCount] = useState(0)
  const [liked, setLiked] = useState(false)
  const [loading, setLoading] = useState(false)

  useEffect(() => {
    likeApi.getVideoLikeCount(videoId).then(setCount)
    if (user) {
      likeApi.isLikedByUser(videoId, user.user_id).then(setLiked)
    }
  }, [videoId, user])

  const toggle = async () => {
    if (!isAuthenticated || !user) return
    setLoading(true)
    try {
      const result = await likeApi.toggleVideoLike(videoId, user.user_id)
      setLiked(result.liked)
      setCount(result.count)
    } finally {
      setLoading(false)
    }
  }

  return (
    <Button
      variant={liked ? 'primary' : 'secondary'}
      size="sm"
      disabled={!isAuthenticated || loading}
      onClick={toggle}
      title={isAuthenticated ? undefined : 'Sign in to like'}
    >
      <Heart className={`size-4 ${liked ? 'fill-current' : ''}`} />
      {count.toLocaleString()}
    </Button>
  )
}

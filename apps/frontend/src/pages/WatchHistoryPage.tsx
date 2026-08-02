import { useEffect, useState } from 'react'
import { History } from 'lucide-react'
import { userApi } from '@/lib/api'
import type { VideoMetadata } from '@/types'
import { Card } from '@/components/ui/Card'
import { LoadingSpinner } from '@/components/ui/LoadingSpinner'
import { VideoCard } from '@/components/video/VideoCard'

export function WatchHistoryPage() {
  const [videos, setVideos] = useState<VideoMetadata[]>([])
  const [error, setError] = useState('')
  const [isLoading, setIsLoading] = useState(true)

  useEffect(() => {
    userApi
      .getWatchHistory()
      .then((res) => setVideos(res.data?.videos ?? []))
      .catch((err) =>
        setError(err instanceof Error ? err.message : 'Failed to load watch history'),
      )
      .finally(() => setIsLoading(false))
  }, [])

  return (
    <div className="animate-fade-in space-y-8">
      <div>
        <h1 className="flex items-center gap-2 font-display text-3xl font-semibold text-ink-900 dark:text-ink-50">
          <History className="size-8 text-sage-500" />
          Watch history
        </h1>
        <p className="mt-2 text-ink-500 dark:text-ink-400">
          Videos you have watched recently.
        </p>
      </div>

      {isLoading ? (
        <LoadingSpinner className="min-h-[30vh]" />
      ) : error ? (
        <Card>
          <p className="text-center text-red-500 dark:text-red-400">{error}</p>
        </Card>
      ) : videos.length === 0 ? (
        <Card className="py-12 text-center text-ink-500">
          No watch history yet. Start watching videos to see them here.
        </Card>
      ) : (
        <div className="grid gap-5 sm:grid-cols-2 lg:grid-cols-3">
          {videos.map((video) => (
            <VideoCard
              key={video.video_id}
              video={video}
              linkTo={`/watch/${video.video_id}`}
            />
          ))}
        </div>
      )}
    </div>
  )
}

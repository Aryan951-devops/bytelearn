import { useEffect, useState } from 'react'
import { Video as VideoIcon } from 'lucide-react'
import { videoApi } from '@/lib/api'
import type { Video } from '@/types'
import { Card } from '@/components/ui/Card'
import { LoadingSpinner } from '@/components/ui/LoadingSpinner'
import { VideoCard } from '@/components/video/VideoCard'

export function VideosPage() {
  const [videos, setVideos] = useState<Video[]>([])
  const [error, setError] = useState('')
  const [isLoading, setIsLoading] = useState(true)

  useEffect(() => {
    videoApi
      .getAll()
      .then((res) => setVideos(res.data?.videos ?? []))
      .catch((err) =>
        setError(err instanceof Error ? err.message : 'Failed to load videos'),
      )
      .finally(() => setIsLoading(false))
  }, [])

  return (
    <div className="animate-fade-in space-y-8">
      <div>
        <h1 className="font-display text-3xl font-semibold text-ink-900 dark:text-ink-50">
          Video library
        </h1>
        <p className="mt-2 text-ink-500 dark:text-ink-400">
          Browse all lessons — click to watch
        </p>
      </div>

      {isLoading ? (
        <LoadingSpinner className="min-h-[30vh]" />
      ) : error ? (
        <Card>
          <p className="text-center text-red-500 dark:text-red-400">{error}</p>
        </Card>
      ) : videos.length === 0 ? (
        <Card className="py-16 text-center">
          <VideoIcon className="mx-auto size-12 text-ink-300 dark:text-ink-600" />
          <p className="mt-4 font-medium text-ink-700 dark:text-ink-300">
            No videos yet
          </p>
        </Card>
      ) : (
        <div className="grid gap-6 sm:grid-cols-2 lg:grid-cols-3">
          {videos.map((video) => (
            <VideoCard
              key={video.video_id}
              video={video}
              durationSeconds={video.duration_seconds}
              linkTo={`/watch/${video.video_id}`}
            />
          ))}
        </div>
      )}
    </div>
  )
}

import { useEffect, useState } from 'react'
import { Clock, Eye, Play, Video as VideoIcon } from 'lucide-react'
import { videoApi } from '@/lib/api'
import type { Video } from '@/types'
import { Card } from '@/components/ui/Card'

function formatDuration(seconds: number | string): string {
  const total = typeof seconds === 'string' ? parseInt(seconds, 10) : seconds
  if (!total || Number.isNaN(total)) return '—'
  const m = Math.floor(total / 60)
  const s = total % 60
  return `${m}:${s.toString().padStart(2, '0')}`
}

function VideoCard({ video }: { video: Video }) {
  return (
    <Card hover className="group overflow-hidden p-0">
      <div className="relative aspect-video overflow-hidden bg-ink-100 dark:bg-ink-800">
        {video.thumbnail_url ? (
          <img
            src={video.thumbnail_url}
            alt=""
            className="size-full object-cover transition-transform duration-500 group-hover:scale-105"
          />
        ) : (
          <div className="flex size-full items-center justify-center bg-gradient-to-br from-sage-200 to-mist-200 dark:from-sage-900 dark:to-mist-900">
            <VideoIcon className="size-12 text-sage-500/50" />
          </div>
        )}
        <div className="absolute inset-0 flex items-center justify-center bg-ink-900/0 opacity-0 transition-all group-hover:bg-ink-900/20 group-hover:opacity-100">
          <span className="flex size-12 items-center justify-center rounded-full bg-white/90 text-sage-600 shadow-lg">
            <Play className="size-5 fill-current" />
          </span>
        </div>
      </div>
      <div className="p-5">
        <h3 className="line-clamp-2 font-medium text-ink-900 dark:text-ink-50">
          {video.title}
        </h3>
        {video.description ? (
          <p className="mt-1 line-clamp-2 text-sm text-ink-500 dark:text-ink-400">
            {video.description}
          </p>
        ) : null}
        <div className="mt-3 flex flex-wrap gap-3 text-xs text-ink-400 dark:text-ink-500">
          <span className="inline-flex items-center gap-1">
            <Eye className="size-3.5" />
            {video.views.toLocaleString()} views
          </span>
          <span className="inline-flex items-center gap-1">
            <Clock className="size-3.5" />
            {formatDuration(video.duration_seconds)}
          </span>
        </div>
      </div>
    </Card>
  )
}

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
          Browse all available lessons and courses
        </p>
      </div>

      {isLoading ? (
        <div className="flex min-h-[30vh] items-center justify-center">
          <span className="size-8 animate-spin rounded-full border-2 border-sage-500 border-t-transparent" />
        </div>
      ) : error ? (
        <Card>
          <p className="text-center text-red-500 dark:text-red-400">{error}</p>
          <p className="mt-2 text-center text-sm text-ink-500">
            Make sure the API gateway is running on port 8080.
          </p>
        </Card>
      ) : videos.length === 0 ? (
        <Card className="py-16 text-center">
          <VideoIcon className="mx-auto size-12 text-ink-300 dark:text-ink-600" />
          <p className="mt-4 font-medium text-ink-700 dark:text-ink-300">
            No videos yet
          </p>
          <p className="mt-1 text-sm text-ink-500">
            Videos will appear here once they are added to the platform.
          </p>
        </Card>
      ) : (
        <div className="grid gap-6 sm:grid-cols-2 lg:grid-cols-3">
          {videos.map((video) => (
            <VideoCard key={video.video_id} video={video} />
          ))}
        </div>
      )}
    </div>
  )
}

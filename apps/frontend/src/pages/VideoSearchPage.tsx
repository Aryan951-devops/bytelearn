import { useEffect, useState } from 'react'
import { useSearchParams } from 'react-router-dom'
import { Search, Video as VideoIcon } from 'lucide-react'
import { videoApi } from '@/lib/api'
import type { VideoMetadata } from '@/types'
import { Card } from '@/components/ui/Card'
import { LoadingSpinner } from '@/components/ui/LoadingSpinner'
import { VideoCard } from '@/components/video/VideoCard'

export function VideoSearchPage() {
  const [searchParams] = useSearchParams()
  const query = searchParams.get('q')?.trim() ?? ''

  const [videos, setVideos] = useState<VideoMetadata[]>([])
  const [error, setError] = useState('')
  const [isLoading, setIsLoading] = useState(false)

  useEffect(() => {
    if (!query) {
      setVideos([])
      setError('')
      setIsLoading(false)
      return
    }

    setIsLoading(true)
    setError('')

    videoApi
      .search(query)
      .then((res) => setVideos(res.data?.videos ?? []))
      .catch((err) =>
        setError(err instanceof Error ? err.message : 'Search failed'),
      )
      .finally(() => setIsLoading(false))
  }, [query])

  return (
    <div className="animate-fade-in space-y-8">
      <div>
        <h1 className="flex items-center gap-2 font-display text-3xl font-semibold text-ink-900 dark:text-ink-50">
          <Search className="size-8 text-sage-500" />
          Search results
        </h1>
        {query ? (
          <p className="mt-2 text-ink-500 dark:text-ink-400">
            Showing results for &ldquo;{query}&rdquo;
          </p>
        ) : (
          <p className="mt-2 text-ink-500 dark:text-ink-400">
            Enter a search term in the navbar to find videos.
          </p>
        )}
      </div>

      {!query ? (
        <Card className="py-16 text-center">
          <Search className="mx-auto size-12 text-ink-300 dark:text-ink-600" />
          <p className="mt-4 font-medium text-ink-700 dark:text-ink-300">
            Search for videos by topic or keyword
          </p>
        </Card>
      ) : isLoading ? (
        <LoadingSpinner className="min-h-[30vh]" />
      ) : error ? (
        <Card>
          <p className="text-center text-red-500 dark:text-red-400">{error}</p>
        </Card>
      ) : videos.length === 0 ? (
        <Card className="py-16 text-center">
          <VideoIcon className="mx-auto size-12 text-ink-300 dark:text-ink-600" />
          <p className="mt-4 font-medium text-ink-700 dark:text-ink-300">
            No videos found for &ldquo;{query}&rdquo;
          </p>
        </Card>
      ) : (
        <div className="grid gap-6 sm:grid-cols-2 lg:grid-cols-3">
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

import { useEffect, useState } from 'react'
import { useParams } from 'react-router-dom'
import { Trash2 } from 'lucide-react'
import { playlistApi } from '@/lib/api'
import type { PlaylistWithVideos } from '@/types'
import { Breadcrumbs } from '@/components/ui/Breadcrumbs'
import { Button } from '@/components/ui/Button'
import { Card } from '@/components/ui/Card'
import { LoadingSpinner } from '@/components/ui/LoadingSpinner'
import { VideoCard } from '@/components/video/VideoCard'

export function UserPlaylistPage() {
  const { playlistId } = useParams<{ playlistId: string }>()
  const [playlist, setPlaylist] = useState<PlaylistWithVideos | null>(null)
  const [error, setError] = useState('')
  const [isLoading, setIsLoading] = useState(true)
  const [removingId, setRemovingId] = useState<string | null>(null)

  const load = () => {
    if (!playlistId) return
    playlistApi
      .getUserPlaylist(playlistId)
      .then((res) => setPlaylist(res.data?.playlist ?? null))
      .catch((err) =>
        setError(err instanceof Error ? err.message : 'Failed to load playlist'),
      )
      .finally(() => setIsLoading(false))
  }

  useEffect(() => {
    load()
  }, [playlistId])

  const removeVideo = async (videoId: string) => {
    if (!playlistId) return
    setRemovingId(videoId)
    try {
      await playlistApi.removeVideo(videoId, playlistId)
      setPlaylist((prev) =>
        prev
          ? {
              ...prev,
              videos: prev.videos.filter((v) => v.video_id !== videoId),
            }
          : null,
      )
    } catch {
      /* keep UI stable */
    } finally {
      setRemovingId(null)
    }
  }

  if (isLoading) return <LoadingSpinner className="min-h-[40vh]" />

  if (error || !playlist) {
    return (
      <Card>
        <p className="text-center text-red-500">{error || 'Playlist not found'}</p>
      </Card>
    )
  }

  return (
    <div className="animate-fade-in space-y-8">
      <Breadcrumbs
        items={[
          { label: 'Home', to: '/' },
          { label: 'My playlists', to: '/my-playlists' },
          { label: playlist.title },
        ]}
      />

      <header>
        <h1 className="font-display text-3xl font-semibold text-ink-900 dark:text-ink-50">
          {playlist.title}
        </h1>
        {playlist.description ? (
          <p className="mt-2 text-ink-600 dark:text-ink-300">{playlist.description}</p>
        ) : null}
      </header>

      <section className="space-y-4">
        <h2 className="font-display text-xl font-semibold text-ink-900 dark:text-ink-50">
          Saved videos
        </h2>
        {playlist.videos.length === 0 ? (
          <Card className="py-12 text-center text-ink-500">
            No videos yet. Open any video and use &quot;Save to playlist&quot;.
          </Card>
        ) : (
          <div className="grid gap-5 sm:grid-cols-2 lg:grid-cols-3">
            {playlist.videos.map((video) => (
              <div key={video.video_id} className="relative">
                <VideoCard video={video} linkTo={`/watch/${video.video_id}`} />
                <Button
                  variant="ghost"
                  size="sm"
                  className="absolute right-2 top-2 bg-white/90 dark:bg-ink-900/90"
                  isLoading={removingId === video.video_id}
                  onClick={() => removeVideo(video.video_id)}
                >
                  <Trash2 className="size-4 text-red-500" />
                </Button>
              </div>
            ))}
          </div>
        )}
      </section>
    </div>
  )
}

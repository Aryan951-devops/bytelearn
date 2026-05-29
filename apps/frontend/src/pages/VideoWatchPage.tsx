import { useEffect, useState } from 'react'
import { Link, useNavigate, useParams } from 'react-router-dom'
import { ListPlus } from 'lucide-react'
import { VideoLikeButton } from '@/components/social/VideoLikeButton'
import { VideoComments } from '@/components/social/VideoComments'
import { playlistApi, videoApi } from '@/lib/api'
import type { Playlist, Video } from '@/types'
import { Breadcrumbs } from '@/components/ui/Breadcrumbs'
import { Card } from '@/components/ui/Card'
import { LoadingSpinner } from '@/components/ui/LoadingSpinner'
import { VideoPlayer } from '@/components/video/VideoPlayer'
import { formatDuration } from '@/components/video/VideoCard'
import { useAuth } from '@/context/AuthContext'
import { Button } from '@/components/ui/Button'

export function VideoWatchPage() {
  const { videoId } = useParams<{ videoId: string }>()
  const navigate = useNavigate()
  const { isAuthenticated } = useAuth()
  const [video, setVideo] = useState<Video | null>(null)
  const [userPlaylists, setUserPlaylists] = useState<Playlist[]>([])
  const [error, setError] = useState('')
  const [message, setMessage] = useState('')
  const [isLoading, setIsLoading] = useState(true)
  const [showPlaylistPicker, setShowPlaylistPicker] = useState(false)

  useEffect(() => {
    if (!videoId) return
    videoApi
      .getById(videoId)
      .then((res) => setVideo(res.data?.video ?? null))
      .catch((err) =>
        setError(err instanceof Error ? err.message : 'Failed to load video'),
      )
      .finally(() => setIsLoading(false))
  }, [videoId])

  useEffect(() => {
    if (!isAuthenticated || !showPlaylistPicker) return
    playlistApi
      .getUserPlaylists()
      .then((res) => {
        const all = res.data?.playlist ?? []
        setUserPlaylists(all.filter((p) => p.type === 'user'))
      })
      .catch(() => setUserPlaylists([]))
  }, [isAuthenticated, showPlaylistPicker])

  const addToPlaylist = async (playlistId: string) => {
    if (!videoId) return
    setMessage('')
    try {
      await playlistApi.addVideo(videoId, playlistId)
      setMessage('Added to playlist!')
      setShowPlaylistPicker(false)
    } catch (err) {
      setMessage(err instanceof Error ? err.message : 'Could not add to playlist')
    }
  }

  if (isLoading) return <LoadingSpinner className="min-h-[40vh]" />

  if (error || !video) {
    return (
      <Card>
        <p className="text-center text-red-500">{error || 'Video not found'}</p>
      </Card>
    )
  }

  return (
    <div className="animate-fade-in mx-auto max-w-4xl space-y-6">
      <Breadcrumbs
        items={[
          { label: 'Home', to: '/' },
          { label: 'Videos', to: '/videos' },
          { label: video.title },
        ]}
      />

      <VideoPlayer
        src={video.videofile_url}
        title={video.title}
        poster={video.thumbnail_url}
      />

      <div className="space-y-3">
        <h1 className="font-display text-2xl font-semibold text-ink-900 dark:text-ink-50">
          {video.title}
        </h1>
        {video.description ? (
          <p className="text-ink-600 dark:text-ink-300">{video.description}</p>
        ) : null}
        <div className="flex flex-wrap items-center gap-3">
          <p className="text-sm text-ink-400">
            {video.views.toLocaleString()} views
            {video.duration_seconds
              ? ` · ${formatDuration(video.duration_seconds)}`
              : ''}
          </p>
          <VideoLikeButton videoId={video.video_id} />
        </div>

        {isAuthenticated ? (
          <div className="relative">
            <Button
              variant="secondary"
              size="sm"
              onClick={() => setShowPlaylistPicker((v) => !v)}
            >
              <ListPlus className="size-4" />
              Save to playlist
            </Button>
            {showPlaylistPicker ? (
              <Card className="absolute left-0 top-full z-10 mt-2 w-72 p-3 shadow-lg">
                {userPlaylists.length === 0 ? (
                  <div className="space-y-2 text-sm text-ink-500">
                    <p>No personal playlists yet.</p>
                    <Link to="/my-playlists">
                      <Button size="sm" className="w-full">
                        Create one
                      </Button>
                    </Link>
                  </div>
                ) : (
                  <ul className="max-h-48 space-y-1 overflow-y-auto">
                    {userPlaylists.map((p) => (
                      <li key={p.playlist_id}>
                        <button
                          type="button"
                          onClick={() => addToPlaylist(p.playlist_id)}
                          className="w-full rounded-lg px-3 py-2 text-left text-sm hover:bg-sage-50 dark:hover:bg-ink-800"
                        >
                          {p.title}
                        </button>
                      </li>
                    ))}
                  </ul>
                )}
              </Card>
            ) : null}
            {message ? (
              <p className="mt-2 text-sm text-sage-600 dark:text-sage-400">{message}</p>
            ) : null}
          </div>
        ) : (
          <Button variant="ghost" size="sm" onClick={() => navigate('/login')}>
            Sign in to save to a playlist
          </Button>
        )}
      </div>

      <VideoComments videoId={video.video_id} />
    </div>
  )
}

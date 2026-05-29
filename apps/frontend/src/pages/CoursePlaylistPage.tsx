import { useCallback, useEffect, useState } from 'react'
import { Link, useParams } from 'react-router-dom'
import { Settings, User } from 'lucide-react'
import { courseApi, playlistApi } from '@/lib/api'
import type { CoursePlaylist } from '@/types'
import { Breadcrumbs } from '@/components/ui/Breadcrumbs'
import { Card } from '@/components/ui/Card'
import { LoadingSpinner } from '@/components/ui/LoadingSpinner'
import { VideoCard } from '@/components/video/VideoCard'
import { useAuth } from '@/context/AuthContext'
import { Button } from '@/components/ui/Button'

export function CoursePlaylistPage() {
  const { courseId, playlistId } = useParams<{
    courseId: string
    playlistId: string
  }>()
  const { user } = useAuth()
  const [playlist, setPlaylist] = useState<CoursePlaylist | null>(null)
  const [courseTitle, setCourseTitle] = useState('Course')
  const [error, setError] = useState('')
  const [isLoading, setIsLoading] = useState(true)

  const load = useCallback(async () => {
    if (!playlistId) return
    try {
      const [playlistRes, courseRes] = await Promise.all([
        playlistApi.getCoursePlaylist(playlistId),
        courseId ? courseApi.getById(courseId).catch(() => null) : Promise.resolve(null),
      ])
      setPlaylist(playlistRes.data?.playlist ?? null)
      if (courseRes?.data?.course?.title) {
        setCourseTitle(courseRes.data.course.title)
      }
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Failed to load playlist')
    } finally {
      setIsLoading(false)
    }
  }, [playlistId, courseId])

  useEffect(() => {
    setIsLoading(true)
    load()
  }, [load])

  if (isLoading) return <LoadingSpinner className="min-h-[40vh]" />

  if (error || !playlist) {
    return (
      <Card>
        <p className="text-center text-red-500">{error || 'Playlist not found'}</p>
      </Card>
    )
  }

  const isOwner =
    user?.role === 'educator' &&
    (user.user_id === playlist.educator_user_id ||
      user.user_id === playlist.user_id)

  return (
    <div className="animate-fade-in space-y-8">
      <Breadcrumbs
        items={[
          { label: 'Home', to: '/' },
          { label: courseTitle, to: courseId ? `/courses/${courseId}` : '/' },
          { label: playlist.title },
        ]}
      />

      <header className="space-y-4">
        <div className="flex flex-wrap items-start justify-between gap-3">
          <div>
            <h1 className="font-display text-3xl font-semibold text-ink-900 dark:text-ink-50">
              {playlist.title}
            </h1>
            {playlist.description ? (
              <p className="mt-2 max-w-2xl text-ink-600 dark:text-ink-300">
                {playlist.description}
              </p>
            ) : null}
          </div>
          {isOwner ? (
            <Link to="/educator/dashboard">
              <Button variant="ghost" size="sm">
                <Settings className="size-4" />
                Educator dashboard
              </Button>
            </Link>
          ) : null}
        </div>

        <Card className="flex items-center gap-4">
          {playlist.educator_profile_pic ? (
            <img
              src={playlist.educator_profile_pic}
              alt=""
              className="size-14 rounded-xl object-cover ring-2 ring-sage-100 dark:ring-sage-900"
            />
          ) : (
            <div className="flex size-14 items-center justify-center rounded-xl bg-gradient-to-br from-sage-400 to-mist-400 text-white">
              <User className="size-7" />
            </div>
          )}
          <div>
            <p className="text-xs font-medium uppercase tracking-wide text-ink-400">
              Instructor
            </p>
            <p className="font-medium text-ink-900 dark:text-ink-50">
              {playlist.educator_name}
            </p>
            <p className="text-sm text-sage-600 dark:text-sage-400">
              @{playlist.educator_username}
            </p>
          </div>
        </Card>
      </header>

      <section className="space-y-4">
        <h2 className="font-display text-xl font-semibold text-ink-900 dark:text-ink-50">
          Videos
        </h2>
        {playlist.videos.length === 0 ? (
          <Card className="py-12 text-center text-ink-500">
            {isOwner
              ? 'No videos yet. Upload from the educator dashboard, then add them to this playlist.'
              : 'No videos in this playlist yet.'}
          </Card>
        ) : (
          <div className="grid gap-5 sm:grid-cols-2 lg:grid-cols-3">
            {playlist.videos.map((video) => (
              <VideoCard
                key={video.video_id}
                video={video}
                linkTo={`/watch/${video.video_id}`}
              />
            ))}
          </div>
        )}
      </section>
    </div>
  )
}

import { useCallback, useEffect, useState } from 'react'
import { Link, useNavigate, useParams } from 'react-router-dom'
import { Code2 } from 'lucide-react'
import { codeApi, courseApi, playlistApi } from '@/lib/api'
import type { CodingPracticeSummary } from '@/types'
import { CreateCodingPracticeCard } from '@/pages/CodingPracticePage'
import { Breadcrumbs } from '@/components/ui/Breadcrumbs'
import { Button } from '@/components/ui/Button'
import { Card } from '@/components/ui/Card'
import { LoadingSpinner } from '@/components/ui/LoadingSpinner'
import { useAuth } from '@/context/AuthContext'

export function PlaylistPracticesPage() {
  const { courseId, playlistId } = useParams<{
    courseId: string
    playlistId: string
  }>()
  const navigate = useNavigate()
  const { user } = useAuth()
  const [practices, setPractices] = useState<CodingPracticeSummary[]>([])
  const [playlistTitle, setPlaylistTitle] = useState('Playlist')
  const [courseTitle, setCourseTitle] = useState('Course')
  const [educatorUserId, setEducatorUserId] = useState('')
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState('')

  const load = useCallback(async () => {
    if (!playlistId) return
    setLoading(true)
    setError('')
    try {
      const [practicesRes, playlistRes, courseRes] = await Promise.all([
        codeApi.getPracticesByPlaylist(playlistId),
        playlistApi.getCoursePlaylist(playlistId),
        courseId ? courseApi.getById(courseId).catch(() => null) : Promise.resolve(null),
      ])
      setPractices(practicesRes.data?.practices ?? [])
      const playlist = playlistRes.data?.playlist
      if (playlist) {
        setPlaylistTitle(playlist.title)
        setEducatorUserId(playlist.educator_user_id || playlist.user_id)
      }
      if (courseRes?.data?.course?.title) {
        setCourseTitle(courseRes.data.course.title)
      }
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Failed to load practices')
    } finally {
      setLoading(false)
    }
  }, [playlistId, courseId])

  useEffect(() => {
    load()
  }, [load])

  const isOwner = user?.role === 'educator' && user.user_id === educatorUserId
  const playlistPath = `/courses/${courseId}/playlists/${playlistId}`

  if (loading) return <LoadingSpinner className="min-h-[40vh]" />

  return (
    <div className="animate-fade-in space-y-8">
      <Breadcrumbs
        items={[
          { label: 'Home', to: '/' },
          { label: courseTitle, to: `/courses/${courseId}` },
          { label: playlistTitle, to: playlistPath },
          { label: 'Coding practice' },
        ]}
      />

      <header>
        <h1 className="flex items-center gap-2 font-display text-3xl font-semibold text-ink-900 dark:text-ink-50">
          <Code2 className="size-8 text-sage-500" />
          Coding practice
        </h1>
        <p className="mt-2 text-ink-500 dark:text-ink-400">
          Choose a practice module to start solving problems.
        </p>
      </header>

      {error ? <p className="text-sm text-red-500">{error}</p> : null}

      {practices.length === 0 ? (
        isOwner ? (
          <CreateCodingPracticeCard
            playlistId={playlistId!}
            onCreated={(id) => {
              navigate(
                `/courses/${courseId}/playlists/${playlistId}/coding/${id}`,
              )
            }}
          />
        ) : (
          <Card className="py-12 text-center text-ink-500">
            No coding practice modules available yet.
          </Card>
        )
      ) : (
        <div className="grid gap-4 sm:grid-cols-2">
          {practices.map((practice) => (
            <Card key={practice.contest_id} hover className="flex flex-col gap-3">
              <div>
                <h2 className="font-display text-lg font-semibold text-ink-900 dark:text-ink-50">
                  {practice.title}
                </h2>
                {practice.description ? (
                  <p className="mt-1 text-sm text-ink-500">{practice.description}</p>
                ) : null}
              </div>
              <Link
                to={`/courses/${courseId}/playlists/${playlistId}/coding/${practice.contest_id}`}
              >
                <Button size="sm" className="w-full sm:w-auto">
                  Open problems
                </Button>
              </Link>
            </Card>
          ))}
        </div>
      )}

      {isOwner && practices.length > 0 ? (
        <CreateCodingPracticeCard
          playlistId={playlistId!}
          onCreated={(id) => {
            navigate(`/courses/${courseId}/playlists/${playlistId}/coding/${id}`)
          }}
        />
      ) : null}

      <Link to={playlistPath}>
        <Button variant="ghost">Back to playlist</Button>
      </Link>
    </div>
  )
}

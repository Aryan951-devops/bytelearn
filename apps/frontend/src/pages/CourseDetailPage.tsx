import { useCallback, useEffect, useState } from 'react'
import { useParams } from 'react-router-dom'
import { GraduationCap } from 'lucide-react'
import { courseApi } from '@/lib/api'
import type { CourseWithPlaylists } from '@/types'
import { Breadcrumbs } from '@/components/ui/Breadcrumbs'
import { Card } from '@/components/ui/Card'
import { LoadingSpinner } from '@/components/ui/LoadingSpinner'
import { EducatorCard } from '@/components/educator/EducatorCard'
import { PlaylistCard } from '@/components/playlist/PlaylistCard'
import { CreateCoursePlaylistForm } from '@/components/educator/CreateCoursePlaylistForm'
import { useAuth } from '@/context/AuthContext'

export function CourseDetailPage() {
  const { courseId } = useParams<{ courseId: string }>()
  const { user } = useAuth()
  const [course, setCourse] = useState<CourseWithPlaylists | null>(null)
  const [error, setError] = useState('')
  const [isLoading, setIsLoading] = useState(true)

  const loadCourse = useCallback(() => {
    if (!courseId) return
    courseApi
      .getById(courseId)
      .then((res) => setCourse(res.data?.course ?? null))
      .catch((err) =>
        setError(err instanceof Error ? err.message : 'Failed to load course'),
      )
      .finally(() => setIsLoading(false))
  }, [courseId])

  useEffect(() => {
    setIsLoading(true)
    loadCourse()
  }, [loadCourse])

  const isEducator = user?.role === 'educator'

  if (isLoading) return <LoadingSpinner className="min-h-[40vh]" />

  if (error || !course) {
    return (
      <Card>
        <p className="text-center text-red-500">{error || 'Course not found'}</p>
      </Card>
    )
  }

  return (
    <div className="animate-fade-in space-y-10">
      <Breadcrumbs
        items={[
          { label: 'Home', to: '/' },
          { label: course.title },
        ]}
      />

      <header className="space-y-3">
        {course.category ? (
          <span className="inline-block rounded-full bg-sage-100 px-3 py-1 text-xs font-medium text-sage-700 dark:bg-sage-900/50 dark:text-sage-300">
            {course.category}
          </span>
        ) : null}
        <h1 className="font-display text-3xl font-semibold text-ink-900 dark:text-ink-50">
          {course.title}
        </h1>
        {course.description ? (
          <p className="max-w-3xl text-lg leading-relaxed text-ink-600 dark:text-ink-300">
            {course.description}
          </p>
        ) : null}
      </header>

      {isEducator && courseId ? (
        <CreateCoursePlaylistForm courseId={courseId} onCreated={loadCourse} />
      ) : null}

      {course.educators.length > 0 ? (
        <section className="space-y-4">
          <h2 className="flex items-center gap-2 font-display text-xl font-semibold text-ink-900 dark:text-ink-50">
            <GraduationCap className="size-5 text-sage-500" />
            Educators
          </h2>
          <div className="grid gap-4 sm:grid-cols-2 lg:grid-cols-3">
            {course.educators.map((educator) => (
              <EducatorCard key={educator.user_id} educator={educator} />
            ))}
          </div>
        </section>
      ) : null}

      <section className="space-y-4">
        <h2 className="font-display text-xl font-semibold text-ink-900 dark:text-ink-50">
          Playlists & modules
        </h2>
        {course.playlists.length === 0 ? (
          <Card className="py-12 text-center text-ink-500">
            {isEducator
              ? 'No playlists yet. Create one above to start uploading lessons.'
              : 'No playlists in this course yet.'}
          </Card>
        ) : (
          <div className="grid gap-4 sm:grid-cols-2">
            {course.playlists.map((playlist) => (
              <PlaylistCard
                key={playlist.playlist_id}
                playlist={playlist}
                href={`/courses/${course.course_id}/playlists/${playlist.playlist_id}`}
              />
            ))}
          </div>
        )}
      </section>
    </div>
  )
}

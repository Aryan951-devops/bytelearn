import { useEffect, useState } from 'react'
import { Link } from 'react-router-dom'
import { BookOpen, Sparkles } from 'lucide-react'
import { courseApi } from '@/lib/api'
import type { Course } from '@/types'
import { CourseCard } from '@/components/course/CourseCard'
import { Card } from '@/components/ui/Card'
import { LoadingSpinner } from '@/components/ui/LoadingSpinner'
import { useAuth } from '@/context/AuthContext'
import { Button } from '@/components/ui/Button'

export function HomePage() {
  const { isAuthenticated, user } = useAuth()
  const [courses, setCourses] = useState<Course[]>([])
  const [error, setError] = useState('')
  const [isLoading, setIsLoading] = useState(true)

  useEffect(() => {
    courseApi
      .getAll()
      .then((res) => setCourses(res.data?.courses ?? []))
      .catch((err) =>
        setError(err instanceof Error ? err.message : 'Failed to load courses'),
      )
      .finally(() => setIsLoading(false))
  }, [])

  return (
    <div className="animate-fade-in space-y-10">
      <section className="relative overflow-hidden rounded-3xl">
        <div className="glass absolute inset-0 rounded-3xl" />
        <div className="relative px-6 py-10 sm:px-10 sm:py-14">
          <p className="mb-3 inline-flex items-center gap-2 rounded-full bg-sage-100/80 px-4 py-1.5 text-sm font-medium text-sage-700 dark:bg-sage-900/40 dark:text-sage-200">
            <Sparkles className="size-4" />
            Learn calmly. Build confidently.
          </p>
          <h1 className="font-display max-w-2xl text-3xl font-semibold leading-tight text-ink-900 sm:text-4xl dark:text-ink-50">
            Explore courses
          </h1>
          <p className="mt-3 max-w-xl text-ink-600 dark:text-ink-300">
            Pick a course to browse modules, meet your educators, and start learning.
          </p>
          {isAuthenticated && user?.role === 'admin' ? (
            <Link to="/admin/courses" className="mt-5 inline-block">
              <Button size="sm" variant="secondary">
                Manage courses
              </Button>
            </Link>
          ) : null}
        </div>
      </section>

      <section className="space-y-6">
        <div className="flex items-end justify-between gap-4">
          <div>
            <h2 className="font-display text-2xl font-semibold text-ink-900 dark:text-ink-50">
              All courses
            </h2>
            <p className="mt-1 text-sm text-ink-500 dark:text-ink-400">
              {courses.length} course{courses.length !== 1 ? 's' : ''} available
            </p>
          </div>
          {isAuthenticated ? (
            <Link to="/my-playlists">
              <Button variant="ghost" size="sm">
                My playlists
              </Button>
            </Link>
          ) : null}
        </div>

        {isLoading ? (
          <LoadingSpinner className="min-h-[30vh]" />
        ) : error ? (
          <Card>
            <p className="text-center text-red-500 dark:text-red-400">{error}</p>
          </Card>
        ) : courses.length === 0 ? (
          <Card className="py-16 text-center">
            <BookOpen className="mx-auto size-12 text-ink-300 dark:text-ink-600" />
            <p className="mt-4 font-medium text-ink-700 dark:text-ink-300">
              No courses yet
            </p>
            <p className="mt-1 text-sm text-ink-500">
              Courses will appear here once an admin creates them.
            </p>
          </Card>
        ) : (
          <div className="grid gap-5 sm:grid-cols-2 lg:grid-cols-3">
            {courses.map((course) => (
              <CourseCard key={course.course_id} course={course} />
            ))}
          </div>
        )}
      </section>
    </div>
  )
}

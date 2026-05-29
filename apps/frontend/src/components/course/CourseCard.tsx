import { Link } from 'react-router-dom'
import { BookOpen } from 'lucide-react'
import type { Course } from '@/types'
import { Card } from '@/components/ui/Card'

export function CourseCard({ course }: { course: Course }) {
  return (
    <Link to={`/courses/${course.course_id}`}>
      <Card hover className="group h-full">
        <div className="mb-4 inline-flex rounded-xl bg-gradient-to-br from-sage-400 to-mist-500 p-3 text-white shadow-md transition-transform duration-300 group-hover:scale-105">
          <BookOpen className="size-6" />
        </div>
        {course.category ? (
          <span className="mb-2 inline-block rounded-full bg-sage-100 px-2.5 py-0.5 text-xs font-medium text-sage-700 dark:bg-sage-900/50 dark:text-sage-300">
            {course.category}
          </span>
        ) : null}
        <h3 className="font-display text-lg font-semibold text-ink-900 dark:text-ink-50">
          {course.title}
        </h3>
        {course.description ? (
          <p className="mt-2 line-clamp-3 text-sm leading-relaxed text-ink-500 dark:text-ink-400">
            {course.description}
          </p>
        ) : (
          <p className="mt-2 text-sm text-ink-400">Explore modules and lessons</p>
        )}
      </Card>
    </Link>
  )
}

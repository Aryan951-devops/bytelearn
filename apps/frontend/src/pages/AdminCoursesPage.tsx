import { type FormEvent, useEffect, useState } from 'react'
import { Shield } from 'lucide-react'
import { courseApi } from '@/lib/api'
import type { Course } from '@/types'
import { Breadcrumbs } from '@/components/ui/Breadcrumbs'
import { Button } from '@/components/ui/Button'
import { Card } from '@/components/ui/Card'
import { Input } from '@/components/ui/Input'
import { LoadingSpinner } from '@/components/ui/LoadingSpinner'
import { CourseCard } from '@/components/course/CourseCard'

export function AdminCoursesPage() {
  const [courses, setCourses] = useState<Course[]>([])
  const [isLoading, setIsLoading] = useState(true)
  const [error, setError] = useState('')
  const [title, setTitle] = useState('')
  const [description, setDescription] = useState('')
  const [category, setCategory] = useState('')
  const [creating, setCreating] = useState(false)
  const [message, setMessage] = useState('')
  const [editingId, setEditingId] = useState<string | null>(null)
  const [editTitle, setEditTitle] = useState('')
  const [editDescription, setEditDescription] = useState('')
  const [editCategory, setEditCategory] = useState('')

  const loadCourses = () => {
    courseApi
      .getAll()
      .then((res) => setCourses(res.data?.courses ?? []))
      .catch((err) =>
        setError(err instanceof Error ? err.message : 'Failed to load courses'),
      )
      .finally(() => setIsLoading(false))
  }

  useEffect(() => {
    loadCourses()
  }, [])

  const handleCreate = async (e: FormEvent) => {
    e.preventDefault()
    setCreating(true)
    setMessage('')
    try {
      await courseApi.create({
        title: title.trim(),
        description: description.trim() || undefined,
        category: category.trim() || undefined,
      })
      setTitle('')
      setDescription('')
      setCategory('')
      setMessage('Course created successfully.')
      setIsLoading(true)
      loadCourses()
    } catch (err) {
      setMessage(err instanceof Error ? err.message : 'Failed to create course')
    } finally {
      setCreating(false)
    }
  }

  const startEdit = (course: Course) => {
    setEditingId(course.course_id)
    setEditTitle(course.title)
    setEditDescription(course.description ?? '')
    setEditCategory(course.category ?? '')
  }

  const handleUpdate = async (courseId: string) => {
    setMessage('')
    try {
      await courseApi.update(courseId, {
        title: editTitle.trim() || undefined,
        description: editDescription.trim() || undefined,
        category: editCategory.trim() || undefined,
      })
      setEditingId(null)
      setMessage('Course updated.')
      loadCourses()
    } catch (err) {
      setMessage(err instanceof Error ? err.message : 'Failed to update course')
    }
  }

  return (
    <div className="animate-fade-in space-y-8">
      <Breadcrumbs
        items={[
          { label: 'Home', to: '/' },
          { label: 'Admin — Courses' },
        ]}
      />

      <div className="flex items-center gap-3">
        <div className="flex size-11 items-center justify-center rounded-xl bg-sage-100 text-sage-600 dark:bg-sage-900/50 dark:text-sage-300">
          <Shield className="size-6" />
        </div>
        <div>
          <h1 className="font-display text-3xl font-semibold text-ink-900 dark:text-ink-50">
            Manage courses
          </h1>
          <p className="text-sm text-ink-500 dark:text-ink-400">
            Only admins can create and update courses.
          </p>
        </div>
      </div>

      <Card>
        <h2 className="mb-4 font-display text-lg font-semibold text-ink-900 dark:text-ink-50">
          Create course
        </h2>
        <form onSubmit={handleCreate} className="space-y-4">
          <Input label="Title" required value={title} onChange={(e) => setTitle(e.target.value)} />
          <Input
            label="Description"
            value={description}
            onChange={(e) => setDescription(e.target.value)}
          />
          <Input
            label="Category"
            value={category}
            onChange={(e) => setCategory(e.target.value)}
          />
          <Button type="submit" isLoading={creating}>
            Create course
          </Button>
        </form>
        {message ? (
          <p className="mt-3 text-sm text-sage-600 dark:text-sage-400">{message}</p>
        ) : null}
      </Card>

      <section className="space-y-4">
        <h2 className="font-display text-xl font-semibold text-ink-900 dark:text-ink-50">
          Existing courses
        </h2>
        {isLoading ? (
          <LoadingSpinner className="min-h-[20vh]" />
        ) : error ? (
          <Card>
            <p className="text-center text-red-500">{error}</p>
          </Card>
        ) : (
          <div className="space-y-6">
            {courses.map((course) =>
              editingId === course.course_id ? (
                <Card key={course.course_id} className="space-y-3">
                  <Input
                    label="Title"
                    value={editTitle}
                    onChange={(e) => setEditTitle(e.target.value)}
                  />
                  <Input
                    label="Description"
                    value={editDescription}
                    onChange={(e) => setEditDescription(e.target.value)}
                  />
                  <Input
                    label="Category"
                    value={editCategory}
                    onChange={(e) => setEditCategory(e.target.value)}
                  />
                  <div className="flex gap-2">
                    <Button onClick={() => handleUpdate(course.course_id)}>Save</Button>
                    <Button variant="ghost" onClick={() => setEditingId(null)}>
                      Cancel
                    </Button>
                  </div>
                </Card>
              ) : (
                <div key={course.course_id} className="flex flex-col gap-3 sm:flex-row sm:items-start">
                  <div className="flex-1">
                    <CourseCard course={course} />
                  </div>
                  <Button variant="secondary" size="sm" onClick={() => startEdit(course)}>
                    Edit
                  </Button>
                </div>
              ),
            )}
          </div>
        )}
      </section>
    </div>
  )
}

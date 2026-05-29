import { type FormEvent, useState } from 'react'
import { ListPlus } from 'lucide-react'
import { playlistApi } from '@/lib/api'
import { Button } from '@/components/ui/Button'
import { Card } from '@/components/ui/Card'
import { Input } from '@/components/ui/Input'

interface CreateCoursePlaylistFormProps {
  courseId: string
  onCreated?: () => void
}

export function CreateCoursePlaylistForm({
  courseId,
  onCreated,
}: CreateCoursePlaylistFormProps) {
  const [title, setTitle] = useState('')
  const [description, setDescription] = useState('')
  const [error, setError] = useState('')
  const [message, setMessage] = useState('')
  const [isLoading, setIsLoading] = useState(false)

  const handleSubmit = async (e: FormEvent) => {
    e.preventDefault()
    setError('')
    setMessage('')
    setIsLoading(true)
    try {
      await playlistApi.create({
        type: 'course',
        title: title.trim(),
        description: description.trim() || undefined,
        course_id: courseId,
      })
      setTitle('')
      setDescription('')
      setMessage('Playlist created! You can upload videos from the playlist page.')
      onCreated?.()
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Failed to create playlist')
    } finally {
      setIsLoading(false)
    }
  }

  return (
    <Card className="border-sage-200/60 dark:border-sage-800/60">
      <h3 className="mb-4 flex items-center gap-2 font-display text-lg font-semibold text-ink-900 dark:text-ink-50">
        <ListPlus className="size-5 text-sage-500" />
        Create course playlist
      </h3>
      <form onSubmit={handleSubmit} className="space-y-3">
        <Input
          label="Playlist title"
          required
          value={title}
          onChange={(e) => setTitle(e.target.value)}
        />
        <div className="space-y-1.5">
          <label className="block text-sm font-medium text-ink-700 dark:text-ink-200">
            Description (optional)
          </label>
          <textarea
            className="w-full rounded-xl border border-ink-200/80 bg-white/80 px-4 py-2.5 text-sm text-ink-900 focus:border-sage-400 focus:outline-none focus:ring-2 focus:ring-sage-300/40 dark:border-ink-700 dark:bg-ink-900/50 dark:text-ink-50"
            rows={2}
            value={description}
            onChange={(e) => setDescription(e.target.value)}
          />
        </div>
        {error ? <p className="text-sm text-red-500">{error}</p> : null}
        {message ? (
          <p className="text-sm text-sage-600 dark:text-sage-400">{message}</p>
        ) : null}
        <Button type="submit" isLoading={isLoading}>
          Create playlist
        </Button>
      </form>
    </Card>
  )
}

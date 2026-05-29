import { type FormEvent, useEffect, useState } from 'react'
import { ListPlus, ListVideo } from 'lucide-react'
import { playlistApi } from '@/lib/api'
import type { Playlist } from '@/types'
import { Breadcrumbs } from '@/components/ui/Breadcrumbs'
import { Button } from '@/components/ui/Button'
import { Card } from '@/components/ui/Card'
import { Input } from '@/components/ui/Input'
import { LoadingSpinner } from '@/components/ui/LoadingSpinner'
import { PlaylistCard } from '@/components/playlist/PlaylistCard'

export function MyPlaylistsPage() {
  const [playlists, setPlaylists] = useState<Playlist[]>([])
  const [error, setError] = useState('')
  const [isLoading, setIsLoading] = useState(true)
  const [showForm, setShowForm] = useState(false)
  const [title, setTitle] = useState('')
  const [description, setDescription] = useState('')
  const [creating, setCreating] = useState(false)
  const [formError, setFormError] = useState('')

  const loadPlaylists = () => {
    playlistApi
      .getUserPlaylists()
      .then((res) => {
        const all = res.data?.playlist ?? []
        setPlaylists(all.filter((p) => p.type === 'user'))
      })
      .catch((err) =>
        setError(err instanceof Error ? err.message : 'Failed to load playlists'),
      )
      .finally(() => setIsLoading(false))
  }

  useEffect(() => {
    loadPlaylists()
  }, [])

  const handleCreate = async (e: FormEvent) => {
    e.preventDefault()
    setFormError('')
    setCreating(true)
    try {
      await playlistApi.create({
        type: 'user',
        title: title.trim(),
        description: description.trim() || undefined,
      })
      setTitle('')
      setDescription('')
      setShowForm(false)
      setIsLoading(true)
      loadPlaylists()
    } catch (err) {
      setFormError(err instanceof Error ? err.message : 'Failed to create playlist')
    } finally {
      setCreating(false)
    }
  }

  const personalPlaylists = playlists

  return (
    <div className="animate-fade-in space-y-8">
      <Breadcrumbs items={[{ label: 'Home', to: '/' }, { label: 'My playlists' }]} />

      <div className="flex flex-wrap items-end justify-between gap-4">
        <div>
          <h1 className="font-display text-3xl font-semibold text-ink-900 dark:text-ink-50">
            My playlists
          </h1>
          <p className="mt-2 text-ink-500 dark:text-ink-400">
            Create personal collections and save videos while you learn.
          </p>
        </div>
        <Button onClick={() => setShowForm((v) => !v)}>
          <ListPlus className="size-4" />
          New playlist
        </Button>
      </div>

      {showForm ? (
        <Card>
          <form onSubmit={handleCreate} className="space-y-4">
            <Input
              label="Title"
              required
              value={title}
              onChange={(e) => setTitle(e.target.value)}
            />
            <Input
              label="Description (optional)"
              value={description}
              onChange={(e) => setDescription(e.target.value)}
            />
            {formError ? (
              <p className="text-sm text-red-500">{formError}</p>
            ) : null}
            <div className="flex gap-2">
              <Button type="submit" isLoading={creating}>
                Create
              </Button>
              <Button
                type="button"
                variant="ghost"
                onClick={() => setShowForm(false)}
              >
                Cancel
              </Button>
            </div>
          </form>
        </Card>
      ) : null}

      {isLoading ? (
        <LoadingSpinner className="min-h-[30vh]" />
      ) : error ? (
        <Card>
          <p className="text-center text-red-500">{error}</p>
        </Card>
      ) : personalPlaylists.length === 0 ? (
        <Card className="py-16 text-center">
          <ListVideo className="mx-auto size-12 text-ink-300 dark:text-ink-600" />
          <p className="mt-4 font-medium text-ink-700 dark:text-ink-300">
            No playlists yet
          </p>
          <p className="mt-1 text-sm text-ink-500">
            Create your first playlist to organize saved videos.
          </p>
        </Card>
      ) : (
        <div className="grid gap-4 sm:grid-cols-2">
          {personalPlaylists.map((playlist) => (
            <PlaylistCard
              key={playlist.playlist_id}
              playlist={playlist}
              href={`/my-playlists/${playlist.playlist_id}`}
              subtitle={`Created ${new Date(playlist.created_at).toLocaleDateString()}`}
            />
          ))}
        </div>
      )}
    </div>
  )
}

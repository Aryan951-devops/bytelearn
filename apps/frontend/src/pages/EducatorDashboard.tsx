import { type FormEvent, useEffect, useState } from 'react'
import { Link } from 'react-router-dom'
import { playlistApi, videoApi } from '@/lib/api'
import {
  registerVideoWithBackend,
  uploadVideoToCloudinary,
} from '@/lib/cloudinaryUpload'
import type { Playlist, Video } from '@/types'
import { Card } from '@/components/ui/Card'
import { Button } from '@/components/ui/Button'
import { Input } from '@/components/ui/Input'
import { FileInput } from '@/components/ui/FileInput'
import { LoadingSpinner } from '@/components/ui/LoadingSpinner'
import {
  Trash2,
  Edit3,
  UploadCloud,
  Film,
  ListVideo,
  ExternalLink,
} from 'lucide-react'

export function EducatorDashboard() {
  const [videos, setVideos] = useState<Video[]>([])
  const [coursePlaylists, setCoursePlaylists] = useState<Playlist[]>([])

  const [title, setTitle] = useState('')
  const [description, setDescription] = useState('')
  const [videoFile, setVideoFile] = useState<File | null>(null)

  const [editingVideo, setEditingVideo] = useState<Video | null>(null)
  const [editTitle, setEditTitle] = useState('')
  const [editDescription, setEditDescription] = useState('')
  const [thumbnailFile, setThumbnailFile] = useState<File | null>(null)

  const [isLoading, setIsLoading] = useState(false)
  const [loadingPlaylists, setLoadingPlaylists] = useState(true)
  const [message, setMessage] = useState('')

  useEffect(() => {
    fetchVideos()
    playlistApi
      .getUserPlaylists()
      .then((res) => {
        const all = res.data?.playlist ?? []
        setCoursePlaylists(all.filter((p) => p.type === 'course'))
      })
      .catch(console.error)
      .finally(() => setLoadingPlaylists(false))
  }, [])

  const fetchVideos = () => {
    videoApi
      .getMyVideos()
      .then((res) => setVideos(res.data?.videos ?? []))
      .catch(console.error)
  }

  const handleCreateVideo = async (e: FormEvent) => {
    e.preventDefault()
    if (!title || !videoFile) return

    setIsLoading(true)
    setMessage('Uploading file to cloud storage…')
    try {
      const uploadData = await uploadVideoToCloudinary(videoFile)
      setMessage('Registering video…')
      await registerVideoWithBackend(uploadData, {
        title,
        description,
      })

      setMessage('Video uploaded! Use “Add to playlist” below to add it to a course playlist.')
      setTitle('')
      setDescription('')
      setVideoFile(null)
      fetchVideos()
    } catch (err) {
      setMessage(err instanceof Error ? err.message : 'Upload failed')
    } finally {
      setIsLoading(false)
    }
  }

  const handleUpdateVideo = async (e: FormEvent) => {
    e.preventDefault()
    if (!editingVideo) return

    setIsLoading(true)
    try {
      const formData = new FormData()
      formData.append('title', editTitle)
      formData.append('description', editDescription)
      if (thumbnailFile) formData.append('thumbnail', thumbnailFile)

      await videoApi.update(editingVideo.video_id, formData)
      setMessage('Video updated.')
      setEditingVideo(null)
      setThumbnailFile(null)
      fetchVideos()
    } catch (err) {
      setMessage(err instanceof Error ? err.message : 'Update failed')
    } finally {
      setIsLoading(false)
    }
  }

  const handleDeleteVideo = async (videoId: string) => {
    if (!confirm('Delete this video?')) return
    try {
      await videoApi.delete(videoId)
      setVideos((v) => v.filter((x) => x.video_id !== videoId))
    } catch (err) {
      alert(err instanceof Error ? err.message : 'Delete failed')
    }
  }

  const addExistingToPlaylist = async (videoId: string, playlistId: string) => {
    try {
      await playlistApi.addVideo(videoId, playlistId)
      setMessage('Video added to playlist.')
    } catch (err) {
      setMessage(err instanceof Error ? err.message : 'Could not add to playlist')
    }
  }

  return (
    <div className="mx-auto max-w-6xl space-y-10 p-4">
      <div>
        <h1 className="font-display text-3xl font-semibold text-ink-900 dark:text-ink-50">
          Educator dashboard
        </h1>
        <p className="mt-2 text-ink-500 dark:text-ink-400">
          Upload videos first, then add them to your course playlists.
        </p>
      </div>

      <div className="grid gap-8 lg:grid-cols-3">
        <div className="space-y-6 lg:col-span-1">
          {!editingVideo ? (
            <Card className="space-y-4 p-6">
              <h2 className="flex items-center gap-2 text-xl font-semibold">
                <UploadCloud className="size-5 text-sage-500" />
                Publish lesson
              </h2>
              <form onSubmit={handleCreateVideo} className="space-y-3">
                <Input
                  label="Video title *"
                  value={title}
                  onChange={(e) => setTitle(e.target.value)}
                  required
                />
                <div className="space-y-1.5">
                  <label className="text-sm font-medium text-ink-700 dark:text-ink-200">
                    Description
                  </label>
                  <textarea
                    className="w-full rounded-xl border border-ink-200/80 bg-white/80 px-3 py-2 text-sm text-ink-900 dark:border-ink-700 dark:bg-ink-900/50 dark:text-ink-50"
                    rows={3}
                    value={description}
                    onChange={(e) => setDescription(e.target.value)}
                  />
                </div>
                <FileInput
                  label="Video file *"
                  accept="video/*"
                  required
                  onChange={(e) => setVideoFile(e.target.files?.[0] ?? null)}
                />

                <Button type="submit" disabled={isLoading} className="w-full">
                  {isLoading ? 'Processing…' : 'Upload video'}
                </Button>
              </form>
            </Card>
          ) : (
            <Card className="space-y-4 border-amber-500/50 p-6">
              <h2 className="flex items-center gap-2 text-xl font-semibold text-amber-600">
                <Edit3 className="size-5" />
                Edit video
              </h2>
              <form onSubmit={handleUpdateVideo} className="space-y-3">
                <Input
                  label="Title"
                  value={editTitle}
                  onChange={(e) => setEditTitle(e.target.value)}
                />
                <textarea
                  className="w-full rounded-xl border border-ink-200/80 p-2 text-sm text-ink-900 dark:border-ink-700 dark:bg-ink-900/50 dark:text-ink-50"
                  rows={3}
                  value={editDescription}
                  onChange={(e) => setEditDescription(e.target.value)}
                />
                <FileInput
                  label="Thumbnail image"
                  accept="image/*"
                  onChange={(e) => setThumbnailFile(e.target.files?.[0] ?? null)}
                />
                <div className="flex gap-2">
                  <Button type="submit" disabled={isLoading} className="flex-1">
                    Save
                  </Button>
                  <Button
                    type="button"
                    variant="secondary"
                    onClick={() => setEditingVideo(null)}
                  >
                    Cancel
                  </Button>
                </div>
              </form>
            </Card>
          )}
          {message ? (
            <p className="text-sm font-medium text-sage-600 dark:text-sage-400">
              {message}
            </p>
          ) : null}
        </div>

        <div className="space-y-8 lg:col-span-2">
          <section className="space-y-4">
            <h2 className="flex items-center gap-2 text-xl font-semibold">
              <ListVideo className="size-5 text-mist-500" />
              Your course playlists
            </h2>
            {loadingPlaylists ? (
              <LoadingSpinner />
            ) : coursePlaylists.length === 0 ? (
              <Card className="py-8 text-center text-sm text-ink-500">
                Open a course and create a playlist, then add uploaded videos below.
              </Card>
            ) : (
              <div className="grid gap-3 sm:grid-cols-2">
                {coursePlaylists.map((p) => (
                  <Card key={p.playlist_id} className="flex flex-col gap-2 p-4">
                    <p className="font-medium text-ink-900 dark:text-ink-50">
                      {p.title}
                    </p>
                    {p.course_id ? (
                      <Link
                        to={`/courses/${p.course_id}/playlists/${p.playlist_id}`}
                        className="inline-flex items-center gap-1 text-sm text-sage-600 hover:underline dark:text-sage-400"
                      >
                        View playlist
                        <ExternalLink className="size-3.5" />
                      </Link>
                    ) : null}
                  </Card>
                ))}
              </div>
            )}
          </section>

          <section className="space-y-4">
            <h2 className="flex items-center gap-2 text-xl font-semibold">
              <Film className="size-5" />
              Your videos
            </h2>
            {videos.length === 0 ? (
              <p className="text-sm text-ink-400">No videos uploaded yet.</p>
            ) : (
              <div className="space-y-3">
                {videos.map((video) => (
                  <div
                    key={video.video_id}
                    className="flex flex-col gap-3 rounded-xl border border-ink-200/60 bg-white p-4 dark:border-ink-800 dark:bg-ink-900 sm:flex-row sm:items-center sm:justify-between"
                  >
                    <div className="flex items-center gap-4">
                      {video.thumbnail_url ? (
                        <img
                          src={video.thumbnail_url}
                          alt=""
                          className="aspect-video w-20 rounded-lg object-cover"
                        />
                      ) : (
                        <div className="flex aspect-video w-20 items-center justify-center rounded-lg bg-ink-100 dark:bg-ink-800">
                          <Film className="size-5 text-ink-400" />
                        </div>
                      )}
                      <div>
                        <h4 className="line-clamp-1 text-sm font-semibold">
                          {video.title}
                        </h4>
                        <Link
                          to={`/watch/${video.video_id}`}
                          className="text-xs text-sage-600 hover:underline"
                        >
                          View
                        </Link>
                      </div>
                    </div>
                    <div className="flex flex-wrap items-center gap-2">
                      {coursePlaylists.length > 0 ? (
                        <select
                          className="rounded-lg border border-ink-200 bg-white px-2 py-1.5 text-xs text-ink-800 dark:border-ink-600 dark:bg-ink-800 dark:text-ink-100"
                          defaultValue=""
                          onChange={(e) => {
                            if (e.target.value) {
                              addExistingToPlaylist(video.video_id, e.target.value)
                              e.target.value = ''
                            }
                          }}
                        >
                          <option value="">Add to playlist…</option>
                          {coursePlaylists.map((p) => (
                            <option key={p.playlist_id} value={p.playlist_id}>
                              {p.title}
                            </option>
                          ))}
                        </select>
                      ) : null}
                      <button
                        type="button"
                        onClick={() => {
                          setEditingVideo(video)
                          setEditTitle(video.title)
                          setEditDescription(video.description || '')
                        }}
                        className="rounded-lg p-2 hover:bg-ink-100 dark:hover:bg-ink-800"
                      >
                        <Edit3 className="size-4" />
                      </button>
                      <button
                        type="button"
                        onClick={() => handleDeleteVideo(video.video_id)}
                        className="rounded-lg p-2 text-red-500 hover:bg-red-50 dark:hover:bg-red-950/30"
                      >
                        <Trash2 className="size-4" />
                      </button>
                    </div>
                  </div>
                ))}
              </div>
            )}
          </section>
        </div>
      </div>
    </div>
  )
}

import { Link } from 'react-router-dom'
import { Clock, Eye, Play, Video as VideoIcon } from 'lucide-react'
import type { Video, VideoMetadata } from '@/types'
import { Card } from '@/components/ui/Card'

export function formatDuration(seconds: number | string): string {
  const total = typeof seconds === 'string' ? parseInt(seconds, 10) : seconds
  if (!total || Number.isNaN(total)) return '—'
  const m = Math.floor(total / 60)
  const s = total % 60
  return `${m}:${s.toString().padStart(2, '0')}`
}

interface VideoCardProps {
  video: Video | VideoMetadata
  durationSeconds?: number | string
  linkTo?: string
  onClick?: () => void
}

export function VideoCard({ video, durationSeconds, linkTo, onClick }: VideoCardProps) {
  const thumbnail = 'thumbnail_url' in video ? video.thumbnail_url : null
  const views = video.views ?? 0

  const content = (
    <Card hover className="group overflow-hidden p-0">
      <div className="relative aspect-video overflow-hidden bg-ink-100 dark:bg-ink-800">
        {thumbnail ? (
          <img
            src={thumbnail}
            alt=""
            className="size-full object-cover transition-transform duration-500 group-hover:scale-105"
          />
        ) : (
          <div className="flex size-full items-center justify-center bg-gradient-to-br from-sage-200 to-mist-200 dark:from-sage-900 dark:to-mist-900">
            <VideoIcon className="size-12 text-sage-500/50" />
          </div>
        )}
        <div className="absolute inset-0 flex items-center justify-center bg-ink-900/0 opacity-0 transition-all group-hover:bg-ink-900/25 group-hover:opacity-100">
          <span className="flex size-12 items-center justify-center rounded-full bg-white/90 text-sage-600 shadow-lg">
            <Play className="size-5 fill-current" />
          </span>
        </div>
      </div>
      <div className="p-4">
        <h3 className="line-clamp-2 font-medium text-ink-900 dark:text-ink-50">
          {video.title}
        </h3>
        <div className="mt-2 flex flex-wrap gap-3 text-xs text-ink-400 dark:text-ink-500">
          <span className="inline-flex items-center gap-1">
            <Eye className="size-3.5" />
            {views.toLocaleString()} views
          </span>
          {durationSeconds !== undefined ? (
            <span className="inline-flex items-center gap-1">
              <Clock className="size-3.5" />
              {formatDuration(durationSeconds)}
            </span>
          ) : null}
        </div>
      </div>
    </Card>
  )

  if (linkTo) {
    return (
      <Link to={linkTo} className="block">
        {content}
      </Link>
    )
  }

  return (
    <button type="button" onClick={onClick} className="block w-full text-left">
      {content}
    </button>
  )
}

import { Link } from 'react-router-dom'
import { ListVideo } from 'lucide-react'
import type { Playlist, PlaylistPreview } from '@/types'
import { Card } from '@/components/ui/Card'

interface PlaylistCardProps {
  playlist: PlaylistPreview | Playlist
  href: string
  subtitle?: string
}

export function PlaylistCard({ playlist, href, subtitle }: PlaylistCardProps) {
  return (
    <Link to={href}>
      <Card hover className="group flex h-full items-start gap-4">
        <div className="flex size-12 shrink-0 items-center justify-center rounded-xl bg-mist-100 text-mist-600 transition-transform group-hover:scale-105 dark:bg-mist-900/40 dark:text-mist-300">
          <ListVideo className="size-6" />
        </div>
        <div className="min-w-0 flex-1">
          <h3 className="font-medium text-ink-900 dark:text-ink-50">{playlist.title}</h3>
          {playlist.description ? (
            <p className="mt-1 line-clamp-2 text-sm text-ink-500 dark:text-ink-400">
              {playlist.description}
            </p>
          ) : null}
          {subtitle ? (
            <p className="mt-2 text-xs text-ink-400">{subtitle}</p>
          ) : null}
        </div>
      </Card>
    </Link>
  )
}

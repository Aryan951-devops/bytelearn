interface VideoPlayerProps {
  src: string
  title?: string
  poster?: string | null
  className?: string
}

export function VideoPlayer({ src, title, poster, className = '' }: VideoPlayerProps) {
  return (
    <div
      className={`overflow-hidden rounded-2xl bg-black shadow-xl ring-1 ring-ink-200/50 dark:ring-ink-700 ${className}`}
    >
      <video
        key={src}
        src={src}
        controls
        playsInline
        poster={poster ?? undefined}
        className="aspect-video w-full bg-black object-contain"
        title={title}
      >
        <track kind="captions" />
        Your browser does not support video playback.
      </video>
    </div>
  )
}

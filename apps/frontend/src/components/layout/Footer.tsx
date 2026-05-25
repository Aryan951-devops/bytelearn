import { GraduationCap } from 'lucide-react'
import { Link } from 'react-router-dom'

export function Footer() {
  return (
    <footer className="mt-auto border-t border-ink-200/50 bg-cream-100/50 dark:border-ink-800 dark:bg-ink-900/30">
      <div className="mx-auto flex max-w-6xl flex-col items-center justify-between gap-4 px-4 py-8 sm:flex-row sm:px-6">
        <div className="flex items-center gap-2 text-sm text-ink-500 dark:text-ink-400">
          <GraduationCap className="size-4 text-sage-500" />
          <span>ByteLearn — Smart Learning & Coding</span>
        </div>
        <div className="flex gap-6 text-sm">
          <Link
            to="/videos"
            className="text-ink-500 transition-colors hover:text-sage-600 dark:text-ink-400 dark:hover:text-sage-300"
          >
            Browse videos
          </Link>
        </div>
      </div>
    </footer>
  )
}

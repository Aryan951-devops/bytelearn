import { type FormEvent, useEffect, useState } from 'react'
import { useNavigate, useSearchParams } from 'react-router-dom'
import { Search } from 'lucide-react'

interface NavbarSearchProps {
  onSearch?: () => void
}

export function NavbarSearch({ onSearch }: NavbarSearchProps) {
  const navigate = useNavigate()
  const [searchParams] = useSearchParams()
  const urlQuery = searchParams.get('q') ?? ''
  const [query, setQuery] = useState(urlQuery)

  useEffect(() => {
    setQuery(urlQuery)
  }, [urlQuery])

  const handleSubmit = (e: FormEvent) => {
    e.preventDefault()
    const trimmed = query.trim()
    if (!trimmed) return
    navigate(`/search?q=${encodeURIComponent(trimmed)}`)
    onSearch?.()
  }

  return (
    <form onSubmit={handleSubmit} className="relative w-full">
      <Search className="pointer-events-none absolute left-3 top-1/2 size-4 -translate-y-1/2 text-ink-400" />
      <input
        type="search"
        value={query}
        onChange={(e) => setQuery(e.target.value)}
        placeholder="Search videos..."
        aria-label="Search videos"
        className="h-9 w-full rounded-xl border border-ink-200/80 bg-white/80 py-2 pl-9 pr-3 text-sm text-ink-900 shadow-sm transition-colors placeholder:text-ink-400 focus:border-sage-400 focus:outline-none focus:ring-2 focus:ring-sage-300/40 dark:border-ink-700 dark:bg-ink-900/50 dark:text-ink-50 dark:placeholder:text-ink-500 dark:focus:border-sage-500 dark:focus:ring-sage-600/30"
      />
    </form>
  )
}

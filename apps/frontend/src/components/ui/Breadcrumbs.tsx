import { Link } from 'react-router-dom'
import { ChevronRight } from 'lucide-react'

export interface BreadcrumbItem {
  label: string
  to?: string
}

export function Breadcrumbs({ items }: { items: BreadcrumbItem[] }) {
  return (
    <nav aria-label="Breadcrumb" className="mb-6 flex flex-wrap items-center gap-1 text-sm">
      {items.map((item, i) => (
        <span key={i} className="flex items-center gap-1">
          {i > 0 ? (
            <ChevronRight className="size-3.5 shrink-0 text-ink-300 dark:text-ink-600" />
          ) : null}
          {item.to ? (
            <Link
              to={item.to}
              className="text-ink-500 transition-colors hover:text-sage-600 dark:text-ink-400 dark:hover:text-sage-300"
            >
              {item.label}
            </Link>
          ) : (
            <span className="font-medium text-ink-800 dark:text-ink-100">
              {item.label}
            </span>
          )}
        </span>
      ))}
    </nav>
  )
}

import { User } from 'lucide-react'
import type { EducatorDetail } from '@/types'
import { Card } from '@/components/ui/Card'

export function EducatorCard({ educator }: { educator: EducatorDetail }) {
  return (
    <Card className="flex items-center gap-4">
      {educator.profile_pic_url ? (
        <img
          src={educator.profile_pic_url}
          alt=""
          className="size-14 shrink-0 rounded-xl object-cover ring-2 ring-sage-100 dark:ring-sage-900"
        />
      ) : (
        <div className="flex size-14 shrink-0 items-center justify-center rounded-xl bg-gradient-to-br from-sage-400 to-mist-400 text-white">
          <User className="size-7" />
        </div>
      )}
      <div className="min-w-0">
        <p className="font-medium text-ink-900 dark:text-ink-50">{educator.name}</p>
        <p className="text-sm text-sage-600 dark:text-sage-400">@{educator.username}</p>
        <p className="mt-0.5 text-xs text-ink-400">Educator</p>
      </div>
    </Card>
  )
}

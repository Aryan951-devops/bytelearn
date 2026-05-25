import { Calendar, Mail, MapPin, User } from 'lucide-react'
import { Link } from 'react-router-dom'
import { useAuth } from '@/context/AuthContext'
import { Button } from '@/components/ui/Button'
import { Card } from '@/components/ui/Card'

function Field({
  label,
  value,
}: {
  label: string
  value: string | null | undefined
}) {
  return (
    <div>
      <dt className="text-xs font-medium uppercase tracking-wide text-ink-400">
        {label}
      </dt>
      <dd className="mt-1 text-ink-800 dark:text-ink-100">
        {value || '—'}
      </dd>
    </div>
  )
}

export function ProfilePage() {
  const { user } = useAuth()

  if (!user) return null

  const joined = new Date(user.created_at).toLocaleDateString(undefined, {
    year: 'numeric',
    month: 'long',
    day: 'numeric',
  })

  return (
    <div className="animate-slide-up mx-auto max-w-2xl space-y-6">
      <Card>
        <div className="flex flex-col items-center gap-6 sm:flex-row sm:items-start">
          <div className="relative">
            {user.profile_pic ? (
              <img
                src={user.profile_pic}
                alt=""
                className="size-24 rounded-2xl object-cover ring-4 ring-sage-100 dark:ring-sage-900"
              />
            ) : (
              <div className="flex size-24 items-center justify-center rounded-2xl bg-gradient-to-br from-sage-400 to-mist-400 text-white ring-4 ring-sage-100 dark:ring-sage-900">
                <User className="size-10" />
              </div>
            )}
          </div>
          <div className="flex-1 text-center sm:text-left">
            <h1 className="font-display text-2xl font-semibold text-ink-900 dark:text-ink-50">
              {user.name}
            </h1>
            <p className="text-sage-600 dark:text-sage-400">@{user.username}</p>
            <div className="mt-3 flex flex-wrap justify-center gap-4 text-sm text-ink-500 sm:justify-start dark:text-ink-400">
              <span className="inline-flex items-center gap-1.5">
                <Mail className="size-4" />
                {user.email}
              </span>
              <span className="inline-flex items-center gap-1.5">
                <Calendar className="size-4" />
                Joined {joined}
              </span>
            </div>
            <Link to="/settings" className="mt-4 inline-block">
              <Button variant="secondary" size="sm">
                Edit profile
              </Button>
            </Link>
          </div>
        </div>
      </Card>

      <Card>
        <h2 className="mb-4 flex items-center gap-2 font-display text-lg font-semibold text-ink-900 dark:text-ink-50">
          <MapPin className="size-5 text-sage-500" />
          Account details
        </h2>
        <dl className="grid gap-4 sm:grid-cols-2">
          <Field label="Phone" value={user.phone_no} />
          <Field label="City" value={user.city} />
          <Field label="State" value={user.state} />
          <Field label="Pincode" value={user.pincode} />
        </dl>
      </Card>
    </div>
  )
}

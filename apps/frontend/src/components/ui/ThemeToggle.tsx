import { Monitor, Moon, Sun } from 'lucide-react'
import { useTheme } from '@/context/ThemeContext'

const options = [
  { value: 'light' as const, icon: Sun, label: 'Light' },
  { value: 'dark' as const, icon: Moon, label: 'Dark' },
  { value: 'system' as const, icon: Monitor, label: 'System' },
]

export function ThemeToggle() {
  const { theme, setTheme } = useTheme()

  return (
    <div
      className="glass flex items-center gap-0.5 rounded-xl p-1"
      role="group"
      aria-label="Theme"
    >
      {options.map(({ value, icon: Icon, label }) => (
        <button
          key={value}
          type="button"
          title={label}
          onClick={() => setTheme(value)}
          className={`rounded-lg p-2 transition-colors ${
            theme === value
              ? 'bg-sage-600 text-white shadow-sm dark:bg-sage-500'
              : 'text-ink-500 hover:bg-sage-100/60 dark:text-ink-400 dark:hover:bg-ink-800'
          }`}
        >
          <Icon className="size-4" aria-hidden />
        </button>
      ))}
    </div>
  )
}

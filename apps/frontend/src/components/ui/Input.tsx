import type { InputHTMLAttributes } from 'react'

interface InputProps extends InputHTMLAttributes<HTMLInputElement> {
  label: string
  error?: string
}

export function Input({ label, error, id, className = '', ...props }: InputProps) {
  const inputId = id ?? label.toLowerCase().replace(/\s+/g, '-')

  return (
    <div className="space-y-1.5">
      <label
        htmlFor={inputId}
        className="block text-sm font-medium text-ink-700 dark:text-ink-200"
      >
        {label}
      </label>
      <input
        id={inputId}
        className={`w-full rounded-xl border border-ink-200/80 bg-white/80 px-4 py-2.5 text-ink-900 shadow-sm transition-colors placeholder:text-ink-400 focus:border-sage-400 focus:outline-none focus:ring-2 focus:ring-sage-300/40 dark:border-ink-700 dark:bg-ink-900/50 dark:text-ink-50 dark:placeholder:text-ink-500 dark:focus:border-sage-500 dark:focus:ring-sage-600/30 ${error ? 'border-red-400 focus:border-red-400 focus:ring-red-300/40' : ''} ${className}`}
        {...props}
      />
      {error ? (
        <p className="text-sm text-red-500 dark:text-red-400">{error}</p>
      ) : null}
    </div>
  )
}

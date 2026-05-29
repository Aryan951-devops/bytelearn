import type { InputHTMLAttributes } from 'react'

interface FileInputProps extends Omit<InputHTMLAttributes<HTMLInputElement>, 'type'> {
  label: string
  hint?: string
}

export function FileInput({ label, hint, id, className = '', ...props }: FileInputProps) {
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
        type="file"
        className={`block w-full cursor-pointer rounded-xl border border-ink-200/80 bg-white px-3 py-2.5 text-sm text-ink-700 shadow-sm file:mr-3 file:cursor-pointer file:rounded-lg file:border file:border-sage-300 file:bg-sage-600 file:px-4 file:py-2 file:text-sm file:font-medium file:text-white hover:file:bg-sage-500 dark:border-ink-600 dark:bg-ink-900 dark:text-ink-200 dark:file:border-sage-600 dark:file:bg-sage-600 dark:file:text-white dark:hover:file:bg-sage-500 ${className}`}
        {...props}
      />
      {hint ? <p className="text-xs text-ink-400">{hint}</p> : null}
    </div>
  )
}

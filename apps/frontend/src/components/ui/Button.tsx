import type { ButtonHTMLAttributes, ReactNode } from 'react'

type Variant = 'primary' | 'secondary' | 'ghost' | 'danger'
type Size = 'sm' | 'md' | 'lg'

interface ButtonProps extends ButtonHTMLAttributes<HTMLButtonElement> {
  variant?: Variant
  size?: Size
  isLoading?: boolean
  children: ReactNode
}

const variants: Record<Variant, string> = {
  primary:
    'bg-sage-600 text-white shadow-md shadow-sage-600/20 hover:bg-sage-500 focus-visible:ring-sage-400 dark:bg-sage-500 dark:hover:bg-sage-400',
  secondary:
    'glass text-ink-800 hover:bg-white/90 dark:text-ink-100 dark:hover:bg-ink-800/80',
  ghost:
    'text-ink-600 hover:bg-sage-100/80 dark:text-ink-300 dark:hover:bg-ink-800',
  danger:
    'bg-red-500/90 text-white hover:bg-red-500 focus-visible:ring-red-400',
}

const sizes: Record<Size, string> = {
  sm: 'h-9 px-3.5 text-sm gap-1.5',
  md: 'h-11 px-5 text-sm gap-2',
  lg: 'h-12 px-6 text-base gap-2',
}

export function Button({
  variant = 'primary',
  size = 'md',
  isLoading,
  className = '',
  disabled,
  children,
  ...props
}: ButtonProps) {
  return (
    <button
      type="button"
      disabled={disabled || isLoading}
      className={`inline-flex items-center justify-center rounded-xl font-medium transition-all duration-200 focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-offset-2 focus-visible:ring-offset-cream-50 disabled:cursor-not-allowed disabled:opacity-50 dark:focus-visible:ring-offset-ink-950 ${variants[variant]} ${sizes[size]} ${className}`}
      {...props}
    >
      {isLoading ? (
        <span className="size-4 animate-spin rounded-full border-2 border-current border-t-transparent" />
      ) : null}
      {children}
    </button>
  )
}

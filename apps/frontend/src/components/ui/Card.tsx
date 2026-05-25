import type { ReactNode } from 'react'

interface CardProps {
  children: ReactNode
  className?: string
  hover?: boolean
}

export function Card({ children, className = '', hover }: CardProps) {
  return (
    <div
      className={`glass rounded-2xl p-6 shadow-sm ${hover ? 'transition-transform duration-300 hover:-translate-y-0.5 hover:shadow-md' : ''} ${className}`}
    >
      {children}
    </div>
  )
}

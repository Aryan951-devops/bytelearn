export function LoadingSpinner({ className = '' }: { className?: string }) {
  return (
    <div className={`flex items-center justify-center ${className}`}>
      <span className="size-8 animate-spin rounded-full border-2 border-sage-500 border-t-transparent" />
    </div>
  )
}

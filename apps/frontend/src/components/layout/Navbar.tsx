import { Link, NavLink, useNavigate } from 'react-router-dom'
import {
  BookOpen,
  GraduationCap,
  LogOut,
  Menu,
  User,
  X,
  ChevronDown,
  LayoutDashboard, // Imported for use as an icon link
} from 'lucide-react'
import { useState, useEffect, useRef } from 'react'
import { useAuth } from '@/context/AuthContext'
import { Button } from '@/components/ui/Button'
import { ThemeToggle } from '@/components/ui/ThemeToggle'

const navLinkClass = ({ isActive }: { isActive: boolean }) =>
  `rounded-lg px-3 py-2 text-sm font-medium transition-colors ${
    isActive
      ? 'bg-sage-100 text-sage-800 dark:bg-sage-900/50 dark:text-sage-200'
      : 'text-ink-600 hover:bg-sage-50 hover:text-sage-700 dark:text-ink-300 dark:hover:bg-ink-800 dark:hover:text-sage-200'
  }`

export function Navbar() {
  const { isAuthenticated, user, logout, isLoading } = useAuth()
  const navigate = useNavigate()
  const [mobileOpen, setMobileOpen] = useState(false)
  const [dropdownOpen, setDropdownOpen] = useState(false)
  const dropdownRef = useRef<HTMLDivElement>(null)

  const handleLogout = async () => {
    setDropdownOpen(false)
    await logout()
    navigate('/login')
  }

  useEffect(() => {
    function handleClickOutside(event: MouseEvent) {
      if (dropdownRef.current && !dropdownRef.current.contains(event.target as Node)) {
        setDropdownOpen(false)
      }
    }
    document.addEventListener('mousedown', handleClickOutside)
    return () => document.removeEventListener('mousedown', handleClickOutside)
  }, [])

  // Modifying the shared links bundle to dynamically evaluate educator role
  const links = (
    <>
      <NavLink to="/" className={navLinkClass} end onClick={() => setMobileOpen(false)}>
        Home
      </NavLink>
      <NavLink to="/videos" className={navLinkClass} onClick={() => setMobileOpen(false)}>
        Videos
      </NavLink>
      {isAuthenticated ? (
        <NavLink to="/my-playlists" className={navLinkClass} onClick={() => setMobileOpen(false)}>
          My playlists
        </NavLink>
      ) : null}
      {isAuthenticated ? (
        <NavLink to="/watch-history" className={navLinkClass} onClick={() => setMobileOpen(false)}>
          Watch history
        </NavLink>
      ) : null}
      {isAuthenticated && user?.role === 'admin' ? (
        <NavLink to="/admin/courses" className={navLinkClass} onClick={() => setMobileOpen(false)}>
          Admin
        </NavLink>
      ) : null}
      {isAuthenticated && user?.role === 'educator' && (
        <NavLink to="/educator/dashboard" className={navLinkClass} onClick={() => setMobileOpen(false)}>
          Dashboard
        </NavLink>
      )}
    </>
  )

  return (
    <header className="sticky top-0 z-50 border-b border-ink-200/50 bg-cream-50/80 backdrop-blur-xl dark:border-ink-800 dark:bg-ink-950/80">
      <nav className="mx-auto flex h-16 max-w-6xl items-center justify-between gap-4 px-4 sm:px-6">
        {/* Logo */}
        <Link
          to="/"
          className="flex items-center gap-2 font-display text-lg font-semibold text-sage-700 dark:text-sage-300"
        >
          <span className="flex size-9 items-center justify-center rounded-xl bg-gradient-to-br from-sage-500 to-mist-500 text-white shadow-md">
            <GraduationCap className="size-5" />
          </span>
          ByteLearn
        </Link>

        {/* Desktop Links */}
        <div className="hidden items-center gap-1 md:flex">{links}</div>

        {/* Right side controls */}
        <div className="flex items-center gap-2 sm:gap-3">
          <ThemeToggle />

          {!isLoading && (
            <div className="relative hidden items-center sm:flex" ref={dropdownRef}>
              <button
                type="button"
                onClick={() => setDropdownOpen((prev) => !prev)}
                className="flex items-center gap-1.5 rounded-full p-0.5 transition hover:opacity-80 focus:outline-none"
              >
                {isAuthenticated && user?.profile_pic_url ? (
                  <img
                    src={user.profile_pic_url}
                    alt={user.name || 'User profile'}
                    className="size-8 rounded-full border border-sage-200 object-cover dark:border-ink-700"
                  />
                ) : (
                  <div className="flex size-8 items-center justify-center rounded-full border border-ink-200 bg-sage-50 text-ink-600 dark:border-ink-800 dark:bg-ink-900 dark:text-ink-300">
                    <User className="size-4" />
                  </div>
                )}
                <ChevronDown className={`size-3.5 text-ink-400 transition-transform duration-200 ${dropdownOpen ? 'rotate-180' : ''}`} />
              </button>

              {/* Dropdown Menu updates */}
              {dropdownOpen && (
                <div className="absolute right-0 top-full mt-2 w-48 rounded-xl border border-ink-200/60 bg-white p-1.5 shadow-xl animate-in fade-in slide-in-from-top-2 duration-150 dark:border-ink-800 dark:bg-ink-900">
                  {isAuthenticated ? (
                    <>
                      <div className="px-3 py-2 text-xs border-b border-ink-100 dark:border-ink-800 mb-1">
                        <p className="font-medium text-ink-800 dark:text-ink-200 truncate">{user?.name}</p>
                        <p className="text-ink-400 truncate text-xs opacity-75">@{user?.username}</p>
                      </div>

                      {/* Dropdown Quick Access Item for Educators */}
                      {user?.role === 'educator' && (
                        <Link
                          to="/educator/dashboard"
                          onClick={() => setDropdownOpen(false)}
                          className="flex items-center gap-2 rounded-lg px-3 py-2 text-sm text-ink-700 hover:bg-sage-50 hover:text-sage-700 dark:text-ink-300 dark:hover:bg-ink-800 dark:hover:text-sage-200"
                        >
                          <LayoutDashboard className="size-4" />
                          Dashboard
                        </Link>
                      )}

                      <Link
                        to="/profile"
                        onClick={() => setDropdownOpen(false)}
                        className="flex items-center gap-2 rounded-lg px-3 py-2 text-sm text-ink-700 hover:bg-sage-50 hover:text-sage-700 dark:text-ink-300 dark:hover:bg-ink-800 dark:hover:text-sage-200"
                      >
                        <User className="size-4" />
                        Profile
                      </Link>
                      <button
                        type="button"
                        onClick={handleLogout}
                        className="flex w-full items-center gap-2 rounded-lg px-3 py-2 text-left text-sm text-rose-600 hover:bg-rose-50 dark:text-rose-400 dark:hover:bg-rose-950/30"
                      >
                        <LogOut className="size-4" />
                        Logout
                      </button>
                    </>
                  ) : (
                    <>
                      <Link
                        to="/login"
                        onClick={() => setDropdownOpen(false)}
                        className="flex items-center gap-2 rounded-lg px-3 py-2 text-sm text-ink-700 hover:bg-sage-50 hover:text-sage-700 dark:text-ink-300 dark:hover:bg-ink-800 dark:hover:text-sage-200"
                      >
                        Sign in
                      </Link>
                      <Link
                        to="/register"
                        onClick={() => setDropdownOpen(false)}
                        className="flex items-center gap-2 rounded-lg bg-sage-600 px-3 py-2 text-sm font-medium text-white hover:bg-sage-700 dark:bg-sage-700 dark:hover:bg-sage-600"
                      >
                        <BookOpen className="size-4" />
                        Get started
                      </Link>
                    </>
                  )}
                </div>
              )}
            </div>
          )}

          <button
            type="button"
            className="rounded-lg p-2 text-ink-600 md:hidden dark:text-ink-300"
            onClick={() => setMobileOpen((o) => !o)}
            aria-label={mobileOpen ? 'Close menu' : 'Open menu'}
          >
            {mobileOpen ? <X className="size-5" /> : <Menu className="size-5" />}
          </button>
        </div>
      </nav>

      {/* Mobile Menu View */}
      {mobileOpen ? (
        <div className="border-t border-ink-200/50 px-4 py-4 md:hidden dark:border-ink-800">
          <div className="flex flex-col gap-1">{links}</div>
          <div className="mt-4 flex flex-col gap-2 border-t border-ink-200/50 pt-4 dark:border-ink-800">
            {isAuthenticated ? (
              <>
                <div className="flex items-center gap-2 px-3 py-2 text-sm text-ink-500">
                  {user?.profile_pic_url ? (
                    <img src={user.profile_pic_url} alt="" className="size-6 rounded-full object-cover" />
                  ) : (
                    <User className="size-4" />
                  )}
                  <span className="truncate">{user?.name}</span>
                </div>
                <Button variant="ghost" onClick={handleLogout} className="justify-start gap-2">
                  <LogOut className="size-4" />
                  Logout
                </Button>
              </>
            ) : (
              <>
                <Link to="/login" onClick={() => setMobileOpen(false)}>
                  <Button variant="secondary" className="w-full">
                    Sign in
                  </Button>
                </Link>
                <Link to="/register" onClick={() => setMobileOpen(false)}>
                  <Button className="w-full">Get started</Button>
                </Link>
              </>
            )}
          </div>
        </div>
      ) : null}
    </header>
  )
}
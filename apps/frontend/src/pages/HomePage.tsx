import { Link } from 'react-router-dom'
import {
  BookOpen,
  Brain,
  Code2,
  PlayCircle,
  Sparkles,
  Video,
} from 'lucide-react'
import { Button } from '@/components/ui/Button'
import { Card } from '@/components/ui/Card'
import { useAuth } from '@/context/AuthContext'

const features = [
  {
    icon: Video,
    title: 'Video learning',
    description:
      'Structured courses with playlists, watch history, and resume playback.',
    color: 'from-mist-400 to-mist-600',
  },
  {
    icon: Code2,
    title: 'Coding practice',
    description:
      'Contests and problems with isolated execution and submission tracking.',
    color: 'from-sage-400 to-sage-600',
  },
  {
    icon: Brain,
    title: 'Smart quizzes',
    description:
      'Flexible JSON-based assessments with automatic scoring.',
    color: 'from-mist-500 to-sage-500',
  },
  {
    icon: Sparkles,
    title: 'Recommendations',
    description:
      'Personalized content based on your watch history and interactions.',
    color: 'from-sage-500 to-mist-500',
  },
]

export function HomePage() {
  const { isAuthenticated } = useAuth()

  return (
    <div className="animate-fade-in space-y-16">
      <section className="relative overflow-hidden rounded-3xl">
        <div className="glass absolute inset-0 rounded-3xl" />
        <div className="relative px-6 py-14 sm:px-12 sm:py-20">
          <p className="mb-4 inline-flex items-center gap-2 rounded-full bg-sage-100/80 px-4 py-1.5 text-sm font-medium text-sage-700 dark:bg-sage-900/40 dark:text-sage-200">
            <Sparkles className="size-4" />
            Learn calmly. Build confidently.
          </p>
          <h1 className="font-display max-w-2xl text-4xl font-semibold leading-tight tracking-tight text-ink-900 sm:text-5xl dark:text-ink-50">
            Your space for videos, code, and growth
          </h1>
          <p className="mt-5 max-w-xl text-lg leading-relaxed text-ink-600 dark:text-ink-300">
            ByteLearn brings together structured video courses, coding contests,
            quizzes, and personalized recommendations — designed for focused,
            stress-free learning.
          </p>
          <div className="mt-8 flex flex-wrap gap-3">
            <Link to={isAuthenticated ? '/videos' : '/register'}>
              <Button size="lg">
                <PlayCircle className="size-5" />
                {isAuthenticated ? 'Browse videos' : 'Start learning'}
              </Button>
            </Link>
            <Link to="/videos">
              <Button variant="secondary" size="lg">
                <BookOpen className="size-5" />
                Explore catalog
              </Button>
            </Link>
          </div>
        </div>
      </section>

      <section className="animate-slide-up space-y-8" style={{ animationDelay: '0.1s' }}>
        <div className="text-center">
          <h2 className="font-display text-2xl font-semibold text-ink-900 dark:text-ink-50">
            Everything in one calm platform
          </h2>
          <p className="mt-2 text-ink-500 dark:text-ink-400">
            Built with a production-style microservice architecture
          </p>
        </div>
        <div className="grid gap-5 sm:grid-cols-2">
          {features.map(({ icon: Icon, title, description, color }) => (
            <Card key={title} hover className="group">
              <div
                className={`mb-4 inline-flex rounded-xl bg-gradient-to-br ${color} p-3 text-white shadow-lg transition-transform duration-300 group-hover:scale-105`}
              >
                <Icon className="size-6" />
              </div>
              <h3 className="font-display text-lg font-semibold text-ink-900 dark:text-ink-50">
                {title}
              </h3>
              <p className="mt-2 text-sm leading-relaxed text-ink-500 dark:text-ink-400">
                {description}
              </p>
            </Card>
          ))}
        </div>
      </section>
    </div>
  )
}

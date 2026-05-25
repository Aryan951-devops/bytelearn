import { type FormEvent, useEffect, useState } from 'react'
import { KeyRound, Settings, Upload } from 'lucide-react'
import { userApi } from '@/lib/api'
import { useAuth } from '@/context/AuthContext'
import { Button } from '@/components/ui/Button'
import { Card } from '@/components/ui/Card'
import { Input } from '@/components/ui/Input'

export function SettingsPage() {
  const { user, refreshUser } = useAuth()
  const [phone, setPhone] = useState('')
  const [city, setCity] = useState('')
  const [state, setState] = useState('')
  const [pincode, setPincode] = useState('')
  const [profilePic, setProfilePic] = useState<File | null>(null)
  const [newPassword, setNewPassword] = useState('')
  const [profileMsg, setProfileMsg] = useState('')
  const [passwordMsg, setPasswordMsg] = useState('')
  const [profileError, setProfileError] = useState('')
  const [passwordError, setPasswordError] = useState('')
  const [profileLoading, setProfileLoading] = useState(false)
  const [passwordLoading, setPasswordLoading] = useState(false)

  useEffect(() => {
    if (!user) return
    setPhone(user.phone_no ?? '')
    setCity(user.city ?? '')
    setState(user.state ?? '')
    setPincode(user.pincode ?? '')
  }, [user])

  const handleProfileSubmit = async (e: FormEvent) => {
    e.preventDefault()
    setProfileError('')
    setProfileMsg('')
    setProfileLoading(true)
    const form = new FormData()
    if (phone) form.append('phone_no', phone)
    if (city) form.append('city', city)
    if (state) form.append('state', state)
    if (pincode) form.append('pincode', pincode)
    if (profilePic) form.append('profile_pic', profilePic)

    try {
      await userApi.updateAccount(form)
      await refreshUser()
      setProfileMsg('Profile updated successfully.')
      setProfilePic(null)
    } catch (err) {
      setProfileError(
        err instanceof Error ? err.message : 'Failed to update profile',
      )
    } finally {
      setProfileLoading(false)
    }
  }

  const handlePasswordSubmit = async (e: FormEvent) => {
    e.preventDefault()
    setPasswordError('')
    setPasswordMsg('')
    setPasswordLoading(true)
    try {
      await userApi.changePassword(newPassword)
      setPasswordMsg('Password updated successfully.')
      setNewPassword('')
    } catch (err) {
      setPasswordError(
        err instanceof Error ? err.message : 'Failed to change password',
      )
    } finally {
      setPasswordLoading(false)
    }
  }

  return (
    <div className="animate-slide-up mx-auto max-w-2xl space-y-8">
      <div>
        <h1 className="font-display text-3xl font-semibold text-ink-900 dark:text-ink-50">
          Settings
        </h1>
        <p className="mt-2 text-ink-500 dark:text-ink-400">
          Update your profile and security preferences
        </p>
      </div>

      <Card>
        <h2 className="mb-4 flex items-center gap-2 font-display text-lg font-semibold text-ink-900 dark:text-ink-50">
          <Settings className="size-5 text-sage-500" />
          Profile
        </h2>
        <form onSubmit={handleProfileSubmit} className="space-y-4">
          <div className="space-y-1.5">
            <label className="block text-sm font-medium text-ink-700 dark:text-ink-200">
              Profile picture
            </label>
            <label className="flex cursor-pointer items-center gap-3 rounded-xl border border-dashed border-ink-200/80 bg-white/50 px-4 py-6 transition-colors hover:border-sage-400 dark:border-ink-700 dark:bg-ink-900/30 dark:hover:border-sage-500">
              <Upload className="size-5 text-ink-400" />
              <span className="text-sm text-ink-500">
                {profilePic ? profilePic.name : 'Choose an image (optional)'}
              </span>
              <input
                type="file"
                accept="image/*"
                className="sr-only"
                onChange={(e) =>
                  setProfilePic(e.target.files?.[0] ?? null)
                }
              />
            </label>
          </div>
          <Input
            label="Phone"
            type="tel"
            value={phone}
            onChange={(e) => setPhone(e.target.value)}
          />
          <div className="grid gap-4 sm:grid-cols-2">
            <Input
              label="City"
              value={city}
              onChange={(e) => setCity(e.target.value)}
            />
            <Input
              label="State"
              value={state}
              onChange={(e) => setState(e.target.value)}
            />
          </div>
          <Input
            label="Pincode"
            value={pincode}
            onChange={(e) => setPincode(e.target.value)}
          />
          {profileError ? (
            <p className="text-sm text-red-500">{profileError}</p>
          ) : null}
          {profileMsg ? (
            <p className="text-sm text-sage-600 dark:text-sage-400">
              {profileMsg}
            </p>
          ) : null}
          <Button type="submit" isLoading={profileLoading}>
            Save profile
          </Button>
        </form>
      </Card>

      <Card>
        <h2 className="mb-4 flex items-center gap-2 font-display text-lg font-semibold text-ink-900 dark:text-ink-50">
          <KeyRound className="size-5 text-mist-500" />
          Change password
        </h2>
        <form onSubmit={handlePasswordSubmit} className="space-y-4">
          <Input
            label="New password"
            type="password"
            required
            minLength={6}
            value={newPassword}
            onChange={(e) => setNewPassword(e.target.value)}
          />
          {passwordError ? (
            <p className="text-sm text-red-500">{passwordError}</p>
          ) : null}
          {passwordMsg ? (
            <p className="text-sm text-sage-600 dark:text-sage-400">
              {passwordMsg}
            </p>
          ) : null}
          <Button type="submit" variant="secondary" isLoading={passwordLoading}>
            Update password
          </Button>
        </form>
      </Card>
    </div>
  )
}

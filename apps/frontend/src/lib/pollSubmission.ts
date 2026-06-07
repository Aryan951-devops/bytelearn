import { codeApi } from '@/lib/api'
import type { SubmissionStatus } from '@/types'

const TERMINAL = new Set(['COMPLETED', 'FAILED'])

export async function pollSubmission(
  submissionId: string,
  options?: { intervalMs?: number; maxAttempts?: number },
): Promise<SubmissionStatus> {
  const intervalMs = options?.intervalMs ?? 1500
  const maxAttempts = options?.maxAttempts ?? 60

  for (let i = 0; i < maxAttempts; i++) {
    const res = await codeApi.pollSubmission(submissionId)
    const status = res.data?.submission
    if (status && TERMINAL.has(status.status)) {
      return status
    }
    await new Promise((r) => setTimeout(r, intervalMs))
  }

  throw new Error('Execution timed out — try again in a moment.')
}

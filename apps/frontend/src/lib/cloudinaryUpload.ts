import { videoApi } from '@/lib/api'

export interface CloudinaryUploadResult {
  secure_url: string
  public_id: string
  duration?: number
}

export async function uploadVideoToCloudinary(
  file: File,
): Promise<CloudinaryUploadResult> {
  const signRes = await videoApi.getUploadSignature()
  if (!signRes.data) throw new Error('Failed to secure upload signature')

  const { signature, timestamp, api_key, cloud_name, folder } = signRes.data

  const formData = new FormData()
  formData.append('file', file)
  formData.append('api_key', api_key)
  formData.append('timestamp', timestamp.toString())
  formData.append('signature', signature)
  formData.append('folder', folder)

  const res = await fetch(
    `https://api.cloudinary.com/v1_1/${cloud_name}/video/upload`,
    { method: 'POST', body: formData },
  )

  if (!res.ok) throw new Error('Cloudinary upload failed')
  return res.json() as Promise<CloudinaryUploadResult>
}

export async function registerVideoWithBackend(
  upload: CloudinaryUploadResult,
  meta: { title: string; description?: string },
) {
  const res = await videoApi.upload({
    title: meta.title,
    description: meta.description,
    videofile_url: upload.secure_url,
    videofile_public_id: upload.public_id,
    duration_seconds: Math.round(upload.duration ?? 0),
  })
  return res.data?.video
}

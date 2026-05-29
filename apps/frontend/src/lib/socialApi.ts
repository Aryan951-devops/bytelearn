/**
 * Placeholder social APIs — routes are commented out in api-gateway.
 * Swap to real endpoints when comment/like handlers are registered.
 *
 * Planned:
 *   GET    /api/v1/comment/video/:videoId
 *   POST   /api/v1/comment/video/:videoId
 *   POST   /api/v1/like/video/:videoId
 *   DELETE /api/v1/like/video/:videoId
 */

import type { Comment } from '@/types'

const API_BASE = import.meta.env.VITE_API_BASE ?? '/api/v1'
const COMMENTS_KEY = 'bytelearn-comments'
const LIKES_KEY = 'bytelearn-likes'

function readJson<T>(key: string, fallback: T): T {
  try {
    const raw = localStorage.getItem(key)
    return raw ? (JSON.parse(raw) as T) : fallback
  } catch {
    return fallback
  }
}

function writeJson(key: string, value: unknown) {
  localStorage.setItem(key, JSON.stringify(value))
}

async function tryRequest<T>(
  path: string,
  options?: RequestInit,
): Promise<T | null> {
  try {
    const res = await fetch(`${API_BASE}${path}`, {
      credentials: 'include',
      headers: { 'Content-Type': 'application/json', ...options?.headers },
      ...options,
    })
    if (!res.ok) return null
    const body = await res.json()
    return body as T
  } catch {
    return null
  }
}

export const commentApi = {
  async getForVideo(videoId: string): Promise<Comment[]> {
    const remote = await tryRequest<{ data?: { comments?: Comment[] } }>(
      `/comment/video/${videoId}`,
    )
    if (remote?.data?.comments) return remote.data.comments

    const store = readJson<Record<string, Comment[]>>(COMMENTS_KEY, {})
    return (store[videoId] ?? []).sort(
      (a, b) =>
        new Date(b.created_at).getTime() - new Date(a.created_at).getTime(),
    )
  },

  async create(
    videoId: string,
    content: string,
    author: { user_id: string; name: string },
  ): Promise<Comment> {
    const remote = await tryRequest<{ data?: { comment?: Comment } }>(
      `/comment/video/${videoId}`,
      {
        method: 'POST',
        body: JSON.stringify({ content }),
      },
    )
    if (remote?.data?.comment) return remote.data.comment

    const comment: Comment = {
      comment_id: crypto.randomUUID(),
      video_id: videoId,
      user_id: author.user_id,
      user_name: author.name,
      content,
      created_at: new Date().toISOString(),
      likes_count: 0,
    }
    const store = readJson<Record<string, Comment[]>>(COMMENTS_KEY, {})
    store[videoId] = [comment, ...(store[videoId] ?? [])]
    writeJson(COMMENTS_KEY, store)
    return comment
  },
}

export const likeApi = {
  async getVideoLikeCount(videoId: string): Promise<number> {
    const remote = await tryRequest<{ data?: { count?: number } }>(
      `/like/video/${videoId}/count`,
    )
    if (remote?.data?.count != null) return remote.data.count

    const store = readJson<Record<string, number>>(LIKES_KEY, {})
    return store[videoId] ?? 0
  },

  async isLikedByUser(videoId: string, userId: string): Promise<boolean> {
    const remote = await tryRequest<{ data?: { liked?: boolean } }>(
      `/like/video/${videoId}/me`,
    )
    if (remote?.data?.liked != null) return remote.data.liked

    const store = readJson<Record<string, string[]>>(`${LIKES_KEY}-users`, {})
    return (store[videoId] ?? []).includes(userId)
  },

  async toggleVideoLike(
    videoId: string,
    userId: string,
  ): Promise<{ liked: boolean; count: number }> {
    const liked = await likeApi.isLikedByUser(videoId, userId)

    const remote = await tryRequest<{ data?: { liked: boolean; count: number } }>(
      `/like/video/${videoId}`,
      { method: liked ? 'DELETE' : 'POST' },
    )
    if (remote?.data) return remote.data

    const usersStore = readJson<Record<string, string[]>>(`${LIKES_KEY}-users`, {})
    const countStore = readJson<Record<string, number>>(LIKES_KEY, {})
    const users = usersStore[videoId] ?? []
    const has = users.includes(userId)

    if (has) {
      usersStore[videoId] = users.filter((id) => id !== userId)
      countStore[videoId] = Math.max(0, (countStore[videoId] ?? 1) - 1)
    } else {
      usersStore[videoId] = [...users, userId]
      countStore[videoId] = (countStore[videoId] ?? 0) + 1
    }

    writeJson(`${LIKES_KEY}-users`, usersStore)
    writeJson(LIKES_KEY, countStore)

    return {
      liked: !has,
      count: countStore[videoId] ?? 0,
    }
  },
}

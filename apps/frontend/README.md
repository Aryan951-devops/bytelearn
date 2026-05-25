# ByteLearn Frontend

React + TypeScript + Vite + Tailwind CSS v4.

## Run locally

```bash
npm install
npm run dev
```

Start the API gateway on port `8080` so the Vite proxy can forward `/api` requests (required for cookie-based auth).

## Pages ↔ API routes

| Page       | Path        | API endpoint                          |
| ---------- | ----------- | ------------------------------------- |
| Home       | `/`         | —                                     |
| Register   | `/register` | `POST /api/v1/auth/register`          |
| Login      | `/login`    | `POST /api/v1/auth/login`             |
| Videos     | `/videos`   | `GET /api/v1/video/all`               |
| Profile    | `/profile`  | `GET /api/v1/user/current-user`       |
| Settings   | `/settings` | `POST /api/v1/user/update-account`, `POST /api/v1/user/change-password` |

Logout is triggered from the navbar via `POST /api/v1/auth/logout`. Session check uses `POST /api/v1/auth/verifytoken`.

## Theme

Light, dark, and system modes — toggle in the navbar. Preference is stored in `localStorage` under `bytelearn-theme`.

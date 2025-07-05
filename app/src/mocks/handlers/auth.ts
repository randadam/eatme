import { http, HttpResponse } from "msw"
import { db, type UserRecord } from "../db"
import { accountForm } from "../../features/auth/signup/schemas/forms"
import { profilePatchApi } from "@/features/auth/signup/schemas/api"

export const authHandlers = [
  http.post("/api/signup", async ({ request }) => {
    const body = await request.json()
    const parse = accountForm.safeParse(body)
    if (!parse.success) {
      return new HttpResponse(JSON.stringify({ message: "Invalid payload" }), { status: 400 })
    }
    if (db.users.findByEmail(parse.data.email)) {
      return new HttpResponse(JSON.stringify({ message: "Email exists" }), { status: 409 })
    }
    const user = db.users.create(parse.data)
    return new HttpResponse(JSON.stringify(user), { status: 201 })
  }),
  http.patch("/api/profile", async ({ request }) => {
    const authToken = request.headers.get("authorization")
    if (!authToken) {
      return new HttpResponse(JSON.stringify({ message: "Missing auth token" }), { status: 401 })
    }
    const currentUser = getUserFromToken(authToken)
    if (!currentUser) {
      return new HttpResponse(JSON.stringify({ message: "Invalid auth token" }), { status: 401 })
    }
    const body = await request.json()
    const parse = profilePatchApi.safeParse(body)
    if (!parse.success) {
      return new HttpResponse(JSON.stringify({ message: "Invalid payload" }), { status: 400 })
    }
    const user = db.users.update({ ...currentUser, ...parse.data })
    return new HttpResponse(JSON.stringify(user), { status: 201 })
  }),
]

function getUserFromToken(token: string): UserRecord | undefined {
  const id = token.replace("Bearer ", "")
  return db.users.findById(id)
}
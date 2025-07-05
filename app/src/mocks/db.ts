import usersSeed from "./data/users.json"

export interface UserRecord {
  id: string
  email: string
  name?: string
  skill?: string
  cuisines?: string[]
  diet?: string[]
  equipment?: string[]
  allergies?: string[]
}

const KEY = "mock-db-v1"

function load(): UserRecord[] {
  const persisted = localStorage.getItem(KEY)
  return persisted ? JSON.parse(persisted) : usersSeed
}
function save(data: UserRecord[]) {
  localStorage.setItem(KEY, JSON.stringify(data))
}

let users = load()

export const db = {
  users: {
    all: () => users,
    create: (u: Omit<UserRecord, "id">) => {
      const user: UserRecord = { ...u, id: crypto.randomUUID() }
      users.push(user)
      save(users)
      return user
    },
    update: (u: UserRecord) => {
      const index = users.findIndex(u => u.id === u.id)
      if (index === -1) {
        return null
      }
      users[index] = u
      save(users)
      return u
    },
    findByEmail: (email: string) => users.find(u => u.email === email),
    findById: (id: string) => users.find(u => u.id === id),
    clear: () => {
      users = []
      save(users)
    },
  },
}

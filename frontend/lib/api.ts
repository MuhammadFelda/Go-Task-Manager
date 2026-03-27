const BASE = "http://localhost:3000"

async function req<T>(path: string, opts?: RequestInit): Promise<T> {
    const res = await fetch(`${BASE}/api${path}`, {
        headers: { "Content-Type": "application/json" },
        ...opts,
    })

    const json = await res.json()
    if (!res.ok) {
        throw new Error(json.error || "Request Failed")
    }

    return json
}

export const userApi = {
    getAll: () => req<{ data: any[] }>("/users"),
    create: (d: any) => req<{ data: any }>("/users", { method: "POST", body: JSON.stringify(d)}),
    update: (id: number, d: any) => req<{ data:any }>(`/users/${id}`, { method: "PUT", body: JSON.stringify(d) }),
    delete: (id: number) => req<{ data: any }>(`/users/${id}`, { method: "DELETE" }),
}

export const skillApi = {
    getAll: () => req<{ data:any[] }>("/skills"),
    create: (d: any) => req<{ data:any }>("/skills", { method: "POST", body: JSON.stringify(d) }),
    delete: (id: number) => req<{ data: any }>(`/skill/${id}`, { method: "DELETE" }),
}

export const projectOwnerApi = {
    getAll: () => req<{ data: any[] }>("/projectOwners"),
    create: (d: any) => req<{ data: any}>("/projectOwners", { method: "POST", body: JSON.stringify(d) }),
    update: (id: number, d: any) => req<{ data: any }>(`/projectOwners/${id}`, { method: "PUT", body: JSON.stringify(d) }),
    delete: (id: number) => req<{ data: any }>(`/projectOwners/${id}`, { method: "DELETE" }),
}

export const logtimeApi = {
    getAll: () => req<{ data: any }>("/logtimes"),
    create: (d: any) => req<{ data: any}>("/logtimes", { method: "POST", body: JSON.stringify(d) }),
    update: (id: number, d: any) => req<{ data: any }>(`/logtimes/${id}`, { method: "PUT", body: JSON.stringify(d) }),
    delete: (id: number) => req<{ data: any }>(`/logtimes/${id}`, { method: "DELETE" })
}
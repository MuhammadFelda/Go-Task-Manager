const BASE = process.env.NEXT_PUBLIC_API_URL

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
    delete: (id: number) => req<{ data: any }>(`/skills/${id}`, { method: "DELETE" }),
}

export const projectOwnerApi = {
    getAll: () => req<{ data: any[] }>("/project-owners"),
    create: (d: any) => req<{ data: any}>("/project-owners", { method: "POST", body: JSON.stringify(d) }),
    update: (id: number, d: any) => req<{ data: any }>(`/project-owners/${id}`, { method: "PUT", body: JSON.stringify(d) }),
    delete: (id: number) => req<{ data: any }>(`/project-owners/${id}`, { method: "DELETE" }),
}

export const logtimeApi = {
    getAll: () => req<{ data: any }>("/logtimes"),
    create: (d: any) => req<{ data: any}>("/logtimes", { method: "POST", body: JSON.stringify(d) }),
    update: (id: number, d: any) => req<{ data: any }>(`/logtimes/${id}`, { method: "PUT", body: JSON.stringify(d) }),
    delete: (id: number) => req<{ data: any }>(`/logtimes/${id}`, { method: "DELETE" })
}
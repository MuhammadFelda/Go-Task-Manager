"use client"
import Link from "next/link"
import { usePathname } from "next/navigation"
import "./globals.css"
import React from "react"

const nav = [
  { href: "/", label: "Dashboard" },
  { href: "/users", label: "Users" },
  { href: "/projectOwners", label: "Project Owners" },
  { href: "/skills", label: "Skills" },
  { href: "/logtimes", label: "Logtimes" },
]

export default function RootLayout({ children }: { children: React.ReactNode }) {
  return (
    <html lang="en">
      <body className="bg-zinc-900 text-zinc-200 min-h-screen">
        <div className="flex min-h-screen">
          <aside className="w-56 flex shrink-0 flex-col bg-zinc-900 border-r border-zinc-800 ">

            <div className="px-5 py-6 border-b border-zinc-800">
              <h1 className="text-white font-semibold text-lg tracking-tight">Task Manager</h1>
            </div>

            <nav className="flex-1 px-3 py-4 space-y-0.5">
              {nav.map(n => (
                <NavLink key={n.href} href={n.href} label={n.label} /> 
              ))}
            </nav>
          </aside>

          <main className="flex-1 overflow-y-auto">
            <div className="max-w mx-5 px-4 py-4">
              {children}
            </div>
          </main>
        </div>
      </body>
    </html>
  )
}

function NavLink({ href, label }: { href: string, label: string }) {
  return (
    <Link
      href={href}
      className="flex items-center gap-3 px-2 py-2 rounded-lg text-zinc-400 hover:text-white hover:bg-zinc-800 transition-all duration-150 group"
    >
      <span className="text-sm font-medium">
        {label}
      </span>
    </Link>
  )
}
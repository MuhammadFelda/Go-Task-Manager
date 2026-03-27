"use client"
import { useState, useEffect, use } from 'react'

interface Field {
    key: string,
    label: string,
    type?: string,
    required?: boolean,
    placeholder?: string
}

interface Props {
    title: string,
    fields: Field[],
    initial?: Record<string, any>
    onSubmit: (data: Record<string, any>) => Promise<void>
    onClose: () => void,
    submitLabel?: string,
}

export default function formModal({ title, fields, initial = {}, onSubmit, onClose, submitLabel = "Save" }: Props) {
    const [ form, setForm ] = useState<Record<string, any>>(
        Object.fromEntries(fields.map(f => [f.key, initial[f.key] ?? ""]))
    )

    const [ loading, setLoading ] = useState(false)
    const [ error, setError ] = useState<string | null>(null)

    const set = (k: string, v: any) => setForm(f => ({...f, [k]: v}))

    const handleSubmit = () => {
        for (const f of fields) {
            if (f.required && !String(form[f.key]).trim) {
                setError(`${f.label} is required!`)
                return
            }
        }

        try {
            setLoading(true)
            setError(null)
        } catch (err: any) {
            setError(err.message)
        } finally {
            setLoading(false)
        }
    }

    const inputClass = "w-full bg-zinc-900 border border-zinc-700 text-zinc-100 text-sm px-3 py-2.5 rounded-lg outline-none focus:border-indigo-500 focus:ring-1 focus:ring-indigo-500/30 transition-all placeholder:text-zinc-600 font-mono"

    return (
        <div 
            className="fixed inset-0 z-50 flex items-center justify-center bg-black/60 backdrop-blur-sm"
            onClick={onClose}
        >
            <div 
                className="bg-zinc-900 border border-zinc-900 rounded-2xl w-full max-w-md mx-4 shadow-2xl"
                onClick={e => e.stopPropagation()}
            >
                <div className="flex items-center justify-between px-6 py-5 border-b border-zinc-800">
                    <div className="">
                        <p className="font-mono text-xs text-indigo-400 tracking-widest uppercase mb-0.5">Form</p>
                        <h2 className="text-white font-semibold text-base">{title}</h2>
                    </div>
                    <button 
                        onClick={onClose}
                        className="text-zinc-500 hover:text-white transition-colors w-8 h-8 flex items-center justify-center rounded-lg hover:bg-zinc-800 text-lg leading-none"
                    >
                        x
                    </button>
                </div>

                <div className="px-6 py-5 space-y-4">
                    {fields.map(f => (
                        <div key={f.key}>
                            <label className="block text-xs text-zinc-400 mb-1.5 font-medium">
                                {f.label}
                                {f.required && <span className='text-red-400 ml-1'>*</span>}
                            </label>

                            {f.type === "textarea" ? (
                                <textarea
                                    rows={3}
                                    placeholder={f.placeholder}
                                    value={form[f.key]}
                                    onChange={e => set(f.key, e.target.value)}
                                    className={`${inputClass} resize-none`}
                                />
                            ) : (
                                <input
                                    type={f.type || "text"}
                                    placeholder={f.placeholder}
                                    value={form[f.key]}
                                    onChange={e => set(f.key, f.type === "number" ? Number(e.target.value) : e.target.value)}
                                    className={inputClass}
                                />
                            )}
                        </div>
                    ))}

                    {error && (
                        <p className="text-xs text-red-400 bg-red-500/60 border border-red-500/20 rounded-xl px-3 py-2">
                            {error}
                        </p>
                    )}
                </div>

                <div className="flex items-center justify-end gap-3 px-6 py-4 border-t border-zinc-800">
                    <button
                        onClick={onClose}
                        className="text-sm text-zinc-400 hover:text-white px-4 py-2 rounded-lg border border-zinc-700 hover:border-zinc-500 transition-all"
                    >
                        Cancel
                    </button>
                    <button
                        onClick={handleSubmit}
                        disabled={loading}
                        className="text-sm bg-indigo-600 hover:bg-indigo-500 disabled:opacity-50 disabled:cursor-not-allowed text-white rounded-lg px-5 py-2 font-medium transition-all"
                    >
                        {loading ? "Saving..." : submitLabel}
                    </button>
                </div>
            </div>
        </div>
    )
}
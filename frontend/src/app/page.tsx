import { APP_NAME } from '@/constants/config'
import Link from 'next/link'

export default function Home() {
  return (
    <div className="flex min-h-screen flex-col items-center justify-center p-24">
      <h1 className="mb-4 text-4xl font-bold text-gray-900">{APP_NAME}</h1>
      <p className="mb-8 text-lg text-gray-600">A modern knowledge sharing platform</p>
      
      <div className="flex flex-wrap gap-4">
        <Link
          href="/error-demo"
          className="rounded-md bg-blue-600 px-4 py-2 text-white transition hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
        >
          Error Handling Demo
        </Link>
      </div>
    </div>
  )
}

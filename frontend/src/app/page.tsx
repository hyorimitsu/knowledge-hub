import { APP_NAME } from '@/constants/config'

export default function Home() {
  return (
    <div className="flex min-h-screen flex-col items-center justify-center p-24">
      <h1 className="mb-4 text-4xl font-bold text-gray-900">{APP_NAME}</h1>
      <p className="text-lg text-gray-600">A modern knowledge sharing platform</p>
    </div>
  )
}

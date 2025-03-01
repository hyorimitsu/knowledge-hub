import React from 'react'
import Link from 'next/link'
import { FiHome, FiAlertCircle } from 'react-icons/fi'

/**
 * Not Found page for Next.js app
 * This component is used by Next.js to handle 404 errors
 */
export default function NotFound() {
  return (
    <div className="flex min-h-screen flex-col items-center justify-center bg-gray-50 p-4">
      <div className="w-full max-w-md rounded-lg bg-white p-8 shadow-md">
        <div className="mb-6 flex items-center justify-center">
          <div className="flex h-16 w-16 items-center justify-center rounded-full bg-blue-100 text-blue-600">
            <FiAlertCircle size={32} aria-hidden="true" />
          </div>
        </div>
        
        <h1 className="mb-2 text-center text-2xl font-bold text-gray-900">Page Not Found</h1>
        <p className="mb-6 text-center text-gray-600">
          The page you are looking for doesn&apos;t exist or has been moved.
        </p>
        
        <div className="flex justify-center">
          <Link
            href="/"
            className="flex items-center rounded-md bg-blue-600 px-4 py-2 text-white transition hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
          >
            <FiHome className="mr-2" aria-hidden="true" />
            Go to Home
          </Link>
        </div>
      </div>
    </div>
  )
}
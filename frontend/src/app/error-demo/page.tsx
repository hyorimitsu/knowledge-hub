import React from 'react'
import Link from 'next/link'
import { ErrorBoundaryExample } from '@/components/error/ErrorBoundaryExample'

/**
 * Demo page for error handling components
 */
export default function ErrorDemoPage() {
  return (
    <div className="container mx-auto max-w-6xl py-8">
      <div className="mb-8 flex items-center justify-between">
        <h1 className="text-3xl font-bold text-gray-900">Error Handling Demo</h1>
        <Link
          href="/"
          className="rounded-md bg-gray-200 px-4 py-2 text-gray-700 transition hover:bg-gray-300 focus:outline-none focus:ring-2 focus:ring-gray-400 focus:ring-offset-2"
        >
          Back to Home
        </Link>
      </div>
      
      <div className="mb-8 rounded-lg bg-blue-50 p-4 text-blue-800">
        <h2 className="mb-2 text-xl font-semibold">About This Demo</h2>
        <p>
          This page demonstrates various error handling techniques in Next.js:
        </p>
        <ul className="ml-6 mt-2 list-disc space-y-1">
          <li>Component-level error boundaries that catch errors in their child components</li>
          <li>Hook-based error handling for async operations</li>
          <li>User-friendly error displays with recovery options</li>
          <li>Error logging and debugging information (visible in development mode)</li>
        </ul>
      </div>
      
      <ErrorBoundaryExample />
      
      <div className="mt-8 rounded-lg bg-yellow-50 p-4 text-yellow-800">
        <h2 className="mb-2 text-xl font-semibold">Testing Global Error Handling</h2>
        <p className="mb-4">
          To test the global error boundary, you can navigate to a non-existent route or trigger
          an error at the page level.
        </p>
        <div className="flex flex-wrap gap-4">
          <Link
            href="/error-demo/non-existent"
            className="rounded-md bg-yellow-200 px-4 py-2 text-yellow-800 transition hover:bg-yellow-300 focus:outline-none focus:ring-2 focus:ring-yellow-400 focus:ring-offset-2"
          >
            Test 404 Page
          </Link>
          <Link
            href="/error-demo/error-page"
            className="rounded-md bg-red-200 px-4 py-2 text-red-800 transition hover:bg-red-300 focus:outline-none focus:ring-2 focus:ring-red-400 focus:ring-offset-2"
          >
            Test Error Page
          </Link>
        </div>
      </div>
    </div>
  )
}
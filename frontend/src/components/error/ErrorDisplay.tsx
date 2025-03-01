'use client'

import React from 'react'
import { FiAlertTriangle, FiRefreshCw } from 'react-icons/fi'

export interface ErrorDisplayProps {
  title?: string
  message?: string
  error?: Error
  reset?: () => void
  showDetails?: boolean
}

/**
 * ErrorDisplay component for showing user-friendly error messages
 * with optional error details and reset functionality
 */
export const ErrorDisplay: React.FC<ErrorDisplayProps> = ({
  title = 'Something went wrong',
  message = 'An error occurred while processing your request.',
  error,
  reset,
  showDetails = false,
}) => {
  return (
    <div
      className="flex flex-col items-center justify-center rounded-lg bg-white p-6 shadow-md"
      role="alert"
      aria-live="assertive"
    >
      <div className="mb-4 flex h-16 w-16 items-center justify-center rounded-full bg-red-100 text-red-600">
        <FiAlertTriangle size={32} aria-hidden="true" />
      </div>
      <h2 className="mb-2 text-xl font-bold text-gray-800">{title}</h2>
      <p className="mb-4 text-center text-gray-600">{message}</p>

      {/* Show error details in development or when explicitly enabled */}
      {showDetails && error && process.env.NODE_ENV !== 'production' && (
        <div className="mb-4 w-full max-w-lg overflow-auto rounded bg-gray-100 p-4">
          <p className="mb-2 font-mono text-sm font-bold text-gray-800">{error.name}</p>
          <p className="font-mono text-sm text-gray-700">{error.message}</p>
          {error.stack && (
            <details className="mt-2">
              <summary className="cursor-pointer text-sm text-gray-500 hover:text-gray-700">
                Stack trace
              </summary>
              <pre className="mt-2 whitespace-pre-wrap break-words text-xs text-gray-600">
                {error.stack}
              </pre>
            </details>
          )}
        </div>
      )}

      {/* Reset button if reset function is provided */}
      {reset && (
        <button
          onClick={reset}
          className="flex items-center rounded-md bg-blue-600 px-4 py-2 text-white transition hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
          aria-label="Try again"
        >
          <FiRefreshCw className="mr-2" aria-hidden="true" />
          Try again
        </button>
      )}
    </div>
  )
}

export default ErrorDisplay
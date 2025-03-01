'use client'

import React, { useEffect } from 'react'
import ErrorDisplay from './ErrorDisplay'

export interface GlobalErrorBoundaryProps {
  error: Error & { digest?: string }
  reset: () => void
}

/**
 * Global error boundary component for Next.js app
 * This component is used by Next.js to handle errors at the root level
 */
export function GlobalErrorBoundary({ error, reset }: GlobalErrorBoundaryProps) {
  // Log the error to the console in development
  useEffect(() => {
    if (process.env.NODE_ENV !== 'production') {
      console.error('Global error caught by GlobalErrorBoundary:', error)
    }
    
    // Here you could send the error to an error tracking service
    // Example: sendToErrorTrackingService(error)
  }, [error])

  return (
    <html lang="en">
      <body>
        <div className="flex min-h-screen items-center justify-center bg-gray-50 p-4">
          <ErrorDisplay
            title="Something went wrong!"
            message="We're sorry, but something went wrong. Our team has been notified."
            error={error}
            reset={reset}
            showDetails={process.env.NODE_ENV !== 'production'}
          />
        </div>
      </body>
    </html>
  )
}

export default GlobalErrorBoundary
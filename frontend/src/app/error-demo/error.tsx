'use client'

import React from 'react'
import { ErrorDisplay } from '@/components/error/ErrorDisplay'
import { logError } from '@/utils/errorLogger'

/**
 * Error component for the error-demo page
 * This component is used by Next.js to handle errors at the page level
 */
export default function ErrorDemoError({
  error,
  reset,
}: {
  error: Error & { digest?: string }
  reset: () => void
}) {
  // Log the error
  React.useEffect(() => {
    logError(error, {
      componentName: 'ErrorDemoError',
      url: typeof window !== 'undefined' ? window.location.href : undefined,
    })
  }, [error])

  return (
    <div className="container mx-auto max-w-6xl py-8">
      <div className="flex flex-col items-center justify-center p-4">
        <div className="w-full max-w-lg">
          <ErrorDisplay
            title="Error in Demo Page"
            message="An error occurred in the error demo page."
            error={error}
            reset={reset}
            showDetails={true} // Always show details in the demo
          />
          
          <div className="mt-8 rounded-lg bg-blue-50 p-4 text-blue-800">
            <h2 className="mb-2 text-lg font-semibold">About This Error</h2>
            <p>
              This error was caught by the page-level error boundary. In Next.js, each page can have
              its own error handling component that catches errors in the page component and its
              children.
            </p>
            <p className="mt-2">
              You can click the &quot;Try again&quot; button above to reset the error boundary and
              attempt to render the page again.
            </p>
          </div>
        </div>
      </div>
    </div>
  )
}
'use client'

import { GlobalErrorBoundary } from '@/components/error/GlobalErrorBoundary'
import { logError } from '@/utils/errorLogger'

/**
 * Global error handler for Next.js app
 * This component is used by Next.js to handle errors at the root level
 */
export default function GlobalError({
  error,
  reset,
}: {
  error: Error & { digest?: string }
  reset: () => void
}) {
  // Log the error
  logError(error, {
    componentName: 'GlobalError',
    url: typeof window !== 'undefined' ? window.location.href : undefined,
  })

  return <GlobalErrorBoundary error={error} reset={reset} />
}
'use client'

import { useState, useCallback } from 'react'
import { logError, ErrorMetadata } from '@/utils/errorLogger'

/**
 * Custom hook for handling errors in functional components
 * Provides error state and functions to handle errors
 */
export function useErrorHandler(componentName?: string) {
  const [error, setError] = useState<Error | null>(null)
  const [isLoading, setIsLoading] = useState<boolean>(false)

  /**
   * Handle an error by setting the error state and logging it
   */
  const handleError = useCallback(
    (err: Error, additionalMetadata: Omit<ErrorMetadata, 'componentName'> = {}) => {
      setError(err)
      
      // Log the error with component information
      logError(err, {
        componentName,
        url: typeof window !== 'undefined' ? window.location.href : undefined,
        ...additionalMetadata,
      })
      
      return err // Return the error for chaining
    },
    [componentName]
  )

  /**
   * Reset the error state
   */
  const resetError = useCallback(() => {
    setError(null)
  }, [])

  /**
   * Wrap an async function with error handling
   */
  const withErrorHandling = useCallback(
    <T>(
      fn: () => Promise<T>,
      additionalMetadata: Omit<ErrorMetadata, 'componentName'> = {}
    ): Promise<T> => {
      setIsLoading(true)
      setError(null)
      
      return fn()
        .catch((err: Error) => {
          handleError(err, additionalMetadata)
          throw err // Re-throw the error for the caller to handle if needed
        })
        .finally(() => {
          setIsLoading(false)
        })
    },
    [handleError]
  )

  return {
    error,
    isLoading,
    handleError,
    resetError,
    withErrorHandling,
  }
}

export default useErrorHandler
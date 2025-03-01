/**
 * Error logging utility for the application
 * This can be extended to send errors to an external service like Sentry
 */

// Interface for error metadata
export interface ErrorMetadata {
  componentName?: string
  userId?: string
  url?: string
  additionalInfo?: Record<string, unknown>
}

/**
 * Log an error to the console and optionally to an error tracking service
 */
export function logError(error: Error, metadata: ErrorMetadata = {}): void {
  // Always log to console in development
  if (process.env.NODE_ENV !== 'production') {
    console.error('Error logged:', error)
    if (Object.keys(metadata).length > 0) {
      console.error('Error metadata:', metadata)
    }
  }

  // In production, we would send this to an error tracking service
  if (process.env.NODE_ENV === 'production') {
    // Example: Send to error tracking service
    // sendToErrorTrackingService(error, metadata)
  }
}

/**
 * Create a standardized error object with additional context
 */
export function createAppError(
  message: string,
  originalError?: Error,
  metadata: ErrorMetadata = {}
): Error {
  // Create a new error with the message
  const error = new Error(message)
  
  // Add the original error's stack trace if available
  if (originalError?.stack) {
    error.stack = `${error.stack}\nCaused by: ${originalError.stack}`
  }
  
  // Add metadata as a non-enumerable property
  Object.defineProperty(error, 'metadata', {
    value: metadata,
    enumerable: false,
    writable: true,
  })
  
  return error
}

/**
 * Example implementation of sending to an error tracking service
 * This would be replaced with actual implementation using a service like Sentry
 */
// Export the function so it can be used elsewhere if needed
export function sendToErrorTrackingService(error: Error, metadata: ErrorMetadata = {}): void {
  // This is a placeholder for the actual implementation
  // Example with Sentry:
  // Sentry.captureException(error, {
  //   extra: metadata
  // })
  
  // For now, just log to console that we would send this to a service
  if (process.env.NODE_ENV !== 'production') {
    console.log('Would send to error tracking service:', error, metadata)
  }
}

/**
 * Format an error for display to users
 * Ensures sensitive information is not exposed
 */
export function formatErrorForUser(error: Error): string {
  // In production, return a generic message
  if (process.env.NODE_ENV === 'production') {
    return 'An unexpected error occurred. Please try again later.'
  }
  
  // In development, return the actual error message
  return error.message || 'An unknown error occurred'
}
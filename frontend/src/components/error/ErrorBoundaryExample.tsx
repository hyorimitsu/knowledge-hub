'use client'

import React, { useState } from 'react'
import { ErrorBoundary } from './ErrorBoundary'
import { useErrorHandler } from '@/hooks/useErrorHandler'
import { FiAlertOctagon } from 'react-icons/fi'

// Component that will throw an error when the button is clicked
const BuggyCounter: React.FC = () => {
  const [counter, setCounter] = useState(0)
  
  const handleClick = () => {
    if (counter === 5) {
      // Simulate an error when counter reaches 5
      throw new Error('Counter reached 5! This is a simulated error.')
    }
    setCounter(counter + 1)
  }
  
  return (
    <div className="rounded-lg bg-white p-6 shadow-md">
      <h3 className="mb-4 text-xl font-semibold text-gray-800">Buggy Counter</h3>
      <p className="mb-4 text-gray-600">
        This component will throw an error when the counter reaches 5.
      </p>
      <p className="mb-4 text-2xl font-bold text-blue-600">{counter}</p>
      <button
        onClick={handleClick}
        className="rounded-md bg-blue-600 px-4 py-2 text-white transition hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
      >
        Increment
      </button>
    </div>
  )
}

// Component that uses the useErrorHandler hook
const AsyncErrorExample: React.FC = () => {
  const { error, isLoading, resetError, withErrorHandling } = useErrorHandler('AsyncErrorExample')
  
  const simulateAsyncError = () => {
    return withErrorHandling(async () => {
      // Simulate an API call that fails
      await new Promise(resolve => setTimeout(resolve, 1000))
      throw new Error('Failed to fetch data! This is a simulated async error.')
    })
  }
  
  return (
    <div className="rounded-lg bg-white p-6 shadow-md">
      <h3 className="mb-4 text-xl font-semibold text-gray-800">Async Error Example</h3>
      <p className="mb-4 text-gray-600">
        This component demonstrates handling async errors with the useErrorHandler hook.
      </p>
      
      {error && (
        <div className="mb-4 rounded-md bg-red-50 p-4 text-red-800">
          <div className="flex items-center">
            <FiAlertOctagon className="mr-2" />
            <span className="font-medium">Error:</span>
            <span className="ml-2">{error.message}</span>
          </div>
          <button
            onClick={resetError}
            className="mt-2 text-sm text-red-600 hover:text-red-800"
          >
            Dismiss
          </button>
        </div>
      )}
      
      <button
        onClick={simulateAsyncError}
        disabled={isLoading}
        className="rounded-md bg-purple-600 px-4 py-2 text-white transition hover:bg-purple-700 focus:outline-none focus:ring-2 focus:ring-purple-500 focus:ring-offset-2 disabled:bg-purple-300"
      >
        {isLoading ? 'Loading...' : 'Simulate Async Error'}
      </button>
    </div>
  )
}

// Main example component that demonstrates both error handling approaches
export const ErrorBoundaryExample: React.FC = () => {
  const [resetKey, setResetKey] = useState(0)
  
  const handleReset = () => {
    // Change the resetKey to reset the ErrorBoundary
    setResetKey(prevKey => prevKey + 1)
  }
  
  return (
    <div className="space-y-8 p-4">
      <h2 className="text-2xl font-bold text-gray-900">Error Boundary Examples</h2>
      
      <div className="grid gap-8 md:grid-cols-2">
        {/* Component-level error boundary example */}
        <div>
          <h3 className="mb-4 text-lg font-semibold text-gray-800">Component Error Boundary</h3>
          <ErrorBoundary
            resetKey={resetKey}
            onError={(error, errorInfo) => {
              console.error('Caught error in component boundary:', error, errorInfo)
            }}
          >
            <BuggyCounter />
          </ErrorBoundary>
        </div>
        
        {/* Hook-based error handling example */}
        <div>
          <h3 className="mb-4 text-lg font-semibold text-gray-800">Hook-based Error Handling</h3>
          <AsyncErrorExample />
        </div>
      </div>
      
      <div className="mt-4 flex justify-center">
        <button
          onClick={handleReset}
          className="rounded-md bg-gray-600 px-4 py-2 text-white transition hover:bg-gray-700 focus:outline-none focus:ring-2 focus:ring-gray-500 focus:ring-offset-2"
        >
          Reset All Error Boundaries
        </button>
      </div>
    </div>
  )
}

export default ErrorBoundaryExample
'use client'

import React, { Component, ErrorInfo, ReactNode } from 'react'
import ErrorDisplay from './ErrorDisplay'

export interface ErrorBoundaryProps {
  children: ReactNode
  fallback?: ReactNode
  onError?: (error: Error, errorInfo: ErrorInfo) => void
  resetKey?: string | number | boolean | null | undefined // When this key changes, the error boundary will reset
}

export interface ErrorBoundaryState {
  hasError: boolean
  error: Error | null
}

/**
 * Component-level error boundary
 * Catches errors in its child component tree and displays a fallback UI
 */
export class ErrorBoundary extends Component<ErrorBoundaryProps, ErrorBoundaryState> {
  constructor(props: ErrorBoundaryProps) {
    super(props)
    this.state = {
      hasError: false,
      error: null,
    }
  }

  static getDerivedStateFromError(error: Error): ErrorBoundaryState {
    // Update state so the next render will show the fallback UI
    return { hasError: true, error }
  }

  componentDidCatch(error: Error, errorInfo: ErrorInfo): void {
    // Log the error to the console in development
    if (process.env.NODE_ENV !== 'production') {
      console.error('Error caught by ErrorBoundary:', error, errorInfo)
    }

    // Call the onError callback if provided
    if (this.props.onError) {
      this.props.onError(error, errorInfo)
    }

    // Here you could send the error to an error tracking service
    // Example: sendToErrorTrackingService(error, errorInfo)
  }

  componentDidUpdate(prevProps: ErrorBoundaryProps): void {
    // If resetKey changes, reset the error boundary
    if (
      this.state.hasError &&
      prevProps.resetKey !== this.props.resetKey
    ) {
      this.setState({
        hasError: false,
        error: null,
      })
    }
  }

  resetErrorBoundary = (): void => {
    this.setState({
      hasError: false,
      error: null,
    })
  }

  render(): ReactNode {
    if (this.state.hasError) {
      // If a custom fallback is provided, use it
      if (this.props.fallback) {
        return this.props.fallback
      }

      // Otherwise, use the default error display
      return (
        <div className="p-4">
          <ErrorDisplay
            error={this.state.error as Error}
            reset={this.resetErrorBoundary}
            showDetails={process.env.NODE_ENV !== 'production'}
          />
        </div>
      )
    }

    // When there's no error, render children normally
    return this.props.children
  }
}

export default ErrorBoundary
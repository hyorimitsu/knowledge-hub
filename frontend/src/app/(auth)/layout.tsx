import React from 'react'
import { ErrorBoundary } from '@/components/error/ErrorBoundary'
import { AuthProvider } from '@/contexts/AuthContext'

export default function AuthLayout({ children }: { children: React.ReactNode }) {
  return (
    <AuthProvider>
      <ErrorBoundary>
        {children}
      </ErrorBoundary>
    </AuthProvider>
  )
}
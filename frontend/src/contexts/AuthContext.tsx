'use client'

import React, { createContext, useContext, useState, useEffect, ReactNode } from 'react'
import { useRouter } from 'next/navigation'
import { User } from '@/types'
import apiClient from '@/services/api'
import { useErrorHandler } from '@/hooks/useErrorHandler'

interface AuthContextType {
  user: User | null
  isLoading: boolean
  isAuthenticated: boolean
  login: (email: string, password: string, tenantId: string) => Promise<void>
  register: (name: string, email: string, password: string, tenantId: string, role: string) => Promise<void>
  logout: () => void
}

const AuthContext = createContext<AuthContextType | undefined>(undefined)

export const AuthProvider: React.FC<{ children: ReactNode }> = ({ children }) => {
  const [user, setUser] = useState<User | null>(null)
  const [isLoading, setIsLoading] = useState<boolean>(true)
  const router = useRouter()
  const { handleError } = useErrorHandler('AuthContext')

  // Check if user is authenticated on mount
  useEffect(() => {
    const checkAuth = async () => {
      try {
        if (apiClient.isAuthenticated()) {
          const userData = await apiClient.getMe()
          setUser(userData)
        }
      } catch {
        // If token is invalid, clear it
        apiClient.clearToken()
      } finally {
        setIsLoading(false)
      }
    }

    checkAuth()
  }, [])

  // Login function
  const login = async (email: string, password: string, tenantId: string) => {
    setIsLoading(true)
    try {
      // Login and get token
      const tokenResponse = await apiClient.login({ email, password, tenant_id: tenantId })
      
      // Ensure token is set before getting user data
      apiClient.setToken(tokenResponse.token)
      
      // Get user data
      const userData = await apiClient.getMe()
      setUser(userData)
      
      // Redirect to dashboard
      router.push('/dashboard')
    } catch (error) {
      handleError(error as Error)
      throw error
    } finally {
      setIsLoading(false)
    }
  }

  // Register function
  const register = async (name: string, email: string, password: string, tenantId: string, role: string) => {
    setIsLoading(true)
    try {
      // Register and get token
      const tokenResponse = await apiClient.register({ name, email, password, tenant_id: tenantId, role })
      
      // Ensure token is set before getting user data
      apiClient.setToken(tokenResponse.token)
      
      // Get user data
      const userData = await apiClient.getMe()
      setUser(userData)
      
      // Redirect to dashboard
      router.push('/dashboard')
    } catch (error) {
      handleError(error as Error)
      throw error
    } finally {
      setIsLoading(false)
    }
  }

  // Logout function
  const logout = () => {
    apiClient.logout()
    setUser(null)
    router.push('/login')
  }

  const value = {
    user,
    isLoading,
    isAuthenticated: !!user,
    login,
    register,
    logout,
  }

  return <AuthContext.Provider value={value}>{children}</AuthContext.Provider>
}

export const useAuth = (): AuthContextType => {
  const context = useContext(AuthContext)
  if (context === undefined) {
    throw new Error('useAuth must be used within an AuthProvider')
  }
  return context
}
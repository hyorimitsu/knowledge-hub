'use client'

import React, { useEffect, useState } from 'react'
import { FiSave } from 'react-icons/fi'
import { useAuth } from '@/contexts/AuthContext'
import { useErrorHandler } from '@/hooks/useErrorHandler'
// import apiClient from '@/services/api' // Will be used in the future for API calls
import { Tenant, TenantSettings } from '@/types'
import { DEFAULT_TENANT_SETTINGS } from '@/constants/config'

export default function SettingsPage() {
  const { user } = useAuth()
  const [tenant, setTenant] = useState<Tenant | null>(null)
  const [settings, setSettings] = useState<TenantSettings>(DEFAULT_TENANT_SETTINGS)
  const [isSubmitting, setIsSubmitting] = useState(false)
  const { error, handleError, resetError, isLoading, withErrorHandling } = useErrorHandler('SettingsPage')

  const isAdmin = user?.role === 'admin'

  useEffect(() => {
    if (!isAdmin) {
      return
    }

    const fetchTenant = async () => {
      if (!user?.tenantId) return

      await withErrorHandling(async () => {
        // In a real app, you'd fetch the tenant from the API
        // For now, we'll create a mock tenant
        const mockTenant: Tenant = {
          id: user.tenantId || 'tenant-1', // Provide a default value to satisfy TypeScript
          name: 'Knowledge Hub',
          domain: 'knowledge-hub.example.com',
          settings: {
            theme: {
              primaryColor: '#3b82f6',
              secondaryColor: '#64748b',
            },
            features: {
              comments: true,
              tags: true,
              ratings: true,
            },
          },
        }

        setTenant(mockTenant)
        setSettings(mockTenant.settings)
      })
    }

    fetchTenant()
  }, [user, isAdmin, withErrorHandling])

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    resetError()

    if (!tenant) {
      handleError(new Error('Tenant not found'))
      return
    }

    setIsSubmitting(true)
    try {
      // In a real app, you'd update the tenant settings via the API
      // For now, we'll just simulate a successful update
      await new Promise((resolve) => setTimeout(resolve, 1000))

      // Update the tenant with the new settings
      const updatedTenant = {
        ...tenant,
        settings,
      }
      setTenant(updatedTenant)

      alert('Settings updated successfully')
    } catch (err) {
      handleError(err as Error)
    } finally {
      setIsSubmitting(false)
    }
  }

  if (!isAdmin) {
    return (
      <div className="rounded-md bg-yellow-50 p-4">
        <div className="text-sm text-yellow-700">
          You don&apos;t have permission to access this page. Only administrators can manage tenant settings.
        </div>
      </div>
    )
  }

  if (isLoading) {
    return (
      <div className="flex h-64 items-center justify-center">
        <div className="h-8 w-8 animate-spin rounded-full border-b-2 border-t-2 border-blue-500"></div>
      </div>
    )
  }

  if (error) {
    return (
      <div className="rounded-md bg-red-50 p-4">
        <div className="text-sm text-red-700">Failed to load tenant settings: {error.message}</div>
      </div>
    )
  }

  if (!tenant) {
    return (
      <div className="rounded-md bg-yellow-50 p-4">
        <div className="text-sm text-yellow-700">Tenant not found</div>
      </div>
    )
  }

  return (
    <div className="space-y-6">
      <div className="flex items-center justify-between">
        <h1 className="text-2xl font-bold text-gray-900">Tenant Settings</h1>
      </div>

      <div className="rounded-lg bg-white shadow">
        <div className="px-4 py-5 sm:p-6">
          <form onSubmit={handleSubmit} className="space-y-6">
            <div>
              <h2 className="text-lg font-medium text-gray-900">Theme</h2>
              <div className="mt-4 grid grid-cols-1 gap-4 sm:grid-cols-2">
                <div>
                  <label htmlFor="primaryColor" className="block text-sm font-medium text-gray-700">
                    Primary Color
                  </label>
                  <div className="mt-1 flex items-center space-x-2">
                    <input
                      type="color"
                      name="primaryColor"
                      id="primaryColor"
                      value={settings.theme.primaryColor}
                      onChange={(e) =>
                        setSettings({
                          ...settings,
                          theme: {
                            ...settings.theme,
                            primaryColor: e.target.value,
                          },
                        })
                      }
                      className="h-8 w-8 rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
                    />
                    <input
                      type="text"
                      value={settings.theme.primaryColor}
                      onChange={(e) =>
                        setSettings({
                          ...settings,
                          theme: {
                            ...settings.theme,
                            primaryColor: e.target.value,
                          },
                        })
                      }
                      className="block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500 sm:text-sm"
                      placeholder="#000000"
                      pattern="^#([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})$"
                      title="Hexadecimal color code (e.g. #ff0000)"
                      required
                    />
                  </div>
                </div>
                <div>
                  <label htmlFor="secondaryColor" className="block text-sm font-medium text-gray-700">
                    Secondary Color
                  </label>
                  <div className="mt-1 flex items-center space-x-2">
                    <input
                      type="color"
                      name="secondaryColor"
                      id="secondaryColor"
                      value={settings.theme.secondaryColor}
                      onChange={(e) =>
                        setSettings({
                          ...settings,
                          theme: {
                            ...settings.theme,
                            secondaryColor: e.target.value,
                          },
                        })
                      }
                      className="h-8 w-8 rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
                    />
                    <input
                      type="text"
                      value={settings.theme.secondaryColor}
                      onChange={(e) =>
                        setSettings({
                          ...settings,
                          theme: {
                            ...settings.theme,
                            secondaryColor: e.target.value,
                          },
                        })
                      }
                      className="block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500 sm:text-sm"
                      placeholder="#000000"
                      pattern="^#([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})$"
                      title="Hexadecimal color code (e.g. #ff0000)"
                      required
                    />
                  </div>
                </div>
              </div>
            </div>

            <div>
              <h2 className="text-lg font-medium text-gray-900">Features</h2>
              <div className="mt-4 space-y-4">
                <div className="flex items-center">
                  <input
                    id="comments"
                    name="comments"
                    type="checkbox"
                    checked={settings.features.comments}
                    onChange={(e) =>
                      setSettings({
                        ...settings,
                        features: {
                          ...settings.features,
                          comments: e.target.checked,
                        },
                      })
                    }
                    className="h-4 w-4 rounded border-gray-300 text-blue-600 focus:ring-blue-500"
                  />
                  <label htmlFor="comments" className="ml-3 block text-sm font-medium text-gray-700">
                    Enable Comments
                  </label>
                </div>
                <div className="flex items-center">
                  <input
                    id="tags"
                    name="tags"
                    type="checkbox"
                    checked={settings.features.tags}
                    onChange={(e) =>
                      setSettings({
                        ...settings,
                        features: {
                          ...settings.features,
                          tags: e.target.checked,
                        },
                      })
                    }
                    className="h-4 w-4 rounded border-gray-300 text-blue-600 focus:ring-blue-500"
                  />
                  <label htmlFor="tags" className="ml-3 block text-sm font-medium text-gray-700">
                    Enable Tags
                  </label>
                </div>
                <div className="flex items-center">
                  <input
                    id="ratings"
                    name="ratings"
                    type="checkbox"
                    checked={settings.features.ratings}
                    onChange={(e) =>
                      setSettings({
                        ...settings,
                        features: {
                          ...settings.features,
                          ratings: e.target.checked,
                        },
                      })
                    }
                    className="h-4 w-4 rounded border-gray-300 text-blue-600 focus:ring-blue-500"
                  />
                  <label htmlFor="ratings" className="ml-3 block text-sm font-medium text-gray-700">
                    Enable Ratings
                  </label>
                </div>
              </div>
            </div>

            <div className="flex justify-end">
              <button
                type="submit"
                disabled={isSubmitting}
                className="inline-flex items-center rounded-md border border-transparent bg-blue-600 px-4 py-2 text-sm font-medium text-white shadow-sm hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2 disabled:bg-blue-300"
              >
                <FiSave className="-ml-1 mr-2 h-5 w-5" aria-hidden="true" />
                {isSubmitting ? 'Saving...' : 'Save Settings'}
              </button>
            </div>
          </form>
        </div>
      </div>
    </div>
  )
}
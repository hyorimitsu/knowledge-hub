'use client'

import React, { ReactNode, useState } from 'react'
import Link from 'next/link'
import { usePathname } from 'next/navigation'
import { FiMenu, FiX, FiHome, FiBook, FiTag, FiSettings, FiLogOut } from 'react-icons/fi'
import { useAuth } from '@/contexts/AuthContext'
import { ErrorBoundary } from '@/components/error/ErrorBoundary'

interface AppLayoutProps {
  children: ReactNode
}

export const AppLayout: React.FC<AppLayoutProps> = ({ children }) => {
  const [sidebarOpen, setSidebarOpen] = useState(false)
  const pathname = usePathname()
  const { user, logout } = useAuth()

  const toggleSidebar = () => {
    setSidebarOpen(!sidebarOpen)
  }

  const closeSidebar = () => {
    setSidebarOpen(false)
  }

  const isActive = (path: string) => {
    return pathname === path
  }

  return (
    <div className="flex h-screen bg-gray-100">
      {/* Sidebar for mobile */}
      <div
        className={`fixed inset-0 z-20 bg-black bg-opacity-50 transition-opacity lg:hidden ${
          sidebarOpen ? 'opacity-100' : 'opacity-0 pointer-events-none'
        }`}
        onClick={closeSidebar}
      />

      {/* Sidebar */}
      <aside
        className={`fixed inset-y-0 left-0 z-30 w-64 transform bg-white shadow-lg transition-transform lg:translate-x-0 lg:static lg:inset-0 ${
          sidebarOpen ? 'translate-x-0' : '-translate-x-full'
        }`}
      >
        <div className="flex h-full flex-col">
          {/* Sidebar header */}
          <div className="flex items-center justify-between border-b px-4 py-5 lg:py-6">
            <Link href="/dashboard" className="text-xl font-bold text-gray-800">
              Knowledge Hub
            </Link>
            <button
              onClick={closeSidebar}
              className="rounded-md p-2 text-gray-500 hover:bg-gray-100 hover:text-gray-600 lg:hidden"
            >
              <FiX size={24} />
            </button>
          </div>

          {/* Sidebar content */}
          <nav className="flex-1 overflow-y-auto p-4">
            <ul className="space-y-2">
              <li>
                <Link
                  href="/dashboard"
                  className={`flex items-center rounded-md px-4 py-2 text-gray-600 hover:bg-blue-50 hover:text-blue-600 ${
                    isActive('/dashboard') ? 'bg-blue-50 text-blue-600' : ''
                  }`}
                  onClick={closeSidebar}
                >
                  <FiHome className="mr-3" />
                  Dashboard
                </Link>
              </li>
              <li>
                <Link
                  href="/knowledge"
                  className={`flex items-center rounded-md px-4 py-2 text-gray-600 hover:bg-blue-50 hover:text-blue-600 ${
                    isActive('/knowledge') || pathname?.startsWith('/knowledge/')
                      ? 'bg-blue-50 text-blue-600'
                      : ''
                  }`}
                  onClick={closeSidebar}
                >
                  <FiBook className="mr-3" />
                  Knowledge
                </Link>
              </li>
              <li>
                <Link
                  href="/tags"
                  className={`flex items-center rounded-md px-4 py-2 text-gray-600 hover:bg-blue-50 hover:text-blue-600 ${
                    isActive('/tags') ? 'bg-blue-50 text-blue-600' : ''
                  }`}
                  onClick={closeSidebar}
                >
                  <FiTag className="mr-3" />
                  Tags
                </Link>
              </li>
              {user?.role === 'admin' && (
                <li>
                  <Link
                    href="/settings"
                    className={`flex items-center rounded-md px-4 py-2 text-gray-600 hover:bg-blue-50 hover:text-blue-600 ${
                      isActive('/settings') ? 'bg-blue-50 text-blue-600' : ''
                    }`}
                    onClick={closeSidebar}
                  >
                    <FiSettings className="mr-3" />
                    Settings
                  </Link>
                </li>
              )}
            </ul>
          </nav>

          {/* Sidebar footer */}
          <div className="border-t p-4">
            <div className="mb-4 flex items-center">
              <div className="h-10 w-10 rounded-full bg-blue-100 text-blue-500">
                <div className="flex h-full w-full items-center justify-center">
                  {user?.name?.charAt(0) || 'U'}
                </div>
              </div>
              <div className="ml-3">
                <p className="font-medium text-gray-800">{user?.name}</p>
                <p className="text-sm text-gray-500">{user?.email}</p>
              </div>
            </div>
            <button
              onClick={logout}
              className="flex w-full items-center rounded-md px-4 py-2 text-gray-600 hover:bg-red-50 hover:text-red-600"
            >
              <FiLogOut className="mr-3" />
              Logout
            </button>
          </div>
        </div>
      </aside>

      {/* Main content */}
      <div className="flex flex-1 flex-col overflow-hidden">
        {/* Header */}
        <header className="bg-white shadow-sm">
          <div className="flex h-16 items-center justify-between px-4">
            <button
              onClick={toggleSidebar}
              className="rounded-md p-2 text-gray-500 hover:bg-gray-100 hover:text-gray-600 lg:hidden"
            >
              <FiMenu size={24} />
            </button>
            <div className="flex items-center">
              <span className="text-sm font-medium text-gray-500">
                {user?.name ? `Welcome, ${user.name}` : 'Welcome'}
              </span>
            </div>
          </div>
        </header>

        {/* Page content */}
        <main className="flex-1 overflow-y-auto p-4">
          <ErrorBoundary>
            {children}
          </ErrorBoundary>
        </main>
      </div>
    </div>
  )
}

export default AppLayout
'use client'

import React, { useEffect, useState } from 'react'
import Link from 'next/link'
import { FiPlus, FiBook, FiTag, FiUser } from 'react-icons/fi'
import { useAuth } from '@/contexts/AuthContext'
import { useErrorHandler } from '@/hooks/useErrorHandler'
import apiClient from '@/services/api'
import { Knowledge, Tag } from '@/types'

export default function DashboardPage() {
  const { user } = useAuth()
  const [recentKnowledge, setRecentKnowledge] = useState<Knowledge[]>([])
  const [tags, setTags] = useState<Tag[]>([])
  const { error, isLoading, withErrorHandling } = useErrorHandler('DashboardPage')

  useEffect(() => {
    const fetchData = async () => {
      await withErrorHandling(async () => {
        // Fetch recent knowledge
        const knowledgeData = await apiClient.getKnowledgeList()
        setRecentKnowledge(knowledgeData.slice(0, 5)) // Get only the 5 most recent

        // Fetch tags
        const tagsData = await apiClient.getTagList()
        setTags(tagsData)
      })
    }

    fetchData()
  }, [withErrorHandling])

  return (
    <div className="space-y-6">
      <div className="flex items-center justify-between">
        <h1 className="text-2xl font-bold text-gray-900">Dashboard</h1>
        <Link
          href="/knowledge/new"
          className="inline-flex items-center rounded-md bg-blue-600 px-4 py-2 text-sm font-medium text-white shadow-sm hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
        >
          <FiPlus className="-ml-1 mr-2 h-5 w-5" aria-hidden="true" />
          New Knowledge
        </Link>
      </div>

      <div className="grid grid-cols-1 gap-6 md:grid-cols-3">
        {/* Stats cards */}
        <div className="rounded-lg bg-white p-6 shadow">
          <div className="flex items-center">
            <div className="flex h-12 w-12 items-center justify-center rounded-md bg-blue-100 text-blue-600">
              <FiBook className="h-6 w-6" />
            </div>
            <div className="ml-4">
              <h2 className="text-lg font-medium text-gray-900">Knowledge</h2>
              <p className="text-3xl font-semibold text-gray-700">{recentKnowledge.length}</p>
            </div>
          </div>
        </div>

        <div className="rounded-lg bg-white p-6 shadow">
          <div className="flex items-center">
            <div className="flex h-12 w-12 items-center justify-center rounded-md bg-green-100 text-green-600">
              <FiTag className="h-6 w-6" />
            </div>
            <div className="ml-4">
              <h2 className="text-lg font-medium text-gray-900">Tags</h2>
              <p className="text-3xl font-semibold text-gray-700">{tags.length}</p>
            </div>
          </div>
        </div>

        <div className="rounded-lg bg-white p-6 shadow">
          <div className="flex items-center">
            <div className="flex h-12 w-12 items-center justify-center rounded-md bg-purple-100 text-purple-600">
              <FiUser className="h-6 w-6" />
            </div>
            <div className="ml-4">
              <h2 className="text-lg font-medium text-gray-900">Role</h2>
              <p className="text-xl font-semibold text-gray-700 capitalize">{user?.role}</p>
            </div>
          </div>
        </div>
      </div>

      {/* Recent knowledge */}
      <div className="rounded-lg bg-white p-6 shadow">
        <h2 className="mb-4 text-lg font-medium text-gray-900">Recent Knowledge</h2>
        {isLoading ? (
          <div className="flex h-40 items-center justify-center">
            <div className="h-8 w-8 animate-spin rounded-full border-b-2 border-t-2 border-blue-500"></div>
          </div>
        ) : error ? (
          <div className="rounded-md bg-red-50 p-4">
            <div className="text-sm text-red-700">
              Failed to load recent knowledge: {error.message}
            </div>
          </div>
        ) : recentKnowledge.length === 0 ? (
          <div className="rounded-md bg-gray-50 p-4 text-center">
            <p className="text-gray-600">No knowledge entries yet.</p>
            <Link
              href="/knowledge/new"
              className="mt-2 inline-flex items-center text-sm font-medium text-blue-600 hover:text-blue-500"
            >
              <FiPlus className="mr-1 h-4 w-4" />
              Create your first knowledge entry
            </Link>
          </div>
        ) : (
          <div className="overflow-hidden rounded-md border border-gray-200">
            <ul className="divide-y divide-gray-200">
              {recentKnowledge.map((knowledge) => (
                <li key={knowledge.id}>
                  <Link
                    href={`/knowledge/${knowledge.id}`}
                    className="block hover:bg-gray-50"
                  >
                    <div className="px-4 py-4 sm:px-6">
                      <div className="flex items-center justify-between">
                        <p className="truncate text-sm font-medium text-blue-600">
                          {knowledge.title}
                        </p>
                        <div className="ml-2 flex flex-shrink-0">
                          <p className="inline-flex rounded-full bg-green-100 px-2 text-xs font-semibold leading-5 text-green-800">
                            {new Date(knowledge.createdAt).toLocaleDateString()}
                          </p>
                        </div>
                      </div>
                      <div className="mt-2 sm:flex sm:justify-between">
                        <div className="sm:flex">
                          <p className="flex items-center text-sm text-gray-500">
                            {knowledge.content.substring(0, 100)}
                            {knowledge.content.length > 100 ? '...' : ''}
                          </p>
                        </div>
                      </div>
                    </div>
                  </Link>
                </li>
              ))}
            </ul>
          </div>
        )}
      </div>

      {/* Tags */}
      <div className="rounded-lg bg-white p-6 shadow">
        <div className="flex items-center justify-between mb-4">
          <h2 className="text-lg font-medium text-gray-900">Tags</h2>
          <Link
            href="/tags"
            className="text-sm font-medium text-blue-600 hover:text-blue-500"
          >
            View all
          </Link>
        </div>
        {isLoading ? (
          <div className="flex h-20 items-center justify-center">
            <div className="h-8 w-8 animate-spin rounded-full border-b-2 border-t-2 border-blue-500"></div>
          </div>
        ) : error ? (
          <div className="rounded-md bg-red-50 p-4">
            <div className="text-sm text-red-700">Failed to load tags: {error.message}</div>
          </div>
        ) : tags.length === 0 ? (
          <div className="rounded-md bg-gray-50 p-4 text-center">
            <p className="text-gray-600">No tags yet.</p>
            <Link
              href="/tags"
              className="mt-2 inline-flex items-center text-sm font-medium text-blue-600 hover:text-blue-500"
            >
              <FiPlus className="mr-1 h-4 w-4" />
              Create your first tag
            </Link>
          </div>
        ) : (
          <div className="flex flex-wrap gap-2">
            {tags.map((tag) => (
              <span
                key={tag.id}
                className="inline-flex items-center rounded-full px-3 py-0.5 text-sm font-medium"
                style={{
                  backgroundColor: `${tag.color}20`, // Add transparency
                  color: tag.color,
                }}
              >
                {tag.name}
              </span>
            ))}
          </div>
        )}
      </div>
    </div>
  )
}
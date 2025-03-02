'use client'

import React, { useEffect, useState } from 'react'
import Link from 'next/link'
import { FiPlus, FiSearch, FiFilter, FiX } from 'react-icons/fi'
import { useErrorHandler } from '@/hooks/useErrorHandler'
import apiClient from '@/services/api'
import { Knowledge, Tag } from '@/types'

export default function KnowledgeListPage() {
  const [knowledge, setKnowledge] = useState<Knowledge[]>([])
  const [tags, setTags] = useState<Tag[]>([])
  const [searchQuery, setSearchQuery] = useState('')
  const [selectedTags, setSelectedTags] = useState<string[]>([])
  const [showFilters, setShowFilters] = useState(false)
  const { error, isLoading, withErrorHandling } = useErrorHandler('KnowledgeListPage')

  useEffect(() => {
    const fetchData = async () => {
      await withErrorHandling(async () => {
        // Fetch knowledge
        const knowledgeData = await apiClient.getKnowledgeList()
        setKnowledge(knowledgeData)

        // Fetch tags
        const tagsData = await apiClient.getTagList()
        setTags(tagsData)
      })
    }

    fetchData()
  }, [withErrorHandling])

  const filteredKnowledge = knowledge.filter((item) => {
    // Filter by search query
    const matchesSearch =
      searchQuery === '' ||
      item.title.toLowerCase().includes(searchQuery.toLowerCase()) ||
      item.content.toLowerCase().includes(searchQuery.toLowerCase())

    // Filter by selected tags
    const matchesTags =
      selectedTags.length === 0 ||
      selectedTags.some((tagId) => item.tags.includes(tagId))

    return matchesSearch && matchesTags
  })

  const handleTagToggle = (tagId: string) => {
    setSelectedTags((prev) =>
      prev.includes(tagId) ? prev.filter((id) => id !== tagId) : [...prev, tagId]
    )
  }

  const clearFilters = () => {
    setSearchQuery('')
    setSelectedTags([])
  }

  return (
    <div className="space-y-6">
      <div className="flex items-center justify-between">
        <h1 className="text-2xl font-bold text-gray-900">Knowledge</h1>
        <Link
          href="/knowledge/new"
          className="inline-flex items-center rounded-md bg-blue-600 px-4 py-2 text-sm font-medium text-white shadow-sm hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
        >
          <FiPlus className="-ml-1 mr-2 h-5 w-5" aria-hidden="true" />
          New Knowledge
        </Link>
      </div>

      {/* Search and filters */}
      <div className="rounded-lg bg-white p-4 shadow">
        <div className="flex flex-col space-y-4 sm:flex-row sm:items-center sm:space-x-4 sm:space-y-0">
          <div className="relative flex-1">
            <div className="pointer-events-none absolute inset-y-0 left-0 flex items-center pl-3">
              <FiSearch className="h-5 w-5 text-gray-400" aria-hidden="true" />
            </div>
            <input
              type="text"
              value={searchQuery}
              onChange={(e) => setSearchQuery(e.target.value)}
              className="block w-full rounded-md border-gray-300 pl-10 focus:border-blue-500 focus:ring-blue-500 sm:text-sm"
              placeholder="Search knowledge..."
            />
          </div>
          <button
            type="button"
            onClick={() => setShowFilters(!showFilters)}
            className="inline-flex items-center rounded-md border border-gray-300 bg-white px-4 py-2 text-sm font-medium text-gray-700 shadow-sm hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
          >
            <FiFilter className="-ml-1 mr-2 h-5 w-5 text-gray-400" aria-hidden="true" />
            Filters
            {selectedTags.length > 0 && (
              <span className="ml-2 inline-flex items-center rounded-full bg-blue-100 px-2.5 py-0.5 text-xs font-medium text-blue-800">
                {selectedTags.length}
              </span>
            )}
          </button>
          {(searchQuery || selectedTags.length > 0) && (
            <button
              type="button"
              onClick={clearFilters}
              className="inline-flex items-center rounded-md border border-gray-300 bg-white px-4 py-2 text-sm font-medium text-gray-700 shadow-sm hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
            >
              <FiX className="-ml-1 mr-2 h-5 w-5 text-gray-400" aria-hidden="true" />
              Clear
            </button>
          )}
        </div>

        {/* Tag filters */}
        {showFilters && (
          <div className="mt-4 flex flex-wrap gap-2">
            {tags.map((tag) => (
              <button
                key={tag.id}
                onClick={() => handleTagToggle(tag.id)}
                className={`inline-flex items-center rounded-full px-3 py-0.5 text-sm font-medium ${
                  selectedTags.includes(tag.id)
                    ? 'bg-blue-100 text-blue-800'
                    : 'bg-gray-100 text-gray-800'
                }`}
                style={
                  selectedTags.includes(tag.id)
                    ? { backgroundColor: `${tag.color}30`, color: tag.color }
                    : {}
                }
              >
                {tag.name}
              </button>
            ))}
          </div>
        )}
      </div>

      {/* Knowledge list */}
      <div className="rounded-lg bg-white shadow">
        {isLoading ? (
          <div className="flex h-64 items-center justify-center">
            <div className="h-8 w-8 animate-spin rounded-full border-b-2 border-t-2 border-blue-500"></div>
          </div>
        ) : error ? (
          <div className="p-6">
            <div className="rounded-md bg-red-50 p-4">
              <div className="text-sm text-red-700">
                Failed to load knowledge: {error.message}
              </div>
            </div>
          </div>
        ) : filteredKnowledge.length === 0 ? (
          <div className="p-6">
            <div className="rounded-md bg-gray-50 p-4 text-center">
              <p className="text-gray-600">No knowledge entries found.</p>
              {knowledge.length > 0 ? (
                <button
                  onClick={clearFilters}
                  className="mt-2 text-sm font-medium text-blue-600 hover:text-blue-500"
                >
                  Clear filters
                </button>
              ) : (
                <Link
                  href="/knowledge/new"
                  className="mt-2 inline-flex items-center text-sm font-medium text-blue-600 hover:text-blue-500"
                >
                  <FiPlus className="mr-1 h-4 w-4" />
                  Create your first knowledge entry
                </Link>
              )}
            </div>
          </div>
        ) : (
          <ul className="divide-y divide-gray-200">
            {filteredKnowledge.map((item) => (
              <li key={item.id}>
                <Link
                  href={`/knowledge/${item.id}`}
                  className="block hover:bg-gray-50"
                >
                  <div className="px-6 py-4">
                    <div className="flex items-center justify-between">
                      <h2 className="text-lg font-medium text-blue-600">{item.title}</h2>
                      <p className="text-sm text-gray-500">
                        {new Date(item.createdAt).toLocaleDateString()}
                      </p>
                    </div>
                    <p className="mt-2 text-sm text-gray-600">
                      {item.content.substring(0, 150)}
                      {item.content.length > 150 ? '...' : ''}
                    </p>
                    {item.tags.length > 0 && (
                      <div className="mt-2 flex flex-wrap gap-2">
                        {item.tags.map((tagId) => {
                          const tag = tags.find((t) => t.id === tagId)
                          if (!tag) return null
                          return (
                            <span
                              key={tag.id}
                              className="inline-flex items-center rounded-full px-2.5 py-0.5 text-xs font-medium"
                              style={{
                                backgroundColor: `${tag.color}20`,
                                color: tag.color,
                              }}
                            >
                              {tag.name}
                            </span>
                          )
                        })}
                      </div>
                    )}
                  </div>
                </Link>
              </li>
            ))}
          </ul>
        )}
      </div>
    </div>
  )
}
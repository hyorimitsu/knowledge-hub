'use client'

import React, { useEffect, useState } from 'react'
import { useRouter } from 'next/navigation'
import Link from 'next/link'
import { FiArrowLeft, FiSave } from 'react-icons/fi'
import { useAuth } from '@/contexts/AuthContext'
import { useErrorHandler } from '@/hooks/useErrorHandler'
import apiClient from '@/services/api'
import { Tag } from '@/types'

export default function NewKnowledgePage() {
  const router = useRouter()
  const { } = useAuth() // We'll use auth context in the future if needed
  const [title, setTitle] = useState('')
  const [content, setContent] = useState('')
  const [selectedTags, setSelectedTags] = useState<string[]>([])
  const [tags, setTags] = useState<Tag[]>([])
  const [isSubmitting, setIsSubmitting] = useState(false)
  const { error, handleError, resetError } = useErrorHandler('NewKnowledgePage')

  useEffect(() => {
    const fetchTags = async () => {
      try {
        const tagsData = await apiClient.getTagList()
        setTags(tagsData)
      } catch (err) {
        handleError(err as Error)
      }
    }

    fetchTags()
  }, [handleError])

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    resetError()

    if (!title.trim() || !content.trim()) {
      handleError(new Error('Title and content are required'))
      return
    }

    setIsSubmitting(true)
    try {
      await apiClient.createKnowledge({
        title,
        content,
        tag_ids: selectedTags,
      })
      router.push('/knowledge')
    } catch (err) {
      handleError(err as Error)
      setIsSubmitting(false)
    }
  }

  const handleTagToggle = (tagId: string) => {
    setSelectedTags((prev) =>
      prev.includes(tagId) ? prev.filter((id) => id !== tagId) : [...prev, tagId]
    )
  }

  return (
    <div className="space-y-6">
      <div className="flex items-center justify-between">
        <Link
          href="/knowledge"
          className="inline-flex items-center text-sm font-medium text-blue-600 hover:text-blue-500"
        >
          <FiArrowLeft className="mr-2 h-4 w-4" />
          Back to Knowledge
        </Link>
      </div>

      <div className="rounded-lg bg-white shadow">
        <div className="px-4 py-5 sm:p-6">
          <h1 className="text-2xl font-bold text-gray-900">Create New Knowledge</h1>

          {error && (
            <div className="mt-4 rounded-md bg-red-50 p-4">
              <div className="text-sm text-red-700">{error.message}</div>
            </div>
          )}

          <form onSubmit={handleSubmit} className="mt-6 space-y-6">
            <div>
              <label htmlFor="title" className="block text-sm font-medium text-gray-700">
                Title
              </label>
              <div className="mt-1">
                <input
                  type="text"
                  name="title"
                  id="title"
                  value={title}
                  onChange={(e) => setTitle(e.target.value)}
                  className="block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500 sm:text-sm"
                  placeholder="Enter a title"
                  required
                />
              </div>
            </div>

            <div>
              <label htmlFor="content" className="block text-sm font-medium text-gray-700">
                Content
              </label>
              <div className="mt-1">
                <textarea
                  id="content"
                  name="content"
                  rows={10}
                  value={content}
                  onChange={(e) => setContent(e.target.value)}
                  className="block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500 sm:text-sm"
                  placeholder="Enter content"
                  required
                />
              </div>
            </div>

            <div>
              <label className="block text-sm font-medium text-gray-700">Tags</label>
              <div className="mt-2 flex flex-wrap gap-2">
                {tags.map((tag) => (
                  <button
                    key={tag.id}
                    type="button"
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
            </div>

            <div className="flex justify-end">
              <Link
                href="/knowledge"
                className="rounded-md border border-gray-300 bg-white px-4 py-2 text-sm font-medium text-gray-700 shadow-sm hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
              >
                Cancel
              </Link>
              <button
                type="submit"
                disabled={isSubmitting}
                className="ml-3 inline-flex items-center rounded-md border border-transparent bg-blue-600 px-4 py-2 text-sm font-medium text-white shadow-sm hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2 disabled:bg-blue-300"
              >
                <FiSave className="-ml-1 mr-2 h-5 w-5" aria-hidden="true" />
                {isSubmitting ? 'Saving...' : 'Save'}
              </button>
            </div>
          </form>
        </div>
      </div>
    </div>
  )
}
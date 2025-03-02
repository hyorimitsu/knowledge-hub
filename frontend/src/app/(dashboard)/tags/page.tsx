'use client'

import React, { useEffect, useState } from 'react'
import { FiPlus, FiEdit, FiTrash2, FiX } from 'react-icons/fi'
import { useAuth } from '@/contexts/AuthContext'
import { useErrorHandler } from '@/hooks/useErrorHandler'
import apiClient from '@/services/api'
import { Tag } from '@/types'

export default function TagsPage() {
  const { user } = useAuth()
  const [tags, setTags] = useState<Tag[]>([])
  const [isModalOpen, setIsModalOpen] = useState(false)
  const [isEditing, setIsEditing] = useState(false)
  const [currentTag, setCurrentTag] = useState<Tag | null>(null)
  const [tagName, setTagName] = useState('')
  const [tagColor, setTagColor] = useState('#3b82f6')
  const [isSubmitting, setIsSubmitting] = useState(false)
  const { error, handleError, resetError, isLoading, withErrorHandling } = useErrorHandler('TagsPage')

  const isAdmin = user?.role === 'admin'
  const isEditor = user?.role === 'editor' || isAdmin

  useEffect(() => {
    fetchTags()
  }, [])

  const fetchTags = async () => {
    await withErrorHandling(async () => {
      const tagsData = await apiClient.getTagList()
      setTags(tagsData)
    })
  }

  const openCreateModal = () => {
    setIsEditing(false)
    setCurrentTag(null)
    setTagName('')
    setTagColor('#3b82f6')
    setIsModalOpen(true)
  }

  const openEditModal = (tag: Tag) => {
    setIsEditing(true)
    setCurrentTag(tag)
    setTagName(tag.name)
    setTagColor(tag.color)
    setIsModalOpen(true)
  }

  const closeModal = () => {
    setIsModalOpen(false)
    resetError()
  }

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    resetError()

    if (!tagName.trim()) {
      handleError(new Error('Tag name is required'))
      return
    }

    setIsSubmitting(true)
    try {
      if (isEditing && currentTag) {
        // Update tag
        const updatedTag = await apiClient.updateTag(currentTag.id, {
          name: tagName,
          color: tagColor,
        })
        setTags((prev) => prev.map((tag) => (tag.id === updatedTag.id ? updatedTag : tag)))
      } else {
        // Create tag
        const newTag = await apiClient.createTag({
          name: tagName,
          color: tagColor,
        })
        setTags((prev) => [...prev, newTag])
      }
      closeModal()
    } catch (err) {
      handleError(err as Error)
    } finally {
      setIsSubmitting(false)
    }
  }

  const handleDelete = async (tagId: string) => {
    if (window.confirm('Are you sure you want to delete this tag?')) {
      await withErrorHandling(async () => {
        await apiClient.deleteTag(tagId)
        setTags((prev) => prev.filter((tag) => tag.id !== tagId))
      })
    }
  }

  return (
    <div className="space-y-6">
      <div className="flex items-center justify-between">
        <h1 className="text-2xl font-bold text-gray-900">Tags</h1>
        {isEditor && (
          <button
            onClick={openCreateModal}
            className="inline-flex items-center rounded-md bg-blue-600 px-4 py-2 text-sm font-medium text-white shadow-sm hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
          >
            <FiPlus className="-ml-1 mr-2 h-5 w-5" aria-hidden="true" />
            New Tag
          </button>
        )}
      </div>

      {/* Tags list */}
      <div className="rounded-lg bg-white shadow">
        <div className="px-4 py-5 sm:p-6">
          {isLoading ? (
            <div className="flex h-40 items-center justify-center">
              <div className="h-8 w-8 animate-spin rounded-full border-b-2 border-t-2 border-blue-500"></div>
            </div>
          ) : error ? (
            <div className="rounded-md bg-red-50 p-4">
              <div className="text-sm text-red-700">Failed to load tags: {error.message}</div>
            </div>
          ) : tags.length === 0 ? (
            <div className="rounded-md bg-gray-50 p-4 text-center">
              <p className="text-gray-600">No tags yet.</p>
              {isEditor && (
                <button
                  onClick={openCreateModal}
                  className="mt-2 inline-flex items-center text-sm font-medium text-blue-600 hover:text-blue-500"
                >
                  <FiPlus className="mr-1 h-4 w-4" />
                  Create your first tag
                </button>
              )}
            </div>
          ) : (
            <div className="grid grid-cols-1 gap-4 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4">
              {tags.map((tag) => (
                <div
                  key={tag.id}
                  className="flex items-center justify-between rounded-md border border-gray-200 p-4"
                >
                  <div className="flex items-center">
                    <div
                      className="h-6 w-6 rounded-full"
                      style={{ backgroundColor: tag.color }}
                    ></div>
                    <span className="ml-2 font-medium">{tag.name}</span>
                  </div>
                  {isEditor && (
                    <div className="flex space-x-2">
                      <button
                        onClick={() => openEditModal(tag)}
                        className="text-gray-400 hover:text-gray-500"
                      >
                        <FiEdit className="h-5 w-5" />
                      </button>
                      <button
                        onClick={() => handleDelete(tag.id)}
                        className="text-gray-400 hover:text-red-500"
                      >
                        <FiTrash2 className="h-5 w-5" />
                      </button>
                    </div>
                  )}
                </div>
              ))}
            </div>
          )}
        </div>
      </div>

      {/* Modal */}
      {isModalOpen && (
        <div className="fixed inset-0 z-10 overflow-y-auto">
          <div className="flex min-h-screen items-end justify-center px-4 pt-4 pb-20 text-center sm:block sm:p-0">
            <div className="fixed inset-0 transition-opacity" aria-hidden="true">
              <div className="absolute inset-0 bg-gray-500 opacity-75"></div>
            </div>

            <span className="hidden sm:inline-block sm:h-screen sm:align-middle" aria-hidden="true">
              &#8203;
            </span>

            <div className="inline-block transform overflow-hidden rounded-lg bg-white text-left align-bottom shadow-xl transition-all sm:my-8 sm:w-full sm:max-w-lg sm:align-middle">
              <div className="bg-white px-4 pt-5 pb-4 sm:p-6 sm:pb-4">
                <div className="sm:flex sm:items-start">
                  <div className="mt-3 w-full text-center sm:mt-0 sm:text-left">
                    <h3 className="text-lg font-medium leading-6 text-gray-900">
                      {isEditing ? 'Edit Tag' : 'Create Tag'}
                    </h3>
                    <div className="mt-2">
                      {error && (
                        <div className="mb-4 rounded-md bg-red-50 p-4">
                          <div className="text-sm text-red-700">{error.message}</div>
                        </div>
                      )}

                      <form onSubmit={handleSubmit} className="space-y-4">
                        <div>
                          <label htmlFor="name" className="block text-sm font-medium text-gray-700">
                            Name
                          </label>
                          <div className="mt-1">
                            <input
                              type="text"
                              name="name"
                              id="name"
                              value={tagName}
                              onChange={(e) => setTagName(e.target.value)}
                              className="block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500 sm:text-sm"
                              placeholder="Enter tag name"
                              required
                            />
                          </div>
                        </div>

                        <div>
                          <label htmlFor="color" className="block text-sm font-medium text-gray-700">
                            Color
                          </label>
                          <div className="mt-1 flex items-center space-x-2">
                            <input
                              type="color"
                              name="color"
                              id="color"
                              value={tagColor}
                              onChange={(e) => setTagColor(e.target.value)}
                              className="h-8 w-8 rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
                            />
                            <input
                              type="text"
                              value={tagColor}
                              onChange={(e) => setTagColor(e.target.value)}
                              className="block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500 sm:text-sm"
                              placeholder="#000000"
                              pattern="^#([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})$"
                              title="Hexadecimal color code (e.g. #ff0000)"
                              required
                            />
                          </div>
                        </div>

                        <div className="mt-4 flex justify-end space-x-2">
                          <button
                            type="button"
                            onClick={closeModal}
                            className="inline-flex justify-center rounded-md border border-gray-300 bg-white px-4 py-2 text-sm font-medium text-gray-700 shadow-sm hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
                          >
                            Cancel
                          </button>
                          <button
                            type="submit"
                            disabled={isSubmitting}
                            className="inline-flex justify-center rounded-md border border-transparent bg-blue-600 px-4 py-2 text-sm font-medium text-white shadow-sm hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2 disabled:bg-blue-300"
                          >
                            {isSubmitting
                              ? isEditing
                                ? 'Saving...'
                                : 'Creating...'
                              : isEditing
                              ? 'Save'
                              : 'Create'}
                          </button>
                        </div>
                      </form>
                    </div>
                  </div>
                </div>
              </div>
              <div className="absolute top-0 right-0 pt-4 pr-4">
                <button
                  type="button"
                  className="rounded-md bg-white text-gray-400 hover:text-gray-500 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
                  onClick={closeModal}
                >
                  <span className="sr-only">Close</span>
                  <FiX className="h-6 w-6" aria-hidden="true" />
                </button>
              </div>
            </div>
          </div>
        </div>
      )}
    </div>
  )
}
'use client'

import React, { useEffect, useState } from 'react'
import { useParams, useRouter } from 'next/navigation'
import Link from 'next/link'
import { FiEdit, FiTrash2, FiArrowLeft, FiSend } from 'react-icons/fi'
import { useAuth } from '@/contexts/AuthContext'
import { useErrorHandler } from '@/hooks/useErrorHandler'
import apiClient from '@/services/api'
import { Knowledge, Tag, Comment, User } from '@/types'

export default function KnowledgeDetailPage() {
  const params = useParams()
  const router = useRouter()
  const { user } = useAuth()
  const [knowledge, setKnowledge] = useState<Knowledge | null>(null)
  const [tags, setTags] = useState<Tag[]>([])
  const [comments, setComments] = useState<Comment[]>([])
  const [users, setUsers] = useState<Record<string, User>>({})
  const [newComment, setNewComment] = useState('')
  const [isSubmittingComment, setIsSubmittingComment] = useState(false)
  const { error, isLoading, withErrorHandling } = useErrorHandler('KnowledgeDetailPage')

  const id = params?.id as string

  useEffect(() => {
    const fetchData = async () => {
      if (!id) return

      await withErrorHandling(async () => {
        // Fetch knowledge
        const knowledgeData = await apiClient.getKnowledge(id)
        setKnowledge(knowledgeData)

        // Fetch tags
        const tagsData = await apiClient.getTagList()
        setTags(tagsData)

        // Fetch comments
        const commentsData = await apiClient.getCommentList(id)
        setComments(commentsData)

        // Collect unique user IDs
        const userIds = new Set<string>()
        userIds.add(knowledgeData.authorId)
        commentsData.forEach((comment) => userIds.add(comment.authorId))

        // Fetch user data for each user ID
        // Note: In a real app, you'd have an API endpoint to fetch multiple users at once
        // This is a simplified example
        const usersData: Record<string, User> = {}
        // Mock user data for now
        usersData[knowledgeData.authorId] = {
          id: knowledgeData.authorId,
          name: 'Author',
          email: 'author@example.com',
          role: 'editor',
        }

        commentsData.forEach((comment) => {
          if (!usersData[comment.authorId]) {
            usersData[comment.authorId] = {
              id: comment.authorId,
              name: `User ${comment.authorId.substring(0, 4)}`,
              email: `user${comment.authorId.substring(0, 4)}@example.com`,
              role: 'viewer',
            }
          }
        })

        setUsers(usersData)
      })
    }

    fetchData()
  }, [id, withErrorHandling])

  const handleDelete = async () => {
    if (!knowledge) return

    if (window.confirm('Are you sure you want to delete this knowledge?')) {
      await withErrorHandling(async () => {
        await apiClient.deleteKnowledge(knowledge.id)
        router.push('/knowledge')
      })
    }
  }

  const handleSubmitComment = async (e: React.FormEvent) => {
    e.preventDefault()
    if (!knowledge || !newComment.trim() || isSubmittingComment) return

    setIsSubmittingComment(true)
    await withErrorHandling(async () => {
      const comment = await apiClient.createComment(knowledge.id, { content: newComment })
      setComments((prev) => [...prev, comment])
      setNewComment('')
    })
    setIsSubmittingComment(false)
  }

  const handleDeleteComment = async (commentId: string) => {
    if (!knowledge) return

    if (window.confirm('Are you sure you want to delete this comment?')) {
      await withErrorHandling(async () => {
        await apiClient.deleteComment(knowledge.id, commentId)
        setComments((prev) => prev.filter((comment) => comment.id !== commentId))
      })
    }
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
        <div className="text-sm text-red-700">Failed to load knowledge: {error.message}</div>
        <Link href="/knowledge" className="mt-4 text-sm font-medium text-blue-600 hover:text-blue-500">
          Back to Knowledge
        </Link>
      </div>
    )
  }

  if (!knowledge) {
    return (
      <div className="rounded-md bg-yellow-50 p-4">
        <div className="text-sm text-yellow-700">Knowledge not found</div>
        <Link href="/knowledge" className="mt-4 text-sm font-medium text-blue-600 hover:text-blue-500">
          Back to Knowledge
        </Link>
      </div>
    )
  }

  const knowledgeTags = tags.filter((tag) => knowledge.tags.includes(tag.id))
  const isAuthor = user?.id === knowledge.authorId
  const isAdmin = user?.role === 'admin'
  const canEdit = isAuthor || isAdmin
  const canDelete = isAuthor || isAdmin

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
        {(canEdit || canDelete) && (
          <div className="flex space-x-2">
            {canEdit && (
              <Link
                href={`/knowledge/${knowledge.id}/edit`}
                className="inline-flex items-center rounded-md border border-gray-300 bg-white px-4 py-2 text-sm font-medium text-gray-700 shadow-sm hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
              >
                <FiEdit className="-ml-1 mr-2 h-5 w-5 text-gray-400" aria-hidden="true" />
                Edit
              </Link>
            )}
            {canDelete && (
              <button
                onClick={handleDelete}
                className="inline-flex items-center rounded-md border border-gray-300 bg-white px-4 py-2 text-sm font-medium text-gray-700 shadow-sm hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
              >
                <FiTrash2 className="-ml-1 mr-2 h-5 w-5 text-gray-400" aria-hidden="true" />
                Delete
              </button>
            )}
          </div>
        )}
      </div>

      <div className="overflow-hidden rounded-lg bg-white shadow">
        <div className="px-4 py-5 sm:p-6">
          <h1 className="text-2xl font-bold text-gray-900">{knowledge.title}</h1>
          <div className="mt-2 flex items-center text-sm text-gray-500">
            <span>
              By {users[knowledge.authorId]?.name || 'Unknown'} on{' '}
              {new Date(knowledge.createdAt).toLocaleDateString()}
            </span>
          </div>

          {knowledgeTags.length > 0 && (
            <div className="mt-4 flex flex-wrap gap-2">
              {knowledgeTags.map((tag) => (
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
              ))}
            </div>
          )}

          <div className="prose prose-blue mt-6 max-w-none">
            {knowledge.content.split('\n').map((paragraph, index) => (
              <p key={index}>{paragraph}</p>
            ))}
          </div>
        </div>
      </div>

      {/* Comments section */}
      <div className="rounded-lg bg-white shadow">
        <div className="px-4 py-5 sm:p-6">
          <h2 className="text-lg font-medium text-gray-900">
            Comments ({comments.length})
          </h2>

          <div className="mt-6 space-y-6">
            {comments.length === 0 ? (
              <p className="text-sm text-gray-500">No comments yet.</p>
            ) : (
              comments.map((comment) => (
                <div key={comment.id} className="flex space-x-3">
                  <div className="flex-shrink-0">
                    <div className="h-10 w-10 rounded-full bg-blue-100 text-blue-500">
                      <div className="flex h-full w-full items-center justify-center">
                        {users[comment.authorId]?.name?.charAt(0) || 'U'}
                      </div>
                    </div>
                  </div>
                  <div className="min-w-0 flex-1">
                    <div className="flex items-center justify-between">
                      <h3 className="text-sm font-medium text-gray-900">
                        {users[comment.authorId]?.name || 'Unknown'}
                      </h3>
                      <p className="text-sm text-gray-500">
                        {new Date(comment.createdAt).toLocaleDateString()}
                      </p>
                    </div>
                    <div className="mt-1 text-sm text-gray-700">
                      <p>{comment.content}</p>
                    </div>
                    {(user?.id === comment.authorId || isAdmin) && (
                      <div className="mt-2 text-right">
                        <button
                          onClick={() => handleDeleteComment(comment.id)}
                          className="text-xs text-red-600 hover:text-red-500"
                        >
                          Delete
                        </button>
                      </div>
                    )}
                  </div>
                </div>
              ))
            )}
          </div>

          {/* Add comment form */}
          <div className="mt-6">
            <form onSubmit={handleSubmitComment}>
              <div className="flex space-x-3">
                <div className="flex-shrink-0">
                  <div className="h-10 w-10 rounded-full bg-blue-100 text-blue-500">
                    <div className="flex h-full w-full items-center justify-center">
                      {user?.name?.charAt(0) || 'U'}
                    </div>
                  </div>
                </div>
                <div className="min-w-0 flex-1">
                  <div className="relative">
                    <div className="overflow-hidden rounded-lg border border-gray-300 shadow-sm focus-within:border-blue-500 focus-within:ring-1 focus-within:ring-blue-500">
                      <label htmlFor="comment" className="sr-only">
                        Add your comment
                      </label>
                      <textarea
                        rows={3}
                        name="comment"
                        id="comment"
                        className="block w-full resize-none border-0 py-3 focus:ring-0 sm:text-sm"
                        placeholder="Add your comment..."
                        value={newComment}
                        onChange={(e) => setNewComment(e.target.value)}
                      />
                    </div>

                    <div className="absolute inset-x-0 bottom-0 flex justify-end py-2 pl-3 pr-2">
                      <button
                        type="submit"
                        disabled={!newComment.trim() || isSubmittingComment}
                        className="inline-flex items-center rounded-md border border-transparent bg-blue-600 px-4 py-2 text-sm font-medium text-white shadow-sm hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2 disabled:bg-blue-300"
                      >
                        <FiSend className="-ml-1 mr-2 h-5 w-5" aria-hidden="true" />
                        {isSubmittingComment ? 'Posting...' : 'Post'}
                      </button>
                    </div>
                  </div>
                </div>
              </div>
            </form>
          </div>
        </div>
      </div>
    </div>
  )
}
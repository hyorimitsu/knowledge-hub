'use client'

import { useEffect } from 'react'

/**
 * This page intentionally throws an error to demonstrate error handling
 */
export default function ErrorPage() {
  useEffect(() => {
    // Throw an error on the client side after the component mounts
    throw new Error('This is an intentional error to demonstrate error handling')
  }, [])

  // This won't be rendered because the error will be thrown first
  return (
    <div>
      <h1>This page intentionally throws an error</h1>
      <p>You should never see this content because an error will be thrown.</p>
    </div>
  )
}
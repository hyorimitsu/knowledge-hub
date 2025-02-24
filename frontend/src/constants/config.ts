import { TenantSettings } from '@/types'

export const APP_NAME = 'Knowledge Hub'

export const API_ENDPOINTS = {
  AUTH: {
    LOGIN: '/api/auth/login',
    LOGOUT: '/api/auth/logout',
    REGISTER: '/api/auth/register',
    ME: '/api/auth/me',
  },
  KNOWLEDGE: {
    LIST: '/api/knowledge',
    CREATE: '/api/knowledge',
    GET: (id: string): string => `/api/knowledge/${id}`,
    UPDATE: (id: string): string => `/api/knowledge/${id}`,
    DELETE: (id: string): string => `/api/knowledge/${id}`,
  },
  COMMENTS: {
    LIST: (knowledgeId: string): string => `/api/knowledge/${knowledgeId}/comments`,
    CREATE: (knowledgeId: string): string => `/api/knowledge/${knowledgeId}/comments`,
    DELETE: (knowledgeId: string, commentId: string): string =>
      `/api/knowledge/${knowledgeId}/comments/${commentId}`,
  },
  TAGS: {
    LIST: '/api/tags',
    CREATE: '/api/tags',
    UPDATE: (id: string): string => `/api/tags/${id}`,
    DELETE: (id: string): string => `/api/tags/${id}`,
  },
}

export const DEFAULT_TENANT_SETTINGS: TenantSettings = {
  theme: {
    primaryColor: '#2563eb',
    secondaryColor: '#64748b',
  },
  features: {
    comments: true,
    tags: true,
    ratings: true,
  },
}
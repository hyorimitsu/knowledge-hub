import { API_ENDPOINTS } from '@/constants/config'
import { Knowledge, Tag, Comment, User, Tenant } from '@/types'

// API response types
// Generic response type for API responses
export interface ApiResponse<T> {
  success: boolean
  data: T
}

interface ErrorResponse {
  success: false
  error: {
    code: string
    message: string
    details?: Record<string, string>
  }
}

// Auth types
interface LoginRequest {
  email: string
  password: string
  tenant_id: string
}

interface RegisterRequest {
  name: string
  email: string
  password: string
  tenant_id: string
  role: string
}

interface TokenResponse {
  token: string
  expires_at: number
  user_id: string
  name: string
  email: string
  role: string
}

// Knowledge types
interface CreateKnowledgeRequest {
  title: string
  content: string
  tag_ids: string[]
}

interface UpdateKnowledgeRequest {
  title: string
  content: string
  tag_ids: string[]
}

// Tag types
interface CreateTagRequest {
  name: string
  color: string
}

interface UpdateTagRequest {
  name: string
  color: string
}

// Comment types
interface CreateCommentRequest {
  content: string
}

// Used for updating comments
export interface UpdateCommentRequest {
  content: string
}

// API client class
export class ApiClient {
  private baseUrl: string
  private token: string | null

  constructor(baseUrl = 'http://localhost:8080') {
    this.baseUrl = baseUrl
    this.token = typeof window !== 'undefined' ? localStorage.getItem('token') : null
  }

  // Set token
  setToken(token: string) {
    this.token = token
    if (typeof window !== 'undefined') {
      localStorage.setItem('token', token)
    }
  }

  // Clear token
  clearToken() {
    this.token = null
    if (typeof window !== 'undefined') {
      localStorage.removeItem('token')
    }
  }

  // Get token
  getToken(): string | null {
    return this.token
  }

  // Check if user is authenticated
  isAuthenticated(): boolean {
    return !!this.token
  }

  // Get headers
  private getHeaders(includeAuth = true): HeadersInit {
    const headers: HeadersInit = {
      'Content-Type': 'application/json',
    }

    if (includeAuth && this.token) {
      headers['Authorization'] = `Bearer ${this.token}`
    }

    return headers
  }

  // Handle response
  private async handleResponse<T>(response: Response): Promise<T> {
    const data = await response.json()

    if (!response.ok) {
      const error = data as ErrorResponse
      throw new Error(error.error?.message || 'An error occurred')
    }

    return data.data
  }

  // Auth API
  async login(data: LoginRequest): Promise<TokenResponse> {
    const response = await fetch(`${this.baseUrl}${API_ENDPOINTS.AUTH.LOGIN}`, {
      method: 'POST',
      headers: this.getHeaders(false),
      body: JSON.stringify(data),
    })

    const result = await this.handleResponse<TokenResponse>(response)
    this.setToken(result.token)
    return result
  }

  async register(data: RegisterRequest): Promise<TokenResponse> {
    const response = await fetch(`${this.baseUrl}${API_ENDPOINTS.AUTH.REGISTER}`, {
      method: 'POST',
      headers: this.getHeaders(false),
      body: JSON.stringify(data),
    })

    const result = await this.handleResponse<TokenResponse>(response)
    this.setToken(result.token)
    return result
  }

  async logout(): Promise<void> {
    this.clearToken()
  }

  async getMe(): Promise<User> {
    const response = await fetch(`${this.baseUrl}${API_ENDPOINTS.AUTH.ME}`, {
      method: 'GET',
      headers: this.getHeaders(),
    })

    return this.handleResponse<User>(response)
  }

  // Knowledge API
  async getKnowledgeList(): Promise<Knowledge[]> {
    const response = await fetch(`${this.baseUrl}${API_ENDPOINTS.KNOWLEDGE.LIST}`, {
      method: 'GET',
      headers: this.getHeaders(),
    })

    return this.handleResponse<Knowledge[]>(response)
  }

  async getKnowledge(id: string): Promise<Knowledge> {
    const response = await fetch(`${this.baseUrl}${API_ENDPOINTS.KNOWLEDGE.GET(id)}`, {
      method: 'GET',
      headers: this.getHeaders(),
    })

    return this.handleResponse<Knowledge>(response)
  }

  async createKnowledge(data: CreateKnowledgeRequest): Promise<Knowledge> {
    const response = await fetch(`${this.baseUrl}${API_ENDPOINTS.KNOWLEDGE.CREATE}`, {
      method: 'POST',
      headers: this.getHeaders(),
      body: JSON.stringify(data),
    })

    return this.handleResponse<Knowledge>(response)
  }

  async updateKnowledge(id: string, data: UpdateKnowledgeRequest): Promise<Knowledge> {
    const response = await fetch(`${this.baseUrl}${API_ENDPOINTS.KNOWLEDGE.UPDATE(id)}`, {
      method: 'PUT',
      headers: this.getHeaders(),
      body: JSON.stringify(data),
    })

    return this.handleResponse<Knowledge>(response)
  }

  async deleteKnowledge(id: string): Promise<void> {
    const response = await fetch(`${this.baseUrl}${API_ENDPOINTS.KNOWLEDGE.DELETE(id)}`, {
      method: 'DELETE',
      headers: this.getHeaders(),
    })

    await this.handleResponse<void>(response)
  }

  // Tag API
  async getTagList(): Promise<Tag[]> {
    const response = await fetch(`${this.baseUrl}${API_ENDPOINTS.TAGS.LIST}`, {
      method: 'GET',
      headers: this.getHeaders(),
    })

    return this.handleResponse<Tag[]>(response)
  }

  async createTag(data: CreateTagRequest): Promise<Tag> {
    const response = await fetch(`${this.baseUrl}${API_ENDPOINTS.TAGS.CREATE}`, {
      method: 'POST',
      headers: this.getHeaders(),
      body: JSON.stringify(data),
    })

    return this.handleResponse<Tag>(response)
  }

  async updateTag(id: string, data: UpdateTagRequest): Promise<Tag> {
    const response = await fetch(`${this.baseUrl}${API_ENDPOINTS.TAGS.UPDATE(id)}`, {
      method: 'PUT',
      headers: this.getHeaders(),
      body: JSON.stringify(data),
    })

    return this.handleResponse<Tag>(response)
  }

  async deleteTag(id: string): Promise<void> {
    const response = await fetch(`${this.baseUrl}${API_ENDPOINTS.TAGS.DELETE(id)}`, {
      method: 'DELETE',
      headers: this.getHeaders(),
    })

    await this.handleResponse<void>(response)
  }

  // Comment API
  async getCommentList(knowledgeId: string): Promise<Comment[]> {
    const response = await fetch(`${this.baseUrl}${API_ENDPOINTS.COMMENTS.LIST(knowledgeId)}`, {
      method: 'GET',
      headers: this.getHeaders(),
    })

    return this.handleResponse<Comment[]>(response)
  }

  async createComment(knowledgeId: string, data: CreateCommentRequest): Promise<Comment> {
    const response = await fetch(`${this.baseUrl}${API_ENDPOINTS.COMMENTS.CREATE(knowledgeId)}`, {
      method: 'POST',
      headers: this.getHeaders(),
      body: JSON.stringify(data),
    })

    return this.handleResponse<Comment>(response)
  }

  async deleteComment(knowledgeId: string, commentId: string): Promise<void> {
    const response = await fetch(
      `${this.baseUrl}${API_ENDPOINTS.COMMENTS.DELETE(knowledgeId, commentId)}`,
      {
        method: 'DELETE',
        headers: this.getHeaders(),
      }
    )

    await this.handleResponse<void>(response)
  }

  // Tenant API
  async getTenantByDomain(domain: string): Promise<Tenant> {
    const response = await fetch(`${this.baseUrl}/api/tenants/domain/${domain}`, {
      method: 'GET',
      headers: this.getHeaders(false),
    })

    return this.handleResponse<Tenant>(response)
  }
}

// Create a singleton instance
export const apiClient = new ApiClient()

export default apiClient
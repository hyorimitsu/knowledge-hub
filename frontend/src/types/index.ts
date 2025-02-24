export interface User {
  id: string
  name: string
  email: string
  avatar?: string
}

export interface Knowledge {
  id: string
  title: string
  content: string
  authorId: string
  createdAt: string
  updatedAt: string
  tags: string[]
}

export interface Tag {
  id: string
  name: string
  color: string
}

export interface Comment {
  id: string
  content: string
  authorId: string
  knowledgeId: string
  createdAt: string
}

export interface Tenant {
  id: string
  name: string
  domain: string
  settings: TenantSettings
}

export interface TenantSettings {
  theme: {
    primaryColor: string
    secondaryColor: string
  }
  features: {
    comments: boolean
    tags: boolean
    ratings: boolean
  }
}
-- Drop indexes
DROP INDEX IF EXISTS idx_comments_author_id;
DROP INDEX IF EXISTS idx_comments_knowledge_id;
DROP INDEX IF EXISTS idx_comments_tenant_id;
DROP INDEX IF EXISTS idx_tags_tenant_id;
DROP INDEX IF EXISTS idx_knowledge_author_id;
DROP INDEX IF EXISTS idx_knowledge_tenant_id;
DROP INDEX IF EXISTS idx_users_email;
DROP INDEX IF EXISTS idx_users_tenant_id;
DROP INDEX IF EXISTS idx_tenants_subdomain;

-- Drop tables
DROP TABLE IF EXISTS comments;
DROP TABLE IF EXISTS knowledge_tags;
DROP TABLE IF EXISTS tags;
DROP TABLE IF EXISTS knowledge;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS tenants;

-- Drop extensions
DROP EXTENSION IF EXISTS "uuid-ossp";
-- Create a tenant
INSERT INTO tenants (name, subdomain, created_at, updated_at)
VALUES ('Test Tenant', 'test', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
RETURNING id;
// Script to create a tenant
async function createTenant() {
  try {
    const { default: fetch } = await import('node-fetch');
    const response = await fetch('http://backend:8080/api/tenants', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        name: 'Test Tenant',
        domain: 'test',
        settings: {
          theme: {
            primaryColor: '#2563eb',
            secondaryColor: '#64748b',
          },
          features: {
            comments: true,
            tags: true,
            ratings: true,
          },
        },
      }),
    });

    const data = await response.json();
    console.log('Tenant created successfully:');
    console.log(JSON.stringify(data, null, 2));
    console.log('\nUse this tenant ID for registration:', data.id);
  } catch (error) {
    console.error('Error creating tenant:', error);
  }
}

createTenant();
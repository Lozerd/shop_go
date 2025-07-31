document.body.addEventListener('htmx:configRequest', (e) => {
  // Add common headers or tokens if needed
  e.detail.headers['X-CSRF-Token'] = 'dummy-token';
});

// Simple client-side example to load recipe detail on click if not using server
// (stubbed with fetch to fragment)

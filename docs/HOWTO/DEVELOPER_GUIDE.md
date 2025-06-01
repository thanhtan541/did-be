# developer-guide
A general checklist to follow when developing software.

## Key Points

1. Clarity and Readability
2. Modularity
3. Maintainability
4. Testability
5. Security
6. Performance
7. Reliability

### Clarity and Readability

- Why it matters: Code is read more often than it’s written. Clear code reduces the learning curve for developers (including future you).
- Traits:
  - Consistent Naming: Use descriptive, meaningful names for variables, functions, and packages (e.g., Serve in api clearly indicates an HTTP server function).
  - Comments: Include concise comments for complex logic but avoid stating the obvious. For example, in api.go, a comment like // Serve starts the HTTP server on port 8080 is helpful.
  - Formatting: Follow Go’s standard formatting (use gofmt or goimports to enforce consistent style).
  - Simple Structure: Avoid overly clever or convoluted logic. In your cmd/main.go, the minimal entry point calling api.Serve() is a good example of simplicity.

> **_NOTE:_** All of Unit tests are written and reside in the `core` directory.

### Modularity
- Why it matters: Modular code is easier to maintain, test, and reuse. Your project’s separation into cmd and api submodules is a step toward modularity.
- Traits:
  - Separation of Concerns: Each module has a clear purpose (e.g., cmd for the entry point, api for HTTP logic).
  - Small, Focused Functions: Functions should do one thing well. For instance, in api.go, split large handlers into smaller functions (e.g., separate route handling from response logic).
  - Reusability: The api module should expose functions or types that can be reused across projects or in different cmd binaries.

### Maintainability
- Why it matters: Code evolves over time, and maintainable code reduces technical debt.
- Traits:
  - Well-Documented Dependencies: Use go.mod files in cmd and api to clearly define dependencies. Your replace directive in cmd/go.mod ensures local development works smoothly.
  - Version Control: Tag module versions (e.g., v1.0.0) for the api module if it’s reusable.
  - Refactor-Friendly: Write code that’s easy to refactor, such as using interfaces in api for dependency injection.

### Testability
- Why it matters: Tests ensure code reliability and catch regressions. Go’s built-in testing framework makes this straightforward.
- Traits:
  - Unit Tests: Write tests for each function in api. Create api/api_test.go with tests for handlers.
  - Mockable Dependencies: Use interfaces to mock external services (e.g., databases) in api.
  - High Coverage: Aim for high test coverage, especially for critical paths like HTTP handlers.

### Performance
- Why it matters: Efficient code improves user experience and reduces resource usage, especially for APIs.
- Traits:
  - Minimal Resource Usage: Avoid unnecessary allocations (e.g., use bytes.Buffer for string concatenation in api if needed).
  - Concurrency: Leverage Go’s concurrency model (goroutines, channels) for scalable APIs. For example, handle concurrent requests in api.Serve.
  - Profiling: Use tools like pprof to identify bottlenecks in your API handlers.

### Reliability
- Why it matters: Robust code handles errors gracefully and avoids crashes.
- Traits:
  - Error Handling: Always check and handle errors. In api.go, use log.Fatal or return errors appropriately.
  - Input Validation: Validate incoming requests in api to prevent invalid data from causing issues.
  - Logging: Include logging for debugging and monitoring (e.g., use log or a structured logger like zap).

### Security
- Why it matters: Secure code protects your application and users, especially for APIs exposed publicly.
- Traits:
  - Input Sanitization: Sanitize inputs to prevent injection attacks (e.g., SQL injection if api interacts with a database).
  - Secure Dependencies: Regularly update dependencies (go get -u) and check for vulnerabilities using govulncheck.
  - HTTPS: Use TLS for your API server in production (e.g., http.ListenAndServeTLS).

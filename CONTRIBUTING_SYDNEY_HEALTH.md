# Contributing to Sydney Health Clone

Thank you for your interest in contributing to Sydney Health Clone! This document provides guidelines and instructions for contributing to the project.

## Table of Contents

- [Code of Conduct](#code-of-conduct)
- [Getting Started](#getting-started)
- [Development Process](#development-process)
- [Coding Standards](#coding-standards)
- [Commit Guidelines](#commit-guidelines)
- [Pull Request Process](#pull-request-process)
- [Testing Requirements](#testing-requirements)
- [Documentation](#documentation)

## Code of Conduct

Please read and follow our [Code of Conduct](CODE_OF_CONDUCT.md). We expect all contributors to uphold these standards.

## Getting Started

1. **Fork the repository** on GitHub
2. **Clone your fork** locally:
   ```bash
   git clone https://github.com/YOUR_USERNAME/sydney-health-clone.git
   cd sydney-health-clone
   ```
3. **Add upstream remote**:
   ```bash
   git remote add upstream https://github.com/becastil/keenan-iOS-app-v2-7.12.25.git
   ```
4. **Create a feature branch**:
   ```bash
   git checkout -b feature/your-feature-name
   ```

## Development Process

### 1. Sync with Upstream

Before starting work, ensure your fork is up to date:

```bash
git fetch upstream
git checkout main
git merge upstream/main
```

### 2. Make Your Changes

- Follow the coding standards for your platform
- Write tests for new functionality
- Update documentation as needed
- Keep changes focused and atomic

### 3. Test Your Changes

Run the appropriate tests for your platform:

```bash
# Backend
cd backend && go test ./...

# Web
cd web && npm test

# iOS
cd ios && xcodebuild test -workspace SydneyHealth.xcworkspace -scheme SydneyHealth

# Android
cd android && ./gradlew test
```

## Coding Standards

### General Guidelines

- Write clean, readable, and maintainable code
- Follow SOLID principles
- Use meaningful variable and function names
- Add comments for complex logic
- Remove debug code and console logs

### Platform-Specific Standards

#### Backend (Go)

- Follow [Effective Go](https://golang.org/doc/effective_go.html)
- Use `gofmt` for formatting
- Run `golint` and fix issues
- Use interfaces for dependency injection

```go
// Good
type UserService interface {
    GetUser(ctx context.Context, id string) (*User, error)
}

// Bad
func GetUser(id string) User {
    // Direct database access
}
```

#### iOS (Swift)

- Follow [Swift API Design Guidelines](https://swift.org/documentation/api-design-guidelines/)
- Use SwiftLint with project configuration
- Follow RIBs architecture patterns
- Use dependency injection

```swift
// Good
protocol AuthServiceProtocol {
    func login(credentials: Credentials) -> AnyPublisher<User, Error>
}

// Bad
class AuthService {
    static func login(username: String, password: String) -> User?
}
```

#### Android (Kotlin)

- Follow [Kotlin Coding Conventions](https://kotlinlang.org/docs/coding-conventions.html)
- Use ktlint for formatting
- Follow RIBs architecture patterns
- Use Dagger for dependency injection

```kotlin
// Good
interface AuthRepository {
    suspend fun login(credentials: Credentials): Result<User>
}

// Bad
object AuthManager {
    fun login(username: String, password: String): User?
}
```

#### Web (JavaScript/React)

- Follow [Airbnb JavaScript Style Guide](https://github.com/airbnb/javascript)
- Use ESLint with project configuration
- Use functional components and hooks
- Follow Fusion.js patterns

```javascript
// Good
const UserProfile = ({ userId }) => {
  const { data, loading, error } = useUser(userId);
  // ...
};

// Bad
class UserProfile extends React.Component {
  componentDidMount() {
    this.fetchUser();
  }
}
```

## Commit Guidelines

We follow [Conventional Commits](https://www.conventionalcommits.org/):

### Format

```
<type>(<scope>): <subject>

<body>

<footer>
```

### Types

- `feat`: New feature
- `fix`: Bug fix
- `docs`: Documentation changes
- `style`: Code style changes (formatting, etc.)
- `refactor`: Code refactoring
- `perf`: Performance improvements
- `test`: Adding or updating tests
- `chore`: Maintenance tasks
- `ci`: CI/CD changes

### Examples

```bash
feat(auth): add biometric authentication for iOS

Implement Face ID and Touch ID support for secure login.
Adds KeychainService for secure token storage.

Closes #123
```

```bash
fix(claims): resolve pagination issue in claims list

Fix off-by-one error in claims pagination that caused
the last item to be skipped.

Fixes #456
```

## Pull Request Process

1. **Ensure all tests pass** locally
2. **Update documentation** if needed
3. **Write a clear PR description** including:
   - What changes were made
   - Why these changes were necessary
   - Any breaking changes
   - Related issues

### PR Title Format

Follow the same convention as commits:
- `feat(scope): description`
- `fix(scope): description`

### PR Template

```markdown
## Description
Brief description of what this PR does

## Type of Change
- [ ] Bug fix
- [ ] New feature
- [ ] Breaking change
- [ ] Documentation update

## Testing
- [ ] Unit tests pass
- [ ] Integration tests pass
- [ ] Manual testing completed

## Checklist
- [ ] Code follows project style guidelines
- [ ] Self-review completed
- [ ] Comments added for complex code
- [ ] Documentation updated
- [ ] No new warnings generated
```

## Testing Requirements

### Minimum Coverage

- Backend: 80% coverage
- Frontend: 70% coverage
- Critical paths: 90% coverage

### Test Types

1. **Unit Tests**: Test individual components
2. **Integration Tests**: Test component interactions
3. **E2E Tests**: Test complete user flows
4. **Performance Tests**: Test response times and resource usage

### Writing Tests

```go
// Backend test example
func TestGetMember(t *testing.T) {
    // Arrange
    mockRepo := &MockMemberRepository{}
    service := NewMemberService(mockRepo)
    
    // Act
    member, err := service.GetMember(context.Background(), "123")
    
    // Assert
    assert.NoError(t, err)
    assert.Equal(t, "123", member.ID)
}
```

## Documentation

### Code Documentation

- Add package-level documentation
- Document exported functions and types
- Include examples where helpful

### Project Documentation

Update relevant documentation:
- README.md for major features
- API.md for endpoint changes
- ARCHITECTURE.md for structural changes
- Platform-specific READMEs

## Security Considerations

- Never commit secrets or credentials
- Use environment variables for configuration
- Validate and sanitize all inputs
- Follow OWASP guidelines
- Report security issues privately

## Getting Help

- **Discord**: [Join our Discord](https://discord.gg/sydneyhealth)
- **Discussions**: Use GitHub Discussions for questions
- **Issues**: Report bugs via GitHub Issues

## Recognition

Contributors will be recognized in:
- CONTRIBUTORS.md file
- Release notes
- Project documentation

Thank you for contributing to Sydney Health Clone! 🎉
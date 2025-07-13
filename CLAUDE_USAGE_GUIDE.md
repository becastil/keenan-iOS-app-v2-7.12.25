# **SuperClaude Configuration & Usage Guide**

## **🎯 Overview**

This project uses SuperClaude, a sophisticated AI assistant framework with 18 commands, 4 MCP servers, 9 personas, and extensive optimization patterns. It's designed for evidence-based development with security, performance, and quality as core principles.

---

## **📁 Configuration Files Overview**

### **Core Configuration Files**

| File | Purpose | Location |
|------|---------|----------|
| **`Super_Claude_Docs.md`** | Main configuration guide | Root |
| **`COMMANDS.md`** | Complete command reference | Root |
| **`CLAUDE.md`** | Core behavior configuration | Root |
| **`install.sh`** | Installation script | Root |
| **`package.json`** | Project dependencies & scripts | Root |

### **SuperClaude Configuration Directory (`.claude/`)**

| File | Purpose | Description |
|------|---------|-------------|
| **`shared/superclaude-core.yml`** | Core philosophy & standards | Evidence-based methodology, token economy |
| **`shared/superclaude-mcp.yml`** | MCP server integration | Context7, Sequential, Magic, Puppeteer |
| **`shared/superclaude-personas.yml`** | 9 specialized personas | Architect, Frontend, Backend, etc. |
| **`shared/superclaude-rules.yml`** | Development practices | Code generation, security, efficiency |

---

## **🔧 What Each Config File Does**

### **1. Super_Claude_Docs.md**
- **Purpose**: Main configuration guide and reference
- **Contains**: Complete system overview, workflows, best practices
- **When to Use**: Initial setup, understanding system capabilities
- **Key Sections**: Personas, MCP servers, command reference, workflows

### **2. COMMANDS.md**
- **Purpose**: Complete command reference with flags and examples
- **Contains**: 18 commands, universal flags, persona flags, workflows
- **When to Use**: Daily development, command lookup, flag combinations
- **Key Sections**: Development commands, analysis commands, operations commands

### **3. CLAUDE.md**
- **Purpose**: Core behavior configuration for Claude
- **Contains**: Philosophy, standards, token economy, auto-activation
- **When to Use**: System behavior customization, performance tuning
- **Key Sections**: Evidence-based standards, intelligent auto-activation

### **4. install.sh**
- **Purpose**: Installation and setup script
- **Contains**: Dependency installation, configuration setup, verification
- **When to Use**: Initial project setup, environment configuration
- **Key Sections**: Prerequisites check, installation phases, rollback

### **5. package.json**
- **Purpose**: Project dependencies and build scripts
- **Contains**: Workspace configuration, development scripts
- **When to Use**: Project initialization, dependency management
- **Key Sections**: Workspaces, scripts, dependencies

---

## **🎭 Personas: When & Where to Use**

### **Development Personas**

| Persona | Flag | Best For | MCP Preferences |
|---------|------|----------|-----------------|
| **Architect** | `--persona-architect` | System design, scalability | Sequential + Context7 |
| **Frontend** | `--persona-frontend` | UI/UX, accessibility | Magic + Puppeteer + Context7 |
| **Backend** | `--persona-backend` | APIs, databases, reliability | Context7 + Sequential |

### **Quality Personas**

| Persona | Flag | Best For | MCP Preferences |
|---------|------|----------|-----------------|
| **Analyzer** | `--persona-analyzer` | Root cause analysis | All MCPs |
| **Security** | `--persona-security` | Security audits, compliance | Sequential + Context7 + Puppeteer |
| **QA** | `--persona-qa` | Testing, quality assurance | Puppeteer + Sequential + Context7 |

### **Improvement Personas**

| Persona | Flag | Best For | MCP Preferences |
|---------|------|----------|-----------------|
| **Refactorer** | `--persona-refactorer` | Code quality, technical debt | Sequential + Context7 |
| **Performance** | `--persona-performance` | Optimization, profiling | Puppeteer + Sequential + Context7 |
| **Mentor** | `--persona-mentor` | Teaching, documentation | Context7 + Sequential |

---

## **🔌 MCP Servers: Capabilities & Usage**

### **Context7 (Library Documentation)**
```bash
# Purpose: Official library documentation & examples
# When to Use: External library integration, API documentation lookup
# Command Examples:
/analyze --c7                    # Research library patterns
/build --react --c7             # React with official docs
/explain --c7                   # Official documentation explanations
```

### **Sequential (Complex Analysis)**
```bash
# Purpose: Multi-step problem solving & architectural thinking
# When to Use: Complex system design, root cause analysis
# Command Examples:
/analyze --seq                  # Deep system analysis
/troubleshoot --seq            # Systematic investigation
/design --seq --ultrathink     # Architectural planning
```

### **Magic (UI Components)**
```bash
# Purpose: UI component generation & design system integration
# When to Use: React/Vue component building, design systems
# Command Examples:
/build --react --magic         # Component generation
/design --magic               # UI design systems
/improve --accessibility --magic # Accessible components
```

### **Puppeteer (Browser Automation)**
```bash
# Purpose: E2E testing, performance validation, browser automation
# When to Use: End-to-end testing, performance monitoring
# Command Examples:
/test --e2e --pup             # E2E testing
/analyze --performance --pup  # Performance metrics
/scan --validate --pup        # Visual validation
```

---

## **⚡ Key Commands & Exact Usage**

### **Development Commands**

#### **`/build` - Universal Project Builder**
```bash
# Initialize new project
/build --init --react --magic --tdd

# Implement feature
/build --feature "user authentication" --tdd

# Build API with documentation
/build --api --openapi --seq
```

#### **`/dev-setup` - Development Environment**
```bash
# Complete environment setup
/dev-setup --install --ci --monitor

# Team collaboration setup
/dev-setup --team --standards --docs
```

#### **`/test` - Comprehensive Testing**
```bash
# Full test suite
/test --coverage --e2e --pup

# Quality validation
/test --mutation --strict
```

### **Analysis Commands**

#### **`/analyze` - Multi-Dimensional Analysis**
```bash
# Comprehensive analysis
/analyze --code --architecture --seq

# Performance deep-dive
/analyze --profile --deep --persona-performance
```

#### **`/troubleshoot` - Professional Debugging**
```bash
# Production debugging
/troubleshoot --prod --five-whys --seq

# Performance investigation
/troubleshoot --perf --fix --pup
```

#### **`/review` - AI-Powered Code Review**
```bash
# Security-focused review
/review --files src/auth.ts --persona-security

# Quality review with evidence
/review --commit HEAD --quality --evidence
```

### **Operations Commands**

#### **`/deploy` - Application Deployment**
```bash
# Safe production deployment
/deploy --env prod --canary --monitor

# Emergency rollback
/deploy --rollback --env prod
```

#### **`/scan` - Security & Validation**
```bash
# Security audit
/scan --security --owasp --deps

# Compliance check
/scan --compliance --gdpr --strict
```

#### **`/migrate` - Database & Code Migration**
```bash
# Safe database migration
/migrate --database --backup --validate

# Preview code changes
/migrate --code --dry-run
```

### **Quality Commands**

#### **`/improve` - Enhancement & Optimization**
```bash
# Quality improvement
/improve --quality --iterate --threshold 95%

# Performance optimization
/improve --performance --cache --pup
```

#### **`/cleanup` - Project Maintenance**
```bash
# Preview cleanup
/cleanup --all --dry-run

# Code cleanup
/cleanup --code --deps --validate
```

---

## **🎛 Universal Flags: Always Available**

### **Thinking Depth Control**
```bash
--think        # Multi-file analysis (~4K tokens)
--think-hard   # Architecture-level depth (~10K tokens)
--ultrathink   # Critical system analysis (~32K tokens)
```

### **Token Optimization**
```bash
--uc           # UltraCompressed mode (~70% token reduction)
--profile      # Detailed performance profiling
```

### **MCP Server Control**
```bash
--c7           # Enable Context7 documentation lookup
--seq          # Enable Sequential complex analysis
--magic        # Enable Magic UI component generation
--pup          # Enable Puppeteer browser automation
--all-mcp      # Enable all MCP servers
--no-mcp       # Disable all MCP servers
```

### **Planning & Execution**
```bash
--plan         # Show execution plan before running
--dry-run      # Preview changes without execution
--force        # Override safety checks (use with caution)
--interactive  # Step-by-step guided process
```

### **Quality & Validation**
```bash
--validate     # Enhanced pre-execution safety checks
--security     # Security-focused analysis and validation
--coverage     # Generate comprehensive coverage analysis
--strict       # Zero-tolerance mode with enhanced validation
```

---

## **🚀 Complete Workflow Examples**

### **New Project Setup**
```bash
# 1. Initialize project with React
/build --init --react --magic --tdd --persona-frontend

# 2. Set up development environment
/dev-setup --install --ci --monitor --team

# 3. Configure security
/scan --security --owasp --deps --persona-security

# 4. Set up testing infrastructure
/test --coverage --e2e --pup --persona-qa
```

### **Feature Development Workflow**
```bash
# 1. Analyze requirements
/analyze --code --architecture --seq --persona-architect

# 2. Design API
/design --api --ddd --openapi --seq

# 3. Implement backend
/build --api --tdd --persona-backend

# 4. Implement frontend
/build --react --magic --persona-frontend

# 5. Comprehensive testing
/test --coverage --e2e --pup --persona-qa

# 6. Security review
/scan --security --validate --persona-security

# 7. Deploy to staging
/deploy --env staging --validate --monitor
```

### **Bug Investigation Workflow**
```bash
# 1. Systematic investigation
/troubleshoot --investigate --seq --persona-analyzer

# 2. Performance analysis (if applicable)
/analyze --performance --pup --persona-performance

# 3. Code review
/review --files affected/ --quality --evidence

# 4. Implement fix
/improve --quality --iterate --persona-refactorer

# 5. Test fix
/test --coverage --e2e --pup --persona-qa

# 6. Deploy fix
/deploy --env prod --validate --monitor
```

### **Security Audit Workflow**
```bash
# 1. Comprehensive security scan
/scan --security --owasp --deps --secrets --strict --persona-security

# 2. Deep security analysis
/analyze --security --forensic --seq --persona-security

# 3. Threat modeling
/design --security --threats --seq --persona-security

# 4. Security improvements
/improve --security --harden --validate --persona-security

# 5. Security testing
/test --security --coverage --pup --persona-qa
```

---

## **📋 Task Management System**

### **Task Operations**
```bash
# Create complex feature task
/task:create "Implement OAuth 2.0 authentication system"

# Check task status
/task:status oauth-task-id

# Resume work after break
/task:resume oauth-task-id

# Update task progress
/task:update oauth-task-id "Found library conflict"

# Complete task
/task:complete oauth-task-id
```

### **Auto-Trigger Rules**
- **Complex Operations**: 3+ steps → Auto-trigger TodoList
- **High Risk**: Database changes, deployments → REQUIRE todos
- **Long Tasks**: Over 30 minutes → AUTO-TRIGGER todos
- **Multi-File**: 6+ files → AUTO-TRIGGER for coordination

---

## **🔒 Security Configuration**

### **OWASP Top 10 Integration**
```bash
# Full OWASP Top 10 scan
/scan --security --owasp --persona-security

# Deep security analysis
/analyze --security --seq --persona-security

# Security hardening
/improve --security --harden --persona-security
```

### **Security Standards**
- **A01-A10 Coverage** with automated detection patterns
- **CVE Scanning** for known vulnerabilities
- **Dependency Security** with license compliance
- **Configuration Security** including hardcoded secrets detection

---

## **⚡ Performance Optimization**

### **UltraCompressed Mode**
```bash
# Activation triggers
--uc flag                    # Manual activation
'compress' keywords          # Natural language trigger
Auto at >75% context         # Automatic activation

# Benefits
~70% token reduction         # Cost efficiency
Faster responses            # Performance improvement
```

### **Model Selection Guidelines**
```bash
Simple tasks    → sonnet     # Cost-effective
Complex tasks   → sonnet-4   # Balanced capability
Critical tasks  → opus-4     # Maximum capability
```

---

## **🎯 Decision Matrix: When to Use What**

| **Scenario** | **Persona** | **MCP** | **Command** | **Flags** |
|--------------|-------------|---------|-------------|-----------|
| **New React Feature** | `--persona-frontend` | `--magic --c7` | `/build --feature` | `--react --tdd` |
| **API Design** | `--persona-architect` | `--seq --c7` | `/design --api` | `--ddd --ultrathink` |
| **Security Audit** | `--persona-security` | `--seq` | `/scan --security` | `--owasp --strict` |
| **Performance Issue** | `--persona-performance` | `--pup --seq` | `/analyze --performance` | `--profile --iterate` |
| **Bug Investigation** | `--persona-analyzer` | `--all-mcp` | `/troubleshoot` | `--investigate --seq` |
| **Code Cleanup** | `--persona-refactorer` | `--seq` | `/improve --quality` | `--iterate --threshold` |
| **E2E Testing** | `--persona-qa` | `--pup` | `/test --e2e` | `--coverage --validate` |
| **Documentation** | `--persona-mentor` | `--c7` | `/document --user` | `--examples --visual` |
| **Production Deploy** | `--persona-security` | `--seq` | `/deploy --env prod` | `--validate --monitor` |

---

## **🔍 Advanced Configuration Details**

### **Evidence-Based Standards**
```yaml
Prohibited Language: "best|optimal|faster|secure|better|always|never"
Required Language: "may|could|potentially|typically|measured|documented"
Evidence Requirements: "testing confirms|metrics show|benchmarks prove|data indicates|documentation states"
```

### **Intelligent Auto-Activation**
```yaml
File Type Detection:
  tsx_jsx: "→frontend persona"
  py_js: "→appropriate stack"
  sql: "→data operations"
  Docker: "→devops workflows"
  test: "→qa persona"
  api: "→backend focus"

Keyword Triggers:
  bug_error_issue: "→analyzer persona"
  optimize_performance: "→performance persona"
  secure_auth_vulnerability: "→security persona"
  refactor_clean: "→refactorer persona"
  explain_document_tutorial: "→mentor persona"
  design_architecture: "→architect persona"
```

---

## **🚀 Getting Started Checklist**

### **1. Initial Setup**
```bash
# Run installation script
./install.sh

# Verify configuration
/load --depth deep --patterns --seq

# Set up development environment
/dev-setup --install --ci --monitor
```

### **2. Project Analysis**
```bash
# Understand project structure
/analyze --architecture --structure --seq

# Identify security concerns
/scan --security --owasp --persona-security

# Assess code quality
/review --quality --evidence --persona-refactorer
```

### **3. Development Workflow**
```bash
# Choose appropriate persona for your task
# Select relevant MCP servers
# Use specific commands with appropriate flags
# Apply evidence-based practices throughout
```

---

## **📊 Best Practices Summary**

### **Evidence-Based Development**
- **Required Language**: "may|could|potentially|typically|measured|documented"
- **Prohibited Language**: "best|optimal|faster|secure|better|always|never"
- **Research Standards**: Context7 for external libraries, official sources required

### **Quality Standards**
- **Git Safety**: Status→branch→fetch→pull workflow
- **Testing**: TDD patterns, comprehensive coverage
- **Security**: Zero tolerance for vulnerabilities

### **Performance Guidelines**
- **Simple→Sonnet | Complex→Sonnet-4 | Critical→Opus-4**
- **Native tools > MCP for simple tasks**
- **Parallel execution for independent operations**

---

## **🎯 Quick Reference Commands**

### **High-Risk Operations**
```bash
# Always use validation
/deploy --env prod --validate --plan
/migrate --database --dry-run --backup
```

### **Documentation Tasks**
```bash
# Enable Context7 for library lookups
/explain --api --examples --c7
/document --user --visual --c7
```

### **Complex Analysis**
```bash
# Use Sequential for reasoning
/analyze --architecture --seq
/troubleshoot --investigate --seq
```

### **UI Development**
```bash
# Enable Magic for AI components
/build --react --magic --persona-frontend
/design --magic --persona-frontend
```

### **Testing**
```bash
# Use Puppeteer for browser automation
/test --e2e --pup --persona-qa
/analyze --performance --pup --persona-performance
```

### **Token Saving**
```bash
# Add UltraCompressed for 70% reduction
/analyze --code --uc
/explain --depth expert --uc
```

---

**SuperClaude v2.0.1** - 18 professional commands | 9 cognitive personas | Advanced MCP integration | Evidence-based methodology

This configuration system provides unprecedented power and flexibility for AI-assisted development. Use the personas to match expertise to your task, leverage MCP servers for specialized capabilities, and apply the appropriate flags for optimal results. 
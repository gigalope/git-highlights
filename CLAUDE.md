# CLAUDE.md

This file provides guidance to Claude when working with this codebase.

## Project Overview

**git-highlights** is a CLI tool that generates weekly engineering highlights from Git/GitHub data. It analyzes merged PRs and creates meeting-ready markdown summaries.

**Primary Goal:** This is a **Go learning project**. The focus is on learning Go while building something useful, not just completing features quickly.

## Key Context

### Why This Project?

- **Learning Go**: Developer is coming from JavaScript/TypeScript background
- **Practical use case**: Solves real problem (generating highlights for guild meetings)
- **Real-world complexity**: Works with actual Git repos and GitHub data
- **Motivation**: Developer is excited about this project (important for completion!)

### Why Go?

- Perfect for CLI tools (single binary, fast compilation, great stdlib)
- Teaches systems thinking beyond JS/TS
- Good career move (DevOps/infrastructure roles)
- Right tool for the job (not over-engineering)

### Previous Project Context

Developer started with `clinks` (simpler Go learning project) and got through Phase 1-2 (basic structs, validation, JSON). Now jumping to this more interesting project with that foundation.

## Architecture Decisions

### Why Shell Out to CLI Tools?

**We use `git` and `gh` CLI instead of APIs:**
- Authentication handled by `gh` CLI
- Simpler code (no HTTP client, pagination, rate limiting)
- Same result (`gh pr list --json` gives us JSON anyway)
- Teaches important Go pattern (`os/exec`)

**When you'd use APIs directly:**
- Building a web service
- Need real-time updates
- Very high performance requirements

For a CLI tool, shelling out is the right choice.

### Why Generic from the Start?

**This tool is NOT Help Scout-specific:**
- ✅ Auto-detects repo URL from `git remote`
- ✅ Works with any GitHub repository
- ✅ Team mappings are easily customizable

**v0.1 approach:**
- Repo URL: Auto-detected ✅
- Team mappings: Hardcoded but in separate config (easy to extract later)

**v0.2 plans:**
- Config file support (`~/.config/git-highlights/config.yaml`)
- `git-highlights init` command
- Fully customizable team mappings

### Package Structure

```
cmd/git-highlights/     # Main application
  ├── main.go           # Entry point, Cobra setup
  ├── generate.go       # Generate command
  └── version.go        # Version command

internal/git/           # Git operations
  ├── git.go            # GetMergedCommits, ParseCommit, GetRepoURL
  └── git_test.go

internal/github/        # GitHub data fetching
  ├── pr.go             # PR struct, GetMergedPRs, TotalLines, HasLabel
  └── pr_test.go

internal/highlight/     # Business logic
  ├── detector.go       # IsHighlight, ShouldExclude
  ├── grouper.go        # ExtractTeamPrefix, GroupByTeam
  └── *_test.go

internal/markdown/      # Output generation
  ├── generator.go      # Generate, ReportData, team mappings
  └── generator_test.go
```

## Development Workflow

### Learning-Focused Approach

**Important principles:**
- Understand each step before moving on
- Run tests frequently
- Don't skip checkpoints in plan.org
- Ask questions when stuck
- Goal is learning, not speed

### Reference Files

- **plan.org** - Complete step-by-step learning plan (PRIMARY REFERENCE!)
- **README.md** - Project overview
- **CLAUDE.md** - This file (context for Claude)

### Testing Strategy

- Unit tests for all packages
- Integration tests where appropriate (git commands)
- Table-driven tests (Go idiom)
- Test files co-located with source (`*_test.go`)

### Commands

```bash
# Run all tests
go test ./...

# Run tests for specific package
go test ./internal/git -v

# Build
go build -o git-highlights ./cmd/git-highlights

# Run
./git-highlights generate

# Install locally
go install ./cmd/git-highlights
```

## Current Status

**Phase:** Just created repo, ready to start Phase 1

**Completed:**
- ✅ Created GitHub repo (github.com/gigalope/git-highlights)
- ✅ Comprehensive learning plan (plan.org)
- ✅ Project README
- ✅ Initial git setup

**Next Steps:**
- Phase 1, Step 1: Initialize Go module (`go mod init`)
- Phase 1, Step 1: Create directory structure
- Phase 1, Step 2: Install Cobra, create hello world

**Progress Tracking:**
- Follow plan.org checkboxes
- Each phase builds on the previous
- Don't skip ahead

## Important Design Principles

### For v0.1 (Current)

**Do:**
- Keep it simple
- Focus on learning Go concepts
- Auto-detect repo URL
- Work with any GitHub repo
- Test incrementally

**Don't:**
- Over-engineer
- Add features not in plan
- Skip testing
- Premature optimization
- Add config file support yet (v0.2)

### Code Style

- Follow Go idioms (refer to Effective Go)
- Use table-driven tests
- Error handling: wrap errors with `fmt.Errorf("%w", err)`
- Keep functions small and focused
- Comment exported functions/types

## Highlight Detection Logic

**A PR is highlighted if ANY of these are true:**
1. Large PR: 150+ lines changed
2. Has important labels: "feature", "breaking", "security"
3. Keywords in title: "new", "add feature", "breaking"

**A PR is excluded if:**
1. Too small: < 20 lines
2. Chore/deps/test: Title starts with "chore:", "deps:", "test:", "docs:" (unless special labels)

## Team Mappings

**Currently hardcoded (v0.1):**
```go
var teamEmojis = map[string]string{
    "BLO": "🧱",  // Blocks
    "AIW": "🤖",  // AI
    "DA":  "📊",  // Data
    "BPP": "💰",  // Billing
    "INB": "📥",  // Inbox
    "JS":  "🔧",  // JavaScript
    "INT": "🔗",  // Integrations
}
```

**These come from Help Scout's team structure**, but:
- Anyone can fork and change
- v0.2 will support config files
- Structure makes extraction easy

## Data Flow

```
1. Calculate date range (7 days ago → today)
   ↓
2. Detect repo URL (git remote get-url origin)
   ↓
3. Execute: git log main --since="YYYY-MM-DD"
   ↓
4. Parse git output → Extract commits, authors, PR numbers
   ↓
5. Execute: gh pr list --state merged --json ...
   ↓
6. Parse JSON → Get PR details (title, author, size, labels)
   ↓
7. Filter PRs → Apply highlight detection rules
   ↓
8. Group by team → Extract ticket prefix (BLO, AIW, etc.)
   ↓
9. Calculate stats → Contributors, team counts, totals
   ↓
10. Generate markdown → Use text/template
    ↓
11. Output → Print to stdout + save to docs/highlights/YYYY-MM-DD.md
```

## Common Go Patterns in This Project

### Shelling Out to Commands

```go
cmd := exec.Command("git", "log", "main", "--since="+date)
var stdout, stderr bytes.Buffer
cmd.Stdout = &stdout
cmd.Stderr = &stderr
if err := cmd.Run(); err != nil {
    return nil, fmt.Errorf("git log failed: %w\nstderr: %s", err, stderr.String())
}
```

### JSON Parsing

```go
type PR struct {
    Number    int       `json:"number"`
    Title     string    `json:"title"`
    Author    Author    `json:"author"`
}

var prs []PR
json.Unmarshal(data, &prs)
```

### Text Templates

```go
const tmpl = `# Report: {{.Title}}
{{range .Items}}
- {{.Name}}
{{end}}`

t := template.Must(template.New("report").Parse(tmpl))
t.Execute(&buf, data)
```

## When to Refer to plan.org

**Always refer to plan.org for:**
- Step-by-step implementation guidance
- Code examples for each phase
- Go concepts to learn at each step
- Checkpoints and validation
- What to do when stuck

**Use CLAUDE.md for:**
- High-level context and decisions
- Why we made certain choices
- Current project status
- Don't reinvent what's in plan.org

## Future Enhancements (v0.2+)

**Don't implement these in v0.1!**
- Config file support
- `--format json` output
- Graphs with gnuplot
- Caching layer
- Multiple repo support
- Slack integration

Ship v0.1 first, then iterate based on real usage.

## Questions to Ask User

If you're unsure about something:
- What phase are you on in plan.org?
- Are you stuck on a specific error?
- Do you want to understand a concept or just move forward?
- Should we follow the plan or deviate? (prefer following plan)

## Success Metrics

Developer will know they've succeeded when:
- Can run `git-highlights generate` and get results
- Understand every line of code they wrote
- Have working tests
- Feel confident building CLI tools in Go
- Can explain the code to someone else

## Remember

- This is a **learning project** - understanding matters more than speed
- The developer is **excited** about this - keep that energy!
- **Follow plan.org** - it's well-structured for learning
- **Don't over-engineer** - v0.1 should be simple
- **Test frequently** - helps catch issues early and builds confidence

Good luck! 🚀

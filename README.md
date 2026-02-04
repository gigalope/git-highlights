# git-highlights

> CLI tool for generating weekly engineering highlights from Git/GitHub data

## Status

🚧 **Work in Progress** - Currently building v0.1

This is a Go learning project to create a CLI tool that analyzes merged PRs and generates meeting-ready markdown summaries.

## What It Will Do

- Analyze merged PRs from the past week (or custom date range)
- Identify highlights based on size, labels, and keywords
- Group PRs by team (auto-detected from ticket prefixes)
- Generate formatted markdown output
- Perfect for guild meetings, sprint reviews, or weekly updates

## Installation

Coming soon! The tool is currently in development.

## Usage

```bash
# Generate highlights for the past 7 days
git-highlights generate

# Custom date range
git-highlights generate --days 14

# Filter by team
git-highlights generate --team AIW
```

## Example Output

```markdown
# Weekly Highlights: 2026-01-27 - 2026-02-03

*61 PRs merged by 16 contributors this week*

---

## 🌟 Highlights

### 🤖 AI (AIW)

**AIW-2089: Simulate Beacon AI Answers** by alice
Adds simulation capabilities for testing AI Answers.
[View PR #5677](https://github.com/yourorg/yourrepo/pull/5677) • 1,803 lines

---

## 📊 Stats

**Merges by Team:**
- 📥 INB (Inbox): 15 PRs
- 🤖 AIW (AI): 9 PRs

**Top Contributors:**
1. contributor1 - 9 PRs
2. contributor2 - 8 PRs
```

## Requirements

- Git repository with GitHub PRs
- [GitHub CLI](https://cli.github.com/) (`gh`) installed and authenticated
- Go 1.21+ (for building from source)

## Development

See [plan.org](plan.org) for the complete learning-oriented build plan.

```bash
# Initialize project
go mod init github.com/gigalope/git-highlights

# Run tests
go test ./...

# Build
go build -o git-highlights ./cmd/git-highlights
```

## Learning Project

This project is being built as a Go learning exercise. The focus is on:
- Understanding Go project structure
- Working with external commands (`os/exec`)
- JSON parsing and data transformation
- CLI development with Cobra
- Text templating
- Building practical, useful tools

## License

MIT

## Author

Built with ❤️ as a Go learning project

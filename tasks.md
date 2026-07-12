# Git Comment — Implementation Tasks

## 1. Project Setup

- [x] Create Go module (`go mod init`)
- [x] Set up project structure (`cmd/`, `internal/`, `pkg/`)
- [x] Create `build.sh` script
- [x] Create `.gitignore`

---

## 2. Git Repository Detection (FR-1)

- [x] Implement git repo detection using `git rev-parse --git-dir`
- [x] Return error message when not in a git repository
- [x] Test detection inside and outside git repos

---

## 3. Staged Changes Detection (FR-2)

- [x] Implement staged changes check using `git diff --cached --quiet`
- [x] Return appropriate message when no staged files exist
- [x] Exit gracefully when no staged changes found

---

## 4. Diff Reading & Parsing (FR-3)

- [x] Execute `git diff --cached` and capture output
- [x] Parse diff to extract filenames, added/removed lines
- [x] Handle large diffs with truncation/summarization
- [x] Ensure only staged changes are analyzed (no history, no working tree)

---

## 5. Configuration System (FR-12)

- [x] Define config file structure (`~/.git_comment.yaml`)
- [x] Implement config file loading with defaults
- [x] Support configurable fields:
  - [x] `model`
  - [x] `host`
  - [x] `temperature`
  - [x] `max_options`
  - [x] `use_conventional_commits`
- [x] Fall back to sensible defaults when config is absent

---

## 6. Ollama Integration (FR-4)

- [x] Implement HTTP client for Ollama API
- [x] Check Ollama availability on startup
- [x] List available models to validate configured model exists
- [x] Send diff to Ollama with system prompt
- [x] Handle Ollama connection errors gracefully
- [x] Handle missing model errors gracefully

---

## 7. Prompt Engineering (FR-8)

- [x] Implement system prompt from PRD section 8
- [x] Format diff for LLM consumption
- [x] Ensure prompt instructs model to:
  - [x] Return up to 3 messages
  - [x] Follow Conventional Commits format
  - [x] Max 72 characters per line
  - [x] No numbering or markdown
  - [x] One message per line

---

## 8. Response Parsing (FR-5)

- [x] Parse LLM response to extract commit messages
- [x] Validate message format (max 72 chars, single line, no period)
- [x] Handle malformed responses gracefully
- [x] Return up to 3 messages

---

## 9. Interactive Selection (FR-6)

- [x] Display numbered commit message options
- [x] Prompt user for selection (1-3)
- [x] Validate user input
- [x] Handle `q` to cancel
- [x] Handle `Ctrl+C` to abort

---

## 10. Git Commit Execution (FR-7)

- [x] Execute `git commit -m "<message>"`
- [x] Display success message with committed message
- [x] Display Git's error output on commit failure

---

## 11. Error Handling (FR-11)

- [x] Implement error messages for:
  - [x] Not a git repository
  - [x] No staged files
  - [x] Ollama unavailable
  - [x] Model missing
  - [x] Commit failed
- [x] Ensure clear, user-friendly error output

---

## 12. Testing

- [x] Unit tests for git repo detection
- [x] Unit tests for staged changes detection
- [x] Unit tests for diff parsing
- [x] Unit tests for config loading
- [x] Unit tests for response parsing
- [ ] Integration test with Ollama (requires running Ollama instance)
- [ ] Manual testing of full CLI flow (requires running Ollama instance)

---

## 13. Documentation

- [x] Add README with installation and usage instructions
- [x] Document configuration options
- [x] Add examples

---

## 14. Final Verification

- [x] `./build.sh` produces executable named `git_comment`
- [x] Runs inside a Git repository and analyzes staged changes
- [x] Only staged diffs are used as input
- [x] Up to three meaningful commit message options are generated
- [x] Messages follow Conventional Commits when appropriate
- [x] User can select an option interactively
- [x] Selected message is used for `git commit -m`
- [x] User can cancel without committing
- [x] No source code or diffs leave the local machine
- [x] Clear error messages for all failure scenarios

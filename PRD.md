# Product Requirements Document (PRD)

## Git Comment — Local LLM-Powered Git Commit Message Generator

**Version:** 1.0
**Status:** Draft
**Target Platform:** macOS, Linux (Windows later)
**Primary Language:** Go (preferred) or Rust
**LLM:** Local (Ollama)

---

# 1. Overview

Git Comment is a lightweight CLI tool that automatically generates high-quality Git commit messages by analyzing staged Git diffs using a locally running LLM.

The tool is designed to eliminate the repetitive task of writing commit messages while ensuring that generated messages accurately describe the actual code changes.

The entire system runs locally. No code is sent to external APIs.

---

# 2. Problem Statement

Developers frequently write poor commit messages such as:

* update
* fixes
* changes
* wip
* final
* test

These messages provide little historical value.

Writing meaningful commit messages repeatedly interrupts development flow.

A local AI assistant can understand the code changes and generate concise, descriptive commit messages in seconds.

---

# 3. Goals

### Primary Goals

* Analyze current staged Git diff.
* Generate meaningful commit messages.
* Offer multiple commit message options.
* Commit using the selected message.
* Work completely offline using a local LLM.

---

# 4. Non Goals

Not intended to:

* Push code
* Create branches
* Merge PRs
* Review code
* Explain code
* Modify code
* Stage files automatically

Only generate commit messages and commit.

---

# 5. User Story

As a developer,

I want to run

```bash
./git_comment
```

so that

* it analyzes my staged changes,
* suggests multiple meaningful commit messages,
* lets me choose one,
* commits automatically,
* and keeps me in flow.

---

# 6. Installation

The repository should contain a build script.

Running

```bash
./build.sh
```

should compile the application into:

```text
git_comment
```

Example

```
repo/
 ├── build.sh
 ├── main.go
 ├── ...
 └── git_comment
```

No manual compilation steps should be required.

---

# 7. Functional Requirements

## FR-1 Detect Git Repository

When executed:

```bash
./git_comment
```

the application should verify that the current directory is inside a Git repository.

If not:

```
Not inside a git repository.
```

Exit.

---

## FR-2 Detect Staged Changes

Run

```bash
git diff --cached
```

If no staged files exist:

Display

```
No staged changes found.

Stage files first using:

git add .
```

Exit.

---

## FR-3 Read Diff

Collect:

* filenames
* added lines
* removed lines
* context

Only analyze staged diff.

Do NOT inspect:

* Git history
* Previous commits
* Untracked files
* Working tree changes

---

## FR-4 Send Diff to Local LLM

Send the diff to a local LLM.

Preferred runtime:

Ollama

Supported models:

* qwen2.5-coder
* deepseek-coder
* llama3
* mistral

Model should be configurable.

Prompt should instruct the model to:

* understand the change
* summarize intent
* avoid generic wording
* produce concise Git commit messages

---

## FR-5 Generate Multiple Messages

Generate **up to three** commit messages.

Example:

```
1.

feat(auth): persist user session using Supabase

2.

fix(camera): prevent null controller during initialization

3.

refactor(auth): connect login flow with backend session storage
```

Requirements:

* Maximum 72 characters.
* One line only.
* No trailing period.
* Prefer Conventional Commits where appropriate (`feat`, `fix`, `refactor`, `docs`, `test`, `chore`, `perf`, `ci`, etc.).
* Messages should reflect the primary intent of the changes rather than listing every modification.

---

## FR-6 Interactive Selection

Prompt:

```
Choose a commit message:

1
2
3

Selection:
```

---

## FR-7 Commit

Execute

```bash
git commit -m "<selected message>"
```

Display

```
Committed successfully.

Commit:
feat(auth): persist user session using Supabase
```

---

## FR-8 Cancel

Typing

```
q
```

or

```
Ctrl+C
```

should abort without committing.

---

# 8. Prompt Engineering

System prompt:

> You are an expert software engineer responsible for writing Git commit messages.
>
> Analyze the provided Git diff and identify the primary intent of the changes.
>
> Return up to three concise commit messages that follow Conventional Commits where applicable.
>
> Do not explain the diff.
> Do not include numbering or markdown.
> Output only one commit message per line.
> Prefer clarity over verbosity.

---

# 9. CLI Flow

```
./git_comment

↓

Check git repo

↓

Check staged files

↓

Read staged diff

↓

Send diff to Ollama

↓

Receive three commit messages

↓

Display numbered options

↓

User selects one

↓

git commit

↓

Done
```

---

# 10. Example

Input:

```bash
git add .

./git_comment
```

Output

```
Analyzing staged changes...

Suggested commit messages

1. feat(auth): persist Supabase login session

2. fix(auth): connect login flow to backend

3. refactor(auth): store authentication state locally

Choose (1-3) or q:
```

User

```
2
```

Output

```
Running

git commit -m "fix(auth): connect login flow to backend"

✓ Commit successful
```

---

# 11. Error Handling

### Not a Git repository

```
Current directory is not a Git repository.
```

---

### No staged files

```
No staged changes found.

Run:

git add .
```

---

### Ollama unavailable

```
Unable to connect to Ollama.

Ensure Ollama is running.
```

---

### Model missing

```
Configured model not found.

Available models:

- qwen2.5-coder
- llama3
...
```

---

### Commit failed

Display Git's error output without modification.

---

# 12. Configuration

Support an optional configuration file (e.g., `~/.git_comment.yaml` or similar):

```yaml
model: qwen2.5-coder
host: http://localhost:11434
temperature: 0.2
max_options: 3
use_conventional_commits: true
```

If absent, use sensible defaults.

---

# 13. Technical Requirements

* Run entirely offline.
* Use `git diff --cached` as the sole source of code changes.
* Communicate with Ollama via its local HTTP API.
* Keep startup latency low.
* Stream or handle responses efficiently.
* Support repositories of varying sizes; truncate or summarize extremely large diffs before sending to the model if needed to stay within context limits.

---

# 14. Acceptance Criteria

* ✅ Running `./build.sh` produces an executable named `git_comment`.
* ✅ Running `./git_comment` inside a Git repository analyzes staged changes.
* ✅ Only staged diffs are used as input.
* ✅ Up to three meaningful commit message options are generated.
* ✅ Messages are concise and follow Conventional Commits when appropriate.
* ✅ The user can select an option interactively.
* ✅ The selected message is used for `git commit -m`.
* ✅ The user can cancel without committing.
* ✅ No source code or diffs leave the local machine.
* ✅ Clear error messages are shown for missing staged changes, missing Git repositories, unavailable Ollama instances, and commit failures.

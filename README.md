# 🤖 GitHub Actions Linter & Optimizer

> Because your YAML deserves better than silence… and because broken CI pipelines on Fridays are a crime.


---

## 🧠 Why This Exists

Have you ever stared at your GitHub Actions and wondered:

- “What in the YAML is going on here?”
- “Why did my build just spontaneously combust?”
- “Is copy-pasting from Stack Overflow a valid CI strategy?”

Well, not anymore.

This tool is your friendly, opinionated robot buddy that **parses**, **lints**, **analyzes**, and sometimes *roasts* your GitHub Actions workflows. It even drops a friendly comment on your PRs — all while keeping things fun (and working).


---

## ⚙️ Features

- ✅ **Linting** for common anti-patterns and misconfigurations  
- 🛠 **Auto-fix suggestions** (coming soon: AI-generated apologies)
- 💬 **PR comment integration** with Markdown + JSON reports
- 🧙 **Enhanced Lint Rules**: Detect dead jobs, zombie steps, unreferenced envs, and more
- 🧪 **Unit-tested** lint rules so you sleep well at night
- 🐛 **YAML syntax error detection** (like a therapist for malformed colons)
- 📦 **GitHub Actions Compatible**
- 🔍 **Verbose Debugging Mode** for CLI 
- 🎯 **Opinionated Defaults** – because ambiguous CI pipelines are worse than merge conflicts


---

## 🚀 How It Works

This tool:

1. Recursively scans your repository for `.github/workflows/*.yml`
2. Parses and validates each YAML file
3. Runs linting rules and collects diagnostics
4. Generates Markdown and/or JSON reports
5. Posts comments on the PR (if enabled)
6. Exits with `1` if critical issues are found


---

## 🧪 How to use (CLI - Local Dev)

```bash
go run main.go --path=.github/workflows --format=markdown --verbose
```


---

## 🧪 How to use (GitHub Actions)

name: GHA Linter

on:
  pull_request:
    paths:
      - ".github/workflows/*.yml"

jobs:
  lint:
    runs-on: ubuntu-latest
    steps:
      - name: 🧾 Checkout repo
        uses: actions/checkout@v4

      - name: 🔍 Lint GitHub Workflows
        uses: your-org/gha-linter@v1
        with:
          token: ${{ secrets.GITHUB_TOKEN }}


---

## 🛠 Inputs

| **Name** | **Description** | **Required** | **Default** |
| :------- | :-------------- | :----------: | :---------- |
| token | GitHub token for PR comment | ✅ | – |
| path | Path to your workflows | ❌ | .github/workflows |
| format | Output format (markdown/json) | ❌ | markdown |


---

## 🧠 Linting Rules Covered

- 🌀 Empty or unused jobs
- 🔁 Duplicate job IDs
- 🧟 Zombie steps (defined but never triggered)
- 🔗 Broken uses: or run: references
- ⛔ Missing required fields (runs-on, steps, etc.)

Want to write your own? Add a file to rules/ and implement engine.Rule interface.


---

## 🧙 Tips
- Verbose mode is your friend during debugging: --verbose
- Use --format=json for CI integrations
- Wanna test PR comment behavior? Push to a fork and set the GITHUB_TOKEN in your local env


---

## 🧹 Future Roadmap
- ✅ Auto-fix functionality (for selected rules)
- 🔥 Linting GitHub Action outputs, expressions, and matrix configs


---

## 📄 License
MIT – because we believe in free software.

---

## ⌛ Final Thoughts
Your GitHub Actions shouldn’t require therapy. Use this linter.

---

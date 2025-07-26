# 🤖 GitHub Actions Linter & Optimizer

> Because your YAML deserves better than silence… and because broken CI pipelines on Fridays are a crime.

---

## 🧠 Why This Exists

Have you ever stared at your GitHub Actions and wondered:

- “What in the YAML is going on here?”
- “Why did my build just spontaneously combust?”
- “Is copy-pasting from StackOverflow a valid CI strategy?”

Well, not anymore.

This tool is your friendly, opinionated robot buddy that **parses**, **lints**, **analyzes**, and **sarcastically judges** your GitHub Actions workflows. It also drops you a friendly PR comment (with just a *hint* of sass) and helps you fix things. Automatically. Because that's the point.

---

## ⚙️ Features

- ✅ **Linting** for common anti-patterns and misconfigurations  
- 🛠 **Auto-fix suggestions** (coming soon: AI-generated apologies)
- 💬 **PR comment integration** with Markdown reports
- 📦 **GitHub Marketplace ready**
- 🐛 **YAML syntax error detection** (like a therapist for malformed colons)
- 🧪 **Unit-tested** lint rules so you sleep well at night
- 📊 **Markdown & JSON** reporters for humans and robots
- 🔍 **Verbose mode** for when you want to feel like a hacker
- 🧠 **Enhanced rules**: catch duplicate jobs, unreferenced steps, zombie workflows, and more
- 😤 **Opinionated**: because life is too short for ambiguous CI

---

## 🚀 How It Works

This tool:

1. Recursively scans your repo for `.github/workflows/*.yml`
2. Parses and validates your YAML files (and roasts them if they’re bad)
3. Runs a series of linting rules on them
4. Generates a nice Markdown/JSON report
5. Comments the results on your PR (if you're into that)
6. Exits with non-zero if it found issues (CI/CD approved)

---

## 🧪 Usage (Local Dev)

```bash
go run main.go --path=.github/workflows --format=markdown --verbose
```

---

## 🧪 GitHub Actions Usage

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

      - name: 🔍 Run GHA Linter
        uses: your-org/gha-linter@v1
        with:
          token: ${{ secrets.GITHUB_TOKEN }}

---

## 🛠 Inputs

Name	Description	                    Required	Default
token	GitHub token for PR comment	      ✅	        –
path	Path to your workflows	          ❌	       .github/workflows
format	Output format (markdown/json)	  ❌	       markdown

---

## 📣 Contributing
Want to add your own lint rules? Just fork and PR — we’ll lint your lint. Please don’t break the linters that lint your linters.

---

## 🧙 Tips
- Verbose mode is your friend during debugging: --verbose
- Use --format=json for CI integrations
- Wanna test PR comment behavior? Push to a fork and set the GITHUB_TOKEN in your local env

---

## 🧹 Future Roadmap
- ✅ Auto-fix YAML patterns with suggestions
- 🧠 AI-generated explanations
- 🦾 Integrate with Dependabot and Renovate
- 🔥 Linting GitHub Action outputs, expressions, and matrix configs

---

## 📄 License
MIT – because we believe in free software.

---

## 🥲 Final Thoughts
Your GitHub Actions shouldn’t require therapy. Use this linter.

---

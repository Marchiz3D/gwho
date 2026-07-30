<div align="center">
  <img src="assets/gwho-logo.png" width="150" alt="Gwho Logo">
  
  # 🎭 gwho (Gw Who?)
  
  **Seamlessly manage and switch between multiple Git profiles across different workspaces.**

[![Go Reference](https://pkg.go.dev/badge/github.com/Marchiz3D/gwho.svg)](https://pkg.go.dev/github.com/Marchiz3D/gwho)
[![Go Report Card](https://goreportcard.com/badge/github.com/Marchiz3D/gwho)](https://goreportcard.com/report/github.com/Marchiz3D/gwho)

_Ever accidentally committed to your company's repository using your anime-avatar personal email? Say no more._

</div>

<br>

> **Note:** 🎓 _This project was built from scratch as a hands-on learning journey to master Clean Architecture and CLI development in Golang. It focuses on clean separation of concerns, SOLID principles, and building beautiful, interactive terminal UIs using Cobra and Fatih/Color._

---

## ✨ Features

- **🚀 Interactive UI:** Beautifully colored terminal prompts for an engaging user experience.
- **🔄 Effortless Switching:** Switch your Git identity locally (per repo) or globally in seconds.
- **🗂️ Profile Management:** Add, edit, list, and remove saved Git profiles with ease.
- **🔍 Active Identity Check:** Instantly verify which profile you are currently using to prevent wrong commits.
- **⚡ Blazing Fast:** Written purely in Go with zero bloated dependencies.

## 📥 Installation

Ensure you have [Node.js](https://nodejs.org/) installed on your machine. Run the following command to install `gwho` globally:

```bash
npm install -g gwho
```

*(Alternatively, if you prefer the Go way: `go install github.com/Marchiz3D/gwho@latest`)*_

## 🎬 Quick Demo

<details>
<summary><b>Click to see what it looks like!</b></summary>

```text
$ gwho list

📋 List of Git Profiles:
========================
[1] personal - John Doe <john.doe@gmail.com>
[2] work     - John Doe <john.doe.dev@companies.com>

$ gwho use
[1] personal
[2] work
Select a profile: 2
> Switched to local profile 'work'
Name: John Doe
Email: john.doe.dev@companies.com
```

</details>

## 🚀 Usage Guide

`gwho` comes with a set of intuitive commands. You can always run `gwho --help` to see the available options.

### 1. Add a Profile

Create a new Git profile by defining an alias (e.g., `work`, `personal`).

```bash
gwho add work
```

### 2. List Profiles

View all your saved Git profiles in a beautifully formatted list.

```bash
gwho list
```

### 3. Switch Profile

Apply a profile to the current Git repository. If you don't provide an alias, an **interactive menu** will pop up!

```bash
gwho use           # Opens interactive selector
gwho use work      # Instantly switches to 'work' locally
gwho use work -g   # Switches to 'work' GLOBALLY across your entire OS
```

### 4. Check Current Profile

Unsure which email is active in your current folder? Just ask:

```bash
gwho current
```

### 5. Edit a Profile

Update the name or email of an existing profile interactively.

```bash
gwho edit work
```

### 6. Remove a Profile

Delete a profile you no longer need.

```bash
gwho remove work
```

## 🛠️ Built With

- [Cobra](https://github.com/spf13/cobra) - The backbone of modern Go CLI applications.
- [Color](https://github.com/fatih/color) - Bringing life to the terminal with beautiful colors.

## 🤝 Contributing

Since this is a learning project, I highly welcome feedback, code reviews, and pull requests! Feel free to open an issue if you find a bug or have a feature request.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

---

<div align="center">
  <i>Built with 🩵 using Go</i>
</div>

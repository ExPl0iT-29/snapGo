
# 📦 SnapGo

> A blazing-fast, Git-independent project snapshot manager built with Go.

**SnapGo** lets you capture and restore the state of your project files at any moment — without touching Git commits. Designed for developers who want lightweight, local versioning and quick experimentation.

---

## 🚀 Features

* 📝 **Create & restore snapshots** instantly
* 💾 **Content-addressable storage** with deduplication
* 🧹 **Auto-cleanup** for stale or unused snapshots
* 🧠 **Custom ignore patterns** (`.snapignore`)
* 🖥️ **Cross-platform support**: Linux, macOS, Windows

---

## ⚙️ Installation

```bash
git clone https://github.com/yourusername/snapgo.git
cd snapgo
go build -o snapgo
```

Or install via `go install`:

```bash
go install github.com/yourusername/snapgo@latest
```

---

## 🧪 Usage

### 📸 Create a Snapshot

```bash
snapgo snap create "my-snapshot"
```

### 🔄 Restore a Snapshot

```bash
snapgo snap restore "my-snapshot"
```

### 🧹 Clean Unused Snapshots

```bash
snapgo snap clean
```

---

## 📚 Project Structure

```
snapgo/
├── main.go
├── snapshot/       # create, restore, clean
├── storage/        # dedup logic
├── config/         # ignore rules, metadata
└── utils/
```

---

## 📜 License

MIT License © \[Your Name]

---

## 🙌 Contributing

Pull requests, feature suggestions, and feedback are always welcome!
Feel free to fork and build your own version.


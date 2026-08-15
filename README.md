<div align="center">
  <h1>🚀 GitDir</h1>
  <p><b>Download specific directories from GitHub repositories at lightning speed.</b></p>
  
  <p>
    <a href="https://github.com/ameen/gitdir/releases/latest"><img alt="Release" src="https://img.shields.io/github/v/release/ameen/gitdir" /></a>
    <a href="https://github.com/ameen/gitdir/actions/workflows/release.yml"><img alt="Build Status" src="https://img.shields.io/github/actions/workflow/status/ameen/gitdir/release.yml" /></a>
    <a href="https://go.dev/"><img alt="Go Version" src="https://img.shields.io/github/go-mod/go-version/ameen/gitdir" /></a>
    <a href="https://github.com/ameen/gitdir/blob/main/LICENSE"><img alt="License" src="https://img.shields.io/github/license/ameen/gitdir" /></a>
  </p>
</div>

<br />

Have you ever wanted to download just a single folder from a massive GitHub repository? 
`gitdir` makes it effortless. Built in Go, it leverages Git's internal `sparse-checkout` functionality to download exactly what you need without cloning the rest of the repository—saving your time and bandwidth.

> 💡 **Why GitDir?**
> Instead of downloading hundreds of megabytes of a monorepo, `gitdir` downloads just the kilobytes you actually want. It's lightning-fast and compiles down to a single binary.

## ✨ Features

- **Blazing Fast**: Uses Git's `blob:none` filter and `sparse-checkout` under the hood.
- **Standalone Binary**: Written in Go. No Node.js, Python, or Ruby environments required.
- **Beautiful DX**: Enjoy smooth loading spinners, vibrant colors, and clear error messages.
- **Cross-Platform**: Works seamlessly on Linux, macOS, and Windows.

## 📦 Installation

### Homebrew (macOS & Linux)
```bash
brew install ameen/tap/gitdir
```

### APT (Debian/Ubuntu)
```bash
echo 'deb [trusted=yes] https://apt.fury.io/ameen-sha/ /' | sudo tee /etc/apt/sources.list.d/gitdir.list
sudo apt update && sudo apt install gitdir
```

### Go Install
If you have Go installed, you can build from source:
```bash
go install github.com/ameen/gitdir@latest
```

### Manual Download
Grab the latest release for your OS from the [Releases Page](https://github.com/ameen/gitdir/releases/latest).

## 🚀 Usage

Simply pass the GitHub URL of the directory you want to download:

```bash
gitdir https://github.com/owner/repo/tree/main/path/to/directory
```

**Options:**
- `-o, --output <dir>`: Specify a custom output directory name. (Defaults to the name of the directory you're downloading).

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## 📄 License

Distributed under the MIT License. See `LICENSE` for more information.

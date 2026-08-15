<div align="center">
  <h1>GitDir</h1>
  <p><b>A CLI tool to download specific subdirectories from GitHub repositories.</b></p>
  
  <p>
    <a href="https://github.com/Ameen-Sha-Cheerangan/GitDir/releases/latest"><img alt="Release" src="https://img.shields.io/github/v/release/Ameen-Sha-Cheerangan/GitDir" /></a>
    <a href="https://github.com/Ameen-Sha-Cheerangan/GitDir/actions/workflows/release.yml"><img alt="Build Status" src="https://img.shields.io/github/actions/workflow/status/Ameen-Sha-Cheerangan/GitDir/release.yml" /></a>
    <a href="https://go.dev/"><img alt="Go Version" src="https://img.shields.io/github/go-mod/go-version/Ameen-Sha-Cheerangan/GitDir" /></a>
    <a href="https://github.com/Ameen-Sha-Cheerangan/GitDir/blob/main/LICENSE"><img alt="License" src="https://img.shields.io/github/license/Ameen-Sha-Cheerangan/GitDir" /></a>
  </p>
</div>

<br />

`gitdir` is a Go-based CLI tool that allows you to download specific directories from GitHub repositories without cloning the entire project. It uses Git's internal `sparse-checkout` and `--filter=blob:none` functionality to minimize bandwidth and storage usage by only fetching the files you request.

## Features

- **Efficient**: Uses Git's `sparse-checkout` feature under the hood.
- **Standalone**: Distributed as a single compiled Go binary with no external language dependencies.
- **Cross-Platform**: Supports Linux, macOS, and Windows.
- **Clear Feedback**: Provides direct terminal output and error handling during downloads.

## 📦 Installation

### APT (Debian/Ubuntu)
```bash
echo 'deb [trusted=yes] https://Ameen-Sha-Cheerangan.github.io/GitDir/ /' | sudo tee /etc/apt/sources.list.d/gitdir.list
sudo apt update && sudo apt install gitdir
```
### Homebrew (macOS & Linux)
> **⚠️ Important:** Because this is a third-party tap, Homebrew requires you to explicitly trust the repository before installing. *(Note: If this tool gains enough users, I will try submitting it to the official Homebrew Core so this extra step won't be necessary! )*

```bash
brew tap Ameen-Sha-Cheerangan/homebrew-tap
brew trust Ameen-Sha-Cheerangan/tap
brew update && brew install gitdir
```

### Go Install
If you have Go installed, you can build from source:
```bash
go install github.com/Ameen-Sha-Cheerangan/GitDir@latest
mv ~/go/bin/GitDir ~/go/bin/gitdir
```

### Manual Download
Grab the latest release for your OS from the [Releases Page](https://github.com/Ameen-Sha-Cheerangan/GitDir/releases/latest).

## 🚀 Usage

Simply pass the GitHub URL of the directory you want to download:

```bash
gitdir https://github.com/owner/repo/tree/main/path/to/directory
```

**Options:**
- `-o, --output <dir>`: Specify a custom output directory name. (Defaults to the name of the directory you're downloading).

## 🤝 Contributing

Patches, bug reports, and feature requests are always welcome! Feel free to open an issue or a pull request on this repository.

## 📄 License

Distributed under the MIT License. See `LICENSE` for more information.

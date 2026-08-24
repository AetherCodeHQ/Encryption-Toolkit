# 🔐 Encryption Toolkit

![Go](https://img.shields.io/badge/Go-1.21%2B-00ADD8?style=flat-square&logo=go&logoColor=white)
![Version](https://img.shields.io/badge/Version-v3.0.0-00ADD8?style=flat-square)
![License](https://img.shields.io/badge/License-MIT-green?style=flat-square)
![PRs](https://img.shields.io/badge/PRs-Welcome-brightgreen?style=flat-square)

> Security tool by [AetherCodeHQ](https://github.com/AetherCodeHQ)

`security` `cryptography` `cli` `golang` `crypto`

---

## What is Encryption-Toolkit?

**Encryption-Toolkit** is a security-focused tool that analyzes and validates code, configurations, or data for vulnerabilities.

## Features

- ✅ Cryptographic operations
- 🚀 **Zero dependencies** — only Go standard library
- 📦 **Single binary** — compile and run anywhere
- 🔄 **Offline capable** — no internet required

## Installation

```bash
# Clone
git clone https://github.com/AetherCodeHQ/Encryption-Toolkit.git
cd Encryption-Toolkit

# Build
go build -o encryption-toolkit .

# Run
./encryption-toolkit <encrypt|decrypt> <file> <key-hex>
```

### Or directly with `go run`:
```bash
go run main.go <encrypt|decrypt> <file> <key-hex>
```

## Usage

```bash
# Basic usage
./encryption-toolkit <encrypt|decrypt> <file> <key-hex>

# With flags
./encryption-toolkit <encrypt|decrypt> <file> <key-hex> value <encrypt|decrypt> <file> <key-hex>
```

### Example Output

```
$ ./encryption-toolkit <encrypt|decrypt> <file> <key-hex>
<encrypt|decrypt> <file> <key-hex>
encrypted %d -> %d bytes -> %s\n
decrypted %d -> %d bytes -> %s\n
```

## Project Structure

```
Encryption-Toolkit/
  main.go          # Entry point (72 lines)
  go.mod            # Go module definition
  go.sum            # Dependency checksums
  README.md         # This file
  LICENSE           # MIT License
  CHANGELOG.md      # Version history
```

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

MIT License - see [LICENSE](LICENSE) for details.

---

Built with ❤️ by [AetherCodeHQ](https://github.com/AetherCodeHQ)

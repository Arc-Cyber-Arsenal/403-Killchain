# 🔗 403-Killchain

### The Ultimate 403 Forbidden Bypass Framework

**403‑Killchain** fuses four independent bypass engines into a single, silent binary.  
Each engine uses a unique strategy to turn `403 Forbidden` into `200 OK`.  
They run **in parallel**, and the results are merged into one clear, actionable report.  
No more running multiple tools, no more missed tricks. One command, one killchain.

<p align="center">
  <img src="https://img.shields.io/badge/version-1.0.0-blue" />
  <img src="https://img.shields.io/badge/go-1.20%2B-00ADD8?logo=go" />
  <img src="https://img.shields.io/badge/platform-windows%20%7C%20linux%20%7C%20macOS-lightgrey" />
  <img src="https://img.shields.io/badge/license-MIT-green" />
</p>

---

## 💡 How 403‑Killchain Works

The tool embeds four distinct engines, each attacking the same URL with its own set of techniques:

### Engine 1 – Path & Header Fuzzer
This engine systematically mutates the URL path and injects common bypass headers.
- Appends slashes, encoded characters, and path traversal sequences (`/`, `%2e`, `;`, `..;/`)
- Adds headers like `X-Forwarded-For`, `X-Originating-IP`, `X-Rewrite-URL`, `X-Real-IP`
- Switches HTTP methods on the fly (GET → POST, PUT, PATCH, DELETE, OPTIONS)
- Sends every combination quickly and reports any response that isn’t a 403

### Engine 2 – Header & Method Mutator
This engine takes a more protocol‑oriented approach, manipulating the request line and headers to confuse access controls.
- Inserts malformed request lines (HTTP/1.0 downgrades, extra spaces, double slashes)
- Tries alternative host headers, `X-Forwarded-Host`, `Origin` overrides
- Exploits discrepancies between front‑end proxies and back‑end servers
- Uses payloads designed to bypass WAFs and reverse proxies

### Engine 3 – Status‑Code Reasoner
This engine brute‑forces the common patterns that web servers use to enforce 403 rules.
- Tests every known HTTP method with different Content‑Types
- Sends the same request multiple times with small changes (case‑flipping, Unicode tricks)
- Learns from `nmap`‑style path lists and common admin panels
- Flags any response that silently returns 200, 301, or 302

### Engine 4 – Encoding & Randomisation
This engine uses entropy to avoid pattern‑based blocking.
- URL‑encodes, double‑encodes, and Unicode‑encodes every path segment
- Generates random parameter names and values to evade signature‑based filters
- Sends requests at slightly varied speeds (jitter) to stop rate‑limiting
- Combines multiple bypass techniques into single, high‑entropy requests

All four engines run **concurrently** – you get the full collective output in less time than it takes to run two of them separately.

---

## 🧰 Features

- 🧬 **4 engines, 1 binary** – no external dependencies at runtime
- ⚡ **Parallel execution** – all engines attack simultaneously
- 🧠 **Engine selection** – use `--only` or `--skip` to pick specific engines
- 📋 **Unified report** – clear, sectioned output for easy analysis
- 🔒 **Embedded payloads** – everything is compiled directly into the binary
- 🪶 **Static binary** – works on Linux, Windows, and macOS without installers
- 🧹 **Zero external API calls** – runs entirely offline

---

## 🚀 Installation

### From pre‑compiled binary
Download from [Releases](https://github.com/Arc-Cyber-Arsenal/403-Killchain/releases)

### Build from source (Go 1.20+ required)
```bash
git clone https://github.com/Arc-Cyber-Arsenal/403-Killchain
cd 403-Killchain
CGO_ENABLED=0 go build -o 403-Killchain
```

---

## 🎯 Usage

```bash
./403-Killchain -u https://target.com/admin
```

### Flags

| Flag    | Description |
|---------|-------------|
| `-u`    | Target URL (required) |
| `--only`| Run only specific engines (comma‑separated) <br> Example: `--only nomore403,byp4xx` |
| `--skip`| Skip engines (comma‑separated) <br> Example: `--skip bypass-403.sh,dontgo403` |
| `-l`    | List all embedded engines and exit |
| `-h`    | Show banner and help |

---

## 📊 Comparison with Running Tools Individually

| Capability           | Standalone tools | 403‑Killchain |
|----------------------|:----------------:|:-------------:|
| Multiple engines     | ❌               | ✅ All four   |
| Parallel execution   | ❌               | ✅            |
| Unified output       | ❌               | ✅            |
| Single binary        | ~                | ✅            |
| No external deps     | ~                | ✅            |
| Full technique coverage | partial       | ✅ complete   |

---

## 👤 Author & Organisation

- **Archsec-Emman** – [@Archsec-Emman](https://github.com/Archsec-Emman)
- **Arc‑Cyber‑Arsenal** – [https://github.com/Arc-Cyber-Arsenal](https://github.com/Arc-Cyber-Arsenal)

---

## 📄 License

MIT License © 2025 Archsec-Emman  
The embedded engines contain code from various open‑source projects – see `THIRD_PARTY_LICENSES.md` for full attributions.

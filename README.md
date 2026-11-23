# prism-foundry-145

Current version: `1.0.0`.

> Work in progress. This checkpoint may be incomplete or fail tests. See PLAN.md and STATUS.md before use.


Every agent's model. One place. Codex on DeepSeek, Claude Code on Kimi, from the menu bar.

## Installation

```bash
# Prerequisites
Go 1.27.4 or later

# Build
go build ./...

# Run tests
go test ./...
```

## Features (planned roadmap)

1. **Session management core** - Fix nil channel panic (#2932) — DONE
2. **GUI titlebar/corners** - Address macOS 26/27 titlebar behavior (#3949)
3. **Router/Provider capacity handling** - "Selected model is at capacity" (#1418)
4. **Release notes i18n** - Chinese release notes support (#2260)

## Contact

- **Website**: https://alanvo.net
- **X**: https://x.com/alanvoasi

## License

MIT

Maintainer: Alan Vo <alanvo@gmail.com>

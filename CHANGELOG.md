# Changelog

## 1.0.0 — work-in-progress checkpoint

Partial code published for continued development. Known issues: STATUS.md.

# prism-foundry-145 Changelog

## v0.1.0 - 2026-11-23

### Fixed
- **sessions: panic: close of nil channel** (#2932)
  - Added nil-channel guard in `publish()`: skips `close()` on nil channels instead of crashing
  - Added duplicate-path guard in `refresh()`: prevents duplicate paths in `todo` from causing channel overwrites
  - Updated `statsIn()`: checks channel is non-nil before operating on it

### Added
- Session management package (`internal/sessions/`) with refresh/publish lifecycle
- GUI integration layer (`internal/gui/sessions.go`)
- Test suite covering the nil-channel fix and related functionality

### Notes
- This is the initial checkpoint; future rounds will add GUI titlebar components, router/provider handling, and i18n release notes.

---

Alan Vo · [alanvo.net](https://alanvo.net) · [X: @alanvoasi](https://x.com/alanvoasi)

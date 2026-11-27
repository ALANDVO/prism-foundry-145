# Changelog

## 1.1.0 — work-in-progress checkpoint

Partial code published for continued development. Known issues: STATUS.md.

## 1.0.0 — work-in-progress checkpoint

Partial code published for continued development. Known issues: STATUS.md.

### Added
- **router: provider capacity fall-through** (#1418 slice)
  - `internal/router.SelectAvailable` returns the first provider in a group that
    is not at capacity, preserving priority order, and reports failure when the
    group is empty or all providers are at capacity.
  - Tests in `internal/router/router_test.go` (skip-at-capacity, all-at-capacity,
    empty group, priority order).

### Fixed (prior round, retained)
- **sessions: panic: close of nil channel** (#2932)
  - Nil-channel guard in `publish()`; duplicate-path guard in `refresh()`.

### Notes
- This is an incremental checkpoint; GUI titlebar (#3949), Chinese release notes
  (#2260), and live router wiring remain planned.

---

Alan Vo · [alanvo.net](https://alanvo.net) · [X: @alanvoasi](https://x.com/alanvoasi)

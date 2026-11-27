# b36date

Prints a date as a compact four-character code: a two-digit year, then the
month and the day as single base36 digits (`0-9`, then `A-Z`).

| Date       | Code   |
|------------|--------|
| 2026-11-26 | `26BQ` |
| 2026-01-01 | `2611` |
| 1999-12-31 | `99CV` |

The year is `year % 100`, so 1926 and 2026 both give `26`. Months run `1`–`C`
and days `1`–`V`.

## Install

```sh
go install github.com/NWoodsman/b36date@latest
```

Or with Nix:

```sh
$ nix run github:NWoodsman/b36date -- 2026-11-26
26BQ
```

Leave off `-- <date>` to get today's code, or use
`nix shell github:NWoodsman/b36date` to put `b36date` on your PATH.

## Usage

With no arguments it prints today's date:

```sh
$ b36date
26BQ
```

Pass one or more dates to convert them, one code per line:

```sh
$ b36date 2026-11-26 "Nov 26, 2026" 1999-12-31
26BQ
26BQ
99CV
```

These input formats are accepted:

- `2026-11-26`
- `2026/11/26`
- `20261126`
- `11/26/2026` (month first)
- `Nov 26 2026` or `November 26, 2026`
- `26 Nov 2026`
- RFC3339, e.g. `2026-11-26T09:30:00Z`

If a date can't be read, b36date prints an error to stderr and exits with
status 1.

# b36date

Agents run amok making branches, in worktrees, in clones, everywhere. To create a consistent workflow, I created a small date format called b36date using base36.

It's a four digit scheme YYMD where M and D are the values in base 36. The resulting dates are human readable, sortable, and easily visible in git branches.

Base 36 uses digits 0-9 and capitalized alphas A-Z; zero and letters W-Z are unused.

The year is `year % 100`, so 1926 and 2026 both give `26`. Months run `1`–`C`
and days `1`–`V`.

Thanksgiving 2026 is `26BQ`.

`26` = 2026

`B` = November (11=B)

`Q` = 26


| Date       | Code   |
|------------|--------|
| 2026-11-26 | `26BQ` |
| 2026-01-01 | `2611` |
| 1999-12-31 | `99CV` |


Your git branches then become:

`26BQ-feature-description-dashed`

`26BQ-fix-broken-link`

`26BQ-staged`

First teach your agents about the format and add it to PATH. Then you can then tell your agents to "push today's b36date branches" and "merge today's b36date branches into a xxxx-staging branch where xxxx is todays b36date."

It's easy for them to spot all accumulated branches and easily delete them at the end of the day.


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

## License

MIT. See [LICENSE](LICENSE).

muh website

## Build

### Requirements

- [Tailwind CSS (v3\*)](https://v3.tailwindcss.com/);
- [SQLite3](https://sqlite.org/index.html);
- [sqlc](https://sqlc.dev/);
- [templ](https://templ.guide/);
- [goose](https://pressly.github.io/goose/) (database migrations);
- [air](https://github.com/air-verse/air) (live reloading);

The `sqlc`, `templ`, `goose`, and `air` packages can be installed directly via Go (recommended):

```bash
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
go install github.com/a-h/templ/cmd/templ@latest
go install github.com/pressly/goose/v3/cmd/goose@latest
go install github.com/air-verse/air@latest
```

The TailwinCSS CLI (MUST be 3.X.X) tool can be installed via the [Node Package Manager](https://tailwindcss.com/docs/installation/tailwind-cli) (_ewwww_) or [as a standalone binary](https://tailwindcss.com/blog/standalone-cli):

```bash
# v3.4.16 Latest v3 release as of writing
curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/download/v3.4.16/tailwindcss-linux-x64
chmod +x tailwindcss-linux-x64
mv tailwindcss-macos-arm64 /otp/tailwindcss # assuming /opt is in your $PATH
```

SQLite3 is available on [most package managers](https://launchpad.net/ubuntu/+source/sqlite3) and can also be downloaded as a binary, built from source, etc.:

```bash
# Debian/Ubuntu
sudo apt install sqlite3
```

## Execution

Run it natively:

```bash
# Builds the binary (PROD)
make build
# Runs it with `air` (live reload) for development
make start
```

Run it in a container:

```bash
make docker-build
make docker-run
```

Both will run on port 8080 by default.

### TODO:

- [ ] `/home`;
- [ ] `/resume`;
- [ ] `/blog/{X}`;

---


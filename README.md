muh website [pedrobinotto.xyz](https://www.pedrobinotto.xyz/)

## Build

### Requirements

- [Tailwind CSS (v3\*)](https://v3.tailwindcss.com/);
- [PostgreSQL](https://www.postgresql.org/);
- [sqlc](https://sqlc.dev/);
- [templ](https://templ.guide/);
- [air](https://github.com/air-verse/air) (live reloading);
- make;
- go;

The `sqlc`, `templ` and `air` packages can be installed directly via Go (recommended):

```bash
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
go install github.com/a-h/templ/cmd/templ@latest
go install github.com/air-verse/air@latest
```

The TailwinCSS CLI (MUST be 3.X.X) tool can be installed via the [Node Package Manager](https://tailwindcss.com/docs/installation/tailwind-cli) (_ewwww_) or [as a standalone binary](https://tailwindcss.com/blog/standalone-cli):

```bash
# v3.4.16 Latest v3 release as of writing
curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/download/v3.4.16/tailwindcss-linux-x64
chmod +x tailwindcss-linux-x64
mv tailwindcss-linux-x64 /otp/tailwindcss # assuming /opt is in your $PATH
```

PostgreSQL is available on [most package managers](https://launchpad.net/ubuntu/+source/postgresql-common) and can also be downloaded as a binary, built from source, etc.:

```bash
# Debian/Ubuntu
sudo apt install postgresql
```

## Configuration

All configuration is done via `.env`; an example file can be found in `.env.example`

```bash
DB_HOST=website-db
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=app_db
DB_PORT=5432
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
make docker-run
```

Both will run on port 8080 by default.

### TODO:

- [ ] Automate `sitemap.xml` generation on build;
- [ ] `/blog/{X}`: 
    - [ ] setup blogs repo + CI actions;
    - [ ] fix DB indexing for files;

---


# Jadwal (Work in Progress) – Multi-List Task Tables

A Go experiment in structured task management,
Jadwal organizes todos across multiple tables or lists—like a lightweight database for your goals.
Designed primarily as a learning project (not a polished tool), it’s a playground for exploring Go, architecture, and open-source collaboration.

## Key Notes:

- 🛠 Work in Progress: Buggy, evolving, and not production-ready.
- 📚 Made for Learning: My goal? To tinker with Go, concurrency, and clean code.
- 🗂 Multi-Table Tasks: Group todos by project, priority, or whimsy (e.g., work, learn-go, groceries).
- 🤝 Open Door Policy: Feel free to use, fork, or critique—merge requests and issues welcome!

## Resources:

- [Go by Example](https://gobyexample.com/)
- [Jason McVetta | Golang for Python programmers](https://golang-for-python-programmers.readthedocs.io/)
- [Standard Go Project Layout](https://github.com/golang-standards/project-layout)
- [joneisen.me | Golang and defalult values](https://joneisen.me/programming/2013/06/23/golang-and-default-values.html)
- [Learn X in Y minutes | Where X=Go](https://learnxinyminutes.com/go/)
- [Atlas GO | docs index](https://atlasgo.io/docs)
- [Atlas GO | SQLc Declarative](https://atlasgo.io/guides/frameworks/sqlc-declarative)

```bash
atlas schema apply \
    --url "sqlite://$HOME/.local/share/.jadwal.db?_journal_mode=WAL&_fk=1" \
    --dev-url "sqlite://$HOME/.local/share/.jadwal.db?_journal_mode=WAL&_fk=1" \
    --to "file://third_party/sqlc/schema.sql"


atlas migrate diff initial_schema \
  --dir "file://migrations" \
  --to "file://third_party/sqlc/schema.sql" \
  --dev-url "sqlite://$HOME/.local/share/.jadwal.db?_journal_mode=WAL&_fk=1"

atlas migrate apply --url "sqlite://$HOME/.local/share/.jadwal.db?_journal_mode=WAL&_fk=1"

atlas migrate new seed_demo_data --dir file://migrations
```

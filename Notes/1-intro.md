# Database from Scratch


Refer: https://build-your-own.org/database/00a_overview

## 1. Durability and atomicity

Traditional database only: MySQL, Postgres, SQLite have in common?

- They persist Data to disk - Durability

- Disk-based, can work with larger than memory data

- Implemented from scratch and not as wrappers over other databases

The number 1 criterion of traditional DB is Durablity. Mobile phones use SQLite (file based db).
But why not use files if DB is a file?
- Cuz of "Persisting" i.e. Data is guranteed to persist even if the machine crashes from any reason

2 requirements to survive the crash:
- Most DB run on filesystem, therefore the used filesystem must also meet these requirements (filesystems ~ databases)

**Difference Between Filesystems and DBs**:
1. Filesystem (writing to files) has no durabiliity gurantee (leads to data loss)


Therefore, making files durablle = 1/2 a DB

`fsync` syscall: filesystem operation that makes all previous written data durable

- Requests and confirms duability
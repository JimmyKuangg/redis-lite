# Redis Lite

A minimal remake of Redis in Go that emphasizes education over robust usage.

RedisLite features only a subset of what Redis has to offer. The goal wasn't to replace Redis. It was to understand the underyling mechanisms and design decisions.

## Disclaimer

RedisLite is intentionally designed to be a learning project. It does not feature all of the robust features that Redis contains, but priotizes education and understanding over having a highly scalable database. For design decisions made, read my notes about how my brain tried to figure out what to do [here](https://jimmy-kuang.com/notes).

## Installation

- Ensure you have the [latest version of go installed](https://go.dev/)
- Clone the repository and run commands using Go.

```bash
git clone https://github.com/JimmyKuangg/redis-lite.git
cd redis-lite

go run .
RedisLite listening on port 6379
```

## Features

- Localhost server to use TCP connections on port 6379
- Commands such as `GET`, `SET`, and others that allow you to interact with the database (full list below)
- Concurrency safe actions by using mutex's and locks
- Data persistence and hydration via AOF and RDB files
- TTL and clean up to allow for some memory clean up

## Why Rebuild Redis?

Because I literally had no idea what it was or how to use it. So, what better way to learn than to jump into the deep end and try?

## Usage

### Table of Commands

| Command | Args | Example         | Response       | Explanation                                          |
| ------- | ---- | --------------- | -------------- | ---------------------------------------------------- |
| PING    | 0    | PING            | PONG           | Ping the server to ensure it's alive (or for fun)    |
| SET     | 2    | SET name Jimmy  | OK             | Set a key value pair within the database             |
| MSET    | even | MSET a 1 b 2    | OK             | Set multiple key value pairs in the database at once |
| GET     | 1    | GET name        | Jimmy          | Retrieve the value of a key from the database        |
| MGET    | >=1  | Get name gender | "Jimmy" "Male" | Retrieve values of multiple keys from the database   |
| DEL     | >=1  | DEL name        | OK             | Delete a key within the database                     |
| EXPIRE  | 2    | EXPIRE name 30  | OK             | Set a TTL on a key. The number is in seconds         |
| PRINT   | 0    | PRINT           | key/value dump | Print out the current state of your database         |

### Example of Usage

```bash
# Terminal one
# Keep an eye on this terminal!
# It spits out operational logs
go run .
RedisLite listening on port 6379

# In a separate terminal, terminal two
nc localhost 6379
SET name Jimmy
OK
GET name
Jimmy
PRINT
name => Jimmy, TTL: <nil>
DELETE name
OK
GET name
key does not exist

# To exit out and stop the server, press Ctrl + C on terminal one
# This applies for both Windows and Mac
```

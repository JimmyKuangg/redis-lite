# RedisLite

A minimal remake of Redis in Go that emphasizes education over covering all its existing features.

RedisLite features only a subset of what Redis has to offer. The goal wasn't to replace Redis. It was to understand the underlying mechanisms and design decisions behind an in-memory database.

## Disclaimer

RedisLite is intentionally designed to be a learning project. It does not feature all of the robust features that Redis contains, but priotizes education and understanding over having a highly scalable database. For design decisions made, read my notes about how my brain tried to figure out what to do [here](https://jimmy-kuang.com/notes).

## Installation

- Ensure you have the [latest version of go installed](https://go.dev/)
- Clone the repository and run commands using Go.

```bash
git clone https://github.com/JimmyKuangg/redis-lite.git
cd redis-lite

# Running the project
go run .
RedisLite listening on port 6379
```

## Features

- Localhost server to use TCP connections on port 6379
- Commands such as `GET`, `SET`, and others that allow you to interact with the database (full list below)
- Concurrent-safe operations using mutexes and read/write locks
- Data persistence and hydration via AOF and RDB files
- TTL support with a background cleanup worker for expired keys

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
# Connects you to the open port at 6379
nc localhost 6379
# Sets a key
SET name Jimmy
OK
# Retrieves a key
GET name
Jimmy
# Prints out the key value pairs of the database
PRINT
name => Jimmy, TTL: <nil>
# Deletes a key
DEL name
OK
GET name
key does not exist

# To exit out and stop the server, press Ctrl + C on terminal one
# This applies for both Windows and Mac
```

## How it Works

Overall structure of the project -

```bash
                     RedisLite

                +-----------------+
                |     Clients     |
                | (nc / programs) |
                +--------+--------+
                         |
                    TCP :6379
                         |
                         ▼
                +-----------------+
                |     Server      |
                | Parse & Execute |
                +--------+--------+
                         |
                         ▼
                +-----------------+
                | In-Memory Store |
                | map[string]Entry|
                +---+---------+---+
                    |         |
          Read/Write|         |Snapshots
                    |         |
                    ▼         ▼
             +----------+  +----------+
             |   AOF    |  |   RDB    |
             +----------+  +----------+

          Background TTL Cleanup Worker
                (Every 10 Seconds)
```

Multiple things happen upon server startup -

```bash
                  go run .
                      │
                      ▼
        +----------------------------+
        |      Load Snapshot         |
        |     (.redislite/rdb)       |
        +----------------------------+
                      │
                      ▼
        +----------------------------+
        |       Restore Data         |
        |    Into Memory (HashMap)   |
        +----------------------------+
                      │
                      ▼
        +----------------------------+
        |        Replay AOF          |
        |    (.redislite/aof)        |
        +----------------------------+
                      │
                      ▼
        +----------------------------+
        | Start Cleanup Worker (TTL) |
        |    Runs Every 10 Seconds   |
        +----------------------------+
                      │
                      ▼
        +----------------------------+
        | Listen on TCP :6379        |
        | Wait for Client Requests   |
        +----------------------------+
                      │
                      ▼
              Ready to Accept Clients
```

Whenever you enter a CLI command, this the flow of the commands -

```bash
        nc localhost 6379
                │
                ▼
        +----------------+
        | TCP Connection |
        +----------------+
                │
                ▼
        +----------------+
        | Parse Command  |
        +----------------+
                │
                ▼
        +----------------+
        | Execute Logic  |
        +----------------+
                │
                ▼
        +----------------+
        | Database (RAM) |
        +----------------+
                │
         ┌──────┴────────┐
         │               │
         ▼               ▼
 Return Response    Append to AOF
   To Client         (if modified)
```

The cleanup worker logic flow -

```bash
           Every 10 Seconds
                  │
                  ▼
        +----------------------+
        | Lock Database         |
        +----------------------+
                  │
                  ▼
        +----------------------+
        | Scan All Entries     |
        +----------------------+
                  │
                  ▼
        +----------------------+
        | Expired?             |
        +----------------------+
           │             │
        Yes│             │No
           ▼             ▼
     Delete Entry     Keep Entry
           │
           ▼
      Unlock Database
```

## Project Structure

`.redislite/`

- Where the storage folders live. Take your time to watch the two files as you use commands in the CLI! It'll write to them automatically.
- Try NOT to edit them. If you have issues with starting the server using `go run .`, delete this folder and run it again. Note that this will delete your data.

`commands/`

- The parsing and execution of the CLI command logic lives here.

`data/`

- Holds the logic behind the commands executed in the database.

`server/`

- All the logic for the actual "server" side of RedisLite sits here. The server starting, what it calls, how it handles the concurrent connections.

`storage/`

- The way the storage is handled within this remake. This includes the AOF and RDB files and also the hydration part of the server startup.

## Limitations

RedisLite intentionally does not contain all the robust features that the brilliant engineers at Redis spent years developing, such as:

- Redis's single-threaded event loop
- LRU memory management
- A RESP protocol parser (I just parse strings)
- Memory limits
- Data sharding

These features are beyond the scope of an educational project, but can be explored down the line.

## Takeaways and Learnings

Learned a lot about:

- Mutex's and concurrency safety
- Design around failure (What happens if a write fails? Too many writes? etc.)
- A general idea of how Redis works

RedisLite answered many of my original questions about Redis, but it also raised new ones. I didn't implement memory eviction, sharding, replication, or Redis's networking protocol, but understanding why those features exist was just as valuable as building the pieces I did. Those questions are what I'll be exploring next.

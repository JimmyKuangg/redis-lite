# Redis Lite

Educational remake of Redis in progress. Follow my shenanigans as I write stuff up [here!](https://jimmy-kuang.com/notes)

## Usage

| Command | Args | Example         | Response       |
| ------- | ---- | --------------- | -------------- |
| PING    | 0    | PING            | PONG           |
| SET     | 2    | SET name Jimmy  | OK             |
| MSET    | even | MSET a 1 b 2    | OK             |
| GET     | 1    | GET name        | Jimmy          |
| MGET    | >=1  | Get name gender | "Jimmy" "Male" |
| DEL     | >=1  | DEL name        | OK             |
| EXPIRE  | 2    | EXPIRE name 30  | OK             |
| PRINT   | 0    | PRINT           | key/value dump |

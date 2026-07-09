# Redis Lite

Educational remake of Redis in progress. Follow my shenanigans as I write stuff up [here!](https://jimmy-kuang.com/notes)

## Usage

| Command | Args | Example         | Response       |
| ------- | ---- | --------------- | -------------- |
| PING    | 0    | PING            | PONG           |
| SET     | 2    | SET name Jimmy  | OK             |
| GET     | 1    | GET name        | Jimmy          |
| MGET    | >=1  | Get name gender | "Jimmy" "Male" |
| MSET    | even | MSET a 1 b 2    | OK             |
| DEL     | >=1  | DEL name        | OK             |
| PRINT   | 0    | PRINT           | key/value dump |

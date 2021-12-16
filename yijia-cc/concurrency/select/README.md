# Select
## Getting Started

```bash
go run select.go
```

Sample Output:
```txt
[Debug] [2021-08-02 13:28:16] Start fetching Blueberry
[Debug] [2021-08-02 13:28:16] Start fetching Banana
[Debug] [2021-08-02 13:28:16] Start fetching Apple
[Debug] [2021-08-02 13:28:16] Start fetching Watermelon
[Debug] [2021-08-02 13:28:17] End fetching Banana
[Info] [2021-08-02 13:28:17] Done fetching Banana
[Debug] [2021-08-02 13:28:18] End fetching Apple
[Info] [2021-08-02 13:28:18] Default: Apple
[Debug] [2021-08-02 13:28:19] End fetching Watermelon
[Info] [2021-08-02 13:28:19] Default: Watermelon
[Debug] [2021-08-02 13:28:19] End fetching Blueberry
[Info] [2021-08-02 13:28:19] Default: Blueberry
[Info] [2021-08-02 13:28:21] Start receiving the first query result
[Info] [2021-08-02 13:28:21] Final result: Banana
```
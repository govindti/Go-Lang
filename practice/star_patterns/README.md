# Star Patterns Practice

LeetCode-style practice: sirf **function ka body** likhna hai, baaki sab likha hua hai.

## Kaise karein

1. Kisi bhi pattern ka folder kholo (e.g. `01_triangle/`)
2. `main.go` me comment me pattern ka example diya hai
3. Function ka body khaali hai (`return nil`). Isse hatao aur apna code likho
4. `go run .` se apna output dekh sakte ho (n=5 print hota hai)
5. Test chalao:

```bash
go test ./...                     # sab patterns
go test ./01_triangle/            # sirf ek pattern
go test -run Triangle ./01_triangle/
```

Test pass hua to pattern sahi hai. Har folder me 3 predefined test cases hain
(`n = 1, 3, 5`).

## Help

- `strings.Repeat("*", n)` se n stars banao
- Function `[]string` return karta hai — ek string har line ke liye
- Spaces bhi matter karte hain (right-aligned / pyramid patterns me)

## Patterns

| # | Folder | Shape |
|---|--------|-------|
| 1 | `01_triangle` | Right triangle (top-left) |
| 2 | `02_inverted_triangle` | Inverted triangle (bottom-left) |
| 3 | `03_square` | Solid square |
| 4 | `04_right_aligned_triangle` | Right triangle (top-right) |
| 5 | `05_pyramid` | Centered pyramid |
| 6 | `06_inverted_pyramid` | Inverted pyramid |
| 7 | `07_diamond` | Diamond |
| 8 | `08_hollow_square` | Hollow square |
| 9 | `09_hollow_triangle` | Hollow triangle |
| 10 | `10_hourglass` | Hourglass |

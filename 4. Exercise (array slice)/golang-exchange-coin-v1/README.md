# Exercise Array Slice 2

## Exchange coin

Buatlah sebuah fungsi yang dapat mengecek pecahan dari kumpulan koin yang di-_input_, pecahan yang tersedia adalah 1000, 500, 200, 100, 50, 20, 10, 5 dan terkecil adalah 1 koin.

Fungsi ini akan mengembalikan berupa tipe data _slice of_ `int`. Perlu diperhatikan bahwa hasil kembalian berupa pecahan koin yang urut dari pecahan terbesar ke terkecil.

Contoh _test case_ :

```txt
input: amount = 1752
output: [1000, 500, 200, 50, 1, 1]

input: amount = 5000
output: [1000, 1000, 1000, 1000, 1000]

input: amount = 1234
output: [1000, 200, 20, 10, 1, 1, 1, 1]

input: amount = 0
output: []

...
```

Berikut _template_ pengerjaannya:

```go
func ExchangeCoin(amount int) []int {
    // kerjakan di sini
}
```

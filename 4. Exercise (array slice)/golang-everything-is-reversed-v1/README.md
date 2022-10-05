# Exercise array slice 1

## Everything is reversed

Terdapat sebuah _array_ yang memiliki _size_ sebesar 5 yang bertipe `int`, dan kita ingin mengubah semuanya menjadi terbalik, baik secara urutan data ataupun urutan isi setiap data. Buatlah sebuah fungsi yang dibutuhkan untuk mengubah semua data menjadi terbalik.

Contoh, jika data berisi `[12, 23, 34, 45, 56]` maka urutan data tersebut dibalik dan urutan setiap data juga dibalik menjadi `[65, 54, 43, 32, 21]`.

Angka di _array_ tersebut hanya berisi angka positif, jadi tidak akan ada angka negatif.

Contoh _test case_ :

```txt
input: arr = [123, 456, 11, 1, 2]
output: [2, 1, 11, 654, 321]

input: arr = [456789, 44332, 2221, 12, 10]
output: [1, 21, 1222, 23344, 987654]

input: arr = [10, 10, 10, 10, 10]
output: [1, 1, 1, 1, 1]

input: arr = [23456, 789, 123, 456, 500]
output: [5, 654, 321, 987, 65432]

input: arr = [0, 0, 0, 0, 0]
output: [0, 0, 0, 0, 0]

...
```

Berikut _template_ pengerjaannya:

```go
func ReverseData(arr [5]int) [5]int {
    // kerjakan di sini
}
```

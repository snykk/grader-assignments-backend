# Exercise array slice 4

## Line of march

Terdapat ekstrakulikuler paskibra yang memiliki beberapa anggota dan sedang melakukan kegiatan baris-berbaris. Sebelum melakukan kegiatan tersebut, mereka harus berdiri sesuai dengan urutan tinggi badan dari yang paling rendah ke yang paling tinggi.

Buatlah sebuah fungsi yang akan mengurutkan anggota paskibra tersebut dari yang paling rendah ke yang paling tinggi. Setiap orang akan memberikan informasi tinggi badannya dan ditampung ke dalam _slice_. Setelah itu, urutkanlah _slice_ tersebut dan kembalikan _slice_ yang sudah terurut.

Contoh _test case_ :

```txt
input: heights : [172, 170, 150, 165, 144, 155, 159]
output: [144, 150, 155, 159, 165, 170, 172]

input: heights: [155, 156, 160, 161, 162, 165, 170, 172]
output: [155, 156, 160, 161, 162, 165, 170, 172]

input: heights: [180, 177, 175, 173, 170, 166, 165, 160]
output: [160, 165, 166, 170, 173, 175, 177, 180]

input: heights: []
output: []

```

Berikut _template_ pengerjaannya :

```go
func SortHeight(heights []int) []int {
    // kerjakan di sini
}
```

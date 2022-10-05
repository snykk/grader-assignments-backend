# Exercise slice mutidimention 1

## Schedulable days 2

Terdapat sebuah desa bernama "Ruangdesa" yang ingin mengadakan kegiatan bersih-bersih desa. Kepala desa ingin mengundang beberapa orang agar banyak orang yang bersedia buat ikut, namun setiap orang memiliki kegiatan masing-masing. Kepala desa berencana menentukan jadwal bersih-bersih ketika semua orang yang bersedia hadir dalam jadwal kosong.

Buatlah sebuah fungsi yang dapat menentukan tanggal kosong yang sama dari setiap orang yang bersedia hadir untuk bersih-bersih. Mereka akan memberikan jadwal kosong masing-masing.

Data yang diberikan adalah sebuah _slice of_ `int` yang mewakili tanggal-tanggal kosong dari mereka selama sebulan. Karena kita tidak tahu berapa orang yang bersedia ikut, maka setiap data yang diberikan akan dibungkus dalam data _slice_.

Contoh, sebagai berikut:

```txt
data = [[7, 12, 19, 22], [12, 19, 21, 23], [7, 12, 19], [12, 19]]
```

Data di atas berarti terdapat 4 warga desa yang ikut dengan memberikan jadwal kosong masing - masing. Maka hasilnya adalah:

```txt
[12, 19]
```

Berikut _template_ yang dapat digunakan:

```go
func SchedulableDays(villager [][]int) []int {
    // kerjakan di sini
}
```

Contoh _test case_ :

```txt
input: villager = [[7, 12, 19, 22], [12, 19, 21, 23], [7, 12, 19], [12, 19]]
output: [12, 19]

input: date1 = [[1, 2, 3, 4, 5], [2, 3, 4, 5], [2, 3, 4, 10, 11, 12, 15]]
output: [2, 3, 4]

input: date1 = [[1, 2, 3, 4], [5, 6, 7, 8], [10, 11, 12], [21, 22, 23, 24], [25]]
output: []

input: villager = []
output: []
```

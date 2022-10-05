# Exercise array slice 3

## Schedulable days

Pada suatu tempat terdapat 2 orang bernama "Imam" dan "Permana". Mereka berdua ingin merencanakan jadwal untuk berlibur bersama selama mereka memiliki jadwal kosong dari tanggal mereka selama sebulan.

Buatlah sebuah fungsi yang dapat menentukan tanggal kosong yang sama dari kedua orang tersebut agar dapat berlibur bersama. Mereka akan memberikan jadwal kosong masing-masing.

Data yang diberikan adalah sebuah _slice of_ `int` yang mewakili tanggal-tanggal kosong dari mereka. Kembalikan sebuah nilai _slice of_ `int` yang mewakili jadwal kosong yang sama dari kedua orang tersebut.

Contoh, sebagai berikut:

```txt
Imam = [11, 12, 13, 14, 15]
Permana = [5, 10, 12, 13, 20, 21]
```

Dari data di atas, dapat disimpulkan bahwa tanggal yang sama-sama kosong adalah:

```txt
[12, 13]
```

Contoh _test case_ :

```txt
input: date1 = [1, 2, 3, 4], date2 = [3, 4, 5]
output: [3, 4]

input: date1 = [11, 12, 13, 14, 15], date2 = [5, 10, 12, 13, 20, 21]
output: [12, 13]

input: date1 = [2,7, 12, 20, 21, 22], date2 = [1, 3, 6, 10]
output: []

input: date1 = [], date2 = []
output: []
```

Berikut _template_ yang dapat digunakan:

```go
func SchedulableDays(date1 []int, date2 []int) []int {
    // kerjakan di sini
}
```

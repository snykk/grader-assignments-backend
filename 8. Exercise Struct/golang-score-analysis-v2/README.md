# Exercise struct 1

## Score analiysis

Terdapat sekolah yang menampung data nilai seluruh siswa. Data tersebut disimpan ke dalam sebuah data _slice of int_. Buatlah sebuah program yang dapat menambahkan data nilai siswa baru dan menghitung rata-rata nilai seluruh siswa tersebut. Nilai siswa paling kecil adalah `0` dan nilai siswa paling besar adalah `100`.

Terdapat _struct_ `School` yang dapat digunakan untuk menyimpan data ke dalam properti sebagai berikut :

- `Name` bertipe `string` digunakan untuk menyimpan nama sekolah.
- `Address` bertipe `string` digunakan untuk menyimpan alamat sekolah.
- `Grades` bertipe `[]int` digunakan untuk menyimpan data nilai-nilai siswa.

Buatlah sebuah method `AddGrade` yang dapat menambahkan nilai siswa baru ke dalam properti `Grades`. Method tersebut menerima minimal satu atau lebih nilai siswa yang akan ditambahkan ke dalam properti `Grades`. Method tersebut tidak mengembalikan nilai apapun.

Terdapat fungsi bernama Analysis yang menerima sebuah parameter bertipe _struct_ `School` dan mengembalikan data sebagai berikut:

- `Average` bertipe `float64` digunakan untuk menyimpan nilai rata-rata nilai seluruh siswa.
- `Min` bertipe `int` digunakan untuk menyimpan nilai terkecil dari seluruh siswa.
- `Max` bertipe `int` digunakan untuk menyimpan nilai terbesar dari seluruh siswa.

Berikut _template pengerjaannya_ :

```go
type School struct {}

func (s *School) AddGrade(grades ...int) {}

func Analysis(s School) (float64, int, int) {
    // kerjakan di sini
}
```

Contoh test case :

```txt
Analysis
input: School with 'Grades' = [100, 90, 100, 90, 100, 90]
output: 95, 90, 100

input: School with 'Grades' = [100, 100, 100, 100]
output: 100, 100, 100

input: School with 'Grades' = [0, 100, 0, 100]
output: 50, 0, 100

input: School with 'Grades' = []
output: 0, 0, 0

...
```

# Exercise package & import 1

## Advance calculator

### Description

Terdapat sebuah program yang dapat membaca input dari user dan melakukan operasi matematika sederhana. Program ini akan menerima input berupa _string_ yang berisi beberapa angka dan operator matematika dengan dipisah menggunakan spasi `" "`.

program ini hanya dapat membaca 4 operasi matematika yaitu penambahan (simbol : `+`), pengurangan (simbol : `-`), perkalian (simbol : `*`), dan pembagian (simbol : `/`). Sehingga input yang diberikan hanya akan berisi 4 operasi matematika tersebut.

Program ini akan mengembalikan hasil dari operasi matematika yang dilakukan dengan tipe data `float32`.

Perlu diketahui, program ini akan melakukan operasi matematika dari urutan paling kiri ke kanan, sehingga tidak perlu memperhatikan urutan operasi matematika.

Contoh, input sebagai berikut `"3 * 4 / 2 + 10 - 5"`

- Pertama, program akan melakukan perkalian antara angka `3` dan `4` sehingga hasilnya adalah `12`.
- Kedua, hasil perkalian sebelumnya (`12`) akan dibagi dengan angka `2` sehingga hasilnya adalah `6`.
- Ketiga, hasil pembagian sebelumnya (`6`) akan ditambahkan dengan angka `10` sehingga hasilnya adalah `16`.
- Terakhir, hasil penambahan sebelumnya (`16`) akan dikurangi dengan angka `5` sehingga hasilnya adalah `11`.

Maka program akan mengembalikan nilai `11` sebagai hasil dari operasi matematika.

Terdapat _package_ `internal` yang harus digunakan untuk membantu menyelesaikan masalah ini. _Package_ tersebut akan berisi _struct_ `Calculator` yang dapat menjalankan 4 operasi sederhana yaitu `Add`, `Subtract`, `Multiply`, `Divide`. Dan dapat memanggil fungsi `Result` untuk mendapatkan hasil dari operasi matematika yang telah dilakukan.

### Constrains

- data yang diberikan hanya berisi angka dan operator matematika dengan angka selalu positif (tidak ada angka negatif).

### Test Case

Contoh _test case_ :

```txt
calculate: "3 * 4 / 2 + 10 - 5"
output: 11.0

calculate: "10 / 4 + 100"
output: 102.5

calculate: "10 + 10 + 10 + 10 + 12 + 12 + 12 + 12"
output: 72.0

calculate: "10"
output: 10.0

calculate: ""
output: 0.0
```

### Template

Silahkan lengkapi code yang ada _package_ `internal` dan lakukan _import package_ tersebut ke dalam _package_ `main` untuk menyelesaikan masalah ini.

```go
func AdvanceCalculator(calculate string) float32 {
    // kerjakan di sini
}
```

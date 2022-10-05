# Assignment 1

## Password Generator

Andi selalu kebingungan ketika membuat _password_ untuk akun-akun yang baru dia buat. Ia ingin membuat sebuah program yang dapat membuat _password_ dengan algoritma tertentu sebagai berikut:

- Andi memasukkan sebuah _string_ yang akan menjadi dasar dalam membuat _password_.
- Pertama, Andi akan mengubah urutan karakter _string_ tersebut menjadi terbalik.
- Kemudian, Andi akan mengubah setiap karakter vokal menjadi karakter vokal berikutnya serta dibalik kapitalisasinya.
  - a menjadi E, e menjadi I, i menjadi O, o menjadi U, u menjadi A
  - A menjadi e, E menjadi i, I menjadi o, O menjadi u, U menjadi a
- Andi juga mengubah semua karakter huruf besar selain vokal menjadi huruf kecil, dan sebaliknya.
- Jika bertemu dengan spasi, maka spasi tersebut dihilangkan, dan karakter lainnya akan tetap (seperti simbol dan angka).

Andi berharap program ini dapat membantu dia membuat _password_ tanpa kebingungan untuk mengingatnya, karena alogritmanya yang terbilang sederhana.

Contoh, Andi memasukkan _string_ dengan nilai `"Semangat Pagi"` maka alogritma yang kita buat akan melakukan:

- Membalik urutan karakter `Semangat Pagi 12!#` => `#!21 !igaP tagnameS`
- Mengubah vokal menjadi vokal berikutnya `#!21 igaP tagnameS` => `#!21 ogeP tegnemiS`
- Mengubah huruf besar menjadi kecil dan sebaliknya `#!21 ogeP tegnemiS` => `#!21 OGEp TEGNEMIs`
- Menghilangkan spasi `#!21 OGEp TEGNEMIs` => `#!21OGEpTEGNEMIs`

Andi juga ingin memastikan apakah _password_ yang dibuat oleh program tersebut sudah aman atau belum. Dia ingin membuat _status_ dari _password_ tersebut sebagai berikut:

- `"sangat lemah"`, jika memiliki kriteria :
  - Panjang _password_ kurang dari 7 karakter
  - Contoh : `"admin"`, `"123ab"` , `""` , `"AdM1n"`, `"ADM!@"`

- `"lemah"`, jika memiliki kriteria :
  - Panjang _password_ minimal 7 karakter
  - Mengandung hanya karakter angka, atau huruf (baik besar atau kecil), atau kombinasi dari keduanya
  - Contoh : `"Admin123"`, `AdM1N342`, `"12345678"`, `"abcdefghasdas"`

- `"sedang"`, jika memiliki kriteria :
  - Panjang _password_ minimal 7 karakter
  - Mengandung minimal 1 karakter simbol
  - Mengandung kombinasi angka (baik besar atau kecil) dan simbol, atau kombinasi huruf dan simbol, atau ketiganya (huruf, angka dan simbol).
  - Contoh : `"Adm1n!@"`, `"AdM1N342!"`, `"12345678!"`, `"abcdefgh!"`, `"Admin!@#"`

- `"kuat"`, jika memiliki kriteria seperti `"sedang"`, namun memiliki panjang minimal 14 karakter.
  - Contoh: `"Adm1n!@#1234567890"`, `"AdM1N342!@#1234567890"`

Andi akan memencah beberapa _subtask_ untuk mempermudah dalam membuat algoritma di atas.

### Phase 1

Sesuai dengan penjelasan di atas, pada tahap ini kita akan membuat program yang dapat mengubah masukan _string_ menjadi terbalik. Program ini akan menerima masukan berupa tipe data _string_ dengan mengembalikan data string. Kerjakan di dalam fungsi `Reverse` untuk membalik posisi karakter dari _string_ yang diberikan.

```go
func Reverse(str string) string {
    // kerjakan di sini
}
```

### Phase 2

Berikutnya, kita akan membuat program yang dibutuhkan untuk melakukan generate _password_ dengan fungsi `Generate`. Program ini akan menerima masukan berupa tipe data _string_ yang akan mengembalikan nilai _string_ yang merupakan hasil dari generate _password_.

Kita akan memakai fungsi `Reverse` untuk membalik _string_ terlebih dahulu, setelah itu buatlah fungsi ini agar dapat melakukan _generate_ berdasarkan algoritma yang sudah dijelaskan di atas.

```go
func Generate(str string) string {
    str = Reverse(str)
    // kerjakan di sini
}
```

### Phase 3

Kita akan membuat program untuk mengecek tingkat keamanan dari _password_ yang dibuat oleh program tersebut. Terdapat fungsi `CheckPassword` yang akan menerima masukan berupa tipe data _string_ yang akan mengembalikan nilai _string_ yang merupakan hasil dari pengecekan _password_ tersebut.

Buatlah fungsi ini untuk mengecek tingkat keamanan dari hasil _generate_ di fungsi `Generate` dengan kriteria yang sudah dijelaskan di atas.

```go
func CheckPassword(str string) string {
    // kerjakan di sini
}
```

Terakhir Andi sudah menyiapkan fungsi bernama `PasswordGenerator` yang akan menerima masukan berupa _string_ dan mengembalikan 2 nilai yaitu:

- _string_ yang merupakan hasil dari generate _password_.
- _string_ yang merupakan hasil dari pengecekan keamanan _password_ tersebut.

```go
func PasswordGenerator(base string) (string, string) {
    // kerjakan di sini
}
```

Contoh _test case_ :

```go
PasswordGenerator
Input: "Semangat Pagi"
Output: "OGEpTEGNEMIs", "lemah"

Input: "Adm1n!@" 
Output: "@!N1MDe", "sedang" 

Input: ""
Output: "", "sangat lemah"

Input: "Admin"
Output: "NOMDe", "lemah"
```

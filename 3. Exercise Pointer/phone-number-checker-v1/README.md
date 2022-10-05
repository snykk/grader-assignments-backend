# Exercise pointer 2

## Phone number checker

Buatlah sebuah program untuk melakukan pengecekan dari nomer telepon yang diberikan, kita diminta untuk mengecek provider dari nomer tersebut, dan mengecek apakah nomer tersebut valid atau tidak.

Kita menggunakan pointer untuk menangkap hasil provider dari nomer tersebut.

Syarat utama untuk nomer telepon yang _valid_ adalah:

- nomer harus diawali dengan kode negara `62` dan angka `8`, atau tanpa menggunakan kode negara dan menggunakan angka awal `08`.
- panjang nomer harus minimal 10 angka jika tidak menggunakan kode negara, dan 11 angka jika menggunakan kode negara.

Jika tidak memenuhi kedua syarat utama di atas, maka kita akan mengisi pointer dengan `invalid`.

Berikut kode 4 nomer awal, untuk mengetahui provider dari nomer tersebut:

- 0811 - 0815 / 62811 - 62815 => Telkomsel
- 0816 - 0819 / 62816 - 62819 => Indosat
- 0821 - 0823 / 62821 - 62823 => XL
- 0827 - 0829 / 62827 - 62829 => Tri
- 0852 - 0853 / 62852 - 62853 => AS
- 0881 - 0888 / 62881 - 62888 => Smartfren

Provider di atas adalah provider yang paling banyak digunakan oleh pengguna nomer telepon. Namun, jika kode dari nomer tersebut tidak ditemukan, maka kita akan mengisi pointer dengan `invalid`.

Contoh _test case_:

```txt
input: number = 081211111111
output: result = Telkomsel

input: number = 08193456123
output: result = Indosat

input: number = 628523456789
output: result = AS

input: number = 081234567
output: result = invalid

input: number = 08222
output: result = invalid
```

```go
func PhoneNumberChecker(number string, result *string) {
    // kerjakan di sini
}

func main() {
    // bisa digunakan untuk pengujian test case
    var number = ""
    var result string

    PhoneNumberChecker(number, &result)
    fmt.Println(result)
}
```

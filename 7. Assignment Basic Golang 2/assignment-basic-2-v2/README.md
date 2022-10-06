# Assigment basic 2

## Rules

- Dilarang mengubah ini file pada _test code_ di `main_test.go` dan `golang_suite_test.go`.
- Dilarang menganti nama fungsi atau mengubah parameter dan return yang diberikan.

## Salary Overview

Terdapat perusahaan bernama "ruangbekerja" yang ingin membuat sistem sederhana tentang perhitungan gaji. Perusahaan ingin mengetahui gaji yang diterima oleh setiap karyawan berdasarkan jarak tanggal yang diberikan. Data karyawan yang hadir akan dicatat setiap hari berdasarkan nama karyawan masing masing.

Setiap karyawan yang hadir akan mendapatkan gaji sebesar `"Rp 50.000"` per hari. Jika karyawan tersebut tidak hadir maka tidak akan mendapatkan gaji.

Untuk mempermudah penyimpanan data di sistem, maka data kehadiran karyawan akan disimpan ke dalam data berformat _slice of string_ yang berisi nama karyawan yang hadir pada hari tersebut. Dan karena data akan menyimpan lebih dari 1 hari, maka data tersebut akan disimpan ke dalam _slice_ multidimensi.

```go
var data = [][]string{
    {"Andi", "Imam", "Eddy", "Deny"},
    {"Imam", "Eddy"},
    {"Deny"}
}
```

Dari data di atas dapat disimpulkan bahwa pada hari pertama, karyawan yang hadir terdapat 4 orang, bernama Andi, Imam, Eddy, dan Deny. Pada hari ke-2, yang hadir terdapat 2 orang, bernama Imam dan Eddy. Dan pada hari ke-3, hanya karyawan bernama Deny yang hadir.

Kemudian perusahaan ingin mengetahui data kehadiran dari tanggal tertentu. Perusahaan akan menginput data dari tanggal tertentu sampai dengan tanggal tertentu. Data tanggal yang diinput akan berformat `"[start-day] [start-month] - [end-day] [end-month] [year]"`

Contoh `21 january - 25 January 2021` maka data yang akan ditampilkan adalah data kehadiran karyawan pada tanggal 21 sampai dengan 25 januari 2021.

Tedapat beberapa aturan dalam menghitung data tanggal, agar tidak terjadi kesalahan :

- Penulisan bulan menggunakan bahasa inggris, dengan awalan huruf kapital (contoh : `January`, `March`).
- Tanggal hanya bisa di cari dalam jarak 1 tahun tersebut (misalnya 21 januari - 25 januari 2021), tidak bisa mencari tanggal di tahun lain (misalnya 25 December 2021 - 25 januari 2022)
- Hari dalam 1 bulan adalah 30 hari, kecuali untuk bulan `"Februari"` yang hanya 28 hari. Jika tahun tersebut adalah tahun kabisat, maka bulan `"February"` akan menjadi 29 hari.

## Phase 1

Pertama yang perlu dilakukan adalah menghitung jarak hari dari tanggal yang dimasukkan. Jarak hari tersebut akan digunakan untuk mengambil data kehadiran karyawan pada tanggal tersebut.

Buatlah fungsi `GetDayDifference` yang akan menghitung jarak hari dari tanggal yang dimasukkan. Fungsi ini akan menerima 1 parameter bertipe `string` yaitu jarak tanggal yang sebelumnya sudah dijelaskan. Fungsi ini akan mengembalikan nilai bertipe `int` berupa jarak tanggal yang mewakili jumlah hari.

Jika kita memasukkan data seperti berikut `"25 January - 30 January 2021"`, maka fungsi ini akan mengembalikan nilai `6` karena jarak tanggal tersebut adalah 6 hari. Contoh lain seperti `"25 February - 10 March 2020"`, maka fungsi ini akan mengembalikan nilai `15` karena jarak tanggal tersebut adalah 15 hari dan tahun tersebut adalah kabisat (25 - 29 Feb + 1 - 10 March).

```go
func GetDayDifference(dateRange string) int {
    // kerjakan di sini
}
```

## Phase 2

Saatnya kita membuat rekap data karyawan dengan menggunakan data yang sudah dijelaskan sebelumnya. Karena data yang dimasukkan tidak memiliki tanggal, maka data tersebut akan mengikuti dari tanggal awal yang dimasukkan. Jika kita memasukkan data `"25 January - 30 January 2021"`, maka data yang akan ditampilkan adalah data kehadiran karyawan dimulai dari tanggal 25 januari 2021.

Buatlah fungsi `GetSalary` yang akan mengembalikan data karyawan yang hadir pada tanggal yang dimasukkan. Fungsi ini akan menerima 2 parameter, yaitu:

- Data yang kehadiran karyawan bertipe `[][]string`
- Data jarak hari yang dimasukkan bertipe `int`

Fungsi ini akan mengembalikan nilai bertipe `map[string]string` yang berisi nama karyawan dan gaji yang diterima berformat mata uang rupiah (contoh: `"Rp 100.000"`).

```go
func GetSalary(data [][]string, dateRange int) map[string]string {
    // kerjakan di sini
}
```

Jika kita memasukkan jarak hari adalah `3`, dan data sebagai berikut:

```txt
data = [
    ["Andi", "Imam", "Eddy", "Deny"],
    ["Andi", "Imam"],
    ["Imam", "Deny"],
    ["Andi", "Deny"]
]
```

Maka kita cukup mengambil data kehadiran karyawan pada hari ke-1, ke-2, dan ke-3. Andi hadir di hari ke 1 dan 2, maka gaji yang diterima adalah `2 x 50.000 = Rp 100.000`. Imam hadir dalam 3 hari, maka gaji yang diterima adalah `3 x 50.000 = Rp 150.000`. Eddy cuma hadir di hari ke 1, maka gaji yang diterima adalah `Rp 50.000`.Terakhir. Deny hadir di hari ke 1 dan ke 3, maka gaji yang diterima adalah `2 x 50.000 = Rp 100.000`.

Jadi, data yang akan ditampilkan adalah sebagai berikut:

```txt
{
    "Andi": "Rp 100.000",
    "Imam": "Rp 150.000",
    "Eddy": "Rp 50.000",
    "Deny": "Rp 100.000"
}
```

### Phase 2.1 OPTIONAL

Kalian dapat membuat fungsi pembantu untuk mengubah data total gaji yang masih bertipe number (`int` / `float`) menjadi format mata uang Rupiah `"Rp xxx.xxx.xxx"` dengan membuat fungsi `FormatRupiah` yang sudah disediakan.

```go
func FormatRupiah(number int) string {
    // kerjakan di sini
}
```

Tapi kalian juga bisa mengerjakannya tanpa fungsi pembantu tersebut, dengan mengubah formatnya langsung di fungsi `GetSalaryOverview`.

## Phase 3

Terakhir Kita akan menggabungkan kedua fungsi yang sudah dibuat sebelumnya. Buatlah fungsi `GetSalaryOverview` yang akan mengembalikan data karyawan yang hadir pada tanggal yang dimasukkan. Fungsi ini akan menerima 2 parameter, yaitu:

- Data yang kehadiran karyawan bertipe `[][]string`
- Data jarak tanggal yang dimasukkan bertipe `string`

Fungsi ini akan mengembalikan nilai bertipe `map[string]string` yang berisi nama karyawan dan gaji yang diterima berformat mata uang rupiah (contoh: `"Rp 100.000"`).

```go
func GetSalaryOverview(data [][]string, dateRange string) map[string]string {
    // kerjakan di sini
}
```

Jika kita memasukkan data `"25 January - 30 January 2021"`, dan data sebagai berikut:

```txt
data = [
    ["Andi", "Imam", "Eddy", "Deny"],
    ["Andi", "Imam"],
    ["Imam", "Deny"],
    ["Andi", "Deny"]
]
```

Harusnya output akan sama seperti pada phase 2.

```txt
{
    "Andi": "Rp 100.000",
    "Imam": "Rp 150.000",
    "Eddy": "Rp 50.000",
    "Deny": "Rp 100.000"
}
```

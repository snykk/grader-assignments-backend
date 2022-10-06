# Execise Interface 3

## Change to Standart Time

Imam tertarik dengan konversi waktu 24 jam menjadi waktu standart (AM / PM ). Oleh karena itu Imam membuat sebuah program untuk mengubah waktu 24 jam menjadi waktu standart tersebut. Terdapat fungsi `ChangeToStandartTime` yang akan mengembalikan tipe data _string_ yang berisi waktu standart dengan format `"HH:MM [AM/PM]"`.

Untuk simplicity, jam `00:00` akan ditampilkan sebagai `00:00 AM`, jam `12:00` sebagai `12:00 PM`, dan semua waktu setelah jam 12 siang adalah `PM`

Imam ingin program ini lebih berguna, sehingga Imam ingin fungsi ini dapat menerima beberapa parameter yang berbeda menggunakan tipe data `interface{}`.

Agar terhindar dari input yang aneh-aneh, Imam membuat beberapa aturan untuk parameter yang bisa diterima adalah sebagai berikut:

- Dapat diisi dengan tipe data `string` yang berformat `"HH:MM"`, ex : `"04:00"`
- Dapat diisi dengan tipe data `slice` yang berisi tipe data `int` dengan 2 data, index ke 0 sebagai `hour` dan index ke 1 sebagai `minute`, ex: `[10, 30]`
- Dapat diisi dengan tipe data `map` yang harus berisi 2 key yaitu `hour` dan `minute` dengan key bertipe `string` dan value bertipe `int`, ex : map["hour": 10, "minute": 30]
- Terakhir, dapat diisi dengan tipe data _struct_ `Time` yang berisi _property_ `hour` bertipe `int` dan `minute` bertipe `int`

  ```go
  type Time struct {
      Hour   int
      Minute int
  }
  ```

- Selain dari itu, maka kembalikan nilai `"Invalid input"`

Contoh, berikut data input:

```txt
"16:00" -> "04:00 PM"
[23, 27] -> "11:27 PM"
map[hour: 11, minute: 30] -> "11:30 AM"
Time{Hour: 9, Minute: 29} -> "09:29 AM"
```

Berikut _template_ pengerjaannya :

```go
type Time struct {
    Hour int
    Minute int
}

func ChangeToStandartTime(time interface{}) string {
    // kerjakan di sini
}
```

Untuk versi Golang 1.18 ke atas, bisa menggunakan tipe `any`

```go
func ChangeToStandartTime(time any) string {
    // kerjakan di sini
}
```

Contoh _test case_ :

```txt
input: time = "04:00"
output: 04:00 AM"

input: time = [4, 15]
output: "04:15 AM"

input: time = map[hour: 16, minute: 47]
output: "04:47 PM"

input: time = [4]
output: "Invalid input"

input: time = map[minute: 47]
output: "Invalid input"
...
```

> hint: Gunakan check _data type_ dari `interface` di dalam kondisi dengan syntax `interface{}.(type)`.

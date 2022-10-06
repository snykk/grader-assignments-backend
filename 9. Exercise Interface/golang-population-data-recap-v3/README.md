# Exercise interface 2

## Population data recap

Terdapat sebuah data berbentuk _slice of_ `string` yang berisi kumpulan data penduduk.

Setiap data akan berisi nama, umur, alamat, tinggi badan dan status pernikahan dengan format `"[name];[age];[address];[height];[is_married]"`.

Terdapat data yang tidak wajib untuk diisi, yaitu tinggi badan dan status pernikahan. Sehingga data yang masuk akan berbeda-beda, contoh : `["Budi;23;Jakarta;;", "Joko;30;Bandung;;true", "Susi;25;Bogor;165.42;"]`

Buatlah code untuk mengubah data tersebut menjadi lebih informatif menggunakan `map`, dan nilainya disesuaikan berdasarkan aturan berikut:

- Terdapat _property_ `name` dengan tipe data `string`
- Terdapat _property_ `age` dengan tipe data `int`
- Terdapat _property_ `address` dengan tipe data `string`
- Jika data `height` diisi, maka terdapat _property_ `height` dengan tipe data `float64`
- Jika data `is_married` diisi, maka terdapat _property_ `isMarried` dengan tipe data `bool`

Berikut _template_ pengerjaanya :

```go
func PopulationData(data []string) []map[string]interface{} {
    // kerjakan di sini
}
```

Untuk versi Golang 1.18 ke atas, bisa menggunakan tipe `any`

```go
func PopulationData(data []string) []map[string]any {
    // kerjakan di sini
}
```

Contoh _test case_ :

```txt
case 1
input: data = ["Budi;23;Jakarta;;", "Joko;30;Bandung;;true", "Susi;25;Bogor;165.42;"]
output: [
    map[name: "Budi", age: 23, address: "jakarta"],
    map[name: "Joko", age: 30, address: "Bandung", isMarried: true],
    map[name: "Susi", age: 25, address: "Bogor", height: 165.42]
]

case 2
input: data = ["Jaka;25;Jakarta;false;170.1", "Anggi;24;Bandung;;"]
output: [
    map[name: "Jaka", age: 25, address: "Jakarta", height: 170.1],
    map[name: "Anggi", age: 24, address: "Bandung"]
]

case 3
input: data = []
output: []
...
```

> hint: Gunakan `strconv.ParseFloat` untuk mengubah `string` ke `float`.

# Assignment map 4

## Formatting Data

Kita diminta memperbaiki format data yang ada di dalam `slice` yang berisi `string`. Untuk saat ini data di dalam `slice` menggunakan format `[header]-[index]-[position]-[value]`.

Terdapat fungsi bernama `formatData` yang berisi parameter `data` dengan tipe data _slice of_ `string`. Buatlah sebuah code yang dapat mengubah format data menjadi sebuah `map` dengan tipe data `map[string][]string`.

Berikut adalah contoh input dan output dari fungsi `formatData`:

```txt
input: data = ["account-0-first-John", "account-0-last-Doe", "account-1-first-Jane", "account-1-last-Doe"]

output: map[account: ["John Doe", "Jane Doe"]]

input: data = ["account-0-first-John", "account-0-last-Doe", "account-1-first-Jane", "account-1-last-Doe", "address-0-first-Jaksel", "address-0-last-Jakarta", "address-1-first-Bandung", "address-1-last-Jabar"]

output: map[account: ["John Doe", "Jane Doe"], address: ["Jaksel Jakarta", "Bandung Jabar"]]
```

Berikut _template_ pengerjaannya:

```go
func changeOutput(data []string) map[string][]string {
    
    // kerjakan di sini
}
```

# Assignment Map 1

## map to slice

Buatlah sebuah code yang dapat mengubah data `map` berisi key dengan tipe data `string` dan value dengan tipe data `string` menjadi list data di `slice`. Setiap data di `map` akan dibungkus menjadi 1 `slice`.

Disediakan fungsi `mapToSlice` yang berisi parameter `mapData` dengan tipe `map[string]string`, dan kerjakan di dalam fungsi tersebut dengan me-_return_ sebuah `slice` dengan tipe data `[]string`.

Berikut adalah contoh input dan output dari fungsi `mapToSlice`:

```txt
input: map[hello: "world", John: "Doe" , age:"14"]
output: [["hello", "world"], ["b" ,"2"], ["c", "3"]]

input: mapData = map[foo: "33", bar: "44", baz: "55"]
output: [["foo", "33"], ["bar", "44"], ["baz", "55"]]

input: mapData = map[]
output: []

...
```

Berikut _template_ pengerjaannya:

```go
func mapToSlice(mapData map[string]string) [][]string {
    
    // kerjakan di sini
}
```

# Exercise pointer 1

## Slurred Talk

Steven adalah seorang yang cadel, dia seringkali mengalami kesulitan untuk mengucapkan beberapa huruf. Huruf yang susah dieja Steven adalah `S`, `R` dan `Z`. Karena itu, Steven seringkali menggantinya dengan huruf tersebut menjadi huruf `L`.

Buatlah sebuah fungsi yang dapat mengubah beberapa huruf yang susah dieja Steven menjadi huruf `L` agar kita dapat mengikuti apa yang Steven ingin katakan. Fungsi ini menerima sebuah pointer bertipe data `string` yang langsung kita ubah menjadi huruf `L` jika huruf tersebut adalah `S`, `R` atau `Z` (baik itu huruf besar atau huruf kecil).

Contoh _test case_ :

```txt
input: words = "Steven"
output: words = "Lteven"

input: words = "Saya Steven"
output: words = "Laya Lteven"

input: words = "Saya Steven, saya suka menggoreng telur dan suka hewan zebra"
output: words = "Laya Lteven, laya luka menggoleng telul dan luka hewan lebra"

input: words = ""
output: words = ""
```

```go
func SlurredTalk(words *string) {
    // kerjakan di sini
}

func main() {
    // bisa dicoba untuk pengujian test case
    var words string = ""
    SlurredTalk(&words)
    fmt.Println(words)
}
```

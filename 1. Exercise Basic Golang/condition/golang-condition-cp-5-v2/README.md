# Exercise Condition

## Ticket Playground

### Description

Terdapat wahana bermain yang memberikan tarif tiket anak-anak sesuai dengan tinggi atau umur (tergantung mana yang lebih besar), dengan ketentuan sebagai berikut:

- Jika Umur di bawah 5 tahun maka tidak dapat membeli tiket wahana dengan menampilkan `-1` (repersentasi dari `error`)
- Jika anak umur 5-7 tahun atau tinggi lebih dari 120 cm maka akan dikenakan tarif Rp 15.000.
- Jika anak umur 8-9 tahun atau tinggi lebih dari 135 cm maka akan dikenakan tarif Rp 25.000.
- Jika anak umur 10-11 tahun atau tinggi lebih dari 150 cm maka akan dikenakan tarif Rp 40.000.
- Jika anak umur 12 tahun atau tinggi lebih dari 160 cm, akan dikenakan tarif Rp 60.000.
- Jika di atas 12 tahun, akan mendapat tiket khusus Remaja yaitu sebesar Rp 100.000

Contoh, jika terdapat anak dengan umur 10 tahun dan tinggi 151 cm maka akan dikenakan tarif tiket sebesar Rp 40.000 (karena tinggi lebih besar dari 150 cm). Atau jika terdapat anak dengan umur 12 tahun dan tinggi 135 cm maka akan dikenakan tarif tiket sebesar Rp 60.000 (karena umur lebih sama 12 tahun).

Terdapat fungsi bernama `TicketPlayground` yang menerima 2 parameter, yaitu `height` bertipe `int` dan `age` bertipe `int`. Buatlah sebuah _code_ yang dapat mengembalikan _output_ berupa total tarif yang harus dibayar dengan kondisi sesuai ketentuan di atas. Cukup tampilkan _output_ dalam tipe data `int`.

### Constraints

- Input `height` akan berupa bilangan bulat positif, dengan rentang 0 - 200
- Input `age` akan berupa bilangan bulat positif, dengan rentang 0 - 20

### Test Case Examples

#### Test Case 1

**Input**:

```txt
height = 160, age = 11
```

**Expected Output / Behavior**:

```txt
40000
```

**Explanation**:

Karena tinggi 160 cm dan umur 11 tahun, maka tarif yang harus dibayar adalah Rp. 40.000. Tarif itu dikenakan karena tinggi anak tersebut lebih besar dari 150 cm.

#### Test Case 2

**Input**:

```txt
height = 165, age = 10
```

**Expected Output / Behavior**:

```txt
60000
```

**Explanation**:

Karena tinggi 165 cm dan umur 10 tahun, maka tarif yang harus dibayar adalah Rp. 60.000. Tarif itu dikenakan karena tinggi anak tersebut lebih besar dari 12 tahun.

#### Test Case 3

**Input**:

```txt
height = 140, age = 13
```

**Expected Output / Behavior**:

```txt
100000
```

**Explanation**:

Karena umur lebih dari 12 tahun, maka akan mendapat tiket khusus Remaja yaitu sebesar Rp. 100.000.

#### Test Case 4

**Input**:

```txt
height = 85, age = 4
```

**Expected Output / Behavior**:

```txt
-1
```

**Explanation**:

Karena umur di bawah 5 tahun, maka tidak dapat membeli tiket wahana dengan menampilkan `-1` (repersentasi dari `error`).

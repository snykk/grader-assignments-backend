# Exercise Function

## Find Shortest Name

### Description

Terdapat sebuah data berupa nama-nama yang terdapat di kelas "X". Data tersebut dapat dipisah dengan 3 cara, yaitu spasi `" "`, koma `","`, dan titik koma `";"`. Kita ingin mencari nama terpendek dari data tersebut.

Gunakan fungsi `FindShortestName` agar dapat mencari nama terpendek dalam data tersebut. Terdapat parameter `names` bertipe `string` yang akan berisi beberapa nama orang di kelas "X". Kembalikan nilai berubah `string` yang berisi nama terpendek. Jika terdapat lebih dari 1 nama terpendek, maka kembalikan nama terpendek yang namanya berada di urutan awal berdasarkan urutan abjad.

### Constraints

- Input `names` akan selalu berupa deret nama orang yang terdiri dari huruf a-z, A-Z, spasi `" "`, koma `","`, dan titik koma `";"`, dengan panjang karakter 2 <= `names` <= 100.000

### Test Case Examples

#### Test Case 1

**Input**:

```txt
names = "Hanif Joko Tio Andi Budi Caca Hamdan"
```

**Expected Output / Behavior**:

```txt
"Tio"
```

**Explanation**:

Output yang diharapkan adalah `"Tio"` karena dari semua nama yang dimasukkan, nama terpendek adalah `Tio` dengan panjang karakter 3 huruf.

```txt
Hanif -> 5 huruf
Joko -> 4 huruf
Tio -> 3 huruf -> nama terpendek
Andi -> 4 huruf
Budi -> 4 huruf
Caca -> 4 huruf
Hamdan -> 6 huruf
```

#### Test Case 2

**Input**:

```txt
"Budi;Tia;Tio"
```

**Expected Output / Behavior**:

```txt
"Tia"
```

**Explanation**:

Output yang diharapkan adalah `"Tia"` karena dari semua nama yang dimasukkan ( `Budi`, `Tia`, dan `Tio`), terdapat 2 nama dengan nama terpendek, yaitu `Tia` dan `Tio`. Namun, karena `Tia` memiliki nama dengan urutan abjad pertama dibandingkan nama lainnya, maka kembalikan nama tersebut.

```txt
Budi -> 4 huruf
Tia -> 3 huruf -> nama dengan urutan abjad awal
Tio -> 3 huruf
```

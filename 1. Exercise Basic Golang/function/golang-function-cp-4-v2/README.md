# Exercise Function

## Find Similar Data

### Description

Andi ingin mencari data yang mirip dengan data yang diberikan. Katakan data yang dicari adalah `"mobil"`, kemudian terdapat sekumpulan data sebagai berikut `"mobil APV", "mobil Avanza", "motor matic", "motor gede"`. Maka data yang mirip dengan `"mobil"` adalah `"mobil APV"` dan `"mobil Avanza"`.

Terdapat sebuah sistem pencarian data yang memiliki fungsi bernama `FindSimilarData` dengan parameter `input` bertipe `string` yang berisi data yang akan dicari dan terdapat parameter ke dua `data` yang akan berisi data dinamis (`...string`). Kembalikan sekumpulan data yang mirip dengan data yang dicari ke dalam bentuk `string` yang dipisahkan dengan koma (`,`).

> hint: Menggunakan `strings.Contains`

### Constraints

- Input `input` akan selalu berupa data yang terdiri dari huruf kecil, besar dan spasi, dengan panjang karakter 1 <= `input` <= 100.
- Input `data` akan selalu berupa sekumpulan data yang terdiri dari huruf kecil, besar dan spasi, dengan panjang karakter 1 <= `data` <= 100.

### Test Case Examples

#### Test Case 1

**Input**:

```txt
input = "mobil", data = ["mobil APV", "mobil Avanza", "motor matic", "motor gede"]
```

**Expected Output / Behavior**:

```txt
"mobil APV,mobil Avanza"
```

**Explanation**:

Output yang diharapkan adalah `"mobil APV,mobil Avanza"` karena data yang mirip / mengandung kata `"mobil"` adalah `"mobil APV"` dan `"mobil Avanza"`. Kemudian kita gabungkan menjadi satu dengan dipisah menggunakan koma (`,`).

#### Test Case 2

**Input**:

```txt
input = "motor", data = ["mobil APV", "mobil Avanza", "motor matic", "motor gede", "iphone 14", "iphone 13", "iphone 12", "pengering baju", "Kemeja flannel"]
```

**Expected Output / Behavior**:

```txt
"motor matic,motor gede"
```

**Explanation**:

Output yang diharapkan adalah `"motor matic,motor gede"` karena data yang mirip / mengandung kata `"motor"` adalah `"motor matic"` dan `"motor gede"`. Kemudian kita gabungkan menjadi satu dengan dipisah menggunakan koma (`,`).

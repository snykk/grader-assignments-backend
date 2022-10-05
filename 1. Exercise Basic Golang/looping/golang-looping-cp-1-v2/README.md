# Exercise Looping

## Counting Odd Even

### Description

Buatlah sebuah _code_ yang dapat menjumlahkan angka dari 1 sampai ke angka `n` dengan peningkatan `0.5` dari angka awal.

Terdapat fungsi `CountingNumber` dengan parameter `n` bertipe `int` yang akan digunakan untuk menentukan batasan angka yang akan dijumlahkan. Kembalikan sebuah nilai bertipe `float64` yang merupakan hasil penjumlahan dari angka 1 sampai ke angka `n` dengan peningkatan `0.5` dari angka awal.

Contoh, jika `n` diisi dengan `5`, maka hasilnya adalah `1 + 1.5 + 2 + 2.5 + 3 + 3.5 + 4 + 4.5 + 5 = 27`.

### Constraints

- Input `n` akan selalu bilangan bulat positif, dengan rentang nilai 1 <= `n` <= 1.000.000

### Test Case Examples

#### Test Case 1

**Input**:

```txt
n = 10
```

**Expected Output / Behavior**:

```txt
104.5
```

**Explanation**:

Output yang diharapkan adalah `55` karena `1 + 1.5 + 2 + 2.5 + 3 + 3.5 + 4 + 4.5 + ... + 10 = 104.5`.

#### Test Case 2

**Input**:

```txt
n = 100
```

**Expected Output / Behavior**:

```txt
10049.5
```

**Explanation**:

Output yang diharapkan adalah `5050` karena `1 + 1.5 + 2 + 2.5 + 3 + 3.5 + 4 + ... + 100 = 10049.5`.

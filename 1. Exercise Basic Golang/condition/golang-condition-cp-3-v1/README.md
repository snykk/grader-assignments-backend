# Exercise Condition

## Predicate score

### Description

Tedapat sekolah yang memberikan predikat pada murid berdasarkan nilai rata-rata yang didapat :

- Jika nilai yang didapat adalah 100, murid mendapat predikat 'Sempurna'
- Jika nilai yang didapat 90 keatas, murid mendapat predikat 'Sangat Baik'
- Jika nilai yang didapat 80 keatas, murid mendapat predikat 'Baik'
- Jika nilai yang didapat 70 keatas, murid mendapat predikat 'Cukup'
- Jika nilai yang didapat 60 keatas, murid mendapat predikat 'Kurang'
- Jika nilai yang didapat kurang dari 60, murid mendapat predikat 'Sangat kurang'

Murid mendapat 4 nilai dari 4 tugas yang berbeda, yaitu `math`, `science`, `english`, dan `indonesia`.

Terdapat fungsi bernama `GetPredicate` yang menerima 4 parameter, yaitu `math`, `science`, `english`, dan `indonesia` bertipe `int` yang berisi nilai dari masing-masing mata pelajaran. Fungsi ini akan mengembalikan nilai bertipe `string` yang berisi predikat yang didapat murid.

### Constraints

- Input `math`, `science`, `english`, `indonesia` bertipe `int` dengan jarak nilai 0 <= N <= 100

### Test Case Examples

#### Test Case 1

**Input**:

```txt
math = 50, science = 80, english = 100, indonesia = 60
```

**Expected Output / Behavior**:

```txt
"Cukup"
```

**Explanation**:

Output `"Cukup"` didapat karena nilai rata-rata dari 4 mata pelajaran adalah `72.5`, sehingga memenuhi syarat predikat "Cukup".

#### Test Case 2

**Input**:

```txt
math = 100, science = 100, english = 100, indonesia = 100
```

**Expected Output / Behavior**:

```txt
"Sempurna"
```

**Explanation**:

Output `"Sempurna"` didapat karena nilai rata-rata dari 4 mata pelajaran adalah 100, sehingga memenuhi syarat predikat "Sempurna".

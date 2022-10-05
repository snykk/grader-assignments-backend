# Exercise Condition

## BMI Calculator

### Description

Terdapat code program yang digunakan untuk melakukan kalkulator BMI (Body Mass Index). BMI merupakan salah satu cara untuk mengetahui apakah seseorang termasuk dalam kategori obesitas atau tidak. Kita menggunakan rumus Broca dengan perhitungan sebagai berikut :

- Jika orang tersebut adalah Laki-laki, adalah (_height_ - 100) - ((_height_ - 100) x 10%)
- Jika orang tersebut adalah Perempuan, adalah (_height_ - 100) - ((_height_ - 100) x 15%)

Terdapat fungsi bernama `BMICalculator` yang akan menerima 2 parameter yaitu `gender` bertipe `string` yang akan berisi `"laki-laki"` jika orang tersebut laki-laki dan `"perempuan"` jika orang tersebut perempuan. Parameter kedua adalah `height` bertipe `int` yang berisi tinggi badan orang tersebut. Kembalikan nilai `float64` yang merupakan hasil perhitungan BMI.

### [Optional] Constraints

- Input `gender` akan berisi `"laki-laki"` atau `"perempuan"`
- Input `height` bertipe `int` dengan nilai 130 <= N <= 250

### Test Case Examples

#### Test Case 1

**Input**:

```txt
gender = "laki-laki", height = 170
```

**Expected Output / Behavior**:

```txt
63
```

**Explanation**:

Output adalah `63` karena orang tersebut adalah laki-laki, maka perhitungannya sebagai berikut :

```txt
(170 - 100) - ((170 - 100) x 10%) = 63
```

#### Test Case 2

**Input**:

```txt
gender = "perempuan", height = 165
```

**Expected Output / Behavior**:

```txt
55.25
```

**Explanation**:

Output adalah `55.25` karena orang tersebut adalah laki-laki, maka perhitungannya sebagai berikut :

```txt
(165 - 100) - ((165 - 100) x 15%) = 55.25
```

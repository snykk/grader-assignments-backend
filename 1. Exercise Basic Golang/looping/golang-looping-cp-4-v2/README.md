# Exercise Looping

## Domain email

### Description

Buatlah sebuah _code_ yang digunakan untuk mendapatkan _domain email_ dan _top level domain email_ (TLD) yang digunakan dari alamat _email_ yang diberikan. Terdapat fungsi `EmailInfo` dengan parameter `email` bertipe `string` yang berisi alamat email yang dimasukkan. Terakhir kembalikan _output_ berupa `string` dengan format "Domain: [domain] dan TLD: [TLD]"  (tidak ada spasi sebelum karakter `:`).

Provider di alamat email adalah nama penyedia layanan email, yang disimpan setelah karakter `@` dan sebelum karakter `.`. Sedangkan domain adalah nama domain penyedia layanan email, yang disimpan setelah karakter `.`. Formatnya sebagai berikut `[nama-email]@[domain].[TLD]`

Contoh sebagai berikut, jika email yang dimasukkan adalah `"admin@yahoo.com"`, maka _domain_-nya adalah `"yahoo"` dengan TLD adalah `"com"`. Maka _output_ yang diharapkan adalah `"Domain: yahoo dan TLD: com"`.

### Constraints

- Input `email` akan selalu berupa _string_ yang terdiri dari karakter huruf, angka, titik (`.`), dan tanda `@`, dengan panjang karakter 1 <= `email` <= 1000

### Test Case Examples

#### Test Case 1

**Input**:

```txt
email = "admin@yahoo.com"
```

**Expected Output / Behavior**:

```txt
"Domain: yahoo dan TLD: com"
```

**Explanation**:

Kita mendapatkan domainnya adalah `"yahoo"` dengan TLD `"com"`. Maka _output_ yang diharapkan adalah `"Domain: yahoo dan TLD: com"`.

#### Test Case 2

**Input**:

```txt
ptmencaricintasejati@gmail.co.id
```

**Expected Output / Behavior**:

```txt
"Domain: gmail dan TLD: co.id"
```

**Explanation**:

Kita mendapatkan domainnya adalah `"gmail"` dengan TLD `"co.id"`. Maka _output_ yang diharapkan adalah `"Domain: gmail dan TLD: co.id"`.

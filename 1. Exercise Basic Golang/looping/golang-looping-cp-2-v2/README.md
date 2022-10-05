# Exercise Looping

## Reverse String

### Description

Buatlah sebuah _code_ sederhana untuk membalik sekumpulan karakter dari _string_ yang diberikan. Setiap pembalikan antar huruf akan ditambah jeda menggunakan karakter `"_"`, namun ini tidak berlaku jika terdapat spasi. Terdapat fungsi `ReverseString` dengan parameter `str` bertipe `string` yang akan menjadi _input_ data sekumpulan karakter yang akan dibalik.

Contoh, jika `str` diisi dengan `"Hello World"`, maka hasilnya adalah `"d_l_r_o_W o_l_l_e_H"`.

### Constraints

- Input `str` akan selalu berupa _string_ yang hanya terdiri dari karakter huruf dan spasi (**tidak ada** karakter simbol seperti `!@#$%^&*()-+=.,`), dengan panjang karakter 1 <= `str` <= 100.000

### Test Case Examples

#### Test Case 1

**Input**:

```txt
str = "Hello World"
```

**Expected Output / Behavior**:

```txt
"d_l_r_o_W o_l_l_e_H"
```

**Explanation**:

Output yang diharapkan adalah `"d_l_r_o_W o_l_l_e_H"`, karena `"Hello World"` jika dibalik akan menjadi `"dlroW olleH"` dan kemudian kita tambahkan jeda `"_"` antar hurufnya.

#### Test Case 2

**Input**:

```txt
str = "I am a student"
```

**Expected Output / Behavior**:

```txt
"t_n_e_d_u_t_s a m_a I"
```

**Explanation**:

Output yang diharapkan adalah `"t_n_e_d_u_t_s a m_a I"`, karena `"I am a student"` jika dibalik akan menjadi `"tneduts a ma I"`, kemudian kita tambahkan jeda `"_"` antar hurufnya.

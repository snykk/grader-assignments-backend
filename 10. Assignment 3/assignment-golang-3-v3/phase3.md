# Phase 3

Kita akan mengimplementasikan _interface_ untuk _struct_ inovice departemen marketing. Buatlah method yang dibutuhkan di file `marketing.go` agar dapat melakukan konversi data invoice marketing ke format yang sebelumnya sudah dijelaskan di **phase 1**.

Untuk `TotalInvoice` dari invoice marketing, diakumulasikan dari lama perjalanan (dalam hari) dikali dengan biaya per hari dan ditambah dengan biaya lainnya. Format untuk `Date`, `StartDate` dan `EndDate` adalah `DD/MM/YYYY`.

```txt
(end date - start date) x price per day + another fee
```

Contoh, jika `StartDate` adalah `"01/01/2022"` dan `EndDate` adalah `"03/01/2022"`, maka lama perjalanan adalah 3 hari. Jika `PricePerDay` adalah 10000 dan `AnotherFee` adalah 5000, maka total invoice adalah 35000.

Jangan lupa untuk melakukan handling error jika terdapat data yang tidak sesuai sebagai berikut:

- Jika `Date` berisi _string_ kosong, maka return `error` dengan pesan `"invoice date is empty"`
- Jika `StartDate` atau `EndDate` berisi string kosong, maka return `error` dengan pesan `"travel date is empty"`
- Jika `PricePerDay` bernilai 0 atau kurang dari 0, maka return `error` dengan pesan `"price per day is not valid"`

## Test Case

**Input**:

```txt
MarketingInvoice{
    Date: "04/01/2022",
    StartDate: "01/01/2022",
    EndDate: "03/01/2022",
    PricePerDay: 10000,
    AnotherFee: 5000,
    Approved: true
}
```

**Output**:

```txt
InvoiceData{
    Date: "04-January-2022",
    TotalInvoice: 35000,
    Department: "marketing",
}
```

Format date akan diubah menjadi `DD-month-YYYY` seperti contoh di atas, dan hasil `TotalInvoice` akan diakumulasikan dari `(end date - start date) x price per day + another fee` dengan perhitungan `3 hari x 10000 + 5000 = 35000`. Dan terakhir `Department` akan bernilai `"marketing"`.

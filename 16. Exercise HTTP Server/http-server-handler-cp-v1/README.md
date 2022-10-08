# Exercise HTTP Server - Basic Handler

## Task

Buatlah sebuah http handler dengan menggunakan `HandlerFunc` yang menampilkan waktu dengan format `"Weekday, Day Month Year"` (according to the current Time).

> **Hint:** gunakan `time.Weekday`, `time.Day`, `time.Month`, dan `time.Year`

Berikut _template_ pengerjaannya, kerjakan pada fungsi `GetHandler()`:

```go
func GetHandler() http.HandlerFunc {
    // kerjakan di sini
    return func(writer http.ResponseWriter, request *http.Request) {} // TODO: replace this
}

func main() {
    http.ListenAndServe("localhost:8080", GetHandler())
}
```

## Sample Input & Output

```bash
# Test
Input  : curl -i http://localhost:8080
Output : HTTP/1.1 200 OK
         Date: Thu, 22 Sep 2022 06:53:16 GMT
         Content-Length: 27
         Content-Type: text/plain; charset=utf-8

         Thursday, 22 September 2022
```

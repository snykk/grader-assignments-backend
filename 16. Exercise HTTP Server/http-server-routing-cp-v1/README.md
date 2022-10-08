# Exercise HTTP Server - Routing

## Task

Buatlah route untuk `TimeHandler` dan `SayHelloHandler`. Buatlah route "/time" pada `TimeHandler` dan "/hello" untuk `SayHelloHandler` dengan menggunakan `http.HandleFunc`

Disini, kita diminta untuk membuat server menggunakan Golang yang membunyai API dengan rute `/time` yang bertujuan untuk melihat waktu saat ini dengan format `hari, tanggal bulan tahun` contohnya adalah `Wednesday, 14 September 2022` atau `Sunday, 26 October 2022` dan seterusnya sesuai dengan waktu di lokasi server berada.

Selain itu, kita diminta untuk membuat API dengan rute `/hello` yang bertujuan untuk menyapa user dengan parameter query di url yaitu `?name=Yourname` dengan output `Hello, Yourname!`. Namun, jika parameter `name` tidak di lampirkan maka mengeluarkan output `Hello there`.

Jadi, buatlah server menggunakan `http.HandleFunc` dengan memiliki route endpoint sebagai berikut

- Route pertama yaitu `/time` yang menghandle fungsi `TimeHandler` untuk melihat waktu saat ini.
- Route kedua yaitu `/hello` yang menghandle fungsi `SayHelloHandler` untuk menyapa user dengan parameter url.

Berikut _template_ pengerjaannya:

```go
func TimeHandler() http.HandlerFunc {
    // kerjakan di sini
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}) // TODO: replace this
}

func SayHelloHandler() http.HandlerFunc {  
    // kerjakan di sini
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}) // TODO: replace this
}

func main() {
    // kerjakan di sini untuk route http.HandleFunc

    http.ListenAndServe("localhost:8080", nil)
}
```

Berikut adalah spesifikasi **Input** dan **Output** dari template code di atas:

## Input Format

- Route **time**: `http://host:port/time`
- Route **hello**:
  - `http://host:port/hello`
  - `http://host:port/hello?name=<Yourname>`

## Output Format

- Route **time**: `"Weekday, Day Month Year"` (according to the current Time)
  > **Hint:** gunakan `time.Weekday`, `time.Day`, `time.Month`, dan `time.Year`

- Route **hello**:
  - Message without url param is "Hello there"
  - Message with url param name is "Hello, `<Yourname>`!"

## Sample Input & Output

```bash
# Test 1 endpoint /time:
Input  : curl -i http://localhost:8080/time
Output : HTTP/1.1 200 OK
         Date: Thu, 22 Sep 2022 06:53:16 GMT
         Content-Length: 27
         Content-Type: text/plain; charset=utf-8

         Thursday, 22 September 2022

# Test 1 endpoint /hello:
Input  : curl -i http://localhost:8080/hello
Output : HTTP/1.1 200 OK
         Date: Thu, 22 Sep 2022 06:58:43 GMT
         Content-Length: 11
         Content-Type: text/plain; charset=utf-8

         Hello there

# Test 2 endpoint /hello?name=Aditira:
Input  : curl -i http://localhost:8080/hello?name=Aditira
Output : HTTP/1.1 200 OK
         Date: Thu, 22 Sep 2022 07:00:06 GMT
         Content-Length: 15
         Content-Type: text/plain; charset=utf-8

         Hello, Aditira!

# Test 3 endpoint /hello?name=Dito:
Input  : curl -i http://localhost:8080/hello?name=Dito
Output : HTTP/1.1 200 OK
         Date: Thu, 22 Sep 2022 07:00:06 GMT
         Content-Length: 15
         Content-Type: text/plain; charset=utf-8

         Hello, Dito!

# Test 4 endpoint /hello?name=Afis:
Input  : curl -i http://localhost:8080/hello?name=Afis
Output : HTTP/1.1 200 OK
         Date: Thu, 22 Sep 2022 07:00:06 GMT
         Content-Length: 15
         Content-Type: text/plain; charset=utf-8

         Hello, Afis!

# Test 5 endpoint /hello?name=Eddy:
Input  : curl -i http://localhost:8080/hello?name=Eddy
Output : HTTP/1.1 200 OK
         Date: Thu, 22 Sep 2022 07:00:06 GMT
         Content-Length: 15
         Content-Type: text/plain; charset=utf-8

         Hello, Eddy!
```

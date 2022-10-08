# Exercise HTTP Server - Custom Handler

## Task

Kita diminta untuk membuat aplikasi yang akan menampilkan random quotes dari data _slice of string_ yang telah disediakan yaitu:

```go
var Quotes = []string{
    "Be yourself; everyone else is already taken. ― Oscar Wilde",
    "Be the change that you wish to see in the world. ― Mahatma Gandhi",
    "I have not failed. I've just found 10,000 ways that won't work. ― Thomas A. Edison",
    "It is never too late to be what you might have been. ― George Eliot",
    "Everything you can imagine is real. ― Pablo Picasso",
    "Nothing is impossible, the word itself says 'I'm possible'! ― Audrey Hepburn",
}
```

Jadi buatlah custom handler `QuotesHandler` yang menampilkan quote secara acak dari data di atas.

Berikut _template_ pengerjaannya:

```go
// kerjakan di sini untuk mendefinisikan struct

func (qh QuotesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    // kerjakan di sini
}

func main() {
    handler := QuotesHandler{}
    http.ListenAndServe("localhost:8080", handler)
}
```

## Sample Input & Output

```bash
# Test 1:
Input  : curl -i http://localhost:8080
Output : HTTP/1.1 200 OK
         Date: Thu, 22 Sep 2022 09:59:48 GMT
         Content-Length: 78
         Content-Type: text/plain; charset=utf-8

         It is never too late to be what you might have been. ― George Eliot

# Test 2:
Input  : curl -i http://localhost:8080
Output : HTTP/1.1 200 OK
         Date: Thu, 22 Sep 2022 10:23:52 GMT
         Content-Length: 67
         Content-Type: text/plain; charset=utf-8

         Be the change that you wish to see in the world. ― Mahatma Gandhi


# Test 3:
Input  : curl -i http://localhost:8080
Output : HTTP/1.1 200 OK
         Date: Thu, 22 Sep 2022 10:24:06 GMT
         Content-Length: 60
         Content-Type: text/plain; charset=utf-8

         Be yourself; everyone else is already taken. ― Oscar Wilde
```

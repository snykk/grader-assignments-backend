# Exercise HTTP Client - GET & POST

## Task

Kita sangat tertarik dengan beberapa quote dari serial anime Naruto. Oleh karena itu kita sebagai seorang backend developer ingin membuat aplikasi yang menampilkan quote tersebut menggunakan 3rd Party API <https://animechan.vercel.app>.

- Jadi buatlah client Golang dengan menggunakan method **GET** ke API ini: <https://animechan.vercel.app/api/quotes/anime?title=naruto> dan lakukan `decode` terhadap response API tersebut ke dalam bentuk Array terhadap struct berikut:

  ```go
  type Animechan struct {
      Anime string
      Character string
      Quote string
  }
  ```

- Lalu, buatlah juga client Golang yang menggunakan method **POST** ke api ini: <https://postman-echo.com/post> dengan body json berikut:

  ```json
  {
      "name":  "Dion",
      "email": "dionbe2022@gmail.com",
  }
  ```

  Dan lakukan `decode` terhadap response API tersebut terhadap struct berikut:

  ```go
  type Postman struct {
      Data data 
      Url string
  }
  ```

Berikut _template_ pengerjaannya:

```go
// Lakukan `decode` dari fungsi `ClientGet()` terhadap struct berikut ini:
type Animechan struct {
    Anime string
    Character string
    Quote string
}

func ClientGet() ([]Animechan, error) {
    // Kerjakan di sini

    return []Animechan, nil // TODO: replace
}

// Lakukan `decode` dari fungsi `ClientPost()` terhadap struct berikut ini:
type data struct {
    Email string
    Name string
}

type Postman struct {
    Data data 
    Url string
}

func ClientPost() (Postman, error) {
    postBody, _ := json.Marshal(map[string]string{
        "name":  "Dion",
        "email": "dionbe2022@gmail.com",
    })
    responseBody := bytes.NewBuffer(postBody)
    fmt.Println(responseBody)

    // Kerjakan di sini

    return Postman, nil // TODO: replace
}
```

Encode **API** to **Struct**:

- API with method **GET**:

    ```json
    // curl -X GET https://animechan.vercel.app/api/quotes/anime?title=naruto
    [
    {
        "anime": "Boruto: Naruto Next Generations",
        "character": "Naruto Uzumaki",
        "quote": "The many lives lost during long years of conflict... because of those selfless sacrifices, we are able to bathe in peace and prosperity now. To ingrain this history within the new generation will be a vital cog in helping to maintain the peace."
    },
    {
        ...
    },
    ...
    ]
    ```

    to

    ```go
    type Animechan struct {
        Anime string `json:"anime"`
        Character string `json:"character"`
        Quote string `json:"quote"`
    }
    ```

- API with method **POST**:

    ```json
    // curl -X POST https://postman-echo.com/post -H "Content-Type: application/json" -d '{"name":"Dion","email":"dionbe2022@gmail.com"}
    {
    "args": {},
    "data": "'{name:Dion,email:dionbe2022@gmail.com}'",
    "files": {},
    "form": {},
    "headers": {
        "x-forwarded-proto": "https",
        "x-forwarded-port": "443",
        "host": "postman-echo.com",
        "x-amzn-trace-id": "Root=1-6329467e-3bb7e4bf7ea9d31b0d6a3edf",
        "content-length": "40",
        "user-agent": "curl/7.83.1",
        "accept": "*/*",
        "content-type": "application/json"
    },
    "json": null,
    "url": "https://postman-echo.com/post"
    }
    ```

    to

    ```go
    type data struct {
        Email string `json:"email"`
        Name string `json:"name"`
    }

    type Postman struct {
        Data data 
        Url string `json:"url"`
    }
    ```

Contoh test case :

```bash
# Method GET:
Input: https://animechan.vercel.app/api/quotes/anime?title=naruto
Output: [{Boruto: Naruto Next Generations Naruto Uzumaki The many lives lost during long years of conflict... because of those selfless sacrifices, we are able to bathe in peace and prosperity now. To ingrain this history within the new generation will be a vital cog in helping to maintain the peace.} {...} ...]

# Method POST:
Input: https://postman-echo.com/post -H "Content-Type: application/json" -d '{"name":"Dion","email":"dionbe2022@gmail.com"}'
Output: {{dionbe2022@gmail.com Dion} https://postman-echo.com/post}
```
  
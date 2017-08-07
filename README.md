# MicroService
A learning project for frontend and backend development.

## backend

### Gorilla

#### Basic usage

```go
# router of gorilla tool chain
r := mux.NewRouter()
r.HandleFunc("/products/{key}", ProductHandler)
r.HandleFunc("/articles/{category}/", ArticlesCategoryHandler)
r.HandleFunc("/articles/{category}/{id:[0-9]+}", ArticleHandler)

func ArticlesCategoryHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category: %v\n", vars["category"])
}
```

#### Router
```go
r := mux.NewRouter()
s := r.Host("www.example.com").Subrouter()
s.HandleFunc("/products/", ProductsHandler)
s.HandleFunc("/products/{key}", ProductHandler)
s.HandleFunc("/articles/{category}/{id:[0-9]+}", ArticleHandler)

```

## frontend

### Basic Demo

The basic jquery framework is used in our micro service. 

```html
<html>
<head>
    <link rel="stylesheet" href="css/main.css" />
    <script src="https://ajax.googleapis.com/ajax/libs/jquery/1.12.4/jquery.min.js">
    </script>
    <script src="js/hello.js"> </script>
</head>

<body>
    Sampple Go Applications! <br/>
    <div>
        <p class="greeting-id"> The ID is </p>
        <p class="greeting-content"> The content is </p>
    </div>
</body>
</html>
```


```js
// ajax refresh local data
$(document).ready(function() {
    $.ajax({
        url:"/api/test"
    }).then(function(data) {
        $('.greeting-id').append(data.id);
        $('.greeting-content').append(data.content);
    })
})
```
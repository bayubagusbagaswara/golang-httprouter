# Golang HTTP Router

# Pengenalan HttpRouter

- HttpRouter merukahan salah satu OpenSource Library yang populer `untuk Http Handler` di Go-Lang
- HttpRouter terkenal dengan kecepatannya dan juga sangat minimalis
- Hal ini dikarenakan HttpRouter hanya memiliki fitur untuk routing saja, tidak memiliki fitur apapun selain itu
- https://github.com/julienschmidt/httprouter

# Menambah HttpRouter ke Project

- go get github.com/julienschmidt/httprouter
- go get github.com/stretchr/testify

# Router

- Inti dari library HttpRouter adalah `struct Router`
- Router ini merupakan implementasi dari `http.Handler`, sehingga kita bisa dengan mudah menambahkan ke dalam http.Server, karena di Server kita butuh handler, dimana handler kita bisa memasukan di router
- Untuk membuat Router, kita bisa menggunakan function `httprouter.New()`, yang akan mengembalikan Router pointer

# HTTP Method

- Router mirip dengan `ServeMux`, dimana kita bisa menambahkan route ke dalam Router
- Kelebihan dibandingkan dengan ServeMux adalah pada Router, kita bisa menentukan HTTP Method yang ingin kita gunakan, misalnya `GET, POST, PUT, dan lain-lain`
- Cara menambahkan route ke dalam Router adalah gunakan function yang sama dengan HTTP Method nya, misal `router.GET()`, `router.POST()`, dan lain-lain

# httprouter.Handle

- Saat kita menggunakan ServeMux, ketika menambah route, kita bisa menambahkan `http.Handler`
- Berbeda dengan Router, pada Router kita tidak menggunakan http.Handler lagi, melainkan menggunakan `type httprouter.Handle`
- Perbedaan dengan http.Handler adalah pada httprouter.Handle terdapat `parameter ke tiga` yaitu `Params`

  ![Handle](img/handle.png)

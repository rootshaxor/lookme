# Implementasi SSH Client dengan Golang

Menjalankan perintah/shell pada server/router menggunakan `ssh` di bahasa pemrograman Go

- #### Requirement
    -  go 
    -  linux / windows / osx

- #### Dependency
    ```
    go get golang.org/x/crypto/ssh
    ```

- #### Konfigurasi Username & Password SSH berada di `config.json`
    ```
    {
        "user" : "username",
        "pass" : "password",
        "host" : "192.168.88.1", 
        "port" : "22"
    }

    ```

- #### Penggunaan 
    ```
    go run lookme.go "<perintah/shell>"

    contoh :
        - go run lookme.go "ip addr pr"
        - go run lookme.go "cat /etc/apt/sources.list"
    ```

- #### Screenshot

    - ![1]("https://github.com/rootshaxor/lookme/blob/master/ss/1.png)
    - ![2]("https://github.com/rootshaxor/lookme/blob/master/ss/1.png")

- #### Build
    > go build lookme
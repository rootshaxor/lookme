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

     - lookme Mikrotik 
        <center> <img src="https://github.com/rootshaxor/lookme/raw/master/ss/1.png"> </center>
     - lookme Debian Server
        <center> <img src="https://github.com/rootshaxor/lookme/raw/master/ss/2.png"> </center>

- #### Build
    > go build lookme.go

- #### Kekurangan
    - [x] Auto complete
    - [x] Tidak bisa menjalankan editor 
        * nano
        * vi
        * pico
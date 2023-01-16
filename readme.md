# Error, Panic, Recovery
## Error
`error` merupakan sebuah tipe data yang mengembalikan detail pesan error dalam string
- Error memiliki 1 buah property berupa method Error()
- Error termasuk tipe yang isinya bisa `nil`
```go
    var input string
    fmt.Print("Type some number: ")
    fmt.Scanln(&input)

    var number int
    var err error
    number, err = strconv.Atoi(input) // mengembalikan data number dan data error

    if err == nil { // jika tidak error
        fmt.Println(number, "is number")
    } else {
        fmt.Println(input, "is not number")
        fmt.Println(err.Error())
    }
```

## Panic
- Panic digunakan untuk menampilkan stack trace error sekaligus menghentikan flow goroutine 
- Setelah ada panic, proses akan terhenti, apapun setelah tidak di-eksekusi kecuali proses yang sudah di-defer sebelumnya (akan muncul sebelum panic error).
- Panic menampilkan pesan error di console, sama seperti fmt.Println()

```go
func main() {
    var name string
    fmt.Print("Type your name: ")
    fmt.Scanln(&name)

    if valid, err := validate(name); valid {
        fmt.Println("halo", name)
    } else {
        panic(err.Error())
        fmt.Println("end")
    }
}
```

## Recover
- Recover berguna untuk meng-handle panic error
- Pada saat panic error muncul, recover men-take-over goroutine yang sedang panic (pesan panic tidak akan muncul)
- Untuk menggunakan recover, fungsi di mana recover() berada harus dieksekusi dengan cara di-defer (diakhirkan)

```go

func catch() {
    if r := recover(); r != nil { // fungsi untuk handle panic
        fmt.Println("Error occured", r)
    } else {
        fmt.Println("Application running perfectly")
    }
}

func main() {
    defer catch() // setelah semua proses berjalan, fungsi catch dijalankan

    var name string
    fmt.Print("Type your name: ")
    fmt.Scanln(&name)

    if valid, err := validate(name); valid {
        fmt.Println("halo", name)
    } else {
        panic(err.Error())
        fmt.Println("end")
    }
}
```

## Race Condition
- Terjadi ketika ada beberapa thread mencoba mengakses dan memodifikasi data yang sama(alamat memori yang sama)

# Preslovi

Konzolna alatka i Go paket za preslovljavanje ćirilice na latinicu.

## Upotreba konzolne alatke

```shell
go install github.com/dejan/preslovi/cmd/preslovi

cat txt/kacusa.txt | preslovi
```

## Upotreba Go paketa

```go
package main

import (
    "fmt"
    "github.com/dejan/preslovi"
)

func main() {
    fmt.Println(preslovi.Latinicom("Ђурђевак"))     // Đurđevak
    fmt.Println(preslovi.Latinicom("Шабан Шаулић")) // Šaban Šaulić
}
```


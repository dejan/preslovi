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
    fmt.Println(preslovi.Latinicom("Ђурђевак"))          // Đurđevak
    fmt.Println(preslovi.Latinicom("Шабан Шаулић"))      // Šaban Šaulić
    fmt.Println(preslovi.Latinicom("ЂОРЂЕ Ђ. Ђорђевић")) // ĐORĐE Đ. Đorđević

    fmt.Println(preslovi.LatinicomAscii("Ђурђевак"))          // Djurdjevak
    fmt.Println(preslovi.LatinicomAscii("Шабан Шаулић"))      // Saban Saulic
    fmt.Println(preslovi.LatinicomAscii("ЂОРЂЕ Ђ. Ђорђевић")) // DJORDJE Dj. Djordjevic
}
```


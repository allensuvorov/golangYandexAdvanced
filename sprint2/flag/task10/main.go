// Реализуйте интерфейс flag.Value для типа
// чтобы разобрать флаг --addr=example.com:60.

package main

import (
	"flag"
	"fmt"
)

type NetAddress struct {
	Host string
	Port int
}

func (na NetAddress) String() string {
	return fmt.Sprint(na.Host + ":" + string(na.Port))
}

// допишите код реализации методов интерфейса

func main() {
	addr := new(NetAddress)
	// если интерфейс не реализован,
	// здесь будет ошибка компиляции
	_ = flag.Value(addr)
	// проверка реализации
	flag.Var(addr, "addr", "Net address host:port")
	flag.Parse()
	fmt.Println(addr.Host)
	fmt.Println(addr.Port)
}

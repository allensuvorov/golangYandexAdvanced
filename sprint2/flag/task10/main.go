// Реализуйте интерфейс flag.Value для типа
// чтобы разобрать флаг --addr=example.com:60.

package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

type NetAddress struct {
	Host string
	Port int
}

// допишите код реализации методов интерфейса
func (na NetAddress) String() string {
	return fmt.Sprint(na.Host + ":" + fmt.Sprint(na.Port))
}

func (na *NetAddress) Set(flagValue string) error {
	na.Host = strings.Split(flagValue, ":")[0]
	na.Port, _ = strconv.Atoi(strings.Split(flagValue, ":")[1])
	return nil
}

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

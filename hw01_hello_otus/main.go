package main

import ("fmt"

        "golang.org/x/example/hello/reverse"
)

func main() {
	reverseHello("Hello, OTUS!")
}

func reverseHello(s string) {
	fmt.Println(reverse.String(s))
}

// package main

// import (
// 	"fmt"
// 	"net/http"
// )

// func main() {
// 	c := make(chan string)
// 	listofsites := []string{"http://google.com", "http://yahoo.com", "http://facebook.com"}

// 	for _, site := range listofsites {
// 		go checklist(site, c)
// 	}

// 	for i := 0; i < len(listofsites); i++ {
// 		fmt.Print(<-c)
// 	}

// }

// func checklist(site string, c chan string) {
// 	_, err := http.Get(site)
// 	if err != nil {
// 		fmt.Println(site, "There's an error", err)
// 		c <- "might be down"
// 	}
// 	fmt.Println(site, "checked")
// 	c <- "\n upppp \n"
// }

package main

import "fmt"

type triangle struct {
	base   float64
	height float64
}

type square struct {
	sideLength float64
}

type shape interface {
	getArea() float64
}

func main() {
	tr := triangle{1.5, 2.5}
	sq := square{1.5}

	printArea(tr)
	printArea(sq)
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

package main

import (
	"fmt"
	"os"
	"strconv"
)

const draw = "#"
const draw_ = " "

func main() {

	var sideDigs, sideLetts int

    // Заложена возможность варьирования размера каждой стороны
	// (с ошибками решил пока не заморачиваться)
	switch len(os.Args) {
	case 2:
		sideLetts, _ = strconv.Atoi(os.Args[1])
		sideDigs, _ = strconv.Atoi(os.Args[1])
	case 3:
		sideLetts, _ = strconv.Atoi(os.Args[1])
		sideDigs, _ = strconv.Atoi(os.Args[2])
	default:
		sideLetts = 8
		sideDigs = 8
	}

	// Подстраиваю так, чтобы левая нижняя клетка была всегда черная (как в шахматах)
	for i := sideDigs; i > 0; i-- {
		for j := 1; j <= sideLetts; j++ {
			// Если сумма индексов четная - черная клетка, иначе - белая
			if (i+j)%2 == 0 {
				fmt.Print(draw)
			} else {
				fmt.Print(draw_)
			}
		}
		fmt.Println()
	}
}

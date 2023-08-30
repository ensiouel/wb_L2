Что выведет программа? Объяснить вывод программы.

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func asChan(vs ...int) <-chan int {
	c := make(chan int)

	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}

		close(c)
	}()
	return c
}

func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for {
			select {
			case v := <-a:
				c <- v
			case v := <-b:
				c <- v
			}
		}
	}()
	return c
}

func main() {

	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4 ,6, 8)
	c := merge(a, b )
	for v := range c {
		fmt.Println(v)
	}
}
```

Ответ:
```
1
2
3
4
5
6
7
8
0
0
```

Из-за того, что в `select` нет проверки на то, закрыты ли каналы `a` и `b`, после непосредственного закрытия каналов из них будет считываться нулевое значение типа (`0` для `int`).

Одним из способов исправления данного кода является использование двух горутин для чтения данных и счетчика `sync.WaitGroup`, который будет ждать до тех пор, пока все каналы не будут закрыты, после чего закроет канал `c`.
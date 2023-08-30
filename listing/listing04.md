Что выведет программа? Объяснить вывод программы.

```go
package main

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	for n := range ch {
		println(n)
	}
}
```

Ответ:
```
0
1
2
3
4
5
6
7
8
9
fatal error: all goroutines are asleep - deadlock!
```

В программе возникает deadlock при чтении данных из канала в цикле `for`, из-за того, что канал не был закрыт, `range ch` будет бесконечно ждать данные. Исправить ошибку можно закрыв канал `ch` после того, как все необходимые данные были в него отправлены.
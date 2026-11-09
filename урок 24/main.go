package main

import "fmt"

// 1
func sendGreeting(ch chan string) {
	ch <- "Привет из горутины!"
}

// 2
func squareWorker(number int, ch chan int) {
	ch <- number * number
}

// 3
func emitNumbers(ch chan int) {
	for i := 1; i <= 5; i++ {
		ch <- i
	}
	close(ch)
}

// 4
func sumReader(ch chan int) {
	sum := 0
	for number := range ch {
		sum += number
	}
	fmt.Println(sum)
}

// 5
func filterEven(input chan int, output chan int) {
	for number := range input {
		if number%2 == 0 {
			output <- number
		}
	}
	close(output)
}

// 6
func checkChannel(ch chan string) {
	value, ok := <-ch
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("Канал закрыт!")
	}
}

func main() {
	//1
	greeting := make(chan string)
	go sendGreeting(greeting)
	fmt.Println(<-greeting)

	//2
	squares := make(chan int)
	go squareWorker(9, squares)
	fmt.Println(<-squares)

	//3
	numbers := make(chan int)
	go emitNumbers(numbers)
	for number := range numbers {
		fmt.Println(number)
	}

	//4
	values := make(chan int)
	done := make(chan bool)
	go func() {
		sumReader(values)
		done <- true
	}()
	values <- 2
	values <- 4
	values <- 6
	close(values)
	<-done

	//5
	input := make(chan int, 10)
	output := make(chan int, 10)
	for number := 1; number <= 10; number++ {
		input <- number
	}
	close(input)
	filterEven(input, output)
	for number := range output {
		fmt.Println(number)
	}

	//6
	openChannel := make(chan string, 1)
	openChannel <- "Сообщение из открытого канала"
	checkChannel(openChannel)
	closedChannel := make(chan string)
	close(closedChannel)
	checkChannel(closedChannel)

	//7
	messages := make(chan string, 3)
	messages <- "Первое сообщение"
	messages <- "Второе сообщение"
	messages <- "Третье сообщение"
	select {
	case messages <- "Четвёртое сообщение":
		fmt.Println("Четвёртое сообщение отправлено")
	default:
		fmt.Println("Канал заполнен, четвёртое сообщение не отправлено")
	}
	for i := 0; i < 3; i++ {
		fmt.Println(<-messages)
	}
}

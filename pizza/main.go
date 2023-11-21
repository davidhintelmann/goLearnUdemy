package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

const NumberOfPizzas uint8 = 10

var pizzasMade, pizzasFailed, total uint8

type Producer struct {
	Data chan PizzaOrder
	Quit chan chan error
}

type PizzaOrder struct {
	PizzaNumber uint8
	Message     string
	Success     bool
}

func main() {
	// sedd the random number generator
	rand.NewSource(time.Now().Unix())

	// print out a message
	fmt.Println("The Pizzeria is open for business!")
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")

	// create a producer
	pizzaJob := &Producer{
		Data: make(chan PizzaOrder),
		Quit: make(chan chan error),
	}

	// run the producer in the background
	go Pizzeria(pizzaJob)

	// create and run consumer
	for i := range pizzaJob.Data {
		if i.PizzaNumber <= NumberOfPizzas {
			if i.Success {
				fmt.Println(i.Message)
				fmt.Printf("Order #%v is out for delivery!\n", i.PizzaNumber)
			} else {
				fmt.Println("The customer is really mad!")
			}
		} else {
			fmt.Println("Done making pizzas...")
			err := pizzaJob.Close()
			if err != nil {
				log.Fatalf("*** Error closing channel: %v\n", err)
			}
		}
	}

	// print out the ending message
	fmt.Println("~~~~~~~~~~~~~~~~~")
	fmt.Println("Done for the day!")
	fmt.Printf("The total number of pizzas made: %v\nThe number of successful pizzas made: %v\nThe number of failed pizzas made: %v\n", total, pizzasMade, pizzasFailed)
}

func (p *Producer) Close() error {
	ch := make(chan error)
	p.Quit <- ch
	return <-ch
}

func MakePizza(pizzaNumber uint8) *PizzaOrder {
	pizzaNumber++
	if pizzaNumber <= NumberOfPizzas {
		delay := rand.Intn(5) + 1
		fmt.Printf("Received order #%v\n", pizzaNumber)

		rnd := rand.Intn(12) + 1
		msg := ""
		success := false
		if rnd < 5 {
			pizzasFailed++
		} else {
			pizzasMade++
		}
		total++

		fmt.Printf("Making pizza #%v. It will take %v seconds...\n", pizzaNumber, delay)
		time.Sleep(time.Second * time.Duration(delay))

		if rnd <= 2 {
			msg = fmt.Sprintf("*** We ran out of ingredients for pizza #%v\n", pizzaNumber)
		} else if rnd <= 4 {
			msg = fmt.Sprintf("*** The cook quit while making pizza #%v\n", pizzaNumber)
		} else {
			success = true
			msg = fmt.Sprintf("*** Pizza order #%v is ready!", pizzaNumber)
		}
		p := PizzaOrder{
			PizzaNumber: pizzaNumber,
			Message:     msg,
			Success:     success,
		}
		return &p
	}
	return &PizzaOrder{PizzaNumber: pizzaNumber}
}

func Pizzeria(pizzaMaker *Producer) {
	// keep track of which pizza we are making
	var i uint8 = 0

	// run forever or until we receive a quit notification
	// try to make pizzas
	for {
		currentPizza := MakePizza(i)

		if currentPizza != nil {
			i = currentPizza.PizzaNumber
			select {
			// we tried to make a pizza (we sent comething to the Data channel)
			case pizzaMaker.Data <- *currentPizza:

			case quitChan := <-pizzaMaker.Quit:
				close(pizzaMaker.Data)
				close(quitChan)
				// exit goroutine
				return
			}
		}

	}
}

package main

import "fmt"

type EvenSource struct {
	consumerList []ConsumerInterface
}

func (e *EvenSource) Register(c ConsumerInterface) {
	e.consumerList = append(e.consumerList, c)
}

func (e *EvenSource) Notify() {
	for _, consumer := range e.consumerList {
		consumer.Notify(fmt.Sprintf("%s notified.", consumer.GetName()))
	}
}

type ConsumerInterface interface {
	GetName() string
	Notify(message string)
}

type Consumer struct {
	Name string
}

func (c Consumer) GetName() string {
	return c.Name
}

func (c Consumer) Notify(message string) {
	fmt.Printf("%s\n", message)
}

func main() {
	e := EvenSource{}
	e.Register(Consumer{Name: "consumer 1"})
	e.Register(Consumer{Name: "consumer 2"})

	e.Notify()
}

package main

import "fmt"

// concrete observer

type Customer struct {
	id string
}

func (c *Customer) update(itemName string) {
	fmt.Println("Sending email to customer %s for item %s\n", c.id, itemName)
}

func (c *Customer) getID() string {
	return c.id
}

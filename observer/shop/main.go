package main

func main() {
	shirtItem := NewItem("Nike Shirt")

	observerFirst := &Customer{id: "abc@gmail.com"}
	observerSecond := &Customer{id: "asdasd@gmail.com"}

	shirtItem.Register(observerFirst)
	shirtItem.Register(observerSecond)

	shirtItem.UpdateAvailability()
}

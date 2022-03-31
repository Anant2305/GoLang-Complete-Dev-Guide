package main

func main() {
	//var card string = "Ace of Spades"
	//card := "Ace of Spades" //Colon is only rewuired first incstance of the cariable being assigned.
	//card := newCard()
	//fmt.Println(card)

	//using Slice
	//cards := deck{"Ace of Diamonds", newCard()}
	//Appending a Slice
	//cards = append(cards, "Six of Spades")

	// cards := newDeck()

	// hand, remaingCards := deal(cards, 5)

	// hand.print()
	// remaingCards.print()

	//cards := newDeck()
	//fmt.Println(cards.toString())

	//save to local file system
	//cards.saveToFile("myCards")

	//get file from local system
	// cards := newDeckFromFile("myCards")
	// cards.print()

	//Shuffle cards
	cards := newDeck()
	cards.shuffle()
	cards.print()

}

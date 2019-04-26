package main

import "fmt"

func main()  {
	doDBoperations()
}

func connectToDB() {
	fmt.Println("Ok, connected to db")
}

func disconnectFromDB() {
	fmt.Println("Ok, disconnected from db")
}

func doDBoperations() {
	connectToDB()
	fmt.Println("Defering the database disconnect")
	defer disconnectFromDB()
	fmt.Println("Doing some DB operation")
	fmt.Println("Oops! some crash or network error")
	fmt.Println("Returning from function here")
	return
}
package main

import "fmt"

type Game struct {
	ActiveKingdoms []Kingdom
	Turn           int
	Player         Kingdom
	isOver         bool
}

func (g *Game) SetPlayerKingdom(PlayerKingdom Kingdom) {
	g.Player = PlayerKingdom
}

type Kingdom struct {
	Name      string
	Resources Resources
	Army      Army
	Towns     []City
}

type Resources struct {
	Gold   int
	Rice   int
	People int
}

type City struct {
	Name    string
	Defense int
	Tribute int
}

type Army struct {
	Soldiers int
}

func (a *Army) RecruitSoldiers(NewSoldiers int) {
	a.Soldiers += NewSoldiers
}

func main() {
	GameSet := LaunchGame()

	DisplayWelcomeMessage()

	ChooseYourKingdom(&GameSet)

	fmt.Println(GameSet)

	CreateMap()

	//fmt.Println(GameMap)
}

func DisplayWelcomeMessage() {
	fmt.Println("Welcome to Warring Kingdoms!")
	fmt.Println("The Country is divided into three realms: Xian, Tan, and Hui.")
	fmt.Println("Each realm fights for control over the Imperial Throne.")
	fmt.Println("Choose your realm and lead your people to victory!")
}

func LaunchGame() Game {
	GameSet := Game{
		ActiveKingdoms: []Kingdom{{Name: "Xian"}, {Name: "Tan"}, {Name: "Hui"}},
		Turn:           0,
		Player:         Kingdom{Name: "Undefined Player Kingdom"},
		isOver:         false,
	}
	return GameSet
}

func CreateMap() [][]string {
	GameMap := [][]string{}
	return GameMap
}

func ChooseYourKingdom(GameSet *Game) {
	for {
		var PlayerChoice string
		fmt.Scan(&PlayerChoice)

		if PlayerChoice == "Xian" || PlayerChoice == "Tan" || PlayerChoice == "Hui" {
			PlayerKingdom := Kingdom{Name: PlayerChoice}
			fmt.Printf("You are ruler of %s kingdom! \n", PlayerKingdom.Name)
			GameSet.SetPlayerKingdom(PlayerKingdom)
			break
		} else {
			fmt.Print("Choose your realm: Xian, Tan or Hui.")
			fmt.Scan(&PlayerChoice)
		}
	}
}

func FinishGame() {
	fmt.Println("Game Over!")
}

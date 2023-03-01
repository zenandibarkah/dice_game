package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var numberOfDice = []int{1, 2, 3, 4, 5, 6}
var onlyOnce sync.Once

func RollDice() int {
	onlyOnce.Do(func() {
		rand.Seed(time.Now().UnixNano()) // only run once
	})
	return numberOfDice[rand.Intn(len(numberOfDice))]
}

func getDices(M int) []int {
	result := []int{}
	for i := 0; i < M; i++ {
		dice := RollDice()
		result = append(result, dice)
	}

	return result
}

func Delete(s []int, data int) []int {
	index := 0
	for _, i := range s {
		if i != data {
			s[index] = i
			index++
		}
	}
	return s[:index]
}

func moveData(s []tempdata, minDice int) []tempdata {
	for i := 0; i < len(s); i++ {
		if i == 0 {
			for _, data := range s[i].dices {
				if data == minDice {
					newDataDice := Delete(s[i].dices, minDice)
					s[i].dices = newDataDice
					s[i+1].dices = append(s[i+1].dices, minDice)

				}
			}
		} else if i < len(s)-1 {
			for _, data := range s[i].dices {
				if data == minDice {
					// fmt.Println("data d", t)
					newDataDice := Delete(s[i].dices, minDice)
					s[i].dices = newDataDice
					s[i+1].dices = append(s[i+1].dices, minDice)

				}
			}
		} else {
			for _, data := range s[i].dices {
				if data == minDice {
					newDataDice := Delete(s[i].dices, minDice)
					s[i].dices = newDataDice
					s[0].dices = append(s[0].dices, minDice)
				}
			}
		}

	}
	return s
}

func getScore(s []int, data int, scores int) int {
	for _, i := range s {
		if i == data {
			scores++
		}
	}
	return scores
}

func getWinner(s []tempdata) int {
	tempScore := 0
	winner := 0
	for i := 0; i < len(s); i++ {
		if s[i].score > tempScore {
			tempScore = s[i].score
			winner = s[i].player
		}
	}
	return winner
}

type tempdata struct {
	player int
	dices  []int
	score  int
}

func Play(N int, M int) {

	var tempdice []int
	var finaldice []int
	var finalscore int
	var tempResults []tempdata

	round := 0
	score := 0
	round_stop := 0

	for true {
		round++
		fmt.Println("======================")
		fmt.Println("Giliran ", round, " lempar dadu:")
		if round == 1 {

			for player := 1; player <= N; player++ {

				tempdice = getDices(M)

				fmt.Printf("Pemain #%d (%d) : %v", player, score, tempdice)
				fmt.Println()

				data := tempdata{player: player, dices: tempdice, score: score}

				tempResults = append(tempResults, data)
			}

			fmt.Println("Setelah Evaluasi:")
			for i := 0; i < len(tempResults); i++ {
				finalscore = getScore(tempResults[i].dices, 6, tempResults[i].score)
				tempResults[i].score = finalscore

				// tempResults = moveData(tempResults, 1)
				finaldice = Delete(tempResults[i].dices, 6)
				tempResults[i].dices = finaldice

				fmt.Printf("Pemain #%d (%d) : %v", tempResults[i].player, finalscore, finaldice)
				fmt.Println()

			}
		} else {
			for player := 0; player < len(tempResults); player++ {
				for i := 0; i < len(tempResults[player].dices); i++ {
					tempdice = getDices(len(tempResults[player].dices))

					tempResults[player].dices = tempdice
				}
				fmt.Printf("Pemain #%d (%d) : %v", tempResults[player].player, tempResults[player].score, tempResults[player].dices)
				fmt.Println()

			}

			fmt.Println("Setelah Evaluasi:")
			for i := 0; i < len(tempResults); i++ {
				finalscore = getScore(tempResults[i].dices, 6, tempResults[i].score)
				tempResults[i].score = finalscore

				// tempResults = moveData(tempResults, 1)
				finaldice = Delete(tempResults[i].dices, 6)
				tempResults[i].dices = finaldice

				fmt.Printf("Pemain #%d (%d) : %v", tempResults[i].player, finalscore, finaldice)
				fmt.Println()

				if len(tempResults[i].dices) == 0 {
					round_stop += 1
				}
			}

			fmt.Println("======================")
			if round_stop > M {
				winner := getWinner(tempResults)
				fmt.Printf("Game dimenagkan oleh pemain #%d karena memiliki point lebih banyak dari pemain lain", winner)
				break
			}

		}
	}

}

func main() {
	var N, M int
	fmt.Print("Jumlah pemain: ")
	fmt.Scanln(&N)

	fmt.Print("Jumlah Dadu: ")
	fmt.Scanln(&M)

	Play(N, M)
}

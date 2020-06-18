package main

/*
Bank Heist______________________________
We often include conditionals for when we need to account for different situations and outcomes.
When it comes to coming up with a plan and executing it, there’s nothing more uncertain than a bank heist!
(Watch any robbery themed movie if you need convincing).
In this project, we’re going to use conditionals to simulate this situation along with hi-jinks and hiccups that may pop up.
Who knows? Maybe, just maybe, if all goes well, we’ll have a few more dollars after.
BY : REDKHAY
*/

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())
	isHeistOn := true
	eludedGuards := rand.Intn(100)

	// 1st: Sneak past guards
	if eludedGuards >= 50 {
		fmt.Println("Good job!, This is the 1st Step")
	} else {
		isHeistOn = false
		fmt.Println("STOP!, Your FAILED")
	}
	// 2nd: Open the Vault
	openedVault := rand.Intn(100)
	if isHeistOn && openedVault >= 70 {
		fmt.Println("The vault has been opened, GO!!")
	} else if isHeistOn = false; isHeistOn {
		fmt.Println("The vault can't be Opened")
	}
	// 3rd: Leaving
	var leftSafely = rand.Intn(5)
	if isHeistOn {
		switch leftSafely {
		case 0:
			isHeistOn = false
			fmt.Println("ALARM!")
		case 1:
			isHeistOn = false
			fmt.Println("POLICE!")
		case 2:
			isHeistOn = false
			fmt.Println("Fingerprint DETECTED")
		case 3:
			isHeistOn = false
			fmt.Println("Vaults DOGS")
		default:
			fmt.Println("Start the Getaway CAR!")
		}
	}
	// Wrapping UP:
	if isHeistOn {
		amStolen := 10000 + rand.Intn(10000)
		fmt.Println(amStolen, "$")
	}

	// current isHeistOn
	fmt.Print("\n********\n")
	fmt.Print("isHeistOn is currently: ", isHeistOn)

}

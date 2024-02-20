package mutexes

import (
	"fmt"
	"sync"
)

var mutex sync.Mutex

func buyTickets(wg *sync.WaitGroup, userId int, remainingTickets *int) {
	defer wg.Done()
	mutex.Lock()
	if *remainingTickets > 0 {
		*remainingTickets--
		fmt.Printf("user %d purchased a ticket. Remaining Tickets: %d\n", userId, *remainingTickets)
	} else {
		fmt.Printf("user %d not found ticket\n", userId)
	}
	mutex.Unlock()
}

func BookTickets() {
	tickets := 500

	var wg sync.WaitGroup

	for userId := 0; userId < 2000; userId++ {
		wg.Add(1)

		go buyTickets(&wg, userId, &tickets)
	}

	wg.Wait()
}

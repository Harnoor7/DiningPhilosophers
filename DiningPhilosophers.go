package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

type Chopstick struct {
	lock        sync.Mutex
	chopStickId int
}

type Philosopher struct {
	leftChopstick  *Chopstick
	rightChopstick *Chopstick
	philosopherId  int
}

func (philosopher *Philosopher) eat(sem *chan int) {
	for i := 0; i < 3; i++ {
		*sem <- 1
		leftChopstick := philosopher.leftChopstick
		rightChopstick := philosopher.rightChopstick
		leftChopstick.lock.Lock()
		rightChopstick.lock.Lock()
		fmt.Println("philosopher with id: ", philosopher.philosopherId, " is starting to eat with chopsticks with ids: ",
			leftChopstick.chopStickId, rightChopstick.chopStickId)
		time.Sleep(1 * time.Second)
		fmt.Println("philosopher with id: ", philosopher.philosopherId, " is fininshing eating with chopsticks with ids: ",
			leftChopstick.chopStickId, rightChopstick.chopStickId)
		leftChopstick.lock.Unlock()
		rightChopstick.lock.Unlock()
		<-*sem
	}
	wg.Done()
}

func startDinnerForPhilosophers(philosophers []Philosopher) {

	sem := make(chan int, 2)
	wg.Add(5)

	for i := 0; i < 5; i++ {
		go philosophers[i].eat(&sem)
	}

	wg.Wait()

}

func main() {

	fmt.Println("Going to create 5 chopsticks.")
	chopsticks := make([]Chopstick, 5)

	for i := 0; i < 5; i++ {
		chopsticks[i] = Chopstick{}
		chopsticks[i].chopStickId = i
		fmt.Println("Chopstick created with id: ", chopsticks[i].chopStickId)
	}

	philosphers := make([]Philosopher, 5)
	for i := 0; i < 5; i++ {
		philosphers[i] = Philosopher{&chopsticks[i], &chopsticks[(i+1)%5], i}
		leftChopstick := *(philosphers[i].leftChopstick)
		rightChopstick := *(philosphers[i].rightChopstick)
		fmt.Println("Philospher created with id: ", philosphers[i].philosopherId,
			" having left chopstick id: ", leftChopstick.chopStickId,
			" having right chopstick id: ", rightChopstick.chopStickId)
	}

	startDinnerForPhilosophers(philosphers)

}

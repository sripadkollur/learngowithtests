package main

import (
	"os"
	"time"

	"github.com/sripadkollur/learngowithtests/pkg/mocking"
)

// func Greet(writer io.Writer, name string) {
// 	fmt.Fprintf(writer, "Hello, %s", name)
// }

// func Countdown(out io.Writer) {
// 	fmt.Fprintf(out, "3")
// }
func main() {
	// fmt.Println(integers.Add(2, 2))
	// Greet(os.Stdout, "Elodie")
	// di.Greet(os.Stdout, "Sripad")

	sleeper := &mocking.ConfigurableSleeper{Duration: 2 * time.Second, CallSleep: time.Sleep}
	mocking.Countdown(os.Stdout, sleeper)
}

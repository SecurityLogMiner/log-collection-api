package greetings 

import "fmt"

func Hello(name string) string {
    message := fmt.Sprintf("Welcome, %s!",name)
    return message
}

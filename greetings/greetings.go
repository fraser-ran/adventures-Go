package greetings

import (
	"fmt"
	)

// Hello returns a greeting for the named person.
func Hello(name string) string {
    // Return a greeting that embeds the name in a message.
    message := fmt.Sprintf("Hi, %v. Welcome!", name)
    return message
}


// //func beans(name string) (string){
//     bean:= fmt.Sprintf("Hi, I am %v and I like beans", name)
//     return bean
// }
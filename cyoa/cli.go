package cyoa

import (
	"fmt"
	"strings"
)

func ShowStoryCli(c Chapter) {
	fmt.Println(c.Title)
	fmt.Println("------------")
	fmt.Println(strings.Join(c.Story, " "))
	fmt.Println("------------")
	if len(c.Options) != 0 {
		fmt.Println("Press the number to choose how to proceed with this story : ")
	}
	for i, st := range c.Options {
		fmt.Printf("%d.  %v \n", i+1, st.Text)
	}
}

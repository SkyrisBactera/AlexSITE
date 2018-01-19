/*
Contains code that does not fit in to the other categories, or entrypoint-level execution
*/

package main

import (
	"encoding/base64"
	"fmt"

	"github.com/gopherjs/gopherjs/js" // Allows the use of JavaScript specific functions, etc.
	"github.com/gopherjs/jquery"      // Bindings for the JavaScript library jQuery
	"honnef.co/go/js/dom"             // Bindings for JavaScript DOM
)

var (
	// Misc
	jQuery    = jquery.NewJQuery // Creating jQuery instance for the rest of the code to work with
	entryPass = randString(5)    // The entry password that is needed to enter the website. This is later converted to Base64 for the user to decode
)

/*
Declares functions to be called by JavaScript. start() takes the role of main() of being the main entrypoint for the application to ensure that the website has loaded
before running code that changes or uses elements on the website to prevent those types of errors from occurring
*/
func main() {
	go func() {
		js.Global.Set("go", map[string]interface{}{ // Specifies that each of the functions is under the variable go
			"checkPass":      checkPass,      // Checks the password the user entered against the entry password (entryPass)
			"start":          start,          // Takes the role of main() and starts the main functionality of the program
			"checkEaster":    checkEaster,    // Runs the entered easter egg against the array of easter eggs and runs addEaster() if they are correct (easterEggs)
			"checkChallenge": checkChallenge, // Has argument id of type int that checks the entered answer against the array of correct answers (answers)
			"debug":          debug,          // Disables entry password
			"birthday":       birthday,       // Rotates alexHead.png and has a rainbow effect with text underneath it that fades in reading "Happy Birthday!". This should never be called by the website, but is for the development console
		})
	}()
}

// When called, disables the entry password and prints "Debugging/Cheat mode activated!" in the console
func debug() {
	fmt.Println("Debugging/Cheat mode activated!")
	js.Global.Get("document").Call("querySelector", "#passwordDiag").Call("close") // Closes password dialogue
	jQuery("#blurme").RemoveClass("fogged")                                        // Removes blurry effect
}

// Checks the password the user entered against the actual entry password (entryPass)
func checkPass() {
	go func() {
		if jQuery("#sample3").Val() == entryPass {
			js.Global.Get("document").Call("querySelector", "#passwordDiag").Call("close") // Closes password dialogue
			jQuery("#blurme").RemoveClass("fogged")                                        // Removes blurry effect
		} else {
			js.Global.Call("alert", "Wrong") // Displays popup box with message "Wrong"
		}
	}()
}

// Takes the role of main() and starts the main functionality of the program
func start() {
	fmt.Println("Page loaded")                                                      // Prints "Page loaded" to the console because start() is called when the page is done loading
	jQuery("#passer").SetText(base64.StdEncoding.EncodeToString([]byte(entryPass))) // Encodes the entry password into a base64 string and sets #passer's contents to the converted password
	hideContent()                                                                   // Blurs the main content
	hideElements()                                                                  // Hides elements that are only shown in certain instances, or until they are ready
	// Adds event listener to save/persist information before the page is closed
	dom.GetWindow().AddEventListener("beforeunload", false, func(event dom.Event) {
		save()
	})
	specialCheck() // Checks if the date is of significance and activates special events in those situations. Returns bool corresponding to if the date is of significance
	loadSave()     // Imports saved information to their corresponding variables
	// If the program is in development mode, then debugging mode is also activated
	if devel {
		debug() // Enable debugging mode
	}
}

//Checks if int e is in int slice s
func contains(s []int, e int) bool {
	for _, a := range s { // Iterates through s
		if a == e { // If the current element is equal to e, then return true
			return true
		}
	}
	return false // If none of the elements are equal to e, then return false
}

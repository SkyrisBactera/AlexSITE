/*
Contains code to save, and load variables with localstorage
*/

package main

import (
	"fmt"

	"github.com/go-humble/locstor" // Local Storage bindings for Go
	"github.com/gopherjs/gopherjs/js"
)

// Loads previously saved variables into the variable placeholders located across the application
func loadSave() {
	go func() {
		// Creates new method of storing using JSON. Binary encoding is ideal as it requires less storage on the user's computer, however it is much buggier
		store := locstor.NewDataStore(locstor.JSONEncoding)
		var previousRev float64 // A placeholder for the revision stored since last time
		// Loads the unlockedEasterEggs, and if there is an error, report it in the console
		if err := store.Find("eastereggs", &unlockedEasterEggs); err != nil {
			fmt.Println("Couldn't load eastereggs") // If there is an error, report it in the console.
		}
		loadedEaster() // After loading the easter eggs, change UI based on the stats
		// Loads the completeChallenges, and if there is an error, report it in the console
		if err := store.Find("challenges", &completeChallenges); err != nil {
			fmt.Println("Couldn't load completeChallenges") // If there is an error, report it in the console.
		}
		loadChallenges() // After loading necessary components for the challenges, load them in
		/*
			Loads the previousRev, and if there is an error, report it in the console. Call showChanges based on the revision's value
		*/
		if err := store.Find("previousRev", &previousRev); err != nil {
			fmt.Println("Couldn't load the previousRev") // If there is an error, report it in the console.
		} else {
			showChanges(previousRev) // Show changelog and upgraded version with the previous revision as float64
		}
	}()
}

// Saves variables in localstorage for next time the user opens the page
func save() {
	go func() {
		// Creates new method of storing using JSON. Binary encoding is ideal as it requires less storage on the user's computer, however it is much buggier
		store := locstor.NewDataStore(locstor.JSONEncoding)
		// Saves the unlockedEasterEggs, and if there is an error, alert the user so they don't lose their work
		if err := store.Save("eastereggs", unlockedEasterEggs); err != nil {
			js.Global.Call("alert", "Cannot save your work! Try using a different browser") // If there is an error, alert the user so they don't lose their work
		}
		// Saves the completeChallenges, and if there is an error, alert the user so they don't lose their work
		if err := store.Save("challenges", completeChallenges); err != nil {
			js.Global.Call("alert", "Cannot save your work! Try using a different browser") // If there is an error, alert the user so they don't lose their work
		}
		// Saves the current revision, and if there is an error, alert the user so they don't lose their work
		if err := store.Save("previousRev", revision); err != nil {
			js.Global.Call("alert", "Cannot save your work! Try using a different browser") // If there is an error, alert the user so they don't lose their work
		}
	}()
}

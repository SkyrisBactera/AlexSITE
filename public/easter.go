/*
Contains code based around the easter egg functionality
*/
package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gopherjs/gopherjs/js"
)

var unlockedEasterEggs []int

// Runs the entered easter egg against the array of easter eggs and runs addEaster() if they are correct (easterEggs)
func checkEaster() {
	go func() {
		query := strings.ToLower(jQuery("#easterText").Val()) // Gets the value of the easter egg textbox and converts it to lowercase
		// Iterates over easterEggs array
		for i, ustr := range easterEggs {
			go func(i int, ustr string) { // Run goroutine to make the easter egg check concurrent and thus faster (by a tiny margin)
				str := strings.ToLower(ustr)          // Converts the easter egg to lowercase to ensure consistency
				stringContains := true                // This value will stay true if the textbox contains all the necessary components, but will become false if one of the components is missing
				pstr := strings.Split(str, "^")       // Splits the string into a slice with ^-delimiting where pstr[0] is the actual easter egg with %-delimiting, and pstr[1] (if it exists), is a custom message if the easter egg is unlocked
				result := strings.Split(pstr[0], "%") // Splits the pstr's first element (actual easter egg) into a slice with %-delimiting
				// For every element in the split string, check if the query contains that element
				for _, easter := range result {
					// If the query does not contain one of the components of the easter egg, set stringContains to false and break out of the for loop to prevent wasting time
					if !strings.Contains(query, easter) {
						stringContains = false // Set stringContains to false so that the user is not awarded
						break                  // Break out of for loop to prevent wasting time checking if it contains other components
					}
				}
				// If the string contained all the necessary components, call addEaster
				if stringContains {
					if len(pstr) > 1 {
						addEasterStr(i, pstr[1]+" ")
					} else {
						addEaster(i)
					}
				}
			}(i, ustr)
		}
	}()
}

// Calls addEasterStr with an empty string for a default message
func addEaster(i int) {
	addEasterStr(i, "") // Calls addEasterStr with an empty string for a default message
}

// Handles adding the easter egg to unlockedEasterEggs, calls updateEasterList to reflect the new easter egg, and checks if the easter egg has already been unlocked, and reports this to the user
func addEasterStr(i int, s string) {
	go func() {
		/*
			If the user hasn't already unlocked the easter egg, then show a message that says that they unlocked it, add the easter egg id to the unlockedEasterEggs slice,
			and call updateEasterList() to update the list of unlocked easter eggs
		*/
		if !contains(unlockedEasterEggs, i) {
			li := jQuery("<li>").SetText(fmt.Sprintf("%sEaster Egg %v Unlocked", s, i)) // Create li element telling the user that they unlocked an easter egg
			jQuery("#resp").Append(li)                                                  // Append the li element
			// Fade the li out, then delete it
			li.FadeOut(4000, func() {
				li.Remove() // Delete the li element
			})
			unlockedEasterEggs = append(unlockedEasterEggs, i) // Add the easter egg ID to the unlockedEasterEggs slice
			go updateEasterList()                              // Update the list on the website with the new easter eggs
		} else { // If the user already unlocked that easter egg, say that
			li := jQuery("<li>").SetText("You have already unlocked this") // Create li element telling the user they already unlocked the easter egg
			jQuery("#resp").Append(li)                                     // Append the li element
			// Fade the li out, then delete it
			li.FadeOut(4000, func() {
				li.Remove() // Delete the li element
			})
		}
	}()
}

// Update the #easterlist unordered list to represent what is in the unlockedEasterEggs slice, and update the counter for how many you have unlocked
func loadedEaster() {
	// If the user has unlocked easter eggs, then hide the message that says that you haven't unlocked any easter eggs
	if len(unlockedEasterEggs) != 0 {
		jQuery("#noeasters").Hide()
	}
	// If the user has unlocked more than half of the available easter eggs, then remove the entry password dialogue
	if len(unlockedEasterEggs) >= easterSize/2 {
		js.Global.Get("document").Call("querySelector", "#passwordDiag").Call("close") // Closes password dialogue
		jQuery("#blurme").RemoveClass("fogged")                                        // Removes blurry effect
	}
	updateEasterList()
}

// Update the #easterlist unordered list to represent what is in the unlockedEasterEggs slice, and update the counter for how many you have unlocked
func updateEasterList() {
	jQuery("#easterlist").Empty() // Empty the list of all it's li's so we don't have duplicate entries
	// Add each unlocked easter egg to the list as an li element
	for _, r := range unlockedEasterEggs {
		jQuery("#easterlist").Append(jQuery("<li>").SetText("Easter Egg " + strconv.Itoa(r+1))) // Add the easter egg to the list as an li element with it's corresponding easter egg ID
	}
	jQuery("#eastersUnlocked").SetText(fmt.Sprintf("You have unlocked %v out of the %v available easter eggs", len(unlockedEasterEggs), easterSize)) // Update the counter for how many you have unlocked
}

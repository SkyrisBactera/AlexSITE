package main

import (
	"fmt"
	"time"

	"github.com/gopherjs/gopherjs/js"
)

// Checks if the date is of significance and activates special events in those situations
func specialCheck() bool {
	returnVal := false // Represents whether or not there is a special event
	now := time.Now()
	day := now.Day()     // The current day
	month := now.Month() // The current month

	// Birthday (January 1st)
	if day == 1 && month == time.January {
		returnVal = true // Sets the returnVal to true to say that there is a special event
		birthday()       // Rotates alexHead.png and has a rainbow effect with text underneath it that fades in reading "Happy Birthday!"
	}
	// Shows h5 #specialEvent which categorizes what is under it as a special event
	if returnVal == true {
		jQuery("#specialEvent").Show()
	}
	/*
		Returns:
		true: There is a special event
		false: There is not a special event
	*/
	return returnVal
}

// Rotates alexHead.png and has a rainbow effect with text underneath it that fades in reading "Happy Birthday!"
func birthday() {
	go func() {
		fmt.Println("Birthday event!")    // Prints to the console "Birthday event!"
		jQuery("#birthdaySpecial").Show() // Show's div #birthdaySpecial which contains the birthday message and the image
		jQuery("#birthdayMessage").Hide() // Hides the message so that it can be faded in later
		var speed float64 = 1             // Sets speed of how fast the hue changes, and how fast the image is rotated, which is incremented by 0.1 every time the for loop runs intil it reaches > 49.9
		go func() {
			for i := 0; true; i = i + 2 { // Increments i by 2 to speed everything up, however this results in it being slightly choppier, but requires less speed and thus less resources
				jQuery("#alexHead").SetAttr("style", fmt.Sprintf("filter: hue-rotate(%vdeg);", i)) // Rotates the hue by whatever i is
				js.Global.Call("$", "#alexHead").Call("rotate", i)                                 // Calls jQueryRotate's rotate function to rotate the image by whatever i is
				// If the speed is less than 49.9, wait for 50 - speed milliseconds and increment speed by 0.1
				if speed < 49.9 {
					time.Sleep(time.Duration(50-speed) * time.Millisecond) // Waits for 50 - speed milliseconds
					speed = speed + 0.1                                    // Increments speed by 0.1
				} else { // If speed is greater than or equal to 49.9, wait for 0.1 milliseconds and do not increment speed further
					time.Sleep(100 * time.Microsecond) // Sleeps for 0.1 milliseconds, or 100 microseconds
				}
			}
		}()
		time.Sleep(3 * time.Second)              // Waits for 3 seconds
		jQuery("#birthdayMessage").FadeIn(10000) // Over the span of 10 seconds, fade the paragraph #birthdayMessage with text "Happy Birthday!" in
	}()
}

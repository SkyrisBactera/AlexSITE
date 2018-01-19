package main

import (
	"fmt"
	// Converting strings to other data types and vice-versa
	"github.com/gopherjs/jquery"
)

// Hides elements that are only shown in certain instances, or until they are ready. Only hides elements
func hideElements() {
	go func() {
		// If development mode is disabled, then hide the development warning message
		if !devel {
			jQuery("#develWarning").Hide() // Hides development warning message
		}
		// For loop that iterates over elementsToHide and hides each of it's elements
		for _, elem := range elementsToHide {
			jQuery(elem).Hide() // Hides the element
		}
	}()
}

// Fogs the main content
func hideContent() {
	go func() {
		jQuery("#blurme").AddClass("fogged") // Fogs the main content
	}()
}

// Show changelog and upgraded version with the previous revision as float64
func showChanges(previousRev float64) {
	go func() {
		// If the previous revision isn't the same as the current revision, show changelog
		if previousRev != revision {
			fmt.Println("Changes detected")                                                                         // Report to the console that the version has upgraded since last time
			jQuery("#changes").Show()                                                                               // Show div #changes which contains #versionChange and #changeList
			jQuery("#versionChange").SetText(fmt.Sprintf("Upgraded from version %v to %v!", previousRev, revision)) // Set the text of h6 #versionChange to show the version upgrade
			// For each change in the changelog add a new list element with that
			for _, val := range changelog {
				jQuery("#changeList").Append(fmt.Sprintf("<li>%s</li>", val)) // Add new list element with val in the #changeList unordered list
			}
		}
	}()
}

// Returns a jQuery object that is a MDL textbox
func generateTextBox(id int) jquery.JQuery {
	textField := jQuery("<div></div>")                           // Create a div that holds the textbox
	textField.SetAttr("class", "mdl-textfield mdl-js-textfield") // Assign the div necessary MDL classes
	input := jQuery("<input></input>")                           // Create input
	input.SetAttr("class", "mdl-textfield__input")               // Assign the input necessary MDL classes
	input.SetAttr("type", "text")                                // Set type as text for a textbox
	input.SetAttr("id", fmt.Sprintf("chal%vText", id))           // Give it an ID based on the int id input so that it can be used elsewhere in the code
	label := jQuery("<label></label>")                           // Create a label
	label.SetAttr("class", "mdl-textfield__label")               // Assign the label necessary MDL classes
	label.SetAttr("for", fmt.Sprintf("chal%vText", id))          // Connect the label with the input
	label.SetText("Text...")                                     // Add placeholder text if nothing has been entered yet
	textField.Append(input, label)                               // Add the input and label to the holder div
	return textField                                             // Return this jQuery object
}

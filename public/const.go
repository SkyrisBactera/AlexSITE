package main

var (
	// Misc
	elementsToHide = [...]string{
		"#changes",         // Changelog
		"#birthdaySpecial", // Div containing the birthday special (rotating head, Happy Birthday text)
		"#specialEvent",    // h5 that says "Special Event"
	}
	// Array of possible easter eggs. If the inputted string must contain two or more different strings, those two or more are delimited by a %
	easterEggs = [...]string{
		"crush%you^<3",
		"morgan",
		"bailey^;-)",
		"alex%root",
		"davis%davalos%delosh",
		"test",
		"consummation",
		"jasica",
		"anal%probe",
	}
	// Element-delimited list of changes since the last revision
	changelog = [...]string{ // Used ... to specify that the amount of elements the array should contain should match the amount of elements specified
		"Cleaned up code",
		"Better in-code documentation",
		"Removed unused functions",
		"Removed redundant code in general",
		"Removed unecessary code",
		"Made non-variable variables constants for better code security",
		"Made repetitive code use for loops for readability and better scalability",
		"Added more easter eggs",
		"Seperated different sections of code into different files for easier reading",
		"Fixed bug where if multiple easter eggs were unlocked at once, only the last one would be reported",
		"With the bug fix above, also made report fade out after 4 seconds",
		"Easter egg parsing now supports custom messages and uses slice for better scalability",
		"Added more readability for challenges by bolding the Background, Info, etc.",
	}
)

const (
	// Amounts that exist of something
	challengeSize = 3               // How many challenges are implemented for use in other parts of the code
	easterSize    = len(easterEggs) // How many easter eggs are implemented for use in other parts of the code, and for the user
	// Version History
	revision = 0.3   // Used to mark updates of the website, and will appear in the changelog. Increases everyday changes are made
	devel    = false // Used to determine if debugging mode should be used, and also activates development warning
)

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
		"bailey^;-) for two reasons",
		"alex%root",
		"davis%davalos%delosh",
		"test",
		"consummation",
		"jasica",
		"anal%probe",
		"booty",
		"rakesh",
		"chris",
		"github",
		"creep",
		"alida",
		"maya",
		"maya%alex^;-)",
		"cooper",
		"maxine",
		"payden^Yeeet",
		"yeet",
		"dead%fish",
		"tag",
		"skyrisbactera",
		"truth%or%dare",
		"never%find^That's right!",
		"ux985",
	}
	// Element-delimited list of changes since the last revision
	changelog = [...]string{ // Used ... to specify that the amount of elements the array should contain should match the amount of elements specified
		"Published to Github for better reliability and openness",
		"Added 17 more easter eggs",
		"New challenges!",
		"Fixed bug where decode appears twice",
	}
)

const (
	// Amounts that exist of something
	challengeSize = 5               // How many challenges are implemented for use in other parts of the code
	easterSize    = len(easterEggs) // How many easter eggs are implemented for use in other parts of the code, and for the user
	// Version History
	revision = 0.4  // Used to mark updates of the website, and will appear in the changelog. Increases everyday changes are made
	devel    = true // Used to determine if debugging mode should be used, and also activates development warning
)

package main

import (
	"fmt"
	"strconv" // Converting strings to other data types and vice-versa

	"github.com/gopherjs/gopherjs/js"
)

var completeChallenges []int      // Used to track the challenges the user has completed, as well as the easter eggs the user has figured out
var answers [challengeSize]string // Array that holds the answers to the challenges

// Holds functions with the substance of the challenges
func loadChallenges() {
	/*
		Possible chalTypes:
		 0: Challenge where all information is a string
		 1: Challenge where you specify a file for usage
	*/
	addChallenge(0, "ASCII Alteration", "To represent characters as a number, ASCII was invented to assign each character to a number. For this, you must simply convert this back to text (after deciding what the cryptic \"Number Multiplier\" means)", "The number multiplier is 2", "230 234 198 198 202 230 230", 0, "success")
	addChallenge(1, "Musical Mystery", "What a mystery?", "Everything is black and white", "shaal.wav", 1, "alex")
	addChallenge(0, "Cryptographic Craziness", "In the age of technology and the ever-growing need for heightened security and privacy, the \"Cryptographic hash function\" became popular across every service you may find on the internet. The purpose is to be able to tell if two pieces of data are the same, while not revealing what those two pieces were. It does this by performing a variety of mathematical operations on the data that cannot easily be reversed. An example function would be \"x * (x/2)\" where x is the input data. While this is a very easy operation to perform, it is very hard to predict what x was from the output.  Most companies use this for password authentication. For example, if your password for Google was \"O! You live!\", it would be stored as 4CE56E4DA84F0EBFCBD51FACAE04C706 (this is a MD5 hash, which is a very popular cryptographic hash function) which cannot be used to figure out the original password. This means that hackers who somehow get a hold of these stored \"hashes\", would be unable to produce or figure out the original password in order to log into the website.", "Your task is to create a cryptographic hash function. The requirement here is to create a function that has an input named x that outputs a number that cannot be multiplied, divided, etc. to figure out x", "When you are complete, zip me an email at davis.davalos.delosh@gmail.com with your function, and I will give you the password for this challenge", 2, "yickerhicker")
	addChallenge(1, "Code Comprehension", "The code for this website is written in a language called Go (by Google). It is a perfect example of a programming idea called abstraction, where repetitive tasks are put into objects called functions, that allow those repetitive tasks to be accomplished in a single line of code, using prewritten code. If my code is descriptive enough, it should be quite easy to find out the answer to this", "Your job is to find out the answer to this question by reading the code for this part of the website. My code is split up into 9 different files pertaining to what it deals with. The following snippet of code handles adding challenges to the website, however I have removed the answers to other challenges for obvious reasons.", "Chal4Snippet.txt", 3, "racketball43")
	addChallenge(0, "Reality", "", "(i'm sorry, this ones impossible without cheating (which I encourage you to try). Why would I do that? I get the same feeling as if I got something off my chest from telling somebody else something I've been holding in, but without actually telling anybody. It won't work. I know)", "Gucci Gang, ooh, yeah, Lil Pump, yeah, Gucci Gang, ooh Gucci gang, Gucci gang, Gucci gang, Gucci gang Gucci gang, Gucci gang, Gucci gang (Gucci gang!) Spend ten racks on a new chain My bitch love do cocaine, ooh I fuck a bitch, I forgot her name I can't buy a bitch no wedding ring Rather go and buy Balmains Gucci gang, Gucci gang, Gucci gang (Gucci gang!) Gucci gang, Gucci gang, Gucci gang, Gucci gang Gucci gang, Gucci gang, Gucci gang (Gucci gang!) Spend ten racks on a new chain My bitch love do cocaine, ooh I fuck a bitch, I forgot her name, yeah I can't buy no bitch no wedding ring Rather go and buy Balmains, aye Gucci gang, Gucci gang, Gucci gang My lean cost more than your rent, ooh Your mama still live in a tent, yeah Still slanging dope in the jets, huh Me and my grandma take meds, ooh None of this shit be new to me Fucking my teacher, call it tutory Bought some red bottoms, cost hella Gs Fuck your airline, fuck your company Bitch, your breath smell like some cigarettes I'd rather fuck a bitch from the projects They kicked me out the plane off a Percocet Now Lil Pump fly a private jet Everybody screaming \"fuck West Jet!\" Lil Pump still sell that meth Hunnid on my wrist sippin on Tech Fuck a lil bitch, make her pussy wet Gucci gang, Gucci gang, Gucci gang, Gucci gang Gucci gang, Gucci gang, Gucci gang (Gucci gang!) Spend ten racks on a new chain My bitch love do cocaine, ooh I fuck a bitch, I forgot her name I can't buy a bitch no wedding ring Rather go and buy Balmains Gucci gang, Gucci gang, Gucci gang (Gucci gang!) Gucci gang, Gucci gang, Gucci gang, Gucci gang Gucci gang, Gucci gang, Gucci gang (Gucci gang!) Spend ten racks on a new chain My bitch love do cocaine, ooh I fuck a bitch, I forgot her name I can't buy no bitch no wedding ring Rather go and buy Balmains, aye Gucci gang, Gucci gang, Gucci gang Lil Pump, yeah, Lil Pump, ooh", 4, "I feel like I'm trapped in a cage of self deprecation. All my actions end up hurting myself, i feel like nothing I will do will fix anything. I need somebody to love, and somebody to love me back other than my parents, but that isn't happening any time soon, now is it?")
}

func checkChallenge(id int) {
	if jQuery("#chal"+strconv.Itoa(id)+"Text").Val() == answers[id] {
		js.Global.Call("alert", "Correct")
		completeChallenges = append(completeChallenges, id)
	} else {
		js.Global.Call("alert", "Wrong")
	}
}

func updateChallengeCompletion() {

}

func addChallenge(chalType int, title string, description string, info string, textToDecode string, id int, answer string) {
	fmt.Println(len(answers))
	answers[id] = answer
	h4Title := jQuery("<h4></h4>")
	if contains(completeChallenges, id) {
		fmt.Println("Completed")
		h4Title.SetText("Challenge " + strconv.Itoa(id+1) + " (Complete)")
	} else {
		h4Title.SetText("Challenge " + strconv.Itoa(id+1))
	}
	jQuery("#content").Append(h4Title)
	divID := "chal" + strconv.Itoa(id)
	div := jQuery("<div></div>")
	div.SetAttr("class", "block")
	div.SetAttr("id", divID)
	h5 := jQuery("<h5></h5>")
	h5.SetText(title)
	backgroundB := jQuery("<b></b>")
	backgroundB.SetText("Background: ")
	descP := jQuery("<p></p>")
	descP.SetText(description)
	infoB := jQuery("<b></b>")
	infoB.SetText("Info: ")
	hintsP := jQuery("<p></p>")
	hintsP.SetText(info)
	decodeB := jQuery("<b></b>")
	decodeB.SetText("Decode: ")
	br := jQuery("<br />")
	// Submit Button
	subButton := jQuery("<button></button>")
	subButton.SetAttr("onclick", fmt.Sprintf("go.checkChallenge(%v)", id))
	subButton.SetAttr("class", "mdl-button mdl-js-button mdl-button--raised mdl-js-ripple-effect mdl-button--accent")
	subButton.SetAttr("type", "submit")
	subButton.SetText("Submit")
	// Hide/Unhide Button
	//<button onclick="go.chal1()" id="chal1But" class="mdl-button mdl-js-button mdl-button--raised mdl-js-ripple-effect mdl-button--accent" type="button">BEGIN</button>
	hideButton := jQuery("<button></button>")
	hideButton.SetAttr("onclick", fmt.Sprintf("$(\"#%s\")", divID)+".toggle();")
	hideButton.SetAttr("class", "mdl-button mdl-js-button mdl-button--raised mdl-js-ripple-effect mdl-button--accent")
	hideButton.SetAttr("type", "button")
	hideButton.SetText("BEGIN")
	// Add generated elements
	if chalType == 1 {
		downloadDiv := jQuery("<div></div>")
		downloadDiv.SetAttr("class", "mdl-cell mdl-cell--1-col")
		downloadButton := jQuery("<button></button>")
		downloadButton.SetAttr("onclick", fmt.Sprintf("downloadFile(\"%v\")", textToDecode))
		downloadButton.SetAttr("class", "mdl-button mdl-js-button mdl-button--raised mdl-js-ripple-effect mdl-button--accent")
		downloadButton.SetAttr("type", "button")
		downloadButton.SetText("DOWNLOAD")
		downloadDiv.Append(downloadButton)
		div.Append(h5, backgroundB, descP, infoB, hintsP, decodeB, downloadDiv, generateTextBox(id), subButton)
	} else {
		textP := jQuery("<p></p>")
		textP.SetText(textToDecode)
		div.Append(h5, backgroundB, descP, infoB, hintsP, decodeB, textP, generateTextBox(id), subButton)
	}
	div.Hide()
	jQuery("#content").Append(div, br, hideButton)
	js.Global.Get("window").Get("componentHandler").Call("upgradeDom")
}

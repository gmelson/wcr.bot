// testWit
package nl

import (
	"log"
	"strings"
	"testing"
)

//{"adoptBeach","do you support community activities"},
//{"availability","where can i buy your coffee"},
//{"buyCoffee","i want to buy"},
//{"buyType","dark roast"},
//{"competition","who do you compete with"},
//{"cuppingEvent","when is your next cupping"},
//{"introduction","Hi"},
//{"roast","what is dark roast"},
//{"shipping","when will I get my coffee"}

func TestGetIntent(t *testing.T) {
	var intents = [2][2]string{{"about", "what do you do"}, {"activities", "what are you doing this weekend"}}
	m := new(Message)
	for _, intent := range intents {
		log.Printf("Ready to send phrase = %s\n", intent[1])
		err := m.GetIntent(intent[1])
		log.Println("Done with GetIntent")
		if err != nil || m.Entities.Intent[0].Value != intent[0] {
			t.Errorf("Intent does not match %s\n for phrase %s\n", m.Entities.Intent[0].Value, intent[0])
		}
		log.Printf("Intent value %s\n", m.Entities.Intent[0].Value)
	}
}

func TestReplaceWords(t *testing.T) {
	var sentences = [...]string{"we work for Weekend Coffee Roasters", "we work for Weekend Coffee", "Weekend Coffee is roasting", "we work for wcr on the Weekend"}
	m := new(Message)
	for _, sentence := range sentences {
		err, replaced := m.ReplaceWords(sentence)
		if err != nil || strings.Contains(replaced, "wcr") == false {
			t.Errorf("String not replaced correctly with wcr; orig = %s\nreplaced = %s\n\n", sentence, replaced)
		}
		log.Printf("Replaced %s\n with new string %s\n\n", sentence, replaced)
	}

}

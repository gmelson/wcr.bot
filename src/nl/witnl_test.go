// testWit
package nl

import (
	"log"
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
			t.Errorf("Intent does not match %s for phrase %s\n", m.Entities.Intent[0].Value, intent[0])
		}
		log.Printf("Intent value %s", m.Entities.Intent[0].Value)
		//fmt.Printf("Intent value %s", m.Entities.Intent[0].Value)
	}
}

// answers
package nl

const (
	about        = "Weekend Coffee Roasters is a small batch roaster located in San Jose California."
	activities   = "Weekend Coffee Roasters will be participating in Makers Market on June 11, located in Santana Row."
	availability = "You can buy on www.amazon.com or our website http://www.weekendcoffeeroasters.com/buy.html"
	adoptBeach   = "We've adopted Seabright!  Our next clean up is in June."
	beanType     = "Thank you!"
	buyCoffee    = "What type of roast would you like?"
	shipping     = "Shipping cost $5.95 and usually takes four days."
	cuppingEvent = "You just missed one.  The next cupping is in three weeks."
	introduction = "Hello, I am the WCR bot.  Ask me a question about our roasting"
	roast        = "We roast an excellent medium roast."
	receipt		 = "The total for your order of %s is %s"
	order		 = "We have recieved your order for %s.  Your order should arrive %s."	
)

type Answer struct {
	Message Message
}

func (a *Answer) GetAnswer() (err error, resp string) {
	val := a.Message.Entities.Intent[0].Value
	switch {
	case val == "about":
		resp = about
	case val == "activities":
		resp = activities
	case val == "availability":
		resp = availability
	case val == "adoptBeach":
		resp = adoptBeach
	case val == "beanType":
		resp = beanType
	case val == "buyCoffee":
		resp = buyCoffee
	case val == "shipping":
		resp = shipping
	case val == "cuppingEvent":
		resp = cuppingEvent
	case val == "introduction":
		resp = introduction
	case val == "roast":
		resp = roast
	}
	return
}

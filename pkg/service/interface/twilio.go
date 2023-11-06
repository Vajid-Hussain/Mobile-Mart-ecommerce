package serviceInterface

type Ijwt interface{
	TwilioSetup()
	SendOtp(string) (string, error)
}
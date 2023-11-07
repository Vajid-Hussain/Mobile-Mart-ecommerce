package serviceInterface

type Ijwt interface{
	TwilioSetup()
	SendOtp(string) (string, error)
	VerifyOtp(string, string) (error)
}
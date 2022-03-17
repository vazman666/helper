package pkg

func Firm(firm string) string {
	var tmp string
	switch firm {
	case "JD":
		tmp = "Just Drive"
	case "Sangsin Brake":
		tmp = "Sangsin"
	case "HSB":
		tmp = "Hong Sung Brake"
	default:
		tmp = firm
	}
	return tmp
}

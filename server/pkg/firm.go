package pkg

func Firm(firm string) string {
	var tmp string
	switch firm {
	case "JD":
		tmp = "Just Drive"
	case "Sangsin Brake":
		tmp = "Sangsin"
	default:
		tmp = firm
	}
	return tmp
}

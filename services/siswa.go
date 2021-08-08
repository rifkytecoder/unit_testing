package services

func grade(nilai int) string {

	if nilai <= 40 {
		return "C"
	} else if nilai > 40 && nilai <= 70 {
		return "B"
	} else {
		return "A"
	}

}

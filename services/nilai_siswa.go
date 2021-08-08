package services

//mengunakan kondisi Switch
func hasil(nilai int) string {

	switch {
	case nilai < 40:
		return "C"
	case nilai > 40 && nilai <= 70:
		return "B"
	default:
		return "A"
	}
}

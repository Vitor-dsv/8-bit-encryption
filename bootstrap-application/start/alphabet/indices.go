package alphabet

func GetAlphabetIndices() map[string]string {
	indices := make(map[string]string)

	indices[" "] = "00"
	indices["A"] = "01"
	indices["B"] = "02"
	indices["C"] = "03"
	indices["D"] = "04"
	indices["E"] = "05"
	indices["F"] = "06"
	indices["G"] = "07"
	indices["H"] = "08"
	indices["I"] = "09"
	indices["J"] = "10"
	indices["K"] = "11"
	indices["L"] = "12"
	indices["M"] = "13"
	indices["N"] = "14"
	indices["O"] = "15"
	indices["P"] = "16"
	indices["Q"] = "17"
	indices["R"] = "18"
	indices["S"] = "19"
	indices["T"] = "20"
	indices["U"] = "21"
	indices["V"] = "22"
	indices["W"] = "23"
	indices["X"] = "24"
	indices["Y"] = "25"
	indices["Z"] = "26"

	return indices
}

func GetKeyOfValue(Indices map[string]string, value string) string {
	for letter, index := range Indices {
		if value == index {
			return letter
		}
	}

	return ""
}

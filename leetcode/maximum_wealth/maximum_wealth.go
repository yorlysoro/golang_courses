package maximumwealth

func maximumWealth(accounts [][]int) int {
	m := len(accounts)
	if m < 1 || m > 50 {
		return 0
	}

	maxWealth := 0
	for _, customer := range accounts {
		n := len(customer)
		if n < 1 || n > 50 {
			return 0
		}

		currentWealth := 0
		for _, bankAccount := range customer {
			if bankAccount < 1 || bankAccount > 100 {
				return 0
			}
			currentWealth += bankAccount
		}

		if currentWealth > maxWealth {
			maxWealth = currentWealth
		}
	}
	return maxWealth
}

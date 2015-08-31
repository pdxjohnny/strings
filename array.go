package strings

func Contains(list []string, check string) (bool) {
	if list != nil {
		for _, item := range list {
			if item == check {
				return true
			}
		}
	}
	return false
}


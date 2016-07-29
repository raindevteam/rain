package rain

func isIn(list []string, el string) bool {
	for _, v := range list {
		if v == el {
			return true
		}
	}
	return false
}

func listPrint(list []string) string {
	str := ""
	for _, v := range list {
		str = str + "    " + v + "\n"
	}
	return str
}

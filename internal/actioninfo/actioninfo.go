package actioninfo

type DataParser interface {
	// TODO: добавить методы
	Parse()
	ActionInfo()
}

func Info(dataset []string, dp DataParser) {
	// TODO: реализовать функцию
	for _, line := range dataset {
		dp.Parse(line)
	}
}

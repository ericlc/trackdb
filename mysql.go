package main

type MySQL struct{}

func (m MySQL) Execute(es ExecutableSQL) bool {
	return true
}

func callClient(sql string) bool {
	return true
}

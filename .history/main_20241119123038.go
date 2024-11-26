package main

func main() {
	err := nunu.Execute()
	if err != nil {
		fmt.Println("execute error: ", err.Error())
	}
}
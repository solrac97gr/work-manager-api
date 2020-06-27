package main

func main() {
	if !database.ConnectionOK() {
		log.Fatal("Not connected to DB")
		return
	}
	handlers.Handlers()
	if err != nil {
		print(err.Error())
	}
}

package main

import "absent.com/absentapi/app"

func main() {
	err := app.SetupAndRunApp()
	if err != nil {
		panic(err)
	}
}

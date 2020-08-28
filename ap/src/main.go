package main

import "echo-sample/infra"

func main() {
	// e := echo.New()
	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello, Echo World!!!!!!")
	// })
	// e.Logger.Fatal(e.Start(":8081"))
	infra.Init()
}

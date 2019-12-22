package iris

import (
	"fmt"
	"net/http"

	"github.com/kataras/iris/v12"
)

func Example() {
	app := iris.New()
	app.Get("/", func(ctx iris.Context) {
		ctx.StatusCode(http.StatusOK)
		_, err := ctx.JSON(iris.Map{"message": "Hello, world!"})
		if err != nil {
			fmt.Println(err)
		}
	})
	if err := app.Run(iris.Addr(":8000")); err != nil {
		panic(err)
	}
}

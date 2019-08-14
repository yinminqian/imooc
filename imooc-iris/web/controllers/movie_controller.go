package controllers

import (
	"github.com/kataras/iris/mvc"
	"imooc-iris/repositories"
	"imooc-iris/services"
)

type MovieController struct {
}

func (c *MovieController) Get() mvc.View {

	movieRepository := repositories.NewMovieManager()
	moiveServe := services.NewMovieServiceManger(movieRepository)
	moiveResult := moiveServe.ShowMovieName()
	return mvc.View{
		Name: "movie/index.html",
		Data: moiveResult,
	}
}

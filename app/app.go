package app

import (
	"fmt"
	"movies/config"

	"github.com/gin-gonic/gin"

	movieHttp "movies/handler/movies"
	movieRepository "movies/repository/movies"
	movieUsecase "movies/usecase/movies"
)

func NewApp() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()

	// connection DB
	// db := config.NewConnection()

	// repository
	movieRepository := movieRepository.NewMovieRepository()

	// usecase
	movieUsecase := movieUsecase.NewMovieUsecase(movieRepository)

	// handler
	movieHandler := movieHttp.NewMovieHandler(movieUsecase)

	// router
	movieHandler.Route(engine)

	fmt.Println("Running on port", config.ENV.AppPort)
	return engine
}

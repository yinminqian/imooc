package repositories

import "imooc-iris/datamodels"

type MovieRepository interface {
	GetMovieName() string
}

type MovieManager struct {
}

func NewMovieManager() MovieRepository {
	return &MovieManager{}
}

func (m *MovieManager) GetMovieName() string {
	//模拟赋值给模型
	movie := &datamodels.Movies{Name: "大闹天宫"}
	return  movie.Name
}

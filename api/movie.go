package api

import (
	"github.com/bosemian/mv-flix/model"
)

type Movier interface {
	List() (*MovieList, error)
	Get(id string) (*MovieData, error)
	AddFavorite(id string) (*MovieData, error)
}

type MovieList struct {
	Movies []*model.Movie `json:"movies"`
}

type MovieData struct {
	Movie *model.Movie `json:"movie"`
}

func (m *MovieList) List() (*MovieList, error) {
	movies := model.ListMovies()
	m.Movies = movies
	return m, nil
}

func (md *MovieData) Get(id string) (*MovieData, error) {
	movie := model.GetMovie(id)
	md.Movie = movie
	return md, nil
}

func (md *MovieData) AddFavorite(id string) (*MovieData, error) {
	movie := model.GetMovie(id)
	md.Movie = movie
	md.Movie.Favorite = !md.Movie.Favorite
	return md, nil
}

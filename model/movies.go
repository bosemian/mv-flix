package model

// struct for movie
type Movie struct {
	Id           string `json:"id"`
	Title        string `json:"title"`
	Subtitle     string `json:"subtitle"`
	Description  string `json:"description"`
	LargeImgSrc  string `json:"largeImgSrc"`
	SmallImgSrc  string `json:"smallImgSrc"`
	RealeaseDate string `json:"releaseDate"`
	Duration     string `json:"duration"`
	Genre        string `json:"genre"`
	TrailerPath  string `json:"trailerPath"`
	Favorite     bool   `json:"favorite"`
}

var movies = []*Movie{
	&Movie{
		Id:           "dunkirk",
		Title:        "Dunkirk",
		Subtitle:     "Dunkirk",
		Description:  "Miraculous evacuation of Allied soldiers from Belgium, Britain, Canada, and France, who were cut off and surrounded by the German army from the beaches and harbor of Dunkirk, France, during the Battle of France in World War II",
		LargeImgSrc:  "https://image.tmdb.org/t/p/w780/fudEG1VUWuOqleXv6NwCExK0VLy.jpg",
		SmallImgSrc:  "https://image.tmdb.org/t/p/w185/fudEG1VUWuOqleXv6NwCExK0VLy.jpg",
		RealeaseDate: "July 21 2017",
		Duration:     "1hr 46min",
		Genre:        "Action, Drama, History",
		TrailerPath:  "https://www.youtube.com/embed/F-eMt3SrfFU",
		Favorite:     false,
	},
	&Movie{
		Id:           "interstellar",
		Title:        "Interstellar",
		Subtitle:     "Interstellar",
		Description:  "Interstellar chronicles the adventures of a group of explorers who make use of a newly discovered wormhole to surpass the limitations on human space travel and conquer the vast distances involved in an interstellar voyage.",
		LargeImgSrc:  "https://image.tmdb.org/t/p/w780/xu9zaAevzQ5nnrsXN6JcahLnG4i.jpg",
		SmallImgSrc:  "https://image.tmdb.org/t/p/w185/xu9zaAevzQ5nnrsXN6JcahLnG4i.jpg",
		RealeaseDate: "November 7 2014",
		Duration:     "2hr 49min",
		Genre:        "Adventure, Drama",
		TrailerPath:  "https://www.youtube.com/embed/zSWdZVtXT7E",
		Favorite:     false,
	},
	&Movie{
		Id:           "the-dark-knight-rises",
		Title:        "The Dark Knight Rises",
		Subtitle:     "TDKR",
		Description:  "Batman encounters the mysterious Selina Kyle and the villainous Bane, a new terrorist leader who overwhelms Gotham's finest. The Dark Knight resurfaces to protect a city that has branded him an enemy.",
		LargeImgSrc:  "https://image.tmdb.org/t/p/w185/3bgtUfKQKNi3nJsAB5URpP2wdRt.jpg",
		SmallImgSrc:  "https://image.tmdb.org/t/p/w780/3bgtUfKQKNi3nJsAB5URpP2wdRt.jpg",
		RealeaseDate: "July 20 2012",
		Duration:     "2hr 44min",
		Genre:        "Action, Thriller",
		TrailerPath:  "https://www.youtube.com/embed/g8evyE9TuYk",
		Favorite:     false,
	},
	&Movie{
		Id:           "inception",
		Title:        "Inception",
		Subtitle:     "Inception",
		Description:  "Cobb, a skilled thief is offered a chance to regain his old life as payment for a task considered to be impossible: \"inception\", the implantation of another person's idea into a target's subconscious.",
		LargeImgSrc:  "https://image.tmdb.org/t/p/w780/s2bT29y0ngXxxu2IA8AOzzXTRhd.jpg",
		SmallImgSrc:  "https://image.tmdb.org/t/p/w185/s2bT29y0ngXxxu2IA8AOzzXTRhd.jpg",
		RealeaseDate: "July 10 2010",
		Duration:     "2hr 28min",
		Genre:        "Action, Adventure, Sci-Fi",
		TrailerPath:  "https://www.youtube.com/embed/8hP9D6kZseM",
		Favorite:     false,
	},
	&Movie{
		Id:           "the-prestige",
		Title:        "The Prestige",
		Subtitle:     "Prestige",
		Description:  "A mysterious story of two magicians whose intense rivalry leads them on a life-long battle for supremacy - to create the ultimate illusion whilst sacrificing everything they have to outwit the other",
		LargeImgSrc:  "https://image.tmdb.org/t/p/w780/c5o7FN2vzI7xlU6IF1y64mgcH9E.jpg",
		SmallImgSrc:  "https://image.tmdb.org/t/p/w185/c5o7FN2vzI7xlU6IF1y64mgcH9E.jpg",
		RealeaseDate: "October 20 2006",
		Duration:     "2hr 10min",
		Genre:        "Drama, Mystery, Sci-Fi",
		TrailerPath:  "https://www.youtube.com/embed/ijXruSzfGEc",
		Favorite:     false,
	},
}

func ListMovies() []*Movie {
	return movies
}

func GetMovie(id string) *Movie {
	for _, movie := range movies {
		if movie.Id == id {
			return movie
		}
	}
	return nil
}

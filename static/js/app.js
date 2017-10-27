// varibale
var api = axios.create({
  headers: { 'Access-Control-Allow-Origin': '*' }
})

// MovieTrailer Component
const MovieTrailer = {
  template: `
    <div class="trailer-body" style="background: #1e1d1d">
      <div class="has-text-centered">
        <div class="columns">
          <div class="column vertical-align">
            <iframe
              allowFullScreen
              frameborder="0"
              height="376"
              :src="trailerUrlPath"
              style="width: 100%; min-width: 536px"
              />
          </div>
        </div>
      </div>
    </div>`,
    created () {
      const id = this.$route.params.id
      api.get(`movie/${id}`)
        .then((res) => {
          this.trailerUrlPath = res.data.movie.trailerPath
        })
    },
    data () {
      return {
        trailerUrlPath: null
      }
    }
  }

  // Movie Component
  const Movie = {
    template:
      `<div
        :class="[{ 'favorite-shadow': selectedMovie.favorite }, 'hero-body']"
        :style="{ 'background-image': 'url('+ selectedMovie.largeImgSrc +')' }">
        <header class="nav">
          <div class="container">
            <div class="nav-left">
            <a class="nav-item">
              <i class="fa fa-bars" aria-hidden="true"></i>
            </a>
            <router-link to="/" class="nav-item is-active">
              Home
            </router-link>
            <a class="nav-item is-active">
              <span class="tag is-rounded">Films</span>
            </a>
            <router-link :to="{path: '/' + $route.params.id + '/trailer'}" class="nav-item is-active">
              Shows
            </router-link>
            </div>
            <div class="nav-right desktop txt-icon">
            <span class="nav-item">
              <a class="title">
              MVFlix
              </a>
            </span>
            </div>
          </div>
        </header>
        <div class="container description-container">
          <div class="columns">
            <div class="column is-three-quarters">
              <h1 class="title">{{ selectedMovie.title }}</h1>
              <h4 class="subtitle">
                <p class="subtitle-tag">{{ selectedMovie.duration }}</p>
                <p class="subtitle-tag">{{ selectedMovie.genre }}</p>
                <p class="subtitle-tag">{{ selectedMovie.releaseDate }}</p>
              </h4>
              <p class="description">{{ selectedMovie.description }}</p>
              <div class="links">
                <router-link
                  :to="{path: '/' + $route.params.id + '/trailer'}"
                  class="button play-button">
                  Play <i class="fa fa-play"></i>
                </router-link>
                <a
                  class="button is-link favorites-button"
                  @click="addToFavorites">
                  <span
                    :class="[{ 'hide': selectedMovie.favorite }]">
                  Add to
                  </span>
                  <span
                    :class="[{ 'hide': !selectedMovie.favorite }]">
                  Remove from
                  </span>
                    &nbsp;favorites
                  <i class="fa fa-plus-square-o"></i>
                </a>
              </div>
            </div>
          </div>
        </div>
      </div>`,
    data () {
      return {
        selectedMovie: null
      }
    },
    created () {
      this.selectMovie()
    },
    watch: {
      $route () {
        this.selectMovie()
      }
    },
    methods: {
      selectMovie () {
        const id = this.$route.params.id
        api.get(`movie/${id}`)
          .then((res) => {
            this.selectedMovie = res.data.movie
          })
      },
      addToFavorites() {
        api.post('favorite', { id: this.$route.params.id })
          .then((res) => {
            this.selectedMovie.favorite = res.data.movie.favorite
            // emit to footer page render select movie favorite
            rootApp.movieChoices.map((movie) => {
              if (this.$route.params.id === movie.id) {
                movie.favorite = res.data.movie.favorite
              }
              return movie
            })
          })
      }
    }
  }
  // Intro Component
  const Intro = {
    template:
      `<div class="hero-body" style="background: #1e1d1d">
        <div class="container has-text-centered">
          <div class="columns">
            <div class="column is-half is-offset-one-quarter vertical-align">
              <h1 class="home-intro">
                MVFlix
              </h1>
              <p class="home-subintro">Select a movie below from the list of critically acclaimed Christopher Nolan films.</p>
            </div>
          </div>
        </div>
      </div>`
  }

  // config router
  const router = new VueRouter({
    routes: [
      { path: '/', component: Intro },
      { path: '/:id', component: Movie },
      { path: '/:id/trailer', component: MovieTrailer },
      { path: '*', redirect: '/' }
    ]
  })
// Vue instance
const rootApp = new Vue({
    el: '#app',
    router,
    data () {
      return {
        movieChoices: null
      }
    },
    created () {
      api.get('movies')
        .then((res) => {
          this.movieChoices = res.data.movies
        })
    }
})

package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

// ask for one here https://www.internationalshowtimes.com/signup.html
const YOUR_API_KEY = "YOUR_API_KEY"

type Showtimes struct {
	Showtime []struct {
		ID               string      `json:"id"`
		CinemaID         string      `json:"cinema_id"`
		MovieID          string      `json:"movie_id"`
		StartAt          time.Time   `json:"start_at"`
		Language         string      `json:"language"`
		SubtitleLanguage interface{} `json:"subtitle_language"`
		Auditorium       interface{} `json:"auditorium"`
		Is3D             bool        `json:"is_3d"`
		IsImax           bool        `json:"is_imax"`
		BookingType      interface{} `json:"booking_type"`
		BookingLink      interface{} `json:"booking_link"`
	} `json:"showtimes"`
}

type MovieWrapper struct {
	Movie struct {
		ID                   string      `json:"id"`
		Slug                 string      `json:"slug"`
		Title                string      `json:"title"`
		PosterImageThumbnail interface{} `json:"poster_image_thumbnail"`
		OriginalTitle        string      `json:"original_title"`
		OriginalLanguage     string      `json:"original_language"`
		Synopsis             string      `json:"synopsis"`
		Runtime              interface{} `json:"runtime"`
		Genres               []struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"genres"`
		PosterImage interface{} `json:"poster_image"`
		SceneImages []struct {
			ImageFiles []struct {
				URL  string `json:"url"`
				Size struct {
					Width  int `json:"width"`
					Height int `json:"height"`
				} `json:"size"`
			} `json:"image_files"`
		} `json:"scene_images"`
		Trailers     interface{} `json:"trailers"`
		Ratings      interface{} `json:"ratings"`
		AgeLimits    interface{} `json:"age_limits"`
		ReleaseDates struct {
			GB []struct {
				Locale string      `json:"locale"`
				Region interface{} `json:"region"`
				Date   string      `json:"date"`
			} `json:"GB"`
		} `json:"release_dates"`
		Website             interface{} `json:"website"`
		ProductionCompanies interface{} `json:"production_companies"`
		Keywords            interface{} `json:"keywords"`
		ImdbID              interface{} `json:"imdb_id"`
		TmdbID              string      `json:"tmdb_id"`
		RentrakFilmID       interface{} `json:"rentrak_film_id"`
		Cast                []struct {
			ID        string `json:"id"`
			Character string `json:"character"`
			Name      string `json:"name"`
		} `json:"cast"`
		Crew interface{} `json:"crew"`
	} `json:"movie"`
}

type CinemaWrapper struct {
	Cinema struct {
		ID        string `json:"id"`
		Slug      string `json:"slug"`
		Name      string `json:"name"`
		ChainID   string `json:"chain_id"`
		Telephone string `json:"telephone"`
		Website   string `json:"website"`
		Location  struct {
			Lat     float64 `json:"lat"`
			Lon     float64 `json:"lon"`
			Address struct {
				DisplayText string `json:"display_text"`
				Street      string `json:"street"`
				House       string `json:"house"`
				Zipcode     string `json:"zipcode"`
				City        string `json:"city"`
				State       string `json:"state"`
				StateAbbr   string `json:"state_abbr"`
				Country     string `json:"country"`
				CountryCode string `json:"country_code"`
			} `json:"address"`
		} `json:"location"`
		BookingType string `json:"booking_type"`
	} `json:"cinema"`
}

func insertNth(s string, n int) string {
	var buffer bytes.Buffer
	var n_1 = n - 1
	var l_1 = len(s) - 1
	for i, rune := range s {
		buffer.WriteRune(rune)
		if i%n == n_1 && i != l_1 {
			buffer.WriteRune(' ')
			buffer.WriteRune('#')
			buffer.WriteRune('\n')
			buffer.WriteRune('#')
			buffer.WriteRune(' ')
		}
	}
	return buffer.String()
}

func main() {

	fmt.Printf("#################################################################\n")
	fmt.Printf("####~~~~~~~~~~------>>> ERLANGE OV CINEMA <<<------~~~~~~~~~~####\n")

	// Create client
	client := &http.Client{}

	// Create request
	// req, err := http.NewRequest("GET", "https://api.internationalshowtimes.com/v4/movies/?countries=GB", nil)
	// req, err := http.NewRequest("GET", "https://api.internationalshowtimes.com/v4/cities/?countries=GB", nil)
	req, err := http.NewRequest("GET", "https://api.internationalshowtimes.com/v4/showtimes/?city_ids=495", nil)

	// Headers
	req.Header.Add("X-API-Key", YOUR_API_KEY)

	parseFormErr := req.ParseForm()
	if parseFormErr != nil {
		fmt.Println(parseFormErr)
	}

	// Fetch Request
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Failure : ", err)
	}

	// Read Response Body
	respBody, _ := ioutil.ReadAll(resp.Body)

	var showtimes Showtimes
	err = json.Unmarshal(respBody, &showtimes)
	if err != nil {
		panic(err)
	}
	for _, showtime := range showtimes.Showtime {
		// fmt.Println(showtime)
		if strings.EqualFold(showtime.Language, "EN") {
			req, err := http.NewRequest("GET", fmt.Sprintf("https://api.internationalshowtimes.com/v4/movies/%s", showtime.MovieID), nil)

			// Headers
			req.Header.Add("X-API-Key", "tiuf1o9NOXNlqm9KomNzxXpd39YKuEd9")

			parseFormErr := req.ParseForm()
			if parseFormErr != nil {
				fmt.Println(parseFormErr)
			}

			// Fetch Request
			resp, err := client.Do(req)

			if err != nil {
				fmt.Println("Failure : ", err)
			}

			// Read Response Body
			respBody, _ := ioutil.ReadAll(resp.Body)

			var movieWrapper MovieWrapper
			err = json.Unmarshal(respBody, &movieWrapper)
			if err != nil {
				panic(err)
			}

			req, err = http.NewRequest("GET", fmt.Sprintf("https://api.internationalshowtimes.com/v4/cinemas/%s", showtime.CinemaID), nil)

			// Headers
			req.Header.Add("X-API-Key", "tiuf1o9NOXNlqm9KomNzxXpd39YKuEd9")

			parseFormErr = req.ParseForm()
			if parseFormErr != nil {
				fmt.Println(parseFormErr)
			}

			// Fetch Request
			resp, err = client.Do(req)

			if err != nil {
				fmt.Println("Failure : ", err)
			}

			// Read Response Body
			respBody, _ = ioutil.ReadAll(resp.Body)

			var cinemaWrapper CinemaWrapper
			err = json.Unmarshal(respBody, &cinemaWrapper)
			if err != nil {
				panic(err)
			}

			fmt.Printf("#################################################################\n")
			fmt.Printf("# Title: %s\n# Description: \n# %s\n# Hour: %s\n", movieWrapper.Movie.Title, insertNth(movieWrapper.Movie.Synopsis, 61), showtime.StartAt.Format("Mon Jan _2 15:04:05 2006"))
			fmt.Printf("# IMDB: https://www.imdb.com/title/%s\n", movieWrapper.Movie.ImdbID)
			fmt.Printf("# Cinema: %s\n# Website: %s\n", cinemaWrapper.Cinema.Name, cinemaWrapper.Cinema.Website)
			fmt.Printf("# Telephone: %s\n# Street: %s\n", cinemaWrapper.Cinema.Telephone, cinemaWrapper.Cinema.Location.Address.Street)

			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Press Enter to continue...\n")
			reader.ReadString('\n')

		}

	}

	fmt.Printf("#################################################################\n")

}

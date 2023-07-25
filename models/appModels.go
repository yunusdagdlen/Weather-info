package models

type Employee struct {
	Id      int
	Name    string
	City    string
	Weather string
}

type apiConfig struct {
	myApiKey string `json:"myApiKey`
}

type weatherData struct {
	Request struct {
		Type     string `json:"type"`
		Query    string `json:"query"`
		Language string `json:"language"`
		Unit     string `json:"unit"`
	} `json:"request"`

	Current struct {
		Temperature         int      `json:"temperature"`
		WeatherIcons        []string `json:"weather_icons"`
		WeatherDescriptions []string `json:"weather_descriptions"`
	} `json:"current"`
}

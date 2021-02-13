package structs

// Config Layout of the config.json file
type Config struct {
	// StartUpSound file name of the sound that plays at startup
	StartUpSound string `json:"startup-sound"`

	// Name name of the user
	Name string
}


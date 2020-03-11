package cfg

import (
	"encoding/json"
	"os"
)

// FileProvider describe the file base loaded which load the configuration
type FileProvider struct {
	Filename string
}

// Provide implents the Provider interface
// for this case is going to open a json file and map the data
func (fp FileProvider) Provide() (map[string]string, error) {
	var config = make(map[string]string)
	file, err := os.Open(fp.Filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	if err := json.NewDecoder(file).Decode(&config); err != nil {
		return nil, err
	}
	return config, nil
}

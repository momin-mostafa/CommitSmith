package models

type Config struct {
	Model                string  `yaml:"model"`
	Host                 string  `yaml:"host"`
	Temperature          float64 `yaml:"temperature"`
	MaxOptions           int     `yaml:"max_options"`
	UseConventionalCommits bool   `yaml:"use_conventional_commits"`
}

type OllamaRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
	Stream   bool      `json:"stream"`
	Options  *Options  `json:"options,omitempty"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Options struct {
	Temperature float64 `json:"temperature"`
}

type OllamaResponse struct {
	Message struct {
		Content string `json:"content"`
	} `json:"message"`
	Done bool `json:"done"`
}

type ModelsResponse struct {
	Models []struct {
		Name string `json:"name"`
	} `json:"models"`
}

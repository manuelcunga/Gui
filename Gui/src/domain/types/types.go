package gui_types

type Messages struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Choices struct {
	Index   int `json:"index"`
	Message struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	}
}

type Request struct {
	Model     string     `json:"model"`
	Messages  []Messages `json:"messages"`
	MaxTokens int        `json:"max_tokens"`
}

type Response struct {
	ID        string    `json:"id"`
	Object    string    `json:"object"`
	CreatedAt int       `json:"created_at"`
	Choices   []Choices `json:"choices"`
}

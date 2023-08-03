package main

type Label struct {
	English  []string `json:"en,omitempty"`
	Japanese []string `json:"jp,omitempty"`
	None     []string `json:"none,omitempty"`
}

type Manifest struct {
	Context   string                   `json:"@context"`
	ID        string                   `json:"id"`
	Type      string                   `json:"type"`
	Label     Label                    `json:"label"`
	Summary   Label                    `json:"summary"`
	Metadata  []map[string]interface{} `json:"metadata,omitempty"`
	Thumbnail []map[string]interface{} `json:"thumbnail,omitempty"`
	Items     []Canvas                 `json:"items"`
}

type Canvas struct {
	ID     string           `json:"id"`
	Type   string           `json:"type"`
	Label  Label            `json:"label"`
	Height int              `json:"height"`
	Width  int              `json:"width"`
	Items  []AnnotationPage `json:"items"`
}

type AnnotationPage struct {
	ID    string       `json:"id"`
	Type  string       `json:"type"`
	Items []Annotation `json:"items"`
}

type Annotation struct {
	ID         string         `json:"id"`
	Type       string         `json:"type"`
	Motivation string         `json:"motivation"`
	Target     string         `json:"target,omitempty"`
	Body       AnnotationBody `json:"body"`
}

type AnnotationBody struct {
	Id      string  `json:"id"`
	Type    string  `json:"type"`
	Format  string  `json:"format"`
	Service Service `json:"service"`
	Height  int     `json:"height,omitempty"`
	Width   int     `json:"width,omitempty"`
}

type ResponseImageAPI struct {
	Id     string `json:"id"`
	Width  int    `json:"width,omitempty"`
	Height int    `json:"height,omitempty"`
}

type Service2 struct {
	Id   string `json:"@id"`
	Type string `json:"@type"`
}

type Service struct {
	Id      string     `json:"id"`
	Type    string     `json:"type"`
	Profile string     `json:"profile"`
	Service []Service2 `json:"service,omitempty"`
}

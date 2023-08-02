package main

type Manifest struct {
	Context   string                   `json:"@context"`
	ID        string                   `json:"id"`
	Type      string                   `json:"type"`
	Label     Label                    `json:"label"`
	Summary   Label                    `json:"summary"`
	Metadata  []map[string]interface{} `json:"metadata"`
	Thumbnail []map[string]interface{} `json:"thumbnail"`
	Items     []CanvasItem             `json:"items"`
}

type Label struct {
	English  []string `json:"en,omitempty"`
	Japanese []string `json:"jp,omitempty"`
	None     []string `json:"none,omitempty"`
}

type AnnotationItem struct {
	ID         string         `json:"id"`
	Type       string         `json:"type"`
	Motivation string         `json:"motivation,omitempty"`
	Target     string         `json:"target,omitempty"`
	Body       AnnotationBody `json:"body"`
}

type CanvasItem struct {
	ID     string           `json:"id"`
	Type   string           `json:"type"`
	Label  Label            `json:"label"`
	Height int              `json:"height"`
	Width  int              `json:"width"`
	Items  []AnnotationItem `json:"items"`
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
	Service []Service2 `json:"service"`
}

type AnnotationBody struct {
	Id      string    `json:"id"`
	Type    string    `json:"type"`
	Format  string    `json:"format"`
	Service []Service `json:"service"`
	Height  int       `json:"height"`
	Width   int       `json:"width"`
}

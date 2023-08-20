// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Config struct {
	ID        string             `json:"id"`
	Name      string             `json:"name"`
	Active    bool               `json:"active"`
	Service   string             `json:"service"`
	Source    string             `json:"source"`
	Context   string             `json:"context"`
	DependsOn []*JobDependencies `json:"dependsOn,omitempty"`
	JobParams string             `json:"jobParams"`
}

type ConfigInput struct {
	Name      string                  `json:"name"`
	Active    bool                    `json:"active"`
	Service   string                  `json:"service"`
	Source    string                  `json:"source"`
	Context   string                  `json:"context"`
	DependsOn []*JobDependenciesInput `json:"dependsOn,omitempty"`
	JobParams string                  `json:"jobParams"`
}

type JobDependencies struct {
	Service string `json:"service"`
	Source  string `json:"source"`
}

type JobDependenciesInput struct {
	Service string `json:"service"`
	Source  string `json:"source"`
}
package config

const (
	HttpMethodGet    = "GET"
	HttpMethodPost   = "POST"
	HttpMethodUpdate = "UPDATE"
	HttpMethodPut    = "PUT"
	HttpMethodDelete = "DELETE"
)

var HttpMethodMapping = map[string]string{
	HttpMethodGet:    "GetMapping",
	HttpMethodPost:   "PostMapping",
	HttpMethodUpdate: "PutMapping",
	HttpMethodPut:    "PutMapping",
	HttpMethodDelete: "DeleteMapping",
}

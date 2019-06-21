package config

var (
	httpMethodGet    = "GET"
	httpMethodPost   = "POST"
	httpMethodUpdate = "UPDATE"
	httpMethodPut    = "PUT"
	httpMethodDelete = "DELETE"
)

var HttpMethodMapping = map[string]string{
	"GET":    "GetMapping",
	"POST":   "PostMapping",
	"UPDATE": "PostMapping",
	"PUT":    "PutMapping",
	"DELETE": "DeleteMapping",
}

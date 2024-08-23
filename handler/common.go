package handler

type message struct {
	InvalidJson string
	InvalidUrl  string
	InvalidFile string
}

var Message = message{
	InvalidJson: "Invalid JSON payload received.",
	InvalidUrl:  "Invalid URL parameter received.",
	InvalidFile: "Invalid file received.",
}

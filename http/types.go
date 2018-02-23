package loanhttpserver


type Err struct {
	Error	bool        `json:"error"`
	Message	error        `json:"message"`
}

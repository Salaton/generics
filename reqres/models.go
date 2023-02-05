package reqres

type RequestObject struct {
	Page       int     `json:"page"`
	PerPage    int     `json:"per_page"`
	Total      int     `json:"total"`
	TotalPages int     `json:"total_pages"`
	Data       []Data  `json:"data"`
	Support    Support `json:"support"`
}

type Data struct {
	ID        int    `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Avatar    string `json:"avatar"`
}

type Support struct {
	URL  string `json:"url"`
	Text string `json:"text"`
}

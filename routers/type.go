package routers

type Article1 struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Article2 struct {
	Title   string `json:"title" xml:"title"`
	Content string `json:"content" xml:"content"`
}

type UserInfo struct {
	Username string `json:"user" form:"username"`
	Password string `json:"pw" form:"password"`
}

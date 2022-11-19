package user

type User struct {
	Name    string `json:"name"`
	Age     string `json:"age"`
	Friends []int  `json:"friends"`
}

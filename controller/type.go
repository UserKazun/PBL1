package controller

// LoginUser ...ログインユーザ情報
type LoginUser struct {
	ID      string `json:"user_id"`
	Name    string `json:"user_name"`
	IsAdmin bool   `json:"is_admin"`
}

// Recipe ...レシピ
type Recipe struct {
	URL string `json:"recipe_url"`
}

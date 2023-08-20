package domain

// @see https://ieyasu.co/docs/api.html#section/Errors
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

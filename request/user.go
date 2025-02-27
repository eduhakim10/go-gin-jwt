package request

type UserRequest struct {
	Nama     string `binding:"required" json:"nama"`
	Email    string `binding:"required" json:"email"`
	Password string `binding:"required" json:"password"`
}

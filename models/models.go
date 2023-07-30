// models/models.go

package models

// User representa os dados de um usuário
type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"` // Papel/Função do usuário: "admin", "user" ou "company"
}

// Vehicle representa os dados de um veículo
type Vehicle struct {
	ID       string `json:"id"`
	Brand    string `json:"brand"`
	Model    string `json:"model"`
	OwnerID  string `json:"ownerId"` // ID do usuário proprietário do veículo
}

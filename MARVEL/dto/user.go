package dto

type Contacto struct {
	IDusuarios int    `json:"idusuarios"`
	Nombre     string `json:"nombre"`
	Usuario    string `json:"usuario"`
	Contraseña string `json:"contraseña"`
}

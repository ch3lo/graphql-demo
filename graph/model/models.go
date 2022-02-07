package model

type Caso struct {
	ID              string `json:"id"`
	OrdenId         string `json:"-"`
	RepresentanteId string `json:"-"`
	FechaCreacion   string `json:"fechaCreacion"`
}

type NuevoCaso struct {
	OrdenID string `json:"ordenId"`
	RepID   string `json:"repId"`
}

type Orden struct {
	ID           string   `json:"id"`
	ProductosIds []string `json:"-"`
}

type Producto struct {
	ID       string `json:"id"`
	Cantidad int    `json:"cantidad"`
	Nuevo    bool   `json:"nuevo"`
}

type Representante struct {
	ID     string `json:"id"`
	Nombre string `json:"nombre"`
}

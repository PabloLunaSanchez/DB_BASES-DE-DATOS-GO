package product

import (
	"github.com/devjaime/golangrest/database"
	"github.com/gofiber/fiber"
	_ "github.com/jinzhu/gorm/dialects/mssql"
)

type Platillos struct {
	ID          uint   `gorm:"column:idPlatillo;primary_key"`
	Nombre      string `json:"nombre"`
	Descripcion string `json:"descripcion"`
	Precio      string `json:"precio"`
}

func GetProducts(c *fiber.Ctx) {
	db := database.DBConn
	var platillos []Platillos
	db.Find(&platillos)
	c.JSON(platillos)
}

func GetProduct(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var platillos Platillos
	db.Find(&platillos, "idPlatillo = ?", id)
	c.JSON(platillos)
}
func NewProduct(c *fiber.Ctx) {
	db := database.DBConn
	var platillos Platillos
	platillos.Nombre = "Carne arrachera"
	platillos.Descripcion = "con arroz y frijoles y guacamole"
	platillos.Precio = "70.00"
	db.Create(&platillos)
	c.JSON(platillos)
}
func UpdateProduct(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var platillo Platillos
	db.First(&platillo, id)

	if platillo.ID == 0 {
		c.Status(fiber.StatusNotFound).Send("No se encontró el platillo")
		return
	}

	// Obtener los nuevos valores del cuerpo de la solicitud
	var newPlatillo Platillos
	if err := c.BodyParser(&newPlatillo); err != nil {
		c.Status(fiber.StatusBadRequest).Send("Error al analizar la solicitud")
		return
	}

	// Actualizar los valores del platillo
	db.Model(&platillo).Updates(&newPlatillo)

	c.JSON(platillo)
}

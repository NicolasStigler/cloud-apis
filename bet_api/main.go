package main

import (
	"log"
  "math/rand"
  "github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	_ "bet_api/docs" // Import for Swagger documentation
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/files"
)

var db *gorm.DB
var err error

// Juego model
type Juego struct {
  ID     uint   `json:"id" gorm:"primaryKey"`
	Nombre string `json:"nombre"`
}

// Apuesta model
type Apuesta struct {
  ID        uint    `json:"id" gorm:"primaryKey"`
	UsuarioID uint    `json:"usuario_id"`
	JuegoID   uint    `json:"juego_id"`
  Juego     Juego `json:"juego"`
	Monto     float64 `json:"monto"`
	Resultado string  `json:"resultado"`
	Fecha     string  `json:"fecha"`
}

var probabilidades = map[string]float64{
	"slots":     0.4,  // 40% de probabilidad de ganar en Slots
	"roulette":  0.47, // 47% de probabilidad de ganar en Ruleta
	"blackjack": 0.49, // 49% de probabilidad de ganar en Blackjack
}

// Conectar a la base de datos
func initDB() {
	dsn := "user:password@tcp(localhost:3306)/casino?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error al conectar con la base de datos:", err)
	}
	// Migrar modelos
	db.AutoMigrate(&Juego{}, &Apuesta{})
}

func main() {
	r := gin.Default()

	initDB() // Inicializar base de datos

	// Endpoints de juegos
	r.GET("/juegos", getJuegos)
	r.POST("/juegos", createJuego)
	r.DELETE("/juegos/:id", deleteJuego)

	// Endpoints de apuestas
	r.GET("/apuestas", getApuestas)
	r.POST("/apuestas", createApuesta)

	// Swagger UI
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080") // Inicia el servidor en el puerto 8080
}

// Determina si se gana o se pierde basado en probabilidades
func determinarResultado(juego string) string {
	probabilidad := probabilidades[juego]
	// Genera un número aleatorio entre 0 y 1
	if rand.Float64() < probabilidad {
		return "ganado"
	}
	return "perdido"
}

// Obtener todos los juegos
// @Summary Listar juegos
// @Description Obtiene todos los juegos disponibles en el casino
// @Produce json
// @Success 200 {array} Juego
// @Router /juegos [get]
func getJuegos(c *gin.Context) {
	var juegos []Juego
	db.Find(&juegos)
	c.JSON(200, juegos)
}

// Crear un nuevo juego
// @Summary Crear un juego
// @Description Crea un nuevo juego en el casino
// @Accept  json
// @Produce  json
// @Param   juego  body  Juego  true  "Datos del juego"
// @Success 201 {object} Juego
// @Failure 400 {object} map[string]string "Error al procesar la solicitud"
// @Router /juegos [post]
func createJuego(c *gin.Context) {
	var juego Juego
	if err := c.ShouldBindJSON(&juego); err == nil {
		db.Create(&juego)
		c.JSON(201, juego)
	} else {
		c.JSON(400, gin.H{"error": err.Error()})
	}
}

// Eliminar un juego
// @Summary Eliminar un juego
// @Description Elimina un juego por su ID
// @Param   id   path  int  true  "ID del juego a eliminar"
// @Success 204  "El juego ha sido eliminado"
// @Failure 404 {object} map[string]string "Juego no encontrado"
// @Router /juegos/{id} [delete]
func deleteJuego(c *gin.Context) {
	id := c.Param("id")
	db.Delete(&Juego{}, id)
	c.Status(204)
}

// Obtener todas las apuestas
// @Summary Listar apuestas
// @Description Obtiene todas las apuestas realizadas en el casino
// @Produce json
// @Success 200 {array} Apuesta
// @Router /apuestas [get]
func getApuestas(c *gin.Context) {
	var apuestas []Apuesta
	db.Find(&apuestas)
	c.JSON(200, apuestas)
}

// Crear una nueva apuesta
// @Summary Crear una apuesta
// @Description Registra una nueva apuesta en el sistema
// @Accept  json
// @Produce  json
// @Param   apuesta  body  Apuesta  true  "Datos de la apuesta a registrar"
// @Success 201 {object} Apuesta
// @Failure 400 {object} map[string]string "Error al procesar la solicitud"
// @Router /apuestas [post]
func createApuesta(c *gin.Context) {
		var apuesta Apuesta
	if err := c.ShouldBindJSON(&apuesta); err == nil {
		// Buscar el juego
		var juego Juego
		db.First(&juego, apuesta.JuegoID)
		if juego.ID == 0 {
			c.JSON(400, gin.H{"error": "Juego no encontrado"})
			return
		}

		// Determinar resultado
		apuesta.Resultado = determinarResultado(juego.Nombre)
		db.Create(&apuesta)
		c.JSON(201, apuesta)
	} else {
		c.JSON(400, gin.H{"error": err.Error()})
	}
}


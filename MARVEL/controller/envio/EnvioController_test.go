package envio



import (
	"github.com/stretchr/testify/assert"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock"
	"github.com/damian2205/GO_Capacitation/MARVEL/models"
	"github.com/damian2205/GO_Capacitation/MARVEL/dto"
	// mocks "github.com/damian2205/GO_Capacitation/MARVEL/services/mocks"
	"github.com/gin-gonic/gin"
)


func TestObtenerUser(t *testing.T){

	var contactoDto dto.Contacto
	var modelContacto models.ModelContacto

	contactoDto = mockContacto()
	// mockEnvioServices := new(mocks.EnvioServices)

	// mockEnvioServices.On("InsertarUser", contactoDto).Return(contactoDto, error).Once()

	// controller := NewEnvioController(mockEnvioServices)
	b, _ := json.Marshal(contactoDto)

	//Crear router
	gin.SetMode(gin.TestMode)
	res := httptest.NewRecorder()
	ctx, router := gin.CreateTestContext(res)

	req, _ := http.NewRequest(http.MethodPost, "/datos", bytes.NewReader(b))
	ctx.NewRequest = req
	router.POST("/datos", func(c *gin.Context)) //{controller.InsertarUser(ctx)})

	router.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)
)

}

func mockContacto() dto.Contacto{
	return dto.Contacto{
	IDusuarios	"23",
	Nombre	"Jose de las casas",
	Usuario	"Josesito",
	Contraseña	"123456789",
	}
	
}
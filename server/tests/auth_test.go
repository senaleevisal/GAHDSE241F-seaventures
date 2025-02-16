package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"seaventures/src/config"
	"seaventures/src/controller"
	"seaventures/src/models"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/users", controller.RegisterUser)
	r.POST("/users/login", controller.LoginUser)
    r.POST("/auth/protected", controller.ProtectedEndpoint)
	return r
}

func setupDatabase() {

	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_USER", "root")
	os.Setenv("DB_PASSWORD", "Nethalk@123")
	os.Setenv("DB_NAME", "testdb")
	os.Setenv("DB_PORT", "3306")


	config.InitDB()


	config.DB.AutoMigrate(&models.User{})
}

func teardownDatabase() {

	config.DB.Migrator().DropTable(&models.User{})
}

func TestMain(m *testing.M) {


	setupDatabase()

	code := m.Run()

	teardownDatabase()

	os.Exit(code)
}

func TestRegisterUser(t *testing.T) {
	router := setupRouter()

	user := map[string]string{
		"userName": "testuser",
		"email":    "testuser@example.com",
		"password": "password123",
		"role":     "user",
	}
	userJSON, _ := json.Marshal(user)

	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(userJSON))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.NotNil(t, response["token"])
	assert.Equal(t, "User created successfully", response["message"])
}

func TestLoginUser(t *testing.T) {
	router := setupRouter()

	user := map[string]string{
		"userName": "testuser",
		"email":    "testuser@example.com",
		"password": "password123",
		"role":     "user",
	}
	userJSON, _ := json.Marshal(user)

	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(userJSON))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)

	// Now, attempt to log in with the same user credentials
	loginPayload := map[string]string{
		"email":    "testuser@example.com",
		"password": "password123",
	}
	loginJSON, _ := json.Marshal(loginPayload)

	loginReq, _ := http.NewRequest("POST", "/users/login", bytes.NewBuffer(loginJSON))
	loginReq.Header.Set("Content-Type", "application/json")

	loginW := httptest.NewRecorder()
	router.ServeHTTP(loginW, loginReq)

	assert.Equal(t, http.StatusOK, loginW.Code)
	var loginResponse map[string]interface{}
	json.Unmarshal(loginW.Body.Bytes(), &loginResponse)

	assert.NotNil(t, loginResponse["token"])
	assert.Equal(t, "Login successful", loginResponse["message"])
}

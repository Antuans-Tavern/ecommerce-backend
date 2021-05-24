package graph

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/Antuans-Tavern/ecommerce-backend/pkg/database"
	"github.com/Antuans-Tavern/ecommerce-backend/pkg/database/model"
	"github.com/Antuans-Tavern/ecommerce-backend/pkg/util"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestMain(m *testing.M) {
	util.PrepareDatabase()

	code := m.Run()
	os.Exit(code)
}

func TestLoginQuery(t *testing.T) {
	db, err := gorm.Open(postgres.Open(database.ConnectionString()), &gorm.Config{})
	assert.Nil(t, err)

	db.Create(&model.User{
		Email:    "admin@testing.com",
		Password: util.Hash("secret"),
	})

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/query", strings.NewReader(`{"query":"{\n\tlogin(email:\"admin@testing.com\", password:\"secret\"){\n  \ttoken,\n    user{\n      id,\n      email,\n      profile {\n        name\n      }\n    }\n  }\n}\n\n"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	h := QueryHandler(db)

	if assert.NoError(t, h(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		response := map[string]interface{}{}
		json.Unmarshal([]byte(rec.Body.String()), &response)

		assert.NotEmpty(t, response["token"])
		assert.NotEmpty(t, response["name"])
	}
}

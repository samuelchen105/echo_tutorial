package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetForm(t *testing.T) {
	assert := assert.New(t)
	e := echo.New()
	var err error
	//test data
	name := "Jack"
	expected := map[string]interface{}{
		"data": map[string]interface{}{"name": name},
	}
	//setup form
	form := &url.Values{}
	form.Set("name", name)
	//setup request
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(form.Encode()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	//setup context
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	//assert
	assert.NoError(getForm(ctx))
	assert.Equal(http.StatusOK, rec.Code)
	var reply map[string]interface{}
	err = json.Unmarshal(rec.Body.Bytes(), &reply)
	assert.NoError(err)
	assert.Equal(expected, reply)
}

func TestGetQueryParam(t *testing.T) {
	assert := assert.New(t)
	e := echo.New()
	var err error
	//test data
	name := "Jack"
	expected := map[string]interface{}{
		"data": map[string]interface{}{"name": name},
	}
	//setup query
	query := &url.Values{}
	query.Set("name", name)
	//setup request
	req := httptest.NewRequest(http.MethodGet, "/?"+query.Encode(), nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	//setup context
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	//assert
	assert.NoError(getQueryParam(ctx))
	assert.Equal(http.StatusOK, rec.Code)
	var reply map[string]interface{}
	err = json.Unmarshal(rec.Body.Bytes(), &reply)
	assert.NoError(err)
	assert.Equal(expected, reply)
}

func TestGetPathParam(t *testing.T) {
	assert := assert.New(t)
	e := echo.New()
	var err error
	//test data
	name := "Jack"
	expected := map[string]interface{}{
		"data": map[string]interface{}{"name": name},
	}
	//setup request
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	//setup context
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetParamNames("name")
	ctx.SetParamValues(name)
	//assert
	assert.NoError(getPathParam(ctx))
	assert.Equal(http.StatusOK, rec.Code)
	var reply map[string]interface{}
	err = json.Unmarshal(rec.Body.Bytes(), &reply)
	assert.NoError(err)
	assert.Equal(expected, reply)
}

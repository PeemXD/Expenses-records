//go:build integration

package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/PeemXD/expenses-gin/model"
	"github.com/stretchr/testify/assert"
)

type Login struct {
	Token string `json:"token"`
}

type Response struct {
	*http.Response
	err error
}

func (r *Response) Decode(v interface{}) error {
	if r.err != nil {
		return r.err
	}

	return json.NewDecoder(r.Body).Decode(v)
}

func loginAuth() string {

	var login Login

	err := request(http.MethodPost, uri("login"), nil).Decode(&login)
	if err != nil {
		log.Fatal("can't login:", err)
	}

	return login.Token
}

func request(method, url string, body io.Reader) *Response {
	req, _ := http.NewRequest(method, url, body)
	// authToken := loginAuth()
	req.Header.Add("Authorization", os.Getenv("AUTH_TOKEN"))
	req.Header.Add("Content-Type", "application/json")
	client := http.Client{}
	res, err := client.Do(req)
	return &Response{res, err}
}

func uri(path ...string) string {
	host := "http://localhost:2565"
	if path == nil {
		return host
	}

	url := append([]string{host}, path...)
	return strings.Join(url, "/")
}

func seedExpense(t *testing.T) model.Expenses {
	var expenses model.Expenses

	body := bytes.NewBufferString(`{
		"title": "strawberry smoothie",
		"amount": 79,
		"note": "night market promotion discount 10 bath", 
		"tags": ["food", "beverage"]
	}`)

	err := request(http.MethodPost, uri("expenses"), body).Decode(&expenses)
	if err != nil {
		t.Fatal("can't create expense:", err)
	}

	return expenses
}

func TestGetsExpenses(t *testing.T) {
	seedExpense(t)
	var expenses []model.Expenses
	res := request(http.MethodGet, uri("expenses"), nil)
	err := res.Decode(&expenses)

	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusOK, res.StatusCode)
	assert.Greater(t, len(expenses), 0)
}

func TestGetExpense(t *testing.T) {
	existedExpense := seedExpense(t)
	var expense model.Expenses
	res := request(http.MethodGet, uri("expenses", strconv.Itoa(int(existedExpense.ID))), nil)
	err := res.Decode(&expense)

	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusOK, res.StatusCode)
	assert.Equal(t, existedExpense.ID, expense.ID)
	assert.EqualValues(t, existedExpense.Title, expense.Title)
	assert.EqualValues(t, existedExpense.Amount, expense.Amount)
	assert.EqualValues(t, existedExpense.Note, expense.Note)
	assert.EqualValues(t, existedExpense.Tags, expense.Tags)
}

func TestCreateExpense(t *testing.T) {
	var expense model.Expenses
	var body model.Expenses
	body.Title = "strawberry smoothie"
	body.Amount = 79
	body.Note = "night market promotion discount 10 bath"
	body.Tags = []string{"food", "beverage"}
	jsonBody, err := json.Marshal(&body)
	if err != nil {
		log.Fatal("cannot marshal")
		return
	}

	res := request(http.MethodPost, uri("expenses"), bytes.NewBuffer(jsonBody)) // post must return what we add
	err = res.Decode(&expense)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, res.StatusCode)
	assert.Equal(t, body.Title, expense.Title)
	assert.Equal(t, body.Amount, expense.Amount)
	assert.Equal(t, body.Note, expense.Note)
	assert.Equal(t, body.Tags, expense.Tags)
}

func TestLogin(t *testing.T) {
	var login Login

	res := request(http.MethodPost, uri("login"), nil)
	err := res.Decode(&login)

	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusOK, res.StatusCode)

}

func TestEditExpense(t *testing.T) {
	var expense model.Expenses
	param := 1
	body := model.Expenses{
		Title:  "apple smoothie",
		Amount: 89,
		Note:   "no discount",
		Tags:   []string{"beverage"}}
	jsonBody, err := json.Marshal(&body)
	if err != nil {
		log.Fatal("cannot marshal")
		return
	}

	res := request(http.MethodPut, uri("expenses", strconv.Itoa(param)), bytes.NewBuffer(jsonBody)) // post must return what we add
	err = res.Decode(&expense)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.Equal(t, uint(param), expense.ID)
	assert.Equal(t, body.Title, expense.Title)
	assert.Equal(t, body.Amount, expense.Amount)
	assert.Equal(t, body.Note, expense.Note)
	assert.Equal(t, body.Tags, expense.Tags)
}

package main

/*
   ____    __
  / __/___/ /  ___
 / _// __/ _ \/ _ \
/___/\__/_//_/\___/ v4.0.0
High performance, minimalist Go web framework
https://echo.labstack.com
____________________________________O/_______
                                    O\
  */

import (
    "log"
    "net/http"
    "github.com/ant0ine/go-json-rest/rest"
)

type postHelloInput struct {
    Name string
}

type postHelloOutput struct {
    Result string
}

/*
type Error struct {
	Status     int64  `json:"status"`
	Type       string `json:"type"`
	Code   string `json:"code,omitempty"`
	Errors []struct {
		Code       string `json:"code,omitempty"`
		Detail     string `json:"detail,omitempty"`
		Field      string `json:"field,omitempty"`
		UserDetail string `json:"user_detail,omitempty"`
	} `json:"errors,omitempty"`
	UserDetail string `json:"user_detail,omitempty"`
}
 */

type DbConf struct {
    Dialect  string
    User     string `env:"MysqlUser"`
    Password string `env:"MysqlPassword"`
    Address  string `env:"MysqlAddress"`
    Port     int    `env:"MysqlPort"`
    Dbname   string `env:"MysqlDbname"`
}

/* 
    dialect:  mysql
    host:     localhost
    port:     3306
    user:     root
    password:
    dbname:   golang_api_sample
 */

func IsTemporary(err error) bool {
        te, ok := err.(temporary)
        return ok && te.Temporary()
}

func (e *MyError) Error() string { 
        return fmt.Sprintf("%s:%d: %s”, e.File, e.Line, e.Msg)
}

func postHello(w rest.ResponseWriter, req *rest.Request) {
    input := postHelloInput{}
    err := req.DecodeJsonPayload(&input)

    if err != nil {
        rest.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    if input.Name == "" {
        rest.Error(w, "Name is required", 400)
    }

    log.Printf("%#v", input)

    w.WriteJson(&postHelloOutput{
        "Hello, " + input.Name,
    })
}

/* 
type mock struct{}
func (m *mock) GetById(id int64) entity.User {
    return testData
}
 */

func hello(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"hello": "world"})
}

func main() {
    api := rest.NewApi()
    api.Use(rest.DefaultDevStack...)
    router, err := rest.MakeRouter(
        rest.Post("/hello", postHello),
    )

    if err != nil {
        log.Fatal(err)
    }
    log.Printf("Server started.")
    api.SetApp(router)
    log.Fatal(http.ListenAndServe(":9999", api.MakeHandler()))
}

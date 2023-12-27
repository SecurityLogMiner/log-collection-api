package main

import (
    "github.com/joho/godotenv"
    "fmt"
    "strings"
    "net/http"
    "io/ioutil"
    "encoding/json"
    "os"
)

type Auth0Response struct {
    AccessToken string `json:"access_token"`
    Expires int`json:"expires_in"` 
    Type string `json:"token_type"` 
}

func testRoute(url string, token string) bool {

    req, _ := http.NewRequest("GET", url, nil)
    tokenstring := fmt.Sprintf("Bearer %s", token)
    req.Header.Add("authorization", tokenstring)
    
    res, err := http.DefaultClient.Do(req)
    if err != nil {
        fmt.Println("error executing default client: %v",err)
        return false
    }

    defer res.Body.Close()
    body, _ := ioutil.ReadAll(res.Body)

    fmt.Println(res)
    fmt.Println(string(body))
    return true
}

func main() {
    godotenv.Load()
    url := "https://log-collection.us.auth0.com/oauth/token"

    sentence := fmt.Sprintf("{\"client_id\": \"%s\", \"client_secret\": \"%s\",\"audience\": \"%s\", \"grant_type\":\"client_credentials\"}", os.Getenv("AUTH0_CLIENTID"), 
                                    os.Getenv("AUTH0_SECRET"), 
                                    os.Getenv("AUTH0_AUDIENCE"))

    payload := strings.NewReader(sentence)

    req, _ := http.NewRequest("POST", url, payload)

    req.Header.Add("content-type", "application/json")

    res, _ := http.DefaultClient.Do(req)

    defer res.Body.Close()
    body, _ := ioutil.ReadAll(res.Body)

    var auth_response Auth0Response
    err := json.Unmarshal(body, &auth_response)
    if err != nil {
        fmt.Println("error reading auth0 response")
        return
    }
    
    // TESTS
    testRoute("http://localhost:6060/api/private",auth_response.AccessToken)
}

package test10

import (
    "net/http"
    "net/url"

    "github.com/go-cas/cas"
    "github.com/julienschmidt/httprouter"
    "golang.org/x/net/context"
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
    if !cas.IsAuthenticated(r) {
        cas.RedirectToLogin(w, r)
    }

    _ := r.Context().Value("params").(httprouter.Params)

    // business logic
}


func main() {

    u, _ := url.Parse("https://cas_example_server.com")

    client := cas.NewClient(&cas.Options{
        URL: u,
    })

    router := httprouter.New()

    //This line fails with the message:

    //"Cannot use defaultHandler (type func(http.ResponseWriter, *http.Request, httprouter.Params))
    //as type func(http.ResponseWriter, *http.Request) in argument to client.HandleFunc"

    router.Handler("GET", "/", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
        newContext := context.WithValue(r.Context(), "params", ps)
        r.WithContext(newContext)
        client.HandleFunc(defaultHandler)(w, r)
    })

    err := http.ListenAndServe(":8080", router)
    if err != nil {
        panic(err)
    }

}
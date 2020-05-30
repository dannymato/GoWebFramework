package framework

import (
	"net/http"

	"github.com/justinas/alice"
)

/*func (c *appContext) authHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		authToken := r.Header.Get("Authorization")
		user, err := getUser(c.db, authToken)

		if err != nil {
			http.Error(w, http.StatusText(401), 401)
			return
		}

		context.Set(r, "user", user)
		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}*/

// WebListener contains the main components for the webservice framework
type WebListener struct {
	middleware alice.Chain
	router     *HandlerRouter
}

// NewWebListener returns a pointer to a new WebListener with the specified database
func NewWebListener() *WebListener {
	listener := new(WebListener)
	listener.middleware = alice.New()
	listener.router = NewRouter()
	return listener
}

// Start initializes the http web service on the specified host address with the routes provided previously
func (listener *WebListener) Start(hostAddress string) {
	http.ListenAndServe(hostAddress, listener.router.GetRouter())
}

// AddMiddleware adds an additional function to the alice Chain for data to pass through on the way to the handler
func (listener *WebListener) AddMiddleware(handler alice.Constructor) {
	listener.middleware = listener.middleware.Append(handler)
}

// GET wraps the router GET function and calls the middleware before the passed HandlerFunc
func (listener *WebListener) GET(path string, handler http.HandlerFunc) {
	listener.router.GET(path, listener.middleware.ThenFunc(handler))
}

// POST wraps the router POST function and calls the middleware before the passed HandlerFunc
func (listener *WebListener) POST(path string, handler http.HandlerFunc) {
	listener.router.POST(path, listener.middleware.ThenFunc(handler))
}

// PUT wraps the router PUT function and calls the middleware before the passed HandlerFunc
func (listener *WebListener) PUT(path string, handler http.HandlerFunc) {
	listener.router.PUT(path, listener.middleware.ThenFunc(handler))
}

// OPTIONS wraps the router OPTIONS function and calls the middleware before the passed HandlerFunc
func (listener *WebListener) OPTIONS(path string, handler http.HandlerFunc) {
	listener.router.OPTIONS(path, listener.middleware.ThenFunc(handler))
}

// DELETE wraps the router DELETE function and calls the middleware before the passed HandlerFunc
func (listener *WebListener) DELETE(path string, handler http.HandlerFunc) {
	listener.router.DELETE(path, listener.middleware.ThenFunc(handler))
}

func (listener *WebListener) SetGlobalOPTIONS(handler http.HandlerFunc) {
	listener.router.router.GlobalOPTIONS = handler
	listener.router.router.HandleOPTIONS = true
}

/*
func main() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	appC := appContext{session.DB("test")}
	commonHandlers := alice.New(context.ClearHandler, loggingHandler, recoverHandler, acceptHandler, corsHandler)
	router := NewRouter()
	router.GET("/teas/:id", commonHandlers.ThenFunc(appC.teaHandler))
	router.POST("/teas", commonHandlers.Append(bodyParserHandler(TeaResource{})).ThenFunc(appC.createTeaHandler))
	router.OPTIONS("/teas", alice.New(context.ClearHandler, loggingHandler, recoverHandler, corsHandler).ThenFunc(optionsHandler))
	router.GET("/teasAll", commonHandlers.ThenFunc(appC.allTeasHandler))
	fmt.Println("Starting WebServer")
	http.ListenAndServe(":8080", router.GetRouter())
}*/

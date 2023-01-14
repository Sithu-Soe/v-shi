package handler

import (
	"github.com/gin-gonic/gin"
)

type graphqlHandler struct {
	R *gin.Engine
}

func newGraphqlHandler(h *Handler) *graphqlHandler {
	return &graphqlHandler{
		R: h.R,
	}
}

func (ctr *graphqlHandler) register() {
	ctr.R.POST("/graphql", ctr.serveGraphQL)
}

// Defining the Graphql handler
func (ctr *graphqlHandler) serveGraphQL(c *gin.Context) {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	// h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	// return func(c *gin.Context) {
	// 	h.ServeHTTP(c.Writer, c.Request)
	// }
}

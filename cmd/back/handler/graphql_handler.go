package handler

import (
	"v-shi/cmd/back/graph"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
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

	group := ctr.R.Group("/graphql")
	group.POST("", ctr.serveGraphQL)
	group.GET("/play", ctr.playGraphQL)
}

// Defining the Graphql handler
func (ctr *graphqlHandler) serveGraphQL(c *gin.Context) {
	handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}})).ServeHTTP(c.Writer, c.Request)
}

func (ctr *graphqlHandler) playGraphQL(c *gin.Context) {
	playground.Handler("GraphQL", "/query").ServeHTTP(c.Writer, c.Request)
}

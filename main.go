package main

import (
	"encoding/json"
	"fmt"
	"github.com/graphql-go/graphql"
	"go-graphql-tutorial/pkg/model"
	"log"
)

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}
	return result
}

var aggregateSchema = graphql.Fields{
	//"tutorial": model.SingleTutorialSchema(),
	"list": model.ListTutorialSchema(),
}

//var createMutations = graphql.NewObject(graphql.ObjectConfig{
//	Name: "Mutation",
//	Fields: graphql.Fields{
//		"create": model.CreateTutorialMutation(),
//	},
//})

var queryMutations = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"query": model.QueryTutorialMutationByID(),
	},
})

func main() {
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: aggregateSchema}
	schema, err := graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    graphql.NewObject(rootQuery),
			Mutation: queryMutations,
		},
	)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	//create := `
	//	mutation {
	//		create(id: 6,title: "Sweet") {
	//			id
	//			title
	//			updated_at
	//			created_at
	//		}
	//	}
	//`
	query := `
		mutation {
			query(id: 15,title: "Sweet") {
				id
				title
				updated_at
				created_at
			}
		}
	`
	result := executeQuery(query, schema)
	rJSON, _ := json.Marshal(result)
	fmt.Printf("%s \n", rJSON)

	//query := `
	//	{
	//		list {id title
	//		}
	//	}
	//`
	//result := executeQuery(query, schema)
	//rJSON, _ := json.Marshal(result)
	//fmt.Printf("%s \n", rJSON)
}

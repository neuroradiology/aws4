package dydb_test

import (
	"fmt"
	"github.com/bmizerany/aws4/dydb"
	"log"
)

func init() {
	log.SetFlags(0)
}

func Example_listTables() {
	db := new(dydb.DB)

	type AttributeDefinition struct {
		AttributeName string
		AttributeType string
	}

	type KeySchema struct {
		AttributeName string
		KeyType       string
	}

	type CreateTable struct {
		TableName             string
		AttributeDefinitions  []AttributeDefinition
		KeySchema             []KeySchema
		ProvisionedThroughput struct {
			ReadCapacityUnits  int
			WriteCapacityUnits int
		}
	}

	posts := new(CreateTable)
	posts.TableName = "Posts"
	posts.AttributeDefinitions = []AttributeDefinition{{"Slug", "S"}}
	posts.KeySchema = []KeySchema{{"Slug", "HASH"}}
	posts.ProvisionedThroughput.ReadCapacityUnits = 4
	posts.ProvisionedThroughput.WriteCapacityUnits = 4

	if err := db.Exec("CreateTable", posts); err != nil {
		if e, ok := err.(*dydb.ResponseError); ok {
			if e.TypeName() != "ResourceInUseException" {
				log.Fatal(err)
			}
		}
	}

	var resp struct{ TableNames []string }
	if err := db.Query("ListTables", nil).Decode(&resp); err != nil {
		log.Fatal(err)
	}

	// Output:
	// ["Posts"]
	fmt.Printf("%q", resp.TableNames)
}

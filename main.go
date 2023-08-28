package main

import (
	"bytes"
	"log"
	"os"
)

var nodePt = `"""Fetches an object given its ID."""
  node(
    """ID of the object."""
    id: ID!
  ): Node`

var nodesPt = `"""Lookup nodes by a list of IDs."""
  nodes(
    """The list of node IDs."""
    ids: [ID!]!
  ): [Node]!`

func main() {
	path, _ := os.Getwd()
	schemaFiles := []string{"/graph/ent.graphql"}

	for _, file := range schemaFiles {
		body, err := os.ReadFile(path + file)
		if err != nil {
			log.Fatalf("failed to run fixer: %s", err)
		}

		body = bytes.Replace(body, []byte(nodePt), []byte(""), -1)
		body = bytes.Replace(body, []byte(nodesPt), []byte(""), -1)

		rewrite, _ := os.OpenFile(path+file, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
		_, err = rewrite.Write(body)
		if err != nil {
			log.Fatalf("failed to run fixer: %s", err)
		}
		rewrite.Close()
	}

	log.Printf("fixer complite in: %s", path)
}

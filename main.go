package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/blevesearch/bleve"
)

func main() {
	indexDirPtr := flag.String("i", "asma.bleve", "A full path to bleve index")
	queryStringPtr := flag.String("q", "", "The search phrase")
	flag.Parse()

	index, ierr := bleve.Open(*indexDirPtr)
	if ierr != nil {
		log.Fatalln("Unable to read index")
	}
	query := bleve.NewMatchQuery(*queryStringPtr)
	search := bleve.NewSearchRequest(query)
	searchResults, searchErr := index.Search(search)
	if searchErr != nil {
		log.Fatalln("Search failed")
	}
	fmt.Println(searchResults)
}

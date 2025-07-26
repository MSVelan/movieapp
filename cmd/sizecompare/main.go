package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"

	"github.com/MSVelan/movieapp/gen"
	"github.com/MSVelan/movieapp/metadata/pkg/model"
	"google.golang.org/protobuf/proto"
)

var metadata = &model.Metadata{
	ID:          "1",
	Title:       "Movie-1",
	Description: "Made by MSVelan.",
	Director:    "Mr. Sivvel",
}

var genMetadata = &gen.Metadata{
	Id:          "1",
	Title:       "Movie-1",
	Description: "Made by MSVelan.",
	Director:    "Mr. Sivvel",
}

func serializeToJSON(m *model.Metadata) ([]byte, error) {
	return json.Marshal(m)
}

func serializeToXML(m *model.Metadata) ([]byte, error) {
	return xml.Marshal(m)
}

func serializeToProto(m *gen.Metadata) ([]byte, error) {
	return proto.Marshal(m)
}

func main() {
	jsonBytes, err := serializeToJSON(metadata)
	if err != nil {
		panic(err)
	}

	xmlBytes, err := serializeToXML(metadata)
	if err != nil {
		panic(err)
	}

	protoBytes, err := serializeToProto(genMetadata)
	if err != nil {
		panic(err)
	}

	fmt.Printf("JSON size: \t%dB\n", len(jsonBytes))
	fmt.Printf("XML size: \t%dB\n", len(xmlBytes))
	fmt.Printf("Proto size: \t%dB\n", len(protoBytes))
}

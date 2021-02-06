package main

import (
	"encoding/json"
	"encoding/xml"
	"encoding/gob"
	"fmt"
	"log"
	"os"
	"time"
	"gopkg.in/yaml.v2"
	"github.com/vmihailenco/msgpack/v5"
	data "github.com/kssenii/soa-curriculum-2021/homeworks/serialization/data"
	"github.com/golang/protobuf/proto"
)

const (
	schemaKey = "avro.schema"
)

type (
	stringMap map[string]interface{}
	Data struct {
		A bool        `json:"a"`
		B uint64      `json:"b"`
		C float64     `json:"c"`
		D string      `json:"d"`
		E []string    `json:"e"`
		F [10]float32 `json:"f"`
		G stringMap   `json:"j"`
	}
)

/// Because xml.Marshal was unable to parse map[string]interface{}
func (s stringMap) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	tokens := []xml.Token{start}
	for k, v := range s {
		v := fmt.Sprintf("%v", v)
		t := xml.StartElement{Name: xml.Name{"", k}}
		tokens = append(tokens, t, xml.CharData(v), xml.EndElement{t.Name})
	}
	tokens = append(tokens, xml.EndElement{start.Name})
	for _, t := range tokens {
		err := e.EncodeToken(t)
		if err != nil {
			return err
		}
	}
	err := e.Flush()
	if err != nil {
		return err
	}

	return nil
}

func main() {
	file_results, err := os.OpenFile("result.txt", os.O_WRONLY | os.O_CREATE | os.O_TRUNC, 0666);
	if err != nil {
		return
	}
	defer file_results.Close()

	file_output, err := os.OpenFile("output.txt", os.O_WRONLY | os.O_CREATE | os.O_TRUNC, 0666);
	if err != nil {
		return
	}
	defer file_output.Close()

	data_1 := &Data{
		A: true,
		B: 18446744073709551614,
		C: 3.145678900,
		D: "Lorem ipsum dolor sit amet, consectetur adipiscingelit." +
			" Mauris adipiscing adipiscing placerat.Vestibulum augue augue,pellentesque" +
			" quis sollicitudin id, adipiscing.",
		E: []string{"Lorem", "ipsum", "dolor", "sit", "amet,", "consectetur", "adipiscingelit",
			"adipiscing", "Mauris", "adipiscing", "placerat.", "Vestibulum", "augue", "augue", "pellentesque",
			"quits", "sollicitudin", "id,", "adipiscing."},
		F: [10]float32{1.12, 2.12, 3.12, 4.12, 5.12, 6.12, 7.12, 8.12, 9.12, 10.12},
		G: stringMap{"name": "Sammy", "animal": "shark", "color": "blue", "location": "ocean"},

	}

	loop := 100000
	data_2 := &Data{};


	/// JSON serialize

	result, err := json.Marshal(data_1);
	if err != nil {
		log.Println(err)
	}
	file_output.WriteString(string(result));

	start := time.Now();
	for i := 0; i < loop; i++ {
		result, _ = json.Marshal(data_1);
	}
	elapsed := time.Since(start);
	file_results.WriteString(fmt.Sprintf("Time to serialize JSON: %s \n", elapsed));


	/// JSON deserialize

	json.Unmarshal(result, data_2);

	start = time.Now();
	for i := 0; i < loop; i++ {
		json.Unmarshal(result, data_2);
	}
	elapsed = time.Since(start);
	file_results.WriteString(fmt.Sprintf("Time to deserialize JSON: %s \n\n", elapsed));


	/// XML serialization

	result, err = xml.MarshalIndent(data_1,"", "  ");
	if err != nil {
		log.Println(err)
	}
	file_output.WriteString(string(result));

	start = time.Now();
	for i := 0; i < loop; i++ {
		result, _ = xml.Marshal(data_1);
	}
	elapsed = time.Since(start);
	file_results.WriteString(fmt.Sprintf("Time to serialize XML: %s \n", elapsed));


	/// XML deserialization

	xml.Unmarshal(result, data_2)

	start = time.Now();
	for i := 0; i < loop; i++ {
		xml.Unmarshal(result, data_2);
	}
	elapsed = time.Since(start);
	file_results.WriteString(fmt.Sprintf("Time to deserialize XML: %s \n\n", elapsed));


	/// YAML serialization

	result, err = yaml.Marshal(data_1)
	if err != nil {
		log.Println(err)
	}
	file_output.WriteString(string(result));

	start = time.Now();
	for i := 0; i < loop; i++ {
		result, _ = yaml.Marshal(data_1);
	}
	elapsed = time.Since(start);
	file_results.WriteString(fmt.Sprintf("Time to serialize YAML: %s \n", elapsed));

	/// YAML deserialization

	yaml.Unmarshal(result, data_2)

	start = time.Now();
	for i := 0; i < loop; i++ {
		yaml.Unmarshal(result, data_2);
	}
	elapsed = time.Since(start);
	file_results.WriteString(fmt.Sprintf("Time to deserialize YAML: %s \n\n", elapsed));


	/// MessagePack serialization

	result, err = msgpack.Marshal(data_1)
	if err != nil {
		log.Println(err)
	}
	file_output.WriteString(string(result));

	start = time.Now();
	for i := 0; i < loop; i++ {
		result, _ = msgpack.Marshal(data_1);
	}
	elapsed = time.Since(start);
	file_results.WriteString(fmt.Sprintf("Time to serialize MessagePack: %s \n", elapsed));


	/// MessagePack deserialization

	msgpack.Unmarshal(result, data_2)

	start = time.Now();
	for i := 0; i < loop; i++ {
		msgpack.Unmarshal(result, data_2);
	}
	elapsed = time.Since(start);
	file_results.WriteString(fmt.Sprintf("Time to deserialize MessagePack: %s \n\n", elapsed));


	/// Native serialization

	e := gob.NewEncoder(file_output)
	err = e.Encode(data_1)
	if err != nil {
		log.Println(err)
	}
	file_output.WriteString(string(result));

	start = time.Now();
	for i := 0; i < loop; i++ {
		err = e.Encode(data_1)
	}
	elapsed = time.Since(start);
	file_results.WriteString(fmt.Sprintf("Time to serialize Native: %s \n", elapsed));


	/// Native deserialization

	d := gob.NewDecoder(file_output)
	err = d.Decode(data_1)

	start = time.Now();
	for i := 0; i < loop; i++ {
		err = d.Decode(data_1)
	}
	elapsed = time.Since(start);
	file_results.WriteString(fmt.Sprintf("Time to deserialize Native: %s \n\n", elapsed));


	/// Protobuf serialization

	message := &data.StructData{
		A: true,
		B: 18446744073709551614,
		C: 3.145678900,
		D: "Lorem ipsum dolor sit amet, consectetur adipiscingelit." +
			" Mauris adipiscing adipiscing placerat.Vestibulum augue augue,pellentesque" +
			" quis sollicitudin id, adipiscing.",
		E: []string{"Lorem", "ipsum", "dolor", "sit", "amet,", "consectetur", "adipiscingelit",
			"adipiscing", "Mauris", "adipiscing", "placerat.", "Vestibulum", "augue", "augue", "pellentesque",
			"quits", "sollicitudin", "id,", "adipiscing."},
		F: []float32{1.12, 2.12, 3.12, 4.12, 5.12, 6.12, 7.12, 8.12, 9.12, 10.12},
		G: map[string]string{"name": "Sammy", "animal": "shark", "color": "blue", "location": "ocean"},
	};

	data_pb, err := proto.Marshal(message)
	if err != nil {
		panic(err)
	}

	start = time.Now();
	for i := 0; i < loop; i++ {
		data_pb, _ = proto.Marshal(message)
	}
	elapsed = time.Since(start);
	file_results.WriteString(fmt.Sprintf("Time to serialize Protobuf: %s \n", elapsed));


	/// Protobuf deserialization

	newMessage := &data.StructData{}
	err = proto.Unmarshal(data_pb, newMessage)
	if err != nil {
		panic(err)
	}

	start = time.Now();
	for i := 0; i < loop; i++ {
		_ = proto.Unmarshal(data_pb, newMessage)
	}
	elapsed = time.Since(start);
	file_results.WriteString(fmt.Sprintf("Time to deserialize Protobuf: %s \n", elapsed));

	log.Println(newMessage.GetA())
	log.Println(newMessage.GetB())
	log.Println(newMessage.GetC())
}


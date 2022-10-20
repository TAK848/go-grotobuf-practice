package main

import (
	"fmt"
	"grpc-ud/pb"
	"log"

	"github.com/golang/protobuf/jsonpb"
)

func main() {
	employee := &pb.Employee{
		Id:          123,
		Name:        "John Doe",
		Email:       "test@example.com",
		Occupation:  pb.Occupation_ENGINEER,
		PhoneNumber: []string{"080-1234-5678", "090-1234-5678"},
		Project:     map[string]*pb.Company_Project{"Project A": &pb.Company_Project{}},
		Profile: &pb.Employee_Text{
			Text: "This is a text profile",
		},
		Birthday: &pb.Date{
			Year:  1990,
			Month: 1,
			Day:   1,
		},
	}
	// binData, err := proto.Marshal(employee)
	// if err != nil {
	// 	log.Fatalln("Failed to encode employee:", err)
	// }
	// if err := ioutil.WriteFile("test.bin", binData, 06666); err != nil {
	// 	log.Fatalln("Failed to write file:", err)
	// }

	// in, err := ioutil.ReadFile("test.bin")
	// if err != nil {
	// 	log.Fatalln("Failed to read file:", err)
	// }
	// readEmployee := &pb.Employee{}
	// err = proto.Unmarshal(in, readEmployee)
	// if err != nil {
	// 	log.Fatalln("Failed to parse address book:", err)
	// }
	// log.Println(readEmployee)
	// fmt.Println(readEmployee)
	m := jsonpb.Marshaler{}
	out, err := m.MarshalToString(employee)
	if err != nil {
		log.Fatalln("Failed to encode employee:", err)
	}
	fmt.Println(out)
	readEmployee := &pb.Employee{}
	if err := jsonpb.UnmarshalString(out, readEmployee); err != nil {
		log.Fatalln("Failed to parse address book:", err)
	}
	fmt.Println(readEmployee)
}

package packet

import (
	"testing"
	"fmt"
)

func Test_Pack(t *testing.T){
	type person struct {
		Name string `json: "name"`
		Sex bool	`json: "sex"`
		Age int8	`json: "age"`
		Grade int16	`json: "grade"`
		Height int32	`json: "height"`
		Weight int64	`json: "weight"`
		Iq	float32	`json: "iq"`
		PhoneNum []string`json: "phoneNum"`
	}

	p := person{Name:"lilei", Sex: true, Age: 100, Grade: 0,
			Height: 100, Weight: 200, Iq: 250.3,
			PhoneNum: []string{"110", "120", "119", "114"}}

	writer := NewPacket()

	Pack(writer, p)

	fmt.Printf("packet data len: %v \ndata: %v\n", len(writer.Data), writer.Data)
	//ret, err := json.Marshal(p)
	//if err != nil {
	//	fmt.Printf("error!!!!!!!!!!!!!!!!!!!!\n")
	//	return
	//} else {
	//	fmt.Printf("packet data len: %v \ndata: %v\n", len(ret), ret)
	//}

	return
}

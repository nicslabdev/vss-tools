package main

import (
	"fmt"
	"log"
	"os"
)

var treeFp *os.File
var fulltree string
var justValidate string
var validateNotEmpty string

func readBytes(numOfBytes uint32) []byte {
	if numOfBytes > 0 {
		buf := make([]byte, numOfBytes)
		treeFp.Read(buf)
		return buf
	}
	return nil
}

func deSerializeUInt(buf []byte) interface{} {
	switch len(buf) {
	case 1:
		var intVal uint8
		intVal = (uint8)(buf[0])
		return intVal
	case 2:
		var intVal uint16
		intVal = (uint16)((uint16)((uint16)(buf[1])*256) + (uint16)(buf[0]))
		return intVal
	case 4:
		var intVal uint32
		intVal = (uint32)((uint32)((uint32)(buf[3])*16777216) + (uint32)((uint32)(buf[2])*65536) + (uint32)((uint32)(buf[1])*256) + (uint32)(buf[0]))
		return intVal
	default:
		fmt.Printf("Buffer length=%d is of an unknown size", len(buf))
		return nil
	}
}

func GetNodes() {
	NameLen := deSerializeUInt(readBytes(1)).(uint8)
	Name := string(readBytes((uint32)(NameLen)))

	NodeTypeLen := deSerializeUInt(readBytes(1)).(uint8)
	NodeType := string(readBytes((uint32)(NodeTypeLen)))

	UuidLen := deSerializeUInt(readBytes(1)).(uint8)
	Uuid := string(readBytes((uint32)(UuidLen)))

	DescrLen := deSerializeUInt(readBytes(2)).(uint16)
	Description := string(readBytes((uint32)(DescrLen)))

	DatatypeLen := deSerializeUInt(readBytes(1)).(uint8)
	Datatype := string(readBytes((uint32)(DatatypeLen)))

	MinLen := deSerializeUInt(readBytes(1)).(uint8)
	Min := string(readBytes((uint32)(MinLen)))

	MaxLen := deSerializeUInt(readBytes(1)).(uint8)
	Max := string(readBytes((uint32)(MaxLen)))

	UnitLen := deSerializeUInt(readBytes(1)).(uint8)
	Unit := string(readBytes((uint32)(UnitLen)))

	allowedStrLen := deSerializeUInt(readBytes(2)).(uint16)
	allowedStr := string(readBytes((uint32)(allowedStrLen)))

	DefaultLen := deSerializeUInt(readBytes(1)).(uint8)
	DefaultAllowed := string(readBytes((uint32)(DefaultLen)))

	ValidateLen := deSerializeUInt(readBytes(1)).(uint8)
	Validate := string(readBytes((uint32)(ValidateLen)))

	Children := int(deSerializeUInt(readBytes(1)).(uint8))

	for childNo := 0; childNo < Children; childNo++ {
		GetNodes()
	}
	fulltree += fmt.Sprintf("Name=%s, NodeType=%s, Uuid=%s, Description=%s, Datatype=%s, Min=%s, Max=%s, Unit=%s, allowedStr=%s, DefaultAllowed=%s, Validate=%s, Children=%d\n", Name, NodeType, Uuid, Description, Datatype, Min, Max, Unit, allowedStr, DefaultAllowed, Validate, Children)
	justValidate += "Name: " + Name + ", Validate: " + Validate + "\n"
	if Validate != "" {
		validateNotEmpty += "Name: " + Name + ", NodeType: " + NodeType + ", Validate: " + Validate + "\n"
	}
}

func main() {
	// Open binary file and read contents
	var err error
	treeFp, err = os.Open("vss_rel_3.1-develop.binary")
	if err != nil {
		log.Fatal(err)
	}
	defer treeFp.Close()

	GetNodes()
	//fmt.Println(fulltree)         // Print full tree
	fmt.Println(justValidate)     // Print just the validate and name field
	fmt.Println(validateNotEmpty) // Print just the validate and name field for nodes that have a validate field
}

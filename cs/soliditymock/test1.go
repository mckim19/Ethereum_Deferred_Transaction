package main

import (
	//"time"
	"fmt"
	"strings"
	"encoding/json"
	proxy "evm"
	statedb "statedb"
)
type (
	Msg struct{
		Flag uint `json:"flag"`
		Words []string `json: "words"`
		Vals []uint `json: "vals"`
	}
)
const (
	ToShuf = iota
	Exit
	ToCS
)
var mockdata []byte
var totalMapperNum = 1

// 1. Solidity Part -> internal
func writeMsg(msg []byte) {
	length := len(msg)
	a := make([]byte, 32)
	msg = append(msg, a...)
	idx := 0
	for {
		tmp := make([]byte, 32)
		if idx>=length {
			tmp[0] = 10
			for i := 1; i<32; i++ {
				tmp[i] = 0
			}
			proxy.Write(tmp)
			break
		}
		tmp = msg[idx:idx+32]
		proxy.Write(tmp)
		idx = idx+32
	}
}
func readMsg() []byte {
	var data []byte
	tmp := make([]byte, 32)
	for {
		tmp = proxy.Read()
		if tmp[0] == 10 && tmp[1] == 0{
			break
		}
		data = append(data, tmp...)
		if tmp[0] == 0 {
			fmt.Println("In readMsg->", data)
			break
		}
	}
	for i:=0; i<len(data); i++ {
		if data[i] == 0 {
			data = data[0:i]
			break
		}
	}
	return data
}

// 2. Solidity Part -> view
func Mapper(input []string){
	// 0. send End Msg
	// 1. tokenizer and map it
	m := make(map[string]uint)
	for _, line := range input {
		for _, token := range strings.Split(line, " ") {
			m[token]++
		}
	}
	// 2. construct message
	msg := Msg {
		Flag: ToShuf,
		Words: make([]string, len(m)),
		Vals: make([]uint, len(m)),
	}
	i := 0
	for key, val := range m {
		msg.Words[i] = key
		msg.Vals[i] = val
		i++
	}

	// 3. encoding the msg from struct to json
	doc, _ := json.Marshal(msg)

	// 4. send the mapping result to the shuffler
	writeMsg(doc)
}
func ShuffleAndReduce(){
	finishedmappernum := 0

	// 1. read data
	var tmp Msg
	m := make(map[string]uint)

	for {
		data:=readMsg()
		json.Unmarshal(data, &tmp)
		if tmp.Flag == Exit && finishedmappernum < totalMapperNum{
			finishedmappernum += 1
		}
		if finishedmappernum == totalMapperNum {
			break
		}
		for j := 0; j<len(tmp.Words); j++ {
			m[tmp.Words[j]]+=tmp.Vals[j] 
		}
	}

	// 2. construct msg
	msg := Msg {
		Flag: ToCS,
		Words: make([]string, len(m)),
		Vals: make([]uint, len(m)),
	}
	i := 0
	for key, val := range m {
		msg.Words[i] = key
		msg.Vals[i] = val
		i++
	}
	
	// 3. encoding the msg from struct to json
	doc, _ := json.Marshal(msg)

	// 4. send data
	fmt.Println(string(doc))
	// writeMsg(doc)
}
func Ending(config 	statedb.Config) {
	msg := Msg {
		Flag: Exit,
		Words: make([]string, 1),
		Vals: make([]uint, 1),
	}
	doc, _ := json.Marshal(msg)
	writeMsg(doc)
}
func Start(config statedb.Config) {
	proxy.Run(config)
}
func Finish(config statedb.Config) {
	proxy.Sleep(config)
}


// 3. Dapp Part
func main(){
	config, res := statedb.ParseFlags()
	if res == false {
		return
	}
	Start(config)
	if config.Role == "mapper" {
		for i:=0; i<2; i++ {
			var book = make([]string, 5)
			book[0] = "aaa bbb ccc ddd"
			book[1] = "aaa bbb ccc eee"
			book[2] = "aaa bbb ccc fff"
			book[3] = "aaa bbb ccc ggg"
			book[4] = "aaa bbb ccc hhh"
			Mapper(book)
		}
		Ending(config)
	}
	if config.Role == "shuffleAndReduce" {
		ShuffleAndReduce()
	}
	Finish(config)
}

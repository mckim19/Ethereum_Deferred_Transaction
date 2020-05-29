package evm

import (
	statedb "statedb"
)

// 1. opEpcRead, opEpcWrite, opEpcRun
func Read() []byte {
	data := <- statedb.Och
	return data
}
func Write(sendData []byte){
	for idx := range statedb.Ich {
		statedb.Ich[idx] <- sendData
	}
}
func Run(config statedb.Config) {
	statedb.RunNode(config)
}

// 2. opEpcSleep
func Sleep(config statedb.Config) {
	statedb.SleepNode(config)
}

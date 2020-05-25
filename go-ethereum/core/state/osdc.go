// Copyright 2014 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package state

import (
	"fmt"
	"net"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

/*
	OSDC parallel project. Hyojin Jeon.
	Description.

	OSDC parallel project. Yoomee Ko.
	Description.

*/
type ChanMessage struct {
	TxHash          common.Hash
	ContractAddress common.Address
	LockName        int64
	LockType        string
	IsLockBusy      bool
	Channel         chan ChanMessage
}

/*
	OSDC parallel project. Hyojin Jeon.
	Description.
	OSDC parallel project. Yoomee Ko.
	Description.
*/
type RecInfoKey struct {
	ContractAddress common.Address
	LockName        int64
}
type RecInfo map[RecInfoKey][]common.Hash

func (statedb *StateDB) StartMutexThread(type_t int, resChannel chan RecInfo) {
	switch type_t {
	case 0: //when syncing
		statedb.SetChannel(make(chan ChanMessage, 10), true)
		go statedb.MutexThread(statedb.GetChannel(true), true, nil)
		break
	case 1: //when commitTransactions
		statedb.SetChannel(make(chan ChanMessage, 10), false)
		go statedb.MutexThread(statedb.GetChannel(false), false, resChannel)
		break
	case 2: //when w.txsCh
		statedb.SetChannel(make(chan ChanMessage, 10), true)
		go statedb.MutexThread(statedb.GetChannel(true), true, nil)
		break
	case 3: //when state_transition.go + api.go
		statedb.SetChannel(make(chan ChanMessage, 10), true)
		go statedb.MutexThread(statedb.GetChannel(true), true, nil)
		break
	}
}
func (statedb *StateDB) TerminateMutexThread(type_t int) {
	var nil_hash common.Hash
	var nil_address common.Address
	msg := ChanMessage{
		TxHash: nil_hash, ContractAddress: nil_address, LockName: 0, LockType: "TERMINATION",
		IsLockBusy: false, Channel: nil,
	}

	switch type_t {
	case 0: //when syncing
		statedb.GetChannel(true) <- msg
		break
	case 1: //when commitTransactions
		statedb.GetChannel(false) <- msg
		break
	case 2: //when w.txsCh
		statedb.GetChannel(true) <- msg
		break
	case 3: //when state_transition.go + api.go
		statedb.GetChannel(true) <- msg
		break
	}

}

/*
	OSDC parallel project. Hyojin Jeon.
	Description.
	OSDC parallel project. Yoomee Ko.
	Description.

*/
func (self *StateDB) MutexThread(com_channel chan ChanMessage, isDoCall bool, resChannel chan RecInfo) {

	LockRequestArray := make(map[RecInfoKey][]ChanMessage)
	CurrentLockArray := make(map[RecInfoKey]ChanMessage)
	recording_info := make(map[RecInfoKey][]common.Hash)
	for {
		msg := <-com_channel
		//fmt.Println("msg.LockType: ",msg.LockType,", msg.LockName: ",msg.LockName)
		key := RecInfoKey{
			ContractAddress: msg.ContractAddress,
			LockName:        msg.LockName,
		}
		if msg.LockType == "EPC" {
			// Networking
			conn, err := net.Dial("tcp", ":8000")
			if nil != err {
				fmt.Println(err)
			}
			// sending msg to server
			var s string
			s = "hello"
			conn.Write([]byte(s))
			time.Sleep(time.Duration(1) * time.Second)

			// receiving msg from server
			data := make([]byte, 4096)
			n, err := conn.Read(data)
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println("Server send : " + string(data[:n]))

			recording_info[key] = append(recording_info[key], msg.TxHash)
			msg.LockType = "OK"
			msg.IsLockBusy = true
			if CurrentLockArray[key].IsLockBusy == false { //nobody holds this lock
				CurrentLockArray[key] = msg
				msg.Channel <- msg
			} else { //somebody holds this lock
				LockRequestArray[key] = append(LockRequestArray[key], msg)
			}
		} else if msg.LockType == "UNLOCK" {
			// Networking

			if CurrentLockArray[key].TxHash == msg.TxHash { //it must be a same transaction who have been locked
				msg.LockType = "OK"
				msg.Channel <- msg
				if len(LockRequestArray[key]) != 0 { //somebody is waiting
					LockRequestArray[key][0].Channel <- msg
					CurrentLockArray[key] = LockRequestArray[key][0]  //change current lock tx
					LockRequestArray[key] = LockRequestArray[key][1:] //get rid of the current lock tx from lock request array
				} else {
					msg.IsLockBusy = false
					CurrentLockArray[key] = msg
				}
			} else { //somebody tries to unlock fakely
				msg.LockType = "NOT_OK"
				msg.Channel <- msg
			}
		} else if msg.LockType == "TERMINATION" {
			if isDoCall != true { //when mining
				resChannel <- recording_info
			}
			return
		}
	}
}

/*
	OSDC parallel project. Hyojin Jeon.
	Description.

*/
func (self *StateDB) GetChannel(isDoCall bool) chan ChanMessage {
	if isDoCall == true {
		return self.ch_com2
	}
	return self.ch_com

}

/*
	OSDC parallel project. Yoomee Ko.
	Description.

*/
func (self *StateDB) SetChannel(new_ch chan ChanMessage, isDoCall bool) {
	if isDoCall == true {
		self.ch_com2 = new_ch
	}
	self.ch_com = new_ch
}

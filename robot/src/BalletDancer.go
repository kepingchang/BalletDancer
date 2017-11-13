package  BalletDancer

import (
	"mind/core/framework/skill"
	"mind/core/framework/log"
	"mind/core/framework/drivers/hexabody"
)

type  BalletDancer struct {
	skill.Base
}

func NewSkill() skill.Interface {
	// Use this method to create a new skill.

	return & BalletDancer{}
}

func (d * BalletDancer) OnStart() {
	// Use this method to do something when this skill is starting.
}

func (d * BalletDancer) OnClose() {
	// Use this method to do something when this skill is closing.
}

func (d * BalletDancer) OnConnect() {
	// Use this method to do something when the remote connected.
}

func (d * BalletDancer) OnDisconnect() {
	// Use this method to do something when the remote disconnected.
}

func (d * BalletDancer) OnRecvJSON(data []byte) {
	// Use this method to do something when skill receive json data from remote client.
}

func (d * BalletDancer) OnRecvString(data string) {

	log.Info.Println("Skill received", len(data), " bytes command", data)

	doSomeSting(data)

	log.Info.Println("Exceute finished.")
}

func doSomeSting(data string) {

	if data == "stop" {
		hexabody.Close()
	}else {
		hexabody.Start()
		hexabody.WalkContinuously(0, 0.5)
	}
}



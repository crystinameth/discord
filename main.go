package main

import(
	"fmt"
	"github.com/crystinameth/discord/bot"	
	"github.com/crystinameth/discord/config"
)

func main(){
	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bot.Start()

	<-make(chan struct{})
	return
}
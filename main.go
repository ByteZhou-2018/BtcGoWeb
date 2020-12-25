package main

import (
	_ "BtcGoWeb/db"
	_ "BtcGoWeb/routers"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {
	//https://www.blockchain.com/btc/block/  +  height / transaction / address
	fmt.Println("hello  world !")
	//if instance, err := btc.GetBestBlockHash();err ==nil{
	//	fmt.Println("instance",instance)
	//}
	//
	//return

	//db.Query()
	//return
	beego.Run()


}

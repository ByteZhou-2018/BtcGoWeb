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

	//db.Query()
	//return
	//beego.SetStaticPath("","")
	beego.Run()


}

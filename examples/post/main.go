package main

import (
	"os"

	"github.com/kr/pretty"

	// Stdlib

	"log"

	// RPC
	client "github.com/smallnest/steem-api"
	// Vendor
)

// private key to wif by cli_wallet:
//  locked >>> get_private_key_from_password username active "password"

var (
	voter = "xeroc"
	key   = "5JLw5dgQAx6rhZEgNN5C2ds1V47RweGshynFSWFbaMohsYsBvE8"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatal("2 arguments required: ", os.Args)
	}
	voter = os.Args[1]
	key = os.Args[2]

	cls, err := client.NewClient([]string{"wss://gtg.steem.house:8090"}, "steem")
	if err != nil {
		log.Fatalln("Error:", err)
	}

	defer cls.Close()

	cls.SetKeys(&client.Keys{PKey: []string{key}})

	if err := run(cls); err != nil {
		log.Fatalln("Error:", err)
	}
}

func run(cls *client.Client) (err error) {

	resp, err := cls.Post("smallnest", "一个中文名字的标题", "明月松间照，清泉石上流", "", "", "", []string{"test"}, &client.PCOptions{})

	if err != nil {
		return
	}
	log.Println(pretty.Sprint(resp))

	return nil
}

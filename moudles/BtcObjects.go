package moudles

//该结构体用于解析比特币节点返回数据
type BTCResult struct {
	Result interface{} `json:"result"`
	Error  interface{} `json:"error"`
	Id     string      `json:"id"`
}

//该结构体用于创建比特币客户端节点连接请求
type BTCJson struct {
	Jsonrpc string        `json:"jsonrpc"`
	Id      string        `json:"id"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
}

//getblock "hash值" 返回的结构体
type Blcok struct {
	Hash          string   `json:"hash"`
	Confirmations int64    `json:"confirmations"`
	Strippedsize  int64    `json:"strippedsize"`
	Size          int64    `json:"size"`
	Weight        int64    `json:"weight"`
	Height        int64    `json:"height"`
	Version       int64    `json:"version"`
	VersionHex    string   `json:"version_hex"`
	Merkleroot    string   `json:"merkleroot"`
	Tx            []string `json:"merkleroot"`
	Time          int64    `json:"time"`
	Mediantime    int64    `json:"mediantime"`
	Nonce         int64    `json:"nonce"`
	Bits          string   `json:"bits"`
	Difficulty    int64    `json:"difficulty"`
	Chainwork     string   `json:"chainwork"`
	NTx           int64    `json:"n_tx"`
	Nextblockhash string   `json:"nextblockhash"`
}

//getaddressinfo "地址" 结构体
type AddressInfo struct {
	Address             string   `json:"address"`
	ScriptPubkey        string   `json:"scriptPubKey"`
	Ismine              bool     `json:"ismine"`
	solvable            bool     `json:"solvable"`
	Desc                string   `json:"desc"`
	Iswatchonly         bool     `json:"iswatchonly"`
	Isscript            bool     `json:"isscript"`
	Iswitness           bool     `json:"iswitness"`
	Witness_version     int64    `json:"witness_version"`
	Witness_program     string   `json:"witness_program"`
	Pubkey              string   `json:"pubkey"`
	Ischange            bool     `json:"ischange"`
	Timestamp           int64    `json:"timestamp"`
	Hdkeypath           string   `json:"hdkeypath"`
	Hdseedid            string   `json:"hdseedid"`
	Hdmasterfingerprint string   `json:"hdmasterfingerprint"`
	Lables              []string `json:"lables"`
}

//getblockchaininfo 命令返回的结构体
type BlockChainInfo struct {
	Chain                string    `json:"chain"`
	Blocks               int64     `json:"blocks"`
	Headers              int64     `json:"headers"`
	Bestblockhash        string    `json:"bestblockhash"`
	Difficulty           int64     `json:"difficulty"`
	Mediantime           int64     `json:"mediantime"`
	Verificationprogress float64   `json:"verificationprogress"`
	Initialblockdownload bool      `json:"initialblockdownload"`
	Chainwork            string    `json:"chainwork"`
	Size_on_disk         int64     `json:"size_on_disk"`
	Pruned               bool      `json:"pruned"`
	Softforks            softforks `json:"softforks"`
	Warnings             string    `json:"warnings"`
}

//getblockchaininfo 命令返回的结构体的数据中的一个子集嵌套 softforks
type softforks struct {
	Bip34  softfork `json:"bip_34"`
	Bip66  softfork `json:"bip_66"`
	Bip65  softfork `json:"bip_65"`
	Csv    softfork `json:"csv"`
	Segwit softfork `json:"segwit"`
}

//getblockchaininfo 命令返回的结构体的数据中的一个子集嵌套 softforks 的嵌套
type softfork struct {
	Type   string `json:"type"`
	Active bool   `json:"active"`
	Height int64  `json:"height"`
}

//getblcokheader 命令返回的结构体
type BlockHeader struct {
	Hash          string `json:"hash"`
	Confirmations int64 `json:"confirmations"`
	Height        int64 `json:"height"`
	Version       int64 `json:"version"`
	VersionHex    string `json:"version_hex"`
	Merkleroot    string `json:"merkleroot"`
	Time          int64 `json:"time"`
	Mediantime    int64 `json:"mediantime"`
	Nonce         int64 `json:"nonce"`
	Bits          string `json:"bits"`
	Difficulty    int64 `json:"difficulty"`
	Chainwork     string `json:"chainwork"`
	NTx           int64 `json:"n_tx"`
	Nextblockhash string `json:"nextblockhash"`
}

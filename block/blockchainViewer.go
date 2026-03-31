package block

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

func View(blocksAmount uint8, hashStrength uint) {
	data, _ := os.ReadFile("block.1")
	//fmt.Println(string(data))
	blockChain := make([]Block, blocksAmount)
	genesis := Block{
		[]byte{},
		data,
		0,
		[]byte{},
	}
	ComputeBlock(&genesis)
	FindNonce(&genesis, hashStrength)
	blockChain[0] = genesis
	var newBlock Block
	for i := 1; i < int(blocksAmount); i++ {
		data, err := os.ReadFile("block." + strconv.Itoa(i+1))
		if err != nil {
			fmt.Println("Error:", err)
			fmt.Println("Looks like there isn't enough data for the blocks...")
			return
		}
		newBlock.Data = data
		newBlock.Nonce = 0
		newBlock.HashPrevious = blockChain[i-1].HashCurrent
		ComputeBlock(&newBlock)
		FindNonce(&newBlock, hashStrength)
		blockChain[i] = newBlock
	}

	for i, bl := range blockChain {
		fmt.Println("Block #", i+1)
		fmt.Println("\tData:", string(bl.Data))
		fmt.Println("\tNonce:", bl.Nonce)
		fmt.Println("\tBlock's Hash:", hex.EncodeToString(bl.HashCurrent))

		f, _ := os.Create("block." + strconv.Itoa(i+1) + ".json")
		encoder := json.NewEncoder(f)
		encoder.SetIndent("", "    ")
		encoder.Encode(bl)
		f.Close()
	}

	fmt.Println("######### VALIDATING ########")
	loadedBlockchain := make([]Block, blocksAmount)
	for i := 0; i < int(blocksAmount); i++ {
		data, _ := os.ReadFile("block." + strconv.Itoa(i+1) + ".json")
		json.Unmarshal(data, &loadedBlockchain[i])
		if !ValidateBlock(&loadedBlockchain[i]) {
			fmt.Printf("Block #%d is not valid...\n", i+1)
			continue
		}
		fmt.Printf("Block #%d is fine!!!\n", i+1)
	}
}

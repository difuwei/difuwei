package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"os"
	"strconv"
	"time"
)

/*

 */
type Block struct {
	Timestamp int64
	Data []byte
	PrevBlockHash []byte
	Hash []byte
}

/**
Block结构体
 */
func (this *Block) SetHash(){
	//将本区块的Timestamp,data,PrevBlockHash
	//整型转为二进制
	timestamp := []byte(strconv.FormatInt(this.Timestamp,10))
	//
	headers :=bytes.Join([][]byte{this.PrevBlockHash,this.Data,timestamp},[]byte{})

	//拼接之后sha256加密
	hash := sha256.Sum256(headers)

	this.Hash = hash[:]
}

/**

 */
func NewBlock(data string,prevBlockHash []byte) *Block {
	//生成一个区块
	block := Block{}
	//给当前区块赋值
	block.Timestamp = time.Now().Unix()
	block.Data = []byte(data)
	block.PrevBlockHash = prevBlockHash
	block.SetHash()

	return &block

}
/**
定义一个区块链结构
 */
 type BlockChain struct {
 	Blocks []*Block
 }

func NewGenesisBlock() *Block  {
	genesisBlock := Block{}
	genesisBlock.Data = []byte("the first block")
	genesisBlock.PrevBlockHash = []byte{}

	return &genesisBlock

}

func NewBlockChain() *BlockChain  {
	//blockChain := BlockChain{}
	//blockChain.Blocks = append(blockChain.Blocks,NewGenesisBlock())
	//
	//return &blockChain
	return &BlockChain{[]*Block{NewGenesisBlock()}}
}
func (this *BlockChain) AddBlock(data string)  {
	prevBlock := this.Blocks[len(this.Blocks)-1]
	newBlock := NewBlock(data,prevBlock.PrevBlockHash)
	this.Blocks = append(this.Blocks,newBlock)
}
func main()  {
	//block := NewBlock("hhh",[]byte)
	//fmt.Print("")
	bc := NewBlockChain()

	var cmd string
	for {
		fmt.Println("按1")
		fmt.Println("按2")
		fmt.Println("按其他退出")
		fmt.Scanf("%s",&cmd)
		switch cmd {
		case "1":
			input := make([]byte,1024)
			fmt.Println("输入区块数据")
			os.Stdin.Read(input)
			bc.AddBlock(string(input))
		case "2":
			for i,block := range bc.Blocks{
				fmt.Println("=====")
				fmt.Println("第",i,"区块:")
				fmt.Printf("PrevHash:%x\n",block.PrevBlockHash)
			}
		default:
			fmt.Println("退出")
			return
			}
	}
}
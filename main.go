package main

import (
	"context"
	"fmt"
	"log"

	// "os"
	"strings"

	// "time"

	// "github.com/spf13/viper"
	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	// "golang.org/x/crypto/bcrypt"
)

func main() {
	// 1. เชื่อมต่อ RPC
	client, err := ethclient.Dial("https://arise-rpc-nonprd.arisetech.dev")
	if err != nil {
		log.Fatal(err)
	}

	contractAddr := common.HexToAddress("0x52C57275A76f687E30912948A8a1Bb2FaafC986D")
	userToTest := common.HexToAddress("0xe8f9f81cb78f6096b10515d9d26750ebfeaffd5d")

	// 2. เตรียมข้อมูลสำหรับ Call (เนื่องจาก Binding ปกติมักจะกลืน Error Data ทิ้ง)
	parsedAbi, _ := abi.JSON(strings.NewReader(TestErrorMetaData.ABI))
	input, err := parsedAbi.Pack("IsOwner", userToTest)
	if err != nil {
		log.Fatal(err)
	}

	// 3. ใช้ CallContract เพื่อให้ได้ Error แบบละเอียด
	msg := ethereum.CallMsg{
		To:   &contractAddr,
		Data: input,
	}

	_, err = client.CallContract(context.Background(), msg, nil)
	if err != nil {
		// ส่ง error ที่ได้ไปแกะรหัส
		decodedErrMsg := DecodeSmartContractError(err)
		fmt.Printf("\n--- ERROR DECODED ---\n%s\n---------------------\n", decodedErrMsg)
	} else {
		fmt.Println("Success: User is the owner!")
	}
}

func DecodeSmartContractError(targetErr error) string {
	contractAbi, _ := abi.JSON(strings.NewReader(TestErrorMetaData.ABI))
	var data []byte

	// 1. ดึงจาก ErrorData
	if rpcErr, ok := targetErr.(interface{ ErrorData() interface{} }); ok {
		if rpcData, ok := rpcErr.ErrorData().(string); ok {
			data, _ = hexutil.Decode(rpcData)
		}
	}

	// 2. ดึงจาก String message
	if len(data) == 0 {
		errStr := targetErr.Error()
		if strings.Contains(errStr, "0x") {
			parts := strings.Split(errStr, "0x")
			hexPart := "0x" + parts[len(parts)-1]
			data, _ = hexutil.Decode(hexPart)
		}
	}

	// 3. Unpack (เวอร์ชัน Arguments.Unpack([]byte) ([]interface{}, error))
	if len(data) >= 4 {
		var methodSig [4]byte
		copy(methodSig[:], data[:4])

		errorObj, err := contractAbi.ErrorByID(methodSig)
		if err == nil {
			// แก้ไขจุดนี้ให้ตรงกับเวอร์ชันของคุณ
			unpacked, err := errorObj.Inputs.Unpack(data[4:])
			if err == nil && len(unpacked) > 0 {
				// unpacked[0] คือค่าแรกใน error NoAuthorization(address user)
				userAddr, ok := unpacked[0].(common.Address)
				if ok {
					return fmt.Sprintf("Reason: %s | Arg User: %s", errorObj.Name, userAddr.Hex())
				}
			}
		}
	}

	return targetErr.Error()
}

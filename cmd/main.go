package main

import (
	"fmt"
	tronWallet "github.com/zx9597446/tron-wallet"
	"github.com/zx9597446/tron-wallet/enums"
)

func main() {

	//c, _ := grpcClient.GetGrpcClient(enums.SHASTA_NODE)
	//
	w, _ := tronWallet.CreateTronWallet(enums.SHASTA_NODE, "88414dbb373a211bc157265a267f3de6a4cec210f3a5da12e89630f2c447ad27")

	fmt.Println(w.EstimateTransferFee("TJnsY5bGiwuPCQFismQDwyVTPAn7M88888", 1000))
	fmt.Println(w.EstimateTransferERC20Fee())
	//priv, _ := w.PrivateKeyRCDSA()
	//
	//fmt.Println(tronWallet.EstimateTrc10TransactionFee(enums.SHASTA_NODE, priv, w.AddressBase58, "TJnsY5bGiwuPCQFismQDwyVTPAn7M88888", 10000))
	//
	//return

	//fmt.Println(c.GetAccountResource(w.Address))
	//fmt.Println(w.Transfer("TJnsY5bGiwuPCQFismQDwyVTPAn7M88888", 10000))
	//fmt.Println(c.GetAccountResource(w.Address))

	fmt.Println(w.TransferTRC20(&tronWallet.Token{
		ContractAddress: enums.SHASTA_Tether_USDT,
	}, "TJnsY5bGiwuPCQFismQDwyVTPAn7M88888", 10))
}

package main

import (
	"github.com/cnvic/huobi_golang/v3/cmd/accountclientexample"
	"github.com/cnvic/huobi_golang/v3/cmd/accountwebsocketclientexample"
	"github.com/cnvic/huobi_golang/v3/cmd/algoorderclientexample"
	"github.com/cnvic/huobi_golang/v3/cmd/commonclientexample"
	"github.com/cnvic/huobi_golang/v3/cmd/crossmarginclientexample"
	"github.com/cnvic/huobi_golang/v3/cmd/etfclientexample"
	"github.com/cnvic/huobi_golang/v3/cmd/isolatedmarginclientexample"
	"github.com/cnvic/huobi_golang/v3/cmd/marketclientexample"
	"github.com/cnvic/huobi_golang/v3/cmd/marketwebsocketclientexample"
	"github.com/cnvic/huobi_golang/v3/cmd/orderclientexample"
	"github.com/cnvic/huobi_golang/v3/cmd/orderwebsocketclientexample"
	"github.com/cnvic/huobi_golang/v3/cmd/stablecoinclientexample"
	"github.com/cnvic/huobi_golang/v3/cmd/subuserclientexample"
	"github.com/cnvic/huobi_golang/v3/cmd/walletclientexample"
	"github.com/cnvic/huobi_golang/v3/logging/perflogger"
)

func main() {
	runAll()
}

// Run all examples
func runAll() {
	commonclientexample.RunAllExamples()
	accountclientexample.RunAllExamples()
	orderclientexample.RunAllExamples()
	algoorderclientexample.RunAllExamples()
	marketclientexample.RunAllExamples()
	isolatedmarginclientexample.RunAllExamples()
	crossmarginclientexample.RunAllExamples()
	walletclientexample.RunAllExamples()
	subuserclientexample.RunAllExamples()
	stablecoinclientexample.RunAllExamples()
	etfclientexample.RunAllExamples()
	marketwebsocketclientexample.RunAllExamples()
	accountwebsocketclientexample.RunAllExamples()
	orderwebsocketclientexample.RunAllExamples()
}

// Run performance test
func runPerfTest() {
	perflogger.Enable(true)
	commonclientexample.RunAllExamples()
	accountclientexample.RunAllExamples()
	orderclientexample.RunAllExamples()
	marketclientexample.RunAllExamples()
	isolatedmarginclientexample.RunAllExamples()
	crossmarginclientexample.RunAllExamples()
	walletclientexample.RunAllExamples()
	etfclientexample.RunAllExamples()
}

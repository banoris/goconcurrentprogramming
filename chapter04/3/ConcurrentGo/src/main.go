package main

import (
    "fmt"
)

func main() {
    po := new(PurchaseOrder)
    po.Value = 42.27

    ch := make(chan *PurchaseOrder)

    // assume SavePO is expensive. With go keyword,
    // we are running it asynchronously
    go SavePO(po, ch)

    // do some operation here...
    // some more work...
    newPo := <- ch
    fmt.Printf("PO Number: %d\n", newPo.Number)
}

type PurchaseOrder struct {
    Number int
    Value float64
}

// Assume this is a long running task. Like do network+db+calculation
// before returning. How can caller exec this function without blocking
// itself? Use go keyword! And how does caller knows when this func
// completes? Simple, by draining the callbackChannel.
func SavePO(po *PurchaseOrder, callbackChannel chan *PurchaseOrder) {
    po.Number = 1234

    callbackChannel <- po
}

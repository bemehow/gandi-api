package operation

import (
    "github.com/kolo/xmlrpc"
    "github.com/prasmussen/gandi-api/client"
)

type Operation struct {
    *client.Client
}

func New(c *client.Client) *Operation {
    return &Operation{c}
}

// Count operations created by this contact
func (self *Operation) Count() (int, error) {
    var result int64
    params := xmlrpc.Params{xmlrpc.Params: []interface{}{self.Key}}
    if err := self.Rpc.Call("operation.count", params, &result); err != nil {
        return -1, err
    }
    return int(result), nil
}

// Get operation information
func (self *Operation) Info(id int) (*OperationInfo, error) {
    var res map[string]interface{}
    params := xmlrpc.Params{xmlrpc.Params: []interface{}{self.Key, id}}
    if err := self.Rpc.Call("operation.info", params, &res); err != nil {
        return nil, err
    }
    return ToOperationInfo(res), nil
}

// Cancel an operation
func (self *Operation) Cancel(id int) (bool, error) {
    var res bool
    params := xmlrpc.Params{xmlrpc.Params: []interface{}{self.Key, id}}
    if err := self.Rpc.Call("operation.cancel", params, &res); err != nil {
        return false, err
    }
    return res, nil
}

// List operations created by this contact
func (self *Operation) List() ([]*OperationInfo, error) {
    var res []interface{}
    params := xmlrpc.Params{xmlrpc.Params: []interface{}{self.Key}}
    if err := self.Rpc.Call("operation.list", params, &res); err != nil {
        return nil, err
    }

    operations := make([]*OperationInfo, len(res), len(res))
    for i, r := range res {
        operations[i] = ToOperationInfo(r.(xmlrpc.Struct))
    }
    return operations, nil
}
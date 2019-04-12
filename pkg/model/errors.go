package model

import (
	"errors"
	"fmt"
	"github.com/go-mesh/openlogging"
)

//ErrAction will wrap raw error to biz error and return
//it record audit log for mongodb operation failure like find, insert, update, deletion
func ErrAction(action, key string, labels Labels, domain string, err error) error {
	msg := fmt.Sprintf("can not [%s] [%s] in [%s] with [%s],err: %s", action, key, domain, labels, err.Error())
	openlogging.Error(msg)
	return errors.New(msg)

}

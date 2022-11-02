/**
* @Author: qinwenjie
* @Date: 2022/11/2 10:34
* @Description:
 */

package exception

import (
	"errors"
	"fmt"
)

func TryCatch(method func() error) (err error) {
	defer func() {
		if p := recover(); p != nil {
			err = errors.New(fmt.Sprintf("occurring error, please note that: %v", p))
		}
	}()
	err = method()
	return
}

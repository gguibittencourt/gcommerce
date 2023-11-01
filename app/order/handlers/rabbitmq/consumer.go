package rabbitmq

import (
	"fmt"
)

func ConsumeMsg(msg any) error {
	fmt.Println(msg)
	return nil
}

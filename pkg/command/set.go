package command

import (
	"errors"
	"strings"
	"time"

	datatypes "github.com/betty2310/redigo/pkg/data-types"
)

func Set(args []string) (key string, value datatypes.RedisValue, err error) {
	if len(args) == 5 {
		if strings.ToLower(args[3]) == "exp" {
			key = args[1]
			value = datatypes.RedisValue{}
			duration, e := time.ParseDuration(args[4])
			if e != nil {
				err = e
				return
			}
			exp := time.Now().Add(duration)
			value = datatypes.RedisValue{Value: datatypes.RedisString(args[2]), Exp: exp}
			return
		} else {
			err = errors.New("wrong arguments")
			return
		}
	} else if len(args) == 3 {
		key = args[1]
		value = datatypes.RedisValue{Value: datatypes.RedisString(args[2]), Exp: time.Now().Add(time.Hour * 24)}
		return
	} else {
		err = errors.New("wrong arguments")
		return
	}
}

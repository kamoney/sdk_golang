package utility

import (
	"time"
)

func GenNonce() int64 {
	mt := time.Now().UnixNano()

	seconds := mt / int64(time.Second)
	nanoseconds := mt % int64(time.Second)

	return seconds + nanoseconds/1e6
}

package support

import "hash/fnv"

func StrToHashCode(s string) int64 {
	f := fnv.New32a()
	_, _ = f.Write([]byte(s))
	return int64(f.Sum32())
}

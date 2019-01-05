package main

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/binary"
	"hash/fnv"
	"os"
	"strconv"

	"github.com/valyala/fastrand"

	"github.com/cespare/xxhash"

	"github.com/lemonwx/log"
)

var (
	dicts map[string]struct{}
	set   map[int]map[int64]struct{}
)

const (
	MinLen = 2
	MaxLen = 50

	Fnv = iota
	xxHash
	md5_
	sha1_
)

func randString() []byte {
	randStrLen := int(fastrand.Uint32n(MaxLen-MinLen) + MinLen)
	randStr := make([]byte, 0, randStrLen)
	for idx := 0; idx < randStrLen; idx += 1 {
		randStr = append(randStr, byte(fastrand.Uint32()))
	}
	return randStr
}

func randUniqueStr() []byte {
	for {
		str := randString()
		_, ok := dicts[string(str)]
		if !ok {
			dicts[string(str)] = struct{}{}
			return str
		}
	}
}

func sha1toInt64(data []byte) int64 {
	b := sha1.Sum(data)
	b1 := b[12:20]
	b2 := b[4:12]
	b3 := b[0:4]
	for idx := 0; idx < 8; idx += 1 {
		b1[idx] = b1[idx] ^ b2[idx]
	}
	for idx := 0; idx < 4; idx += 1 {
		b1[7-idx] = b1[7-idx] ^ b3[idx]
	}
	return int64(binary.LittleEndian.Uint64(b1))
}

func md5toInt64(data []byte) int64 {
	b := md5.Sum(data)
	b1 := b[8:16]
	b2 := b[0:8]
	for idx := 0; idx < 8; idx += 1 {
		b1[idx] = b1[idx] ^ b2[idx]
	}
	return int64(binary.LittleEndian.Uint64(b1))
}

func Hash(data []byte, hashType int) int64 {
	switch hashType {
	case Fnv: //0.02
		h := fnv.New64()
		h.Reset()
		h.Write(data)
		return int64(h.Sum64())
	case xxHash: // 0.02
		h := xxhash.New()
		h.Reset()
		h.Write(data)
		return int64(h.Sum64())
	case sha1_: // 0.02
		return sha1toInt64(data)
	case md5_:
		return md5toInt64(data)
	}
	return 0
}

func conflictTest(hashType int) {
	dicts = map[string]struct{}{}
	conflictCount := 0
	for idx := 0; idx < count; idx += 1 {
		x := Hash(randUniqueStr(), hashType)
		//x := Hash(randString(), hashType)
		if _, ok := set[hashType][x]; ok {
			conflictCount += 1
		} else {
			set[hashType][x] = struct{}{}
		}
	}
	log.Debug(hashType, conflictCount, float64(conflictCount)/float64(count), count)
}

var count int

func mainHashTest() {
	x, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	count = x
	// hashTypes := []int{Fnv, xxHash, md5_, sha1_}
	hashTypes := []int{Fnv}

	set = map[int]map[int64]struct{}{}
	for _, t := range hashTypes {
		set[t] = map[int64]struct{}{}
	}
	for _, t := range hashTypes {
		conflictTest(t)
	}
	log.Debug("----------------")
}

func mainRandStrConflict() {
	dicts = map[string]struct{}{}
	x, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	count = x
	m := map[string]struct{}{}
	conflictCnt := 0
	for idx := 0; idx < count; idx += 1 {
		// x := string(randUniqueStr())
		x := string(randString())
		// x = uuid.Gen()
		_, ok := m[x]
		if ok {
			conflictCnt += 1
		} else {
			m[x] = struct{}{}
		}
	}
	log.Debug(conflictCnt, float64(conflictCnt)/float64(count), count)
}

func main() {
	mainHashTest()
	//mainRandStrConflict()
}

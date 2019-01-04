package main

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/binary"
	"hash/fnv"
	"os"
	"strconv"

	"code.byted.org/inf/hb/bytable_master/utils/uuid"

	"github.com/valyala/fastrand"

	"github.com/cespare/xxhash"

	"github.com/lemonwx/log"
)

var (
	dict []byte
	set  []map[int64]struct{}
)

func init() {
	set = []map[int64]struct{}{
		map[int64]struct{}{}, map[int64]struct{}{},
		map[int64]struct{}{}, map[int64]struct{}{},
	}
	for idx := 'a'; idx < 'z'; idx += 1 {
		dict = append(dict, byte(idx))
	}
	for idx := 'A'; idx < 'Z'; idx += 1 {
		dict = append(dict, byte(idx))
	}

	str := "哈希算法是一个大杂烩，除了 MD5、SHA1 这一类加密哈希算法（cryptographic hash），还有很多或家喻户晓或寂寂无闻的算法。哈希算法只需满足把一个元素映射到另一个区间的要求。鉴于该要求是如此之低，像 Java 的 hashCode 其实也算一种哈希算法。不过今天当然不会讲 Java 的 hashCode，我们来聊聊一些更具实用性的算法，某些场景下它们不仅可以替代加密哈希算法，甚至还有额外的优势。评价一个哈希算法的好坏，人们通常会引用 SMHasher 测试集的运行结果。从这里可以看到关于它的介绍：SMHasher wiki。此外还有人在 SMHasher 的基础上，添加更多的测试和哈希算法：demerphq/smhasher。demerphq/smhasher 这个项目有个好处，作者把 SMhasher 运行结果放到 doc 目录下，为本文的分析提供了丰富的数据。在需要进行比较的时候，我会选择 MD5 作为加密哈希算法阵营的代表。原因有三：MD5 通常用在不需要加密的哈希计算中，以致于有些场合下“哈希”就意味着计算 MD5。MD5 计算速度跟其他加密哈希算法差不多。MD5 生成的哈希值是 128 比特的。这里的哈希值指的是二进制的值，而不是 HEX 或 base64 格式化后的人类可读的值。通常我们提到的 32 位 MD5 是指由 32 个字符组成的，HEX 格式的 MD5。下面提到的 32 位、128 位，如果没有特殊说明，都是指比特数。MurMurHash设想这样的场景：当数据中有几个字段相同，就可以把它当作同类型数据。对于同类型的数据，我们需要做去重。一个通常的解决办法是，使用 MD5 计算这几个字段的指纹，然后把指纹存起来。如果当前指纹存在，则表示有同类型的数据。这种情况下，由于不涉及故意的哈希碰撞，并不一定要采用加密哈希算法。（加密哈希算法的一个特点是，即使你知道哈希值，也很难伪造有同样哈希值的文本）而非加密哈希算法通常要比加密哈希算法快得多。如果数据量小，或者不太在意哈希碰撞的频率，甚至可以选择生成哈希值小的哈希算法，占用更小的空间。就个人而言，我偏好在这种场景里采用 MurMurHash3，128 位的版本。理由有三：MurMurHash3 久经沙场，主流语言里面都能找到关于它的库。MurMurHash3，128 位版本的哈希值是 128 位的，跟 MD5 一样。128 位的哈希值，在数据量只有千万级别的情况下，基本不用担心碰撞。MurMurHash3 的计算速度非常快。MurMurHash3 哈希算法是 MurMurHash 算法家族的最新一员。虽说是“最新一员”，但距今也有五年的历史了。无论从运算速度、结果碰撞率，还是结果的分布均衡程度进行评估，MurMurHash3 都算得上一个优秀的哈希算法。除了 128 位版本以外，它还有生成 32 位哈希值的版本。在某些场景下，比如哈希的对象长度小于 128 位，或者存储空间要求占用小，或者需要把字符串转换成一个整数，这一特性就能帮上忙。当然，32 位哈希值发生碰撞的可能性就比 128 位的要高得多。当数据量达到十万时，就很有可能发生碰撞。贴一个简略的 MurMurHash3 和 MD5、MurMurHash2 的 benchmark：lua-resty-murmurhash3#benchmark可以看到，MurMurHash3 128 位版本的速度是 MD5 的十倍。有趣的是，MurMurHash3 生成 32 位哈希的用时比生成 128 位哈希的用时要长。原因在于生成 128 位哈希的实现受益于现代处理器的特性。CRC32MurMurHash3 是我的心头好，另一个值得关注的是 CRC 系列哈希算法。诸位都知道，CRC 可以用来算校验和。除此以外，CRC 还可以当作一个哈希函数使用。目前常用的 CRC 算法是 CRC32，它可以把一个字符串哈希成 32 位的值。CRC32 的碰撞率要比 MurMurHash3（32位）低，可惜它的运算速度跟 MD5 差不多。一个 32 位的哈希值，再怎么样，碰撞的概率还是比 128 位的多得多。更何况选用非加密哈希算法，运算速度往往是首先考虑的。看样子 CRC32 要出局了。好在 CRC 系列一向有硬件加持。只要 CPU 支持 sse4.2 特性，就能使用 _mm_crc32_* 这一类硬件原语。（参见 Intel 的 文档，更详细的用法可以在网上搜到）硬件加持之后，CRC32 的计算速度可以达到十数倍以上的提升。换个说法，就是有硬件加持的 CRC32，比 MurMurHash3 要快。不过要注意的是，有 sse4.2 加持的是 CRC32c，它是 CRC32 的一个变种。CRC32c 跟我们通常用的 CRC32 并不兼容。所以如果你要编写的程序会持久化 CRC32 哈希值，在使用硬件加速之前先关注这一点。FNV跟 FNV 的初次邂逅，是在 Go 的标准库文档里。我很惊讶，一门主流语言的标准库里，居然会有不太知名的哈希算法。从另一个角度看，非加密哈希算法还是有其必要性的，不然 Go 这门以实用著称的语言也不会内置 FNV 算法了。可惜跟其他哈希算法相比，FNV 并无出彩之处（参考 SMHasher 测试结果）。FNV 家族较新的 FNV-1a 版本，对比于 FNV-1 做了一些改善。尽管如此，除非写的是 Go 程序，我通常都不会考虑使用它。SipHash大部分非加密哈希算法的改良，都集中在让哈希速度更快更好上。SipHash 则是个异类，它的提出是为了解决一类安全问题：hash flooding。通过让输出随机化，SipHash 能够有效减缓 hash flooding 攻击。凭借这一点，它逐渐成为 Ruby、Python、Rust 等语言默认的 hash 表实现的一部分。如果你愿意尝试下新技术，可以看看 2016 新出的 HighwayHash。它宣称可以达到 SipHash 一样的效果，并且凭借 SIMD 的加持，在运算速度上它是 SipHash 的 5.2 倍（参考来源：https://arxiv.org/abs/1612.06257）。xxHash为什么要用一个不知名的哈希函数？最首要的理由就是性能。对性能的追求是无止境的。如果不考虑硬件加持的 CRC32，xxHash可以说是哈希函数性能竞赛的最新一轮优胜者。xxHash 支持生成 32 位和 64 位哈希值，多个 benchmark 显示，其性能比 MurMurHash 的 32 位版本快接近一倍。如果程序的热点在于哈希操作，作为一种优化手段，xxHash 值得一试。顺便一提，MetroHash 声称其速度在 xxHash 之上。不过前者用的人不多，如果想尝鲜，可以关注下。"
	for _, s := range str {
		dict = append(dict, byte(s))
	}
}

const (
	MinLen = 2
	MaxLen = 50
)

func randString() []byte {

	randStrLen := int(fastrand.Uint32n(MaxLen-MinLen) + MinLen)
	// dictLen := len(dict)

	var randStr []byte
	for idx := 0; idx < randStrLen; idx += 1 {
		// randIdx := fastrand.Uint32n(uint32(dictLen - 1))
		// randStr = append(randStr, dict[randIdx])
		randStr = append(randStr, byte(fastrand.Uint32()))
	}
	return randStr

	//return []byte(uuid.Gen())
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

func Hash(data []byte, flag int) int64 {
	switch flag {
	case 0: //0.02
		h := fnv.New64()
		h.Reset()
		h.Write(data)
		return int64(h.Sum64())
	case 1: // 0.02
		h := xxhash.New()
		h.Reset()
		h.Write(data)
		return int64(h.Sum64())
	case 2: // 0.02
		return sha1toInt64(data)
	case 3:
		return md5toInt64(data)
	}
	return 0
}

func conflictTest(flag int) {
	conflictCount := 0
	for idx := 0; idx < count; idx += 1 {
		x := Hash(randString(), flag)
		if _, ok := set[flag][x]; ok {
			conflictCount += 1
		} else {
			set[flag][x] = struct{}{}
		}
	}
	log.Debug(flag, conflictCount, float64(conflictCount)/float64(count), count)
}

var count int

func mainHashTest() {
	x, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	count = x
	loop := 10
	for j := 0; j < loop; j += 1 {
		set = []map[int64]struct{}{
			map[int64]struct{}{}, map[int64]struct{}{},
			map[int64]struct{}{}, map[int64]struct{}{},
		}
		for idx := 0; idx < 4; idx += 1 {
			conflictTest(idx)
		}
		log.Debug("----------------")
	}
}

func mainRandStrConflict() {
	m := map[string]struct{}{}
	count := 1000000
	conflictCnt := 0
	for idx := 0; idx < count; idx += 1 {
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
	log.Debug(uuid.Gen())
	log.Debug(uuid.Gen())
	log.Debug(uuid.Gen())
	mainRandStrConflict()
}

package common

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"github.com/weilinks/gotcp/message"
	"io"
	"strconv"
)

var (
	// 完整的数据包
	packet = make([]byte, 0)
	// 缓存
	cache = make([]byte, 0)
)

// 数据格式： 0x02+数据类型+0x03
// 0x027b224964223a312c226d0x03
const (
	// 分隔符长度
	SeparatorSize = 4
	// 分隔符前缀
	PacketPrefix = 0x02
	// 分隔符后缀
	PacketSuffix = 0x03
)

// NewDefaultPacker create a *DefaultPacker with initial field value.
func NewCustomPacker() *CustomPacker {
	return &CustomPacker{
		MaxDataSize: 1 << 10 << 10, // 1MB
	}
}

// DefaultPacker is the default Packer used in session.
// Treats the packet with the format:
//
// head(4)|data(n)|tail(4)
//
// | segment    | type   | size    | remark                  |
// | ---------- | ------ | ------- | ----------------------- |
// | `head` | uint32 | 4       | the size of `data` only |
// | `data`     | []byte | dynamic |                         |
// | `tail`       | uint32 | 4       |                         |
// .
type CustomPacker struct {
	// MaxDataSize represents the max size of `data`
	MaxDataSize int
}

func (d *CustomPacker) bytesOrder() binary.ByteOrder {
	return binary.BigEndian
}

// 0x02+content+0x03
// 0x02{"devId":"CHZD12XHSO190923001","devType":2}0x03
// Pack implements the Packer Pack method.
func (d *CustomPacker) Pack(entry *message.Entry) ([]byte, error) {
	// 将源二进制数据转成16进制的字符串，再转成二进制数组，得到最终发送的二进制数据
	hexStringData := hex.EncodeToString(entry.Data)
	data := []byte(hexStringData)

	dataSize := len(data)

	buffer := make([]byte, 4+dataSize+4)
	d.bytesOrder().PutUint32(buffer[:4], uint32(0x02)) // write dataSize
	copy(buffer[4:dataSize+4], data)              // write data
	d.bytesOrder().PutUint32(buffer[4+dataSize:], uint32(0x03))    // write idSize

	return buffer, nil
}

// 解包
func unpack(buf []byte) ([]byte) {
	buf = append(cache, buf...)
	length := len(buf)

	if length < SeparatorSize * 2 {
		cache = buf
		return []byte{}
	}

	pIndex := bytes.Index(buf, intTo4Bytes(PacketPrefix))
	sIndex := bytes.Index(buf, intTo4Bytes(PacketSuffix))

	if pIndex == -1 || sIndex == -1 {
		cache = buf
		return []byte{}
	}

	if pIndex > sIndex {
		cache = []byte{}
		return []byte{}
	}
	cache = buf[sIndex+SeparatorSize:]

	data := buf[pIndex + SeparatorSize : sIndex]

	hexBytes, _ := hex.DecodeString(string(data))

	return hexBytes
}

func (d *CustomPacker) Unpack(reader io.Reader) (*message.Entry, error) {
	// 每次读取的字节流
	buf := make([]byte, 4)

	for {
		n, err := reader.Read(buf)
		if err != nil {
			if err == io.EOF {
				continue
			} else {
				return nil, err
			}
		} else {
			packet := unpack(buf[:n])
			if len(packet) > 0 {
				msgType := &MsgType{}
				if err := json.Unmarshal(packet, msgType); err == nil {
					entry := &message.Entry{
						ID:   msgType.MsgType,
						Data: packet,
					}
					return entry, nil
				} else {
					return nil, err
				}
			}
		}
	}
}


// bytes to int 32
func bytesTo32Int(b []byte) int {
	buf := bytes.NewBuffer(b)
	var tmp uint32
	binary.Read(buf, binary.BigEndian, &tmp)
	return int(tmp)
}

// bytes to int 16
func bytesTo16Int(b []byte) int {
	buf := bytes.NewBuffer(b)
	var tmp uint16
	binary.Read(buf, binary.BigEndian, &tmp)
	return int(tmp)
}

// int to 4 bytes
func intTo4Bytes(i int) []byte {
	buf := bytes.NewBuffer([]byte{})
	tmp := uint32(i)
	binary.Write(buf, binary.BigEndian, tmp)
	return buf.Bytes()
}

// int to 2 bytes
func intTo2Bytes(i int) []byte {
	buf := bytes.NewBuffer([]byte{})
	tmp := uint16(i)
	binary.Write(buf, binary.BigEndian, tmp)
	return buf.Bytes()
}

// bytes to hex string
func bytesToHexString(b []byte) string {
	var buf bytes.Buffer
	for _, v := range b {
		t := strconv.FormatInt(int64(v), 16)
		if len(t) > 1 {
			buf.WriteString(t)
		} else {
			buf.WriteString("0" + t)
		}
	}
	return buf.String()
}

// hex string to bytes
func hexStringToBytes(s string) []byte {
	bs := make([]byte, 0)
	for i := 0; i < len(s); i = i + 2 {
		b, _ := strconv.ParseInt(s[i:i+2], 16, 16)
		bs = append(bs, byte(b))
	}
	return bs
}

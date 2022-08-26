package lsm

import (
	"bytes"
	"encoding/binary"
	"io"
	"kvdb/pkg/utils"

	"github.com/golang/snappy"
)

const restartInterval = 16
const blockTrailerlen = 4

type Block struct {
	restartInterval int

	header             [30]byte
	record             *bytes.Buffer
	trailer            *bytes.Buffer
	nEntries           int
	prevKey            []byte
	compressionScratch []byte
}

func (d *Block) Append(key, value []byte) {
	keyLen := len(key)
	valueLen := len(value)
	nSharePrefix := 0

	// restart point
	if d.nEntries%d.restartInterval == 0 {
		buf4 := make([]byte, 4)
		binary.LittleEndian.PutUint32(buf4, uint32(d.record.Len()))
		d.trailer.Write(buf4)
	} else {
		nSharePrefix = SharedPrefixLen(d.prevKey, key)
	}

	n := binary.PutUvarint(d.header[0:], uint64(nSharePrefix))
	n += binary.PutUvarint(d.header[n:], uint64(keyLen-nSharePrefix))
	n += binary.PutUvarint(d.header[n:], uint64(valueLen))

	// data
	d.record.Write(d.header[:n])
	d.record.Write(key[nSharePrefix:])
	d.record.Write(value)

	d.prevKey = append(d.prevKey[:0], key...)
	d.nEntries++
}

func (d *Block) FlushBlockTo(dest io.Writer) (uint64, error) {

	defer d.clear()

	buf4 := make([]byte, 4)
	binary.LittleEndian.PutUint32(buf4, uint32(d.trailer.Len())/4)
	d.trailer.Write(buf4)

	n, err := dest.Write(d.compress())
	return uint64(n), err
}

func (d *Block) compress() []byte {

	d.record.Write(d.trailer.Bytes())
	n := snappy.MaxEncodedLen(d.record.Len())

	if n > len(d.compressionScratch) {
		d.compressionScratch = make([]byte, n+blockTrailerlen)
	}

	compressed := snappy.Encode(d.compressionScratch, d.record.Bytes())
	crc := utils.Checksum(compressed)

	size := len(compressed)
	compressed = compressed[:size+blockTrailerlen]

	binary.LittleEndian.PutUint32(compressed[size:], crc)

	return compressed
}

func (d *Block) clear() {

	d.nEntries = 0
	d.prevKey = d.prevKey[:0]
	d.record.Reset()
	d.trailer.Reset()
}

func (d *Block) Size() int {
	return d.record.Len() + d.trailer.Len() + 4
}

func NewBlock() *Block {

	return &Block{
		record:          bytes.NewBuffer(make([]byte, 0)),
		trailer:         bytes.NewBuffer(make([]byte, 0)),
		restartInterval: restartInterval,
	}

}
package desync

import "github.com/klauspost/compress/zstd"

// Create a writer that caches compressors.
// For this operation type we supply a nil Reader.
var encoder, _ = zstd.NewWriter(nil)

// Compress a buffer. 
// If you have a destination buffer, the allocation in the call can also be eliminated.
func Compress(b []byte) ([]byte, error) {
        return encoder.EncodeAll(b, make([]byte, 0, len(b))), nil
} 

var decoder, _ = zstd.NewReader(nil)

// Decompress a buffer. We don't supply a destination buffer,
// so it will be allocated by the decoder.
func Decompress(out, in []byte) ([]byte, error) {
        return decoder.DecodeAll(in, out)
} 

// import (
// 	"github.com/datadog/zstd"
// )

// // Compress a block using the only (currently) supported algorithm
// func Compress(b []byte) ([]byte, error) {
// 	return zstd.CompressLevel(nil, b, 3)
// }

// // Decompress a block using the only supported algorithm. If you already have
// // a buffer it can be passed into out and will be used. If out=nil, a buffer
// // will be allocated.
// func Decompress(out, in []byte) ([]byte, error) {
// 	return zstd.Decompress(out, in)
// }

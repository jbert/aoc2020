// Code generated by go2go; DO NOT EDIT.

//line /home/john/dev/jbert/aoc2020/src/aoc/bitmap/bitmap.go2:1
package bitmap

//line /home/john/dev/jbert/aoc2020/src/aoc/bitmap/bitmap.go2:1
import (
//line /home/john/dev/jbert/aoc2020/src/aoc/bitmap/bitmap.go2:1
	"fmt"
//line /home/john/dev/jbert/aoc2020/src/aoc/bitmap/bitmap.go2:1
	"testing"

//line /home/john/dev/jbert/aoc2020/src/aoc/bitmap/bitmap.go2:1
)

//line /home/john/dev/jbert/aoc2020/src/aoc/bitmap/bitmap.go2:5
type Bitmap struct {
	size int
	bits []byte
}

func New(size int) *Bitmap {
	b := Bitmap{size: size}
	b.bits = make([]byte, (size+7)/8)
	return &b
}

func (b *Bitmap) Set(n int) error {
	i, mask, err := b.indexAndMask(n)
	if err != nil {
		return err
	}
	b.bits[i] |= mask
	return nil
}

func (b *Bitmap) Clear(n int) error {
	i, mask, err := b.indexAndMask(n)
	if err != nil {
		return err
	}
	b.bits[i] &= ^mask
	return nil
}

func (b *Bitmap) Get(n int) (bool, error) {
	i, mask, err := b.indexAndMask(n)
	if err != nil {
		return false, err
	}
}

func (b *Bitmap) indexAndMask(n int) (int, byte, error) {
	if n <= 0 {
		return 0, 0, fmt.Errorf("Can't set negative index [%d]", n)
	}
	if n >= b.size {
		return 0, 0, fmt.Errorf("Can't set n >= size [%d >= %d]", n, b.size)
	}
	i := n / 8
	j := n % 8
	mask := byte(1) << j
	return i, mask, nil
}

//line /home/john/dev/jbert/aoc2020/src/aoc/bitmap/bitmap.go2:45
type Importable୦ int

//line /home/john/dev/jbert/aoc2020/src/aoc/bitmap/bitmap.go2:45
var _ = fmt.Errorf

//line /home/john/dev/jbert/aoc2020/src/aoc/bitmap/bitmap.go2:45
var _ = testing.AllocsPerRun
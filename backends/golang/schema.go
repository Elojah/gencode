package golang

import (
	"fmt"

	"github.com/elojah/gencode/schema"
)

type Walker struct {
	Needs     []string
	Offset    int
	IAdjusted bool
	Unsafe    bool
	DefStruct bool
	Ignore    map[string]struct{}
}

func (w *Walker) WalkSchema(s *schema.Schema, Package string) (parts *StringBuilder, err error) {
	parts = &StringBuilder{}
	parts.Append(fmt.Sprintf(`package %s

	import (
		"unsafe"
		"io"
		"time"
	)

	var (
		_ = unsafe.Sizeof(0)
		_ = io.ReadFull
		_ = time.Now()
	)
	`, Package))
	for _, st := range s.Structs {
		if _, ok := w.Ignore[st.Name]; ok {
			continue
		}
		p, err := w.WalkStruct(st)
		if err != nil {
			return nil, err
		}
		parts.Join(p)
	}
	return
}

package inputFile

import "io"

type mockFile struct {
	pos      int
	file     string
	didClose bool
}

func (m *mockFile) Read(buff []byte) (int, error) {
	if m.pos >= len(m.file) {
		return 0, io.EOF
	}

	count := copy(buff, m.file[m.pos:])
	m.pos += count
	return count, nil
}

func (m *mockFile) Close() error {
	m.didClose = true
	return nil
}

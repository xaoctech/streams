package streams

import "bufio"

func Scanner(scanner *bufio.Scanner) Stream[string] {
	return Stream[string]{&scannerStream{scanner}}
}

type scannerStream struct {
	scanner *bufio.Scanner
}

func (s scannerStream) Err() error {
	return s.scanner.Err()
}

func (s scannerStream) Next(value *string) bool {
	if s.scanner.Scan() {
		*value = s.scanner.Text()
		return true
	}
	return false
}

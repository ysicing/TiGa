// Copyright (c) 2023 ysicing(ysicing.me, ysicing@12306.work) All rights reserved.
// Use of this source code is covered by the following dual licenses:
// (1) Y PUBLIC LICENSE 1.0 (YPL 1.0)
// (2) Affero General Public License 3.0 (AGPL 3.0)
// License that can be found in the LICENSE file.

package log

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"math/rand"
	"sync"
)

var (
	globalItem  string
	globalMutex sync.Mutex
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// generateRandomString returns a random strin containing only letters
func generateRandomString(length int) string {
	b := make([]rune, length)
	for i := range b {
		//nolint:gosec
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func AcquireGlobalSilence() (string, error) {
	globalMutex.Lock()
	defer globalMutex.Unlock()

	if globalItem != "" {
		return "", fmt.Errorf("seems like there is already another terminal or question being asked currently")
	}

	globalItem = generateRandomString(12)
	return globalItem, nil
}

func ReleaseGlobalSilence(id string) {
	globalMutex.Lock()
	defer globalMutex.Unlock()

	if globalItem == id {
		globalItem = ""
	}
}

func NewScanner(r io.Reader) *bufio.Scanner {
	scanner := bufio.NewScanner(r)
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 1024*1024)
	scanner.Split(ScanLines)
	return scanner
}

// ScanLines is a split function for a Scanner that returns each line of
// text, stripped of any trailing end-of-line marker. The returned line may
// be empty. The end-of-line marker is one optional carriage return followed
// by one mandatory newline. In regular expression notation, it is `\r?\n`.
// The last non-empty line of input will be returned even if it has no
// newline.
func ScanLines(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := bytes.IndexByte(data, '\n'); i >= 0 {
		// We have a full newline-terminated line.
		return i + 1, data[0:i], nil
	}
	// If we're at EOF, we have a final, non-terminated line. Return it.
	if atEOF {
		return len(data), data, nil
	}
	// Request more data.
	return 0, nil, nil
}

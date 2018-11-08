package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"os"
	"unicode"
	"unicode/utf8"

	unidecode "github.com/mozillazg/go-unidecode"
	"github.com/pkg/term"
)

const (
	remotePort        = 4210
	ctrla      uint32 = 0x01
	ctrlc      uint32 = 0x03
)

func main() {
	if len(os.Args[1:]) != 1 {
		fmt.Fprintln(os.Stderr, "usage: reader <ip>")
		os.Exit(1)
	}

	remote := net.ParseIP(os.Args[1])
	if remote == nil {
		fmt.Fprintf(os.Stderr, "invalid ip specified: %s\n", os.Args[1])
		os.Exit(1)
	}

	t, err := term.Open("/dev/tty", term.RawMode)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error opening /dev/tty: %s\n", err)
		os.Exit(1)
	}
	defer t.Close()
	defer t.Restore()

	fmt.Print("lintening for input.\r\npress CTRL+A, CTRL+C to quit.\r\n\r\n")

	keys := make(chan uint32, 10)
	chars := make(chan uint8, 10)

	go read(t, keys)
	go convert(keys, chars)
	send(remote, chars)

	fmt.Print("program has quit\r\n")
}

func read(r io.Reader, keys chan<- uint32) {
	rawKey := make([]byte, 4)
	quitMode := false

	defer close(keys)

	for {
		n, err := r.Read(rawKey)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error reading input: %s\r\n", err)
			continue
		}

		key := binary.BigEndian.Uint32(rawKey) >> ((4 - uint(n)) * 8)

		if quitMode {
			if key == ctrlc {
				return
			}
			keys <- ctrla
			quitMode = false
		} else if key == ctrla {
			quitMode = true
			continue
		}

		keys <- key
	}
}

func convert(keys <-chan uint32, chars chan<- uint8) {
	defer close(chars)

	for k := range keys {
		if c, ok := subs[k]; ok {
			chars <- c
			continue
		}

		for _, r := range toASCII(k) {
			chars <- uint8(r)
		}
	}
}

func send(ip net.IP, chars <-chan uint8) {
	conn, _ := net.DialUDP("udp", nil, &net.UDPAddr{IP: ip, Port: remotePort})
	defer conn.Close()

	for c := range chars {
		fmt.Printf("%#x (%c) \r", c, c)
		_, err := conn.Write([]byte{c})
		if err != nil {
			fmt.Fprintf(os.Stderr, "\r\nerror sending character: %s\r\n", err)
		}
	}
}

func toASCII(key uint32) string {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, key)
	b = bytes.TrimLeft(b, "\x00")
	r, _ := utf8.DecodeRune(b)

	if r == utf8.RuneError || !unicode.IsPrint(r) {
		return ""
	}

	return unidecode.Unidecode(string(r))
}

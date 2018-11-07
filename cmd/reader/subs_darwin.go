package main

var subs = map[uint32]byte{
	0x1b5b41:   ArduinoKeyUpArrow,
	0x1b5b42:   ArduinoKeyDownArrow,
	0x1b5b44:   ArduinoKeyLeftArrow,
	0x1b5b43:   ArduinoKeyRightArrow,
	0x7f:       ArduinoKeyBackspace,
	0x09:       ArduinoKeyTab,
	0x0d:       ArduinoKeyReturn,
	0x1b:       ArduinoKeyEsc,
	0x1b5b337e: ArduinoKeyDelete,
	0x1b4f50:   ArduinoKeyF1,
	0x1b4f51:   ArduinoKeyF2,
	0x1b4f52:   ArduinoKeyF3,
	0x1b4f53:   ArduinoKeyF4,
	0x1b5b3135: ArduinoKeyF5,
	0x1b5b3137: ArduinoKeyF6,
	0x1b5b3138: ArduinoKeyF7,
	0x1b5b3139: ArduinoKeyF8,
	0x1b5b3230: ArduinoKeyF9,
	0x1b5b3231: ArduinoKeyF10,
}

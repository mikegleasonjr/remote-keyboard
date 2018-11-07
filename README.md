# remote-keyboard

I will document more in the coming days.

* The `Go` program in `cmd/reader` reads key presses and converts the characters to ascii. The characters are sent via `UDP` to the `ESP8266`.
* The `ESP8266` receives the ascii characters. and sends them to the `Pro Micro` via serial.
* The `Pro Micro` reads from the serial port the characters and sends them to the computer as an USB keyboard.

See https://www.youtube.com/watch?v=F7WRIw0N6Gg

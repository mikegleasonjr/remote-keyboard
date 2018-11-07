#include <Keyboard.h>

void setup() {
  delay(5000);
  Serial1.begin(115200);
}

void loop() {
  if (Serial1.available()) {
    Keyboard.write(Serial1.read());
  }
}

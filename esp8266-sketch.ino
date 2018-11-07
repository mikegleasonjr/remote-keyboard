#include <ESP8266WiFi.h>
#include <WiFiUdp.h>
#include <ArduinoOTA.h>

const char*         SSID     = "**********";
const char*         PASSWORD = "**********";
const unsigned int  UDP_PORT = 4210;

WiFiUDP Udp;
unsigned long lastLit = 0;

void setup() {
  pinMode(LED_BUILTIN, OUTPUT);
  digitalWrite(LED_BUILTIN, LOW);

  WiFi.mode(WIFI_STA);
  WiFi.begin(SSID, PASSWORD);
  while (WiFi.waitForConnectResult() != WL_CONNECTED) {
    delay(4000);
    digitalWrite(LED_BUILTIN, HIGH);
    delay(1000);
    ESP.restart();
  }

  ArduinoOTA.begin();
  Udp.begin(UDP_PORT);
  Serial.begin(115200);

  digitalWrite(LED_BUILTIN, HIGH);
}

void loop() {
  ArduinoOTA.handle();

  if (Udp.parsePacket()) {
    digitalWrite(LED_BUILTIN, LOW);
    lastLit = millis();
    Serial.print((char)Udp.read());
  }

  if (lastLit != 0 && (millis() - lastLit) >= 35) {
    digitalWrite(LED_BUILTIN, HIGH);
    lastLit = 0;
  }
}

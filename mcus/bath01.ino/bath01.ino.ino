#include <ESP8266WiFi.h>
#include <IRremoteESP8266.h>
#include <IRsend.h>
#include <ESP8266WebServer.h>

ESP8266WebServer server(80);

void setup() {
  // put your setup code here, to run once:
  Serial.begin(115200);
  Serial.println("begin");
  connectWiFi();

  server.on("/light/value", handleLightValue);
  server.on("/human/is_in", handleHumanValue);
  server.on("/ir/send", handleSendIRCmd);
  server.begin();
}

void handleSendIRCmd() {
  if(!server.hasArg("cmd")) {
    server.send(500, "text/plain", "need param cmd");
    return;
  }
  String value = server.arg("cmd");
  long cmd = value.toInt();
  if(cmd <= 0) {
    server.send(400, "text/plain", "param cmd must be active int, got" + value);
    return;
  }
  sendIRCmd(cmd);
  server.send(200, "application/json", "{\"success\": true}");
}

void handleHumanValue() {
  int isHumanInValue = isHumanIn();
  String isHumanInStr = isHumanInValue > 0 ? "true" : "false";
  server.send(200, "application/json", "{\"is_in\": " + isHumanInStr + "}");
}

void handleLightValue() {
  int lightValue = getLightValue();
  server.send(200, "application/json", "{\"value\":" + String(lightValue) + "}");
}

void connectWiFi() {
  Serial.println("connect WiFi");

  WiFi.begin("daxigua", "Wuqi4321");

  Serial.print("Connecting");
  int i = 0;
  while (WiFi.status() != WL_CONNECTED && i < 20)
  {
    i++;
    delay(500);
    Serial.print(".");
  }
  Serial.println();

  Serial.print("Connected, IP address: ");
  Serial.println(WiFi.localIP());
}

void loop() {
  server.handleClient();
}

int PIN_IR = 7;
IRsend irsend(PIN_IR);

void sendIRCmd(long cmd) {
  Serial.println("send IR cmd:");
  Serial.println(String(cmd, HEX));
  irsend.sendNEC(0x1FE40BFL, 32);
}

int PIN_HUMAN = 2;

int isHumanIn() {
  int human = digitalRead(PIN_HUMAN);
  Serial.println("human value:");
  Serial.println(human);
  return human;
}

int PIN_LIGHT = 0;
int getLightValue() {
  int light = analogRead(PIN_LIGHT);
  Serial.println("light value:");
  Serial.println(light);
  return light;
}


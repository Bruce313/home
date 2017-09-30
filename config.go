package main

import "os"
import "github.com/naoina/toml"

type LightSensorConf struct {
	RoomId uint8 `toml:room_id`
	ID uint8 `toml:id`
	HttpURL string `toml:http_url`
	LightValuePath string `toml:light_value_path`
}

type RoomConf struct {
	Id uint8	`toml:id`
	Name string `toml:name`
}

type IRSender struct {
	RoomId uint8 `toml:room_id`
	ID uint8 `toml:id`
	HttpURL string `toml:http_url`
	CmdSendPath string `toml:cmd_send_path`
}

type MotionSensorConf struct {
	RoomId uint8 `toml:room_id`
	ID uint8 `toml:id`
}

type Config struct {
	Rooms []RoomConf `toml:rooms`
	LightSensors []LightSensorConf `toml:light_sensors`
	IRSenders []IRSender `toml:ir_senders`
	MotionSensors []MotionSensorConf `toml:motion_sensors`
}

func ParseConfig(confPath string) (conf Config, err error)  {
	f, err := os.Open(confPath)
	if err != nil {
		return
	}
	defer f.Close()
	if err = toml.NewDecoder(f).Decode(&conf); err != nil {
		return
	}
	return
}
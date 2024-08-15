package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/google/uuid"
)

type Owntracks struct {
	Type       string  `json:"_type"`
	BSSID      string  `json:"BSSID"`
	SSID       string  `json:"SSID"`
	ID         string  `json:"_id"`
	Acc        uint64  `json:"acc"`
	Alt        uint64  `json:"alt"`
	Batt       uint64  `json:"batt"`
	Bs         uint64  `json:"bs"`
	Conn       string  `json:"conn"`
	Created_At string  `json:"created_at"`
	Lat        float64 `json:"lat"`
	Lon        float64 `json:"lon"`
	M          uint64  `json:"m"`
	T          string  `json:"t"`
	Tid        string  `json:"tid"`
	Tst        uint64  `json:"tst"`
	Vac        uint64  `json:"vac"`
	Vel        uint64  `json:"vel"`
}

// General
// accuracy
// altitude
// deviceBattery
// batteryStatus
// latitude
// longitude
// trigger
// trackerId
// timestamp
// velocity
// internetConnectivity
// createdAt

// Android
// _id

// IoS
// courseOverGround
// radius
// verticalAccuracy
// barometricPressure
// pointOfInterestName
// tag
// ssid
// bssid
// monitoringMode

type Pd07_dcu struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Timestamp uint64  `json:"timestamp"`
}

// InfluxDB ->
// Organization -> Buckets (deviceType: IC, H2, EV) -> Measurements (rmc, gga, powerTrain, dynamics, peripherals)

type GPS struct {
	GpsAltitude  uint64  `json:"gps_altitude"`
	GpsLatitude  float64 `json:"gps_latitude"`
	GpsLongitude float64 `json:"gps_longitude"`
	GpsSpeed     float64 `json:"gps_speed"`
}

type IcPowerTrain struct {
	EngineRpm                    uint64  `json:"engine_rpm"`
	Gear                         uint64  `json:"gear"`
	EngineOilPressure            float64 `json:"engine_oil_pressure"`
	EngineTemperature            float64 `json:"engine_temperature"`
	EngineIntakeAirTemperature   float64 `json:"engine_intake_air_temperature"`
	EngineEcuTemperature         float64 `json:"engine_ecu_temperature"`
	ManifoldAirPressure          float64 `json:"manifold_air_pressure"`
	FuelPressure                 float64 `json:"fuel_pressure"`
	ExhaustTemperatureCylinder_1 float64 `json:"exhaust_temperature_cylinder_1"`
	ExhaustTemperatureCylinder_2 float64 `json:"exhaust_temperature_cylinder_2"`
	ExhaustTemperatureCylinder_3 float64 `json:"exhaust_temperature_cylinder_3"`
	Lambda_1                     float64 `json:"lambda_1"`
	ThrottlePosition             float64 `json:"throttle_position"`
	FuelUsed                     float64 `json:"fuel_used"`
}

type H2PowerTrain struct {
	EngineRpm                    uint64  `json:"engine_rpm"`
	Gear                         uint64  `json:"gear"`
	EngineOilPressure            float64 `json:"engine_oil_pressure"`
	EngineTemperature            float64 `json:"engine_temperature"`
	EngineIntakeAirTemperature   float64 `json:"engine_intake_air_temperature"`
	EngineEcuTemperature         float64 `json:"engine_ecu_temperature"`
	ManifoldAirPressure          float64 `json:"manifold_air_pressure"`
	FuelPressure                 float64 `json:"fuel_pressure"`
	ExhaustTemperatureCylinder_1 float64 `json:"exhaust_temperature_cylinder_1"`
	ExhaustTemperatureCylinder_2 float64 `json:"exhaust_temperature_cylinder_2"`
	ExhaustTemperatureCylinder_3 float64 `json:"exhaust_temperature_cylinder_3"`
	Lambda_1                     float64 `json:"lambda_1"`
	ThrottlePosition             float64 `json:"throttle_position"`
	FuelUsed                     float64 `json:"fuel_used"`
}

type EvPowerTrain struct {
	EngineRpm                    uint64  `json:"engine_rpm"`
	Gear                         uint64  `json:"gear"`
	EngineOilPressure            float64 `json:"engine_oil_pressure"`
	EngineTemperature            float64 `json:"engine_temperature"`
	EngineIntakeAirTemperature   float64 `json:"engine_intake_air_temperature"`
	EngineEcuTemperature         float64 `json:"engine_ecu_temperature"`
	ManifoldAirPressure          float64 `json:"manifold_air_pressure"`
	FuelPressure                 float64 `json:"fuel_pressure"`
	ExhaustTemperatureCylinder_1 float64 `json:"exhaust_temperature_cylinder_1"`
	ExhaustTemperatureCylinder_2 float64 `json:"exhaust_temperature_cylinder_2"`
	ExhaustTemperatureCylinder_3 float64 `json:"exhaust_temperature_cylinder_3"`
	Lambda_1                     float64 `json:"lambda_1"`
	ThrottlePosition             float64 `json:"throttle_position"`
	FuelUsed                     float64 `json:"fuel_used"`
}

type Dynamics struct {
	BrakePressureFront float64 `json:"brake_pressure_front"`
	BrakePressureRear  float64 `json:"brake_pressure_rear"`
	Acc_X              float64 `json:"acc_x"`
	Acc_Y              float64 `json:"axx_y"`
	Acc_Z              float64 `json:"acc_z"`
	Beacon             uint64  `json:"beacon"`
}

type Peripherals struct {
	DataLoggerVoltage     float64 `json:"sdl3_battery_voltage"`
	DataLoggerTemperature float64 `json:"sdl3_temperature"`
}
type IC struct {
	Dynamics     Dynamics
	GPS          GPS
	Peripherals  Peripherals
	IcPowerTrain IcPowerTrain
}

type H2 struct {
	Dynamics     Dynamics
	GPS          GPS
	Peripherals  Peripherals
	H2PowerTrain H2PowerTrain
}

type EV struct {
	Dynamics     Dynamics
	GPS          GPS
	Peripherals  Peripherals
	EvPowerTrain EvPowerTrain
}

// return influx line protocol
// measurement,tags fields timestamp
func parseOwntracks(deviceId string, direction string, etc string, message string) string {
	var sb strings.Builder
	var owntracks Owntracks

	if message == "" {
		// return message, errors.New("empty message to parse")
		return message
	}
	json.Unmarshal([]byte(message), &owntracks)

	if direction == "rx" {
		// message to struct
		// Measurement
		sb.WriteString("owntracks.Type,")

		// Tags
		sb.WriteString(",deviceId=")
		sb.WriteString(deviceId)
		sb.WriteString(",direction=")
		sb.WriteString(direction)
		sb.WriteString(",etc=")
		sb.WriteString(etc)
		// sb.WriteString(",ssid=")
		// sb.WriteString(owntracks.SSID)
		// sb.WriteString(",bssid=")
		// sb.WriteString(owntracks.BSSID)
		// sb.WriteString(",id=")
		// sb.WriteString(owntracks.ID)
		sb.WriteString(",trigger=")
		sb.WriteString(owntracks.T)
		sb.WriteString(",trackerId=")
		sb.WriteString(owntracks.Tid)
		sb.WriteString(",internetConnectivity=")
		sb.WriteString(owntracks.Conn)

		// Fields
		sb.WriteString(" accuracy=")
		sb.WriteString(strconv.FormatUint(uint64(owntracks.Acc), 10))
		sb.WriteString(",verticalAccuracy=")
		sb.WriteString(strconv.FormatUint(uint64(owntracks.Vac), 10))
		sb.WriteString(",deviceBatteryLevel=")
		sb.WriteString(strconv.FormatUint(uint64(owntracks.Batt), 10))
		sb.WriteString(",deviceBatteryStatus=")
		sb.WriteString(strconv.FormatUint(uint64(owntracks.Bs), 10))
		sb.WriteString(",altitude=")
		sb.WriteString(strconv.FormatUint(uint64(owntracks.Alt), 10))
		sb.WriteString(",latitude=")
		sb.WriteString(strconv.FormatFloat(owntracks.Lat, 'f', -1, 64))
		sb.WriteString(",longitude=")
		sb.WriteString(strconv.FormatFloat(owntracks.Lon, 'f', -1, 64))
		sb.WriteString(",velocity=")
		sb.WriteString(strconv.FormatUint(uint64(owntracks.Vel), 10))

		// if etc == "android" {

		// }
		// if etc == "ios" {

		// }

		// Timestamp_ms
		sb.WriteString(" timestamp=")
		sb.WriteString(strconv.FormatUint(uint64(owntracks.Tst), 10))
	}
	return sb.String()
}

func parseIc(deviceId string, direction string, etc string, message string) string {
	var s string

	return s
}

func parseH2(deviceId string, direction string, etc string, message string) string {
	var s string

	return s
}

func parseEv(deviceId string, direction string, etc string, message string) string {
	var s string

	return s
}

func parseEvse(deviceId string, direction string, etc string, message string) string {
	var s string

	return s
}

func parseHealthPack(deviceId string, direction string, etc string, message string) string {
	var s string

	return s
}

func parseLns(deviceId string, direction string, etc string, message string) string {
	var s string

	return s
}

func parseHttp(deviceId string, direction string, etc string, message string) string {
	var s string

	return s
}

// TODO:RECEIVE MQTT MESSAGE AND PARSE Owntracks TO LNS
// func parseOwntracks(message string) string {
// 	if message == "" {
// 		// return message, errors.New("empty message to parse")
// 		return message
// 	}

// 	var om Owntracks // owntracks message
// 	var sb strings.Builder
// 	// var pd string // parsed data

// 	json.Unmarshal([]byte(message), &om)

// TODO: JUST PARSE DATA TO LNS STRUCTURE
// om.Type =
// om.BSSID =
// om.SSID =
// om.ID =
// om.Acc =
// om.Alt =
// om.Batt =
// om.Bs =
// om.Conn =
// om.Created_At =
// pm.Latitude = om.Lat
// pm.Longitude = om.Lon
// om.M =
// om.P =
// om.Tid =
// om.Tst =
// om.Vac =
// om.Vel

// Show NodeName before Panic
// fmt.Printf("\n pm.Latitude: %v\n", om.Lat)
// fmt.Printf("\n pm.Longitude: %v\n", om.Lon)
// hex, _ := b64ToHex(lnsImt.Data)
// fmt.Printf("Hex: %v\n", hex)

// Set measurement
// sb.WriteString("owntracks,")
// deviceType := strings.Split(lnsImt.NodeName, "_")
// sb.WriteString(deviceType[0])
// sb.WriteString(",")

// // Set tags
// sb.WriteString("applicationName=")
// sb.WriteString(lnsImt.ApplicationName)
// sb.WriteString(",")

// sb.WriteString("applicationID=")
// sb.WriteString(lnsImt.ApplicationID)
// sb.WriteString(",")

// sb.WriteString("nodeName=")
// sb.WriteString(lnsImt.NodeName)
// sb.WriteString(",")

// sb.WriteString("devEUI=")
// sb.WriteString(lnsImt.DevEUI)
// sb.WriteString(",")

// Just for the first gateway
// for i, v := range lnsImt.RxInfo {

// 	if i == 0 {
// 		lns.RxInfo_0_mac = v.Mac            // string
// 		lns.RxInfo_0_time = v.Time.String() // string
// 		lns.RxInfo_0_rssi = v.Rssi          // int64
// 		lns.RxInfo_0_snr = v.LoRaSNR        // float64
// 		lns.RxInfo_0_lat = v.Latitude       // float64
// 		lns.RxInfo_0_lon = v.Longitude      // float64
// 		lns.RxInfo_0_alt = v.Altitude       // int64
// sb.WriteString("lora_rxInfo_mac_")
// sb.WriteString(strconv.FormatUint(uint64(i), 10))
// sb.WriteString("=")
// sb.WriteString(v.Mac)
// sb.WriteString(",")

// sb.WriteString("lora_rxInfo_name_")
// sb.WriteString(strconv.FormatUint(uint64(i), 10))
// sb.WriteString("=")
// sb.WriteString(v.Name)
// sb.WriteString(",")

// sb.WriteString("rxInfo_time_")
// sb.WriteString(strconv.FormatUint(uint64(i), 10))
// sb.WriteString("=")
// sb.WriteString(v.Time.String())
// sb.WriteString(",")

// sb.WriteString("lora_rxInfo_rssi_")
// sb.WriteString(strconv.FormatUint(uint64(i), 10))
// sb.WriteString("=")
// sb.WriteString(strconv.FormatInt(v.Rssi, 10))
// sb.WriteString(",")

// sb.WriteString("lora_rxInfo_loRaSNR_")
// sb.WriteString(strconv.FormatUint(uint64(i), 10))
// sb.WriteString("=")
// sb.WriteString(strconv.FormatInt(int64(v.LoRaSNR), 10))
// sb.WriteString(",")

// sb.WriteString("lora_rxInfo_latitude_")
// sb.WriteString(strconv.FormatUint(uint64(i), 10))
// sb.WriteString("=")
// sb.WriteString(strconv.FormatFloat(v.Latitude, 'f', -1, 64))
// sb.WriteString(",")

// sb.WriteString("lora_rxInfo_longitude_")
// sb.WriteString(strconv.FormatUint(uint64(i), 10))
// sb.WriteString("=")
// sb.WriteString(strconv.FormatFloat(v.Longitude, 'f', -1, 64))
// sb.WriteString(",")

// sb.WriteString("lora_rxInfo_altitude_")
// sb.WriteString(strconv.FormatUint(uint64(i), 10))
// sb.WriteString("=")
// sb.WriteString(strconv.FormatInt(int64(v.Altitude), 10))
// sb.WriteString(",")
// 	}
// }

// sb.WriteString("lora_txInfo_dataRate_modulation=")
// sb.WriteString(lnsImt.TxInfo.DataRate.Modulation)
// sb.WriteString(",")

// sb.WriteString("lora_txInfo_dataRate_bandwidth=")
// sb.WriteString(strconv.FormatUint(lnsImt.TxInfo.DataRate.Bandwidth, 10))
// sb.WriteString(",")

// sb.WriteString("lora_txInfo_adr=")
// sb.WriteString(strconv.FormatBool(lnsImt.TxInfo.Adr))
// sb.WriteString(",")

// sb.WriteString("lora_txInfo_codeRate=")
// sb.WriteString(lnsImt.TxInfo.CodeRate)
// sb.WriteString(",")

// sb.WriteString("lora_fPort=")
// sb.WriteString(strconv.FormatUint(lnsImt.FPort, 10))
// sb.WriteString("")

// Set fields
// sb.WriteString(" ")

// for i, v := range lnsImt.RxInfo {
// if i == 0 {
// sb.WriteString("rxInfo_time_")
// sb.WriteString(strconv.FormatUint(uint64(i), 10))
// sb.WriteString("=")
// sb.WriteString(v.Time.String())
// sb.WriteString(",")

// sb.WriteString("lora_rxInfo_rssi_")
// sb.WriteString(strconv.FormatUint(uint64(i), 10))
// sb.WriteString("=")
// sb.WriteString(strconv.FormatInt(v.Rssi, 10))
// sb.WriteString(",")

// sb.WriteString("lora_rxInfo_loRaSNR_")
// sb.WriteString(strconv.FormatUint(uint64(i), 10))
// sb.WriteString("=")
// sb.WriteString(strconv.FormatInt(int64(v.LoRaSNR), 10))
// sb.WriteString(",")

// sb.WriteString("lora_rxInfo_latitude_")
// sb.WriteString(strconv.FormatUint(uint64(i), 10))
// sb.WriteString("=")
// sb.WriteString(strconv.FormatFloat(v.Latitude, 'f', -1, 64))
// sb.WriteString(",")

// sb.WriteString("lora_rxInfo_longitude_")
// sb.WriteString(strconv.FormatUint(uint64(i), 10))
// sb.WriteString("=")
// sb.WriteString(strconv.FormatFloat(v.Longitude, 'f', -1, 64))
// sb.WriteString(",")

// sb.WriteString("lora_rxInfo_altitude_")
// sb.WriteString(strconv.FormatUint(uint64(i), 10))
// sb.WriteString("=")
// sb.WriteString(strconv.FormatInt(int64(v.Altitude), 10))
// sb.WriteString(",")
// }
// }

// sb.WriteString("lora_txInfo_frequency=")
// sb.WriteString(strconv.FormatUint(lnsImt.TxInfo.Frequency, 10))
// sb.WriteString(",")

// sb.WriteString("lora_txInfo_dataRate_spreadFactor=")
// sb.WriteString(strconv.FormatUint(lnsImt.TxInfo.DataRate.SpreadFactor, 10))
// sb.WriteString(",")

// Decode and Parse data
// data := lnsImt.Data
// b, err := b64ToByte(data)
// if err != nil {
// 	fmt.Print(lnsImt.Data)
// 	log.Panic(err)
// }

// switch lnsImt.FPort {
// case 100:
// 	pd = imtIotProtocolParser(b)

// 	// case 200:
// 	// 	jsonProtocolParser()

// 	// case 3:
// 	// 	khompProtocolParser()

// 	// case 4:
// 	// 	khompProtocolParser()
// }
// sb.WriteString(pd)

// sb.WriteString("fCnt=")
// sb.WriteString(strconv.FormatUint(lnsImt.FCnt, 10))

// Set time
// sb.WriteString(" ")
// t := time.Now().UnixNano()
// sb.WriteString(strconv.FormatInt(t, 10))

// pd, err := json.Marshal(lns)
// if err != nil {
// 	fmt.Println(err)
// 	return "No parsed"
// }
// // fmt.Println(string(pd))
// // return sb.String(), nil
// return string(pd)
// }

func main() {
	id := uuid.New().String()

	// measurements format
	var location Location

	// MqttSubscriberClient
	var sbMqttSubClientId strings.Builder
	sbMqttSubClientId.WriteString("parse-owntracks-sub-")
	sbMqttSubClientId.WriteString(id)

	// MqttSubscriberTopic
	var sbMqttSubTopic strings.Builder
	sbMqttSubTopic.WriteString("OpenDataTelemetry/")
	sbMqttSubTopic.WriteString(ORGANIZATION)
	sbMqttSubTopic.WriteString("/")
	sbMqttSubTopic.WriteString(DEVICETYPE)
	sbMqttSubTopic.WriteString("/+/")
	sbMqttSubTopic.WriteString(DIRECTION)
	sbMqttSubTopic.WriteString("/+")

	// KafkaProducerClient
	var sbKafkaProdClientId strings.Builder
	sbKafkaProdClientId.WriteString("parse-owntracks-prod-")
	sbKafkaProdClientId.WriteString(id)

	// KafkaProducerClient
	var sbKafkaProdTopic strings.Builder
	sbKafkaProdTopic.WriteString(ORGANIZATION)
	sbKafkaProdTopic.WriteString(".")
	sbKafkaProdTopic.WriteString(BUCKET)

	// var lns LNS
	// var smartlight SmartLight
	// var decodedData Decoded
	// var decodedDataJson []byte

	// var pd string // parsed data

	// var lpd string
	// var ppd string

	// var sbKafkaClientId strings.Builder
	// var sbKafkaProducerTopic strings.Builder
	// sbKafkaClientId.WriteString("parse-owntracks-mqtt-producer-")
	// sbKafkaClientId.WriteString(id)

	// MQTT
	mqttSubBroker := "mqtt://mqtt.maua.br:1883"
	mqttSubClientId := sbMqttSubClientId.String()
	mqttSubUser := "public"
	mqttSubPassword := "public"
	mqttSubQos := 0

	mqttSubOpts := MQTT.NewClientOptions()
	mqttSubOpts.AddBroker(mqttSubBroker)
	mqttSubOpts.SetClientID(mqttSubClientId)
	mqttSubOpts.SetUsername(mqttSubUser)
	mqttSubOpts.SetPassword(mqttSubPassword)

	// pBroker := "mqtt://mqtt.maua.br:1883"
	// pClientId := sbMqttPubClientId.String()
	// pUser := "public"
	// pPassword := "public"
	// pQos := 0

	// pOpts := MQTT.NewClientOptions()
	// pOpts.AddBroker(pBroker)
	// pOpts.SetClientID(pClientId)
	// pOpts.SetUsername(pUser)
	// pOpts.SetPassword(pPassword)

	c := make(chan [2]string)

	mqttSubOpts.SetDefaultPublishHandler(func(mqttClient MQTT.Client, msg MQTT.Message) {
		c <- [2]string{msg.Topic(), string(msg.Payload())}
	})

	mqttSubClient := MQTT.NewClient(mqttSubOpts)
	if token := mqttSubClient.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	} else {
		fmt.Printf("Connected to %s\n", mqttSubBroker)
	}

	// pClient := MQTT.NewClient(pOpts)
	// if token := pClient.Connect(); token.Wait() && token.Error() != nil {
	// 	panic(token.Error())
	// } else {
	// 	fmt.Printf("Connected to %s\n", pBroker)
	// }

	if token := mqttSubClient.Subscribe(sbMqttSubTopic.String(), byte(mqttSubQos), nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	// KAFKA
	// p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "my-cluster-kafka-bootstrap.test-kafka.svc.cluster.local"})
	// if err != nil {
	// 	panic(err)
	// }
	// defer p.Close()

	// Delivery report handler for produced messages
	// go func() {
	// 	for e := range p.Events() {
	// 		switch ev := e.(type) {
	// 		case *kafka.Message:
	// 			if ev.TopicPartition.Error != nil {
	// 				fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
	// 			} else {
	// 				fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
	// 			}
	// 		}
	// 	}
	// }()

	// MQTT -> KAFKA
	for {
		// 1. Input
		incoming := <-c

		// 2. Process
		// 2.1. Process Topic
		s := strings.Split(incoming[0], "/")
		deviceId := s[3]
		etc := s[5]

		// influxLineProtocol
		var ilp string

		// 2.2. Process Message
		switch organization {
		case "FSAELive": // Owntracks || IC || H2 || EV
			switch deviceType {
			case "Owntracks":
				ilp = parseOwntracks(deviceId, direction, etc, incoming[1])

			case "IC":
				ilp = parseIc(deviceId, direction, etc, incoming[1])

			case "H2":
				ilp = parseH2(deviceId, direction, etc, incoming[1])

			case "EV":
				ilp = parseEv(deviceId, direction, etc, incoming[1])

			default:
			}

		case "IMT": // LNS || EVSE
			switch deviceType {
			case "LNS":
				ilp = parseLns(deviceId, direction, etc, incoming[1])

			case "EVSE":
				ilp = parseEvse(deviceId, direction, etc, incoming[1])

			default:
			}

		case "SaoRafael": // HealthPack
			switch deviceType {
			case "HealthPack":
				ilp = parseHealthPack(deviceId, direction, etc, incoming[1])

			default:
			}

		case "Hilight": // LNS || HTTP
			switch deviceType {
			case "LNS":
				ilp = parseLns(deviceId, direction, etc, incoming[1])

			case "HTTP":
				ilp = parseHttp(deviceId, direction, etc, incoming[1])

			default:
			}

		default: //
		}

		// sbKafkaProducerTopic.Reset()
		// sbKafkaProducerTopic.WriteString(organization)
		// sbKafkaProducerTopic.WriteString(".")
		// sbKafkaProducerTopic.WriteString("SmartCampusMaua")

		// fmt.Printf("Organization: %s\n", organization)
		// // fmt.Printf("KafkaTopic: %s\n", sbKafkaProducerTopic.String())
		// fmt.Printf("deviceType: %s\n", deviceType)
		// fmt.Printf("DeviceId: %s\n", deviceId)
		// fmt.Printf("Message: %s\n", incoming[1])
		// fmt.Printf("ETC: %s\n", etc)
		// fmt.Printf("\n")

		// return influx line protocol
		// measurement,tags fields timestamp
		fmt.Printf("InfluxLineProtocol: %s\n", ilp)

		// Data TAG_KEYS shall be given by the application API using a Redis database. So the correct information shall be stored alongside with sensor
		// MAP deviceId vs deviceType to understand what decode really means for each one then write to kafka after decoded
		// Parse (IMT vs ATC) -> Map (deviceId vs deviceType) -> decode payload by port (0dCCCCCC)
		// -> measurement=,tag_key1=,tag_key2=, field_key_1= timestamp_ms
		// smartlight, raw=,temperature=,humidity=,lux=,movement=,box_battery=,board_voltage= timestamp_ms
		// gaugepressure, raw=pressure_in=,pressure_out=,board_voltage= timestamp_ms
		// watertanklevel, raw=distance=,pressure_out=,board_voltage= timestamp_ms
		// healthpack_alarm, raw=emergency= timestamp_ms
		// healthpack_tracking, raw=latitude=,longitude= timestamp_ms
		// evse_startTransaction, raw= timestamp_ms
		// evse_heartbeat, raw= timestamp_ms

		// switch etc {
		// case "lns_imt":
		// 	lpd = parseLnsImt(incoming[1]) // lns parsed data
		// 	// if err != nil {
		// 	// 	log.Panic(err)
		// 	// }
		// 	fmt.Printf("PARSED DATA: %s", lpd)
		// case "lns_atc":
		// 	// parseLnsAtc()
		// case "android":
		// 	lpd = parseLnsImt(incoming[1]) // lns parsed data
		// 	// if err != nil {
		// 	// 	log.Panic(err)
		// 	// }
		// 	fmt.Printf("PARSED DATA: %s", lpd)
		// default:
		// }
		// json.Unmarshal([]byte(lpd), &lns)
		// TODO: DECODE DATA ACCORDING MEASUREMENT THEN RETURN A JSON STRING
		// dd := decodeData(lns.Data) // decoded data
		// dd: {"01":24.1,"02":95,"03_0":0,"03_1":0,"04":0,"05":0,"06":0,"07":0,"08":0,"09":0,"0A_0":0,"0A_1":0,"0B":518744,"0C":3.294,"0D_0":4095,"0D_1":0,"0D_2":0,"0D_3":0,"0E_0":0,"0E_1":0,"10":0,"11":0,"12":0,"13":0}
		// if err != nil {
		// }
		// fmt.Printf("\n\n### PROTOCOL PARSED: %s", dd)

		// var decodedDataJson []byte

		switch deviceType {
		// case "SmartLight":
		// 	json.Unmarshal([]byte(dd), &decodedData)

		// 	smartlight.Temperature = float64(decodedData.X_01)
		// 	smartlight.Humidity = float64(decodedData.X_02)
		// 	smartlight.Lux = float64(decodedData.X_03_0)
		// 	smartlight.Movement = decodedData.X_0D_0
		// 	smartlight.Battery = float64(decodedData.X_0D_1)
		// 	smartlight.BoardVoltage = float64(decodedData.X_0C)

		// 	var err any
		// 	decodedDataJson, err = json.Marshal(smartlight)
		// 	if err != nil {
		// 		fmt.Println(err)
		// 	}
		case "Owntracks":
			measurement = "owntracks"
			json.Unmarshal([]byte(incoming[1]), &om)
		default:
		}

		// fmt.Printf("\n\n###ecodedDataJson:%s", string(decodedDataJson))
		fmt.Printf("\n\n###ecodedDataJson:%s", string(om.BSSID))
		fmt.Printf("\n\n###measurement:%s", measurement)

		// TODO: JSON TO INFLUX

		// FINISH

		// fmt.Println("topic:" + incoming[0])

		// sbKafkaProducerTopic.Reset()
		// sbKafkaProducerTopic.WriteString("lns_imt/")
		// sbKafkaProducerTopic.WriteString(devEUI)
		// sbKafkaProducerTopic.WriteString("/rx")

		// lp, err := decode.LoraImt(incoming[1])
		// if err != nil {
		// 	log.Panic(err)
		// }
		// fmt.Println(lp)

		// topic := "SmartCampusMaua.smartcampusmaua"

		// pClient.Publish(sbPubTopic.String(), byte(pQos), false, incoming[1])

		// p.Produce(&kafka.Message{
		// 	TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		// 	Value:          []byte(lp),
		// }, nil)

		// p.Flush(15 * 1000)

		sbPubTopic.Reset()
		sbPubTopic.WriteString("FSAELive/Owntracks/")
		sbPubTopic.WriteString(deviceId)
		sbPubTopic.WriteString("/rx")
		// fmt.Printf("RECEIVED TOPIC: %s MESSAGE: %s\n", incoming[0], incoming[1])
		token := pClient.Publish(sbPubTopic.String(), byte(pQos), false, incoming[1])
		token.Wait()
	}
}

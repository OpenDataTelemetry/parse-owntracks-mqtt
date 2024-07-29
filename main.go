package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/OpenDataTelemetry/decode"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/google/uuid"
)

type Owntracks struct {
	Type       string  `json:"_type"`
	BSSID      string  `json:"BSSID"`
	SSID       string  `json:"SSID"`
	ID         string  `json:"_id"`
	Acc        float64 `json:"acc"`
	Alt        float64 `json:"alt"`
	Batt       float64 `json:"batt"`
	Bs         string  `json:"bs"`
	Conn       string  `json:"conn"`
	Created_At string  `json:"created_at"`
	Lat        float64 `json:"lat"`
	Lon        float64 `json:"lon"`
	M          string  `json:"m"`
	T          string  `json:"t"`
	Tid        string  `json:"tid"`
	Tst        string  `json:"tst"`
	Vac        string  `json:"vac"`
	Vel        float64 `json:"vel"`
}
type Pd07_dcu struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Timestamp uint64  `json:"timestamp"`
}

// TODO:RECEIVE MQTT MESSAGE AND PARSE LNS_IMT TO LNS
func parseOwntracks(message string) string {
	if message == "" {
		// return message, errors.New("empty message to parse")
		return message
	}

	var om Owntracks // owntracks message
	var pm Pd07_dcu  // parsed message
	var sb strings.Builder
	// var pd string // parsed data

	json.Unmarshal([]byte(message), &om)

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
	pm.Latitude = om.Lat
	pm.Longitude = om.Lon
	// om.M =
	// om.T =
	// om.Tid =
	// om.Tst =
	// om.Vac =
	// om.Vel

	// Show NodeName before Panic
	fmt.Printf("\n pm.Latitude: %v\n", pm.Latitude)
	fmt.Printf("\n pm.Longitude: %v\n", pm.Longitude)
	// hex, _ := b64ToHex(lnsImt.Data)
	// fmt.Printf("Hex: %v\n", hex)

	// Set measurement
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
}

// CONVERT B64 to BYTE
func b64ToByte(b64 string) ([]byte, error) {
	b, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		log.Fatal(err)
	}
	return b, err
}

// RECEIVE DATA FROM PARSE LNS AND DECODE INTO _01
func decodeData(data string) string {
	if data == "" {
		return "No data"
	}

	// B64 to Byte
	b, err := b64ToByte(data)
	if err != nil {
		fmt.Print(data)
		log.Panic(err)
	}

	// var smartlight SmartLight
	// var decoded Decoded
	// var decodedData string

	// protocol parsed
	pp := decode.ImtIotProtocolParser(b)
	// fmt.Printf("PROTOCOL PARSED: [%x]", pp)
	// d is a json
	// decoded is a JSON string to be unmarshal according to deviceType
	// 01: TTTT
	// 02: HHHH
	// 0c: VVVV

	// json.Unmarshal([]byte(d), &decoded)
	// json.Unmarshal([]byte(d), &decoded)

	// switch deviceType {
	// case "SmartLight":
	// 	smartlight.Temperature = float64(decoded._01)
	// 	smartlight.Humidity = float64(decoded._02)
	// 	smartlight.Lux = float64(decoded._03)
	// 	smartlight.Movement = decoded._0D_1
	// 	smartlight.Battery = float64(decoded._0D_2)
	// 	smartlight.BoardVoltage = float64(decoded._0C)

	// 	decodedData, err = json.Marshal(smartlight)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return "SmartLight data wrongly decoded", err
	// 	}
	// default:
	// }

	// fmt.Println(string(pd))
	// return sb.String(), nil
	return string(pp)
}

func main() {
	id := uuid.New().String()
	// var lns LNS
	// var smartlight SmartLight
	// var decodedData Decoded
	// var decodedDataJson []byte

	// var pd string // parsed data

	var lpd string
	// var ppd string
	var sbMqttClientIdSub strings.Builder
	var sbKafkaClientId strings.Builder
	var sbKafkaProducerTopic strings.Builder
	sbMqttClientIdSub.WriteString("parse-lns-lnsImt-mqtt-sub")
	sbKafkaClientId.WriteString("parse-lns-imt-mqtt-")
	sbMqttClientIdSub.WriteString(id)
	sbKafkaClientId.WriteString(id)

	// MQTT
	sTopic := "OpenDataTelemetry/IMT/+/+/rx/+"
	sBroker := "mqtt://mqtt.maua.br:1883"
	sClientId := sbMqttClientIdSub.String()
	sUser := "public"
	sPassword := "public"
	sQos := 0

	sOpts := MQTT.NewClientOptions()
	sOpts.AddBroker(sBroker)
	sOpts.SetClientID(sClientId)
	sOpts.SetUsername(sUser)
	sOpts.SetPassword(sPassword)

	c := make(chan [2]string)

	sOpts.SetDefaultPublishHandler(func(mqttClient MQTT.Client, msg MQTT.Message) {
		c <- [2]string{msg.Topic(), string(msg.Payload())}
	})

	sClient := MQTT.NewClient(sOpts)
	if token := sClient.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	} else {
		fmt.Printf("Connected to %s\n", sBroker)
	}

	if token := sClient.Subscribe(sTopic, byte(sQos), nil); token.Wait() && token.Error() != nil {
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
		incoming := <-c
		s := strings.Split(incoming[0], "/")
		organization := s[1]
		deviceType := s[2]
		deviceId := s[3]
		lnsOrigin := s[5]
		sbKafkaProducerTopic.Reset()
		sbKafkaProducerTopic.WriteString(organization)
		sbKafkaProducerTopic.WriteString(".")
		sbKafkaProducerTopic.WriteString("SmartCampusMaua")

		fmt.Printf("Organization: %s\n", organization)
		fmt.Printf("KafkaTopic: %s\n", sbKafkaProducerTopic.String())
		fmt.Printf("deviceType: %s\n", deviceType)
		fmt.Printf("DeviceId: %s\n", deviceId)
		fmt.Printf("Message: %s\n", incoming[1])
		fmt.Printf("LNS: %s\n", lnsOrigin)
		fmt.Printf("\n")

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

		switch lnsOrigin {
		case "lns_imt":
			lpd = parseLnsImt(incoming[1]) // lns parsed data
			// if err != nil {
			// 	log.Panic(err)
			// }
			fmt.Printf("PARSED DATA: %s", lpd)
		case "lns_atc":
			// parseLnsAtc()
		default:
		}
		json.Unmarshal([]byte(lpd), &lns)
		// TODO: DECODE DATA ACCORDING MEASUREMENT THEN RETURN A JSON STRING
		dd := decodeData(lns.Data) // decoded data
		// dd: {"01":24.1,"02":95,"03_0":0,"03_1":0,"04":0,"05":0,"06":0,"07":0,"08":0,"09":0,"0A_0":0,"0A_1":0,"0B":518744,"0C":3.294,"0D_0":4095,"0D_1":0,"0D_2":0,"0D_3":0,"0E_0":0,"0E_1":0,"10":0,"11":0,"12":0,"13":0}
		// if err != nil {
		// }
		fmt.Printf("\n\n### PROTOCOL PARSED: %s", dd)

		var decodedDataJson []byte

		switch deviceType {
		case "SmartLight":
			json.Unmarshal([]byte(dd), &decodedData)

			smartlight.Temperature = float64(decodedData.X_01)
			smartlight.Humidity = float64(decodedData.X_02)
			smartlight.Lux = float64(decodedData.X_03_0)
			smartlight.Movement = decodedData.X_0D_0
			smartlight.Battery = float64(decodedData.X_0D_1)
			smartlight.BoardVoltage = float64(decodedData.X_0C)

			var err any
			decodedDataJson, err = json.Marshal(smartlight)
			if err != nil {
				fmt.Println(err)
			}
		default:
		}

		fmt.Printf("\n\n###ecodedDataJson:%s", string(decodedDataJson))

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
	}
}

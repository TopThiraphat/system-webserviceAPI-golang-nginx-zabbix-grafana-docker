package http

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"math"

	// "net/http"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	domain "web/pkg/domain"
)

func Place_T2_DepartureSecurity(router fiber.Router) {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	ENV_PLACE_T2_DEPARTURESECURITY := os.Getenv("ENV_PLACE_T2_DEPARTURESECURITY")
	fmt.Println(ENV_PLACE_T2_DEPARTURESECURITY)

	router.Get("/place_T2_DepartureSecurity", func(c *fiber.Ctx) error {

		// resq, err := http.Get(ENV_PLACE_T2_DEPARTURESECURITY)
		// if err != nil {
		// 	log.Fatalln(err)
		// }

		// byteValue, _ := ioutil.ReadAll(resq.Body)
		// xml.Unmarshal(byteValue, &place)

		xmlFile1, err1 := os.Open("./place_T2_DepartureSecurity.xml")
		if err1 != nil {
			log.Panic("!!! Unable to connect link Xovis : ", err1)
		}
		defer xmlFile1.Close()
		// defer resp.Body.Close()

		byteValue, _ := ioutil.ReadAll(xmlFile1)

		// byteValue, _ := ioutil.ReadAll(resp.Body)

		var placeT2DepartureSecurity domain.Place

		xml.Unmarshal(byteValue, &placeT2DepartureSecurity)

		if placeT2DepartureSecurity.Place == "Place_T2_DepartureSecurity" &&
			placeT2DepartureSecurity.Value[0].Itemname == "Departure Security" &&
			placeT2DepartureSecurity.Value[1].Itemname == "Departure Security" &&
			placeT2DepartureSecurity.Value[2].Itemname == "Departure Security" &&
			placeT2DepartureSecurity.Value[3].Itemname == "Departure Security" &&
			placeT2DepartureSecurity.Value[4].Itemname == "DOMDSEC 01" &&
			placeT2DepartureSecurity.Value[5].Itemname == "DOMDSEC 02" &&
			placeT2DepartureSecurity.Value[6].Itemname == "DOMDSEC 03" &&
			placeT2DepartureSecurity.Value[7].Itemname == "DOMDSEC 04" &&
			placeT2DepartureSecurity.Value[8].Itemname == "DOMDSEC 05" &&
			placeT2DepartureSecurity.Value[9].Itemname == "DOMDSEC 06" &&
			placeT2DepartureSecurity.Value[10].Itemname == "DOMDSEC 07" &&
			placeT2DepartureSecurity.Value[11].Itemname == "DOMDSEC 08" &&
			placeT2DepartureSecurity.Value[12].Itemname == "DOMDSEC 09" &&
			placeT2DepartureSecurity.Value[13].Itemname == "DOMDSEC 10" &&
			placeT2DepartureSecurity.Value[14].Itemname == "DOMDSEC 11" {

			///
			minute := 0.016666
			nameT2DepartureSecurity := placeT2DepartureSecurity.Place
			queueLengthT2DepartureSecurity, _ := strconv.Atoi(placeT2DepartureSecurity.Value[1].Value)
			waitingTimeT2DepartureSecurity, _ := strconv.Atoi(placeT2DepartureSecurity.Value[2].Value)
			waitDisplayT2DepartureSecurity, _ := strconv.Atoi(placeT2DepartureSecurity.Value[3].Value)
			openChannelsT2DepartureSecurity := 0
			processTimeT2DepartureSecurity_4, _ := strconv.Atoi(placeT2DepartureSecurity.Value[4].Value)
			if processTimeT2DepartureSecurity_4 != 0 {
				openChannelsT2DepartureSecurity += 1
			}
			processTimeT2DepartureSecurity_5, _ := strconv.Atoi(placeT2DepartureSecurity.Value[5].Value)
			if processTimeT2DepartureSecurity_5 != 0 {
				openChannelsT2DepartureSecurity += 1
			}
			processTimeT2DepartureSecurity_6, _ := strconv.Atoi(placeT2DepartureSecurity.Value[6].Value)
			if processTimeT2DepartureSecurity_6 != 0 {
				openChannelsT2DepartureSecurity += 1
			}
			processTimeT2DepartureSecurity_7, _ := strconv.Atoi(placeT2DepartureSecurity.Value[7].Value)
			if processTimeT2DepartureSecurity_7 != 0 {
				openChannelsT2DepartureSecurity += 1
			}
			processTimeT2DepartureSecurity_8, _ := strconv.Atoi(placeT2DepartureSecurity.Value[8].Value)
			if processTimeT2DepartureSecurity_8 != 0 {
				openChannelsT2DepartureSecurity += 1
			}
			processTimeT2DepartureSecurity_9, _ := strconv.Atoi(placeT2DepartureSecurity.Value[9].Value)
			if processTimeT2DepartureSecurity_9 != 0 {
				openChannelsT2DepartureSecurity += 1
			}
			processTimeT2DepartureSecurity_10, _ := strconv.Atoi(placeT2DepartureSecurity.Value[10].Value)
			if processTimeT2DepartureSecurity_10 != 0 {
				openChannelsT2DepartureSecurity += 1
			}
			processTimeT2DepartureSecurity_11, _ := strconv.Atoi(placeT2DepartureSecurity.Value[11].Value)
			if processTimeT2DepartureSecurity_11 != 0 {
				openChannelsT2DepartureSecurity += 1
			}
			processTimeT2DepartureSecurity_12, _ := strconv.Atoi(placeT2DepartureSecurity.Value[12].Value)
			if processTimeT2DepartureSecurity_12 != 0 {
				openChannelsT2DepartureSecurity += 1
			}
			processTimeT2DepartureSecurity_13, _ := strconv.Atoi(placeT2DepartureSecurity.Value[13].Value)
			if processTimeT2DepartureSecurity_13 != 0 {
				openChannelsT2DepartureSecurity += 1
			}
			processTimeT2DepartureSecurity_14, _ := strconv.Atoi(placeT2DepartureSecurity.Value[14].Value)
			if processTimeT2DepartureSecurity_14 != 0 {
				openChannelsT2DepartureSecurity += 1
			}
			var waitingT2DepartureSecurity, avgProcessingTimeT2DepartureSecurity float64
			sumprocessTimeT2DepartureSecurity := ((processTimeT2DepartureSecurity_4) +
				(processTimeT2DepartureSecurity_5) +
				(processTimeT2DepartureSecurity_6) +
				(processTimeT2DepartureSecurity_7) +
				(processTimeT2DepartureSecurity_8) +
				(processTimeT2DepartureSecurity_9) +
				(processTimeT2DepartureSecurity_10) +
				(processTimeT2DepartureSecurity_11) +
				(processTimeT2DepartureSecurity_12) +
				(processTimeT2DepartureSecurity_13) +
				(processTimeT2DepartureSecurity_14))
			if openChannelsT2DepartureSecurity == 0 {
				waitingT2DepartureSecurity = math.Ceil((float64((sumprocessTimeT2DepartureSecurity)+waitDisplayT2DepartureSecurity) * minute))
				avgProcessingTimeT2DepartureSecurity = math.Ceil(float64(sumprocessTimeT2DepartureSecurity))
			}
			if openChannelsT2DepartureSecurity != 0 {
				waitingT2DepartureSecurity = math.Ceil((float64((sumprocessTimeT2DepartureSecurity)/openChannelsT2DepartureSecurity+waitDisplayT2DepartureSecurity) * minute))
				avgProcessingTimeT2DepartureSecurity = math.Ceil(float64((sumprocessTimeT2DepartureSecurity) / openChannelsT2DepartureSecurity))
			}
			var conditionT2DepartureSecurity float64
			if queueLengthT2DepartureSecurity == 1 && waitingT2DepartureSecurity == 1 {
				conditionT2DepartureSecurity = 1
			} else if queueLengthT2DepartureSecurity != 0 && waitingT2DepartureSecurity <= 1 {
				conditionT2DepartureSecurity = 1
			} else if queueLengthT2DepartureSecurity == 0 && waitingT2DepartureSecurity != 0 {
				conditionT2DepartureSecurity = waitingT2DepartureSecurity
			} else if queueLengthT2DepartureSecurity == 0 && waitingT2DepartureSecurity == 0 {
				conditionT2DepartureSecurity = 0
			} else {
				conditionT2DepartureSecurity = waitingT2DepartureSecurity
			}

			setupResponse := []domain.Response{}
			setupResponse = append(setupResponse, domain.Response{
				Name:                nameT2DepartureSecurity,
				QueueLength:         queueLengthT2DepartureSecurity,
				PredictWaiting:      waitDisplayT2DepartureSecurity,
				WaitingTime:         waitingTimeT2DepartureSecurity,
				OpenChannel:         openChannelsT2DepartureSecurity,
				AvgProcessingTime:   avgProcessingTimeT2DepartureSecurity,
				WaitingTimeOnscreen: conditionT2DepartureSecurity,
			})
			setupResponse = append(setupResponse, domain.Response{
				Name:                nameT2DepartureSecurity,
				QueueLength:         queueLengthT2DepartureSecurity,
				PredictWaiting:      waitDisplayT2DepartureSecurity,
				WaitingTime:         waitingTimeT2DepartureSecurity,
				OpenChannel:         openChannelsT2DepartureSecurity,
				AvgProcessingTime:   avgProcessingTimeT2DepartureSecurity,
				WaitingTimeOnscreen: conditionT2DepartureSecurity,
			})
			setupResponse = append(setupResponse, domain.Response{
				Name:                nameT2DepartureSecurity,
				QueueLength:         queueLengthT2DepartureSecurity,
				PredictWaiting:      waitDisplayT2DepartureSecurity,
				WaitingTime:         waitingTimeT2DepartureSecurity,
				OpenChannel:         openChannelsT2DepartureSecurity,
				AvgProcessingTime:   avgProcessingTimeT2DepartureSecurity,
				WaitingTimeOnscreen: conditionT2DepartureSecurity,
			})
			return c.Status(fiber.StatusOK).JSON(&fiber.Map{
				"result":  fiber.StatusOK,
				"message": setupResponse,
			})
		} else {
			return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
				"result":  fiber.StatusNotFound,
				"message": "!!! The link structure has been fixed.",
			})
		}
	})
}

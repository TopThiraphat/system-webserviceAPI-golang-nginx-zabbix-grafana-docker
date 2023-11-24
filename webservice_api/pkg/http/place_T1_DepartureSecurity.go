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

func Place_T1_DepartureSecurity(router fiber.Router) {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	ENV_PLACE_T2_DEPARTURESECURITY := os.Getenv("ENV_PLACE_T2_DEPARTURESECURITY")
	fmt.Println(ENV_PLACE_T2_DEPARTURESECURITY)

	router.Get("/place_T1_DepartureSecurity", func(c *fiber.Ctx) error {

		// resq, err := http.Get(ENV_PLACE_T2_DEPARTURESECURITY)
		// if err != nil {
		// 	log.Fatalln(err)
		// }

		// byteValue, _ := ioutil.ReadAll(resq.Body)
		// xml.Unmarshal(byteValue, &place)

		xmlFile1, err1 := os.Open("./place_T1_DepartureSecurity.xml")
		if err1 != nil {
			log.Panic("!!! Unable to connect link Xovis : ", err1)
		}
		defer xmlFile1.Close()
		// defer resp.Body.Close()

		byteValue, _ := ioutil.ReadAll(xmlFile1)

		// byteValue, _ := ioutil.ReadAll(resp.Body)

		var placeT1DepartureSecurity domain.Place

		xml.Unmarshal(byteValue, &placeT1DepartureSecurity)

		if placeT1DepartureSecurity.Place == "Place_T1_DepartureSecurity" &&
			placeT1DepartureSecurity.Value[0].Itemname == "PAX Queue" &&
			placeT1DepartureSecurity.Value[1].Itemname == "PAX Queue" &&
			placeT1DepartureSecurity.Value[2].Itemname == "PAX Queue" &&
			placeT1DepartureSecurity.Value[3].Itemname == "PAX Queue" &&
			placeT1DepartureSecurity.Value[4].Itemname == "INTDSEC 01" &&
			placeT1DepartureSecurity.Value[5].Itemname == "INTDSEC 02" &&
			placeT1DepartureSecurity.Value[6].Itemname == "INTDSEC 03" &&
			placeT1DepartureSecurity.Value[7].Itemname == "INTDSEC 04" &&
			placeT1DepartureSecurity.Value[8].Itemname == "INTDSEC 05" &&
			placeT1DepartureSecurity.Value[9].Itemname == "INTDSEC 06" &&
			placeT1DepartureSecurity.Value[10].Itemname == "INTDSEC 07" &&
			placeT1DepartureSecurity.Value[11].Itemname == "INTDSEC 08" &&
			placeT1DepartureSecurity.Value[12].Itemname == "INTDSEC 09" {

			///
			minute := 0.016666
			nameT1DepartureSecurity := placeT1DepartureSecurity.Place
			queueLengthT1DepartureSecurity, _ := strconv.Atoi(placeT1DepartureSecurity.Value[1].Value)
			waitingTimeT1DepartureSecurity, _ := strconv.Atoi(placeT1DepartureSecurity.Value[2].Value)
			waitDisplayT1DepartureSecurity, _ := strconv.Atoi(placeT1DepartureSecurity.Value[3].Value)
			openChannelsT1DepartureSecurity := 0
			processTimeT1DepartureSecurity_4, _ := strconv.Atoi(placeT1DepartureSecurity.Value[4].Value)
			if processTimeT1DepartureSecurity_4 != 0 {
				openChannelsT1DepartureSecurity += 1
			}
			processTimeT1DepartureSecurity_5, _ := strconv.Atoi(placeT1DepartureSecurity.Value[5].Value)
			if processTimeT1DepartureSecurity_5 != 0 {
				openChannelsT1DepartureSecurity += 1
			}
			processTimeT1DepartureSecurity_6, _ := strconv.Atoi(placeT1DepartureSecurity.Value[6].Value)
			if processTimeT1DepartureSecurity_6 != 0 {
				openChannelsT1DepartureSecurity += 1
			}
			processTimeT1DepartureSecurity_7, _ := strconv.Atoi(placeT1DepartureSecurity.Value[7].Value)
			if processTimeT1DepartureSecurity_7 != 0 {
				openChannelsT1DepartureSecurity += 1
			}
			processTimeT1DepartureSecurity_8, _ := strconv.Atoi(placeT1DepartureSecurity.Value[8].Value)
			if processTimeT1DepartureSecurity_8 != 0 {
				openChannelsT1DepartureSecurity += 1
			}
			processTimeT1DepartureSecurity_9, _ := strconv.Atoi(placeT1DepartureSecurity.Value[9].Value)
			if processTimeT1DepartureSecurity_9 != 0 {
				openChannelsT1DepartureSecurity += 1
			}
			var waitingT1DepartureSecurity, avgProcessingTimeT1DepartureSecurity float64
			sumprocessTimeT1DepartureSecurity := ((processTimeT1DepartureSecurity_4) +
				(processTimeT1DepartureSecurity_5) +
				(processTimeT1DepartureSecurity_6) +
				(processTimeT1DepartureSecurity_7) +
				(processTimeT1DepartureSecurity_8) +
				(processTimeT1DepartureSecurity_9))

			if openChannelsT1DepartureSecurity == 0 {
				waitingT1DepartureSecurity = math.Ceil((float64((sumprocessTimeT1DepartureSecurity)+waitDisplayT1DepartureSecurity) * minute))
				avgProcessingTimeT1DepartureSecurity = math.Ceil(float64(sumprocessTimeT1DepartureSecurity))
			}
			if openChannelsT1DepartureSecurity != 0 {
				waitingT1DepartureSecurity = math.Ceil((float64((sumprocessTimeT1DepartureSecurity)/openChannelsT1DepartureSecurity+waitDisplayT1DepartureSecurity) * minute))
				avgProcessingTimeT1DepartureSecurity = math.Ceil(float64((sumprocessTimeT1DepartureSecurity) / openChannelsT1DepartureSecurity))
			}
			///

			///
			var conditionT1DepartureSecurity float64
			if queueLengthT1DepartureSecurity == 1 && waitingT1DepartureSecurity == 1 {
				conditionT1DepartureSecurity = 1
			} else if queueLengthT1DepartureSecurity != 0 && waitingT1DepartureSecurity <= 1 {
				conditionT1DepartureSecurity = 1
			} else if queueLengthT1DepartureSecurity == 0 && waitingT1DepartureSecurity != 0 {
				conditionT1DepartureSecurity = waitingT1DepartureSecurity
			} else if queueLengthT1DepartureSecurity == 0 && waitingT1DepartureSecurity == 0 {
				conditionT1DepartureSecurity = 0
			} else {
				conditionT1DepartureSecurity = waitingT1DepartureSecurity
			}

			setupResponse := []domain.Response{}
			setupResponse = append(setupResponse, domain.Response{
				Name:                nameT1DepartureSecurity,
				QueueLength:         queueLengthT1DepartureSecurity,
				PredictWaiting:      waitDisplayT1DepartureSecurity,
				WaitingTime:         waitingTimeT1DepartureSecurity,
				OpenChannel:         openChannelsT1DepartureSecurity,
				AvgProcessingTime:   avgProcessingTimeT1DepartureSecurity,
				WaitingTimeOnscreen: conditionT1DepartureSecurity,
			})
			setupResponse = append(setupResponse, domain.Response{
				Name:                nameT1DepartureSecurity,
				QueueLength:         queueLengthT1DepartureSecurity,
				PredictWaiting:      waitDisplayT1DepartureSecurity,
				WaitingTime:         waitingTimeT1DepartureSecurity,
				OpenChannel:         openChannelsT1DepartureSecurity,
				AvgProcessingTime:   avgProcessingTimeT1DepartureSecurity,
				WaitingTimeOnscreen: conditionT1DepartureSecurity,
			})
			setupResponse = append(setupResponse, domain.Response{
				Name:                nameT1DepartureSecurity,
				QueueLength:         queueLengthT1DepartureSecurity,
				PredictWaiting:      waitDisplayT1DepartureSecurity,
				WaitingTime:         waitingTimeT1DepartureSecurity,
				OpenChannel:         openChannelsT1DepartureSecurity,
				AvgProcessingTime:   avgProcessingTimeT1DepartureSecurity,
				WaitingTimeOnscreen: conditionT1DepartureSecurity,
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

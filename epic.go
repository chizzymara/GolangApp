package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

//Slack Request Body
type SlackRequestBody struct {
	Text string `json:"text"`
}

// A Response struct to map the Entire Response
type Response struct {
	Data Data `json:"data"`
}

// A Data struct to map the data key
type Data struct {
	Catalog Catalog `json:"Catalog"`
}

// A Catalog struct to map the Catalog key
type Catalog struct {
	SearchStore SearchStore `json:"searchStore"`
}

type SearchStore struct {
	Element []Element `json:"elements"`
}

type Element struct {
	Title      string     `json:"title"`
	Promotions Promotions `json:"promotions"`
}

type Promotions struct {
	PromotionalOffers []PromotionalOffers `json:"promotionalOffers"`
}

type PromotionalOffers struct {
	MainPromotionalOffers []MainPromotionalOffers `json:"promotionalOffers"`
}

type MainPromotionalOffers struct {
	StartDate       string          `json:"startDate"`
	EndDate         string          `json:"endDate"`
	DiscountSetting DiscountSetting `json:"discountSetting"`
}

type DiscountSetting struct {
	DiscountPercentage int `json:"discountPercentage"`
}

func main() {

	webhookUrl := os.Getenv("SLACK_URL")
	freeGames := strings.Join(GetFreeGames(), "\n")
	// Slack message
	message := " Hurry the following games are free: \n" + freeGames
	err := SendSlackNotification(webhookUrl, message)
	if err != nil {
		log.Fatal(err)
	}

}

func GetFreeGames() []string {

	var FreeGames []string
	response, err := http.Get("https://store-site-backend-static-ipv4.ak.epicgames.com/freeGamesPromotions?locale=en-US&country=PL&allowCountries=PL")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)

	for i := 0; i < len(responseObject.Data.Catalog.SearchStore.Element); i++ {
		for x := 0; x < len(responseObject.Data.Catalog.SearchStore.Element[i].Promotions.PromotionalOffers); x++ {
			for y := 0; y < len(responseObject.Data.Catalog.SearchStore.Element[i].Promotions.PromotionalOffers[x].MainPromotionalOffers); y++ {
				discount := responseObject.Data.Catalog.SearchStore.Element[i].Promotions.PromotionalOffers[x].MainPromotionalOffers[y].DiscountSetting.DiscountPercentage

				if discount == 0 {

					FreeGames = append(FreeGames, responseObject.Data.Catalog.SearchStore.Element[i].Title)
				}

			}
		}
	}
	return FreeGames
}

func SendSlackNotification(webhookUrl string, msg string) error {

	slackBody, _ := json.Marshal(SlackRequestBody{Text: msg})
	req, err := http.NewRequest(http.MethodPost, webhookUrl, bytes.NewBuffer(slackBody))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	if buf.String() != "ok" {
		return errors.New("Non-ok response returned from Slack")
	}
	return nil
}

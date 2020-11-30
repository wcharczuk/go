package main

import (
	"context"
	"os"
	"os/signal"
	"strings"
	"time"

	firefox "github.com/njasm/marionette_client"

	"go.charczuk.com/sdk/graceful"
	"go.charczuk.com/sdk/log"
	"go.charczuk.com/sdk/uuid"
)

func main() {
	log.Info("sentinel starting")

	sessionID := uuid.V4().String()

	fc := firefox.NewClient()
	if err := fc.Connect("127.0.0.1", 2828); err != nil {
		log.Error(err)
		os.Exit(1)
	}
	if _, err := fc.NewSession(sessionID, nil); err != nil {
		log.Error(err)
		os.Exit(1)
	}

	ctx, cancel := context.WithCancel(context.Background())
	go checkForUpdates(ctx, fc)
	quit := graceful.Notify(graceful.DefaultSignals...)
	select {
	case <-quit:
		log.Info("sentinel received quit signal")
		signal.Stop(quit)
		cancel()
	}
	log.Info("sentinel deleting firefox session")
	if _, err := fc.CloseWindow(); err != nil {
		log.Error(err)
	}
	log.Info("sentinel exiting")
}

var page = "https://www.bestbuy.com/site/gigabyte-geforce-rtx-3080-10g-gddr6x-pci-express-4-0-graphics-card-black/6436223.p?skuId=6436223"

func checkForUpdates(ctx context.Context, fc *firefox.Client) error {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return nil
		case <-ticker.C:
			if _, err := fc.Navigate(page); err != nil {
				return err
			}
			button, err := fc.FindElement(firefox.XPATH, "/html/body/div[3]/main/div[2]/div[3]/div[2]/div/div/div[7]/div[1]/div/div/div/button")
			if err != nil {
				return err
			}
			if button == nil {
				log.Warning("buy page has no buy button")
			} else {
				if strings.TrimSpace(button.Text()) != "Sold Out" {
					log.Infof("buy page text: %s", button.Text())
				}
			}
		}

	}
}

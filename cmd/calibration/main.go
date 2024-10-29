package main

import (
	"flag"
	"jackson/internal/mover"
	"os"

	"github.com/sirupsen/logrus"
)

const (
	sensorsIP = "localhost:8080"
	motorsIP  = sensorsIP
	robotID   = "F535AF9628574A53"
)

func main() {
	sip := flag.String("sip", sensorsIP, "")
	mip := flag.String("bip", motorsIP, "")
	id := flag.String("id", robotID, "")

	log := logrus.New()

	log.SetLevel(logrus.DebugLevel)
	log.SetFormatter(&logrus.JSONFormatter{})

	log.SetOutput(os.Stdout)

	sm := mover.NewSmartMover(log.WithField("entity", "smart-mover-calibration"), *sip, *mip, *id)
	sm.Calibrate()
}

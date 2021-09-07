package logger

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"io"
	"io/ioutil"
	"os"
)

func SetUpLogger() {
	tempFile, err := ioutil.TempFile(os.TempDir(), "conference_")
	if err != nil {
		log.Fatal().Err(err).Msg("there was an error creating a temporary file for log")
	}
	fmt.Printf("The log file is allocated at %s\n", tempFile.Name())
	wrt := io.MultiWriter(os.Stdout, tempFile)
	log.Logger = zerolog.New(wrt)
}

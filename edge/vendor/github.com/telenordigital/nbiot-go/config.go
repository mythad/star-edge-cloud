package nbiot

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const (
	// DefaultAddr is the default address of the Telenor NB-IoT API. You normally won't
	// have to change this.
	DefaultAddr = "https://api.nbiot.telenor.io"

	// ConfigFile is the name for the config file. The configuration file is a
	// plain text file that contains the Telenor NB-IoT configuration.
	// The configuration file is expected to be in the current home directory
	// and contain a "address=<value>" line and/or a "token=<value>" line.
	ConfigFile = ".telenor-nbiot"

	// AddressEnvironmentVariable is the name of the environment variable that
	// can be used to override the address set in the configuration file.
	// If the  environment variable isn't set (or is empty) the configuration
	// file settings will be used.
	AddressEnvironmentVariable = "TELENOR_NBIOT_ADDRESS"

	// TokenEnvironmentVariable is the name of the environment variable that
	// can be used to override the token set in the configuration file.
	TokenEnvironmentVariable = "TELENOR_NBIOT_TOKEN"
)

// These are the configuration file directives.
const (
	addressKey = "address"
	tokenKey   = "token"
)

// Return both address and token from configuration file. The file name is
// for testing purposes; use the ConfigFile constant when calling the functino.
func addressTokenFromConfig(filename string) (string, string, error) {
	address, token, err := readConfig(getFirstMatchingPath(filename))
	if err != nil {
		return "", "", err
	}

	envAddress := os.Getenv(AddressEnvironmentVariable)
	if envAddress != "" {
		address = envAddress
	}

	envToken := os.Getenv(TokenEnvironmentVariable)
	if envToken != "" {
		token = envToken
	}

	return address, token, nil
}

// getFirstMatchingPath traverses the directories from the current
// working directory to the home directory looking for the config
// file.  It will return the path of the first file encountered.
func getFirstMatchingPath(filename string) string {
	dir, err := os.Getwd()
	if err != nil {
		return ""
	}

	cwdParts := strings.Split(dir, string(os.PathSeparator))

	for i := len(cwdParts); i >= 1; i-- {
		s := strings.Join(cwdParts[:i], string(os.PathSeparator)) + string(os.PathSeparator)
		f := filepath.Join(s, filename)
		if _, err := os.Stat(f); os.IsNotExist(err) {
			continue
		}
		return f
	}

	return ""
}

// readConfig reads the config file and returns the address and token
// settings from the file.
func readConfig(filename string) (string, string, error) {
	address := DefaultAddr
	token := ""

	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return address, token, nil
	}
	scanner := bufio.NewScanner(bytes.NewReader(buf))
	scanner.Split(bufio.ScanLines)
	lineno := 0
	for scanner.Scan() {
		lineno++
		line := strings.ToLower(scanner.Text())
		if len(line) == 0 || line[0] == '#' {
			// ignore comments and empty lines
			continue
		}
		words := strings.Split(scanner.Text(), "=")
		if len(words) != 2 {
			return "", "", fmt.Errorf("Not a key value expression on line %d in %s: %s\n", lineno, filename, scanner.Text())
		}
		switch words[0] {
		case addressKey:
			address = strings.TrimSpace(words[1])
		case tokenKey:
			token = strings.TrimSpace(words[1])
		default:
			return "", "", fmt.Errorf("Unknown keyword on line %d in %s: %s\n", lineno, filename, scanner.Text())
		}
	}
	return address, token, nil
}

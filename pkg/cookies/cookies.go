package cookies

import (
	"encoding/json"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/zellyn/kooky"
	_ "github.com/zellyn/kooky/allbrowsers" // register cookie store finders!
	"github.com/zellyn/kooky/chrome"
	"github.com/zellyn/kooky/safari"
	"io/ioutil"
	"os"
	"runtime"
)

func ScannLocalCookies(domainStr string, namesStr string, storageDir string) error {
	var (
		err          error
		cookiefiltes []kooky.Filter
		cookiesFile  string
	)
	currentOs := runtime.GOOS
	if domainStr != "" {
		cookiefiltes = append(cookiefiltes, kooky.DomainHasSuffix(domainStr))
	}
	if namesStr != "" {
		cookiefiltes = append(cookiefiltes, kooky.NameContains(namesStr))
	}
	if storageDir == "" {
		storageDir, _ = os.Getwd()
	}

	if currentOs == "windows" { //windows
		cookiesFile = storageDir + "\\" + currentOs + "cookies.json"

		cookies := kooky.ReadCookies(cookiefiltes...)
		file, _ := json.MarshalIndent(cookies, "", " ")
		err = ioutil.WriteFile(cookiesFile, file, 0644)
		if err != nil {
			log.Error(err)
			return err
		}

	} else if currentOs == "darwin" { //mac
		dir, err := os.UserConfigDir() // "/<USER>/Library/Application Support/"
		if err != nil {
			log.Error(err)
			return err
		}
		cookiesFileWithChrome := dir + "/Google/Chrome/Default/Cookies"
		cookiesInChrome, err := chrome.ReadCookies(cookiesFileWithChrome)
		if err != nil {
			log.Error(err)
			return err
		}
		cookiesFile = storageDir + "/" + currentOs + "cookies.json"
		fileCookieInChrome, err := json.MarshalIndent(cookiesInChrome, "", " ")
		if err != nil {
			log.Error(err)
			return err
		}
		err = ioutil.WriteFile(cookiesFile, fileCookieInChrome, 0644)
		if err != nil {
			log.Error(err)
			return err
		}

		cookiesFileWithSafari := dir + "/Library/Cookies/Cookies.binarycookies"
		cookiesInSafari, err := safari.ReadCookies(cookiesFileWithSafari)
		if err != nil {
			log.Error(err)
			return err
		}
		fileCookieInSafari, err := json.MarshalIndent(cookiesInSafari, "", " ")
		if err != nil {
			log.Error(err)
			return err
		}
		err = ioutil.WriteFile(cookiesFile, fileCookieInSafari, 0644)
		if err != nil {
			log.Error(err)
			return err
		}

	} else {
		return errors.New(fmt.Sprintf("current os %s is not support now!", currentOs))
	}

	return err
}

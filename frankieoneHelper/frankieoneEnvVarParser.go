package frankieoneHelper

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/golang/glog"
	"github.com/spf13/viper"
)

const (
	_ErrorPrefix = "Fatal error config file: %s"
)

type configMap struct {
	label string
	value interface{}
}

func ParseFlags(configPath string) {
	flags := map[string]string{
		"v":               "10",
		"stderrthreshold": "INFO",
		"logtostderr":     "true",
		"cnf":             configPath,
	}

	for k, v := range flags {
		f := flag.Lookup(k)

		if f == nil {
			flag.String(k, v, "")
		} else {
			if err := f.Value.Set(v); err != nil {
				glog.Error("unable to set the value", "error", err)
			}
		}
	}

	flag.Parse()
}

func LoadConfig() (*viper.Viper, error) {

	files, err := ioutil.ReadDir(flag.Lookup("cnf").Value.String())
	if err != nil {
		glog.Error(fmt.Errorf(_ErrorPrefix, err))
		return nil, err
	}

	config := viper.New()

	for _, file := range files {
		cfg := load(file.Name())
		config.Set(cfg.label, cfg.value)
	}

	return config, nil

}

func load(path string) (cfmap *configMap) {
	var (
		name       string
		ext        string
		cfgContent string
	)

	if ext = filepath.Ext(path); ext != ".yaml" && ext != ".yml" {
		return nil
	}

	name = strings.TrimSuffix(path, ext)
	content, err := ioutil.ReadFile(filepath.Join(flag.Lookup("cnf").Value.String(), path))
	if err != nil {
		log.Fatal(fmt.Errorf(_ErrorPrefix, err))
	}

	cfgContent = os.ExpandEnv(string(content[:]))

	v := viper.New()
	v.SetConfigName(name)
	v.SetConfigType("yaml")

	// Read the config file
	err = v.ReadConfig(bytes.NewBuffer([]byte(cfgContent)))

	if err != nil {
		// Handle errors reading the config file
		log.Fatal(fmt.Errorf(_ErrorPrefix, err))
	}

	return &configMap{label: name, value: v.AllSettings()}
}

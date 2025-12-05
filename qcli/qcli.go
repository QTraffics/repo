package qcli

import (
	"os"
	"path/filepath"
	"sort"

	"github.com/qtraffics/qtfra/enhancements/slicelib"
	"github.com/qtraffics/qtfra/ex"
	"github.com/qtraffics/qtfra/log"
	"github.com/qtraffics/repo/qcli/options"

	"github.com/goccy/go-yaml"
)

type QCli struct {
	logger log.Logger
}

func Create(configFile, configFileDir string) (*QCli, error) {
	var (
		option options.Option
		err    error
	)

	if len(configFileDir) != 0 {
		option, err = readConfigAndMerge(configFileDir)
	} else if len(configFile) != 0 {
		option, err = readConfig(configFile)
	}

	if err != nil {
		return nil, ex.Cause(err, "Configure")
	}

	return createWithOption(option)
}

func createWithOption(option options.Option) (*QCli, error) {
	panic("implement me")
}

func readConfigAndMerge(ConfigFileDir string) (options.Option, error) {
	var root options.Option
	dir, err := os.ReadDir(ConfigFileDir)
	if err != nil {
		return root, ex.Cause(err, "ReadDir")
	}
	fileToRead := slicelib.Map(dir, func(it os.DirEntry) string {
		if it.IsDir() || (filepath.Ext(it.Name()) != ".yaml" && filepath.Ext(it.Name()) != ".yml") {
			return ""
		}
		return it.Name()
	})
	sort.Strings(fileToRead)
	for _, file := range fileToRead {
		if file == "" {
			continue
		}
		option, err := readConfig(file)
		if err != nil {
			return options.Option{}, err
		}
		root = options.Merge(root, option)
	}
	return root, nil
}

func readConfig(path string) (options.Option, error) {
	openFile, err := os.Open(path)
	if err != nil {
		return options.Option{}, ex.Cause(err, "Open")
	}
	defer openFile.Close()

	var current options.Option
	decoder := yaml.NewDecoder(openFile, yaml.DisallowUnknownField())
	err = decoder.Decode(&current)
	if err != nil {
		return options.Option{}, ex.Cause(err, "Decode")
	}
	return current, nil
}

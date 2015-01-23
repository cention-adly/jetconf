// Package jetconf (json-etc-configuration) simplifies reading of
// JSON-formatted configuration files in /etc. It uses the basename of the
// program name as the configuration filename - for example if you program is
// /foo/bar then then it will read /etc/bar.conf.
package jetconf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// Load loads /etc/$(basename $0).conf.
func Load(v interface{}) error {
	return LoadAtPrefix("", v)
}

// LoadAtPrefix loads $prefix/etc/$(basename $0).conf.
func LoadAtPrefix(prefix string, v interface{}) error {
	return readConfig(prefix, v)
}

// MustLoad is same as Load but panics on error.
func MustLoad(v interface{}) {
	MustLoadAtPrefix("", v)
}

// MustLoadAtPrefix is same as LoadAtPrefix but panics on error.
func MustLoadAtPrefix(prefix string, v interface{}) {
	err := readConfig(prefix, v)
	if err != nil {
		panic(err)
	}
}

func readConfig(prefix, v interface{}) error {
	name := fmt.Sprintf("%v/etc/%s.conf",
		prefix,
		baseName(os.Args[0]))
	buf, err := ioutil.ReadFile(name)
	if err != nil {
		return err
	}

	err = json.Unmarshal(buf, v)
	if err != nil {
		panic("Error: " + err.Error())
	}
	return nil
}

func baseName(filename string) string {
	lastSlash := strings.LastIndex(filename, "/")
	if lastSlash == -1 {
		return filename
	}
	return filename[lastSlash+1:]
}

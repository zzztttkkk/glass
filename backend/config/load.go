package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/imdario/mergo"
	"github.com/zzztttkkk/sha/utils"
)

func fromTomlBytes(conf interface{}, data []byte) error {
	_, err := toml.Decode(string(data), conf)
	return err
}

var envReg = regexp.MustCompile(`\$ENV{.*?}`)

func doReplace(fp, name string, value *reflect.Value, path []string) {
	key := strings.Join(path, ".") + "." + name

	rawValue := (*value).Interface().(string)
	if strings.HasPrefix(rawValue, "file://") {
		_fp := rawValue[7:]
		var f *os.File
		var e error

		if !strings.HasPrefix(_fp, "/") {
			_fp = filepath.Join(filepath.Dir(fp), _fp)
		}

		f, e = os.Open(_fp)
		if e != nil {
			log.Fatalf("glass.config: key: `%s`; raw: `%s`; err: `%s`\n", key, rawValue, e.Error())
		}
		defer f.Close()

		data, e := ioutil.ReadAll(f)
		if e != nil {
			log.Fatalf("glass.config: key: `%s`; raw: `%s`; err: `%s`\n", key, rawValue, e.Error())
		}
		value.SetString(string(data))
		return
	}

	s := envReg.ReplaceAllFunc(
		utils.B(rawValue),
		func(data []byte) []byte {
			envK := strings.TrimSpace(string(data[5 : len(data)-1]))
			v := os.Getenv(envK)
			return []byte(v)
		},
	)
	value.SetString(string(s))
}

func reflectMap(filePath string, value reflect.Value, path []string) {
	ele := value.Elem()
	t := ele.Type()

	for i := 0; i < ele.NumField(); i++ {
		filed := ele.Field(i)

		tf := t.Field(i)
		if tf.Tag.Get("toml") == "-" {
			continue
		}
		switch filed.Type().Kind() {
		case reflect.String:
			doReplace(filePath, tf.Name, &filed, path)
		case reflect.Struct:
			cp := path[:]
			cp = append(cp, tf.Name)
			reflectMap(filePath, filed.Addr(), cp)
		}
	}
}

func FromFile(conf interface{}, fp string) error {
	f, e := os.Open(fp)
	if e != nil {
		panic(e)
	}
	defer f.Close()

	v, e := ioutil.ReadAll(f)
	if e != nil {
		panic(e)
	}

	if strings.HasSuffix(fp, ".toml") {
		e = fromTomlBytes(conf, v)
	} else {
		e = json.Unmarshal(v, conf)
	}

	if e != nil {
		return e
	}
	reflectMap(fp, reflect.ValueOf(conf), []string{})
	return nil
}

func FromFiles(dist interface{}, fps ...string) {
	t := dist
	if reflect.TypeOf(t).Kind() != reflect.Ptr {
		panic(fmt.Errorf("glass.config: dist is not a pointer"))
	}

	ct := reflect.TypeOf(dist).Elem()
	if ct.Kind() != reflect.Struct {
		panic(fmt.Errorf("glass.config: dist is not a struct pointer"))
	}

	for _, fp := range fps {
		ele := reflect.New(ct).Interface()
		err := FromFile(ele, fp)
		if err != nil {
			panic(err)
		}
		if t == nil {
			t = ele
		} else {
			if err := mergo.Merge(t, ele); err != nil {
				panic(err)
			}
		}
		log.Printf("glass.conf: load from file `%s`\n", fp)
	}

	defV := _Default()
	defaultValuePtr := &defV

	if reflect.TypeOf(defaultValuePtr).Kind() != reflect.Ptr {
		panic(fmt.Errorf("config: default value is not a pointer"))
	}
	dct := reflect.TypeOf(defaultValuePtr).Elem()
	if dct != ct {
		panic(fmt.Errorf("glass.config: default value type error"))
	}
	if defaultValuePtr != nil {
		if err := mergo.Merge(t, reflect.ValueOf(defaultValuePtr).Elem().Interface()); err != nil {
			panic(err)
		}
	}

	reflectMap("", reflect.ValueOf(t), []string{})
}

package session

import (
	"encoding/json"
	"glass/config"
	"glass/internal"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
)

type TangPoem struct {
	Author    string   `json:"author"`
	Title     string   `json:"title"`
	Contents  []string `json:"contents"`
	nContents [][]string
}

var allTangPoetry []TangPoem
var tangPoetryJsonData []byte

// todo embed when 1.16
func init() {
	internal.DigContainer.Append(
		func(cfg *config.Type) {
			f, e := os.Open(cfg.Static.DistPath + "/300_tang_poems.json")
			if e != nil {
				panic(e)
			}
			tangPoetryJsonData, e = ioutil.ReadAll(f)
			if e != nil {
				panic(e)
			}
			e = json.Unmarshal(tangPoetryJsonData, &allTangPoetry)
			if e != nil {
				panic(e)
			}

			for i := 0; i < len(allTangPoetry); i++ {
				v := &(allTangPoetry[i])
				for _, l := range v.Contents {
					v.nContents = append(v.nContents, strings.Split(l, " "))
				}
			}
		},
	)
}

func RandTangPoem() *TangPoem { return &(allTangPoetry[int(rand.Uint32())%len(allTangPoetry)]) }

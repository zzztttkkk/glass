package session

import (
	"encoding/json"
	"glass/config"
	"glass/dist"
	"glass/internal"
	"math/rand"
	"strings"
)

type TangPoem struct {
	Author    string   `json:"author"`
	Title     string   `json:"title"`
	Contents  []string `json:"contents"`
	nContents [][]string
}

var allTangPoetry []TangPoem

func init() {
	internal.DigContainer.Append(
		func(cfg *config.Type) {
			e := json.Unmarshal(dist.TangPoems, &allTangPoetry)
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

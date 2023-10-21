package main

import "fmt"

type Pair struct {
	Groups [2]GroupPair `json:"groups"`
	Start  PairTime     `json:"start,omitempty"`
	End    PairTime     `json:"end,omitempty"`
}

type GroupPair struct {
	Name string   `json:"name"`
	Room int16    `json:"room"`
	Type PairType `json:"type"`
}

type PairType string

func (t PairType) String() string {
	if t != "" {
		return fmt.Sprintf("(%s)", string(t))
	}
	return ""
}

func (p *Pair) Answer() string {
	if p.Groups[0] == p.Groups[1] {
		return "    " + p.Groups[0].Name
	} else {
		var pairString string
		for n, pair := range p.Groups {
			pairString += fmt.Sprintf(
				"    У ПС-1%d %s %s\n",
				n+1, pair.Name, pair.Type,
			)
		}
		return pairString
	}
}

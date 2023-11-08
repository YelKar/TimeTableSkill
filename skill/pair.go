package main

import "fmt"

type Pair struct {
	Groups    []GroupPair     `json:"groups,omitempty"`
	WeekSplit *[2][]GroupPair `json:"week_split,omitempty"`
	Start     PairTime        `json:"start,omitempty"`
	End       PairTime        `json:"end,omitempty"`
}

type GroupPair struct {
	Name  string   `json:"name"`
	Room  int16    `json:"room"`
	Group string   `json:"group"`
	Type  PairType `json:"type"`
}

type PairType string

func (t PairType) String() string {
	if t != "" {
		return fmt.Sprintf("(%s)", string(t))
	}
	return ""
}

func (p *Pair) Answer() string {
	var pairString string
	for _, pair := range p.Groups {
		pairString += "    "
		if pair.Group != "" {
			pairString += fmt.Sprintf(
				"У %s ",
				pair.Group,
			)
		}
		pairString += fmt.Sprintf(
			"%s %s\n",
			pair.Name, pair.Type,
		)
	}
	return pairString
}

package meechum

import (
	"encoding/json"
	"strings"
	"time"
)

type Check struct {
	Name   string `json:"name"`
	Every  int    `json:"every"`
	Repeat int    `json:"repeat"`
	Cmd    string `json:"cmd"`
	File   string
	Args   []string
	Result *chan Result
}

func NewCheck(data []byte, r *chan Result) (*Check, error) {
	c := &Check{
		Result: r,
	}

	err := json.Unmarshal(data, c)
	if err != nil {
		return nil, err
	}
	// Extract the right path for the check
	tmp := strings.Split(Cmd, " ")
	c.File = tmp[0]
	if len(tmp) > 0 {
		c.Args = tmp[1 : len(tmp)-1]
	}
	// Check if the plugin is installed,
	// TODO:
	//   - ifnot we should ask for a download
	if _, err := os.Stat(c.File); os.IsNotExist(err) {
		return nil, fmt.Errorf("The check at %s is not installed on this node", filename)
	}

	if c.Every == 0 {
		c.Every = 10
	}

	if c.Repeat == 0 {
		c.Repeat == 120
	}

	return c, nil
}

func (c *Check) Run() {
	// Create the tickers
	m := time.NewTicker(c.Every * time.Second)
	for {
		select {
		case <-m.C:
			c.Execute()
		}
	}
}

func (c *Check) Execute() error {

}

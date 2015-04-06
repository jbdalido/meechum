package meechum

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type Check struct {
	Name     string       `json:"name"`
	Every    int          `json:"every"`
	Repeat   int          `json:"repeat"`
	Cmd      string       `json:"cmd"`
	Handlers []string     `json:"handlers"`
	File     string       `json:"-"`
	Args     []string     `json:"-"`
	Exec     *Executor    `json:"-"`
	Result   chan *Result `json:"-"`
}

type Result struct {
	Name     string    `json:"name"`
	Code     ErrorCode `json:"errorcode"`
	Error    error     `json:"error"`
	Level    string    `json:"level"`
	StdOut   string    `json:"stdout"`
	Handlers []string  `json:"-"`
}

func (r *Result) MarshalJson() []byte {
	data, err := json.Marshal(r)
	if err != nil {
		log.Printf("[ENGINE] Cannot save check as json %s", err)
		return []byte{}
	}
	return data
}
func NewCheck(data []byte, r chan *Result) (*Check, error) {
	c := &Check{
		Result: r,
	}

	err := json.Unmarshal(data, c)
	if err != nil {
		return nil, err
	}
	// Extract the right path for the check
	tmp := strings.Split(c.Cmd, " ")
	c.File = tmp[0]
	if len(tmp) > 0 {
		c.Args = tmp[1:len(tmp)]
	}
	// Check if the plugin is installed,
	// TODO:
	//   - ifnot we should ask for a download
	if _, err := os.Stat(c.File); os.IsNotExist(err) {
		return nil, fmt.Errorf("The check at %s is not installed on this node", c.File)
	}

	if len(c.Handlers) == 0 {
		c.Handlers = append(c.Handlers, "log")
	}

	if c.Every == 0 {
		c.Every = 1
	}

	if c.Repeat == 0 {
		c.Repeat = 120
	}

	c.Exec, err = NewExecutor(c.File, "")
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (c *Check) Run() {
	// Create the tickers
	m := time.NewTicker(time.Duration(c.Every) * time.Second)

	for {
		select {
		case <-m.C:
			std, err, code := c.Execute()
			log.Printf("[ERROR] Code: %d Level: %s", code, code.String())
			r := &Result{
				Name:   c.Name,
				Code:   code,
				Level:  code.String(),
				StdOut: std,
				Error:  err,
			}
			go func() {
				c.Result <- r
			}()
		}
	}
}

func (c *Check) Execute() (string, error, ErrorCode) {
	std, err, code := c.Exec.Do(c.Args)
	if err != nil {
		return std, err, ErrorCode(code)
	}
	return std, nil, 3
}

package slack

func Fire(r *meechum.Result) error {

}
func Levels() []string {
	return []string{
		"warning",
		"critical",
	}
}
func String() string {
	return "slack"
}


package main

import (
  "bytes"
  "crypto/tls"
  "encoding/json"
  "fmt"
  "io/ioutil"
  "log"
  "net/http"
  "net/url"
  "strconv"
)

// Client is containing the configured http.Client
// and the host url
type Slack struct {
  Token      string
  Host       *url.URL
  HTTPClient *http.Client
}


type RequestOptions struct {
  Method string
  Path   string
  Data   *Message
}

type Message struct {
  Channel  string `json:"channel"`
  Username string `json:"username,omitempty"`
  Title    string
  Text     string `json:"text"`
  Parse    string `json:"parse"`
  Content  string `json:"content,omitempty"`
}

func NewSlack(host string, tlsConfig *tls.Config) (*Slack, error) {
  h, err := url.Parse(host)
  if err != nil {
    return nil, fmt.Errorf("can't parse host %s", host)
  }

  return &Slack{
    Host:       h,
    HTTPClient: newHTTPClient(h, tlsConfig),
  }, nil
}

func (s *Slack) Fire(a *Alert) error {
  if a.Options.Slack == nil {
    return fmt.Errorf("[SLACK] No options given")
  }

  if a.Options.Slack.User == "" {
    a.Options.Slack.User = "Noise"
  }

  m := &Message{
    Username: a.Options.Slack.User,
    Text:     a.Message,
    Parse:    "full",
    Channel:  a.Options.Slack.Channel,
  }

  o := &RequestOptions{
    Method: "POST",
    Path:   "",
    Data:   m,
  }
  _, _, err := s.request(o)
  if err != nil {
    log.Printf("Slack failed to send %s\n", a.Message)
  }

  return nil
}

func newHTTPClient(u *url.URL, tlsConfig *tls.Config) *http.Client {
  httpTransport := &http.Transport{
    TLSClientConfig: tlsConfig,
  }
  u.Path = ""
  return &http.Client{Transport: httpTransport}
}

// do the actual prepared request in request()
func (s *Slack) do(method, path string, data *url.Values) ([]byte, int, error) {
  var resp *http.Response
  req, err := http.NewRequest(method, s.Host.String()+path, bytes.NewBufferString(data.Encode()))
  if err != nil {
    return nil, -1, err
  }
  // Prepare and do the request
  req.Header.Set("User-Agent", "SlackClient")
  req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
  req.Header.Set("Content-Length", strconv.Itoa(len(data.Encode())))

  resp, err = s.HTTPClient.Do(req)
  if err != nil {
    return nil, -1, err
  }

  defer resp.Body.Close()

  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return nil, -1, err
  }
  if resp.StatusCode >= 400 {
    return nil, resp.StatusCode, fmt.Errorf("SLACKAPI ERROR %d: %s", resp.StatusCode, body)
  }
  return body, resp.StatusCode, nil
}

// request prepare the request by setting the correct methods and parameters
func (s *Slack) request(options *RequestOptions) (int, string, error) {

  if options.Method == "" {
    options.Method = "GET"
  }

  path := options.Path

  v := url.Values{}

  data, err := json.Marshal(options.Data)
  if err != nil {
    return -1, "", err
  }

  v.Add("payload", string(data))
  body, code, err := s.do(options.Method, path, &v)

  if err != nil {
    return -1, "", err
  }

  if code != 200 {
    return -1, "", fmt.Errorf("%s", string(body))
  }

  return code, string(body), nil
}
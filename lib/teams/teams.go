package teams


import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"../errormessages"
	"github.com/pkg/errors"
)

// API
type API struct {
	IncomingWebhookURL string
}

// APIRequest Request to teams api
type APIRequest struct {
	Type     string `json:"@type"`
	Context  string `json:"@context"`
	Summary  string `json:"summary"`
	Sections []APISection `json:"sections"`
	PotentialActions []APIPotentialAction `json:"potentialAction"`
}

type APISection struct {
	Title    string `json:"activityTitle"`
	SubTitle    string `json:"activitySubtitle"`
	Image    string `json:"activityImage"`
	Facts            []APIFact `json:"facts"`
}

type APIFact struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type APIPotentialAction struct {
	Type     string `json:"@type"`
	Name     string `json:"name"`
	Targets  []APIOpenUriTarget `json:"targets"`
}

type APIOpenUriTarget struct {
	OS		string `json:"os"`
	URI		string `json:"uri"`
}


func New(webhookUrl string) (*API, error) {

	if webhookUrl == ""  {
		return nil, errors.New(errormessages.EmptyWebhookURL)
	}

	api := &API{
		IncomingWebhookURL: webhookUrl,
	}

	return api, nil

}
func (api *API) PerformAPIRequest(request *APIRequest) (string, error) {

	//payload := strings.NewReader(string(scriptContent))

	if request == nil {
		request = &APIRequest{}
	}

	if request.Type == "" {
		request.Type = "MessageCard"
	}

	if request.Context == "" {
		request.Context = "http://schema.org/extensions"
	}

	var requestPayload io.Reader

	payloadJSON, _ := json.Marshal(request)
			
	payloadString := string(payloadJSON)
	requestPayload = strings.NewReader(payloadString)

	req, _ := http.NewRequest("POST", api.IncomingWebhookURL, requestPayload)
	req.Header.Add("Content-Type", "application/json")


	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	if res.StatusCode != 200 {
		return "", errors.New(string(body))
	}

	var response string

	json.Unmarshal([]byte(body), &response)

	return response, nil
}
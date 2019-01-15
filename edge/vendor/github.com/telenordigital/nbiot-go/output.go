package nbiot

import "fmt"

// Output represents a data output for a collection
type Output interface {
	toOutput() output
}

type WebHookOutput struct {
	ID                string
	CollectionID      string
	URL               string
	BasicAuthUser     string
	BasicAuthPass     string
	CustomHeaderName  string
	CustomHeaderValue string
}

type MQTTOutput struct {
	ID               string
	CollectionID     string
	Endpoint         string
	DisableCertCheck bool
	Username         string
	Password         string
	ClientID         string
	TopicName        string
}

type IFTTTOutput struct {
	ID           string
	CollectionID string
	Key          string
	EventName    string
	AsIsPayload  bool
}

type UDPOutput struct {
	ID           string
	CollectionID string
	Host         string
	Port         int
}

// Output retrieves an output
func (c *Client) Output(collectionID, outputID string) (Output, error) {
	var output output
	err := c.get(fmt.Sprintf("/collections/%s/outputs/%s", collectionID, outputID), &output)
	if err != nil {
		return nil, err
	}
	return output.toOutput()
}

// Outputs retrieves a list of outputs on a collection
func (c *Client) Outputs(collectionID string) ([]Output, error) {
	var outputs struct {
		Outputs []output `json:"outputs"`
	}
	err := c.get(fmt.Sprintf("/collections/%s/outputs", collectionID), &outputs)
	if err != nil {
		return nil, err
	}

	ret := make([]Output, len(outputs.Outputs))
	for i, o := range outputs.Outputs {
		ret[i], err = o.toOutput()
		if err != nil {
			return nil, err
		}
	}
	return ret, err
}

// CreateOutput creates an output
func (c *Client) CreateOutput(collectionID string, output Output) (Output, error) {
	o := output.toOutput()
	err := c.create(fmt.Sprintf("/collections/%s/outputs", collectionID), &o)
	if err != nil {
		return nil, err
	}
	return o.toOutput()
}

// UpdateOutput updates an output. The type field can't be modified
func (c *Client) UpdateOutput(collectionID string, output Output) (Output, error) {
	o := output.toOutput()
	err := c.update(fmt.Sprintf("/collections/%s/outputs/%s", collectionID, *o.ID), &o)
	if err != nil {
		return nil, err
	}
	return o.toOutput()
}

// DeleteOutput removes an output
func (c *Client) DeleteOutput(collectionID, outputID string) error {
	return c.delete(fmt.Sprintf("/collections/%s/outputs/%s", collectionID, outputID))
}

func (o WebHookOutput) toOutput() output {
	typ := "webhook"
	return output{
		ID:           &o.ID,
		CollectionID: &o.CollectionID,
		Type:         &typ,
		Config: map[string]interface{}{
			"url":               o.URL,
			"basicAuthUser":     o.BasicAuthUser,
			"basicAuthPass":     o.BasicAuthPass,
			"customHeaderName":  o.CustomHeaderName,
			"customHeaderValue": o.CustomHeaderValue,
		},
	}
}

func (o MQTTOutput) toOutput() output {
	typ := "mqtt"
	return output{
		ID:           &o.ID,
		CollectionID: &o.CollectionID,
		Type:         &typ,
		Config: map[string]interface{}{
			"endpoint":         o.Endpoint,
			"disableCertCheck": o.DisableCertCheck,
			"username":         o.Username,
			"password":         o.Password,
			"clientId":         o.ClientID,
			"topicName":        o.TopicName,
		},
	}
}

func (o IFTTTOutput) toOutput() output {
	typ := "ifttt"
	return output{
		ID:           &o.ID,
		CollectionID: &o.CollectionID,
		Type:         &typ,
		Config: map[string]interface{}{
			"key":         o.Key,
			"eventName":   o.EventName,
			"asIsPayload": o.AsIsPayload,
		},
	}
}

func (o UDPOutput) toOutput() output {
	typ := "udp"
	return output{
		ID:           &o.ID,
		CollectionID: &o.CollectionID,
		Type:         &typ,
		Config: map[string]interface{}{
			"host": o.Host,
			"port": o.Port,
		},
	}
}

type output struct {
	ID           *string                `json:"outputId"`
	CollectionID *string                `json:"collectionId"`
	Type         *string                `json:"type"`
	Config       map[string]interface{} `json:"config"`
}

func (o *output) toOutput() (Output, error) {
	switch *o.Type {
	case "webhook":
		return WebHookOutput{
			ID:                *o.ID,
			CollectionID:      *o.CollectionID,
			URL:               o.str("url"),
			BasicAuthUser:     o.str("basicAuthUser"),
			BasicAuthPass:     o.str("basicAuthPass"),
			CustomHeaderName:  o.str("customHeaderName"),
			CustomHeaderValue: o.str("customHeaderValue"),
		}, nil
	case "mqtt":
		return MQTTOutput{
			ID:               *o.ID,
			CollectionID:     *o.CollectionID,
			Endpoint:         o.str("endpoint"),
			DisableCertCheck: o.bool("disableCertCheck"),
			Username:         o.str("username"),
			Password:         o.str("password"),
			ClientID:         o.str("clientId"),
			TopicName:        o.str("topicName"),
		}, nil
	case "ifttt":
		return IFTTTOutput{
			ID:           *o.ID,
			CollectionID: *o.CollectionID,
			Key:          o.str("key"),
			EventName:    o.str("eventName"),
			AsIsPayload:  o.bool("asIsPayload"),
		}, nil
	case "udp":
		return UDPOutput{
			ID:           *o.ID,
			CollectionID: *o.CollectionID,
			Host:         o.str("host"),
			Port:         o.int("port"),
		}, nil
	}
	return nil, fmt.Errorf("unknown output type %q", *o.Type)
}

func (o *output) str(name string) string {
	s, _ := o.Config[name].(string)
	return s
}

func (o *output) bool(name string) bool {
	b, _ := o.Config[name].(bool)
	return b
}

func (o *output) int(name string) int {
	b, _ := o.Config[name].(float64)
	return int(b)
}

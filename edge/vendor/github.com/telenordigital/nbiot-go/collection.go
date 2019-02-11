package nbiot

import (
	"fmt"
	"time"
)

// Collection represents a collection.
type Collection struct {
	CollectionID string            `json:"collectionId"`
	TeamID       *string           `json:"teamId"`
	Tags         map[string]string `json:"tags,omitempty"`
}

// Datapoint represents a data point in a collection
type Datapoint struct {
	Type     string `json:"type"`
	Device   Device `json:"device"`
	Payload  []byte `json:"payload"`
	Received int64  `json:"received"`
}

// Collection gets a collection.
func (c *Client) Collection(id string) (Collection, error) {
	var collection Collection
	err := c.get("/collections/"+id, &collection)
	return collection, err
}

// Collections gets all collections that the user has access to.
func (c *Client) Collections() ([]Collection, error) {
	var collections struct {
		Collections []Collection `json:"collections"`
	}
	err := c.get("/collections", &collections)
	return collections.Collections, err
}

// CreateCollection creates a collection.
func (c *Client) CreateCollection(collection Collection) (Collection, error) {
	err := c.create("/collections", &collection)
	return collection, err
}

// UpdateCollection updates a collection.
func (c *Client) UpdateCollection(collection Collection) (Collection, error) {
	err := c.update("/collections/"+collection.CollectionID, &collection)
	return collection, err
}

// DeleteCollection deletes a collection.
func (c *Client) DeleteCollection(id string) error {
	return c.delete("/collections/" + id)
}

// Data returns all the stored data for the collection
func (c *Client) Data(collectionID string, since time.Time, until time.Time, limit int) ([]Datapoint, error) {
	var data struct {
		Datapoints []Datapoint `json:"messages"`
	}

	var s, u int64

	if !since.IsZero() {
		s = since.UnixNano() / int64(time.Millisecond)
	}

	if !until.IsZero() {
		u = until.UnixNano() / int64(time.Millisecond)
	}

	err := c.get(fmt.Sprintf("/collections/%s/data?since=%d&until=%d&limit=%d", collectionID, s, u, limit), &data)
	return data.Datapoints, err
}

package nbiot

// Team represents a team.
type Team struct {
	TeamID  string            `json:"teamId"`
	Members []Member          `json:"members,omitempty"`
	Tags    map[string]string `json:"tags,omitempty"`
}

// Member is the member element of the Team type
type Member struct {
	UserID        *string `json:"userId"`
	Role          *string `json:"role"`
	Email         *string `json:"email"`
	Phone         *string `json:"phone"`
	VerifiedEmail *bool   `json:"verifiedEmail"`
	VerifiedPhone *bool   `json:"verifiedPhone"`
	ConnectID     *string `json:"connectId"`
}

// Team gets a team.
func (c *Client) Team(id string) (Team, error) {
	var team Team
	err := c.get("/teams/"+id, &team)
	return team, err
}

// Teams gets all teams that the user belongs to.
func (c *Client) Teams() ([]Team, error) {
	var teams struct {
		Teams []Team `json:"teams"`
	}
	err := c.get("/teams", &teams)
	return teams.Teams, err
}

// CreateTeam creates a team.
func (c *Client) CreateTeam(team Team) (Team, error) {
	err := c.create("/teams", &team)
	return team, err
}

// UpdateTeam updates a team.
func (c *Client) UpdateTeam(team Team) (Team, error) {
	err := c.update("/teams/"+team.TeamID, &team)
	return team, err
}

// DeleteTeam deletes a team.
func (c *Client) DeleteTeam(id string) error {
	return c.delete("/teams/" + id)
}

package curseforge

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const apiHost = "https://www.curseforge.com/api/v1"
const gameId = 1

type Client struct {
	apiHost string
	gameId  int
	Client  *http.Client
}

func NewClient() *Client {
	return &Client{
		apiHost: apiHost,
		gameId:  gameId,
		Client:  &http.Client{},
	}
}

type File struct {
	Id           int       `json:"id"`
	Filename     string    `json:"filename"`
	GameVersions []string  `json:"gameVersions"`
	DateCreated  time.Time `json:"dateCreated"`
}

type FileSet struct {
	Data []File `json:"data"`
}

func (c *Client) GetFiles(modId int) FileSet {
	endpoint := fmt.Sprintf("/mods/%d/files", modId)
	req, err := http.NewRequest("GET", c.apiHost+endpoint, nil)
	if err != nil {
		fmt.Errorf("Request is malformed: %w", err)
	}

	req.Header.Set("Accept", "application/json")
	resp, err := c.Client.Do(req)

	if err != nil || resp.StatusCode != 200 {
		fmt.Errorf("Request failed (Status code %s): %w", resp.Status, err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Errorf("Failed to read response body: %w", err)
	}

	fileSet := FileSet{}
	err = json.Unmarshal(body, &fileSet)
	if err != nil {
		fmt.Errorf("Could not map reponse to type FileSet: %w", err)
	}

	return fileSet
}

func (c *Client) DownloadFile(modId int, fileId int) bool {
	endpoint := fmt.Sprintf("/mods/%d/files/%d/download", modId, fileId)
	fmt.Println(endpoint)

	return true
}

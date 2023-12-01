package curseforge

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
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
	Location     string    `json:"location"`
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

func (c *Client) DownloadFile(modId int, file File) File {
	endpoint := fmt.Sprintf("/mods/%d/files/%d/download", modId, file.Id)
	dest := "/tmp/" + file.Filename
	out, err := os.Create(dest)

	if err != nil {
		fmt.Errorf("Failed to create temporary file: %w", err)
	}
	defer func(out *os.File) {
		err := out.Close()
		if err != nil {
			fmt.Errorf("Could not close file: %w", err)
		}
	}(out)

	resp, err := http.Get(c.apiHost + endpoint)

	if err != nil {
		fmt.Errorf("File could not be downloaded: %w", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Errorf("Could not close response body: %w", err)
		}
	}(resp.Body)

	_, writeErr := io.Copy(out, resp.Body)
	if writeErr != nil {
		fmt.Errorf("File could not be copied to target destination (%s): %w", dest, err)
	}

	file.Location = dest

	return file
}

package curseforge

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"golang.org/x/mod/semver"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"
)

const apiHost = "https://www.curseforge.com/api/v1"
const gameId = 1

type Client struct {
	apiHost      string
	gameId       int
	DownloadPath string
	Client       *http.Client
}

var c *Client

func init() {
	c = NewClient()
}

func NewClient() *Client {
	return &Client{
		apiHost:      apiHost,
		gameId:       gameId,
		DownloadPath: "/tmp",
		Client:       &http.Client{},
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

func GetFiles(modId int) (FileSet, error) {
	return c.GetFiles(modId)
}

func (c *Client) GetFiles(modId int) (FileSet, error) {
	endpoint := fmt.Sprintf("/mods/%d/files", modId)
	req, err := http.NewRequest("GET", c.apiHost+endpoint, nil)
	if err != nil {
		return FileSet{}, fmt.Errorf("Request is malformed: %w", err)
	}

	req.Header.Set("Accept", "application/json")
	resp, err := c.Client.Do(req)

	if err != nil || resp.StatusCode != 200 {
		return FileSet{}, fmt.Errorf("Request failed (Status code %s): %w", resp.Status, err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return FileSet{}, fmt.Errorf("Failed to read response body: %w", err)
	}

	fileSet := FileSet{}
	err = json.Unmarshal(body, &fileSet)
	if err != nil {
		return FileSet{}, fmt.Errorf("Could not map reponse to type FileSet: %w", err)
	}

	return fileSet, nil
}

func DownloadFile(modId int, file File) (File, error) {
	return c.DownloadFile(modId, file)
}

func (c *Client) DownloadFile(modId int, file File) (File, error) {
	endpoint := fmt.Sprintf("/mods/%d/files/%d/download", modId, file.Id)
	dest := path.Clean(c.DownloadPath + "/" + file.Filename)
	out, err := os.Create(dest)

	if err != nil {
		return File{}, fmt.Errorf("Failed to create temporary file: %w", err)
	}
	defer func(out *os.File) {
		err := out.Close()
		if err != nil {
			panic(fmt.Errorf("Could not close file: %w", err))
		}
	}(out)

	resp, err := http.Get(c.apiHost + endpoint)

	if err != nil {
		return File{}, fmt.Errorf("File could not be downloaded: %w", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(fmt.Errorf("Could not close response body: %w", err))
		}
	}(resp.Body)

	_, writeErr := io.Copy(out, resp.Body)
	if writeErr != nil {
		return File{}, fmt.Errorf("File could not be copied to target destination (%s): %w", dest, err)
	}

	file.Location = dest

	return file, nil
}

func InstallAddon(file File, dest string) error {
	return c.InstallAddon(file, dest)
}

func (c *Client) InstallAddon(file File, dest string) error {
	addon, err := zip.OpenReader(file.Location)
	if err != nil {
		return fmt.Errorf("Failed to read Zip file: %w", err)
	}
	defer addon.Close()

	destination, err := filepath.Abs(dest)
	if err != nil {
		return fmt.Errorf("Install file path is invalid %w", err)
	}

	for _, zipFile := range addon.File {
		zippedFile, err := zipFile.Open()
		if err != nil {
			return fmt.Errorf("Failed to open zipped file: %w", err)
		}
		defer zippedFile.Close()

		targetFilePath := filepath.Join(destination, zipFile.Name)
		if zipFile.FileInfo().IsDir() {
			os.MkdirAll(targetFilePath, zipFile.Mode())
		} else {
			targetFile, err := os.Create(targetFilePath)
			if err != nil {
				return fmt.Errorf("Failed to create target file: %w", err)
			}
			defer targetFile.Close()

			_, err = io.Copy(targetFile, zippedFile)
			if err != nil {
				return fmt.Errorf("Failed to copy content to target file: %w", err)
			}
		}
	}

	return nil
}

func (c *Client) NegotiateFile(files FileSet) File {
	latest := File{GameVersions: make([]string, 1)}
	fmt.Println(latest)
	for _, file := range files.Data {
		if semver.Compare("v"+file.GameVersions[0], "v"+latest.GameVersions[0]) == 1 {
			latest = file
		}
	}
	return latest
}

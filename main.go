package main

import (
	"encoding/json" // For JSON operations
	"flag"
	"fmt"
	"io/ioutil" // If you're using Go version < 1.16. Otherwise, use "io" and "os" as needed.
	"net/http"  // For HTTP requests
	"os"        // For writing to stdout and potentially for replacing ioutil if using Go >= 1.16
	"github.com/cheggaaa/pb/v3"
	"github.com/olekukonko/tablewriter"
	"github.com/common-nighthawk/go-figure"
)

type Project struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Updated functions to include gitLabURL as a parameter
func getGitlabProjects(gitLabURL string, groupID int, token string) ([]Project, error) {
	var projects []Project
	url := fmt.Sprintf("%s/api/v4/groups/%d/projects", gitLabURL, groupID)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	query := req.URL.Query()
	query.Add("per_page", "100")
	query.Add("simple", "true")
	req.URL.RawQuery = query.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return projects, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		err = json.NewDecoder(resp.Body).Decode(&projects)
	} else {
		return projects, fmt.Errorf("error retrieving projects: %s", resp.Status)
	}
	return projects, err
}

func getComposerFile(gitLabURL string, projectID int, token string) (string, error) {
	url := fmt.Sprintf("%s/api/v4/projects/%d/repository/files/composer.json/raw", gitLabURL, projectID)
	branches := []string{"main", "master"}

	for _, branch := range branches {
		req, _ := http.NewRequest("GET", url, nil)
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
		query := req.URL.Query()
		query.Add("ref", branch)
		req.URL.RawQuery = query.Encode()

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return "", err
		}
		defer resp.Body.Close()

		if resp.StatusCode == 200 {
			body, err := ioutil.ReadAll(resp.Body)
			return string(body), err
		}
	}
	return "", fmt.Errorf("file not found")
}

func parseSymfonyVersion(composerContent string) (string, error) {
	var data struct {
		Require map[string]string `json:"require"`
	}
	if err := json.Unmarshal([]byte(composerContent), &data); err != nil {
		return "", fmt.Errorf("failed to parse composer.json: %v", err)
	}

	// Assuming framework-bundle is the package we're interested in, adjust as necessary
	version, exists := data.Require["symfony/framework-bundle"]
	if !exists {
		return "Not specified", nil
	}
	return version, nil
}

func main() {
	gitLabURL := flag.String("gitlab_url", "https://gitlab.com", "The GitLab instance URL")
	groupID := flag.Int("group_id", 0, "The GitLab group ID")
	token := flag.String("token", "", "Your GitLab personal access token")
	flag.Parse()

	if *groupID == 0 || *token == "" {
		fmt.Println("group_id and token are required")
		return
	}

	projects, err := getGitlabProjects(*gitLabURL, *groupID, *token)
	if err != nil {
		fmt.Printf("Error fetching projects: %v\n", err)
		return
	}

	// ASCII Art
	myFigure := figure.NewFigure("Check Symfony", "", true)
	myFigure.Print()

	// Start Message
	fmt.Println("\n")
	fmt.Println("Fetching Symfony versions for projects in GitLab group...")
	fmt.Println("\n")

	// Initialize progress bar
	bar := pb.StartNew(len(projects))

	var symfonyProjects [][]string
	var notSymfonyProjects []string

	for _, project := range projects {
		composerContent, err := getComposerFile(*gitLabURL, project.ID, *token)
		if err != nil {
			notSymfonyProjects = append(notSymfonyProjects, project.Name)
			bar.Increment()
			continue
		}
		version, err := parseSymfonyVersion(composerContent)
		if err != nil {
			notSymfonyProjects = append(notSymfonyProjects, project.Name)
		} else {
			symfonyProjects = append(symfonyProjects, []string{project.Name, version})
		}
		bar.Increment()
	}

	bar.Finish()

	// Display summary table
	fmt.Println("\n")
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Project Name", "Symfony Version"})
	table.SetBorder(false) // Set to true if you want table with border
	table.AppendBulk(symfonyProjects) // Add Bulk Data
	table.Render()

	// Display non-Symfony projects
	fmt.Println("\n")
	fmt.Println("\nNon-Symfony Projects:")
	for _, projectName := range notSymfonyProjects {
		fmt.Println(projectName)
	}
}

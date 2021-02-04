package upload

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/urfave/cli/v2"

	"github.com/covergates/covergates/cmd/cli/modules"
	"github.com/covergates/covergates/core"
	"github.com/covergates/covergates/modules/git"
	"github.com/covergates/covergates/modules/util"
	"github.com/covergates/covergates/service/coverage"
)

// Command for upload report
var Command = &cli.Command{
	Name:      "upload",
	Usage:     "upload coverage report",
	ArgsUsage: "report",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "report",
			Usage:    "report id",
			EnvVars:  []string{"REPORT_ID"},
			Value:    "",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "type",
			Usage:    "report type",
			Value:    "",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "branch",
			Usage:    "branch to upload the report",
			EnvVars:  []string{"GITHUB_HEAD_REF", "DRONE_SOURCE_BRANCH"},
			Value:    "",
			Required: false,
		},
		&cli.StringSliceFlag{
			Name:     "skip-file",
			Usage:    "skip regexp compliant files",
			Required: false,
		},
	},
	Action: upload,
}

func upload(c *cli.Context) error {
	if c.NArg() <= 0 {
		_ = cli.ShowCommandHelp(c, "upload")
		return fmt.Errorf("report path is required")
	}

	data, err := findReportData(c.Context, c.String("type"), c.Args().First())
	if err != nil {
		return err
	}

	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	gitService := &git.Service{}
	repo, err := gitService.PlainOpen(c.Context, cwd)
	if err != nil {
		return err
	}

	branch := c.String("branch")
	if branch == "" {
		branch = repo.Branch()
	}

	files, err := repo.ListAllFiles(repo.HeadCommit())
	if err != nil {
		return err
	}

	files, err = removeSkipFile(files, c.StringSlice("skip-file"))
	if err != nil {
		return err
	}

	filesData, err := json.Marshal(files)
	if err != nil {
		return err
	}

	form := util.FormData{
		"type":   c.String("type"),
		"commit": repo.HeadCommit(),
		"ref":    branch,
		"files":  string(filesData),
		"root":   repo.Root(),
		"file": util.FormFile{
			Name: "report",
			Data: data,
		},
	}

	url := fmt.Sprintf(
		"%s/reports/%s",
		c.String("url"),
		c.String("report"),
	)

	log.Printf("upload commit %s, %s\n", repo.HeadCommit(), c.String("type"))

	request, err := util.CreatePostFormRequest(url, form)
	if err != nil {
		return nil
	}
	respond, err := modules.GetHTTPClient(c).Do(request)
	if err != nil {
		return err
	}

	var text []byte
	defer func() {
		_ = respond.Body.Close()
		if respond.StatusCode >= 400 {
			log.Fatal(string(text))
		} else {
			log.Println(string(text))
		}
	}()
	text, err = ioutil.ReadAll(respond.Body)
	return err
}

func findReportData(ctx context.Context, reportType, path string) ([]byte, error) {
	t := core.ReportType(reportType)
	service := &coverage.Service{}
	report, err := service.Find(ctx, t, path)
	if err != nil {
		return nil, err
	}
	r, err := service.Open(ctx, t, report)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(r)
}

func removeSkipFile(files []string, patterns []string) ([]string, error) {
	var result []string
	for _, s := range files {
		matched, err := matchPatterns(s, patterns)
		if err != nil {
			return nil, err
		}

		if !matched {
			result = append(result, s)
		}
	}

	return result, nil
}

func matchPatterns(s string, patterns []string) (bool, error) {
	var matched bool
	for _, pattern := range patterns {
		p := normalizePathInRegex(pattern)
		patternRe, err := regexp.Compile(p)
		if err != nil {
			return false, err
		}

		matched = patternRe.MatchString(s)
	}

	return matched, nil
}

var separatorToReplace = regexp.QuoteMeta(string(filepath.Separator))

func normalizePathInRegex(path string) string {
	if filepath.Separator == '/' {
		return path
	}

	// This replacing should be safe because "/" are disallowed in Windows
	// https://docs.microsoft.com/ru-ru/windows/win32/fileio/naming-a-file
	return strings.ReplaceAll(path, "/", separatorToReplace)
}

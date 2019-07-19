package main

import (
	"fmt"
	"log"
	"os"
	"encoding/base64"

	"github.com/urfave/cli"
	"./lib/teams"
)

var (
	version = "0.0.0"
	build   = "0"
)

type Icons struct {
	OK string
	Error string
}

func icon(iconPath string) string {
	icon, _ := Asset(iconPath)
	return "data:image/png;base64," + base64.StdEncoding.EncodeToString(icon);
}

var icons = Icons{
	OK: icon("assets/icon-ok.png"),
	Error: icon("assets/icon-error.png"),
}


func main() {

	app := cli.NewApp()
	app.Name = "teams plugin"
	app.Usage = "teams plugin"
	app.Action = run
	app.Version = fmt.Sprintf("%s+%s", version, build)
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "webhook",
			Usage:  "teams webhook url",
			EnvVar: "TEAMS_WEBHOOK,PLUGIN_WEBHOOK",
	
		},
		cli.StringFlag{
			Name:   "repo.owner",
			Usage:  "repository owner",
			EnvVar: "DRONE_REPO_OWNER",
		},
		cli.StringFlag{
			Name:   "repo.name",
			Usage:  "repository name",
			EnvVar: "DRONE_REPO_NAME",
		},
		cli.StringFlag{
			Name:   "commit.sha",
			Usage:  "git commit sha",
			EnvVar: "DRONE_COMMIT_SHA",
			Value:  "00000000",
		},
		cli.StringFlag{
			Name:   "commit.ref",
			Value:  "refs/heads/master",
			Usage:  "git commit ref",
			EnvVar: "DRONE_COMMIT_REF",
		},
		cli.StringFlag{
			Name:   "commit.branch",
			Value:  "master",
			Usage:  "git commit branch",
			EnvVar: "DRONE_COMMIT_BRANCH",
		},
		cli.StringFlag{
			Name:   "commit.author",
			Usage:  "git author name",
			EnvVar: "DRONE_COMMIT_AUTHOR",
		},
		cli.StringFlag{
			Name:   "commit.pull",
			Usage:  "git pull request",
			EnvVar: "DRONE_PULL_REQUEST",
		},
		cli.StringFlag{
			Name:   "commit.message",
			Usage:  "commit message",
			EnvVar: "DRONE_COMMIT_MESSAGE",
		},
		cli.StringFlag{
			Name:   "build.event",
			Value:  "push",
			Usage:  "build event",
			EnvVar: "DRONE_BUILD_EVENT",
		},
		cli.IntFlag{
			Name:   "build.number",
			Usage:  "build number",
			Value:  0,
			EnvVar: "DRONE_BUILD_NUMBER",
		},
		cli.StringFlag{
			Name:   "build.status",
			Usage:  "build status",
			Value:  "success",
			EnvVar: "DRONE_BUILD_STATUS",
		},
		cli.StringFlag{
			Name:   "build.link",
			Usage:  "build link",
			EnvVar: "DRONE_BUILD_LINK",
		},
		cli.Int64Flag{
			Name:   "build.started",
			Usage:  "build started",
			EnvVar: "DRONE_BUILD_STARTED",
		},
		cli.Int64Flag{
			Name:   "build.created",
			Usage:  "build created",
			EnvVar: "DRONE_BUILD_CREATED",
		},
		cli.StringFlag{
			Name:   "build.tag",
			Usage:  "build tag",
			EnvVar: "DRONE_TAG",
		},
		cli.StringFlag{
			Name:   "build.deployTo",
			Usage:  "environment deployed to",
			EnvVar: "DRONE_DEPLOY_TO",
		},
		cli.Int64Flag{
			Name:   "job.started",
			Usage:  "job started",
			EnvVar: "DRONE_JOB_STARTED",
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(c *cli.Context) error {

	api, err := teams.New(c.String("webhook"))

	var title = "Build " 
	var icon string
	
	if c.String("build.status") == "success" {
		title = title + "Completed"
		icon = icons.OK
	} else {
		title = title + "Failed"
		icon = icons.Error
	}

	resp, err := api.PerformAPIRequest(&teams.APIRequest{
		Summary: title,
		Sections: []teams.APISection{
			{
				Title: title,
				SubTitle: string("\\#" + c.String("build.number") + " " + c.String("repo.name")),
				Image: icon,
				Facts: []teams.APIFact{
					{
						Name: "Branch",
						Value: c.String("commit.branch"),
					},
					{
						Name: "Author",
						Value: c.String("commit.author"),
					},
					{
						Name: "Message",
						Value: c.String("commit.message"),
					},
					{
						Name: "Started",
						Value: c.String("build.started"),
					},
				},
			},
		},
		PotentialActions: []teams.APIPotentialAction{
			{
				Type: "OpenUri",
				Name: "View Result",
				Targets: []teams.APIOpenUriTarget{
					{
						OS: "default",
						URI:  c.String("build.link"),
					},
				},
			},
		},
		
	})

	log.Println(resp);

	return err

}

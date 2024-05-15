package internal

import (
	"strconv"
	"strings"
)

type RootProjectAction interface {
	setupEnv()
	bump(releaseType ReleaseType)
}

type RootProject struct {
	Workspace string
	Name      string
	//currentVersion string
	//currentTag     string
}

func (p RootProject) SetupEnv() {

}

func (p RootProject) Build() {

}

func (p RootProject) getLatestTag() string {
	out, err := RunCommand("git describe --tags --abbrev=0", p.Workspace)
	if err != nil {
		return ""
	}
	return strings.TrimSpace(out)
}

func (p RootProject) getCommitCount(tag string, subdir *string) int {

	command := "git rev-list --count " + tag + "..HEAD --pretty=oneline --count"

	if subdir != nil {
		command += " " + *subdir
	}

	out, err := RunCommand(command, p.Workspace)
	if err != nil {
		return -1
	}

	num, err := strconv.Atoi(strings.TrimSpace(out))

	if err != nil {
		return -1
	}

	return num
}

func (p RootProject) gitCommit(message string) {
	_, _ = RunCommand("git commit -m '"+message+"' -a", p.Workspace)
}

func (p RootProject) gitTag(tag string) {
	_, _ = RunCommand("git tag -a "+tag+" -m '"+tag+"'", p.Workspace)
	_, _ = RunCommand("git push origin "+tag+" --no-verify", p.Workspace)
}

func (p RootProject) gitPush(string) {
	_, _ = RunCommand("git push --no-verify", p.Workspace)
}

package app

import "runtime"

var (
    majorVer  = "0"
    minorVer  = "0"
    patchVer  = "0"
    gitCommit string
)

type buildInfo struct {
    Version   string
    GitCommit string
    GoVersion string
}

func VersionSlim() string {
    return majorVer + "." + minorVer + "." + patchVer
}

func Version() string {
    return Name + " v" + VersionSlim()
}

func BuildInfo() buildInfo {
    return buildInfo{
        Version:   VersionSlim(),
        GitCommit: gitCommit,
        GoVersion: runtime.Version(),
    }
}

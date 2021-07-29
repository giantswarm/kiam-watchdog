package project

var (
	description = "Checks health of kiam and restarts it when it's failing"
	gitSHA      = "n/a"
	name        = "kiam-watchdog"
	source      = "https://github.com/giantswarm/kiam-watchdog"
	version     = "0.1.2-dev"
)

func Description() string {
	return description
}

func GitSHA() string {
	return gitSHA
}

func Name() string {
	return name
}

func Source() string {
	return source
}

func Version() string {
	return version
}

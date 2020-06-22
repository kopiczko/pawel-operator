package project

var (
	description = "The template-operator does something."
	gitSHA      = "n/a"
	name        = "template-operator"
	source      = "https://github.com/giantswarm/template-operator"
	version     = "1.0.0"
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

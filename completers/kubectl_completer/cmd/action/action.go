package action

import (
	"os/exec"
	"strings"

	"github.com/rsteube/carapace"
)

func ActionApiResources() carapace.Action {
	return carapace.ActionCallback(func(args []string) carapace.Action {
		if output, err := exec.Command("kubectl", "api-resources", "--output=name", "--cached").Output(); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			lines := strings.Split(string(output), "\n")
			for index, line := range lines {
				lines[index] = strings.SplitN(line, ".", 2)[0]
			}
			return carapace.ActionValues(lines[:len(lines)-1]...)
		}
	})
}

func ActionResources(types string) carapace.Action {
	return carapace.ActionCallback(func(args []string) carapace.Action {
		if output, err := exec.Command("kubectl", "get", "-o", "go-template={{range .items}}{{.metadata.name}}\n{{.kind}}\n{{end}}", types).Output(); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			lines := strings.Split(string(output), "\n")
			return carapace.ActionValuesDescribed(lines[:len(lines)-1]...)
		}
	})
}

func ActionDryRunOptions() carapace.Action {
	return carapace.ActionValues("none", "server", "client")
}

func ActionOutputFormats() carapace.Action {
	return carapace.ActionValues("json", "yaml", "name", "go-template", "go-template-file", "template", "templatefile", "jsonpath", "jsonpath-as-json", "jsonpath-file")
}

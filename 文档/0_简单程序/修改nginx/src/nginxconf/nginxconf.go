package nginxconf

import (
	"bytes"
	"fmt"
	"nginx_reload"
	"os/exec"
)

func Nginxconf(path, name string) (string, string, error) {
	var outMsg, errMsg bytes.Buffer
	cmd := exec.Command("/bin/sh", "-c", fmt.Sprintf(`sed -i 's/404.html;/&\nif  (\$request_uri  \
~* "%s") {return 403;}/' %s`, name, path))
	cmd.Stdout = &outMsg
	cmd.Stderr = &errMsg
	err := cmd.Run()
	return outMsg.String(), errMsg.String(), err

	nginx_reload.Reload()
}

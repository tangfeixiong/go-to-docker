/*
   Inspired from
    https://github.com/helm/helm-classic/blob/master/kubectl/get.go
*/
// package kubectl
package manipulate

func python_check(wd string, opts ...string) ([]byte, error) {
	cmd := command(opts...)

	if len(wd) != 0 {
		cmd.Dir = wd
	}

	Path = "python"
	return cmd.CombinedOutput()
}

func (r RealRunner) Python_00_check(wd string, opts ...string) ([]byte, error) {
	return python_check(wd, opts...)
}

func (r RealRunner) Web1_2_check(wd string, opts ...string) ([]byte, error) {
	return python_check(wd, opts...)
}

func webXcheck(wd, f string) ([]byte, error) {
	args := []string{f}

	cmd := command(args...)

	if len(wd) != 0 {
		cmd.Dir = wd
	}
	return cmd.CombinedOutput()
}

func (r RealRunner) Web1Check(wd string) ([]byte, error) {
	return webXcheck(wd, "web1check.py")
}

func (r RealRunner) Web2Check(wd string) ([]byte, error) {
	return webXcheck(wd, "web2check.py")
}

//// Get returns Kubernetes resources
//func (r RealRunner) Get(stdin []byte, ns string) ([]byte, error) {
//	args := []string{"get", "-f", "-"}

//	if ns != "" {
//		args = append([]string{"--namespace=" + ns}, args...)
//	}
//	cmd := command(args...)
//	assignStdin(cmd, stdin)

//	return cmd.CombinedOutput()
//}

//// Get returns the commands to kubectl
//func (r PrintRunner) Get(stdin []byte, ns string) ([]byte, error) {
//	args := []string{"get", "-f", "-"}

//	if ns != "" {
//		args = append([]string{"--namespace=" + ns}, args...)
//	}
//	cmd := command(args...)
//	assignStdin(cmd, stdin)

//	return []byte(cmd.String()), nil
//}

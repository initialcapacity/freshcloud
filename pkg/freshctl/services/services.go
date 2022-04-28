package services

import "github.com/initialcapacity/freshcloud/pkg/templatesupport"

func InstallContourCmd(resourcesDirectory, domain string) []string {
	name := "install_contour"
	data := struct {
		Domain string
	}{
		Domain: domain,
	}
	return []string{
		templatesupport.Parse(resourcesDirectory, name, data),
	}
}

func InstallCertManagerCmd() []string {
	return []string{"echo todo"}
}

func InstallHarborCmd() []string {
	return []string{"echo todo"}
}

func InstallConcourseCmd() []string {
	return []string{"echo todo"}
}

func InstallKpackCmd() []string {
	return []string{"echo todo"}
}

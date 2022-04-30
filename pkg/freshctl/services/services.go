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

func InstallCertManagerCmd(resourcesDirectory, emailAddress string) []string {
	name := "install_cert_manager"
	data := struct {
		EmailAddress string
	}{
		EmailAddress: emailAddress,
	}
	return []string{
		templatesupport.Parse(resourcesDirectory, name, data),
	}
}

func InstallHarborCmd(resourcesDirectory, domain, emailAddress, password string) []string {
	name := "install_harbor"
	data := struct {
		Domain       string
		EmailAddress string
		Password     string
	}{
		Domain:       domain,
		EmailAddress: emailAddress,
		Password:     password,
	}
	return []string{
		templatesupport.Parse(resourcesDirectory, name, data),
	}
}

func InstallConcourseCmd(resourcesDirectory, domain, password string) []string {
	name := "install_concourse"
	data := struct {
		Domain   string
		Password string
	}{
		Domain:   domain,
		Password: password,
	}
	return []string{
		templatesupport.Parse(resourcesDirectory, name, data),
	}
}

func InstallKpackCmd() []string {
	return []string{"echo todo"}
}

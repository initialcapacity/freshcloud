package services

import "github.com/initialcapacity/freshcloud/pkg/templatesupport"

func InstallContourCmd(resourcesLocation string, envMap map[string]string) []string {
	return []string{templatesupport.Parse(resourcesLocation, "services_install_contour", envMap)}
}

func InstallCertManagerCmd(resourcesLocation string, envMap map[string]string) []string {
	return []string{templatesupport.Parse(resourcesLocation, "services_install_cert_manager", envMap)}
}

func InstallHarborCmd(resourcesLocation string, envMap map[string]string) []string {
	return []string{templatesupport.Parse(resourcesLocation, "services_install_harbor", envMap)}
}

func InstallConcourseCmd(resourcesLocation string, envMap map[string]string) []string {
	return []string{templatesupport.Parse(resourcesLocation, "services_install_concourse", envMap)}
}

func InstallKpackCmd(resourcesLocation string, envMap map[string]string) []string {
	return []string{templatesupport.Parse(resourcesLocation, "services_install_kpack", envMap)}
}

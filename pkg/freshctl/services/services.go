package services

import "github.com/initialcapacity/freshcloud/pkg/templatesupport"

func InstallContourCmd(resourcesDirectory string, envMap map[string]string) []string {
	return []string{templatesupport.Parse(resourcesDirectory, "services_install_contour", envMap)}
}

func InstallCertManagerCmd(resourcesDirectory string, envMap map[string]string) []string {
	return []string{templatesupport.Parse(resourcesDirectory, "services_install_cert_manager", envMap)}
}

func InstallHarborCmd(resourcesDirectory string, envMap map[string]string) []string {
	return []string{templatesupport.Parse(resourcesDirectory, "services_install_harbor", envMap)}
}

func InstallConcourseCmd(resourcesDirectory string, envMap map[string]string) []string {
	return []string{templatesupport.Parse(resourcesDirectory, "services_install_concourse", envMap)}
}

func InstallKpackCmd(resourcesDirectory string, envMap map[string]string) []string {
	return []string{templatesupport.Parse(resourcesDirectory, "services_install_kpack", envMap)}
}

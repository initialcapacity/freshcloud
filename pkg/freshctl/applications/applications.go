package applications

import "github.com/initialcapacity/freshcloud/pkg/templatesupport"

func PushImageCmd(resourcesDirectory, registryDomain, password, app, image string) []string {
	name := "push_image"
	data := struct {
		RegistryDomain string
		Password       string
		App            string
		Image          string
	}{
		RegistryDomain: registryDomain,
		Password:       password,
		App:            app,
		Image:          image,
	}
	return []string{
		templatesupport.Parse(resourcesDirectory, name, data),
	}
}

func DeployAppCmd(resourcesDirectory, registryDomain, password, app, appDomain, image, path string) []string {
	name := "deploy_app"
	data := struct {
		RegistryDomain       string
		Password             string
		App                  string
		AppDomain            string
		Image                string
		AppConfigurationPath string
	}{
		RegistryDomain:       registryDomain,
		Password:             password,
		App:                  app,
		AppDomain:            appDomain,
		Image:                image,
		AppConfigurationPath: path,
	}

	return []string{
		templatesupport.Parse(resourcesDirectory, name, data),
	}
}

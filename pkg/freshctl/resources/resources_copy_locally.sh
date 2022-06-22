resources=( 'applications_deploy_app.sh'
'applications_push_image.sh'
'aws_clusters_configure.sh'
'aws_clusters_configure_registry.sh'
'aws_clusters_create.sh'
'aws_clusters_delete.sh'
'azure_clusters_configure.sh'
'azure_clusters_configure_registry.sh'
'azure_clusters_create.sh'
'azure_clusters_create_resource_group.sh'
'azure_clusters_delete_resource_group.sh'
'google_clusters_configure.sh'
'google_clusters_configure_registry.sh'
'google_clusters_create.sh'
'google_clusters_create_service_account.sh'
'google_clusters_delete.sh'
'google_clusters_enable_services.sh'
'google_clusters_list.sh'
'pipelines_delete_pipeline.sh'
'pipelines_deploy_pipeline.sh'
'pipelines_push_build_image.sh'
'services_install_cert_manager.sh'
'services_install_concourse.sh'
'services_install_gitlab.sh'
'services_install_contour.sh'
'services_install_harbor.sh'
'services_install_kpack.sh' )
mkdir -p local_resources
for resource in ${resources[@]}; do
  echo -n "Installing: local_resources/${resource}: "
  curl -s -L https://raw.githubusercontent.com/initialcapacity/freshcloud/main/pkg/freshctl/resources/${resource} -o local_resources/${resource}
  if [ $? != 0 ]; then
    echo "Failed"
  else
    echo "Success"
  fi
done

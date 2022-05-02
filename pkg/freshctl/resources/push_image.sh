curl --user "admin:{{.Password}}" -X POST \
  https://registry.{{.RegistryDomain}}/api/v2.0/projects \
  -H "Content-type: application/json" --data \
  '{ "project_name": "'{{.App}}'", "metadata":
   { "auto_scan": "true", "enable_content_trust":
     "false", "prevent_vul": "false", "public":
     "true", "reuse_sys_cve_whitelist": "true",
     "severity": "high" }
   }'
docker login -u admin -p {{.Password}} https://registry.{{.RegistryDomain}}
docker tag {{.Image}} registry.{{.RegistryDomain}}/{{.App}}/{{.Image}}:latest
docker push registry.{{.RegistryDomain}}/{{.App}}/{{.Image}}:latest
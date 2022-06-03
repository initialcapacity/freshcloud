eksctl create cluster -f -
apiVersion: eksctl.io/v1alpha5
kind: ClusterConfig
metadata:
  name: {{index . "AWS_CLUSTER_NAME" }}
  region: {{index . "AWS_REGION" }}
nodeGroups:
  - name: controllers
    labels: { role: controllers }
    instanceTypes:
    - "t3.small"
    desiredCapacity: 2
    iam:
      withAddonPolicies:
        certManager: true
        albIngress: true
    taints:
      controllers: "true:NoSchedule"
managedNodeGroups:
  - name: workers
#! /bin/bash


kubectl apply -f cluster-configuration.yaml 
kubectl apply -f kubesphere-installer.yaml
# fix user account not activated bug
kubectl apply -f https://raw.githubusercontent.com/kubesphere/ks-installer/2c4b479ec65110f7910f913734b3d069409d72a8/roles/ks-core/prepare/files/ks-init/users.iam.kubesphere.io.yaml
kubectl apply -f https://raw.githubusercontent.com/kubesphere/ks-installer/2c4b479ec65110f7910f913734b3d069409d72a8/roles/ks-core/prepare/files/ks-init/webhook-secret.yaml
kubectl -n kubesphere-system rollout restart deploy ks-controller-manager

## 日常开发记录信息
> Kubernetes安装成功提示信息
```shell
Your Kubernetes control-plane has initialized successfully!

To start using your cluster, you need to run the following as a regular user:

mkdir -p $HOME/.kube
sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
sudo chown $(id -u):$(id -g) $HOME/.kube/config

Alternatively, if you are the root user, you can run:

export KUBECONFIG=/etc/kubernetes/admin.conf

You should now deploy a pod network to the cluster.
Run "kubectl apply -f [podnetwork].yaml" with one of the options listed at:
https://kubernetes.io/docs/concepts/cluster-administration/addons/

You can now join any number of the control-plane node running the following command on each as root:

kubeadm join 172.17.0.1:6443 --token abcdef.0123456789abcdef \
--discovery-token-ca-cert-hash sha256:b2bc02ba38bc90838e670677a69b93a0d4e818d45c7645e4ed7305c13274ee23 \
--control-plane --certificate-key 9ed56eb5bef58d28ba6e6f490e0b39958415194dd4a3672df3094f95ded36a72

Please note that the certificate-key gives access to cluster sensitive data, keep it secret!
As a safeguard, uploaded-certs will be deleted in two hours; If necessary, you can use
"kubeadm init phase upload-certs --upload-certs" to reload certs afterward.

Then you can join any number of worker nodes by running the following on each as root:

kubeadm join 172.17.0.1:6443 --token abcdef.0123456789abcdef \
--discovery-token-ca-cert-hash sha256:b2bc02ba38bc90838e670677a69b93a0d4e818d45c7645e4ed7305c13274ee23
```


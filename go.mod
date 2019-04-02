require (
	bitbucket.org/ww/goautoneg v0.0.0-20120707110453-75cd24fc2f2c // indirect
	github.com/google/gofuzz v0.0.0-20170612174753-24818f796faf
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/spf13/cobra v0.0.3
	k8s.io/apimachinery v0.0.0-20190402200122-9d32a9f7dc5b
	k8s.io/apiserver v0.0.0-20190319190228-a4358799e4fe
	k8s.io/client-go v0.0.0-20190402200509-a625628b4097
	k8s.io/component-base v0.0.0-20190402201002-d41b764c175d
	k8s.io/klog v0.2.1-0.20190311220638-291f19f84ceb
	k8s.io/kube-openapi v0.0.0-20190320154901-c59034cc13d5 // indirect
	k8s.io/utils v0.0.0-20190308190857-21c4ce38f2a7 // indirect
)

replace (
	k8s.io/api => github.com/kubernetes-nightly/api v0.0.0-20190402200219-0d432f26bdee
	k8s.io/apimachinery => github.com/kubernetes-nightly/apimachinery v0.0.0-20190402200122-9d32a9f7dc5b
	k8s.io/apiserver => github.com/kubernetes-nightly/apiserver v0.0.0-20190402201334-a156787b8b91
	k8s.io/client-go => github.com/kubernetes-nightly/client-go v0.0.0-20190402200509-a625628b4097
	k8s.io/component-base => github.com/kubernetes-nightly/component-base v0.0.0-20190402201002-d41b764c175d
)

module github.com/programming-kubernetes/custom-apiserver

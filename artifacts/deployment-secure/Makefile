SHELL = bash

OUTPUT := serving-cert-secret.yaml v1alpha1-apiservice.yaml v1beta1-apiservice.yaml

all: tls.key tls.crt $(OUTPUT)

tls.key tls.crt:
		openssl req -new -x509 -subj "/CN=api.pizza-apiserver.svc" -nodes -newkey rsa:4096 -keyout tls.key -out tls.crt -days 365

$(OUTPUT): tls.key tls.crt
		sed 's,CERT,$(shell base64 tls.crt),;s,KEY,$(shell base64 tls.key),' $@.template > $@

clean:
		rm -f tls.key tls.crt $(OUTPUT)

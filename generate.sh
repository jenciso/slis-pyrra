#!/bin/bash

if [[ $1 == apps ]]; then
go run main.go apps -t all -a ../manifests/apps/all.yaml -o ../manifests/apps/all.yaml
fi

if [[ $1 == slis ]]; then
for i in $(yq .[].name < ../manifests/apps/all.yaml); do go run main.go slis -n $i -a ../manifests/apps/all.yaml -o ../manifests/slis/$i.yaml ; done
fi

if [[ $1 == slos ]]; then
for i in $(yq .[].name < ../manifests/apps/all.yaml); do go run main.go slos -s ../manifests/slis/$i.yaml -o ../manifests/slos/$i.yaml ; done
fi

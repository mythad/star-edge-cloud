#!/bin/bash

rm -rf cloud
mkdir -p cloud/

cd ../cloud/caas/

mvn clean package  -Dmaven.test.skip=true
mv target/lib/ ../../deploy/cloud/
mv target/caas* ../../deploy/cloud/


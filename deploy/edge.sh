#!/bin/bash
rm -rf edge
mkdir -p edge/plugins/device/
mkdir -p edge/plugins/extension/
mkdir -p edge/data/sqlite3/
mkdir -p edge/data/badger/
mkdir -p edge/compile/
mkdir -p edge/conf/
cp -r ../edge/core/website/ edge/

cd ../edge/core/
go build main.go
mv main ../../deploy/edge/core
cp ./conf/conf.env ../../deploy/edge/conf/

cd ../store/
go build main.go
mv main ../../deploy/edge/store

cd ../log/
go build main.go
mv main ../../deploy/edge/log
cp ./conf/log_conf.env ../../deploy/edge/conf/

cd ../rules_engine/
go build main.go
mv main ../../deploy/edge/rules_engine
cp ./conf/rules_engine_conf.env ../../deploy/edge/conf/

cd ../scheduler/
go build main.go
mv main ../../deploy/edge/scheduler
cp ./conf/scheduler_conf.env ../../deploy/edge/conf/

cd ../device/
go build main.go
mv main ../../deploy/edge/compile/device

cd ../extension/
go build main.go
mv main ../../deploy/edge/compile/extension

cd ../../deploy/edge/

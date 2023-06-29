budget:
	 cd ../../src/stupidestCache; go test -bench=. -count 25 -run=^# | tee s.csv

minimal:
	go run main.go testdata/minimal.csv 
fail:
	-go run main.go testdata/fail.csv 

curl:
	go run main.go --daemonic &
	if [ $$? -ne 0 ]; then echo "main.go returned $$?"; exit 1; fi
	curl localhost:8080 


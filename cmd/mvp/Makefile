default:
	cd ../../src/mvp; go test

budget: # just the benchmark
	 cd ../../src/mvp; go test -bench=. -count 25 -run=BenchmarkMVP | tee s.csv


# some other possibilities
minimal:
	go run main.go testdata/minimal.csv 
fail:
	-go run main.go testdata/fail.csv 

curl:
	go run main.go --daemonic &
	if [ $$? -ne 0 ]; then echo "main.go returned $$?"; exit 1; fi
	curl localhost:8080 


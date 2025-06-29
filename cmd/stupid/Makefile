default: duplicate

budget:
	 cd ../../src/stupidestCache; go test -bench=. -count 25 -run=^# | tee s.csv


minimal:
	go run main.go testdata/minimal.csv 
fail:
	-go run main.go testdata/fail.csv 

curl: # FIXME, this fails
	go run main.go --daemonic &
	if [ $$? -ne 0 ]; then echo "main.go returned $$?"; exit 1; fi
	curl http://localhost:8080/1_en.12up.com_934664_US_HAS_USERID_1

duplicate:
	go run main.go testdata/p.csv testdata/g.csv 

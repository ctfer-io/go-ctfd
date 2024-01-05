.PHONY: tests
tests:
	@echo "--- Unitary tests ---"
	go test ./api -run=^Test_U_ -json | tee -a gotest.json

	@echo "--- Functional tests ---"
	go test ./api -run=^Test_F_ -coverprofile=functional.out -json | tee -a gotest.json

.PHONY: clean
clean:
	rm gotest.json functional.out

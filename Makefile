.PHONY: tests
tests:
	@echo "--- Unitary tests ---"
	go test ./api -run=^Test_U_ -json | tee -a gotest.json

	@echo "--- Functional tests ---"
	go test ./deploy/integration -run=^Test_F_ -json -coverpkg "github.com/ctfer-io/go-ctfd/api" -coverprofile=functional.out | tee -a gotest.json

.PHONY: clean
clean:
	rm gotest.json unitary.out functional.out

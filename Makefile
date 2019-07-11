BIN := go-libra-cli

BUILDENV := CGO_ENABLED=0
BUILDENV += GO111MODULE=on

LINKFLAGS :=-s -extldflags "-static"

$(BIN): clean
	$(BUILDENV) go build -o $(BIN) -a -ldflags '$(LINKFLAGS)' .

.PHONY: clean
clean:
	@rm -f go-libra-cli

AC_SCHEMA_DIR := libra/admission_control/admission_control_proto/src/proto
TYPES_SCHEMA_DIR := libra/types/src/proto
MEMPOOL_SCHEMA_DIR := libra/mempool/src/proto/shared

GENERATED_DIR := pb
GENERATED_AC_DIR := $(GENERATED_DIR)/admission_control
GENERATED_TYPES_DIR := $(GENERATED_DIR)/types
GENERATED_MEMPOOL_DIR := $(GENERATED_DIR)/mempool

PB_PKG := github.com/KentaKudo/go-libra-cli/pb
AC_PROTO_MAPPINGS := Mget_with_proof.proto=$(PB_PKG)/types,Mmempool_status.proto=$(PB_PKG)/mempool,Mtransaction.proto=$(PB_PKG)/types,Mvm_errors.proto=$(PB_PKG)/types
TYPES_PROTO_MAPPINGS := Mgoogle/protobuf/wrappers.proto=github.com/gogo/protobuf/types

.PHONY: .protos
protos:
	mkdir -pv $(GENERATED_DIR) $(GENERATED_AC_DIR) $(GENERATED_TYPES_DIR) $(GENERATED_MEMPOOL_DIR)
	protoc \
		-I $(TYPES_SCHEMA_DIR) \
		--gogoslick_out=$(TYPES_PROTO_MAPPINGS),paths=source_relative:$(GENERATED_TYPES_DIR) \
		access_path.proto account_state_blob.proto events.proto get_with_proof.proto language_storage.proto ledger_info.proto proof.proto transaction.proto transaction_info.proto validator_change.proto vm_errors.proto
	protoc \
		-I $(MEMPOOL_SCHEMA_DIR) \
		--gogoslick_out=paths=source_relative:$(GENERATED_MEMPOOL_DIR) \
		mempool_status.proto
	protoc \
		-I $(AC_SCHEMA_DIR) \
		-I $(TYPES_SCHEMA_DIR) \
		-I $(MEMPOOL_SCHEMA_DIR) \
		--gogoslick_out=$(AC_PROTO_MAPPINGS),paths=source_relative:$(GENERATED_AC_DIR) \
		admission_control.proto

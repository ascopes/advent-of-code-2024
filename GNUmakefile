SOLUTIONS := $(wildcard day*)

define delegate_to_each_solution
	$(foreach solution,$(SOLUTIONS),$(MAKE) -C $(solution) $@;)
endef

.PHONY: run

run:
	$(call delegate_to_each_solution)

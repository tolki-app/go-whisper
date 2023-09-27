.PHONY: push
push:
	git add . && git commit -m "$(if $(filter-out $@,$(MAKECMDGOALS)), $(filter-out $@,$(MAKECMDGOALS)), $(shell date))" && git push origin main
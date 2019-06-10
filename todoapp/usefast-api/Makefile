run:
	uvicorn main:app --reload
fmt:
	black *.py

see-docs: OPEN ?= $(shell which xdg-open 2>/dev/null || open)
see-docs:
	$(OPEN) http://localhost:8000/docs
	# http :8000/openapi.json

create-user:
	echo '{"name": "foo"}' | http --json post :8000/users


# Run inside all folders, run a `docker build -t <folder> .` and `docker run <folder>`
run-all:
	@for dir in $(shell ls -d *server/); do \
		dir=$${dir%/}; \
		$(MAKE) run/$$dir; \
	done

# run inside specific folder as parameter, run a `docker build -t <folder> .` and `docker run <folder>`
run/%:
	@cd $* ; \
	docker build -t $* .; \
	docker run -d --rm --name $* -p8080:8080 $* ; \
	sleep 5; \
	cd ..; \
	pytest api_test.py; \
	docker stop $* ;

up/%:
	@cd $* ; \
	docker build -t $* .; \
	docker run -it --rm --name $* -p8080:8080 $* ;
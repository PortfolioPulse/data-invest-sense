
check-project:
	npx nx test $(project)

image:
	npx nx run $(project):image

run: image
	docker-compose up -d

logs: run
	docker logs -f $(project)

check-project:
	npx nx test $(project)

image:
  npx nx run-many --target=image --projects=tag:scope:lake-orchestration
  npx nx run-many --target=image --projects=tag:scope:lake-ingestors

run: image
	docker-compose up -d

logs: run
	docker logs -f $(project)

set-controller-configs:
  npx nx g @nx-plugins/repo-settings:controller-configs

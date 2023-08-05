load('ext://uibutton', 'cmd_button', 'location', 'text_input')

local_resource(
  'lake-manager-api-compile',
  cmd='cd apps/apis/lake-manager; go install github.com/swaggo/swag/cmd/swag@latest && swag init -g cmd/server/main.go && GOOS=linux CGO_ENABLED=0 go build -ldflags="-w -s" -o ./build/server ./cmd/server/main.go',
  deps=['apps/apis/lake-manager/cmd/server/main.go'],
  labels=['lake-manager']
)

local_resource(
  'clean-database',
  cmd='echo "database droped"',
  deps=['lake-manager-api-compile'],
  labels=['lake-manager']
)


docker_build(
  'lake-manager-api',
  'apps/apis/lake-manager/',
  dockerfile='./apps/apis/lake-manager/Dockerfile',
  live_update=[
    sync('./apps/apis/lake-manager/build/server', '/app/server'),
  ],
)

# docker_compose('./apps/apis/lake-manager/docker-compose.yaml')

k8s_yaml(
  ['./apps/infra/k8s/services/deployment.yaml'],
)

k8s_resource(
  'lake-manager',
  port_forwards='8000:8000',
  labels=['lake-manager'],
)

cmd_button(
  'restart_service_group',
  text='Restart Service Group',
  icon_name='restart_alt',
  argv=['/bin/sh', '-c',
    'tilt get uiresource -l$RESOURCE=$RESOURCE --no-headers -ocustom-columns=:.metadata.name | xargs -L1 tilt trigger'
  ],
  location=location.NAV,
  inputs=[
    text_input(
      'RESOURCE',
      label='Resource',
    ),
  ]
)

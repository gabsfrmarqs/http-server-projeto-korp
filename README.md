# http-server-projeto-korp

Para execução local:
```
go run httpserver.go
```

Comando para facilitar minha vida:
```
sudo docker build -t http-server-projeto-korp:latest . && sudo docker compose up -d
```

## Parte 2
Grafana ativo em http://localhost:3000

Prometheus pode ser adicionado como datasource em http://prometheus:9090

http://localhost:3000/d/adqw9sx/dashboard-principal?from=2026-09-10T03:00:00.000Z&to=2026-09-10T04:59:59.000Z&timezone=browser

## Parte 3
Estarei testando o Ansible num container Fedora
```
sudo docker run -it --rm \
  -v /var/run/docker.sock:/var/run/docker.sock \
  -v $(pwd):/app \
  -w /app \
  fedora:latest bash
  
sudo docker exec -it fedora-test bash
dnf install -y ansible git
ansible-playbook -i ansible/inventory ansible/playbook.yml
```
# http-server-projeto-korp

Este repositório contém a automação de implantação e monitoramento para a aplicação **Projeto Korp**, construída em Go e orquestrada via **Docker Compose** e **Ansible**.


O Ansible foi testado num container Fedora
```
sudo docker run -it --rm \
  -v /var/run/docker.sock:/var/run/docker.sock \
  -v $(pwd):/app \
  -w /app \
  fedora:latest bash

#Dentro do container:
dnf install -y ansible git
ansible-playbook -i ansible/inventory ansible/playbook.yml
```

# Grafana
![](imagens/Screenshot_20260915_190246.png)
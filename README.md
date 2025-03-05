<h2>Instalação do banco de dados</h2>

<h4>1º - instalando docker, para instalar no windows previamente é necessário baixar o ubuntu na loja da microsoft e utilizar o script no prompt de comando do linux.</h4>

```
  # Add Docker's official GPG key:
sudo apt-get update
sudo apt-get install ca-certificates curl
sudo install -m 0755 -d /etc/apt/keyrings
sudo curl -fsSL https://download.docker.com/linux/ubuntu/gpg -o /etc/apt/keyrings/docker.asc
sudo chmod a+r /etc/apt/keyrings/docker.asc

# Add the repository to Apt sources:
echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.asc] https://download.docker.com/linux/ubuntu \
  $(. /etc/os-release && echo "${UBUNTU_CODENAME:-$VERSION_CODENAME}") stable" | \
  sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
sudo apt-get update
```

```
sudo apt-get install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin
```
<h4>usar script para verificar se funcionou: </h4>

```
sudo docker run hello-world
```

<h4>instalação do banco postgres</h4>

```
sudo docker pull postgres
```

```
sudo docker run --name chatbot -e POSTGRES_PASSWORD=[senha] -d -p 5432:5432 postgres
```

<h4>Como entrar no Postgres via terminal</h4>
```
sudo docker exec -it chatbot bash
```

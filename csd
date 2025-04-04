New-Item -ItemType Directory -Path go-api-engine | Out-Null
Set-Location go-api-engine
go mod init go-api-engine
New-Item -ItemType Directory -Path "cmd\go-api-engine"
New-Item -ItemType Directory -Path "internal\tcp"
New-Item -ItemType Directory -Path "pkg\utils"
New-Item -ItemType Directory -Path "config"
New-Item -ItemType File -Path "cmd\go-api-engine\main.go"
New-Item -ItemType File -Path "internal\tcp\server.go"
New-Item -ItemType File -Path "internal\tcp\client.go"
New-Item -ItemType File -Path "pkg\utils\logger.go"
New-Item -ItemType File -Path "config\config.yaml"
New-Item -ItemType File -Path "README.md"
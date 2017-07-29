

Build
```[vagrant@localhost hydra]$ go install -v ./
github.com/ory/hydra/vendor/github.com/julienschmidt/httprouter
github.com/ory/hydra/vendor/github.com/jmoiron/sqlx/reflectx
github.com/ory/hydra/vendor/github.com/jmoiron/sqlx
github.com/ory/hydra/vendor/github.com/imdario/mergo
github.com/asaskevich/govalidator
github.com/ory/hydra/vendor/github.com/pborman/uuid
github.com/ory/hydra/vendor/github.com/pkg/errors
golang.org/x/crypto/blowfish
golang.org/x/sys/unix
golang.org/x/crypto/bcrypt
github.com/ory/hydra/vendor/github.com/ory/fosite
github.com/ory/hydra/vendor/github.com/sirupsen/logrus
github.com/ory/hydra/firewall
github.com/ory/hydra/vendor/github.com/dgrijalva/jwt-go
github.com/ory/hydra/vendor/github.com/ory/herodot
github.com/ory/hydra/vendor/github.com/ory/fosite/token/hmac
github.com/ory/hydra/vendor/github.com/ory/fosite/token/jwt
github.com/ory/hydra/vendor/github.com/ory/fosite/storage
github.com/ory/hydra/vendor/github.com/ory/fosite/handler/oauth2
github.com/ory/hydra/vendor/github.com/moul/http2curl
github.com/ory/hydra/pkg/helper
github.com/ory/hydra/rand/sequence
github.com/hashicorp/golang-lru
github.com/ory/hydra/vendor/github.com/ory/fosite/handler/openid
github.com/ory/hydra/vendor/github.com/ory/ladon/compiler
github.com/ory/hydra/vendor/github.com/stretchr/testify/vendor/github.com/davecgh/go-spew/spew
github.com/ory/hydra/vendor/github.com/stretchr/testify/vendor/github.com/pmezard/go-difflib/difflib
github.com/ory/hydra/vendor/github.com/rubenv/sql-migrate/sqlparse
github.com/ory/hydra/vendor/github.com/stretchr/testify/assert
gopkg.in/gorp.v1
github.com/ory/hydra/vendor/github.com/stretchr/testify/require
github.com/ory/hydra/vendor/github.com/ory/ladon
github.com/ory/hydra/vendor/github.com/rubenv/sql-migrate
github.com/ory/hydra/vendor/github.com/ory/ladon/manager/memory
github.com/ory/hydra/pkg
github.com/ory/hydra/vendor/github.com/go-sql-driver/mysql
github.com/ory/hydra/client
github.com/ory/hydra/vendor/github.com/lib/pq/oid
github.com/ory/hydra/vendor/github.com/lib/pq
gopkg.in/square/go-jose.v1/cipher
gopkg.in/square/go-jose.v1/json
github.com/ory/hydra/vendor/github.com/square/go-jose
github.com/jehiah/go-strftime
github.com/segmentio/backo-go
github.com/xtgo/uuid
github.com/segmentio/analytics-go
github.com/ory/hydra/vendor/github.com/urfave/negroni
github.com/ory/hydra/jwk
github.com/ory/hydra/metrics
github.com/ory/hydra/warden/group
github.com/ory/hydra/vendor/github.com/ory/ladon/manager/sql
github.com/ory/hydra/vendor/github.com/spf13/cobra
github.com/fsnotify/fsnotify
github.com/hashicorp/hcl/hcl/strconv
github.com/hashicorp/hcl/hcl/token
github.com/magiconair/properties
github.com/hashicorp/hcl/hcl/ast
github.com/hashicorp/hcl/hcl/scanner
github.com/hashicorp/hcl/hcl/parser
github.com/hashicorp/hcl/json/token
github.com/mitchellh/mapstructure
github.com/hashicorp/hcl/json/scanner
github.com/hashicorp/hcl/json/parser
github.com/pelletier/go-buffruneio
github.com/hashicorp/hcl
github.com/pelletier/go-toml
github.com/kr/fs
golang.org/x/crypto/curve25519
golang.org/x/crypto/ed25519/internal/edwards25519
github.com/spf13/afero/mem
golang.org/x/text/transform
golang.org/x/crypto/ed25519
golang.org/x/crypto/ssh
golang.org/x/text/unicode/norm
github.com/spf13/cast
github.com/spf13/jwalterweatherman
github.com/ory/hydra/vendor/gopkg.in/yaml.v2
github.com/pkg/sftp
github.com/ory/hydra/vendor/golang.org/x/oauth2/internal
github.com/spf13/afero/sftp
github.com/ory/hydra/vendor/golang.org/x/oauth2
github.com/spf13/afero
github.com/ory/hydra/vendor/golang.org/x/oauth2/clientcredentials
github.com/ory/hydra/vendor/github.com/gorilla/context
github.com/gorilla/securecookie
github.com/ory/hydra/vendor/github.com/spf13/viper
github.com/ory/hydra/vendor/github.com/gorilla/sessions
github.com/ory/hydra/vendor/github.com/mohae/deepcopy
github.com/ory/hydra/config
github.com/ory/hydra/vendor/github.com/oleiade/reflections
github.com/ory/hydra/oauth2
github.com/ory/hydra/policy
github.com/ory/hydra/vendor/github.com/square/go-jose/json
github.com/ory/hydra/vendor/github.com/meatballhat/negroni-logrus
github.com/ory/hydra/vendor/github.com/ory/fosite/compose
github.com/ory/hydra/vendor/github.com/ory/graceful
github.com/ory/hydra/cmd/cli
github.com/ory/hydra/health
github.com/ory/hydra/warden
github.com/ory/hydra/vendor/github.com/toqueteos/webbrowser
github.com/ory/hydra/cmd/server
github.com/ory/hydra/vendor/github.com/pkg/profile
github.com/ory/hydra/cmd
github.com/ory/hydra
```

Test
```
[vagrant@localhost hydra]$ DATABASE_URL=memory hydra host --dangerous-auto-logon --dangerous-force-http
INFO[0000] DATABASE_URL set to memory, connecting to ephermal in-memory database. 
WARN[0000] Expected system secret to be at least 32 characters long, got 0 characters. 
INFO[0000] Generating a random system secret...         
INFO[0000] Generated system secret: HrlgoHh7NeUvmzDCDHtRr9J0qjX-jq0b 
WARN[0000] WARNING: DO NOT generate system secrets in production. The secret will be leaked to the logs. 
INFO[0000] Key pair for signing hydra.openid.id-token is missing. Creating new one. 
INFO[0003] Setting up telemetry - for more information please visit https://ory.gitbooks.io/hydra/content/telemetry.html 
INFO[0003] Key pair for signing hydra.consent.response is missing. Creating new one. 
INFO[0013] Key pair for signing hydra.consent.challenge is missing. Creating new one. 
WARN[0014] No clients were found. Creating a temporary root client... 
INFO[0014] Temporary root client created.               
INFO[0014] client_id: 439cb36b-eba6-4207-8577-1f1d210c971c 
INFO[0014] client_secret: r12Gw4DagKe9EW-8              
WARN[0014] WARNING: YOU MUST delete this client once in production, as credentials may have been leaked in your logfiles. 
WARN[0014] Do not use flag --dangerous-auto-logon in production. 
INFO[0014] Persisting config in file /home/vagrant/.hydra.yml 
WARN[0014] No TLS Key / Certificate for HTTPS found. Generating self-signed certificate. 
INFO[0014] Setting up http server on :4444              
WARN[0014] HTTPS disabled. Never do this in production. 
```
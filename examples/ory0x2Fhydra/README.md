

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

### Consent-Test

Server
```
fanhonglingdeMacBook-Pro:go-to-openstack-bootcamp fanhongling$ DATABASE_URL=memory ISSUER=http://localhost:4444/ FORCE_ROOT_CLIENT_CREDENTIALS=demo:demo CONSENT_URL=http://localhost:4445/consent hydra host --dangerous-force-http --skip-tls-verify 
INFO[0000] DATABASE_URL set to memory, connecting to ephermal in-memory database. 
WARN[0000] Expected system secret to be at least 32 characters long, got 0 characters. 
INFO[0000] Generating a random system secret...         
INFO[0000] Generated system secret: ~XXlGC4oZJ7uFHSXsCBYalCU1rX-VlOS 
WARN[0000] WARNING: DO NOT generate system secrets in production. The secret will be leaked to the logs. 
INFO[0000] Key pair for signing hydra.openid.id-token is missing. Creating new one. 
INFO[0002] Setting up telemetry - for more information please visit https://ory.gitbooks.io/hydra/content/telemetry.html 
INFO[0002] Key pair for signing hydra.consent.response is missing. Creating new one. 
INFO[0004] Key pair for signing hydra.consent.challenge is missing. Creating new one. 
WARN[0008] No clients were found. Creating a temporary root client... 
INFO[0008] Temporary root client created.               
WARN[0008] No TLS Key / Certificate for HTTPS found. Generating self-signed certificate. 
INFO[0008] Setting up http server on :4444              
WARN[0008] HTTPS disabled. Never do this in production. 


```

Client
```
fanhonglingdeMacBook-Pro:ory fanhongling$ git clone --depth=1 https://github.com/ory/hydra-consent-app-go hydra-consent-app-go
Cloning into 'hydra-consent-app-go'...
remote: Counting objects: 15, done.
remote: Compressing objects: 100% (13/13), done.
remote: Total 15 (delta 0), reused 15 (delta 0), pack-reused 0
Unpacking objects: 100% (15/15), done.
Checking connectivity... done.
fanhonglingdeMacBook-Pro:hydra-consent-app-go fanhongling$ glide install
[INFO]	Downloading dependencies. Please wait...
[INFO]	--> Found desired version locally github.com/asaskevich/govalidator 7b3beb6df3c42abd3509abfc3bcacc0fbfb7c877!
[INFO]	--> Found desired version locally github.com/davecgh/go-spew 6d212800a42e8ab5c146b8ace3490ee17e5225f9!
[INFO]	--> Found desired version locally github.com/dgrijalva/jwt-go d2709f9f1f31ebcda9651b03077758c1f3a0018c!
[INFO]	--> Found desired version locally github.com/fsnotify/fsnotify a904159b9206978bb6d53fcc7a769e5cd726c737!
[INFO]	--> Found desired version locally github.com/go-sql-driver/mysql a0583e0143b1624142adab07e0e97fe106d99561!
[INFO]	--> Found desired version locally github.com/golang/protobuf 69b215d01a5606c843240eab4937eab3acee6530!
[INFO]	--> Found desired version locally github.com/gorilla/context 08b5f424b9271eedf6f9f0ce86cb9396ed337a42!
[INFO]	--> Found desired version locally github.com/gorilla/mux 392c28fe23e1c45ddba891b0320b3b5df220beea!
[INFO]	--> Found desired version locally github.com/gorilla/securecookie fa5329f913702981df43dcb2a380bac429c810b5!
[INFO]	--> Found desired version locally github.com/gorilla/sessions ca9ada44574153444b00d3fd9c8559e4cc95f896!
[INFO]	--> Found desired version locally github.com/hashicorp/hcl 630949a3c5fa3c613328e1b8256052cbc2327c9b!
[INFO]	--> Found desired version locally github.com/imdario/mergo 3e95a51e0639b4cf372f2ccf74c86749d747fbdc!
[INFO]	--> Found desired version locally github.com/inconshreveable/mousetrap 76626ae9c91c4f2a10f34cad8ce83ea42c93bb75!
[INFO]	--> Found desired version locally github.com/jmoiron/sqlx f4076845477b10ac2453a16377a8383467aafe72!
[INFO]	--> Found desired version locally github.com/julienschmidt/httprouter 8c199fb6259ffc1af525cc3ad52ee60ba8359669!
[INFO]	--> Found desired version locally github.com/lib/pq ba5d4f7a35561e22fbdf7a39aa0070f4d460cfc0!
[INFO]	--> Found desired version locally github.com/magiconair/properties b3b15ef068fd0b17ddf408a23669f20811d194d2!
[INFO]	--> Found desired version locally github.com/meatballhat/negroni-logrus 7c570a907cfc69cdc004ad506c6f5e234815b936!
[INFO]	--> Found desired version locally github.com/mitchellh/mapstructure db1efb556f84b25a0a13a04aad883943538ad2e0!
[INFO]	--> Found desired version locally github.com/moul/http2curl 4e24498b31dba4683efb9d35c1c8a91e2eda28c8!
[INFO]	--> Found desired version locally github.com/oleiade/reflections 2b6ec3da648e3e834dc41bad8d9ed7f2dc6a9496!
[INFO]	--> Found desired version locally github.com/pborman/uuid a97ce2ca70fa5a848076093f05e639a89ca34d06!
[INFO]	--> Found desired version locally github.com/pelletier/go-buffruneio df1e16fde7fc330a0ca68167c23bf7ed6ac31d6d!
[INFO]	--> Found desired version locally github.com/pelletier/go-toml 22139eb5469018e7374b3e7ef653de37ffb44f72!
[INFO]	--> Found desired version locally github.com/pkg/errors 645ef00459ed84a119197bfb8d8205042c6df63d!
[INFO]	--> Found desired version locally github.com/pmezard/go-difflib d8ed2627bdf02c080bf22230dbb337003b7aba2d!
[INFO]	--> Found desired version locally github.com/rubenv/sql-migrate f64b6080c334adaf843209164107439e92bb170b!
[INFO]	--> Found desired version locally github.com/Sirupsen/logrus ba1b36c82c5e05c4f912a88eab0dcd91a171688f!
[INFO]	--> Found desired version locally github.com/spf13/afero 9be650865eab0c12963d8753212f4f9c66cdcf12!
[INFO]	--> Found desired version locally github.com/spf13/cast f820543c3592e283e311a60d2a600a664e39f6f7!
[INFO]	--> Found desired version locally github.com/spf13/cobra 92ea23a837e66f46ac9e7d04fa826602b7b0a42d!
[INFO]	--> Found desired version locally github.com/spf13/jwalterweatherman fa7ca7e836cf3a8bb4ebf799f472c12d7e903d66!
[INFO]	--> Found desired version locally github.com/spf13/pflag 9ff6c6923cfffbcd502984b8e0c80539a94968b7!
[INFO]	--> Found desired version locally github.com/spf13/viper 7538d73b4eb9511d85a9f1dfef202eeb8ac260f4!
[INFO]	--> Found desired version locally github.com/square/go-jose aa2e30fdd1fe9dd3394119af66451ae790d50e0d!
[INFO]	--> Found desired version locally github.com/stretchr/testify 69483b4bd14f5845b5a1e55bca19e954e827f1d0!
[INFO]	--> Found desired version locally github.com/urfave/negroni fde5e16d32adc7ad637e9cd9ad21d4ebc6192535!
[INFO]	--> Found desired version locally golang.org/x/crypto 453249f01cfeb54c3d549ddb75ff152ca243f9d8!
[INFO]	--> Found desired version locally golang.org/x/net dd2d9a67c97da0afa00d5726e28086007a0acce5!
[INFO]	--> Found desired version locally golang.org/x/oauth2 b9780ec78894ab900c062d58ee3076cd9b2a4501!
[INFO]	--> Found desired version locally golang.org/x/sys e4594059fe4cde2daf423055a596c2cd1e6c9adf!
[INFO]	--> Found desired version locally golang.org/x/text 4687d739464a2d0af89a25be0318456e0776f3ef!
[INFO]	--> Found desired version locally google.golang.org/appengine 3a452f9e00122ead39586d68ffdb9c6e1326af3c!
[INFO]	--> Found desired version locally gopkg.in/gorp.v1 c87af80f3cc5036b55b83d77171e156791085e2e!
[INFO]	--> Found desired version locally gopkg.in/square/go-jose.v1 aa2e30fdd1fe9dd3394119af66451ae790d50e0d!
[INFO]	--> Found desired version locally gopkg.in/yaml.v2 a3f3340b5840cee44f372bddb5880fcbc419b46a!
[INFO]	--> Fetching github.com/cenkalti/backoff
[INFO]	--> Fetching github.com/ory-am/hydra
[INFO]	--> Fetching github.com/hailocab/go-hostpool
[INFO]	--> Fetching github.com/ory-am/common
[INFO]	--> Fetching github.com/ory-am/ladon
[INFO]	--> Fetching github.com/ory/common
[INFO]	--> Fetching gopkg.in/dancannon/gorethink.v2
[INFO]	--> Fetching gopkg.in/gorethink/gorethink.v3
[INFO]	--> Fetching gopkg.in/redis.v5
[INFO]	--> Fetching gopkg.in/fatih/pool.v2
[INFO]	--> Fetching github.com/go-errors/errors
[INFO]	--> Fetching github.com/ory-am/fosite
[INFO]	Setting references.
[INFO]	--> Setting version for github.com/fsnotify/fsnotify to a904159b9206978bb6d53fcc7a769e5cd726c737.
[INFO]	--> Setting version for github.com/hashicorp/hcl to 630949a3c5fa3c613328e1b8256052cbc2327c9b.
[INFO]	--> Setting version for github.com/inconshreveable/mousetrap to 76626ae9c91c4f2a10f34cad8ce83ea42c93bb75.
[INFO]	--> Setting version for github.com/lib/pq to ba5d4f7a35561e22fbdf7a39aa0070f4d460cfc0.
[INFO]	--> Setting version for github.com/gorilla/context to 08b5f424b9271eedf6f9f0ce86cb9396ed337a42.
[INFO]	--> Setting version for github.com/davecgh/go-spew to 6d212800a42e8ab5c146b8ace3490ee17e5225f9.
[INFO]	--> Setting version for github.com/jmoiron/sqlx to f4076845477b10ac2453a16377a8383467aafe72.
[INFO]	--> Setting version for github.com/hailocab/go-hostpool to e80d13ce29ede4452c43dea11e79b9bc8a15b478.
[INFO]	--> Setting version for github.com/gorilla/sessions to ca9ada44574153444b00d3fd9c8559e4cc95f896.
[INFO]	--> Setting version for github.com/go-errors/errors to 8fa88b06e5974e97fbf9899a7f86a344bfd1f105.
[INFO]	--> Setting version for github.com/magiconair/properties to b3b15ef068fd0b17ddf408a23669f20811d194d2.
[INFO]	--> Setting version for github.com/gorilla/mux to 392c28fe23e1c45ddba891b0320b3b5df220beea.
[INFO]	--> Setting version for github.com/julienschmidt/httprouter to 8c199fb6259ffc1af525cc3ad52ee60ba8359669.
[INFO]	--> Setting version for github.com/imdario/mergo to 3e95a51e0639b4cf372f2ccf74c86749d747fbdc.
[INFO]	--> Setting version for github.com/gorilla/securecookie to fa5329f913702981df43dcb2a380bac429c810b5.
[INFO]	--> Setting version for github.com/asaskevich/govalidator to 7b3beb6df3c42abd3509abfc3bcacc0fbfb7c877.
[INFO]	--> Setting version for github.com/go-sql-driver/mysql to a0583e0143b1624142adab07e0e97fe106d99561.
[INFO]	--> Setting version for github.com/golang/protobuf to 69b215d01a5606c843240eab4937eab3acee6530.
[INFO]	--> Setting version for github.com/cenkalti/backoff to b02f2bbce11d7ea6b97f282ef1771b0fe2f65ef3.
[INFO]	--> Setting version for github.com/dgrijalva/jwt-go to d2709f9f1f31ebcda9651b03077758c1f3a0018c.
[INFO]	--> Setting version for github.com/meatballhat/negroni-logrus to 7c570a907cfc69cdc004ad506c6f5e234815b936.
[INFO]	--> Setting version for github.com/mitchellh/mapstructure to db1efb556f84b25a0a13a04aad883943538ad2e0.
[INFO]	--> Setting version for github.com/moul/http2curl to 4e24498b31dba4683efb9d35c1c8a91e2eda28c8.
[INFO]	--> Setting version for github.com/oleiade/reflections to 2b6ec3da648e3e834dc41bad8d9ed7f2dc6a9496.
[INFO]	--> Setting version for github.com/ory-am/common to eaf2f2a2e18295ffc4e89d07a9556903d22b4e79.
[INFO]	--> Setting version for github.com/ory-am/fosite to 9b33931ee14ae0768ea46a423d569330a85b482e.
[INFO]	--> Setting version for github.com/ory-am/hydra to 93bb521963ab05557553de6a5b0b34f7a1b8def4.
[INFO]	--> Setting version for github.com/ory-am/ladon to ef4a3e7a29dc84bb3fabb23142e350fbb5bfdd42.
[INFO]	--> Setting version for github.com/ory/common to ba06ec2f738cb3a55608657c2e998a1eef675423.
[INFO]	--> Setting version for github.com/pborman/uuid to a97ce2ca70fa5a848076093f05e639a89ca34d06.
[INFO]	--> Setting version for github.com/pelletier/go-buffruneio to df1e16fde7fc330a0ca68167c23bf7ed6ac31d6d.
[INFO]	--> Setting version for github.com/pelletier/go-toml to 22139eb5469018e7374b3e7ef653de37ffb44f72.
[INFO]	--> Setting version for github.com/pkg/errors to 645ef00459ed84a119197bfb8d8205042c6df63d.
[INFO]	--> Setting version for github.com/rubenv/sql-migrate to f64b6080c334adaf843209164107439e92bb170b.
[INFO]	--> Setting version for github.com/pmezard/go-difflib to d8ed2627bdf02c080bf22230dbb337003b7aba2d.
[INFO]	--> Setting version for github.com/Sirupsen/logrus to ba1b36c82c5e05c4f912a88eab0dcd91a171688f.
[INFO]	--> Setting version for github.com/spf13/afero to 9be650865eab0c12963d8753212f4f9c66cdcf12.
[INFO]	--> Setting version for github.com/spf13/cast to f820543c3592e283e311a60d2a600a664e39f6f7.
[INFO]	--> Setting version for github.com/spf13/jwalterweatherman to fa7ca7e836cf3a8bb4ebf799f472c12d7e903d66.
[INFO]	--> Setting version for github.com/spf13/cobra to 92ea23a837e66f46ac9e7d04fa826602b7b0a42d.
[INFO]	--> Setting version for github.com/spf13/pflag to 9ff6c6923cfffbcd502984b8e0c80539a94968b7.
[INFO]	--> Setting version for github.com/spf13/viper to 7538d73b4eb9511d85a9f1dfef202eeb8ac260f4.
[INFO]	--> Setting version for github.com/square/go-jose to aa2e30fdd1fe9dd3394119af66451ae790d50e0d.
[INFO]	--> Setting version for github.com/stretchr/testify to 69483b4bd14f5845b5a1e55bca19e954e827f1d0.
[INFO]	--> Setting version for github.com/urfave/negroni to fde5e16d32adc7ad637e9cd9ad21d4ebc6192535.
[INFO]	--> Setting version for golang.org/x/crypto to 453249f01cfeb54c3d549ddb75ff152ca243f9d8.
[INFO]	--> Setting version for golang.org/x/net to dd2d9a67c97da0afa00d5726e28086007a0acce5.
[INFO]	--> Setting version for golang.org/x/oauth2 to b9780ec78894ab900c062d58ee3076cd9b2a4501.
[INFO]	--> Setting version for golang.org/x/sys to e4594059fe4cde2daf423055a596c2cd1e6c9adf.
[INFO]	--> Setting version for golang.org/x/text to 4687d739464a2d0af89a25be0318456e0776f3ef.
[INFO]	--> Setting version for google.golang.org/appengine to 3a452f9e00122ead39586d68ffdb9c6e1326af3c.
[INFO]	--> Setting version for gopkg.in/fatih/pool.v2 to 6e328e67893eb46323ad06f0e92cb9536babbabc.
[INFO]	--> Setting version for gopkg.in/dancannon/gorethink.v2 to 610fcc04c971a9fa42f1b6625c7c52b5bb472c51.
[INFO]	--> Setting version for gopkg.in/square/go-jose.v1 to aa2e30fdd1fe9dd3394119af66451ae790d50e0d.
[INFO]	--> Setting version for gopkg.in/gorp.v1 to c87af80f3cc5036b55b83d77171e156791085e2e.
[INFO]	--> Setting version for gopkg.in/gorethink/gorethink.v3 to 610fcc04c971a9fa42f1b6625c7c52b5bb472c51.
[INFO]	--> Setting version for gopkg.in/redis.v5 to a16aeec10ff407b1e7be6dd35797ccf5426ef0f0.
[INFO]	--> Setting version for gopkg.in/yaml.v2 to a3f3340b5840cee44f372bddb5880fcbc419b46a.
[INFO]	Exporting resolved dependencies...
[INFO]	--> Exporting github.com/dgrijalva/jwt-go
[INFO]	--> Exporting github.com/davecgh/go-spew
[INFO]	--> Exporting github.com/cenkalti/backoff
[INFO]	--> Exporting github.com/magiconair/properties
[INFO]	--> Exporting github.com/fsnotify/fsnotify
[INFO]	--> Exporting github.com/hashicorp/hcl
[INFO]	--> Exporting github.com/golang/protobuf
[INFO]	--> Exporting github.com/go-errors/errors
[INFO]	--> Exporting github.com/gorilla/securecookie
[INFO]	--> Exporting github.com/gorilla/sessions
[INFO]	--> Exporting github.com/gorilla/mux
[INFO]	--> Exporting github.com/go-sql-driver/mysql
[INFO]	--> Exporting github.com/lib/pq
[INFO]	--> Exporting github.com/imdario/mergo
[INFO]	--> Exporting github.com/hailocab/go-hostpool
[INFO]	--> Exporting github.com/julienschmidt/httprouter
[INFO]	--> Exporting github.com/jmoiron/sqlx
[INFO]	--> Exporting github.com/asaskevich/govalidator
[INFO]	--> Exporting github.com/gorilla/context
[INFO]	--> Exporting github.com/inconshreveable/mousetrap
[INFO]	--> Exporting github.com/meatballhat/negroni-logrus
[INFO]	--> Exporting github.com/mitchellh/mapstructure
[INFO]	--> Exporting github.com/moul/http2curl
[INFO]	--> Exporting github.com/oleiade/reflections
[INFO]	--> Exporting github.com/ory-am/common
[INFO]	--> Exporting github.com/ory-am/fosite
[INFO]	--> Exporting github.com/ory-am/hydra
[INFO]	--> Exporting github.com/ory-am/ladon
[INFO]	--> Exporting github.com/pborman/uuid
[INFO]	--> Exporting github.com/pelletier/go-buffruneio
[INFO]	--> Exporting github.com/ory/common
[INFO]	--> Exporting github.com/pelletier/go-toml
[INFO]	--> Exporting github.com/pkg/errors
[INFO]	--> Exporting github.com/pmezard/go-difflib
[INFO]	--> Exporting github.com/Sirupsen/logrus
[INFO]	--> Exporting github.com/rubenv/sql-migrate
[INFO]	--> Exporting github.com/spf13/afero
[INFO]	--> Exporting github.com/spf13/cobra
[INFO]	--> Exporting github.com/spf13/cast
[INFO]	--> Exporting github.com/spf13/jwalterweatherman
[INFO]	--> Exporting github.com/spf13/pflag
[INFO]	--> Exporting github.com/spf13/viper
[INFO]	--> Exporting github.com/square/go-jose
[INFO]	--> Exporting github.com/stretchr/testify
[INFO]	--> Exporting github.com/urfave/negroni
[INFO]	--> Exporting golang.org/x/crypto
[INFO]	--> Exporting golang.org/x/oauth2
[INFO]	--> Exporting golang.org/x/sys
[INFO]	--> Exporting golang.org/x/net
[INFO]	--> Exporting golang.org/x/text
[INFO]	--> Exporting google.golang.org/appengine
[INFO]	--> Exporting gopkg.in/dancannon/gorethink.v2
[INFO]	--> Exporting gopkg.in/fatih/pool.v2
[INFO]	--> Exporting gopkg.in/gorethink/gorethink.v3
[INFO]	--> Exporting gopkg.in/square/go-jose.v1
[INFO]	--> Exporting gopkg.in/gorp.v1
[INFO]	--> Exporting gopkg.in/redis.v5
[INFO]	--> Exporting gopkg.in/yaml.v2
[INFO]	Replacing existing vendor dependencies

fanhonglingdeMacBook-Pro:hydra-consent-app-go fanhongling$ go install -v ./
github.com/ory/hydra-consent-app-go/vendor/github.com/cenkalti/backoff
github.com/ory/hydra-consent-app-go/vendor/golang.org/x/crypto/blowfish
github.com/ory/hydra-consent-app-go/vendor/github.com/gorilla/mux
github.com/ory/hydra-consent-app-go/vendor/github.com/gorilla/context
github.com/ory/hydra-consent-app-go/vendor/github.com/gorilla/securecookie
github.com/ory/hydra-consent-app-go/vendor/github.com/Sirupsen/logrus
github.com/ory/hydra-consent-app-go/vendor/github.com/urfave/negroni
github.com/ory/hydra-consent-app-go/vendor/github.com/gorilla/sessions
github.com/ory/hydra-consent-app-go/vendor/github.com/dgrijalva/jwt-go
github.com/ory/hydra-consent-app-go/vendor/github.com/pborman/uuid
github.com/ory/hydra-consent-app-go/vendor/github.com/meatballhat/negroni-logrus
github.com/ory/hydra-consent-app-go/vendor/github.com/pkg/errors
github.com/ory/hydra-consent-app-go/vendor/github.com/imdario/mergo
github.com/ory/hydra-consent-app-go/vendor/github.com/jmoiron/sqlx/reflectx
github.com/ory/hydra-consent-app-go/vendor/github.com/julienschmidt/httprouter
github.com/ory/hydra-consent-app-go/vendor/github.com/ory-am/fosite/token/jwt
github.com/ory/hydra-consent-app-go/vendor/github.com/jmoiron/sqlx
github.com/ory/hydra-consent-app-go/vendor/github.com/ory-am/common/rand/sequence
github.com/ory/hydra-consent-app-go/vendor/github.com/asaskevich/govalidator
github.com/ory/hydra-consent-app-go/vendor/github.com/square/go-jose/json
github.com/ory/hydra-consent-app-go/vendor/golang.org/x/crypto/bcrypt
github.com/ory/hydra-consent-app-go/vendor/golang.org/x/net/context
github.com/ory/hydra-consent-app-go/vendor/github.com/ory-am/hydra/firewall
github.com/ory/hydra-consent-app-go/vendor/github.com/go-errors/errors
github.com/ory/hydra-consent-app-go/vendor/github.com/oleiade/reflections
github.com/ory/hydra-consent-app-go/vendor/github.com/stretchr/testify/vendor/github.com/davecgh/go-spew/spew
github.com/ory/hydra-consent-app-go/vendor/github.com/stretchr/testify/vendor/github.com/pmezard/go-difflib/difflib
github.com/ory/hydra-consent-app-go/vendor/github.com/moul/http2curl
github.com/ory/hydra-consent-app-go/vendor/github.com/stretchr/testify/assert
github.com/ory/hydra-consent-app-go/vendor/github.com/ory-am/hydra/pkg/helper
github.com/ory/hydra-consent-app-go/vendor/github.com/ory-am/common/compiler
github.com/ory/hydra-consent-app-go/vendor/github.com/rubenv/sql-migrate/sqlparse
github.com/ory/hydra-consent-app-go/vendor/gopkg.in/gorp.v1
github.com/ory/hydra-consent-app-go/vendor/github.com/hailocab/go-hostpool
github.com/ory/hydra-consent-app-go/vendor/golang.org/x/crypto/pbkdf2
github.com/ory/hydra-consent-app-go/vendor/gopkg.in/fatih/pool.v2
github.com/ory/hydra-consent-app-go/vendor/gopkg.in/gorethink/gorethink.v3/encoding
github.com/ory/hydra-consent-app-go/vendor/github.com/ory-am/fosite
github.com/ory/hydra-consent-app-go/vendor/github.com/stretchr/testify/require
github.com/ory/hydra-consent-app-go/vendor/github.com/ory-am/common/pkg
github.com/ory/hydra-consent-app-go/vendor/github.com/golang/protobuf/proto
github.com/ory/hydra-consent-app-go/vendor/github.com/rubenv/sql-migrate
github.com/ory/hydra-consent-app-go/vendor/gopkg.in/gorethink/gorethink.v3/types
github.com/ory/hydra-consent-app-go/vendor/gopkg.in/redis.v5/internal
github.com/ory/hydra-consent-app-go/vendor/gopkg.in/redis.v5/internal/consistenthash
github.com/ory/hydra-consent-app-go/vendor/github.com/ory-am/hydra/herodot
github.com/ory/hydra-consent-app-go/vendor/github.com/ory-am/fosite/token/hmac
github.com/ory/hydra-consent-app-go/vendor/github.com/ory-am/fosite/storage
github.com/ory/hydra-consent-app-go/vendor/github.com/ory-am/fosite/handler/oauth2
github.com/ory/hydra-consent-app-go/vendor/gopkg.in/redis.v5/internal/hashtag
github.com/ory/hydra-consent-app-go/vendor/gopkg.in/redis.v5/internal/proto
github.com/ory/hydra-consent-app-go/vendor/gopkg.in/square/go-jose.v1/cipher
github.com/ory/hydra-consent-app-go/vendor/gopkg.in/square/go-jose.v1/json
github.com/ory/hydra-consent-app-go/vendor/gopkg.in/redis.v5/internal/pool
github.com/ory/hydra-consent-app-go/vendor/gopkg.in/redis.v5
github.com/ory/hydra-consent-app-go/vendor/github.com/ory-am/fosite/handler/openid
github.com/ory/hydra-consent-app-go/vendor/golang.org/x/oauth2/internal
github.com/ory/hydra-consent-app-go/vendor/github.com/square/go-jose
github.com/ory/hydra-consent-app-go/vendor/golang.org/x/oauth2
github.com/ory/hydra-consent-app-go/vendor/golang.org/x/oauth2/clientcredentials
github.com/ory/hydra-consent-app-go/vendor/github.com/go-sql-driver/mysql
github.com/ory/hydra-consent-app-go/vendor/gopkg.in/gorethink/gorethink.v3/ql2
github.com/ory/hydra-consent-app-go/vendor/gopkg.in/dancannon/gorethink.v2
github.com/ory/hydra-consent-app-go/vendor/github.com/lib/pq/oid
github.com/ory/hydra-consent-app-go/vendor/github.com/lib/pq
github.com/ory/hydra-consent-app-go/vendor/github.com/spf13/pflag
github.com/ory/hydra-consent-app-go/vendor/github.com/spf13/cobra
github.com/ory/hydra-consent-app-go/vendor/golang.org/x/sys/unix
github.com/ory/hydra-consent-app-go/vendor/github.com/hashicorp/hcl/hcl/strconv
github.com/ory/hydra-consent-app-go/vendor/github.com/hashicorp/hcl/hcl/token
github.com/ory/hydra-consent-app-go/vendor/github.com/hashicorp/hcl/hcl/ast
github.com/ory/hydra-consent-app-go/vendor/github.com/hashicorp/hcl/hcl/scanner
github.com/ory/hydra-consent-app-go/vendor/github.com/hashicorp/hcl/hcl/parser
github.com/ory/hydra-consent-app-go/vendor/github.com/hashicorp/hcl/json/token
github.com/ory/hydra-consent-app-go/vendor/github.com/hashicorp/hcl/json/scanner
github.com/ory/hydra-consent-app-go/vendor/github.com/magiconair/properties
github.com/ory/hydra-consent-app-go/vendor/github.com/hashicorp/hcl/json/parser
github.com/ory/hydra-consent-app-go/vendor/github.com/fsnotify/fsnotify
github.com/ory/hydra-consent-app-go/vendor/github.com/ory-am/ladon
github.com/ory/hydra-consent-app-go/vendor/github.com/hashicorp/hcl
github.com/ory/hydra-consent-app-go/vendor/github.com/mitchellh/mapstructure
github.com/ory/hydra-consent-app-go/vendor/github.com/pelletier/go-buffruneio
github.com/ory/hydra-consent-app-go/vendor/github.com/spf13/afero/mem
github.com/ory/hydra-consent-app-go/vendor/github.com/pelletier/go-toml
github.com/ory/hydra-consent-app-go/vendor/golang.org/x/text/transform
github.com/ory/hydra-consent-app-go/vendor/github.com/spf13/cast
github.com/ory/hydra-consent-app-go/vendor/golang.org/x/text/unicode/norm
github.com/ory/hydra-consent-app-go/vendor/github.com/ory-am/hydra/pkg
github.com/ory/hydra-consent-app-go/vendor/github.com/spf13/jwalterweatherman
github.com/ory/hydra-consent-app-go/vendor/gopkg.in/yaml.v2
github.com/ory/hydra-consent-app-go/vendor/github.com/ory-am/hydra/client
github.com/ory/hydra-consent-app-go/vendor/github.com/ory-am/hydra/jwk
github.com/ory/hydra-consent-app-go/vendor/github.com/ory-am/hydra/policy
github.com/ory/hydra-consent-app-go/vendor/github.com/ory-am/hydra/warden/group
github.com/ory/hydra-consent-app-go/vendor/github.com/spf13/afero
github.com/ory/hydra-consent-app-go/vendor/github.com/ory/common/env
github.com/ory/hydra-consent-app-go/vendor/github.com/ory-am/hydra/oauth2
github.com/ory/hydra-consent-app-go/vendor/github.com/spf13/viper
github.com/ory/hydra-consent-app-go/vendor/github.com/ory-am/hydra/config
github.com/ory/hydra-consent-app-go/vendor/github.com/ory-am/hydra/warden
github.com/ory/hydra-consent-app-go/vendor/github.com/ory-am/hydra/sdk
github.com/ory/hydra-consent-app-go

fanhonglingdeMacBook-Pro:hydra-consent-app-go fanhongling$ hydra-consent-app-go 
```

Then open _http://localhost:4445_
## CI/CD HTTP Agent
The convenient agent for continuous delivery and deployment to your servers.

## Overview

* Http delivery channel, POST method
* Upload file up to MAX_SIZE
* Run a custom bash script after upload file
* Return stdout of bash script to CI/CD system
* Multiple projects per server
* Security is primarily provided by NGINX
* Allow to opt out Cloud Buckets services, Docker Registry, e.t.c


## Getting Started

* A CI/CD pipeline just run a script bellow after build JS assets, Docker image, e.t.c.
  
```sh
curl --location 'httpS://Login:${{ secrets.DEPLOY_PASS }}@YourDomain/YourNativeToken?project=YourProject' --form 'file=@"pathToFileOnBuildServer/deployArchiveFileName"'
```

*  A nginx configuration file can consist parts like this
  
  ```conf
  location = /YourNativeToken {
            auth_basic "Restricted area";
            auth_basic_user_file /etc/nginx/foo_bar_pass;
            proxy_pass http://127.0.0.1:9777;
            proxy_redirect http://127.0.0.1:9777 /;
            proxy_set_header Host $host;
            proxy_connect_timeout 1200;
 }

 listen 443 ssl;
 ssl_certificate /etc/path_key/fullchain.pem;
 ssl_certificate_key /etc/path_key/privkey.pem;

  ```
* The custom bash script file is stored in ```/etc/cicd_agent``` folder.
```sh
#!/bin/bash
# /etc/cicd_agent/YourProject.sh (must be executable)

cd workspace_directory
tar -xzf $1
cp foobar.file web_directory
# or
docker load image.file

***

```

## License

© Morph, 2024~time.Now

Released under the [MIT License](https://github.com/go-gorm/gorm/blob/master/LICENSE)

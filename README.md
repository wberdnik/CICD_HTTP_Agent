## CICD_HTTP_Agent
The convenient agent for continuous delivery and deployment to your servers.

## Overview

* Http delivery channel, POST method
* Upload file up to MAX_SIZE
* Run a custom bash script after upload file
* Return stdout of bash script to CI/CD system
* Multiple projects per server
* Security is primarily provided by NGINX 


## Getting Started

* A CI/CD pipeline just run
  
```
curl --location 'httpS://YourDomain/YourNativeToken?project=YourProject' --form 'file=@"pathToFile/deployArchiveFileName"'
```

* after build JS assets, Docker image, e.t.c. A nginx configuration can consist parts like
  
  ```
  location = /FooBarToken {
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



## License

© Morph, 2024~time.Now

Released under the [MIT License](https://github.com/go-gorm/gorm/blob/master/LICENSE)

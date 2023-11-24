########################################################################
## Nginx base
## Create environment
ARG NGINX_VERSION=1.24.0
FROM nginx:${NGINX_VERSION} as build

ARG MODSECURITY_VERSION=3.0.9
ENV MODSECURITY_VERSION=${MODSECURITY_VERSION}

ARG OWASP_CRS_VERSION=3.2.0
ENV OWASP_CRS_VERSION=${OWASP_CRS_VERSION}

ARG MODSECURITY_NGINX=1.0.3

ARG MODULE_HEADERS_MORE_NGINX_FILTER=0.34

########################################################################
## Install package for Ubuntu  
RUN apt-get update && \
    apt-get install -y --no-install-recommends \
    ## 
    git \
    curl \
    wget \ 
    gcc \
    make \
    zlib1g-dev \
    libpcre3-dev \
    ##
    openssh-client \
    libxml2 \
    libxslt1-dev \
    libpcre3 \
    zlib1g \
    openssl \
    libssl-dev \
    libtool \
    automake \
    g++ \
    libmaxminddb-dev \
    ##
    python3-pip \
    ##
    apache2-utils && \
    ln -fs /usr/share/zoneinfo/Asia/Bangkok /etc/localtime && \
    dpkg-reconfigure -f noninteractive tzdata && \
    rm -rf /var/lib/apt/lists/* && \
    rm -rf /var/cache/apt

RUN pip install gixy==0.1.20 pyparsing==2.4.7

########################################################################
## Install module
## Download : nginx
WORKDIR /usr/src
RUN wget "http://nginx.org/download/nginx-${NGINX_VERSION}.tar.gz" && \
    tar -xzvf nginx-${NGINX_VERSION}.tar.gz && \
    rm nginx-${NGINX_VERSION}.tar.gz
    # nginx-1.24.0.tar.gz
    # nginx-1.24.0
########################################################################
## Download module
## Module : Headers more nginx
WORKDIR /usr/src/nginx-${NGINX_VERSION}
RUN wget "https://github.com/openresty/headers-more-nginx-module/archive/refs/tags/v${MODULE_HEADERS_MORE_NGINX_FILTER}.tar.gz" && \
    tar -xzvf v${MODULE_HEADERS_MORE_NGINX_FILTER}.tar.gz
    # v0.34.tar.gz
    # headers-more-nginx-module-0.34


## Module : ModSecurity 
WORKDIR /usr/src/nginx-${NGINX_VERSION}
RUN wget "https://github.com/SpiderLabs/ModSecurity/releases/download/v${MODSECURITY_VERSION}/modsecurity-v${MODSECURITY_VERSION}.tar.gz" && \
    tar -xzvf modsecurity-v${MODSECURITY_VERSION}.tar.gz
    # modsecurity-v3.0.4.tar.gz
    # modsecurity-v3.0.4
WORKDIR /usr/src/nginx-${NGINX_VERSION}/modsecurity-v${MODSECURITY_VERSION}
RUN ./build.sh && \
    ./configure && \
    make && \
    make install 


## Module : ModSecurity nginx Connector
WORKDIR /usr/src/nginx-${NGINX_VERSION}
RUN wget "https://github.com/SpiderLabs/ModSecurity-nginx/releases/download/v${MODSECURITY_NGINX}/modsecurity-nginx-v${MODSECURITY_NGINX}.tar.gz" && \
    tar -xzvf modsecurity-nginx-v${MODSECURITY_NGINX}.tar.gz
    # modsecurity-nginx-v1.0.1.tar.gz
    # modsecurity-nginx-v1.0.1


## Module : ModSecurity for OWASP Core Rule Set 
WORKDIR /usr/src/nginx-${NGINX_VERSION}
RUN wget "https://github.com/SpiderLabs/owasp-modsecurity-crs/archive/v${OWASP_CRS_VERSION}.tar.gz" && \
    tar -xzvf v${OWASP_CRS_VERSION}.tar.gz
    # v3.2.0.tar.gz
    # owasp-modsecurity-crs-3.2.0


########################################################################
## Add and Compiler module
## Module : Headers more nginx
WORKDIR /usr/src/nginx-${NGINX_VERSION}
RUN NGINX_ARGS=$(nginx -V 2>&1 | sed -n -e 's/^.*arguments: //p') \
    ./configure  --with-compat --add-dynamic-module=/usr/src/nginx-${NGINX_VERSION}/headers-more-nginx-module-${MODULE_HEADERS_MORE_NGINX_FILTER} ${NGINX_ARGS} && \
    make modules 

## Module : ModSecurity nginx Connector
WORKDIR /usr/src/nginx-${NGINX_VERSION}
RUN NGINX_ARGS=$(nginx -V 2>&1 | sed -n -e 's/^.*arguments: //p') \
    ./configure --with-compat --with-http_dav_module --add-dynamic-module=/usr/src/nginx-${NGINX_VERSION}/modsecurity-nginx-v${MODSECURITY_NGINX} ${NGINX_ARGS} && \
    make modules
########################################################################


## Nginx 
## Create environment
ARG NGINX_VERSION=1.24.0
FROM nginx:${NGINX_VERSION} 

ARG MODSECURITY_VERSION=3.0.4
ENV MODSECURITY_VERSION=${MODSECURITY_VERSION}

ARG OWASP_CRS_VERSION=3.2.0
ENV OWASP_CRS_VERSION=${OWASP_CRS_VERSION}

ARG MODSECURITY_NGINX=1.0.1
########################################################################
## Create user for Nginx
RUN groupadd groupuserbe && \
    useradd -g groupuserbe userbe
RUN chown -R userbe:groupuserbe /var/cache/nginx 
RUN touch /var/run/nginx.pid && \
    chown -R userbe:groupuserbe /var/run/nginx.pid
########################################################################
## Logs Nginx
RUN rm  /var/log/nginx/access.log && \
    rm  /var/log/nginx/error.log
RUN touch /var/log/nginx/nginx_portainer.access.log && \
    touch /var/log/nginx/nginx_portainer.error.log && \
    touch /var/log/nginx/nginx_webservice_api.access.log && \
    touch /var/log/nginx/nginx_webservice_api.error.log 
RUN chmod 777 /var/log/nginx/nginx_portainer.access.log && \
    chmod 777 /var/log/nginx/nginx_portainer.error.log && \
    chmod 777 /var/log/nginx/nginx_webservice_api.access.log && \
    chmod 777 /var/log/nginx/nginx_webservice_api.error.log 

## Logs Modsecurity
RUN touch /var/log/modsec-debug.log && \
    touch /var/log/modsec-audit.log
RUN chmod 777 /var/log/modsec-debug.log && \
    chmod 777 /var/log/modsec-audit.log 
########################################################################
## Create file for network_internal
RUN touch /etc/nginx/network_internal.conf && \
    chmod 777  /etc/nginx/network_internal.conf
########################################################################
## Install package for Ubuntu  
RUN apt-get update && \
    apt-get install -y --no-install-recommends \
    libmaxminddb0 \
    iputils-ping \
    net-tools \
    git \
    curl \
    wget \ 
    vim \
    htop \
    fail2ban \
    iptables \
    procps \ 
    apache2-utils && \
    ln -fs /usr/share/zoneinfo/Asia/Bangkok /etc/localtime && \
    dpkg-reconfigure -f noninteractive tzdata && \
    rm -rf /var/lib/apt/lists/* && \
    rm -rf /var/cache/apt
########################################################################
## Copy folder from Nginx base to Nginx for Headers more nginx
COPY --from=build /usr/src/nginx-${NGINX_VERSION}/objs/ngx_http_headers_more_filter_module.so /usr/lib/nginx/modules/ngx_http_headers_more_filter_module.so
## Copy folder from Nginx base to Nginx for Modsecurity
RUN mkdir /etc/nginx/modsecurity
RUN touch /etc/nginx/modsecurity/main.conf && \
    touch /etc/nginx/modsecurity/modsecurity.conf && \
    touch /etc/nginx/modsecurity/unicode.mapping 
RUN chmod 777 /etc/nginx/modsecurity/main.conf && \
    chmod 777 /etc/nginx/modsecurity/modsecurity.conf && \
    chmod 777 /etc/nginx/modsecurity/unicode.mapping 
COPY ./nginx/WAF/modsecurity/main.conf /etc/nginx/modsecurity/main.conf 
COPY ./nginx/WAF/modsecurity/modsecurity.conf /etc/nginx/modsecurity/modsecurity.conf
COPY ./nginx/WAF/modsecurity/unicode.mapping /etc/nginx/modsecurity/unicode.mapping

RUN sed -i "s/_OWASP_CRS_VERSION_/${OWASP_CRS_VERSION}/g" /etc/nginx/modsecurity/main.conf
COPY --from=build /usr/src/nginx-${NGINX_VERSION}/objs/ngx_http_modsecurity_module.so /usr/lib/nginx/modules/ngx_http_modsecurity_module.so
COPY --from=build /usr/local/modsecurity/ /usr/local/modsecurity/
COPY --from=build /usr/src/nginx-${NGINX_VERSION}/owasp-modsecurity-crs-${OWASP_CRS_VERSION} /usr/local/owasp-modsecurity-crs-${OWASP_CRS_VERSION}
########################################################################
## Move folder Nginx for Modsecurity
RUN mv /usr/local/owasp-modsecurity-crs-${OWASP_CRS_VERSION}/crs-setup.conf.example /usr/local/owasp-modsecurity-crs-${OWASP_CRS_VERSION}/crs-setup.conf && \
    mv /usr/local/owasp-modsecurity-crs-${OWASP_CRS_VERSION}/rules/REQUEST-900-EXCLUSION-RULES-BEFORE-CRS.conf.example /usr/local/owasp-modsecurity-crs-${OWASP_CRS_VERSION}/rules/REQUEST-900-EXCLUSION-RULES-BEFORE-CRS.conf && \
    mv /usr/local/owasp-modsecurity-crs-${OWASP_CRS_VERSION}/rules/RESPONSE-999-EXCLUSION-RULES-AFTER-CRS.conf.example /usr/local/owasp-modsecurity-crs-${OWASP_CRS_VERSION}/rules/RESPONSE-999-EXCLUSION-RULES-AFTER-CRS.conf
########################################################################
WORKDIR /home
USER userbe
########################################################################



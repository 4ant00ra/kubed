package ingress

const haproxyTemplate = `# This file uses golang pongo2 templates to
# dynamically configure the haproxy loadbalancer.
global
    daemon
    stats socket /tmp/haproxy
    server-state-file global
    server-state-base /var/state/haproxy/
    maxconn 4000
    # log using a syslog socket
    log /dev/log local0 info
    log /dev/log local0 notice
    {% if SSLCert %}
    tune.ssl.default-dh-param 2048
    ssl-default-bind-ciphers ECDHE-RSA-AES128-GCM-SHA256:ECDHE-ECDSA-AES128-GCM-SHA256:ECDHE-RSA-AES256-GCM-SHA384:ECDHE-ECDSA-AES256-GCM-SHA384:DHE-RSA-AES128-GCM-SHA256:DHE-DSS-AES128-GCM-SHA256:kEDH+AESGCM:ECDHE-RSA-AES128-SHA256:ECDHE-ECDSA-AES128-SHA256:ECDHE-RSA-AES128-SHA:ECDHE-ECDSA-AES128-SHA:ECDHE-RSA-AES256-SHA384:ECDHE-ECDSA-AES256-SHA384:ECDHE-RSA-AES256-SHA:ECDHE-ECDSA-AES256-SHA:DHE-RSA-AES128-SHA256:DHE-RSA-AES128-SHA:DHE-DSS-AES128-SHA256:DHE-RSA-AES256-SHA256:DHE-DSS-AES256-SHA:DHE-RSA-AES256-SHA:!aNULL:!eNULL:!EXPORT:!DES:!RC4:!3DES:!MD5:!PSK
    {% endif %}

defaults
    log global

    option http-server-close

    # Disable logging of null connections (haproxy connections like checks).
    # This avoids excessive logs from haproxy internals.
    option dontlognull

    # Maximum time to wait for a connection attempt to a server to succeed.
    timeout connect         50000

    # Maximum inactivity time on the client side.
    # Applies when the client is expected to acknowledge or send data.
    timeout client          50000

    # Inactivity timeout on the client side for half-closed connections.
    # Applies when the client is expected to acknowledge or send data
    # while one direction is already shut down.
    timeout client-fin      50000

    # Maximum inactivity time on the server side.
    timeout server          50000

    # timeout to use with WebSocket and CONNECT
    timeout tunnel          50000

    # default traffic mode is http
    # mode is overwritten in case of tcp services
    mode http

    # errorloc 400 https://appscode.com/errors/400
    # errorloc 403 https://appscode.com/errors/403
    # errorloc 408 https://appscode.com/errors/408
    # errorloc 500 https://appscode.com/errors/500
    # errorloc 502 https://appscode.com/errors/502
    # errorloc 503 https://appscode.com/errors/503
    # errorloc 504 https://appscode.com/errors/504


{% if Stats %}
listen stats
    bind *:1936
    mode http
    stats enable
    stats realm Haproxy\ Statistics
    stats uri /
    {% if StatsUserName %}stats auth {{ StatsUserName }}:{{ StatsPassWord }}{% endif %}
{% endif %}

{% if DefaultBackend %}
# default backend
backend default-backend
    {% if Sticky %}cookie SERVERID insert indirect nocache{% endif %}
    {% for rule in DefaultBackend.RewriteRules %}
    reqrep {{ rule }}
    {% endfor %}
    {% for rule in DefaultBackend.HeaderRules %}
    acl h_x_{{ forloop.Counter }}_exists req.hdr({{ rule|header_name }}) -m found
    http-request add-header {{ rule }} unless h_x_{{ forloop.Counter }}_exists
    {% endfor %}
    {% for e in DefaultBackend.Endpoints %}
    server {{ e.Name }} {{ e.IP }}:{{ e.Port }} {% if Sticky %} cookie {{ e.Name }} {% endif %}
    {% endfor %}
{% endif %}

{% if HttpsService %}
# https service
frontend https-frontend
    bind *:443 ssl no-sslv3 no-tlsv10 crt /etc/ssl/private/haproxy/
    rspadd  Strict-Transport-Security:\ max-age=15768000

    mode http
    option httplog
    option forwardfor

{% for svc in HttpsService %}
    {% set both = 0 %}
    {% if svc.AclMatch %}acl url_acl_{{ svc.Name }} path_beg {{ svc.AclMatch }} {% set both = both + 1 %}{% endif %}
    {% if svc.Host %}acl host_acl_{{ svc.Name }} {{ svc.Host|host_name }} {% set both = both + 1 %}{% endif %}
    use_backend https-{{ svc.Name }} {% if both != 0 %}if {% endif %}{% if svc.AclMatch %}url_acl_{{ svc.Name }}{% endif %} {% if svc.Host %}host_acl_{{ svc.Name }}{% endif %}
{% endfor %}
    {% if DefaultBackend %}default_backend default-backend{% endif %}
{% endif %}

{% for svc in HttpsService %}
backend https-{{ svc.Name }}
    {% if Sticky %}cookie SERVERID insert indirect nocache{% endif %}
    {% for rule in svc.Backends.RewriteRules %}
    reqrep {{ rule }}
    {% endfor %}
    {% for rule in svc.Backends.HeaderRules %}
    acl h_x_{{ forloop.Counter }}_exists req.hdr({{ rule|header_name }}) -m found
    http-request add-header {{ rule }} unless h_x_{{ forloop.Counter }}_exists
    {% endfor %}
    {% for e in svc.Backends.Endpoints %}
    server {{ e.Name }} {{ e.IP }}:{{ e.Port }} {% if Sticky %} cookie {{ e.Name }} {% endif %}
    {% endfor %}
{% endfor %}

{% if HttpService %}
# http services.
frontend http-frontend
    bind *:80
    mode http

    mode http
    option httplog
    option forwardfor

{% for svc in HttpService %}
    {% set both = 0 %}
    {% if svc.AclMatch %}acl url_acl_{{ svc.Name }} path_beg {{ svc.AclMatch }} {% set both = both + 1 %}{% endif %}
    {% if svc.Host %}acl host_acl_{{ svc.Name }} {{ svc.Host|host_name }} {% set both = both + 1 %}{% endif %}
    use_backend http-{{ svc.Name }} {% if both != 0 %}if {% endif %}{% if svc.AclMatch %}url_acl_{{ svc.Name }}{% endif %} {% if svc.Host %}host_acl_{{ svc.Name }}{% endif %}
{% endfor %}
    {% if DefaultBackend %}default_backend default-backend{% endif %}
{% endif %}

{% for svc in HttpService %}
backend http-{{ svc.Name }}
    {% if Sticky %}cookie SERVERID insert indirect nocache{% endif %}
    {% for rule in svc.Backends.RewriteRules %}
    reqrep {{ rule }}
    {% endfor %}
    {% for rule in svc.Backends.HeaderRules %}
    acl h_xff_exists req.hdr({{ rule|header_name }}) -m found
    http-request add-header {{ rule }} unless h_xff_exists
    {% endfor %}
    {% for e in svc.Backends.Endpoints %}
    server {{ e.Name }} {{ e.IP }}:{{ e.Port }} {% if Sticky %}cookie {{ e.Name }} {% endif %}
    {% endfor %}
{% endfor %}


{% if TCPService %}
# tcp service
{% for svc in TCPService %}
frontend tcp-frontend-key-{{ svc.Port }}
    bind *:{{ svc.Port }} {% if svc.SecretName %}ssl no-sslv3 no-tlsv10 crt /etc/ssl/private/haproxy/{{ svc.SecretName }}.pem{% endif %}
    mode tcp

    default_backend tcp-{{ svc.Name }}
{% endfor %}
{% endif %}

{% for svc in TCPService %}
backend tcp-{{ svc.Name }}
    mode tcp
    {% if Sticky %}
    stick-table type ip size 100k expire 30m
    stick on src
    {% endif %}
    {% for e in svc.Backends.Endpoints %}
    server {{ e.Name }} {{ e.IP }}:{{ e.Port }}
    {% endfor %}
{% endfor %}

{% if !HttpService and !HttpsService and DefaultBackend %}
frontend http-frontend
    bind *:80
    mode http

    option forwardfor
    default_backend default-backend
{% endif %}`

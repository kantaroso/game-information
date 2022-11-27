vcl 4.1;

import std;

# 参考
# https://qiita.com/koudaiii/items/6a0efa024842cd48e2fb

backend default {
    .host = "go";
    .port = "80";
    .max_connections = 100;
    .probe = {
        # .url = "/health";
        .request =
            "GET /health HTTP/1.1"
            "Host: localhost"
            "Connection: close"
            "User-Agent: Varnish Health Probe";
        .interval  = 10s;
        .timeout   = 5s;
        .window    = 5;
        .threshold = 3;
    }
    .connect_timeout        = 5s;
    .first_byte_timeout     = 90s;
    .between_bytes_timeout  = 2s;
}

sub vcl_init {
    std.log("vcl_init");
}

sub vcl_recv {
    std.log("vcl_recv");
    # 固定レスポンステスト
    if (req.url ~ "^/test_v") {
        return (synth(200));
    }
    # キャッシュしないパス
    if (req.url ~ "^/common/pv_nocache") {
        std.log("no cache");
        return (pass);
    }
    # パージ
    if (req.method == "PURGE") {
        ban("req.http.host == " + req.http.host +
            " && req.url ~ " + req.url);
        return(synth(200, "PURGE accepted"));
    }

    # cookieが設定されているとキャッシュしないのでキャッシュさせたい場合は外す必要あり
    if (req.url ~ "^/*") {
        unset req.http.Cookie;
    }
}

# キャッシュキーの設定
sub vcl_hash {
    hash_data(req.url);
    if (req.http.host) {
        hash_data(req.http.host);
    }
    hash_data(req.http.User-Agent);
}

# vclのエラーページのカスタマイズ
sub vcl_synth {
    std.log("vcl_synth");
    set resp.http.Content-Type = "text/html; charset=utf-8";
    synthetic({"<!DOCTYPE html><html><body>"} + resp.reason + {"</body></html>"});
    return (deliver);
}

sub vcl_backend_response {
    # キャッシュしないパス
    if (bereq.url ~ "^/common/pv_cachepurge") {
        set beresp.ttl = 1s;
    }
}

sub vcl_deliver {
    set resp.http.cache-control = "max-age=300";
}

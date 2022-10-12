Benchmark Test is as below On mac M1
Test Utility : https://github.com/grafana/k6


          /\      |‾‾| /‾‾/   /‾‾/
     /\  /  \     |  |/  /   /  /
    /  \/    \    |     (   /   ‾‾\
   /          \   |  |\  \ |  (‾)  |
  / __________ \  |__| \__\ \_____/ .io

  execution: local
     script: get-http-ping-local.js
     output: -

  scenarios: (100.00%) 1 scenario, 80 max VUs, 1m0s max duration (incl. graceful stop):
           * default: 80 looping VUs for 30s (gracefulStop: 30s)


running (0m30.0s), 00/80 VUs, 1721532 complete and 0 interrupted iterations
default ✓ [======================================] 80 VUs  30s

     data_received..................: 210 MB  7.0 MB/s
     data_sent......................: 145 MB  4.8 MB/s
     http_req_blocked...............: avg=775ns  min=0s      med=0s       max=7.83ms   p(90)=1µs    p(95)=1µs
     http_req_connecting............: avg=165ns  min=0s      med=0s       max=6.26ms   p(90)=0s     p(95)=0s
     http_req_duration..............: avg=1.37ms min=18µs    med=636µs    max=113.86ms p(90)=2.37ms p(95)=4.37ms
       { expected_response:true }...: avg=1.37ms min=18µs    med=636µs    max=113.86ms p(90)=2.37ms p(95)=4.37ms
     http_req_failed................: 0.00%   ✓ 0            ✗ 1721532
     http_req_receiving.............: avg=5.95µs min=2µs     med=4µs      max=12.5ms   p(90)=8µs    p(95)=10µs
     http_req_sending...............: avg=2.44µs min=1µs     med=2µs      max=8.9ms    p(90)=3µs    p(95)=6µs
     http_req_tls_handshaking.......: avg=0s     min=0s      med=0s       max=0s       p(90)=0s     p(95)=0s
     http_req_waiting...............: avg=1.37ms min=13µs    med=629µs    max=113.85ms p(90)=2.35ms p(95)=4.35ms
     http_reqs......................: 1721532 57376.493519/s
     iteration_duration.............: avg=1.39ms min=24.58µs med=646.87µs max=113.87ms p(90)=2.38ms p(95)=4.38ms
     iterations.....................: 1721532 57376.493519/s
     vus............................: 80      min=80         max=80
     vus_max........................: 80      min=80         max=80

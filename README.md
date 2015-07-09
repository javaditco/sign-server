aftasardemhemorroidasidem

to test it:


````curl -F "file1=@test.txt" -X PUT http://127.0.0.1:8080/upload````

where test.txt is a local file


build docker image


run docker

````sudo docker run  --dns 172.17.42.1 --name sign-server -i -t -v $PWD:/sign-server vpereira/golang````

update dnsmasq config:

````sudo ./update-dns.sh````

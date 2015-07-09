Tarball signer server

to test it:


````curl -F "file1=@test.tgz" -X PUT http://127.0.0.1:8080/upload````

or if you are running the dns config below:

````curl -F "file1=@test.tgz" -X PUT http://sign-server:8080/upload````

you will get a response like:

````
{
    "bytes": 120342,
    "id": "2519adc0-261b-11e5-a051-0242ac110052",
    "name": "test.tgz",
    "path": "/uploaded/test.tgz",
    "sha256": "f6f24a11d7cbbbc6d9440aca2eba0f6498755ca90adea14c5e233bf4c04bd928",
    "url": "http://sign-server:8080/file/2519adc0-261b-11e5-a051-0242ac110052"
}


````

curling then http://sign-server:8080/file/2519adc0-261b-11e5-a051-0242ac110052

you get the clearsigned sha256 from the package:

````
-----BEGIN PGP SIGNED MESSAGE-----
Hash: SHA256

f6f24a11d7cbbbc6d9440aca2eba0f6498755ca90adea14c5e233bf4c04bd928
-----BEGIN PGP SIGNATURE-----

wl4EAREIABAFAlWeO90JEDzLZuJFzu0IAADYGgEAkCynaxOTwhrvpZEm4YoAlcx6
T9937KjxRJI9ILlNJscBAMPU4/33jBuGyEGt0cGejaofiA3e3CYP+4HKWUDuGpJO
=Ab5t
-----END PGP SIGNATURE-----vpereira@sagres:~>

````


build docker image


run docker

````sudo docker run  --dns 172.17.42.1 --name sign-server -i -t -v $PWD:/sign-server vpereira/golang````

update dnsmasq config:

````sudo ./update-dns.sh````


TODO:

* configuration file
* support not just for upload, but file download

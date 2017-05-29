swig -go -cgo -c++ -intgosize 64 db.i; 
CGO_LDFLAGS='-ldb'  go  install -x
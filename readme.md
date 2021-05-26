начать стоит с протокол buffers, ктрый будет описывать внешний интерфейс сервиса
api/proto/adder.proto

скомпилировать go пакет
protoc -I api/proto --go_out=plugins=grpc:./pkg/api api/proto/adder.proto
утилита protoc, ктрая компилирует из proto файлов, файлы на требуемом языке
```--go_out=plugins=grpc:``` - использовать плагин grpc
./pkg/api - папка куда скомпилировать код
api/proto/adder.proto - указание proto файла

будет сгенерирован adder.pb.go, в нем не надо делать изменения
создам структуру pkg/adder/grpcserver.go, которая будет реализовывать интерфейс AdderServer
требуется описать метод Add

создам main файл, ктрый будет запускать сервер и создавать grpc слушателя

скомпилирую проект
go build -v ./cmd/server
пробую запустить
./server

во втором терминале по тому же пути, использую клиент evans пробую подключиться к grpc серверу
указываю путь к описывающему файлу схеме сервиса  
evans api/proto/adder.proto -p 8080

вызываю метод
call Add

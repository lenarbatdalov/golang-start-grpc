<h1>Что за суета вообще происходит, опишу))</h1>

<p>начать стоит с протокол buffers, ктрый будет описывать внешний интерфейс сервиса</br>
api/proto/adder.proto</p>

<p>protoc -I api/proto --go_out=plugins=grpc:./pkg/api api/proto/adder.proto</br>
утилита protoc, ктрая компилирует из proto файлов, файлы на требуемом языке</br>
```--go_out=plugins=grpc:``` - использовать плагин grpc</br>
./pkg/api - папка куда скомпилировать код</br>
api/proto/adder.proto - указание proto файла</p>

<p>будет сгенерирован adder.pb.go, в нем не надо делать изменения</br>
создам структуру pkg/adder/grpcserver.go, которая будет реализовывать интерфейс AdderServer</br>
требуется описать метод Add</p>

<p>создам main файл, ктрый будет запускать сервер и создавать grpc слушателя</p>

<p>скомпилирую проект</br>
go build -v ./cmd/server</br>
пробую запустить</br>
./server</p>

<p>во втором терминале по тому же пути, использую клиент evans пробую подключиться к grpc серверу</br>
указываю путь к описывающему файлу схеме сервиса</br>
evans api/proto/adder.proto -p 8080</p>

<p>вызываю метод</br>
call Add</p>

<p></p>

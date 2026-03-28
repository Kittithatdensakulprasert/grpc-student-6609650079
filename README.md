# grpc-student-6609650079


เมื่อแก้ไฟล์ .proto ต้องสั่ง Gen โค้ดใหม่เสมอ : protoc --go_out=. --go-grpc_out=. proto/student.proto

เปิด รัน Server:
go run server/server.go

เปิด รัน Client:
go run client/client.go


ไฟล์ .proto (Protocol Buffers): เปรียบเหมือน "เมนู"

ฝั่ง Server (server.go): เปรียบเหมือน "ห้องทำงาน" มีหน้าที่รอรับคำสั่ง พอคำสั่งมาถึงก็ไปประมวลผล

ฝั่ง Client (client.go): เปรียบเหมือน "คนเรียก" มีหน้าที่ดูเมนู (.proto) แล้วส่งคำสั่งไป

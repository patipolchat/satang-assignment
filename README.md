1 ใช้ make task เปิด service mongodb ด้วย `make db-create`

2 รัน go run main.go

*Note

1 ใช้ go-ethereum ที่ implement json rpc แทนการใช้ json rpc โดยตรง โดยมีปัญหาดังนี้

    - ข้อมูล transaction ที่ได้รับมาไม่ครบเมื่อเทียบกับการใช้ json rpc ติดต่อไปโดยตรง
    - Bug แปลง type common.address to string หากเป็น nil จะค้าง
    แก้ไขโดยทำการใช้ json rpc โดยตรง โดยใช้ websocket และ subscribe newHead {"jsonrpc":"2.0", "id": 1, "method": "eth_subscribe", "params": ["newHeads"]}
    
2 ใช้ INFURA เพื่อติดต่อไปยัง ethereum network โดยจะมี api key ติดไปด้วย จะ revoke ทิ้งในภายหล้ง
